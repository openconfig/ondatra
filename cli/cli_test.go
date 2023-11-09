// Copyright 2023 Google LLC
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

package cli

import (
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"
)

var (
	fakeClient = &fakeCLIClient{}
	dut        = &fakebind.DUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{Name: "fakeDUT"}},
		DialCLIFn: func(context.Context) (binding.CLIClient, error) {
			return fakeClient, nil
		},
	}
)

func TestRun(t *testing.T) {
	cli := New(dut)

	t.Run("pass", func(t *testing.T) {
		want := "fake output"
		fakeClient.result = &fakeCommandResult{output: want}
		got := cli.Run(t, "cmd")
		if got != want {
			t.Errorf("Run() got %q, want %q", got, want)
		}
	})

	t.Run("fail", func(t *testing.T) {
		wantErr := "fake error"
		fakeClient.result = &fakeCommandResult{output: "output", error: wantErr}
		gotErr := testt.ExpectFatal(t, func(t testing.TB) {
			cli.Run(t, "cmd")
		})
		if !strings.Contains(gotErr, wantErr) {
			t.Errorf("Run() got error %q, want %q", gotErr, wantErr)
		}
	})
}

func TestRunResult(t *testing.T) {
	cli := New(dut)

	t.Run("pass", func(t *testing.T) {
		want := &fakeCommandResult{output: "fake output"}
		fakeClient.result = want
		got := cli.RunResult(t, "cmd")
		if got != want {
			t.Errorf("RunResult() got %q, want %q", got, want)
		}
	})

	t.Run("error", func(t *testing.T) {
		want := &fakeCommandResult{output: "fake output", error: "fake error"}
		fakeClient.result = want
		got := cli.RunResult(t, "cmd")
		if got != want {
			t.Errorf("RunResult() got %q, want %q", got, want)
		}
	})
}

type fakeCLIClient struct {
	binding.CLIClient
	result binding.CommandResult
}

func (c *fakeCLIClient) RunCommand(ctx context.Context, cmd string) (binding.CommandResult, error) {
	return c.result, nil
}

type fakeCommandResult struct {
	*binding.AbstractCommandResult
	output, error string
}

func (r *fakeCommandResult) Output() string {
	return r.output
}

func (r *fakeCommandResult) Error() string {
	return r.error
}
