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

package console

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
)

var (
	fakeClient *fakebind.ConsoleClient
	dut        = &fakebind.DUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{Name: "fakeDUT"}},
		DialConsoleFn: func(context.Context) (binding.ConsoleClient, error) {
			return fakeClient, nil
		},
	}
)

type action struct {
	stdout, stderr []byte
	outErr, errErr bool
}

func TestConsoleCapture(t *testing.T) {
	overBuffer := strings.Repeat("a", 8000)
	tests := []struct {
		desc    string
		actions []action
		wantOut string
		wantErr string
	}{{
		desc: "simple writes",
		actions: []action{
			{stdout: []byte("write out 1\n")},
			{stderr: []byte("write err 1\n")},
			{stderr: []byte("write err 2\n")},
			{stdout: []byte("write out 2\n")},
		},
		wantOut: "write out 1\nwrite out 2\n",
		wantErr: "write err 1\nwrite err 2\n",
	}, {
		desc: "write errors",
		actions: []action{
			{stdout: []byte("write out 1\n")},
			{stderr: []byte("write err 1\n")},
			{outErr: true},
			{errErr: true},
			{stderr: []byte("write err 2\n")},
			{stdout: []byte("write out 2\n")},
		},
		wantOut: "write out 1\n",
		wantErr: "write err 1\n",
	}, {
		desc: "oversize buffer",
		actions: []action{
			{stdout: []byte(overBuffer)},
			{stderr: []byte(overBuffer)},
			{stdout: []byte(overBuffer)},
			{stderr: []byte(overBuffer)},
		},
		wantOut: overBuffer + overBuffer,
		wantErr: overBuffer + overBuffer,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			console := New(dut)
			fakeClient = fakebind.NewConsoleClient()
			outBuf := bytes.NewBuffer([]byte{})
			errBuf := bytes.NewBuffer([]byte{})
			stopCap := console.StartCapture(t, outBuf, errBuf)

			for _, a := range test.actions {
				switch {
				case a.outErr:
					fakeClient.OutWriter.Close()
				case a.errErr:
					fakeClient.ErrWriter.Close()
				case len(a.stdout) > 0:
					fakeClient.OutWriter.Write(a.stdout)
				case len(a.stderr) > 0:
					fakeClient.ErrWriter.Write(a.stderr)
				}
			}
			stopCap(t)

			if gotOut := string(outBuf.Bytes()); test.wantOut != gotOut {
				t.Fatalf("Unexpected output buffer: got %s, want %s", gotOut, test.wantOut)
			}
			if gotErr := string(errBuf.Bytes()); test.wantErr != gotErr {
				t.Fatalf("Unexpected error buffer: got %s, want %s", gotErr, test.wantErr)
			}
		})
	}
}
