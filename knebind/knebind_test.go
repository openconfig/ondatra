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
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"errors"
	"io"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/openconfig/kne/topo/node"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/knebind/creds"
	"github.com/openconfig/ondatra/knebind/solver"
	"golang.org/x/crypto/ssh"
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
				1234: {Name: "gnmi"},
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

	wantDUTServices := map[string]*tpb.Service{"gnmi": {Name: "gnmi"}}
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
		desc: "missing gnpsi",
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
			if _, err := d.DialGNPSI(context.Background()); err == nil {
				t.Fatalf("DialGNPSI() got unexpected error: %v", err)
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
					9339: {
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
					9339: {
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
					9336: {
						Name:      "p4rt",
						Outside:   9336,
						OutsideIp: "1.1.1.1",
					},
					9339: {
						Name:      "gnmi",
						Outside:   9339,
						OutsideIp: "1.1.1.1",
					},
					4242: {
						Name:      "gribi",
						Outside:   4242,
						OutsideIp: "1.1.1.1",
					},
					9340: {
						Name:      "gnsi",
						Outside:   9339,
						OutsideIp: "1.1.1.1",
					},
					4739: {
						Name:      "gnpsi",
						Outside:   4739,
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
			if _, err := d.DialGNPSI(context.Background()); err != nil {
				t.Fatalf("DialGNPSI() got unexpected error: %v", err)
			}
			if _, err := d.DialGNSI(context.Background()); err != nil {
				t.Fatalf("DialGNSI() got unexpected error: %v", err)
			}
			if _, err := d.DialP4RT(context.Background()); err != nil {
				t.Fatalf("DialP4RT() got unexpected error: %v", err)
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

type testSSHServer struct {
	ip        string
	port      uint32
	hostKey   []byte
	clientPEM []byte
	clientPub []byte
	certBytes []byte
}

func startTestSSHServer(t *testing.T) *testSSHServer {
	t.Helper()

	_, hostPriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate host key: %v", err)
	}
	hostSigner, err := ssh.NewSignerFromKey(hostPriv)
	if err != nil {
		t.Fatalf("failed to create host signer: %v", err)
	}

	_, clientPriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate client key: %v", err)
	}
	clientSigner, err := ssh.NewSignerFromKey(clientPriv)
	if err != nil {
		t.Fatalf("failed to create client signer: %v", err)
	}
	clientPEMBlock, err := ssh.MarshalPrivateKey(clientPriv, "")
	if err != nil {
		t.Fatalf("failed to marshal client private key: %v", err)
	}
	clientPEM := pem.EncodeToMemory(clientPEMBlock)

	_, caPriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate CA key: %v", err)
	}
	caSigner, err := ssh.NewSignerFromKey(caPriv)
	if err != nil {
		t.Fatalf("failed to create CA signer: %v", err)
	}
	cert := &ssh.Certificate{
		Key:             clientSigner.PublicKey(),
		CertType:        ssh.UserCert,
		ValidPrincipals: []string{"certuser"},
		ValidAfter:      0,
		ValidBefore:     ssh.CertTimeInfinity,
	}
	if err := cert.SignCert(rand.Reader, caSigner); err != nil {
		t.Fatalf("failed to sign certificate: %v", err)
	}

	serverCfg := &ssh.ServerConfig{
		PasswordCallback: func(conn ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if conn.User() == "testuser" && string(pass) == "testpass" {
				return nil, nil
			}
			return nil, errors.New("password rejected")
		},
		KeyboardInteractiveCallback: func(conn ssh.ConnMetadata, client ssh.KeyboardInteractiveChallenge) (*ssh.Permissions, error) {
			ans, err := client(conn.User(), "", []string{"Password: "}, []bool{false})
			if err != nil {
				return nil, err
			}
			if conn.User() == "kiuser" && len(ans) == 1 && ans[0] == "kipass" {
				return nil, nil
			}
			return nil, errors.New("keyboard-interactive rejected")
		},
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			if c, ok := key.(*ssh.Certificate); ok {
				if conn.User() == "certuser" && bytes.Equal(c.SignatureKey.Marshal(), caSigner.PublicKey().Marshal()) {
					return nil, nil
				}
				return nil, errors.New("certificate rejected")
			}
			if conn.User() == "keyuser" && bytes.Equal(key.Marshal(), clientSigner.PublicKey().Marshal()) {
				return nil, nil
			}
			return nil, errors.New("public key rejected")
		},
	}
	serverCfg.AddHostKey(hostSigner)

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen on loopback: %v", err)
	}
	t.Cleanup(func() { listener.Close() })

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				_, chans, reqs, err := ssh.NewServerConn(c, serverCfg)
				if err != nil {
					return
				}
				go ssh.DiscardRequests(reqs)
				for newChan := range chans {
					if newChan.ChannelType() != "session" {
						newChan.Reject(ssh.UnknownChannelType, "unknown channel type")
						continue
					}
					ch, requests, err := newChan.Accept()
					if err != nil {
						continue
					}
					go func(ch ssh.Channel, in <-chan *ssh.Request) {
						defer ch.Close()
						for req := range in {
							if req.Type != "exec" {
								if req.WantReply {
									req.Reply(false, nil)
								}
								continue
							}
							var payload struct{ Cmd string }
							if err := ssh.Unmarshal(req.Payload, &payload); err != nil {
								req.Reply(false, nil)
								return
							}
							req.Reply(true, nil)
							var exitStatus uint32
							switch payload.Cmd {
							case "echo hello":
								ch.Write([]byte("hello\n"))
							case "fail":
								ch.Write([]byte("failed output\n"))
								exitStatus = 1
							case "hang":
								continue
							default:
								ch.Write([]byte("unknown\n"))
							}
							ch.SendRequest("exit-status", false, ssh.Marshal(struct{ Status uint32 }{Status: exitStatus}))
							return
						}
					}(ch, requests)
				}
			}(conn)
		}
	}()

	tcpAddr := listener.Addr().(*net.TCPAddr)
	return &testSSHServer{
		ip:        tcpAddr.IP.String(),
		port:      uint32(tcpAddr.Port),
		hostKey:   ssh.MarshalAuthorizedKey(hostSigner.PublicKey()),
		clientPEM: clientPEM,
		clientPub: ssh.MarshalAuthorizedKey(clientSigner.PublicKey()),
		certBytes: ssh.MarshalAuthorizedKey(cert),
	}
}

