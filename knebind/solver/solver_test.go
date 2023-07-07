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
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/binding"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"

	tpb "github.com/openconfig/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

func TestSolve(t *testing.T) {
	const topoText = `
		nodes: {
		  name: "node1"
		  vendor: ARISTA
			model: "ceos"
			os: "eos"
		  services: {
		    key: 1234
		    value: {
		      name: "gnmi"
		    }
		  }
		  interfaces: {
		    key: "eth1"
		    value: {
		      name: "Ethernet1"
		    }
		  }
		  interfaces: {
		    key: "eth2"
		    value: {
		      name: "Ethernet2"
		    }
		  }
		}
		nodes: {
		  name: "node2"
		  vendor: CISCO
		  services: {
		    key: 2345
		    value: {
		      name: "gnmi"
		    }
		  }
		  interfaces: {
		    key: "eth1"
		    value: {
					name: "GigabitEthernet0/0/0/0"
				}
		  }
		  interfaces: {
		    key: "eth2"
		    value: {}
		  }
		}
		nodes: {
		  name: "node3"
		  vendor: JUNIPER
			model: "cptx"
			os: "evo"
		  services: {
		    key: 3456
		    value: {
		      name: "gnmi"
		    }
		  }
		  interfaces: {
		    key: "eth1"
		    value: {}
		  }
		}
		nodes: {
		  name: "node4"
		  vendor: KEYSIGHT
		  interfaces: {
		    key: "eth1"
		    value: {}
		  }
		}
		nodes: {
		  name: "node5"
		  vendor: OPENCONFIG
			model: "MAGNA"
		  labels: {
				key: "ondatra-role"
				value: "ATE"
			}
		}
		links: {
		  a_node: "node1"
		  a_int: "eth1"
		  z_node: "node2"
		  z_int: "eth1"
		}
		links: {
		  a_node: "node2"
		  a_int: "eth2"
		  z_node: "node3"
		  z_int: "eth1"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth2"
		  z_node: "node4"
		  z_int: "eth1"
		}`
	topo := unmarshalTopo(t, topoText)

	dut1 := &opb.Device{
		Id:                   "dut1",
		Vendor:               opb.Device_ARISTA,
		HardwareModelValue:   &opb.Device_HardwareModel{"ceos"},
		SoftwareVersionValue: &opb.Device_SoftwareVersion{"eos"},
		Ports:                []*opb.Port{{Id: "port1"}, {Id: "port2"}},
	}
	dut2 := &opb.Device{
		Id:    "dut2",
		Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}},
	}
	dut3 := &opb.Device{
		Id:                   "dut3",
		Vendor:               opb.Device_JUNIPER,
		SoftwareVersionValue: &opb.Device_SoftwareVersionRegex{"^evo$"},
		Ports:                []*opb.Port{{Id: "port1"}},
	}
	ate1 := &opb.Device{
		Id:    "ate1",
		Ports: []*opb.Port{{Id: "port1"}},
	}
	ate2 := &opb.Device{
		Id:     "ate2",
		Vendor: opb.Device_OPENCONFIG,
	}
	link12 := &opb.Link{
		A: "dut1:port1",
		B: "dut2:port1",
	}
	link23 := &opb.Link{
		A: "dut2:port2",
		B: "dut3:port1",
	}
	link14 := &opb.Link{
		A: "dut1:port2",
		B: "ate1:port1",
	}

	wantDUTServices := map[string]*tpb.Service{
		"gnmi": &tpb.Service{Name: "gnmi"},
	}
	wantDUT1 := &ServiceDUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{
			Name:            "node1",
			Vendor:          opb.Device_ARISTA,
			HardwareModel:   "ceos",
			SoftwareVersion: "eos",
			Ports: map[string]*binding.Port{
				"port1": {Name: "Ethernet1"},
				"port2": {Name: "Ethernet2"},
			},
			CustomData: map[string]any{KNEServiceMapKey: wantDUTServices},
		}},
		Services:   wantDUTServices,
		NodeVendor: tpb.Vendor_ARISTA,
	}
	wantDUT2 := &ServiceDUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{
			Name:   "node2",
			Vendor: opb.Device_CISCO,
			Ports: map[string]*binding.Port{
				"port1": {Name: "GigabitEthernet0/0/0/0"},
				"port2": {Name: "eth2"},
			},
			CustomData: map[string]any{KNEServiceMapKey: wantDUTServices},
		}},
		Services:   wantDUTServices,
		NodeVendor: tpb.Vendor_CISCO,
	}
	wantDUT3 := &ServiceDUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{
			Name:            "node3",
			Vendor:          opb.Device_JUNIPER,
			HardwareModel:   "cptx",
			SoftwareVersion: "evo",
			Ports: map[string]*binding.Port{
				"port1": {Name: "eth1"},
			},
			CustomData: map[string]any{KNEServiceMapKey: wantDUTServices},
		}},
		Services:   wantDUTServices,
		NodeVendor: tpb.Vendor_JUNIPER,
	}
	wantATEServices := make(map[string]*tpb.Service)
	wantATE1 := &ServiceATE{
		AbstractATE: &binding.AbstractATE{&binding.Dims{
			Name:   "node4",
			Vendor: opb.Device_IXIA,
			Ports: map[string]*binding.Port{
				"port1": {Name: "eth1"},
			},
			CustomData: map[string]any{KNEServiceMapKey: wantATEServices},
		}},
		Services:   wantATEServices,
		NodeVendor: tpb.Vendor_KEYSIGHT,
	}
	wantATE2 := &ServiceATE{
		AbstractATE: &binding.AbstractATE{&binding.Dims{
			Name:          "node5",
			Vendor:        opb.Device_OPENCONFIG,
			HardwareModel: "MAGNA",
			Ports:         map[string]*binding.Port{},
			CustomData:    map[string]any{KNEServiceMapKey: wantATEServices},
		}},
		Services:   wantATEServices,
		NodeVendor: tpb.Vendor_OPENCONFIG,
	}

	tests := []struct {
		desc    string
		tb      *opb.Testbed
		partial map[string]string
		wantRes *binding.Reservation
	}{{
		desc: "one dut",
		tb: &opb.Testbed{
			Duts: []*opb.Device{dut3},
		},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{
				"dut3": wantDUT3,
			},
			ATEs: map[string]binding.ATE{},
		},
	}, {
		desc: "one dut with partial",
		tb: &opb.Testbed{
			Duts: []*opb.Device{
				&opb.Device{
					Id: "dut",
				},
			},
		},
		partial: map[string]string{"dut": "node3"},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{
				"dut": &ServiceDUT{
					AbstractDUT: &binding.AbstractDUT{&binding.Dims{
						Name:            "node3",
						Vendor:          opb.Device_JUNIPER,
						HardwareModel:   "cptx",
						SoftwareVersion: "evo",
						Ports:           map[string]*binding.Port{},
						CustomData:      map[string]any{KNEServiceMapKey: wantDUTServices},
					}},
					Services:   wantDUTServices,
					NodeVendor: tpb.Vendor_JUNIPER,
				},
			},
			ATEs: map[string]binding.ATE{},
		},
	}, {
		desc: "one dut with partial + port intf name",
		tb: &opb.Testbed{
			Duts: []*opb.Device{
				&opb.Device{
					Id:    "dut",
					Ports: []*opb.Port{&opb.Port{Id: "port"}},
				},
			},
		},
		partial: map[string]string{"dut": "node2", "dut:port": "eth2"},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{
				"dut": &ServiceDUT{
					AbstractDUT: &binding.AbstractDUT{&binding.Dims{
						Name:   "node2",
						Vendor: opb.Device_CISCO,
						Ports: map[string]*binding.Port{
							"port": {Name: "eth2"},
						},
						CustomData: map[string]any{KNEServiceMapKey: wantDUTServices},
					}},
					Services:   wantDUTServices,
					NodeVendor: tpb.Vendor_CISCO,
				},
			},
			ATEs: map[string]binding.ATE{},
		},
	}, {
		desc: "one dut with partial + port vendor name",
		tb: &opb.Testbed{
			Duts: []*opb.Device{
				&opb.Device{
					Id:    "dut",
					Ports: []*opb.Port{&opb.Port{Id: "port"}},
				},
			},
		},
		partial: map[string]string{"dut": "node2", "dut:port": "GigabitEthernet0/0/0/0"},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{
				"dut": &ServiceDUT{
					AbstractDUT: &binding.AbstractDUT{&binding.Dims{
						Name:   "node2",
						Vendor: opb.Device_CISCO,
						Ports: map[string]*binding.Port{
							"port": {Name: "GigabitEthernet0/0/0/0"},
						},
						CustomData: map[string]any{KNEServiceMapKey: wantDUTServices},
					}},
					Services:   wantDUTServices,
					NodeVendor: tpb.Vendor_CISCO,
				},
			},
			ATEs: map[string]binding.ATE{},
		},
	}, {
		desc: "one ate",
		tb: &opb.Testbed{
			Ates: []*opb.Device{ate1},
		},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{},
			ATEs: map[string]binding.ATE{
				"ate1": wantATE1,
			},
		},
	}, {
		desc: "two ates",
		tb: &opb.Testbed{
			Ates: []*opb.Device{ate1, ate2},
		},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{},
			ATEs: map[string]binding.ATE{
				"ate1": wantATE1,
				"ate2": wantATE2,
			},
		},
	}, {
		desc: "two duts",
		tb: &opb.Testbed{
			Duts:  []*opb.Device{dut1, dut2},
			Links: []*opb.Link{link12},
		},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{
				"dut1": wantDUT1,
				"dut2": wantDUT2,
			},
			ATEs: map[string]binding.ATE{},
		},
	}, {
		desc: "dut and ate",
		tb: &opb.Testbed{
			Duts:  []*opb.Device{dut1},
			Ates:  []*opb.Device{ate1},
			Links: []*opb.Link{link14},
		},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{
				"dut1": wantDUT1,
			},
			ATEs: map[string]binding.ATE{
				"ate1": wantATE1,
			},
		},
	}, {
		desc: "three duts",
		tb: &opb.Testbed{
			Duts:  []*opb.Device{dut1, dut2, dut3},
			Links: []*opb.Link{link12, link23},
		},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{
				"dut1": wantDUT1,
				"dut2": wantDUT2,
				"dut3": wantDUT3,
			},
			ATEs: map[string]binding.ATE{},
		},
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotRes, err := Solve(test.tb, topo, test.partial)
			if err != nil {
				t.Fatalf("Solve() got unexpected error: %v", err)
			}
			gotRes.ID = ""
			if diff := cmp.Diff(test.wantRes, gotRes, protocmp.Transform()); diff != "" {
				t.Errorf("Solve() got unexpected diff in reservation (-want,+got): %s", diff)
			}
		})
	}
}

