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
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/gnmi"
	"github.com/openconfig/ondatra/gnmi/oc"
	kinit "github.com/openconfig/ondatra/knebind/init"
)

func TestMain(m *testing.M) {
	ondatra.RunTests(m, kinit.Init)
}

func TestGNMIConfig(t *testing.T) {
	dut := ondatra.DUT(t, "dut")
	port := dut.Port(t, "port1")
	iface := &oc.Interface{}
	iface.SetName(port.Name())

	addr := iface.GetOrCreateSubinterface(0).GetOrCreateIpv4().GetOrCreateAddress("10.0.0.1")
	addr.SetPrefixLength(30)

	gnmi.Replace(t, dut, gnmi.OC().Interface(port.Name()).Config(), iface)
}

func TestGNMISubscribe(t *testing.T) {
	dut := ondatra.DUT(t, "dut")
	sys := gnmi.Lookup(t, dut, gnmi.OC().System().State())
	if !sys.IsPresent() {
		t.Fatalf("No System telemetry for %v", dut)
	}
}

func TestOTG(t *testing.T) {
	ate := ondatra.ATE(t, "ate")
	cfg := ate.OTG().NewConfig(t)
	cfg.Ports().Add().SetName("port1")
	cfg.Ports().Add().SetName("port2")
	ate.OTG().PushConfig(t, cfg)

	portNames := gnmi.GetAll(t, ate.OTG(), gnmi.OTG().PortAny().Name().State())
	sort.Strings(portNames)
	if want := []string{"port1", "port2"}; !cmp.Equal(portNames, want) {
		t.Errorf("Telemetry got port names %v, want %v", portNames, want)
	}
}
