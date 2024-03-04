// Copyright 2024 Google LLC
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

package portgraph

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	opb "github.com/openconfig/ondatra/proto"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestTestbedToAbstractGraph(t *testing.T) {
	port := &opb.Port{
		Id:             "port1",
		Speed:          opb.Port_S_100GB,
		CardModelValue: &opb.Port_CardModel{CardModel: "model"},
		PmdValue:       &opb.Port_Pmd_{Pmd: opb.Port_PMD_100GBASE_LR4},
	}
	portEmpty := &opb.Port{
		Id:             "portEmpty",
		Speed:          opb.Port_SPEED_UNSPECIFIED,
		CardModelValue: &opb.Port_CardModel{},
		PmdValue:       &opb.Port_Pmd_{Pmd: opb.Port_PMD_UNSPECIFIED},
	}
	portRegex := &opb.Port{
		Id:             "portRegex",
		CardModelValue: &opb.Port_CardModelRegex{CardModelRegex: "modelregex"},
		PmdValue:       &opb.Port_PmdRegex{PmdRegex: ".*"},
	}
	dut := &opb.Device{
		Id:                   "dut",
		Vendor:               opb.Device_ARISTA,
		HardwareModelValue:   &opb.Device_HardwareModel{HardwareModel: "hw"},
		SoftwareVersionValue: &opb.Device_SoftwareVersion{SoftwareVersion: "sw"},
		Ports:                []*opb.Port{port},
		ExtraDimensions:      map[string]string{"extra_dimension": "hello"},
	}
	dutEmpty := &opb.Device{
		Id:                   "dutEmpty",
		Vendor:               opb.Device_VENDOR_UNSPECIFIED,
		HardwareModelValue:   &opb.Device_HardwareModel{},
		SoftwareVersionValue: &opb.Device_SoftwareVersion{},
		Ports:                []*opb.Port{portEmpty},
	}
	dutRegex := &opb.Device{
		Id:                   "dutRegex",
		HardwareModelValue:   &opb.Device_HardwareModelRegex{HardwareModelRegex: "hwreg"},
		SoftwareVersionValue: &opb.Device_SoftwareVersionRegex{SoftwareVersionRegex: "swreg"},
		Ports:                []*opb.Port{portRegex},
	}
	ate := &opb.Device{
		Id:    "ate",
		Ports: []*opb.Port{port},
	}
	link := &opb.Link{
		A: "dut:port1",
		B: "ate:port1",
	}

	tb := &opb.Testbed{
		Duts:  []*opb.Device{dut, dutEmpty, dutRegex},
		Ates:  []*opb.Device{ate},
		Links: []*opb.Link{link},
	}
	wantDev := map[string]*opb.Device{
		"dut":      dut,
		"dutEmpty": dutEmpty,
		"dutRegex": dutRegex,
		"ate":      ate,
	}
	wantPort := map[string]*opb.Port{
		"dut:port1":          port,
		"dutEmpty:portEmpty": portEmpty,
		"dutRegex:portRegex": portRegex,
		"ate:port1":          port,
	}

	graph, node2Dev, port2Port, err := TestbedToAbstractGraph(tb, nil)
	if err != nil {
		t.Fatalf("TestbedToAbstractGraph() got error %v, want nil", err)
	}
	if wantNodes := 4; len(graph.Nodes) != wantNodes {
		t.Fatalf("TestbedToAbstractGraph() got %d nodes, want %v", len(graph.Nodes), wantNodes)
	}
	for _, node := range graph.Nodes {
		if got, ok := node2Dev[node]; !ok {
			t.Errorf("TestbedToAbstractGraph() got node %q not mapped to any device", node.Desc)
		} else if diff := cmp.Diff(wantDev[node.Desc], got, protocmp.Transform()); diff != "" {
			t.Errorf("TestbedToAbstractGraph() returned unexpected device diff (-want +got):\n%s", diff)
		}
		for _, port := range node.Ports {
			if got, ok := port2Port[port]; !ok {
				t.Errorf("TestbedToAbstractGraph() got port %q not mapped to any port", port.Desc)
			} else if diff := cmp.Diff(wantPort[port.Desc], got, protocmp.Transform()); diff != "" {
				t.Errorf("TestbedToAbstractGraph() returned unexpected port diff (-want +got):\n%s", diff)
			}
		}
	}
}
