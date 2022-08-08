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
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/knebind/solver"

	tpb "github.com/openconfig/kne/proto/topo"
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
    }
		nodes: {
		  name: "node2"
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
		}`

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
	tb := &opb.Testbed{
		Duts: []*opb.Device{{
			Id:    "dut",
			Ports: []*opb.Port{{Id: "port1"}},
		}},
		Ates: []*opb.Device{{
			Id:    "ate",
			Ports: []*opb.Port{{Id: "port1"}},
		}},
		Links: []*opb.Link{{
			A: "dut:port1",
			B: "ate:port1",
		}},
	}

	wantRes := &binding.Reservation{
		DUTs: map[string]binding.DUT{
			"dut": &kneDUT{
				ServiceDUT: &solver.ServiceDUT{
					AbstractDUT: &binding.AbstractDUT{&binding.Dims{
						Name:            "node1",
						Vendor:          opb.Device_ARISTA,
						HardwareModel:   "ARISTA_CEOS",
						SoftwareVersion: "ARISTA_CEOS",
						Ports: map[string]*binding.Port{
							"port1": {Name: "Ethernet1"},
						},
					}},
					Services: map[string]*tpb.Service{
						"gnmi": &tpb.Service{Name: "gnmi"},
					},
				},
				cfg: bind.cfg,
			},
		},
		ATEs: map[string]binding.ATE{
			"ate": &kneATE{
				ServiceATE: &solver.ServiceATE{
					AbstractATE: &binding.AbstractATE{&binding.Dims{
						Name:            "node2",
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
			},
		},
	}

	gotResets = 0
	gotRes, err := bind.Reserve(context.Background(), tb, time.Minute, time.Minute, nil)
	if err != nil {
		t.Fatalf("Reserve() got error: %v", err)
	}
	if gotRes.ID == "" {
		t.Errorf("Reserve() got reservation missing ID: %v", gotRes)
	}
	gotRes.ID = ""
	if diff := cmp.Diff(wantRes, gotRes, protocmp.Transform(), cmp.AllowUnexported(kneDUT{}, kneATE{})); diff != "" {
		t.Errorf("Reserve() got unexpected diff in reservation (-want,+got): %s", diff)
	}
	if wantResets := len(wantRes.DUTs); wantResets != gotResets {
		t.Errorf("Reserve() got unexpected DUT resets: want %v, got %v", wantResets, gotResets)
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
