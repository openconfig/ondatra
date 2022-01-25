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
	"golang.org/x/crypto/ssh"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/knebind/solver"

	tpb "github.com/google/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

func TestReserve(t *testing.T) {
	topo := &tpb.Topology{
		Nodes: []*tpb.Node{{
			Name:     "node1",
			Type:     tpb.Node_ARISTA_CEOS,
			Services: map[uint32]*tpb.Service{1234: {Name: "gnmi"}},
			Interfaces: map[string]*tpb.Interface{
				"eth1": {
					Name: "Ethernet1",
				},
				"eth2": {
					Name: "Ethernet2",
				},
			},
		}, {
			Name:     "node2",
			Type:     tpb.Node_CISCO_CXR,
			Services: map[uint32]*tpb.Service{2345: {Name: "gnmi"}},
			Interfaces: map[string]*tpb.Interface{
				"eth1": {},
				"eth2": {},
			},
		}, {
			Name:     "node3",
			Type:     tpb.Node_JUNIPER_CEVO,
			Services: map[uint32]*tpb.Service{3456: {Name: "gnmi"}},
			Interfaces: map[string]*tpb.Interface{
				"eth1": {},
			},
		}, {
			Name: "node4",
			Type: tpb.Node_IXIA_TG,
			Interfaces: map[string]*tpb.Interface{
				"eth1": {},
			},
		}},
		Links: []*tpb.Link{
			{ANode: "node1", AInt: "eth1", ZNode: "node2", ZInt: "eth1"},
			{ANode: "node2", AInt: "eth2", ZNode: "node3", ZInt: "eth1"},
			{ANode: "node1", AInt: "eth2", ZNode: "node4", ZInt: "eth1"},
		},
	}
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
	fetchTopoFn = func(*Config) (*tpb.Topology, error) {
		return topo, nil
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
			b, err := New(&Config{})
			if err != nil {
				t.Fatalf("New failed: %v", err)
			}
			gotRes, err := b.Reserve(context.Background(), tt.tb, time.Minute, time.Minute)
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
	tests := []struct {
		desc        string
		tb          *opb.Testbed
		topo        *tpb.Topology
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
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name: "node1",
				Type: tpb.Node_ARISTA_CEOS,
			}},
		},
		wantGNMIErr: "gnmi",
	}, {
		desc: "no match for DUT",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
			}},
		},
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name: "node1",
				Type: tpb.Node_CISCO_CXR,
			}},
		},
		wantErr: "No node in KNE topology to match testbed",
	}, {
		desc: "no match for ATE",
		tb: &opb.Testbed{
			Ates: []*opb.Device{{
				Id: "ate1",
			}},
		},
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name: "node1",
				Type: tpb.Node_CISCO_CXR,
			}},
		},
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
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{
				{Name: "node1", Type: tpb.Node_ARISTA_CEOS},
				{Name: "node2", Type: tpb.Node_CISCO_CXR},
			},
		},
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
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{
				{Name: "node1", Type: tpb.Node_ARISTA_CEOS},
				{Name: "node2", Type: tpb.Node_ARISTA_CEOS},
				{Name: "node3", Type: tpb.Node_JUNIPER_VMX},
				{Name: "node4", Type: tpb.Node_JUNIPER_VMX},
			},
			Links: []*tpb.Link{
				{ANode: "node1", AInt: "eth1", ZNode: "node3", ZInt: "eth1"},
				{ANode: "node2", AInt: "eth1", ZNode: "node4", ZInt: "eth1"},
			},
		},
		wantErr: "No KNE topology",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			fetchTopoFn = func(*Config) (*tpb.Topology, error) {
				return tt.topo, nil
			}
			b, err := New(&Config{})
			if err != nil {
				t.Fatalf("New failed: %v", err)
			}
			res, gotErr := b.Reserve(context.Background(), tt.tb, time.Minute, time.Minute)
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
			_, gnmiErr := b.DialGNMI(context.Background(), d)
			if diff := errdiff.Substring(gnmiErr, tt.wantGNMIErr); diff != "" {
				t.Errorf("DialGNMI() got unexpected error diff: %s", diff)
			}
		})
	}
}

func TestServices(t *testing.T) {
	tests := []struct {
		desc         string
		tb           *opb.Testbed
		topo         *tpb.Topology
		serviceCheck func(t *testing.T, b binding.Binding, d *reservation.DUT)
	}{{
		desc: "missing gnmi",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{Id: "dut1"}},
		},
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name: "node1",
				Type: tpb.Node_ARISTA_CEOS,
			}},
		},
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
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name:     "node1",
				Type:     tpb.Node_CISCO_CXR,
				Services: map[uint32]*tpb.Service{9339: {Name: "gnmi", Outside: 9339, OutsideIp: "1.1.1.1"}},
			}},
		},
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
		topo: &tpb.Topology{
			Nodes: []*tpb.Node{{
				Name: "node1",
				Type: tpb.Node_CISCO_CXR,
				Services: map[uint32]*tpb.Service{
					9336: {Name: "p4rt", Outside: 9336, OutsideIp: "1.1.1.1"},
					9339: {Name: "gnmi", Outside: 9339, OutsideIp: "1.1.1.1"},
				},
			}},
		},
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
			fetchTopoFn = func(*Config) (*tpb.Topology, error) {
				return tt.topo, nil
			}
			b, err := New(&Config{})
			if err != nil {
				t.Fatalf("New failed: %v", err)
			}
			res, err := b.Reserve(context.Background(), tt.tb, time.Minute, time.Minute)
			if err != nil {
				t.Fatalf("Reserve() failed: %v", err)
			}
			d, err := res.DUT("dut1")
			if err != nil {
				t.Fatalf("Node %q not found in topology", "node1")
			}
			tt.serviceCheck(t, b, d)
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
	sshExecFn = func(addr string, cfg *ssh.ClientConfig, cmd string) (_ string, rerr error) {
		return "", nil
	}

	tests := []struct {
		desc    string
		dut     *reservation.DUT
		opts    *binding.ConfigOptions
		wantErr string
	}{{
		desc: "success",
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
		desc:    "no replace support",
		dut:     &reservation.DUT{&reservation.Dims{Name: dutName, Vendor: opb.Device_ARISTA}},
		opts:    &binding.ConfigOptions{},
		wantErr: "config replace",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			err := bind.PushConfig(context.Background(), tt.dut, "my config", tt.opts)
			if (err == nil) != (tt.wantErr == "") || (err != nil && !strings.Contains(err.Error(), tt.wantErr)) {
				t.Fatalf("PushConfig got error %v, want %v", err, tt.wantErr)
			}
		})
	}
}
