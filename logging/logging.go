// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package logging utilities for logging from Ondatra tests.
package logging

import (
	"fmt"
	"io"
	"sync"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/raw"
)

// StartStreamLogger creates and starts a new StreamLogger for recording output
// from a StreamClient to an io.Writer. The caller is responsible for calling
// Close() on the StreamLogger to properly shutdown the wrapper.
func StartStreamLogger(sc raw.StreamClient, out io.Writer) *StreamLogger {
	sl := &StreamLogger{sc: sc, out: out}
	sl.start()
	return sl
}

// StreamLogger wraps a provided StreamClient and sends the output of
// Stdout and Stderr to a provided io.Writer.
type StreamLogger struct {
	sc      raw.StreamClient
	mu      sync.Mutex
	closed  bool
	stopped bool
	out     io.Writer
	wg      sync.WaitGroup
}

func (sl *StreamLogger) start() {
	// Start stdout reader.
	sl.wg.Add(1)
	go sl.writeStream("output", sl.sc.Stdout())
	// Start stderr reader.
	sl.wg.Add(1)
	go sl.writeStream("error", sl.sc.Stderr())
}

func (sl *StreamLogger) writeStream(prefix string, in io.Reader) {
	defer sl.wg.Done()
	buf := make([]byte, 4096)
	for !sl.isStopped() {
		if err := sl.readAndWrite(prefix, in, buf); err != nil {
			if err == io.EOF {
				sl.close()
			} else {
				log.Error(err)
			}
			return
		}
	}
}

func (sl *StreamLogger) isStopped() bool {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	return sl.stopped
}

func (sl *StreamLogger) readAndWrite(prefix string, in io.Reader, buf []byte) error {
	n, err := in.Read(buf)
	if err != nil && err != io.EOF {
		return fmt.Errorf("StreamLogger got error reading: %w", err)
	}
	buf = buf[:n]
	if len(buf) > 0 {
		l := len(buf) - 1
		if buf[l] == '\n' {
			buf = buf[:l]
		}
		if err := sl.write(prefix, buf); err != nil {
			return err
		}
	}
	return err
}

func (sl *StreamLogger) write(prefix string, buf []byte) error {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	if _, err := fmt.Fprintf(sl.out, "[%s] %s\n", prefix, string(buf)); err != nil {
		return fmt.Errorf("StreamLogger got error writing: %w", err)
	}
	return nil
}

// Stop will stop the logging.
// The provided io.Writer will still need to be closed by the user.
func (sl *StreamLogger) Stop() {
	func() {
		sl.mu.Lock()
		defer sl.mu.Unlock()
		sl.stopped = true
	}()
	sl.close()
	sl.wg.Wait()
}

func (sl *StreamLogger) close() {
	sl.mu.Lock()
	defer sl.mu.Unlock()
	if !sl.closed {
		if err := sl.sc.Close(); err != nil {
			log.Errorf("Failed to close StreamClient: %v", err)
		}
		sl.closed = true
	}
}
