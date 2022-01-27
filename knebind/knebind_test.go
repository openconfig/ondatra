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
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
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
	kneCmdFn = func(cfg *Config, args ...string) ([]byte, error) {
		return []byte(topo), nil
	}
	wantDUT1 := &reservation.DUT{&reservation.Dims{
		Name:            "node1",
		Vendor:          opb.Device_ARISTA,
		HardwareModel:   "ARISTA_CEOS",
		SoftwareVersion: "ARISTA_CEOS",
		Ports: map[string]*reservation.Port{
			"port1": {Name: "Ethernet1"},
			"port2": {Name: "Ethernet2"},
		},
	}}
	wantDUT2 := &reservation.DUT{&reservation.Dims{
		Name:            "node2",
		Vendor:          opb.Device_CISCO,
		HardwareModel:   "CISCO_CXR",
		SoftwareVersion: "CISCO_CXR",
		Ports: map[string]*reservation.Port{
			"port1": {Name: "eth1"},
			"port2": {Name: "eth2"},
		},
	}}
	wantDUT3 := &reservation.DUT{&reservation.Dims{
		Name:            "node3",
		Vendor:          opb.Device_JUNIPER,
		HardwareModel:   "JUNIPER_CEVO",
		SoftwareVersion: "JUNIPER_CEVO",
		Ports: map[string]*reservation.Port{
			"port1": {Name: "eth1"},
		},
	}}
	wantATE := &reservation.ATE{&reservation.Dims{
		Name:            "node4",
		Vendor:          opb.Device_IXIA,
		HardwareModel:   "IXIA_TG",
		SoftwareVersion: "IXIA_TG",
		Ports: map[string]*reservation.Port{
			"port1": {Name: "eth1"},
		},
	}}
	bind := &Bind{cfg: &Config{}}

	tests := []struct {
		desc    string
		tb      *opb.Testbed
		wantRes *reservation.Reservation
	}{{
		desc: "one dut",
		tb: &opb.Testbed{
			Duts: []*opb.Device{dut3},
		},
		wantRes: &reservation.Reservation{
			DUTs: map[string]*reservation.DUT{
				"dut3": wantDUT3,
			},
			ATEs: map[string]*reservation.ATE{},
		},
	}, {
		desc: "one ate",
		tb: &opb.Testbed{
			Ates: []*opb.Device{ate},
		},
		wantRes: &reservation.Reservation{
			DUTs: map[string]*reservation.DUT{},
			ATEs: map[string]*reservation.ATE{
				"ate": wantATE,
			},
		},
	}, {
		desc: "two duts",
		tb: &opb.Testbed{
			Duts:  []*opb.Device{dut1, dut2},
			Links: []*opb.Link{link12},
		},
		wantRes: &reservation.Reservation{
			DUTs: map[string]*reservation.DUT{
				"dut1": wantDUT1,
				"dut2": wantDUT2,
			},
			ATEs: map[string]*reservation.ATE{},
		},
	}, {
		desc: "dut and ate",
		tb: &opb.Testbed{
			Duts:  []*opb.Device{dut1},
			Ates:  []*opb.Device{ate},
			Links: []*opb.Link{link14},
		},
		wantRes: &reservation.Reservation{
			DUTs: map[string]*reservation.DUT{
				"dut1": wantDUT1,
			},
			ATEs: map[string]*reservation.ATE{
				"ate": wantATE,
			},
		},
	}, {
		desc: "three duts",
		tb: &opb.Testbed{
			Duts:  []*opb.Device{dut1, dut2, dut3},
			Links: []*opb.Link{link12, link23},
		},
		wantRes: &reservation.Reservation{
			DUTs: map[string]*reservation.DUT{
				"dut1": wantDUT1,
				"dut2": wantDUT2,
				"dut3": wantDUT3,
			},
			ATEs: map[string]*reservation.ATE{},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotRes, err := bind.Reserve(context.Background(), tt.tb, time.Minute, time.Minute)
			if err != nil {
				t.Fatalf("Reserve() got error: %v", err)
			}
			if gotRes.ID == "" {
				t.Errorf("Reserve() got reservation missing ID: %v", gotRes)
			}
			gotRes.ID = ""
			if diff := cmp.Diff(tt.wantRes, gotRes); diff != "" {
				t.Errorf("Reserve() got unexpected diff in reservation (-want,+got): %s", diff)
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
		wantErr: "Not enough nodes",
	}, {
		desc:    "too few links",
		tb:      &opb.Testbed{Links: []*opb.Link{{A: "dut1:port1", B: "dut2:port1"}}},
		wantErr: "Not enough links",
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
		wantErr: "No node in KNE topology to match testbed",
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
		wantErr: "No node in KNE topology to match testbed",
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
		wantErr: "No combination of nodes",
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
		wantErr: "No KNE topology",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			kneCmdFn = func(cfg *Config, args ...string) ([]byte, error) {
				return []byte(tt.topo), nil
			}
			res, gotErr := bind.Reserve(context.Background(), tt.tb, time.Minute, time.Minute)
			if diff := errdiff.Substring(gotErr, tt.wantErr); diff != "" {
				t.Fatalf("Reserve() got unexpected error diff: %s", diff)
			}
			if tt.wantErr != "" {
				return
			}
			d, err := res.DUT("dut1")
			if err != nil {
				t.Fatalf("Node %q not found in topology", "node1")
			}
			_, gnmiErr := bind.DialGNMI(context.Background(), d)
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
		serviceCheck func(t *testing.T, b binding.Binding, d *reservation.DUT)
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
		serviceCheck: func(t *testing.T, b binding.Binding, d *reservation.DUT) {
			t.Helper()
			if _, err := b.DialGNMI(context.Background(), d); err == nil {
				t.Fatalf("DialGNMI() got unexpected error: %v", err)
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
		serviceCheck: func(t *testing.T, b binding.Binding, d *reservation.DUT) {
			t.Helper()
			if _, err := b.DialP4RT(context.Background(), d); err == nil {
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
		  }
		`,
		serviceCheck: func(t *testing.T, b binding.Binding, d *reservation.DUT) {
			t.Helper()
			if _, err := b.DialGNMI(context.Background(), d); err != nil {
				t.Fatalf("DialGNMI() got unexpected error: %v", err)
			}
			if _, err := b.DialP4RT(context.Background(), d); err != nil {
				t.Fatalf("DialP4RT() got unexpected error: %v", err)
			}
		},
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			kneCmdFn = func(cfg *Config, args ...string) ([]byte, error) {
				return []byte(tt.topo), nil
			}
			res, err := bind.Reserve(context.Background(), tt.tb, time.Minute, time.Minute)
			if err != nil {
				t.Fatalf("Reserve() failed: %v", err)
			}
			d, err := res.DUT("dut1")
			if err != nil {
				t.Fatalf("Node %q not found in topology", "node1")
			}
			tt.serviceCheck(t, bind, d)
		})
	}
}

func TestPushConfig(t *testing.T) {
	const dutName = "dut"
	bind := &Bind{
		cfg: &Config{},
		services: solver.ServiceMap{dutName: map[string]*tpb.Service{
			"ssh": &tpb.Service{OutsideIp: "1.2.3.4", Outside: 1234},
		}},
	}
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
		dut       *reservation.DUT
		opts      *binding.ConfigOptions
		sshErr    error
		wantReset bool
		wantErr   string
	}{{
		desc:      "replace success",
		dut:       &reservation.DUT{&reservation.Dims{Name: dutName, Vendor: opb.Device_ARISTA}},
		opts:      &binding.ConfigOptions{},
		wantReset: true,
	}, {
		desc: "append success",
		dut:  &reservation.DUT{&reservation.Dims{Name: dutName, Vendor: opb.Device_ARISTA}},
		opts: &binding.ConfigOptions{Append: true},
	}, {
		desc:    "only arista support",
		dut:     &reservation.DUT{&reservation.Dims{Name: dutName, Vendor: opb.Device_CISCO}},
		opts:    &binding.ConfigOptions{Append: true},
		wantErr: "supports Arista",
	}, {
		desc:    "no openconfig support",
		dut:     &reservation.DUT{&reservation.Dims{Name: dutName, Vendor: opb.Device_ARISTA}},
		opts:    &binding.ConfigOptions{OpenConfig: true},
		wantErr: "OpenConfig",
	}, {
		desc:    "ssh error",
		dut:     &reservation.DUT{&reservation.Dims{Name: dutName, Vendor: opb.Device_ARISTA}},
		opts:    &binding.ConfigOptions{Append: true},
		sshErr:  errors.New("ssh error"),
		wantErr: "ssh error",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sshErr = tt.sshErr
			gotReset = false
			err := bind.PushConfig(context.Background(), tt.dut, "my config", tt.opts)
			if (err == nil) != (tt.wantErr == "") || (err != nil && !strings.Contains(err.Error(), tt.wantErr)) {
				t.Errorf("PushConfig got error %v, want %v", err, tt.wantErr)
			}
			if gotReset != tt.wantReset {
				t.Errorf("PushConfig got reset %v, want %v", gotReset, tt.wantReset)
			}
		})
	}
}
