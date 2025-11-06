// Copyright 2025 Google LLC
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
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/binding/portgraph"
	opb "github.com/openconfig/ondatra/proto"
)

// Mock DUT implementation for testing
type mockDUT struct {
	*binding.AbstractDUT
}

// Mock ATE implementation for testing
type mockATE struct {
	*binding.AbstractATE
}

// TestInventoryToConcreteGraph tests the conversion of an Inventory to a ConcreteGraph, including link canonicalization and edge deduplication.
func TestInventoryToConcreteGraph(t *testing.T) {
	// Shared devices and ports for test cases
	// For original test case (node/port attributes)
	dutPort1Orig := &binding.Port{Name: "Ethernet1", Speed: opb.Port_S_10GB}
	atePort1Orig := &binding.Port{Name: "eth2", Speed: opb.Port_S_10GB}
	dut1Orig := &mockDUT{
		AbstractDUT: &binding.AbstractDUT{
			Dims: &binding.Dims{
				Name: "node", Vendor: opb.Device_ARISTA, HardwareModel: "ceos", SoftwareVersion: "eos",
				Ports: map[string]*binding.Port{"port1": dutPort1Orig},
			},
		},
	}
	ate1Orig := &mockATE{
		AbstractATE: &binding.AbstractATE{
			Dims: &binding.Dims{
				Name: "ate1", Vendor: opb.Device_IXIA,
				Ports: map[string]*binding.Port{"port1": atePort1Orig},
			},
		},
	}
	ate2Orig := &mockATE{
		AbstractATE: &binding.AbstractATE{
			Dims: &binding.Dims{Name: "ate2", Vendor: opb.Device_OPENCONFIG, Ports: make(map[string]*binding.Port)},
		},
	}

	// For link canonicalization test cases
	linkDut1P1 := &binding.Port{Name: "port1"}
	linkDut1 := &mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1", Ports: map[string]*binding.Port{"port1": linkDut1P1}}}}
	linkDut2P1 := &binding.Port{Name: "port1"}
	linkDut2 := &mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut2", Ports: map[string]*binding.Port{"port1": linkDut2P1}}}}
	linkDut3P1 := &binding.Port{Name: "port1"}
	linkDut3 := &mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut3", Ports: map[string]*binding.Port{"port1": linkDut3P1}}}}

	// For error cases
	errLinkDut1P1 := &binding.Port{Name: "port1"}
	errLinkDut1 := &mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1", Ports: map[string]*binding.Port{"port1": errLinkDut1P1}}}}
	nonExistentPort := &binding.Port{Name: "nonexistent"}

	// Shared concrete ports and nodes for expected graphs
	cpDut1Link := &portgraph.ConcretePort{Desc: "dut1:port1", Attrs: map[string]string{"name": "port1", "speed": "SPEED_UNSPECIFIED", "pmd": "PMD_UNSPECIFIED"}}
	cnDut1Link := &portgraph.ConcreteNode{Desc: "dut1", Ports: []*portgraph.ConcretePort{cpDut1Link}, Attrs: map[string]string{"role": "DUT", "vendor": "VENDOR_UNSPECIFIED", "hardware_model": "", "software_version": "", "name": "dut1"}}
	cpDut2Link := &portgraph.ConcretePort{Desc: "dut2:port1", Attrs: map[string]string{"name": "port1", "speed": "SPEED_UNSPECIFIED", "pmd": "PMD_UNSPECIFIED"}}
	cnDut2Link := &portgraph.ConcreteNode{Desc: "dut2", Ports: []*portgraph.ConcretePort{cpDut2Link}, Attrs: map[string]string{"role": "DUT", "vendor": "VENDOR_UNSPECIFIED", "hardware_model": "", "software_version": "", "name": "dut2"}}
	cpDut3Link := &portgraph.ConcretePort{Desc: "dut3:port1", Attrs: map[string]string{"name": "port1", "speed": "SPEED_UNSPECIFIED", "pmd": "PMD_UNSPECIFIED"}}
	cnDut3Link := &portgraph.ConcreteNode{Desc: "dut3", Ports: []*portgraph.ConcretePort{cpDut3Link}, Attrs: map[string]string{"role": "DUT", "vendor": "VENDOR_UNSPECIFIED", "hardware_model": "", "software_version": "", "name": "dut3"}}

	tests := []struct {
		desc        string
		inv         *Inventory
		wantGraph   *portgraph.ConcreteGraph
		wantNodeMap map[string]string
		wantPortMap map[string]string
		wantErr     bool
	}{{
		desc: "original node and port attributes",
		inv: &Inventory{
			DUTs:  []binding.DUT{dut1Orig},
			ATEs:  []binding.ATE{ate1Orig, ate2Orig},
			Links: nil,
		},
		wantGraph: &portgraph.ConcreteGraph{
			Desc: "Generic Inventory",
			Nodes: []*portgraph.ConcreteNode{{
				Desc:  "node",
				Ports: []*portgraph.ConcretePort{{Desc: "node:Ethernet1", Attrs: map[string]string{"name": "Ethernet1", "speed": "S_10GB", "pmd": "PMD_UNSPECIFIED"}}},
				Attrs: map[string]string{"role": "DUT", "vendor": "ARISTA", "hardware_model": "ceos", "software_version": "eos", "name": "node"},
			}, {
				Desc:  "ate1",
				Ports: []*portgraph.ConcretePort{{Desc: "ate1:eth2", Attrs: map[string]string{"name": "eth2", "speed": "S_10GB", "pmd": "PMD_UNSPECIFIED"}}},
				Attrs: map[string]string{"role": "ATE", "vendor": "IXIA", "hardware_model": "", "software_version": "", "name": "ate1"},
			}, {
				Desc:  "ate2",
				Ports: []*portgraph.ConcretePort{},
				Attrs: map[string]string{"role": "ATE", "vendor": "OPENCONFIG", "hardware_model": "", "software_version": "", "name": "ate2"},
			}},
			Edges: []*portgraph.ConcreteEdge{},
		},
		wantNodeMap: map[string]string{"node": "node", "ate1": "ate1", "ate2": "ate2"},
		wantPortMap: map[string]string{"node:Ethernet1": "Ethernet1", "ate1:eth2": "eth2"},
	}, {
		desc: "link from non-existent port",
		inv: &Inventory{
			DUTs:  []binding.DUT{errLinkDut1},
			Links: map[*binding.Port]*binding.Port{nonExistentPort: errLinkDut1P1},
		},
		wantErr: true,
	}, {
		desc: "link to non-existent port",
		inv: &Inventory{
			DUTs:  []binding.DUT{errLinkDut1},
			Links: map[*binding.Port]*binding.Port{errLinkDut1P1: nonExistentPort},
		},
		wantErr: true,
	},
		{
			desc: "device with nil port",
			inv: &Inventory{
				DUTs: []binding.DUT{
					&mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{
						Name:  "dut1",
						Ports: map[string]*binding.Port{"port1": nil},
					}}},
				},
			},
			wantErr: true,
		}, {
			desc: "link order preserved",
			inv: &Inventory{
				DUTs:  []binding.DUT{linkDut1, linkDut2},
				Links: map[*binding.Port]*binding.Port{linkDut1P1: linkDut2P1},
			},
			wantGraph: &portgraph.ConcreteGraph{
				Desc:  "Generic Inventory",
				Nodes: []*portgraph.ConcreteNode{cnDut1Link, cnDut2Link},
				Edges: []*portgraph.ConcreteEdge{{Src: cpDut1Link, Dst: cpDut2Link}},
			},
			wantNodeMap: map[string]string{"dut1": "dut1", "dut2": "dut2"},
			wantPortMap: map[string]string{"dut1:port1": "port1", "dut2:port1": "port1"},
		}, {
			desc: "link order swapped",
			inv: &Inventory{
				DUTs:  []binding.DUT{linkDut1, linkDut2},
				Links: map[*binding.Port]*binding.Port{linkDut2P1: linkDut1P1},
			},
			wantGraph: &portgraph.ConcreteGraph{
				Desc:  "Generic Inventory",
				Nodes: []*portgraph.ConcreteNode{cnDut1Link, cnDut2Link},
				Edges: []*portgraph.ConcreteEdge{{Src: cpDut1Link, Dst: cpDut2Link}},
			},
			wantNodeMap: map[string]string{"dut1": "dut1", "dut2": "dut2"},
			wantPortMap: map[string]string{"dut1:port1": "port1", "dut2:port1": "port1"},
		}, {
			desc: "redundant links",
			inv: &Inventory{
				DUTs:  []binding.DUT{linkDut1, linkDut2},
				Links: map[*binding.Port]*binding.Port{linkDut1P1: linkDut2P1, linkDut2P1: linkDut1P1},
			},
			wantGraph: &portgraph.ConcreteGraph{
				Desc:  "Generic Inventory",
				Nodes: []*portgraph.ConcreteNode{cnDut1Link, cnDut2Link},
				Edges: []*portgraph.ConcreteEdge{{Src: cpDut1Link, Dst: cpDut2Link}},
			},
			wantNodeMap: map[string]string{"dut1": "dut1", "dut2": "dut2"},
			wantPortMap: map[string]string{"dut1:port1": "port1", "dut2:port1": "port1"},
		}, {
			desc: "multiple links",
			inv: &Inventory{
				DUTs:  []binding.DUT{linkDut1, linkDut2, linkDut3},
				Links: map[*binding.Port]*binding.Port{linkDut1P1: linkDut2P1, linkDut2P1: linkDut3P1},
			},
			wantGraph: &portgraph.ConcreteGraph{
				Desc:  "Generic Inventory",
				Nodes: []*portgraph.ConcreteNode{cnDut1Link, cnDut2Link, cnDut3Link},
				Edges: []*portgraph.ConcreteEdge{{Src: cpDut1Link, Dst: cpDut2Link}, {Src: cpDut2Link, Dst: cpDut3Link}},
			},
			wantNodeMap: map[string]string{"dut1": "dut1", "dut2": "dut2", "dut3": "dut3"},
			wantPortMap: map[string]string{"dut1:port1": "port1", "dut2:port1": "port1", "dut3:port1": "port1"},
		}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotGraph, node2BindDev, port2BindPort, err := inventoryToConcreteGraph(test.inv)
			if test.wantErr {
				if err == nil {
					t.Fatal("inventoryToConcreteGraph() got no error, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("inventoryToConcreteGraph() got unexpected error: %v", err)
			}

			sortNodes := cmpopts.SortSlices(func(a, b *portgraph.ConcreteNode) bool { return a.Desc < b.Desc })
			sortEdges := cmpopts.SortSlices(func(a, b *portgraph.ConcreteEdge) bool { return a.Src.Desc < b.Src.Desc })
			if diff := cmp.Diff(test.wantGraph, gotGraph, sortNodes, sortEdges, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("inventoryToConcreteGraph() returned diff in graph (-want +got):\n%s", diff)
			}

			gotNodeMap := make(map[string]string)
			for node, dev := range node2BindDev {
				gotNodeMap[node.Desc] = (*dev).Name()
			}
			if diff := cmp.Diff(test.wantNodeMap, gotNodeMap); diff != "" {
				t.Errorf("inventoryToConcreteGraph() returned diff in node map (-want +got):\n%s", diff)
			}

			gotPortMap := make(map[string]string)
			for port, bindPort := range port2BindPort {
				gotPortMap[port.Desc] = bindPort.Name
			}
			if diff := cmp.Diff(test.wantPortMap, gotPortMap); diff != "" {
				t.Errorf("inventoryToConcreteGraph() returned diff in port map (-want +got):\n%s", diff)
			}
		})
	}
}

