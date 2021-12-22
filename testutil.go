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
	"sync"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/fakebind"
	"github.com/openconfig/ondatra/internal/reservation"

	opb "github.com/openconfig/ondatra/proto"
)

var (
	mu       sync.Mutex
	fakeBind *fakebind.Binding

	fakeTBPath = filepath.Join("testdata", "testbed.pb.txt")

	fakeRes = &reservation.Reservation{
		DUTs: map[string]*reservation.DUT{
			"dut": &reservation.DUT{&reservation.Dims{
				Name:            "pf01.xxx01",
				Vendor:          opb.Device_ARISTA,
				HardwareModel:   "aristaModel",
				SoftwareVersion: "aristaVersion",
				Ports: map[string]*reservation.Port{
					"port1": &reservation.Port{Name: "Et1/2/3", Speed: opb.Port_S_10GB},
					"port2": &reservation.Port{Name: "Et4/5/6", Speed: opb.Port_S_100GB},
				},
			}},
			"dut_cisco": &reservation.DUT{&reservation.Dims{
				Name:            "pf02.xxx01",
				Vendor:          opb.Device_CISCO,
				HardwareModel:   "ciscoModel",
				SoftwareVersion: "ciscoVersion",
			}},
			"dut_juniper": &reservation.DUT{&reservation.Dims{
				Name:            "pf03.xxx01",
				Vendor:          opb.Device_JUNIPER,
				HardwareModel:   "juniperModel",
				SoftwareVersion: "juniperVersion",
			}},
		},
		ATEs: map[string]*reservation.ATE{
			"ate": &reservation.ATE{&reservation.Dims{
				Name:            "ix1",
				Vendor:          opb.Device_IXIA,
				HardwareModel:   "ixiaModel",
				SoftwareVersion: "ixiaVersion",
				Ports: map[string]*reservation.Port{
					"port1": &reservation.Port{Name: "1/1", Speed: opb.Port_S_10GB},
					"port2": &reservation.Port{Name: "1/2", Speed: opb.Port_S_100GB},
				},
			}},
		},
	}
)

// Initializes Ondatra with a fake binding implementation.
func initFakeBinding(t *testing.T) {
	t.Helper()
	mu.Lock()
	defer mu.Unlock()
	if fakeBind == nil {
		fakeBind = &fakebind.Binding{}
		binding.Init(fakeBind)
	} else {
		fakeBind.Reset()
	}
	fakeBind.Reservation = fakeRes
	if err := release(); err != nil {
		t.Fatalf("Failed to release prior testbed: %v", err)
	}
}

// Reserves the fake testbed.
func reserveFakeTestbed(t *testing.T) {
	t.Helper()
	if err := reserve(fakeTBPath, time.Hour, 0); err != nil {
		t.Fatalf("Failed to reserve testbed: %v", err)
	}
}
