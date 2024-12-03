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

package integration_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/openconfig/ondatra"
	dninit "github.com/openconfig/ondatra/dnbind/init"
)

func TestMain(m *testing.M) {
	ondatra.RunTests(m, dninit.Init)
}

func TestDrivenetsConfig(t *testing.T) {
	dut := ondatra.DUT(t, "dut")

	dut.Config().New().
		WithDrivenetsText(
			`interfaces
               ge100-0/0/0.321
                 admin-state enabled
                 vlan-id 321
               !
             !`).
		Append(t)

	// check duplicate commit
	dut.Config().New().
		WithDrivenetsText(
			`interfaces
               ge100-0/0/0.321
                 admin-state enabled
                 vlan-id 321
               !
             !`).
		Append(t)

	// updates DUT config with replace config below
	dut.Config().New().
		WithDrivenetsText(
			`interfaces
               ge100-0/0/0.{{ var "vlan" }}
                 admin-state {{ var "state" }}
                 vlan-id {{ var "vlan" }}
               !
             !`).
		WithVarMap(map[string]string{
			"vlan":  "888",
			"state": "disabled",
		}).
		Append(t)

	// update DUT with multi-vendor config
	dut.Config().New().
		WithCienaText(`should skip this`).
		WithCiscoText(`should also skip this`).
		WithAristaText(`should also skip this`).
		WithJuniperText(`should also skip this`).
		WithDrivenetsText(
			`interfaces
               ge100-0/0/0.333
                 admin-state enabled
                 vlan-id 333
               !
             !`).
		Append(t)

	// replace DUT config with static config below
	dut.Config().New().
		WithDrivenetsText(
			`interfaces
               ge100-0/0/0
                 admin-state enabled
                 description {{ var "desc" }}
               !
             !`).
		WithVarValue("desc", "ondatra").
		Push(t)

	// replace DUT config with static config below
	dut.Config().New().
		WithDrivenetsFile(filepath.Join("../testdata", "example_config_1.txt")).
		Push(t)
}

func TestDrivenetsCLI(t *testing.T) {
	cli := ondatra.DUT(t, "dut").CLI()

	result := cli.RunResult(t, "show system name")
	if !strings.Contains(result.Output(), "System Name: ") {
		t.Fatalf("unexpected command output: stdout: %s stderr: %s",
			result.Output(), result.Error())
	}

	result = cli.RunResult(t, "show interfaces ge100-0/0/a")
	if len(result.Error()) == 0 {
		t.Fatalf("unexpected command output: stdout: %s stderr: %s",
			result.Output(), result.Error())
	}

	stdout := cli.Run(t, "show system name")
	if !strings.Contains(stdout, "System Name: ") {
		t.Fatalf("unexpected command output: %s", stdout)
	}
}