func TestPortGroupsSolve(t *testing.T) {
	const topoText = `
		nodes: {
		  name: "node1"
		  vendor: KEYSIGHT
		  interfaces: {
		    key: "eth1"
		    value: {group: "lag1"}
		  }
		  interfaces: {
		    key: "eth2"
		    value: {group: "lag1"}
		  }
		  interfaces: {
		    key: "eth3"
		    value: {group: "lag2"}
		  }
		  interfaces: {
		    key: "eth4"
		    value: {group: "lag2"}
		  }
		  interfaces: {
		    key: "eth5"
		    value: {group: "lag3"}
		  }
		}
		nodes: {
		  name: "node2"
		  vendor: ARISTA
		}
		links: {
		  a_node: "node1"
		  a_int: "eth1"
		  z_node: "node2"
		  z_int: "eth1"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth2"
		  z_node: "node2"
		  z_int: "eth2"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth3"
		  z_node: "node2"
		  z_int: "eth3"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth4"
		  z_node: "node2"
		  z_int: "eth4"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth5"
		  z_node: "node2"
		  z_int: "eth5"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth6"
		  z_node: "node2"
		  z_int: "eth6"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth7"
		  z_node: "node2"
		  z_int: "eth7"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth8"
		  z_node: "node2"
		  z_int: "eth8"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth9"
		  z_node: "node2"
		  z_int: "eth9"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth10"
		  z_node: "node2"
		  z_int: "eth10"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth11"
		  z_node: "node2"
		  z_int: "eth11"
		}
		links: {
			a_node: "node1"
			a_int: "eth12"
			z_node: "node2"
			z_int: "eth12"
		}
		links: {
			a_node: "node1"
			a_int: "eth13"
			z_node: "node2"
			z_int: "eth13"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth14"
		  z_node: "node2"
		  z_int: "eth14"
		}
		links: {
		  a_node: "node1"
		  a_int: "eth15"
		  z_node: "node2"
		  z_int: "eth15"
		}`
	topo := unmarshalTopo(t, topoText)

	dut1 := &opb.Device{
		Id:     "dut1",
		Vendor: opb.Device_ARISTA,
		Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}, {Id: "port3"}, {Id: "port4"}, {Id: "port5"}, {Id: "port6"},
			{Id: "port7"}, {Id: "port8"}, {Id: "port9"}, {Id: "port10"}, {Id: "port11"}, {Id: "port12"},
			{Id: "port13"}, {Id: "port14"}, {Id: "port15"}},
	}
	ate := &opb.Device{
		Id: "ate",
		Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}, {Id: "port3"}, {Id: "port4"}, {Id: "port5"}, {Id: "port6"},
			{Id: "port7"}, {Id: "port8"}, {Id: "port9"}, {Id: "port10"}, {Id: "port11", Group: "G1"},
			{Id: "port12", Group: "G2"}, {Id: "port13", Group: "G2"}, {Id: "port14", Group: "G3"}, {Id: "port15", Group: "G3"}},
	}
	link1 := &opb.Link{
		A: "dut1:port1",
		B: "ate:port1",
	}
	link2 := &opb.Link{
		A: "dut1:port2",
		B: "ate:port2",
	}
	link3 := &opb.Link{
		A: "dut1:port3",
		B: "ate:port3",
	}
	link4 := &opb.Link{
		A: "dut1:port4",
		B: "ate:port4",
	}
	link5 := &opb.Link{
		A: "dut1:port5",
		B: "ate:port5",
	}
	link6 := &opb.Link{
		A: "dut1:port6",
		B: "ate:port6",
	}
	link7 := &opb.Link{
		A: "dut1:port7",
		B: "ate:port7",
	}
	link8 := &opb.Link{
		A: "dut1:port8",
		B: "ate:port8",
	}
	link9 := &opb.Link{
		A: "dut1:port9",
		B: "ate:port9",
	}
	link10 := &opb.Link{
		A: "dut1:port10",
		B: "ate:port10",
	}
	link11 := &opb.Link{
		A: "dut1:port11",
		B: "ate:port11",
	}
	link12 := &opb.Link{
		A: "dut1:port12",
		B: "ate:port12",
	}
	link13 := &opb.Link{
		A: "dut1:port13",
		B: "ate:port13",
	}
	link14 := &opb.Link{
		A: "dut1:port14",
		B: "ate:port14",
	}
	link15 := &opb.Link{
		A: "dut1:port15",
		B: "ate:port15",
	}

	tb := &opb.Testbed{
		Duts:  []*opb.Device{dut1},
		Ates:  []*opb.Device{ate},
		Links: []*opb.Link{link1, link2, link3, link4, link5, link6, link7, link8, link9, link10, link11, link12, link13, link14, link15},
	}

	res, err := Solve(tb, topo, nil)
	if err != nil {
		t.Fatalf("Solve() got unexpected error: %v", err)
	}
	if res == nil {
		t.Fatalf("Solve() found no solution")
	}
}

