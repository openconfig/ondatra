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

package knebind

import (
	"errors"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/openconfig/kne/topo/node"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/knebind/creds"
	"github.com/openconfig/ondatra/knebind/solver"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding/introspect"
	"google.golang.org/grpc"

	cpb "github.com/openconfig/kne/proto/controller"
	tpb "github.com/openconfig/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

type fakeGosnappi struct {
	gosnappi.Api
	capturedGrpcTransport gosnappi.GrpcTransport
}

func (f *fakeGosnappi) NewGrpcTransport() gosnappi.GrpcTransport {
	f.capturedGrpcTransport = f.Api.NewGrpcTransport()
	return f.capturedGrpcTransport
}

func TestDialOTG(t *testing.T) {
	tests := []struct {
		name        string
		otgTimeout  time.Duration
		servicesMap map[string]*tpb.Service
		wantTimeout time.Duration
		wantErr     bool
	}{
		{
			name:       "Success with Provided Timeout",
			otgTimeout: 111 * time.Second,
			servicesMap: map[string]*tpb.Service{
				"grpc": {Name: "otg"},
			},
			wantTimeout: 111 * time.Second,
		},
		{
			name: "Success with Zero Timeout",
			servicesMap: map[string]*tpb.Service{
				"grpc": {Name: "otg"},
			},
			wantTimeout: 0,
		},
		{
			name:        "OTG Service Not Found",
			otgTimeout:  60 * time.Second,
			servicesMap: map[string]*tpb.Service{},
			wantErr:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			origAteDialGRPCFn := ateDialGRPCFn
			origGosnappiNewAPIFn := gosnappiNewAPIFn
			defer func() {
				ateDialGRPCFn = origAteDialGRPCFn
				gosnappiNewAPIFn = origGosnappiNewAPIFn
			}()

			ateDialGRPCFn = func(a *kneATE, ctx context.Context, svc introspect.Service, opts []grpc.DialOption) (*grpc.ClientConn, error) {
				if svc != introspect.OTG {
					return nil, errors.New("fakeAteDialGRPC: unexpected service, want OTG")
				}
				return &grpc.ClientConn{}, nil
			}

			fakeAPI := &fakeGosnappi{Api: gosnappi.NewApi()}
			gosnappiNewAPIFn = func() gosnappi.Api { return fakeAPI }

			// Prepare kneATE with the desired OTGRequestTimeout in the config
			ate := &kneATE{
				ServiceATE: &solver.ServiceATE{
					Services: tc.servicesMap,
					AbstractATE: &binding.AbstractATE{
						Dims: &binding.Dims{Name: "fakeATE"},
					},
				},
				bind: &Bind{
					otgRequestTimeout: tc.otgTimeout,
				},
			}

			_, err := ate.DialOTG(context.Background())

			if (err != nil) != tc.wantErr {
				t.Fatalf("DialOTG() got error %v, want error? %v", err, tc.wantErr)
			}
			if err != nil {
				return
			}
			if gotTimeout := fakeAPI.capturedGrpcTransport.RequestTimeout(); gotTimeout != tc.wantTimeout {
				t.Errorf("DialOTG() SetRequestTimeout got %v, want %v", gotTimeout, tc.wantTimeout)
			}
		})
	}
}