func TestDialSSH_Success(t *testing.T) {
	srv := startTestSSHServer(t)
	dut := &kneDUT{
		ServiceDUT: &solver.ServiceDUT{
			AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1"}},
			Services: map[string]*tpb.Service{
				"ssh": {
					Name:      "ssh",
					OutsideIp: srv.ip,
					Outside:   srv.port,
				},
			},
		},
		bind: &Bind{cfg: &Config{}},
	}

	tests := []struct {
		name string
		auth binding.SSHAuth
	}{
		{
			name: "password_auth",
			auth: binding.PasswordAuth{User: "testuser", Password: "testpass"},
		},
		{
			name: "keyboard_interactive_auth",
			auth: binding.PasswordAuth{User: "kiuser", Password: "kipass"},
		},
		{
			name: "key_auth",
			auth: binding.KeyAuth{User: "keyuser", Key: srv.clientPEM},
		},
		{
			name: "certificate_auth",
			auth: binding.CertificateAuth{
				User:        "certuser",
				PrivateKey:  srv.clientPEM,
				Certificate: srv.certBytes,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client, err := dut.DialSSH(t.Context(), tc.auth)
			if err != nil {
				t.Fatalf("DialSSH() unexpected error: %v", err)
			}
			defer client.Close()

			if diff := cmp.Diff(srv.hostKey, client.HostKey()); diff != "" {
				t.Errorf("HostKey() mismatch (-want +got):\n%s", diff)
			}

			res, err := client.RunCommand(t.Context(), "echo hello")
			if err != nil {
				t.Fatalf("RunCommand(\"echo hello\") unexpected error: %v", err)
			}
			if got, want := res.Output(), "hello\n"; got != want {
				t.Errorf("RunCommand(\"echo hello\") Output() = %q, want %q", got, want)
			}
			if got := res.Error(); got != "" {
				t.Errorf("RunCommand(\"echo hello\") Error() = %q, want empty", got)
			}

			failRes, err := client.RunCommand(t.Context(), "fail")
			if err != nil {
				t.Fatalf("RunCommand(\"fail\") unexpected error: %v", err)
			}
			if got, want := failRes.Output(), "failed output\n"; got != want {
				t.Errorf("RunCommand(\"fail\") Output() = %q, want %q", got, want)
			}
			if got := failRes.Error(); got == "" {
				t.Errorf("RunCommand(\"fail\") Error() = empty, want non-empty exit error")
			}

			canceledCtx, cancel := context.WithCancel(t.Context())
			cancel()
			if _, err := client.RunCommand(canceledCtx, "echo hello"); err == nil {
				t.Errorf("RunCommand() with pre-canceled context = nil, want error")
			}

			hangCtx, cancelHang := context.WithTimeout(t.Context(), 50*time.Millisecond)
			defer cancelHang()
			if _, err := client.RunCommand(hangCtx, "hang"); err == nil {
				t.Errorf("RunCommand(\"hang\") with expiring context = nil, want error")
			}

			if err := client.Close(); err != nil {
				t.Fatalf("Close() unexpected error: %v", err)
			}
			if _, err := client.RunCommand(t.Context(), "echo hello"); err == nil {
				t.Errorf("RunCommand() on closed client = nil, want error")
			}
		})
	}
}