func TestTestbedToAbstractGraph(t *testing.T) {
	port := &opb.Port{
		Id:             "port1",
		Speed:          opb.Port_S_100GB,
		CardModelValue: &opb.Port_CardModel{"model"},
		PmdValue:       &opb.Port_Pmd_{opb.Port_PMD_100GBASE_LR4},
	}
	portEmpty := &opb.Port{
		Id:             "portEmpty",
		Speed:          opb.Port_SPEED_UNSPECIFIED,
		CardModelValue: &opb.Port_CardModel{},
		PmdValue:       &opb.Port_Pmd_{opb.Port_PMD_UNSPECIFIED},
	}
	portRegex := &opb.Port{
		Id:             "portRegex",
		CardModelValue: &opb.Port_CardModelRegex{"modelregex"},
		PmdValue:       &opb.Port_PmdRegex{".*"},
	}
	dut := &opb.Device{
		Id:                   "dut",
		Vendor:               opb.Device_ARISTA,
		HardwareModelValue:   &opb.Device_HardwareModel{"hw"},
		SoftwareVersionValue: &opb.Device_SoftwareVersion{"sw"},
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
		HardwareModelValue:   &opb.Device_HardwareModelRegex{"hwreg"},
		SoftwareVersionValue: &opb.Device_SoftwareVersionRegex{"swreg"},
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

	graph, node2Dev, port2Port, err := testbedToAbstractGraph(tb, nil)
	if err != nil {
		t.Fatalf("testbedToAbstractGraph() got error %v, want nil", err)
	}
	if wantNodes := 4; len(graph.Nodes) != wantNodes {
		t.Fatalf("testbedToAbstractGraph() got %d nodes, want %v", len(graph.Nodes), wantNodes)
	}
	for _, node := range graph.Nodes {
		if got, ok := node2Dev[node]; !ok {
			t.Errorf("testbedToAbstractGraph() got node %q not mapped to any device", node.Desc)
		} else if diff := cmp.Diff(wantDev[node.Desc], got, protocmp.Transform()); diff != "" {
			t.Errorf("testbedToAbstractGraph() returned unexpected device diff (-want +got):\n%s", diff)
		}
		for _, port := range node.Ports {
			if got, ok := port2Port[port]; !ok {
				t.Errorf("testbedToAbstractGraph() got port %q not mapped to any port", port.Desc)
			} else if diff := cmp.Diff(wantPort[port.Desc], got, protocmp.Transform()); diff != "" {
				t.Errorf("testbedToAbstractGraph() returned unexpected port diff (-want +got):\n%s", diff)
			}
		}
	}
}

func TestTopologyToConcreteGraph(t *testing.T) {
	intf1 := &tpb.Interface{
		Name:  "Ethernet1",
		Mtu:   1000,
		Group: "group1",
	}
	intf2 := &tpb.Interface{
		Name:  "eth2",
		Mtu:   2000,
		Group: "ategroup",
	}
	node := &tpb.Node{
		Name:       "node",
		Vendor:     tpb.Vendor_ARISTA,
		Model:      "ceos",
		Os:         "eos",
		Interfaces: map[string]*tpb.Interface{"eth1": intf1},
	}
	ate1 := &tpb.Node{
		Name:       "ate1",
		Vendor:     tpb.Vendor_KEYSIGHT,
		Interfaces: map[string]*tpb.Interface{"eth1": intf2},
	}
	ate2 := &tpb.Node{
		Name:   "ate2",
		Vendor: tpb.Vendor_OPENCONFIG,
		Labels: map[string]string{"ondatra-role": "ATE"},
	}
	topo := &tpb.Topology{
		Nodes: []*tpb.Node{node, ate1, ate2},
	}

	wantNode := map[string]*tpb.Node{"node": node, "ate1": ate1, "ate2": ate2}
	wantIntf := map[string]*tpb.Interface{"node:eth1": intf1, "ate1:eth1": intf2}

	graph, node2Node, _, err := topoToConcreteGraph(topo)
	if err != nil {
		t.Fatalf("topoToConcreteGraph() got error %v, want nil", err)
	}
	if len(graph.Nodes) != 3 {
		t.Fatalf("topoToConcreteGraph() got %d nodes, want 3", len(graph.Nodes))
	}
	for _, node := range graph.Nodes {
		got, ok := node2Node[node]
		if !ok {
			t.Errorf("topoToConcreteGraph() got node %q not mapped to any device", node.Desc)
		} else if diff := cmp.Diff(wantNode[node.Desc], got, protocmp.Transform()); diff != "" {
			t.Errorf("topoToConcreteGraph() returned unexpected device diff (-want +got):\n%s", diff)
		}
		if gotRole, wantRole := node.Attrs["role"], role(got); gotRole != wantRole {
			t.Errorf("node %q got role %q, want role %q", node.Desc, gotRole, wantRole)
		}
		for _, port := range node.Ports {
			wantIntf, ok := wantIntf[port.Desc]
			if !ok {
				t.Fatalf("topoToConcreteGraph() got unexpected port %q", port.Desc)
			}
			if gotMTU, wantMTU := port.Attrs["mtu"], strconv.FormatUint(uint64(wantIntf.GetMtu()), 10); gotMTU != wantMTU {
				t.Errorf("port %q got mtu %q, want mtu %q", port.Desc, gotMTU, wantMTU)
			}
			if gotGroup, wantGroup := port.Attrs["group"], wantIntf.GetGroup(); gotGroup != wantGroup {
				t.Errorf("port %q got group %q, want group %q", port.Desc, gotGroup, wantGroup)
			}
		}
	}
}

func TestRole(t *testing.T) {
	tests := []struct {
		desc string
		node *tpb.Node
		want string
	}{{
		desc: "default DUT",
		node: &tpb.Node{
			Name: "node",
		},
		want: roleDUT,
	}, {
		desc: "vendor DUT",
		node: &tpb.Node{
			Name:   "node",
			Vendor: tpb.Vendor_ARISTA,
		},
		want: roleDUT,
	}, {
		desc: "vendor ATE",
		node: &tpb.Node{
			Name:   "node",
			Vendor: tpb.Vendor_KEYSIGHT,
		},
		want: roleATE,
	}, {
		desc: "label ATE",
		node: &tpb.Node{
			Name:   "node",
			Vendor: tpb.Vendor_ARISTA,
			Labels: map[string]string{
				roleLabel: roleATE,
			},
		},
		want: roleATE,
	}, {
		desc: "label DUT",
		node: &tpb.Node{
			Name:   "node",
			Vendor: tpb.Vendor_KEYSIGHT,
			Labels: map[string]string{
				roleLabel: roleDUT,
			},
		},
		want: roleDUT,
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			if got := role(tc.node); got != tc.want {
				t.Errorf("role(%v) = %v, want %v", tc.node.GetName(), got, tc.want)
			}
		})
	}
}