func TestNew(t *testing.T) {
	userHomeDir := t.TempDir()

	tests := []struct {
		desc              string
		configPath        string
		user              *user.User
		otgRequestTimeout time.Duration
		setOTGTimeout     bool
		wantPath          string
		wantErr           bool
		wantOTGTimeout    time.Duration
	}{{
		desc:           "kubeconfig provided",
		configPath:     userHomeDir + "/config",
		wantPath:       userHomeDir + "/config",
		wantOTGTimeout: DefaultOTGRequestTimeout,
	}, {
		desc:           "default kubeconfig",
		user:           &user.User{HomeDir: userHomeDir},
		wantPath:       userHomeDir + "/.kube/config",
		wantOTGTimeout: DefaultOTGRequestTimeout,
	}, {
		desc:              "zero otgRequestTimeout uses default",
		user:              &user.User{HomeDir: userHomeDir},
		otgRequestTimeout: 0,
		setOTGTimeout:     true,
		wantPath:          userHomeDir + "/.kube/config",
		wantOTGTimeout:    DefaultOTGRequestTimeout,
	}, {
		desc:              "positive otgRequestTimeout",
		user:              &user.User{HomeDir: userHomeDir},
		otgRequestTimeout: 15 * time.Second,
		setOTGTimeout:     true,
		wantPath:          userHomeDir + "/.kube/config",
		wantOTGTimeout:    15 * time.Second,
	}, {
		desc:              "negative otgRequestTimeout",
		user:              &user.User{HomeDir: userHomeDir},
		otgRequestTimeout: -10 * time.Second,
		setOTGTimeout:     true,
		wantPath:          userHomeDir + "/.kube/config",
		wantErr:           true,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			userCurrFn = func() (*user.User, error) {
				return test.user, nil
			}
			writeDummyKubeconfig(t, test.wantPath)
			cfg := &Config{
				Topology:   "integration/topology.textproto",
				Kubeconfig: test.configPath,
			}

			var opts []Option
			if test.setOTGTimeout {
				opts = append(opts, WithOTGRequestTimeout(test.otgRequestTimeout))
			}

			bind, err := New(cfg, opts...)
			if (err != nil) != test.wantErr {
				t.Errorf("New() with OTGRequestTimeout %v got error %v, want error? %v", test.otgRequestTimeout, err, test.wantErr)
			}
			if err != nil {
				return
			}
			if gotPath := cfg.Kubeconfig; gotPath != test.wantPath {
				t.Errorf("New() got unexpected kubeconfig path %q, want %q", gotPath, test.wantPath)
			}
			if bind.otgRequestTimeout != test.wantOTGTimeout {
				t.Errorf("New() got OTGRequestTimeout %v, want %v", bind.otgRequestTimeout, test.wantOTGTimeout)
			}
		})
	}
}

func writeDummyKubeconfig(t *testing.T, path string) {
	const text = `apiVersion: v1
clusters:
  - cluster:
      server: http://kubeconfig.dir.com
    name: foo
contexts:
  - context:
      cluster: foo
    name: foo-context
current-context: foo-context
`
	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		t.Fatalf("MkdirAll() got unexpected error: %v", err)
	}
	if err := os.WriteFile(path, []byte(text), 0644); err != nil {
		t.Fatalf("WriteFile() got unexpected error: %v", err)
	}
}

func TestReserve(t *testing.T) {
	top := &tpb.Topology{
		Nodes: []*tpb.Node{{
			Name:   "node1",
			Model:  "ceos",
			Os:     "eos",
			Vendor: tpb.Vendor_ARISTA,
			Services: map[uint32]*tpb.Service{
				1234: &tpb.Service{Name: "gnmi"},
			},
			Interfaces: map[string]*tpb.Interface{
				"eth1": {Name: "Ethernet1"},
			},
		}, {
			Name:   "node2",
			Vendor: tpb.Vendor_KEYSIGHT,
			Interfaces: map[string]*tpb.Interface{
				"eth1": {},
			},
		}},
		Links: []*tpb.Link{{
			ANode: "node1",
			AInt:  "eth1",
			ZNode: "node2",
			ZInt:  "eth1",
		}},
	}
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

	wantDUTServices := map[string]*tpb.Service{"gnmi": &tpb.Service{Name: "gnmi"}}
	wantATEServices := make(map[string]*tpb.Service)
	wantRes := &binding.Reservation{
		DUTs: map[string]binding.DUT{
			"dut": &kneDUT{
				ServiceDUT: &solver.ServiceDUT{
					AbstractDUT: &binding.AbstractDUT{&binding.Dims{
						Name:            "node1",
						Vendor:          opb.Device_ARISTA,
						HardwareModel:   "ceos",
						SoftwareVersion: "eos",
						Ports: map[string]*binding.Port{
							"port1": {Name: "Ethernet1"},
						},
					}},
					Services:   wantDUTServices,
					NodeVendor: tpb.Vendor_ARISTA,
				},
			},
		},
		ATEs: map[string]binding.ATE{
			"ate": &kneATE{
				ServiceATE: &solver.ServiceATE{
					AbstractATE: &binding.AbstractATE{&binding.Dims{
						Name:   "node2",
						Vendor: opb.Device_IXIA,
						Ports: map[string]*binding.Port{
							"port1": {Name: "eth1"},
						},
					}},
					Services:   wantATEServices,
					NodeVendor: tpb.Vendor_KEYSIGHT,
				},
			},
		},
	}

	tests := []struct {
		desc      string
		skipReset bool
	}{{
		desc: "with reset",
	}, {
		desc:      "skip reset",
		skipReset: true,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tm := &fakeTopoManager{top: top}
			bind := &Bind{cfg: &Config{}, tm: tm}
			wantResets := len(wantRes.DUTs)
			if tt.skipReset {
				bind.cfg.SkipReset = true
				wantResets = 0
			}
			gotRes, err := bind.Reserve(context.Background(), tb, time.Minute, time.Minute, nil)
			if err != nil {
				t.Fatalf("Reserve() got error: %v", err)
			}
			if gotRes.ID == "" {
				t.Errorf("Reserve() got reservation missing ID: %v", gotRes)
			}
			if diff := cmp.Diff(wantRes, gotRes, protocmp.Transform(), cmp.AllowUnexported(kneDUT{}, kneATE{}), cmpopts.IgnoreFields(kneDUT{}, "bind"), cmpopts.IgnoreFields(kneATE{}, "bind"), cmpopts.IgnoreFields(binding.Reservation{}, "ID")); diff != "" {
				t.Errorf("Reserve() got unexpected diff in reservation (-want,+got): %s", diff)
			}
			if wantResets != tm.gotResets {
				t.Errorf("Reserve() got unexpected resets: want %v, got %v", wantResets, tm.gotResets)
			}
		})
	}
}

