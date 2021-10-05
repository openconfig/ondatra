// Package clientwrapper provides a wrapper for writing output from a
// StreamClient to io.Writer.
package clientwrapper

import (
	"golang.org/x/net/context"
	"fmt"
	"io"
	"sync"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra"
)

// IOWrappedClient wraps the provided StreamClient and sends the output of
// Stdout and Stderr to the provided io.Writer.
type IOWrappedClient struct {
	ctx    context.Context
	cancel func()
	sc     ondatra.StreamClient
	mu     sync.Mutex
	closed bool
	out    io.Writer
	wg     sync.WaitGroup
}

// Start creates a new log wrapped StreamClient for recording output from the
// client to an io.Writer. The caller is responsible for calling Close() on the
// Wrapper to properly shutdown the wrapper.
func Start(ctx context.Context, sc ondatra.StreamClient, o io.Writer) *IOWrappedClient {
	if err := sc.Stdin().Close(); err != nil {
		log.Warningf("clientwrapper failed to close stdin: %v", err)
	}
	c := &IOWrappedClient{
		sc:  sc,
		out: o,
	}
	c.start(ctx)
	return c
}

func (c *IOWrappedClient) write(prefix string, buf []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(buf) == 0 {
		return nil
	}
	l := len(buf) - 1
	if buf[l] == '\n' {
		buf = buf[:l]
	}
	_, err := fmt.Fprintf(c.out, "[%s] %s\n", prefix, string(buf))
	return err
}

func (c *IOWrappedClient) writeOut(prefix string, in io.Reader) {
	buf := make([]byte, 4096)
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}
		n, err := in.Read(buf)
		switch {
		default:
			log.Errorf("console unknown error: %v", err)
		case err == nil:
			if err := c.write(prefix, buf[:n]); err != nil {
				log.Errorf("console write failed for message: %v", err)
			}
		case err == io.EOF:
			c.mu.Lock()
			defer c.mu.Unlock()
			if !c.closed {
				if err := c.sc.Close(); err != nil {
					log.Warningf("Failed to close streamclient: %v", err)
				}
				c.closed = true
			}
			if n > 0 {
				c.write(prefix, buf[:n])
			}
			return
		}
	}
}

func (c *IOWrappedClient) start(ctx context.Context) {
	c.ctx, c.cancel = context.WithCancel(ctx)
	// Start Stdout reader.
	c.wg.Add(1)
	go func() {
		c.writeOut("output", c.sc.Stdout())
		c.wg.Done()
	}()
	// Start Stderr reader.
	c.wg.Add(1)
	go func() {
		c.writeOut("error", c.sc.Stderr())
		c.wg.Done()
	}()
}

// Stop will stop the wrapped logging. The underlying io.Writer will still need
// to be closed by the user.
func (c *IOWrappedClient) Stop() {
	c.cancel()
	if err := c.sc.Close(); err != nil {
		log.Warningf("Failed to close streamclient: %v", err)
	}
	c.wg.Wait()
}
