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

package fakebind

import (
	"fmt"
	"io"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
)

// StreamClient is a fake implementation of StreamClient for use in testing.
type StreamClient struct {
	*binding.AbstractStreamClient
	stdin     *io.PipeWriter
	stdout    *io.PipeReader
	stderr    *io.PipeReader
	InReader  *io.PipeReader
	OutWriter *io.PipeWriter
	ErrWriter *io.PipeWriter
	cErr      error
}

// NewStreamClient returns a new StreamClient. Users should just access the test
// sides of the pipes for testing via exported fields.
func NewStreamClient() *StreamClient {
	inReader, inWriter := io.Pipe()
	outReader, outWriter := io.Pipe()
	errReader, errWriter := io.Pipe()
	return &StreamClient{
		InReader:  inReader,
		OutWriter: outWriter,
		ErrWriter: errWriter,
		stdin:     inWriter,
		stdout:    outReader,
		stderr:    errReader,
	}
}

// SendCommand will echo whatever command is sent. If the cmd is "error" an
// error till be returned.
func (f *StreamClient) SendCommand(_ context.Context, cmd string) (string, error) {
	if cmd == "error" {
		return "", fmt.Errorf("error")
	}
	return cmd, nil
}

// Close will return a configured err.
func (f *StreamClient) Close() error {
	f.InReader.Close()
	f.OutWriter.Close()
	f.ErrWriter.Close()
	return f.cErr
}

// Stdin returns the configured stdin buffer.
func (f *StreamClient) Stdin() io.WriteCloser {
	return f.stdin
}

// Stdout returns the configured stdout buffer.
func (f *StreamClient) Stdout() io.ReadCloser {
	return f.stdout
}

// Stderr returns the configured stderr buffer.
func (f *StreamClient) Stderr() io.ReadCloser {
	return f.stderr
}