// TestSolve tests the mapping of a testbed to a concrete graph with a variety of test cases, using the Solve function.
func TestSolve(t *testing.T) {
	// Build inventory from topology:
	// Define ports first so they can be reused in device definitions and the links map.
	var (
		portNode1Eth1 = binding.Port{Name: "Ethernet1"}
		portNode1Eth2 = binding.Port{Name: "Ethernet2"}
		portNode2Eth1 = binding.Port{Name: "GigabitEthernet0/0/0/0"}
		portNode2Eth2 = binding.Port{Name: "eth2"}
		portNode3Eth1 = binding.Port{Name: "eth1"}
		portNode4Eth1 = binding.Port{Name: "eth1"}
		portNode5Eth1 = binding.Port{Name: "eth1"}
	)

	inventory := &Inventory{
		DUTs: []binding.DUT{
			&mockDUT{
				AbstractDUT: &binding.AbstractDUT{
					Dims: &binding.Dims{
						Name:            "node1",
						Vendor:          opb.Device_ARISTA,
						HardwareModel:   "ceos",
						SoftwareVersion: "eos",
						Ports: map[string]*binding.Port{
							"port1": &portNode1Eth1,
							"port2": &portNode1Eth2,
						},
					},
				},
			},
			&mockDUT{
				AbstractDUT: &binding.AbstractDUT{
					Dims: &binding.Dims{
						Name:   "node2",
						Vendor: opb.Device_CISCO,
						Ports: map[string]*binding.Port{
							"port1": &portNode2Eth1,
							"port2": &portNode2Eth2,
						},
					},
				},
			},
			&mockDUT{
				AbstractDUT: &binding.AbstractDUT{
					Dims: &binding.Dims{
						Name:            "node3",
						Vendor:          opb.Device_JUNIPER,
						HardwareModel:   "cptx",
						SoftwareVersion: "evo",
						Ports: map[string]*binding.Port{
							"port1": &portNode3Eth1,
						},
					},
				},
			},
		},
		ATEs: []binding.ATE{
			&mockATE{
				AbstractATE: &binding.AbstractATE{
					Dims: &binding.Dims{
						Name:   "node4",
						Vendor: opb.Device_IXIA,
						Ports: map[string]*binding.Port{
							"port1": &portNode4Eth1,
						},
					},
				},
			},
			&mockATE{
				AbstractATE: &binding.AbstractATE{
					Dims: &binding.Dims{
						Name:          "node5",
						Vendor:        opb.Device_OPENCONFIG,
						HardwareModel: "MAGNA",
						Ports: map[string]*binding.Port{
							"port1": &portNode5Eth1,
						},
					},
				},
			},
		},
		Links: map[*binding.Port]*binding.Port{
			&portNode1Eth1: &portNode2Eth1,
			&portNode2Eth2: &portNode3Eth1,
			&portNode1Eth2: &portNode5Eth1,
		},
	}

	tests := []struct {
		desc             string
		tb               *opb.Testbed
		inv              *Inventory
		partial          map[string]string
		wantDeviceAssign map[string]string
		wantPortAssign   map[string]string
		wantErr          bool
	}{
		{
			desc: "one dut",
			tb: &opb.Testbed{
				Duts: []*opb.Device{{
					Id:     "dut",
					Vendor: opb.Device_JUNIPER,
					HardwareModelValue: &opb.Device_HardwareModel{
						HardwareModel: "cptx",
					},
					SoftwareVersionValue: &opb.Device_SoftwareVersion{
						SoftwareVersion: "evo",
					},
					Ports: []*opb.Port{{Id: "port1"}},
				}},
			},
			inv: inventory,
			wantDeviceAssign: map[string]string{
				"dut": "node3",
			},
			wantPortAssign: map[string]string{
				"dut:port1": "node3:eth1",
			},
		}, {
			desc: "one dut with partial",
			tb: &opb.Testbed{
				Duts: []*opb.Device{{
					Id:     "dut",
					Vendor: opb.Device_JUNIPER,
					HardwareModelValue: &opb.Device_HardwareModel{
						HardwareModel: "cptx",
					},
					SoftwareVersionValue: &opb.Device_SoftwareVersion{
						SoftwareVersion: "evo",
					},
					Ports: []*opb.Port{{Id: "port1"}},
				}},
			},
			inv: inventory,
			partial: map[string]string{
				"dut": "node3",
			},
			wantDeviceAssign: map[string]string{
				"dut": "node3",
			},
			wantPortAssign: map[string]string{
				"dut:port1": "node3:eth1",
			},
		}, {
			desc: "one dut with partial + port intf name",
			tb: &opb.Testbed{
				Duts: []*opb.Device{{
					Id:     "dut",
					Vendor: opb.Device_CISCO,
					Ports:  []*opb.Port{{Id: "port1"}},
				}},
			},
			inv: inventory,
			partial: map[string]string{
				"dut":       "node2",
				"dut:port1": "eth2",
			},
			wantDeviceAssign: map[string]string{
				"dut": "node2",
			},
			wantPortAssign: map[string]string{
				"dut:port1": "node2:eth2",
			},
		}, {
			desc: "one dut with partial + port vendor name",
			tb: &opb.Testbed{
				Duts: []*opb.Device{{
					Id:     "dut",
					Vendor: opb.Device_CISCO,
					Ports:  []*opb.Port{{Id: "port1"}},
				}},
			},
			inv: inventory,
			partial: map[string]string{
				"dut":       "node2",
				"dut:port1": "GigabitEthernet0/0/0/0",
			},
			wantDeviceAssign: map[string]string{
				"dut": "node2",
			},
			wantPortAssign: map[string]string{
				"dut:port1": "node2:GigabitEthernet0/0/0/0",
			},
		}, {
			desc: "one ate",
			tb: &opb.Testbed{
				Ates: []*opb.Device{{
					Id:     "ate1",
					Vendor: opb.Device_IXIA,
					Ports:  []*opb.Port{{Id: "port1"}},
				}},
			},
			inv: inventory,
			wantDeviceAssign: map[string]string{
				"ate1": "node4",
			},
			wantPortAssign: map[string]string{
				"ate1:port1": "node4:eth1",
			},
		}, {
			desc: "two ates",
			tb: &opb.Testbed{
				Ates: []*opb.Device{{
					Id:     "ate1",
					Vendor: opb.Device_IXIA,
					Ports:  []*opb.Port{{Id: "port1"}},
				}, {
					Id:     "ate2",
					Vendor: opb.Device_OPENCONFIG,
				}},
			},
			inv: inventory,
			wantDeviceAssign: map[string]string{
				"ate1": "node4",
				"ate2": "node5",
			},
			wantPortAssign: map[string]string{
				"ate1:port1": "node4:eth1",
			},
		}, {
			desc: "two duts",
			tb: &opb.Testbed{
				Duts: []*opb.Device{{
					Id:     "dut1",
					Vendor: opb.Device_ARISTA,
					HardwareModelValue: &opb.Device_HardwareModel{
						HardwareModel: "ceos",
					},
					SoftwareVersionValue: &opb.Device_SoftwareVersion{
						SoftwareVersion: "eos",
					},
					Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}},
				}, {
					Id:     "dut2",
					Vendor: opb.Device_CISCO,
					Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}},
				}},
				Links: []*opb.Link{{
					A: "dut1:port1",
					B: "dut2:port1",
				}},
			},
			inv: inventory,
			wantDeviceAssign: map[string]string{
				"dut1": "node1",
				"dut2": "node2",
			},
			wantPortAssign: map[string]string{
				"dut1:port1": "node1:Ethernet1",
				"dut2:port1": "node2:GigabitEthernet0/0/0/0",
				"dut1:port2": "node1:Ethernet2",
				"dut2:port2": "node2:eth2",
			},
		}, {
			desc: "dut and ate",
			tb: &opb.Testbed{
				Duts: []*opb.Device{{
					Id:     "dut1",
					Vendor: opb.Device_ARISTA,
					HardwareModelValue: &opb.Device_HardwareModel{
						HardwareModel: "ceos",
					},
					SoftwareVersionValue: &opb.Device_SoftwareVersion{
						SoftwareVersion: "eos",
					},
					Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}},
				}},
				Ates: []*opb.Device{{
					Id:    "ate1",
					Ports: []*opb.Port{{Id: "port1"}},
				}},
				Links: []*opb.Link{{
					A: "dut1:port2",
					B: "ate1:port1",
				}},
			},
			inv: inventory,
			wantDeviceAssign: map[string]string{
				"dut1": "node1",
				"ate1": "node5",
			},
			wantPortAssign: map[string]string{
				"dut1:port1": "node1:Ethernet1",
				"dut1:port2": "node1:Ethernet2",
				"ate1:port1": "node5:eth1",
			},
		}, {
			desc: "three duts",
			tb: &opb.Testbed{
				Duts: []*opb.Device{{
					Id:     "dut1",
					Vendor: opb.Device_ARISTA,
					Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}},
				}, {
					Id:     "dut2",
					Vendor: opb.Device_CISCO,
					Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}},
				}, {
					Id:     "dut3",
					Vendor: opb.Device_JUNIPER,
					Ports:  []*opb.Port{{Id: "port1"}},
				}},
				Links: []*opb.Link{{
					A: "dut1:port1",
					B: "dut2:port1",
				}, {
					A: "dut2:port2",
					B: "dut3:port1",
				}},
			},
			inv: inventory,
			wantDeviceAssign: map[string]string{
				"dut1": "node1",
				"dut2": "node2",
				"dut3": "node3",
			},
			wantPortAssign: map[string]string{
				"dut1:port1": "node1:Ethernet1",
				"dut1:port2": "node1:Ethernet2",
				"dut2:port1": "node2:GigabitEthernet0/0/0/0",
				"dut2:port2": "node2:eth2",
				"dut3:port1": "node3:eth1",
			},
		}, {
			desc: "exact solve",
			tb: &opb.Testbed{
				Duts: []*opb.Device{
					{Id: "r1", Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}}},
					{Id: "a2", Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}}},
					{Id: "r_3", Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}}},
					{Id: "r4", Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}}},
				},
				Ates: []*opb.Device{
					{Id: "r5", Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}}},
				},
				Links: []*opb.Link{
					{A: "r1:port1", B: "a2:port2"},
					{A: "a2:port1", B: "r_3:port2"},
					{A: "r_3:port1", B: "r4:port2"},
					{A: "r4:port1", B: "r5:port2"},
					{A: "r5:port1", B: "r1:port2"},
				},
			},
			inv: func() *Inventory {
				var (
					r1eth1 = binding.Port{Name: "eth1"}
					r1eth2 = binding.Port{Name: "eth2"}
					r2eth1 = binding.Port{Name: "eth1"}
					r2eth2 = binding.Port{Name: "eth2"}
					r3eth1 = binding.Port{Name: "eth1"}
					r3eth2 = binding.Port{Name: "eth2"}
					r4eth1 = binding.Port{Name: "eth1"}
					r4eth2 = binding.Port{Name: "eth2"}
					r5eth1 = binding.Port{Name: "eth1"}
					r5eth2 = binding.Port{Name: "eth2"}
				)
				return &Inventory{
					DUTs: []binding.DUT{
						&mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "r1", Vendor: opb.Device_OPENCONFIG, HardwareModel: "LEMMING", Ports: map[string]*binding.Port{"port1": &r1eth1, "port2": &r1eth2}}}},
						&mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "r2", Vendor: opb.Device_OPENCONFIG, HardwareModel: "LEMMING", Ports: map[string]*binding.Port{"port1": &r2eth1, "port2": &r2eth2}}}},
						&mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "r-3", Vendor: opb.Device_OPENCONFIG, HardwareModel: "LEMMING", Ports: map[string]*binding.Port{"port1": &r3eth1, "port2": &r3eth2}}}},
						&mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "r4", Vendor: opb.Device_OPENCONFIG, HardwareModel: "LEMMING", Ports: map[string]*binding.Port{"port1": &r4eth1, "port2": &r4eth2}}}},
					},
					ATEs: []binding.ATE{
						&mockATE{AbstractATE: &binding.AbstractATE{Dims: &binding.Dims{Name: "r5", Vendor: opb.Device_IXIA, Ports: map[string]*binding.Port{"port1": &r5eth1, "port2": &r5eth2}}}},
					},
					Links: map[*binding.Port]*binding.Port{
						&r1eth1: &r2eth2,
						&r2eth1: &r3eth2,
						&r3eth1: &r4eth2,
						&r4eth1: &r5eth2,
						&r5eth1: &r1eth2,
					},
				}
			}(),
			partial: map[string]string{
				"r1":  "r1",
				"a2":  "r2",
				"r_3": "r-3",
				"r4":  "r4",
				"r5":  "r5",
			},
			wantDeviceAssign: map[string]string{
				"r1":  "r1",
				"a2":  "r2",
				"r_3": "r-3",
				"r4":  "r4",
				"r5":  "r5",
			},
			wantPortAssign: map[string]string{
				"r1:port1":  "r1:eth1",
				"r1:port2":  "r1:eth2",
				"a2:port1":  "r2:eth1",
				"a2:port2":  "r2:eth2",
				"r_3:port1": "r-3:eth1",
				"r_3:port2": "r-3:eth2",
				"r4:port1":  "r4:eth1",
				"r4:port2":  "r4:eth2",
				"r5:port1":  "r5:eth1",
				"r5:port2":  "r5:eth2",
			},
		},
		{
			desc: "no solution",
			tb: &opb.Testbed{
				Duts: []*opb.Device{{
					Id:    "dut1",
					Ports: []*opb.Port{{Id: "port1"}},
				}, {
					Id:    "dut2",
					Ports: []*opb.Port{{Id: "port1"}},
				}},
				Links: []*opb.Link{{
					A: "dut1:port1",
					B: "dut2:port1",
				}},
			},
			inv: &Inventory{
				DUTs: []binding.DUT{
					&mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "nodeA", Ports: map[string]*binding.Port{"p1": {Name: "eth1"}}}}},
					&mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "nodeB", Ports: map[string]*binding.Port{"p1": {Name: "eth1"}}}}},
				},
			},
			wantErr: true,
		},
		{
			desc: "inventoryToConcreteGraph error",
			tb:   &opb.Testbed{},
			inv: &Inventory{
				Links: map[*binding.Port]*binding.Port{
					&binding.Port{Name: "p1"}: &binding.Port{Name: "p2"},
				},
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotRes, err := Solve(context.Background(), test.tb, test.inv, test.partial)
			if test.wantErr {
				if err == nil {
					t.Fatal("Solve() got no error, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("Solve() got error %v, want nil", err)
			}
			assignment := gotRes.Assignment
			absNode2Dev := gotRes.AbsNode2Dev
			absPort2Port := gotRes.AbsPort2Port
			conNode2BindDev := gotRes.ConNode2BindDev
			conPort2BindPort := gotRes.ConPort2BindPort
			if assignment == nil {
				t.Fatalf("Solve() returned no assignment, but expected one.")
			}

			gotDeviceAssign := make(map[string]string)
			for absNode, conNode := range assignment.Node2Node {
				gotDeviceAssign[absNode2Dev[absNode].GetId()] = (*conNode2BindDev[conNode]).Name()
			}
			gotPortAssign := make(map[string]string)

			findAbsNode := func(targetPort *portgraph.AbstractPort, nodeMap map[*portgraph.AbstractNode]*opb.Device) *portgraph.AbstractNode {
				for node := range nodeMap {
					for _, p := range node.Ports {
						if p == targetPort {
							return node
						}
					}
				}
				return nil
			}

			findConNode := func(targetPort *portgraph.ConcretePort, nodeMap map[*portgraph.ConcreteNode]*binding.Device) *portgraph.ConcreteNode {
				for node := range nodeMap {
					for _, p := range node.Ports {
						if p == targetPort {
							return node
						}
					}
				}
				return nil
			}

			for absPort, conPort := range assignment.Port2Port {
				absNode := findAbsNode(absPort, absNode2Dev)
				if absNode == nil {
					t.Errorf("FATAL: absPort %s not found in any absNode", absPort.Desc)
					continue
				}
				tbDevID := absNode2Dev[absNode].GetId()
				tbPortID := absPort2Port[absPort].GetId()

				conNode := findConNode(conPort, conNode2BindDev)
				if conNode == nil {
					t.Errorf("FATAL: conPort %s not found in any conNode", conPort.Desc)
					continue
				}
				bindDevName := (*conNode2BindDev[conNode]).Name()
				bindPortName := conPort2BindPort[conPort].Name
				gotPortAssign[fmt.Sprintf("%s:%s", tbDevID, tbPortID)] = fmt.Sprintf("%s:%s", bindDevName, bindPortName)
			}
			if diff := cmp.Diff(test.wantDeviceAssign, gotDeviceAssign); diff != "" {
				t.Errorf("Solve() returned diff in device assignments (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(test.wantPortAssign, gotPortAssign); diff != "" {
				t.Errorf("Solve() returned diff in port assignments (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAssignmentToReservation(t *testing.T) {
	tDUT1 := &opb.Device{Id: "dut1", Ports: []*opb.Port{{Id: "port1"}}}
	tATE1 := &opb.Device{Id: "ate1", Ports: []*opb.Port{{Id: "port1"}}}
	tb := &opb.Testbed{
		Duts:  []*opb.Device{tDUT1},
		Ates:  []*opb.Device{tATE1},
		Links: []*opb.Link{{A: "dut1:port1", B: "ate1:port1"}},
	}

	bDUT1Port1 := &binding.Port{Name: "eth1"}
	bATE1Port1 := &binding.Port{Name: "eth1"}
	invDUT1 := &mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "inventory_dut1", Ports: map[string]*binding.Port{"p1": bDUT1Port1}}}}
	invATE1 := &mockATE{AbstractATE: &binding.AbstractATE{Dims: &binding.Dims{Name: "inventory_ate1", Ports: map[string]*binding.Port{"p1": bATE1Port1}}}}
	bdevDUT1 := binding.Device(invDUT1)
	bdevATE1 := binding.Device(invATE1)

	absDUT1Node := &portgraph.AbstractNode{Desc: "dut1"}
	absATE1Node := &portgraph.AbstractNode{Desc: "ate1"}
	absDUT1Port1 := &portgraph.AbstractPort{Desc: "dut1:port1"}
	absATE1Port1 := &portgraph.AbstractPort{Desc: "ate1:port1"}

	conDUT1Node := &portgraph.ConcreteNode{Desc: invDUT1.Name()}
	conATE1Node := &portgraph.ConcreteNode{Desc: invATE1.Name()}
	conDUT1Port1 := &portgraph.ConcretePort{Desc: invDUT1.Name() + ":" + bDUT1Port1.Name}
	conATE1Port1 := &portgraph.ConcretePort{Desc: invATE1.Name() + ":" + bATE1Port1.Name}

	solveResult := &SolveResult{
		Assignment: &portgraph.Assignment{
			Node2Node: map[*portgraph.AbstractNode]*portgraph.ConcreteNode{
				absDUT1Node: conDUT1Node,
				absATE1Node: conATE1Node,
			},
			Port2Port: map[*portgraph.AbstractPort]*portgraph.ConcretePort{
				absDUT1Port1: conDUT1Port1,
				absATE1Port1: conATE1Port1,
			},
		},
		AbsNode2Dev: map[*portgraph.AbstractNode]*opb.Device{
			absDUT1Node: tDUT1,
			absATE1Node: tATE1,
		},
		AbsPort2Port: map[*portgraph.AbstractPort]*opb.Port{
			absDUT1Port1: tDUT1.Ports[0],
			absATE1Port1: tATE1.Ports[0],
		},
		ConNode2BindDev: map[*portgraph.ConcreteNode]*binding.Device{
			conDUT1Node: &bdevDUT1,
			conATE1Node: &bdevATE1,
		},
		ConPort2BindPort: map[*portgraph.ConcretePort]*binding.Port{
			conDUT1Port1: bDUT1Port1,
			conATE1Port1: bATE1Port1,
		},
	}

	res, err := AssignmentToReservation(solveResult, tb)
	if err != nil {
		t.Fatalf("AssignmentToReservation() got unexpected error: %v", err)
	}

	wantRes := &binding.Reservation{
		ID: res.ID,
		DUTs: map[string]binding.DUT{
			"dut1": &binding.AbstractDUT{
				Dims: &binding.Dims{
					Name: "inventory_dut1",
					Ports: map[string]*binding.Port{
						"port1": bDUT1Port1,
					},
				},
			},
		},
		ATEs: map[string]binding.ATE{
			"ate1": &binding.AbstractATE{
				Dims: &binding.Dims{
					Name: "inventory_ate1",
					Ports: map[string]*binding.Port{
						"port1": bATE1Port1,
					},
				},
			},
		},
	}
	if diff := cmp.Diff(wantRes, res); diff != "" {
		t.Errorf("AssignmentToReservation() returned diff (-want +got):\n%s", diff)
	}
}

func TestAssignmentToReservationErrors(t *testing.T) {
	tDUT1 := &opb.Device{Id: "dut1", Ports: []*opb.Port{{Id: "port1"}}}
	tATE1 := &opb.Device{Id: "ate1", Ports: []*opb.Port{{Id: "port1"}}}
	bDUT1Port1 := &binding.Port{Name: "eth1"}
	bATE1Port1 := &binding.Port{Name: "eth1"}
	invDUT1 := &mockDUT{AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "inventory_dut1", Ports: map[string]*binding.Port{"p1": bDUT1Port1}}}}
	invATE1 := &mockATE{AbstractATE: &binding.AbstractATE{Dims: &binding.Dims{Name: "inventory_ate1", Ports: map[string]*binding.Port{"p1": bATE1Port1}}}}
	bdevDUT1 := binding.Device(invDUT1)
	bdevATE1 := binding.Device(invATE1)
	absDUT1Node := &portgraph.AbstractNode{Desc: "dut1"}
	absATE1Node := &portgraph.AbstractNode{Desc: "ate1"}
	conDUT1Node := &portgraph.ConcreteNode{Desc: invDUT1.Name()}
	conATE1Node := &portgraph.ConcreteNode{Desc: invATE1.Name()}

	tests := []struct {
		desc        string
		tb          *opb.Testbed
		solveResult *SolveResult
		wantErr     string
	}{
		{
			desc: "dut not resolved",
			tb:   &opb.Testbed{Duts: []*opb.Device{tDUT1}},
			solveResult: &SolveResult{
				Assignment:       &portgraph.Assignment{},
				AbsNode2Dev:      map[*portgraph.AbstractNode]*opb.Device{},
				AbsPort2Port:     map[*portgraph.AbstractPort]*opb.Port{},
				ConNode2BindDev:  map[*portgraph.ConcreteNode]*binding.Device{},
				ConPort2BindPort: map[*portgraph.ConcretePort]*binding.Port{},
			},
			wantErr: "not resolved",
		},
		{
			desc: "ate not resolved",
			tb:   &opb.Testbed{Ates: []*opb.Device{tATE1}},
			solveResult: &SolveResult{
				Assignment:       &portgraph.Assignment{},
				AbsNode2Dev:      map[*portgraph.AbstractNode]*opb.Device{},
				AbsPort2Port:     map[*portgraph.AbstractPort]*opb.Port{},
				ConNode2BindDev:  map[*portgraph.ConcreteNode]*binding.Device{},
				ConPort2BindPort: map[*portgraph.ConcretePort]*binding.Port{},
			},
			wantErr: "not resolved",
		},
		{
			desc: "dut assigned ate device",
			tb:   &opb.Testbed{Duts: []*opb.Device{tDUT1}},
			solveResult: &SolveResult{
				Assignment: &portgraph.Assignment{
					Node2Node: map[*portgraph.AbstractNode]*portgraph.ConcreteNode{
						absDUT1Node: conATE1Node,
					},
				},
				AbsNode2Dev: map[*portgraph.AbstractNode]*opb.Device{
					absDUT1Node: tDUT1,
				},
				ConNode2BindDev: map[*portgraph.ConcreteNode]*binding.Device{
					conATE1Node: &bdevATE1,
				},
				AbsPort2Port:     map[*portgraph.AbstractPort]*opb.Port{},
				ConPort2BindPort: map[*portgraph.ConcretePort]*binding.Port{},
			},
			wantErr: "is not a DUT",
		},
		{
			desc: "ate assigned dut device",
			tb:   &opb.Testbed{Ates: []*opb.Device{tATE1}},
			solveResult: &SolveResult{
				Assignment: &portgraph.Assignment{
					Node2Node: map[*portgraph.AbstractNode]*portgraph.ConcreteNode{
						absATE1Node: conDUT1Node,
					},
				},
				AbsNode2Dev: map[*portgraph.AbstractNode]*opb.Device{
					absATE1Node: tATE1,
				},
				ConNode2BindDev: map[*portgraph.ConcreteNode]*binding.Device{
					conDUT1Node: &bdevDUT1,
				},
				AbsPort2Port:     map[*portgraph.AbstractPort]*opb.Port{},
				ConPort2BindPort: map[*portgraph.ConcretePort]*binding.Port{},
			},
			wantErr: "is not an ATE",
		},
		{
			desc: "nil assignment",
			tb:   &opb.Testbed{Duts: []*opb.Device{tDUT1}},
			solveResult: &SolveResult{
				Assignment:       nil,
				AbsNode2Dev:      map[*portgraph.AbstractNode]*opb.Device{},
				AbsPort2Port:     map[*portgraph.AbstractPort]*opb.Port{},
				ConNode2BindDev:  map[*portgraph.ConcreteNode]*binding.Device{},
				ConPort2BindPort: map[*portgraph.ConcretePort]*binding.Port{},
			},
			wantErr: "solveResult.Assignment is nil",
		},
		{
			desc:        "nil solveResult",
			tb:          &opb.Testbed{Duts: []*opb.Device{tDUT1}},
			solveResult: nil,
			wantErr:     "solveResult is nil",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			_, err := AssignmentToReservation(test.solveResult, test.tb)
			if err == nil || !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("AssignmentToReservation() got err %v, want err containing %q", err, test.wantErr)
			}
		})
	}
}
