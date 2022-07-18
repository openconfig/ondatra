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
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/testing/protocmp"
	tpb "github.com/openconfig/kne/proto/topo"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/binding"

	opb "github.com/openconfig/ondatra/proto"
)

func TestSolveDutError(t *testing.T) {
	const topoText = `
		nodes: {
			name: "node1"
			type: ARISTA_CEOS
			services: {
				key: 1
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
			name: "node2"
			type: ARISTA_CEOS
			services: {
				key: 2
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
			type: CISCO_CXR
			services: {
				key: 3
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
			type: CISCO_CXR
			services: {
				key: 4
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
			name: "node5"
			type: JUNIPER_CEVO
			services: {
				key: 5
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
			name: "node6"
			type: JUNIPER_CEVO
			services: {
				key: 6
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
			name: "node7"
			type: IXIA_TG
			interfaces: {
				key: "eth1"
				value: {}
			}
			interfaces: {
				key: "eth2"
				value: {}
			}
			interfaces: {
				key: "eth3"
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
			z_node: "node7"
			z_int: "eth1"
		}
		links: {
			a_node: "node3"
			a_int: "eth1"
			z_node: "node4"
			z_int: "eth1"
		}
		links: {
			a_node: "node4"
			a_int: "eth2"
			z_node: "node7"
			z_int: "eth2"
		}
		links: {
			a_node: "node5"
			a_int: "eth1"
			z_node: "node6"
			z_int: "eth1"
		}
		links: {
			a_node: "node6"
			a_int: "eth2"
			z_node: "node7"
			z_int: "eth3"
		}`
	topo := unmarshalTopo(t, topoText)

	dut1 := &opb.Device{
		Id:     "dut1",
		Vendor: opb.Device_ARISTA,
		Ports:  []*opb.Port{{Id: "port1"}},
	}
	dut2 := &opb.Device{
		Id:    "dut2",
		Vendor: opb.Device_CISCO,
		Ports: []*opb.Port{{Id: "port1"}},
	}
	dut3 := &opb.Device{
		Id:    "dut3",
		Vendor: opb.Device_JUNIPER,
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

	tb := &opb.Testbed{
			Duts:  []*opb.Device{dut1, dut2, dut3},
			Links: []*opb.Link{link12, link23},
		}

	_, err := Solve(tb, topo)
	if err == nil {
		t.Fatalf("Solve() got unexpected error: %v", err)
	}
}

func TestSolve(t *testing.T) {
	const topoText = `
		nodes: {
		  name: "node1"
      type: ARISTA_CEOS
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
      type: CISCO_CXR
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
      type: JUNIPER_CEVO
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
      type: IXIA_TG
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
			HardwareModel:   "ARISTA_CEOS",
			SoftwareVersion: "ARISTA_CEOS",
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
			Name:            "node2",
			Vendor:          opb.Device_CISCO,
			HardwareModel:   "CISCO_CXR",
			SoftwareVersion: "CISCO_CXR",
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
			HardwareModel:   "JUNIPER_CEVO",
			SoftwareVersion: "JUNIPER_CEVO",
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
			Name:            "node4",
			Vendor:          opb.Device_IXIA,
			HardwareModel:   "IXIA_TG",
			SoftwareVersion: "IXIA_TG",
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
		desc:    "too few links",
		tb:      &opb.Testbed{Links: []*opb.Link{{A: "dut1:port1", B: "dut2:port1"}}},
		wantErr: "not enough links",
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
        type: CISCO_CXR
		  }
		`,
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
        type: CISCO_CXR
		  }
		`,
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
        type: ARISTA_CEOS
		  }
		  nodes: {
		    name: "node2"
        type: CISCO_CXR
		  }
		`,
		wantErr: "no combination of nodes",
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
        type: ARISTA_CEOS
		  }
		  nodes: {
		    name: "node2"
        type: ARISTA_CEOS
		  }
		  nodes: {
		    name: "node3"
        type: JUNIPER_VMX
		  }
		  nodes: {
		    name: "node4"
        type: JUNIPER_VMX
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
		  }
		`,
		wantErr: "no KNE topology",
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

func unmarshalTopo(t *testing.T, text string) *tpb.Topology {
	topo := new(tpb.Topology)
	if err := prototext.Unmarshal([]byte(text), topo); err != nil {
		t.Fatalf("Error unmarshalling topology text: %v", err)
	}
	return topo
}
