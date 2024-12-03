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

package dnbind

import (
	"context"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/cli"
	"github.com/openconfig/ondatra/config"
	"github.com/openconfig/ondatra/dnbind/creds"
	opb "github.com/openconfig/ondatra/proto"
)

var (
	credFlags = creds.DefineFlags()
)

// requires passing DUT credentials via --node_creds=hostname/user/pass
func mockReserve() (*binding.Reservation, error) {
	creds, err := credFlags.Parse()
	if err != nil {
		return nil, err
	}

	res := &binding.Reservation{
		DUTs: make(map[string]binding.DUT),
	}

	for node, _ := range creds.Node {
		dims := &binding.Dims{
			Name:   node,
			Vendor: opb.Device_DRIVENETS,
		}
		res.DUTs[node] = &dnDUT{
			AbstractDUT: &binding.AbstractDUT{Dims: dims},
			bind:        &Bind{cfg: &Config{Credentials: creds}},
		}
	}

	return res, nil
}

func TestDrivenetsBinding(t *testing.T) {
	res, err := mockReserve()
	if err != nil {
		t.Fatalf("failed to create binding: %s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, dut := range res.DUTs {
		cli, err := dut.DialCLI(ctx)
		if err != nil {
			t.Fatal(err.Error())
		}

		res, err := cli.RunCommand(ctx, "show system")
		if err != nil {
			t.Fatal(err.Error())
		}
		t.Log(res.Output())

		config := `interfaces
                     ge100-0/0/0.123
                       admin-state enabled
                       vlan-id 123
                     !
                   !`
		dut.PushConfig(ctx, config, false)

		res, err = cli.RunCommand(ctx, "show interfaces ge100-0/0/0.123")
		if err != nil {
			t.Fatal(err.Error())
		}
		t.Log(res.Output())

		config = `no interfaces ge100-0/0/0.123`
		dut.PushConfig(ctx, config, false)

		res, err = cli.RunCommand(ctx, "show interfaces ge100-0/0/0.123")
		if err != nil {
			t.Fatal(err.Error())
		}
		t.Log(res.Error())
	}
}

// TODO: check push config outputs?
func TestDrivenetsVendorConfig(t *testing.T) {
	res, err := mockReserve()
	if err != nil {
		t.Fatalf("failed to create binding: %s", err.Error())
	}

	for _, dut := range res.DUTs {
		// updates DUT config with static config below
		config.NewVendorConfig(dut).
			WithDrivenetsText(
				`interfaces
                   ge100-0/0/0.321
                     admin-state enabled
                     vlan-id 321
                   !
                 !`).
			Append(t)

		// updates DUT config with replace config below
		config.NewVendorConfig(dut).
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
		config.NewVendorConfig(dut).
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
		config.NewVendorConfig(dut).
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
		config.NewVendorConfig(dut).
			WithDrivenetsFile(filepath.Join("testdata", "example_config_1.txt")).
			Push(t)
	}
}

func TestDrivenetsCLI(t *testing.T) {
	res, err := mockReserve()
	if err != nil {
		t.Fatalf("failed to create binding: %s", err.Error())
	}

	for _, dut := range res.DUTs {
		cli := cli.New(dut)

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
}
