// Copyright 2022 Google LLC
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

package solver

import (
	"testing"

	opb "github.com/openconfig/ondatra/proto"
)

func BenchmarkSolveScale(b *testing.B) {
	const topoText = `
		nodes: {
			name: "node1"
			type: NOKIA_SRL
		}
		nodes: {
			name: "node2"
			type: NOKIA_SRL
		}
		nodes: {
			name: "node3"
			type: NOKIA_SRL
		}
		nodes: {
			name: "node4"
			type: NOKIA_SRL
		}
		nodes: {
			name: "node5"
			type: NOKIA_SRL
		}
		nodes: {
			name: "node6"
			type: NOKIA_SRL
		}
		nodes: {
			name: "node7"
			type: NOKIA_SRL
		}
		nodes: {
			name: "node8"
			type: NOKIA_SRL
		}
		nodes: {
			name: "node9"
			type: CISCO_CXR
		}
		nodes: {
			name: "node10"
			type: CISCO_CXR
		}
		nodes: {
			name: "node11"
			type: CISCO_CXR
		}
		nodes: {
			name: "node12"
			type: JUNIPER_CEVO
		}
		nodes: {
			name: "node13"
			type: JUNIPER_CEVO
		}
		nodes: {
			name: "node14"
			type: JUNIPER_CEVO
		}
		nodes: {
			name: "node15"
			type: JUNIPER_CEVO
		}
		nodes: {
			name: "node16"
			type: ARISTA_CEOS
		}
		nodes: {
			name: "node17"
			type: ARISTA_CEOS
		}
		nodes: {
			name: "node18"
			type: ARISTA_CEOS
		}
		nodes: {
			name: "node19"
			type: IXIA_TG
		}
		links: {
			a_node: "node1"
			a_int: "eth1"
			z_node: "node2"
			z_int: "eth1"
		}
		links: {
			a_node: "node3"
			a_int: "eth1"
			z_node: "node4"
			z_int: "eth1"
		}
		links: {
			a_node: "node5"
			a_int: "eth1"
			z_node: "node6"
			z_int: "eth1"
		}
		links: {
			a_node: "node7"
			a_int: "eth1"
			z_node: "node8"
			z_int: "eth1"
		}
		links: {
			a_node: "node1"
			a_int: "eth2"
			z_node: "node9"
			z_int: "eth1"
		}
		links: {
			a_node: "node2"
			a_int: "eth2"
			z_node: "node9"
			z_int: "eth2"
		}
		links: {
			a_node: "node9"
			a_int: "eth3"
			z_node: "node16"
			z_int: "eth1"
		}
		links: {
			a_node: "node10"
			a_int: "eth1"
			z_node: "node7"
			z_int: "eth2"
		}
		links: {
			a_node: "node10"
			a_int: "eth2"
			z_node: "node8"
			z_int: "eth2"
		}
		links: {
			a_node: "node10"
			a_int: "eth3"
			z_node: "node17"
			z_int: "eth1"
		}
		links: {
			a_node: "node11"
			a_int: "eth1"
			z_node: "node18"
			z_int: "eth1"
		}
		links: {
			a_node: "node11"
			a_int: "eth2"
			z_node: "node5"
			z_int: "eth2"
		}
		links: {
			a_node: "node12"
			a_int: "eth1"
			z_node: "node3"
			z_int: "eth2"
		}
		links: {
			a_node: "node12"
			a_int: "eth2"
			z_node: "node4"
			z_int: "eth2"
		}
		links: {
			a_node: "node12"
			a_int: "eth3"
			z_node: "node16"
			z_int: "eth2"
		}
		links: {
			a_node: "node13"
			a_int: "eth1"
			z_node: "node3"
			z_int: "eth3"
		}
		links: {
			a_node: "node13"
			a_int: "eth2"
			z_node: "node5"
			z_int: "eth3"
		}
		links: {
			a_node: "node13"
			a_int: "eth3"
			z_node: "node17"
			z_int: "eth2"
		}
		links: {
			a_node: "node14"
			a_int: "eth1"
			z_node: "node6"
			z_int: "eth2"
		}
		links: {
			a_node: "node14"
			a_int: "eth2"
			z_node: "node17"
			z_int: "eth3"
		}
		links: {
			a_node: "node15"
			a_int: "eth1"
			z_node: "node7"
			z_int: "eth3"
		}
		links: {
			a_node: "node15"
			a_int: "eth2"
			z_node: "node8"
			z_int: "eth3"
		}
		links: {
			a_node: "node15"
			a_int: "eth13"
			z_node: "node18"
			z_int: "eth2"
		}
		links: {
			a_node: "node16"
			a_int: "eth3"
			z_node: "node19"
			z_int: "eth1"
		}
		links: {
			a_node: "node17"
			a_int: "eth4"
			z_node: "node19"
			z_int: "eth2"
		}
		links: {
			a_node: "node18"
			a_int: "eth3"
			z_node: "node19"
			z_int: "eth3"
		}`
	topo := unmarshalTopo(b, topoText)
	tb := &opb.Testbed{
		Duts: []*opb.Device{{
			Id:     "dut1",
			Vendor: opb.Device_ARISTA,
			Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}, {Id: "port3"}},
		}, {
			Id:     "dut2",
			Vendor: opb.Device_CISCO,
			Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}, {Id: "port3"}},
		}, {
			Id:     "dut3",
			Vendor: opb.Device_JUNIPER,
			Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}, {Id: "port3"}},
		}, {
			Id:     "dut4",
			Vendor: opb.Device_NOKIA,
			Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}},
		}, {
			Id:     "dut5",
			Vendor: opb.Device_NOKIA,
			Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}},
		}, {
			Id:     "dut6",
			Vendor: opb.Device_NOKIA,
			Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}},
		}, {
			Id:     "dut7",
			Vendor: opb.Device_NOKIA,
			Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}},
		}},
		Ates: []*opb.Device{{
			Id:    "ate",
			Ports: []*opb.Port{{Id: "port1"}},
		}},
		Links: []*opb.Link{{
			A: "dut1:port1",
			B: "dut2:port1",
		}, {
			A: "dut1:port2",
			B: "dut3:port1",
		}, {
			A: "dut1:port3",
			B: "ate:port1",
		}, {
			A: "dut2:port2",
			B: "dut4:port1",
		}, {
			A: "dut2:port3",
			B: "dut5:port1",
		}, {
			A: "dut4:port2",
			B: "dut5:port2",
		}, {
			A: "dut3:port2",
			B: "dut6:port1",
		}, {
			A: "dut3:port3",
			B: "dut7:port1",
		}, {
			A: "dut6:port2",
			B: "dut7:port2",
		}},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_, err := Solve(tb, topo)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Solve() got unexpected error: %v", err)
		}
	}
}
