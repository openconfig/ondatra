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

// Package cli provides an API to send CLI commands to a DUT.
package cli

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/events"
)

// New constructs a new instance of the CLI API.
// Tests must not call this directly
func New(dut binding.DUT) *CLI {
	return &CLI{dut: dut}
}

// CLI is the device CLI API.
type CLI struct {
	dut binding.DUT
}

// Run runs the specified CLI command on the DUT and returns its output.
// Run fails fatally if either (a) the command runs and reports an error, or
// (b) an error occurs that prevents the command from being run at all.
// To capture the error from case (a) instead, use [CLI.RunResult].
func (c *CLI) Run(t testing.TB, cmd string) string {
	t.Helper()
	t = events.ActionStarted(t, "Running CLI command on %s", c.dut)
	res, err := c.run(context.Background(), cmd)
	if err != nil {
		t.Fatalf("Run(t, %q) on %s: %v", cmd, c.dut.Name(), err)
	}
	if res.Error() != "" {
		t.Fatalf("Run(t, %q) on %s: %v", cmd, c.dut.Name(), res.Error())
	}
	return res.Output()
}

// RunResult runs the specified CLI command on the DUT and returns its result.
// RunResult fails fatally if an error occurs that prevents the command from
// being run. If the command runs and reports an error, that error is available
// in the result. To fail fatally in the latter case instead, use [CLI.Run].
func (c *CLI) RunResult(t testing.TB, cmd string) binding.CommandResult {
	t.Helper()
	t = events.ActionStarted(t, "Running CLI command on %s", c.dut)
	res, err := c.run(context.Background(), cmd)
	if err != nil {
		t.Fatalf("RunResult(t, %q) on %s: %v", cmd, c.dut.Name(), err)
	}
	return res
}

func (c *CLI) run(ctx context.Context, cmd string) (binding.CommandResult, error) {
	cli, err := c.dut.DialCLI(ctx)
	if err != nil {
		return nil, err
	}
	return cli.RunCommand(ctx, cmd)
}