func TestDialSSH_Errors(t *testing.T) {
	srv := startTestSSHServer(t)
	_, otherPriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("failed to generate other key: %v", err)
	}
	otherPEMBlock, err := ssh.MarshalPrivateKey(otherPriv, "")
	if err != nil {
		t.Fatalf("failed to marshal other private key: %v", err)
	}
	otherPEM := pem.EncodeToMemory(otherPEMBlock)

	dutWithSSH := &kneDUT{
		ServiceDUT: &solver.ServiceDUT{
			AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1"}},
			Services: map[string]*tpb.Service{
				"ssh": {
					Name:      "ssh",
					OutsideIp: srv.ip,
					Outside:   srv.port,
				},
			},
		},
		bind: &Bind{cfg: &Config{}},
	}
	dutWithoutSSH := &kneDUT{
		ServiceDUT: &solver.ServiceDUT{
			AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1"}},
			Services:    map[string]*tpb.Service{},
		},
		bind: &Bind{cfg: &Config{}},
	}

	tests := []struct {
		name    string
		dut     *kneDUT
		auth    binding.SSHAuth
		wantErr string
	}{
		{
			name:    "missing_ssh_service",
			dut:     dutWithoutSSH,
			auth:    binding.PasswordAuth{User: "testuser", Password: "testpass"},
			wantErr: "service \"ssh\" not found",
		},
		{
			name:    "nil_auth",
			dut:     dutWithSSH,
			auth:    nil,
			wantErr: "unsupported SSH auth type",
		},
		{
			name:    "invalid_password",
			dut:     dutWithSSH,
			auth:    binding.PasswordAuth{User: "testuser", Password: "wrongpassword"},
			wantErr: "could not dial SSH server",
		},
		{
			name:    "invalid_private_key_in_key_auth",
			dut:     dutWithSSH,
			auth:    binding.KeyAuth{User: "keyuser", Key: []byte("not-a-private-key")},
			wantErr: "parsing private key",
		},
		{
			name: "invalid_private_key_in_cert_auth",
			dut:  dutWithSSH,
			auth: binding.CertificateAuth{
				User:        "certuser",
				PrivateKey:  []byte("not-a-private-key"),
				Certificate: srv.certBytes,
			},
			wantErr: "parsing private key",
		},
		{
			name: "invalid_certificate_in_cert_auth",
			dut:  dutWithSSH,
			auth: binding.CertificateAuth{
				User:        "certuser",
				PrivateKey:  srv.clientPEM,
				Certificate: []byte("not-a-certificate"),
			},
			wantErr: "parsing certificate",
		},
		{
			name: "public_key_instead_of_certificate_in_cert_auth",
			dut:  dutWithSSH,
			auth: binding.CertificateAuth{
				User:        "certuser",
				PrivateKey:  srv.clientPEM,
				Certificate: srv.clientPub,
			},
			wantErr: "want *ssh.Certificate",
		},
		{
			name: "mismatched_key_and_certificate_in_cert_auth",
			dut:  dutWithSSH,
			auth: binding.CertificateAuth{
				User:        "certuser",
				PrivateKey:  otherPEM,
				Certificate: srv.certBytes,
			},
			wantErr: "creating certificate signer",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client, err := tc.dut.DialSSH(t.Context(), tc.auth)
			if err == nil {
				client.Close()
				t.Fatalf("DialSSH() = nil, want error containing %q", tc.wantErr)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Errorf("DialSSH() error = %v, want error containing %q", err, tc.wantErr)
			}
		})
	}
}