func TestNewRPCCredentials(t *testing.T) {
	tests := []struct {
		desc string
		cfg  *Config
		want *creds.UserPass
	}{{
		desc: "nil credentials",
		cfg:  &Config{},
		want: nil,
	}, {
		desc: "empty credentials",
		cfg: &Config{
			Credentials: &creds.Credentials{},
		},
		want: nil,
	}, {
		desc: "empty credentials with fallback",
		cfg: &Config{
			Credentials: &creds.Credentials{},
			Username:    "deprecatedUser",
			Password:    "deprecatedPass",
		},
		want: &creds.UserPass{Username: "deprecatedUser", Password: "deprecatedPass"},
	}, {
		desc: "from credentials",
		cfg: &Config{
			Credentials: &creds.Credentials{
				Default: &creds.UserPass{
					Username: "defaultUser",
					Password: "defaultPass",
				},
			},
			Username: "deprecatedUser",
			Password: "deprecatedPass",
		},
		want: &creds.UserPass{Username: "defaultUser", Password: "defaultPass"},
	}, {
		desc: "deprecated credentials",
		cfg: &Config{
			Username: "deprecatedUser",
			Password: "deprecatedPass",
		},
		want: &creds.UserPass{Username: "deprecatedUser", Password: "deprecatedPass"},
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			dut := &kneDUT{
				ServiceDUT: &solver.ServiceDUT{
					AbstractDUT: &binding.AbstractDUT{&binding.Dims{}},
				},
				bind: &Bind{cfg: tt.cfg},
			}
			var got *creds.UserPass
			if gotRPCCreds := dut.newRPCCredentials(); gotRPCCreds != nil {
				got = gotRPCCreds.UserPass
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("newRPCCredentials() got unexpected diff (-want,+got): %s", diff)
			}
		})
	}
}

