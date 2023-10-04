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

// Package console provides an API to interact with the serial console of a DUT.
//
// To capture the contents of the serial console to a file, use the
// [Console.StartCapture] method:
//
//	stdoutFile, err := os.Create("stdout.log")
//	if err != nil {
//		return err
//	}
//	stderrFile, err := os.Create("stderr.log")
//	if err != nil {
//		return err
//	}
//	dut.Console().StartCapture(t, stdoutFile, stderrFile)
//
// StartCapture automatically registers a test cleanup function to stop the
// capture, but if a test needs fine-grained control over when precisely it is
// stopped, it can call the [StopCaptureFunc] function returned by StartCapture:
//
//	stopFn := dut.Console().StartCapture(t, stdoutFile, stderrFile)
//	doStuffToTriggerConsoleOutput()
//	stopFn(t)
package console

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"testing"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/events"
)

// New constructs a new instance of the Console API.
// Tests must not call this directly
func New(dut binding.DUT) *Console {
	return &Console{dut: dut}
}

// Console is the device Console API.
type Console struct {
	dut binding.DUT
}

// StartCapture starts copying console stdout and stderr to outw and errw,
// respectively. It registers a cleanup function to stop the capture but also
// returns that same function, in case the user wants to stop it earlier.
func (c *Console) StartCapture(t testing.TB, outw, errw io.Writer) StopCaptureFunc {
	t.Helper()
	t = events.ActionStarted(t, "Starting console capture on %s", c.dut)
	cap, err := c.startCapture(context.Background(), outw, errw)
	if err != nil {
		t.Fatalf("StartCapture(t, out, err) on %s: %v", c.dut, err)
	}
	t.Cleanup(func() {
		if err := cap.stop(); err != nil {
			log.Errorf("StopCapture(t) on %s: %v", c.dut, err)
		}
	})
	return func(t testing.TB) {
		t.Helper()
		if err := cap.stop(); err != nil {
			t.Fatalf("StopCapture(t) on %s: %v", c.dut, err)
		}
	}
}

// StopCaptureFunc is a function that stops copying console stdout and stderr
type StopCaptureFunc func(t testing.TB)

func (c *Console) startCapture(ctx context.Context, outw, errw io.Writer) (*capturer, error) {
	client, err := c.dut.DialConsole(ctx)
	if err != nil {
		return nil, err
	}
	cap := &capturer{client: client}
	cap.wg.Add(1)
	go cap.writeStream(client.Stdout(), outw)
	cap.wg.Add(1)
	go cap.writeStream(client.Stderr(), errw)
	return cap, nil
}

type capturer struct {
	client  binding.ConsoleClient
	wg      sync.WaitGroup
	stopped atomic.Bool
}

func (c *capturer) readAndWrite(in io.Reader, buf []byte, out io.Writer) error {
	n, err := in.Read(buf)
	if err != nil && err != io.EOF {
		return fmt.Errorf("console capture got error reading: %w", err)
	}
	buf = buf[:n]
	if len(buf) > 0 {
		if _, err := out.Write(buf); err != nil {
			return fmt.Errorf("console capture got error writing: %w", err)
		}
	}
	return err
}

func (c *capturer) writeStream(in io.Reader, out io.Writer) {
	defer c.wg.Done()
	buf := make([]byte, 4096)
	for !c.stopped.Load() {
		if err := c.readAndWrite(in, buf, out); err != nil {
			if err != io.EOF {
				log.Error(err)
			}
			return
		}
	}
}

func (c *capturer) stop() error {
	if c.stopped.Swap(true) {
		return nil
	}
	if err := c.client.Close(); err != nil {
		return fmt.Errorf("failed to close the ConsoleClient: %w", err)
	}
	c.wg.Wait()
	return nil
}