func TestSolveErrors(t *testing.T) {
	tests := []struct {
		desc    string
		tb      *opb.Testbed
		topo    string
		wantErr string
	}{{
		desc:    "too few nodes",
		tb:      &opb.Testbed{Duts: []*opb.Device{{Id: "dut1"}}},
		wantErr: "not enough nodes",
	}, {
		desc: "too few nodes after filter",
		tb:   &opb.Testbed{Duts: []*opb.Device{{Id: "dut1"}}},
		topo: `
			nodes: {
			  name: "node1"
			  vendor: UNKNOWN
			}`,
		wantErr: "not enough nodes",
	}, {
		desc:    "too few links",
		tb:      &opb.Testbed{Links: []*opb.Link{{A: "dut1:port1", B: "dut2:port1"}}},
		wantErr: "not enough links",
	}, {
		desc:    "too few links after filter",
		tb:      &opb.Testbed{Links: []*opb.Link{{A: "dut1:port1", B: "dut2:port1"}}},
		wantErr: "not enough links",
		topo: `
			links: {
			  a_node: "node1"
			  a_int: "eth1"
			  z_node: "node2"
			  z_int: "eth1"
			}`,
	}, {
		desc: "no match for DUT",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
			}},
		},
		topo: `
			nodes: {
			  name: "node1"
			  vendor: CISCO
			}`,
		wantErr: "Node \"dut1\" was not assigned",
	}, {
		desc: "no match for DUT - wrong hardware model",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id:                 "dut1",
				Vendor:             opb.Device_ARISTA,
				HardwareModelValue: &opb.Device_HardwareModel{"ceos"},
			}},
		},
		topo: `
				nodes: {
					name: "node1"
					vendor: ARISTA
					model: "bad"
					os: "eos"
				}`,
		wantErr: "Node \"dut1\" was not assigned",
	}, {
		desc: "no match for DUT - wrong software model",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id:                   "dut1",
				Vendor:               opb.Device_ARISTA,
				SoftwareVersionValue: &opb.Device_SoftwareVersion{"eos"},
			}},
		},
		topo: `
					nodes: {
						name: "node1"
						vendor: ARISTA
						model: "ceos"
						os: "bad"
					}`,
		wantErr: "Node \"dut1\" was not assigned",
	}, {
		desc: "no match for DUT - role label override",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id: "dut1",
			}},
		},
		topo: `
				nodes: {
					name: "node1"
					vendor: ARISTA
					labels: {
						key: "ondatra-role"
						value: "ATE"
					}
				}`,
		wantErr: "Node \"dut1\" was not assigned",
	}, {
		desc: "no match for ATE",
		tb: &opb.Testbed{
			Ates: []*opb.Device{{
				Id: "ate1",
			}},
		},
		topo: `
			nodes: {
			  name: "node1"
			  vendor: CISCO
			}`,
		wantErr: "Node \"ate1\" was not assigned",
	}, {
		desc: "no match for ATE - role label override",
		tb: &opb.Testbed{
			Ates: []*opb.Device{{
				Id: "ate1",
			}},
		},
		topo: `
				nodes: {
					name: "node1"
					vendor: KEYSIGHT
					labels: {
						key: "ondatra-role"
						value: "DUT"
					}
				}`,
		wantErr: "Node \"ate1\" was not assigned",
	}, {
		desc: "no node combination",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
			}, {
				Id:     "dut2",
				Vendor: opb.Device_ARISTA,
			}},
		},
		topo: `
				nodes: {
					name: "node1"
					vendor: ARISTA
				}
				nodes: {
					name: "node2"
					vendor: CISCO
				}`,
		wantErr: "could not solve for specified testbed",
	}, {
		desc: "no link combination",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
				Ports:  []*opb.Port{{Id: "port1"}},
			}, {
				Id:     "dut2",
				Ports:  []*opb.Port{{Id: "port1"}},
				Vendor: opb.Device_ARISTA,
			}, {
				Id:     "dut3",
				Ports:  []*opb.Port{{Id: "port1"}},
				Vendor: opb.Device_JUNIPER,
			}, {
				Id:     "dut4",
				Ports:  []*opb.Port{{Id: "port1"}},
				Vendor: opb.Device_JUNIPER,
			}},
			Links: []*opb.Link{
				{A: "dut1:port1", B: "dut2:port1"},
				{A: "dut3:port1", B: "dut4:port1"},
			},
		},
		topo: `
			nodes: {
			  name: "node1"
			  vendor: ARISTA
			}
			nodes: {
			  name: "node2"
			  vendor: ARISTA
			}
			nodes: {
			  name: "node3"
			  vendor: JUNIPER
			}
			nodes: {
			  name: "node4"
			  vendor: JUNIPER
			}
			links: {
			  a_node: "node1"
			  a_int: "eth1"
			  z_node: "node3"
			  z_int: "eth1"
			}
			links: {
			  a_node: "node2"
			  a_int: "eth1"
			  z_node: "node4"
			  z_int: "eth1"
			}`,
		wantErr: "could not solve for specified testbed",
	}, {
		desc: "bad port group config in testbed; no groups in topology",
		tb: &opb.Testbed{
			Ates: []*opb.Device{{
				Id:    "ate1",
				Ports: []*opb.Port{{Id: "port1", Group: "G1"}, {Id: "port2", Group: "G1"}},
			}},
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
				Ports:  []*opb.Port{{Id: "port1"}},
			}, {
				Id:     "dut2",
				Ports:  []*opb.Port{{Id: "port1"}},
				Vendor: opb.Device_ARISTA,
			}},
			Links: []*opb.Link{
				{A: "dut1:port1", B: "ate1:port1"},
				{A: "dut2:port1", B: "ate1:port2"},
			},
		},
		topo: `
			nodes: {
			  name: "node1"
			  vendor: ARISTA
			}
			nodes: {
			  name: "node2"
			  vendor: KEYSIGHT
			}
			nodes: {
			  name: "node3"
			  vendor: ARISTA
			}
			links: {
			  a_node: "node1"
			  a_int: "eth1"
			  z_node: "node2"
			  z_int: "eth1"
			}
			links: {
			  a_node: "node2"
			  a_int: "eth2"
			  z_node: "node3"
			  z_int: "eth1"
			}`,
		wantErr: "could not solve for specified testbed",
	}, {
		desc: "required port group size not found",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
				Ports:  []*opb.Port{{Id: "port1", Group: "G1"}, {Id: "port2", Group: "G1"}},
			}, {
				Id:     "dut2",
				Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}},
				Vendor: opb.Device_ARISTA,
			}},
			Links: []*opb.Link{
				{A: "dut1:port1", B: "dut2:port1"},
				{A: "dut1:port2", B: "dut2:port2"},
			},
		},
		topo: `
			nodes: {
			  name: "node1"
			  vendor: ARISTA
			  interfaces: {
			    key: "eth1"
			  }
			  interfaces: {
			    key: "eth2"
			    value: {group: "lag"}
			  }
			}
			nodes: {
			  name: "node2"
			  vendor: ARISTA
			}
			links: {
			  a_node: "node1"
			  a_int: "eth1"
			  z_node: "node2"
			  z_int: "eth1"
			}
			links: {
			  a_node: "node1"
			  a_int: "eth2"
			  z_node: "node2"
			  z_int: "eth2"
			}`,
		wantErr: "could not solve for specified testbed",
	}, {
		desc: "required number of port groups not found",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
				Ports:  []*opb.Port{{Id: "port1", Group: "G1"}, {Id: "port2", Group: "G2"}},
			}, {
				Id:     "dut2",
				Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}},
				Vendor: opb.Device_ARISTA,
			}},
			Links: []*opb.Link{
				{A: "dut1:port1", B: "dut2:port1"},
				{A: "dut1:port2", B: "dut2:port2"},
			},
		},
		topo: `
			nodes: {
			  name: "node1"
			  vendor: ARISTA
			  interfaces: {
			    key: "eth1"
			  }
			  interfaces: {
			    key: "eth2"
			    value: {group: "lag"}
			  }
			}
			nodes: {
			  name: "node2"
			  vendor: ARISTA
			}
			links: {
			  a_node: "node1"
			  a_int: "eth1"
			  z_node: "node2"
			  z_int: "eth1"
			}
			links: {
			  a_node: "node1"
			  a_int: "eth2"
			  z_node: "node2"
			  z_int: "eth2"
			}`,
		wantErr: "could not solve for specified testbed",
	}, {
		desc: "dut without vendor does not match ate",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{Id: "dut"}},
		},
		topo: `
			nodes: {
			  name: "node"
			  vendor: KEYSIGHT
			}`,
		wantErr: "could not solve for specified testbed",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			topo := unmarshalTopo(t, test.topo)
			_, gotErr := Solve(test.tb, topo, nil)
			if diff := errdiff.Substring(gotErr, test.wantErr); diff != "" {
				t.Fatalf("Reserve() got unexpected error diff: %s", diff)
			}
		})
	}
}

func unmarshalTopo(t testing.TB, text string) *tpb.Topology {
	topo := new(tpb.Topology)
	if err := prototext.Unmarshal([]byte(text), topo); err != nil {
		t.Fatalf("Error unmarshalling topology text: %v", err)
	}
	return topo
}