func TestServices(t *testing.T) {
	tests := []struct {
		desc         string
		tb           *opb.Testbed
		topo         *tpb.Topology
		serviceCheck func(t *testing.T, b binding.Binding, d binding.DUT)
	}{{
		desc: "missing gnmi",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{Id: "dut1"}},
		},
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name:   "node1",
				Vendor: tpb.Vendor_ARISTA,
			}},
		},
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
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name:   "node1",
				Vendor: tpb.Vendor_ARISTA,
			}},
		},
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
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name:   "node1",
				Vendor: tpb.Vendor_CISCO,
				Services: map[uint32]*tpb.Service{
					9339: &tpb.Service{
						Name:      "gnmi",
						Outside:   9339,
						OutsideIp: "1.1.1.1",
					},
				},
			}},
		},
		serviceCheck: func(t *testing.T, b binding.Binding, d binding.DUT) {
			t.Helper()
			if _, err := d.DialP4RT(context.Background()); err == nil {
				t.Fatalf("DialP4RT() got unexpected error: %v", err)
			}
		},
	}, {
		desc: "missing gnsi",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id: "dut1",
			}},
		},
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name:   "node1",
				Vendor: tpb.Vendor_CISCO,
				Services: map[uint32]*tpb.Service{
					9339: &tpb.Service{
						Name:      "gnmi",
						Outside:   9339,
						OutsideIp: "1.1.1.1",
					},
				},
			}},
		},
		serviceCheck: func(t *testing.T, b binding.Binding, d binding.DUT) {
			t.Helper()
			if _, err := d.DialGNSI(context.Background()); err == nil {
				t.Fatalf("DialGNSI() got unexpected error: %v", err)
			}
		},
	}, {
		desc: "valid",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id: "dut1",
			}},
		},
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name:   "node1",
				Vendor: tpb.Vendor_CISCO,
				Services: map[uint32]*tpb.Service{
					9336: &tpb.Service{
						Name:      "p4rt",
						Outside:   9336,
						OutsideIp: "1.1.1.1",
					},
					9339: &tpb.Service{
						Name:      "gnmi",
						Outside:   9339,
						OutsideIp: "1.1.1.1",
					},
					4242: &tpb.Service{
						Name:      "gribi",
						Outside:   4242,
						OutsideIp: "1.1.1.1",
					},
					9340: &tpb.Service{
						Name:      "gnsi",
						Outside:   9339,
						OutsideIp: "1.1.1.1",
					},
				},
			}},
		},
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
			if _, err := d.DialGNSI(context.Background()); err != nil {
				t.Fatalf("DialGNSI() got unexpected error: %v", err)
			}
		},
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			bind := &Bind{
				cfg: &Config{},
				tm:  &fakeTopoManager{top: tt.topo},
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
	top := &tpb.Topology{
		Nodes: []*tpb.Node{{
			Name:   dutName,
			Config: &tpb.Config{ConfigData: &tpb.Config_Data{[]byte("init")}},
		}},
	}

	tests := []struct {
		desc      string
		reset     bool
		resetErr  error
		pushErr   error
		wantReset bool
		wantErr   string
	}{{
		desc: "success",
	}, {
		desc:      "reset success",
		reset:     true,
		wantReset: true,
	}, {
		desc:     "reset error",
		reset:    true,
		resetErr: errors.New("reset error"),
		wantErr:  "reset error",
	}, {
		desc:    "push error",
		pushErr: errors.New("push error"),
		wantErr: "push error",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			tm := &fakeTopoManager{top: top, resetErr: tt.resetErr, pushErr: tt.pushErr}
			dut := &kneDUT{
				ServiceDUT: &solver.ServiceDUT{
					AbstractDUT: &binding.AbstractDUT{&binding.Dims{Name: dutName}},
				},
				bind: &Bind{tm: tm, cfg: &Config{}},
			}
			err := dut.PushConfig(context.Background(), "my config", tt.reset)
			if (err == nil) != (tt.wantErr == "") || (err != nil && !strings.Contains(err.Error(), tt.wantErr)) {
				t.Errorf("PushConfig got error %v, want %v", err, tt.wantErr)
			}
			if gotReset := tm.gotResets > 0; gotReset != tt.wantReset {
				t.Errorf("PushConfig got reset %v, want %v", gotReset, tt.wantReset)
			}
		})
	}
}

