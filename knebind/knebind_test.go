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

package knebind

import (
	"golang.org/x/net/context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/crypto/ssh"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/knebind/solver"

	tpb "github.com/google/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

func TestReserve(t *testing.T) {
	const topo = `
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

	var gotResets int
	kneCmdFn = func(cfg *Config, args ...string) ([]byte, error) {
		switch args[1] {
		case "reset":
			gotResets++
			return nil, nil
		case "service":
			return []byte(topo), nil
		default:
			return nil, fmt.Errorf("unexpected args: %s", args)
		}
	}

	bind := &Bind{cfg: &Config{}}
	wantDUT1 := &kneDUT{
		ServiceDUT: &solver.ServiceDUT{
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
		},
		cfg: bind.cfg,
	}
	wantDUT2 := &kneDUT{
		ServiceDUT: &solver.ServiceDUT{
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
		},
		cfg: bind.cfg,
	}
	wantDUT3 := &kneDUT{
		ServiceDUT: &solver.ServiceDUT{
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
		},
		cfg: bind.cfg,
	}
	wantATE := &kneATE{
		ServiceATE: &solver.ServiceATE{
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
		},
		cfg: bind.cfg,
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
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotResets = 0
			gotRes, err := bind.Reserve(context.Background(), tt.tb, time.Minute, time.Minute, nil)
			if err != nil {
				t.Fatalf("Reserve() got error: %v", err)
			}
			if gotRes.ID == "" {
				t.Errorf("Reserve() got reservation missing ID: %v", gotRes)
			}
			gotRes.ID = ""
			if diff := cmp.Diff(tt.wantRes, gotRes, protocmp.Transform(), cmp.AllowUnexported(kneDUT{}, kneATE{})); diff != "" {
				t.Errorf("Reserve() got unexpected diff in reservation (-want,+got): %s", diff)
			}
			if wantResets := len(tt.wantRes.DUTs); wantResets != gotResets {
				t.Errorf("Reserve() got unexpected DUT resets: want %v, got %v", wantResets, gotResets)
			}
		})
	}
}

func TestReserveErrors(t *testing.T) {
	bind := &Bind{cfg: &Config{}}

	tests := []struct {
		desc        string
		tb          *opb.Testbed
		topo        string
		wantErr     string
		wantGNMIErr string
	}{{
		desc:    "too few nodes",
		tb:      &opb.Testbed{Duts: []*opb.Device{{Id: "dut1"}}},
		wantErr: "not enough nodes",
	}, {
		desc:    "too few links",
		tb:      &opb.Testbed{Links: []*opb.Link{{A: "dut1:port1", B: "dut2:port1"}}},
		wantErr: "not enough links",
	}, {
		desc: "missing gnmi",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{Id: "dut1"}},
		},
		topo: `
		  nodes: {
		    name: "node1"
        type: ARISTA_CEOS
		  }
		`,
		wantGNMIErr: "gnmi",
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
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			kneCmdFn = func(cfg *Config, args ...string) ([]byte, error) {
				return []byte(tt.topo), nil
			}
			res, gotErr := bind.Reserve(context.Background(), tt.tb, time.Minute, time.Minute, nil)
			if diff := errdiff.Substring(gotErr, tt.wantErr); diff != "" {
				t.Fatalf("Reserve() got unexpected error diff: %s", diff)
			}
			if tt.wantErr != "" {
				return
			}
			d, ok := res.DUTs["dut1"]
			if !ok {
				t.Fatalf("Node %q not found in topology", "node1")
			}
			_, gnmiErr := d.DialGNMI(context.Background())
			if diff := errdiff.Substring(gnmiErr, tt.wantGNMIErr); diff != "" {
				t.Errorf("DialGNMI() got unexpected error diff: %s", diff)
			}
		})
	}
}

func TestServices(t *testing.T) {
	bind := &Bind{cfg: &Config{}}

	tests := []struct {
		desc         string
		tb           *opb.Testbed
		topo         string
		serviceCheck func(t *testing.T, b binding.Binding, d binding.DUT)
	}{{
		desc: "missing gnmi",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{Id: "dut1"}},
		},
		topo: `
		  nodes: {
		    name: "node1"
        type: ARISTA_CEOS
		  }
		`,
		serviceCheck: func(t *testing.T, b binding.Binding, d binding.DUT) {
			t.Helper()
			if _, err := d.DialGNMI(context.Background()); err == nil {
				t.Fatalf("DialGNMI() got unexpected error: %v", err)
			}
		},
	}, {
		desc: "missing gribi",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{Id: "dut1"}},
		},
		topo: `
		  nodes: {
		    name: "node1"
        type: ARISTA_CEOS
		  }
		`,
		serviceCheck: func(t *testing.T, b binding.Binding, d binding.DUT) {
			t.Helper()
			if _, err := d.DialGRIBI(context.Background()); err == nil {
				t.Fatalf("DialGRIBI() got unexpected error: %v", err)
			}
		},
	}, {
		desc: "missing p4rt",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id: "dut1",
			}},
		},
		topo: `
		  nodes: {
		    name: "node1"
        type: CISCO_CXR
				services {
				  key: 9339
					value {
					  name: "gnmi"
						outside: 9339
						outside_ip: "1.1.1.1"
					}
				}
		  }
		`,
		serviceCheck: func(t *testing.T, b binding.Binding, d binding.DUT) {
			t.Helper()
			if _, err := d.DialP4RT(context.Background()); err == nil {
				t.Fatalf("DialP4RT() got unexpected error: %v", err)
			}
		},
	}, {
		desc: "valid",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id: "dut1",
			}},
		},
		topo: `
		  nodes: {
		    name: "node1"
        type: CISCO_CXR
				services {
				  key: 9336
					value {
					  name: "p4rt"
						outside: 9336
						outside_ip: "1.1.1.1"
					}
				}
				services {
				  key: 9339
					value {
					  name: "gnmi"
						outside: 9339
						outside_ip: "1.1.1.1"
					}
				}
				services {
				  key: 4242
					value {
					  name: "gribi"
						outside: 4242
						outside_ip: "1.1.1.1"
					}
				}
		  }
		`,
		serviceCheck: func(t *testing.T, b binding.Binding, d binding.DUT) {
			t.Helper()
			if _, err := d.DialGNMI(context.Background()); err != nil {
				t.Fatalf("DialGNMI() got unexpected error: %v", err)
			}
			if _, err := d.DialGRIBI(context.Background()); err != nil {
				t.Fatalf("DialGRIBI() got unexpected error: %v", err)
			}
			if _, err := d.DialP4RT(context.Background()); err != nil {
				t.Fatalf("DialP4RT() got unexpected error: %v", err)
			}
		},
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			kneCmdFn = func(cfg *Config, args ...string) ([]byte, error) {
				return []byte(tt.topo), nil
			}
			res, err := bind.Reserve(context.Background(), tt.tb, time.Minute, time.Minute, nil)
			if err != nil {
				t.Fatalf("Reserve() failed: %v", err)
			}
			d, ok := res.DUTs["dut1"]
			if !ok {
				t.Fatalf("Node %q not found in topology", "node1")
			}
			tt.serviceCheck(t, bind, d)
		})
	}
}

func TestPushConfig(t *testing.T) {
	const dutName = "dut"
	var sshErr error
	sshExecFn = func(addr string, cfg *ssh.ClientConfig, cmd string) (_ string, rerr error) {
		return "", sshErr
	}
	var gotReset bool
	kneCmdFn = func(cfg *Config, args ...string) ([]byte, error) {
		for _, a := range args {
			if a == "reset" {
				gotReset = true
			}
		}
		return nil, nil
	}

	tests := []struct {
		desc      string
		vendor    opb.Device_Vendor
		reset     bool
		sshErr    error
		noSSH     bool
		wantReset bool
		wantErr   string
	}{{
		desc:   "success",
		vendor: opb.Device_ARISTA,
	}, {
		desc:      "reset success",
		vendor:    opb.Device_ARISTA,
		reset:     true,
		wantReset: true,
	}, {
		desc:    "only arista support",
		vendor:  opb.Device_CISCO,
		wantErr: "supports Arista",
	}, {
		desc:    "ssh error",
		vendor:  opb.Device_ARISTA,
		sshErr:  errors.New("ssh error"),
		wantErr: "ssh error",
	}, {
		desc:    "no ssh",
		vendor:  opb.Device_ARISTA,
		noSSH:   true,
		wantErr: "\"ssh\" not found",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sshErr = tt.sshErr
			gotReset = false
			dut := &kneDUT{
				ServiceDUT: &solver.ServiceDUT{
					AbstractDUT: &binding.AbstractDUT{&binding.Dims{Name: dutName, Vendor: tt.vendor}},
				},
				cfg: &Config{},
			}
			if !tt.noSSH {
				dut.Services = map[string]*tpb.Service{
					"ssh": &tpb.Service{OutsideIp: "1.2.3.4", Outside: 1234},
				}
			}
			err := dut.PushConfig(context.Background(), "my config", tt.reset)
			if (err == nil) != (tt.wantErr == "") || (err != nil && !strings.Contains(err.Error(), tt.wantErr)) {
				t.Errorf("PushConfig got error %v, want %v", err, tt.wantErr)
			}
			if gotReset != tt.wantReset {
				t.Errorf("PushConfig got reset %v, want %v", gotReset, tt.wantReset)
			}
		})
	}
}
