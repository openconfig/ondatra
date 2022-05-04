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
package clientwrapper

import (
	"bytes"
	"golang.org/x/net/context"
	"strings"
	"testing"
	"time"

	"github.com/openconfig/ondatra/fakebind"
)

type action struct {
	wait   time.Duration
	stdout []byte
	stderr []byte
	err    bool
}

func TestClientWrapper(t *testing.T) {
	overBuffer := strings.Repeat("a", 8000)
	tests := []struct {
		desc    string
		actions []action
	}{{
		desc: "simple writes",
		actions: []action{
			{stdout: []byte("1 multiple writes\n")},
			{stdout: []byte("2 multiple writes\n")},
			{stdout: []byte("3 multiple writes\n")},
			{stdout: []byte("4 multiple writes\n")},
			{stdout: []byte("5 multiple writes\n")},
			{stdout: []byte("6 multiple writes\n")},
			{stdout: []byte("7 multiple writes\n")},
			{stdout: []byte("8 multiple writes\n")},
		},
	}, {
		desc: "simple err and out",
		actions: []action{
			{stdout: []byte("1 multiple writes\n")},
			{stderr: []byte("2 multiple writes\n")},
			{stdout: []byte("3 multiple writes\n")},
			{stdout: []byte("4 multiple writes\n")},
			{stderr: []byte("5 multiple writes\n")},
			{stdout: []byte("6 multiple writes\n")},
			{stdout: []byte("7 multiple writes\n")},
			{stderr: []byte("8 multiple writes\n")},
		},
	}, {
		desc: "simple err and out with err",
		actions: []action{
			{stdout: []byte("1 multiple writes\n")},
			{stderr: []byte("2 multiple writes\n")},
			{stdout: []byte("3 multiple writes\n")},
			{stdout: []byte("4 multiple writes\n")},
			{stderr: []byte("5 multiple writes\n")},
			{stdout: []byte("6 multiple writes\n")},
			{stdout: []byte("7 multiple writes\n")},
			{err: true},
		},
	}, {
		desc: "oversize buffer",
		actions: []action{
			{stdout: []byte(overBuffer)},
			{stderr: []byte(overBuffer)},
			{stdout: []byte(overBuffer)},
			{stdout: []byte(overBuffer)},
			{stderr: []byte(overBuffer)},
			{stdout: []byte(overBuffer)},
			{stdout: []byte(overBuffer)},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			b := bytes.NewBuffer([]byte{})
			sc := fakebind.NewStreamClient()
			ctx := context.Background()
			l := Start(ctx, sc, b)
			bLen := 0
			for _, a := range tt.actions {
				time.Sleep(a.wait)
				switch {
				case a.err:
					sc.OutWriter.Close()
				case len(a.stdout) != 0:
					bLen += len(a.stdout) + 9 // prefix
					sc.OutWriter.Write(a.stdout)
				case len(a.stderr) != 0:
					bLen += len(a.stdout) + 8 // prefix
					sc.ErrWriter.Write(a.stderr)
				}
			}
			l.Stop()
			if bLen > len(b.Bytes()) {
				t.Fatalf("Unexpected buffer output length: got %d, want %d", len(b.Bytes()), bLen)
			}
		})
	}
}