func TestKneDUT(t *testing.T) {
	t.Run("RPCUsername and RPCPassword", func(t *testing.T) {
		dut := &kneDUT{
			ServiceDUT: &solver.ServiceDUT{
				AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1", Vendor: opb.Device_ARISTA}},
				NodeVendor:  tpb.Vendor_ARISTA,
			},
			bind: &Bind{
				cfg: &Config{},
			},
		}
		tests := []struct {
			desc     string
			cfg      *Config
			wantUser string
			wantPass string
		}{{
			desc: "no creds",
			cfg:  &Config{},
		}, {
			desc: "default user pass",
			cfg: &Config{
				Username: "user",
				Password: "password",
			},
			wantUser: "user",
			wantPass: "password",
		}, {
			desc: "creds by node name",
			cfg: &Config{
				Username: "user",
				Password: "password",
				Credentials: &creds.Credentials{
					Node: map[string]*creds.UserPass{
						"dut1": {Username: "node_user", Password: "node_password"},
					},
				},
			},
			wantUser: "node_user",
			wantPass: "node_password",
		}, {
			desc: "creds by vendor",
			cfg: &Config{
				Username: "user",
				Password: "password",
				Credentials: &creds.Credentials{
					Vendor: map[tpb.Vendor]*creds.UserPass{
						tpb.Vendor_ARISTA: {Username: "vendor_user", Password: "vendor_password"},
					},
				},
			},
			wantUser: "vendor_user",
			wantPass: "vendor_password",
		}, {
			desc: "creds by node name precedence",
			cfg: &Config{
				Username: "user",
				Password: "password",
				Credentials: &creds.Credentials{
					Node: map[string]*creds.UserPass{
						"dut1": {Username: "node_user", Password: "node_password"},
					},
					Vendor: map[tpb.Vendor]*creds.UserPass{
						tpb.Vendor_ARISTA: {Username: "vendor_user", Password: "vendor_password"},
					},
				},
			},
			wantUser: "node_user",
			wantPass: "node_password",
		}}

		for _, tt := range tests {
			t.Run(tt.desc, func(t *testing.T) {
				dut.bind.cfg = tt.cfg
				if got := dut.RPCUsername(); got != tt.wantUser {
					t.Errorf("RPCUsername() got %q, want %q", got, tt.wantUser)
				}
				if got := dut.RPCPassword(); got != tt.wantPass {
					t.Errorf("RPCPassword() got %q, want %q", got, tt.wantPass)
				}
			})
		}
	})

	t.Run("DialGRPCWithPort success", func(t *testing.T) {
		d := &kneDUT{
			ServiceDUT: &solver.ServiceDUT{
				AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1"}},
				Services: map[string]*tpb.Service{
					"gnmi": {
						Name:      "gnmi",
						Inside:    9339,
						Outside:   9339,
						OutsideIp: "1.1.1.1",
					},
				},
			},
			bind: &Bind{cfg: &Config{}},
		}
		_, err := d.DialGRPCWithPort(context.Background(), 9339)
		if err != nil {
			t.Fatalf("DialGRPCWithPort() got unexpected error: %v", err)
		}
	})

	t.Run("DialGRPCWithPort not found", func(t *testing.T) {
		d := &kneDUT{
			ServiceDUT: &solver.ServiceDUT{
				AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1"}},
				Services: map[string]*tpb.Service{
					"gnmi": {
						Name:   "gnmi",
						Inside: 9339,
					},
				},
			},
		}
		_, err := d.DialGRPCWithPort(context.Background(), 1234)
		if err == nil {
			t.Fatal("DialGRPCWithPort() got nil error, want non-nil")
		}
		if wantErr := "service for port 1234 not found on DUT \"dut1\""; err.Error() != wantErr {
			t.Errorf("DialGRPCWithPort() got error %q, want %q", err, wantErr)
		}
	})
}

type fakeTopoManager struct {
	top       *tpb.Topology
	gotResets int
	resetErr  error
	pushErr   error
}

func (m *fakeTopoManager) Nodes() map[string]node.Node {
	nodes := make(map[string]node.Node)
	for _, npb := range m.top.GetNodes() {
		nodes[npb.GetName()] = &fakeNode{npb: npb}
	}
	return nodes
}

func (m *fakeTopoManager) Show(context.Context) (*cpb.ShowTopologyResponse, error) {
	return &cpb.ShowTopologyResponse{Topology: m.top}, nil
}

func (m *fakeTopoManager) ConfigPush(context.Context, string, io.Reader) error {
	return m.pushErr
}

func (m *fakeTopoManager) ResetCfg(context.Context, string) error {
	if m.resetErr == nil {
		m.gotResets++
	}
	return m.resetErr
}

type fakeNode struct {
	node.Node
	npb *tpb.Node
}

func (n *fakeNode) GetProto() *tpb.Node {
	return n.npb
}
