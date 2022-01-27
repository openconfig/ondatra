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
	"testing"

	kinit "github.com/openconfig/ondatra/knebind/init"
	"github.com/openconfig/ondatra"
)

func TestMain(m *testing.M) {
	ondatra.RunTests(m, kinit.Init)
}

func TestConfig(t *testing.T) {
	dut := ondatra.DUT(t, "dut1")
	const config = `
	  interface Ethernet1
	    no switchport
	    ip address 1.1.1.1/24
	  !
	  interface Ethernet2
	    no switchport
	    ip address 2.2.2.1/24
	  !
	  ip routing
	  !
	  router bgp 1111
	    no bgp enforce-first-as
	    neighbor 1.1.1.2 remote-as 2222
	    neighbor 2.2.2.2 remote-as 3333
	  !`
	dut.Config().New().WithAristaText(config).Push(t)
}

func TestGNMI(t *testing.T) {
	dut1 := ondatra.DUT(t, "dut1")
	dut2 := ondatra.DUT(t, "dut2")
	sys1 := dut1.Telemetry().System().Lookup(t)
	if !sys1.IsPresent() {
		t.Fatalf("No System telemetry for %v", dut1)
	}
	sys2 := dut2.Telemetry().System().Lookup(t)
	if !sys2.IsPresent() {
		t.Fatalf("No System telemetry for %v", dut2)
	}
}
