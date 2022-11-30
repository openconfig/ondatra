// Copyright 2022 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"

	opb "github.com/openconfig/ondatra/proto"
)

func TestPush(t *testing.T) {
	var gotConfig string
	var gotReset bool
	dut := &fakebind.DUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{
			Vendor: opb.Device_ARISTA,
			Ports: map[string]*binding.Port{
				"port1": &binding.Port{Name: "Eth1/2/3"},
				"port2": &binding.Port{Name: "Eth4/5/6"},
			},
		}},
		PushConfigFn: func(_ context.Context, config string, reset bool) error {
			gotConfig = config
			gotReset = reset
			return nil
		},
	}

	testsPass := []struct {
		desc       string
		config     *VendorConfig
		wantConfig string
		wantFatal  string
	}{{
		desc:       "correct text",
		config:     NewVendorConfig(dut).WithText("generated"),
		wantConfig: "generated",
	}, {
		desc:       "correct file",
		config:     NewVendorConfig(dut).WithFile(filepath.Join("testdata", "example_config_1.txt")),
		wantConfig: "example_config_1",
	}, {
		desc: "correct per-vendor text",
		config: NewVendorConfig(dut).
			WithAristaText("Arista config").
			WithCiscoText("Cisco config").
			WithJuniperText("Juniper config"),
		wantConfig: "Arista config",
	}, {
		desc: "correct per-vendor file",
		config: NewVendorConfig(dut).
			WithAristaFile(filepath.Join("testdata", "example_config_1.txt")).
			WithCiscoText("Cisco config").
			WithJuniperFile(filepath.Join("testdata", "example_config_2.txt")),
		wantConfig: "example_config_1",
	}, {
		desc: "port template",
		config: NewVendorConfig(dut).
			WithAristaText(`reconfigure {{ port "port1" }} and {{ port "port2" }}`),
		wantConfig: "reconfigure Eth1/2/3 and Eth4/5/6",
	}, {
		desc: "secrets template",
		config: NewVendorConfig(dut).
			WithAristaText(`shh {{ secrets "hello" "there" }} wink`),
		wantConfig: `shh {{ secrets "hello" "there" }} wink`,
	}, {
		desc: "var template",
		config: NewVendorConfig(dut).
			WithAristaText(`hello {{ var "foo" }} there`).
			WithVarValue("foo", "bar"),
		wantConfig: `hello bar there`,
	}, {
		desc: "var map template",
		config: NewVendorConfig(dut).
			WithAristaText(`hello {{ var "x" }} and {{ var "y" }}`).
			WithVarMap(map[string]string{"x": "apple", "y": "orange"}),
		wantConfig: `hello apple and orange`,
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			gotConfig = ""
			gotReset = false
			tt.config.Push(t)
			if diff := cmp.Diff(tt.wantConfig, gotConfig); diff != "" {
				t.Errorf("Push(t) got unexpected config diff(-want,+got):\n %s", diff)
			}
			if !gotReset {
				t.Errorf("Push(t) got unexpected reset %v, want true", gotReset)
			}
		})
	}

	testsFail := []struct {
		desc      string
		config    *VendorConfig
		wantFatal string
	}{{
		desc:      "no config",
		config:    NewVendorConfig(dut).WithCiscoText("gaga"),
		wantFatal: "no config",
	}, {
		desc: "all-vendor and per-vendor",
		config: NewVendorConfig(dut).
			WithText("text1").
			WithAristaText("text2"),
		wantFatal: "all-vendor and per-vendor",
	}, {
		desc:      "template function does not exist",
		config:    NewVendorConfig(dut).WithAristaText(`{{ qwerty "port1" }}`),
		wantFatal: `function "qwerty" not defined`,
	}, {
		desc:      "port name does not exist",
		config:    NewVendorConfig(dut).WithAristaText(`{{ port "port10" }}`),
		wantFatal: "not found",
	}, {
		desc:      "template malformed",
		config:    NewVendorConfig(dut).WithAristaText(`{{ port"port10" }}`),
		wantFatal: "bad character",
	}, {
		desc:      "var has no value",
		config:    NewVendorConfig(dut).WithAristaText(`{{ var "key1" }}`),
		wantFatal: "no value for key",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			got := testt.ExpectFatal(t, func(t testing.TB) {
				tt.config.Push(t)
			})
			if !strings.Contains(got, tt.wantFatal) {
				t.Errorf("Push(t) failed with message %q, want %q", got, tt.wantFatal)
			}
		})
	}

	t.Run("append", func(t *testing.T) {
		wantConfig := "appended config"
		NewVendorConfig(dut).WithAristaText(wantConfig).Append(t)
		if gotConfig != wantConfig {
			t.Errorf("Append(t) got pushed config %v, want %v", gotConfig, wantConfig)
		}
		if gotReset {
			t.Errorf("Append(t) got unexpected reset %v, want false", gotReset)
		}
	})
}