func TestDialSSH_ContextCanceled(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen on loopback: %v", err)
	}
	t.Cleanup(func() { listener.Close() })
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			t.Cleanup(func() { conn.Close() })
		}
	}()

	tcpAddr := listener.Addr().(*net.TCPAddr)
	dut := &kneDUT{
		ServiceDUT: &solver.ServiceDUT{
			AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1"}},
			Services: map[string]*tpb.Service{
				"ssh": {
					Name:      "ssh",
					OutsideIp: tcpAddr.IP.String(),
					Outside:   uint32(tcpAddr.Port),
				},
			},
		},
		bind: &Bind{cfg: &Config{}},
	}
	auth := binding.PasswordAuth{User: "testuser", Password: "testpass"}

	canceledCtx, cancel := context.WithCancel(t.Context())
	cancel()
	if client, err := dut.DialSSH(canceledCtx, auth); err == nil {
		client.Close()
		t.Errorf("DialSSH() with pre-canceled context = nil, want error")
	}

	handshakeCtx, cancelHandshake := context.WithTimeout(t.Context(), 50*time.Millisecond)
	defer cancelHandshake()
	if client, err := dut.DialSSH(handshakeCtx, auth); err == nil {
		client.Close()
		t.Errorf("DialSSH() with stalled handshake = nil, want error")
	}
}

func TestDialCLI(t *testing.T) {
	srv := startTestSSHServer(t)
	dut := &kneDUT{
		ServiceDUT: &solver.ServiceDUT{
			AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1"}},
			Services: map[string]*tpb.Service{
				"ssh": {
					Name:      "ssh",
					OutsideIp: srv.ip,
					Outside:   srv.port,
				},
			},
		},
		bind: &Bind{
			cfg: &Config{
				Username: "testuser",
				Password: "testpass",
			},
		},
	}
	cli, err := dut.DialCLI(t.Context())
	if err != nil {
		t.Fatalf("DialCLI() unexpected error: %v", err)
	}
	res, err := cli.RunCommand(t.Context(), "echo hello")
	if err != nil {
		t.Fatalf("RunCommand(\"echo hello\") unexpected error: %v", err)
	}
	if got, want := res.Output(), "hello\n"; got != want {
		t.Errorf("RunCommand(\"echo hello\") Output() = %q, want %q", got, want)
	}
}

func TestDialCLI_Errors(t *testing.T) {
	srv := startTestSSHServer(t)
	tests := []struct {
		name string
		dut  *kneDUT
	}{
		{
			name: "missing_credentials",
			dut: &kneDUT{
				ServiceDUT: &solver.ServiceDUT{
					AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1"}},
					Services: map[string]*tpb.Service{
						"ssh": {
							Name:      "ssh",
							OutsideIp: srv.ip,
							Outside:   srv.port,
						},
					},
				},
				bind: &Bind{cfg: &Config{}},
			},
		},
		{
			name: "dial_error",
			dut: &kneDUT{
				ServiceDUT: &solver.ServiceDUT{
					AbstractDUT: &binding.AbstractDUT{Dims: &binding.Dims{Name: "dut1"}},
					Services:    map[string]*tpb.Service{},
				},
				bind: &Bind{
					cfg: &Config{
						Username: "testuser",
						Password: "testpass",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cli, err := tc.dut.DialCLI(t.Context())
			if err != nil {
				t.Fatalf("DialCLI() unexpected error: %v", err)
			}
			if _, err := cli.RunCommand(t.Context(), "echo hello"); err == nil {
				t.Errorf("RunCommand() = nil, want error")
			}
		})
	}
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
