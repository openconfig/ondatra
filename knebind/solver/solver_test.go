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

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gnmi/errdiff"
	tpb "github.com/openconfig/kne/proto/topo"
	"github.com/openconfig/ondatra/binding"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"

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
		    value: {}
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
		Id:     "dut1",
		Vendor: opb.Device_ARISTA,
		Ports:  []*opb.Port{{Id: "port1"}, {Id: "port2"}},
	}
	dut2 := &opb.Device{
		Id:    "dut2",
		Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}},
	}
	dut3 := &opb.Device{
		Id:     "dut3",
		Vendor: opb.Device_JUNIPER,
		Ports:  []*opb.Port{{Id: "port1"}},
	}
	ate := &opb.Device{
		Id:    "ate",
		Ports: []*opb.Port{{Id: "port1"}},
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
		B: "ate:port1",
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
		}},
		Services: map[string]*tpb.Service{
			"gnmi": &tpb.Service{Name: "gnmi"},
		},
	}
	wantDUT2 := &ServiceDUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{
			Name:   "node2",
			Vendor: opb.Device_CISCO,
			Ports: map[string]*binding.Port{
				"port1": {Name: "eth1"},
				"port2": {Name: "eth2"},
			},
		}},
		Services: map[string]*tpb.Service{
			"gnmi": &tpb.Service{Name: "gnmi"},
		},
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
		}},
		Services: map[string]*tpb.Service{
			"gnmi": &tpb.Service{Name: "gnmi"},
		},
	}
	wantATE := &ServiceATE{
		AbstractATE: &binding.AbstractATE{&binding.Dims{
			Name:   "node4",
			Vendor: opb.Device_IXIA,
			Ports: map[string]*binding.Port{
				"port1": {Name: "eth1"},
			},
		}},
		Services: make(map[string]*tpb.Service),
	}

	tests := []struct {
		desc    string
		tb      *opb.Testbed
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
		desc: "one ate",
		tb: &opb.Testbed{
			Ates: []*opb.Device{ate},
		},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{},
			ATEs: map[string]binding.ATE{
				"ate": wantATE,
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
			Ates:  []*opb.Device{ate},
			Links: []*opb.Link{link14},
		},
		wantRes: &binding.Reservation{
			DUTs: map[string]binding.DUT{
				"dut1": wantDUT1,
			},
			ATEs: map[string]binding.ATE{
				"ate": wantATE,
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
			gotRes, err := Solve(test.tb, topo)
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

func TestSolvePortGroups(t *testing.T) {
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

	res, err := Solve(tb, topo)
	if err != nil {
		t.Fatalf("Solve() got unexpected error: %v", err)
	}
	if res == nil {
		t.Fatalf("Solve() found no solution")
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
			  vendor: HOST
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
		wantErr: "no node in KNE topology to match testbed",
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
		wantErr: "no node in KNE topology to match testbed",
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
		wantErr: "no KNE topology",
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
		wantErr: "no KNE topology matches the testbed",
	}, {
		desc: "port group specified on both ends of link in topology",
		tb: &opb.Testbed{
			Ates: []*opb.Device{{
				Id:    "ate1",
				Ports: []*opb.Port{{Id: "port1"}},
			}},
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
				Ports:  []*opb.Port{{Id: "port1"}},
			}},
			Links: []*opb.Link{
				{A: "dut1:port1", B: "ate1:port1"},
			},
		},
		topo: `
			nodes: {
			  name: "node1"
			  vendor: KEYSIGHT
			  interfaces: {
			    key: "eth1"
			    value: {group: "lag"}
			  }
			}
			nodes: {
			  name: "node2"
			  vendor: ARISTA
			  interfaces: {
			    key: "eth1"
			    value: {group: "lag"}
			  }
			}
			links: {
			  a_node: "node1"
			  a_int: "eth1"
			  z_node: "node2"
			  z_int: "eth1"
			}`,
		wantErr: "unsupported configuration",
	}, {
		desc: "bad port group config in topology",
		tb: &opb.Testbed{
			Ates: []*opb.Device{{
				Id:    "ate1",
				Ports: []*opb.Port{{Id: "port1"}, {Id: "port2"}},
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
			  interfaces: {
			    key: "eth1"
			    value: {group: "lag"}
			  }
			  interfaces: {
			    key: "eth2"
			    value: {group: "lag"}
			  }
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
		wantErr: "inconsistent port group",
	}, {
		desc: "port group specified on both ends of link in testbed",
		tb: &opb.Testbed{
			Ates: []*opb.Device{{
				Id:    "ate1",
				Ports: []*opb.Port{{Id: "port1", Group: "G1"}},
			}},
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
				Ports:  []*opb.Port{{Id: "port1", Group: "G1"}},
			}},
			Links: []*opb.Link{
				{A: "dut1:port1", B: "ate1:port1"},
			},
		},
		topo: `
			nodes: {
			  name: "node1"
			  vendor: KEYSIGHT
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
			}`,
		wantErr: "unsupported configuration",
	}, {
		desc: "bad port group config in testbed",
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
		wantErr: "inconsistent port group",
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
		wantErr: "no KNE topology matches the testbed",
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
		wantErr: "no KNE topology matches the testbed",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			topo := unmarshalTopo(t, test.topo)
			_, gotErr := Solve(test.tb, topo)
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
