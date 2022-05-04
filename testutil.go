// Copyright 2020 Google LLC
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

package ondatra

import (
	"path/filepath"
	"testing"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/ondatra/internal/flags"

	opb "github.com/openconfig/ondatra/proto"
)

var (
	fakeBind *fakebind.Binding

	fakeTBPath = filepath.Join("testdata", "testbed.pb.txt")

	fakeArista = &fakebind.DUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{
			Name:            "pf01.xxx01",
			Vendor:          opb.Device_ARISTA,
			HardwareModel:   "aristaModel",
			SoftwareVersion: "aristaVersion",
			Ports: map[string]*binding.Port{
				"port1": &binding.Port{Name: "Et1/2/3", Speed: opb.Port_S_10GB},
				"port2": &binding.Port{Name: "Et4/5/6", Speed: opb.Port_S_100GB},
			},
		}},
	}
	fakeCisco = &fakebind.DUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{
			Name:            "pf02.xxx01",
			Vendor:          opb.Device_CISCO,
			HardwareModel:   "ciscoModel",
			SoftwareVersion: "ciscoVersion",
			Ports: map[string]*binding.Port{
				"port1": &binding.Port{Name: "Et1/2/3", Speed: opb.Port_S_10GB},
			},
		}},
	}
	fakeJuniper = &fakebind.DUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{
			Name:            "pf03.xxx01",
			Vendor:          opb.Device_JUNIPER,
			HardwareModel:   "juniperModel",
			SoftwareVersion: "juniperVersion",
			Ports: map[string]*binding.Port{
				"port1": &binding.Port{Name: "Et1/2/3", Speed: opb.Port_S_10GB},
			},
		}},
	}
	fakeIxia = &fakebind.ATE{
		AbstractATE: &binding.AbstractATE{&binding.Dims{
			Name:            "ix1",
			Vendor:          opb.Device_IXIA,
			HardwareModel:   "ixiaModel",
			SoftwareVersion: "ixiaVersion",
			Ports: map[string]*binding.Port{
				"port1": &binding.Port{Name: "1/1", Speed: opb.Port_S_10GB},
				"port2": &binding.Port{Name: "1/2", Speed: opb.Port_S_100GB},
			},
		}},
	}
	fakeRes = &binding.Reservation{
		DUTs: map[string]binding.DUT{
			"dut_arista":  fakeArista,
			"dut_cisco":   fakeCisco,
			"dut_juniper": fakeJuniper,
		},
		ATEs: map[string]binding.ATE{
			"ate_ixia": fakeIxia,
		},
	}
)

// Initializes Ondatra with a fake binding implementation.
func initFakeBinding(t *testing.T) {
	if fakeBind != nil {
		if err := release(); err != nil {
			t.Fatalf("Failed to release prior testbed: %v", err)
		}
	}
	fakeBind = fakebind.Setup().StubReservation(fakeRes)
}

// Reserves the fake testbed.
func reserveFakeTestbed(t *testing.T) {
	t.Helper()
	if err := reserve(&flags.Values{TestbedPath: fakeTBPath}); err != nil {
		t.Fatalf("Failed to reserve testbed: %v", err)
	}
}
