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
	"io"

	"github.com/openconfig/ondatra/binding"
)

// ConsoleClient is a fake implementation of ConsoleClient for use in testing.
type ConsoleClient struct {
	*binding.AbstractConsoleClient
	stdin     *io.PipeWriter
	stdout    *io.PipeReader
	stderr    *io.PipeReader
	InReader  *io.PipeReader
	OutWriter *io.PipeWriter
	ErrWriter *io.PipeWriter
	cErr      error
}

// NewConsoleClient returns a new ConsoleClient.
// Tests should access the stub side of the pipes via exported fields.
func NewConsoleClient() *ConsoleClient {
	inReader, inWriter := io.Pipe()
	outReader, outWriter := io.Pipe()
	errReader, errWriter := io.Pipe()
	return &ConsoleClient{
		InReader:  inReader,
		OutWriter: outWriter,
		ErrWriter: errWriter,
		stdin:     inWriter,
		stdout:    outReader,
		stderr:    errReader,
	}
}

// Close will return a configured err.
func (f *ConsoleClient) Close() error {
	f.InReader.Close()
	f.OutWriter.Close()
	f.ErrWriter.Close()
	return f.cErr
}

// Stdin returns the configured stdin buffer.
func (f *ConsoleClient) Stdin() io.WriteCloser {
	return f.stdin
}

// Stdout returns the configured stdout buffer.
func (f *ConsoleClient) Stdout() io.ReadCloser {
	return f.stdout
}

// Stderr returns the configured stderr buffer.
func (f *ConsoleClient) Stderr() io.ReadCloser {
	return f.stderr
}
