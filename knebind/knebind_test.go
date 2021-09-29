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
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/internal/reservation"

	kpb "github.com/google/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

func TestReserve(t *testing.T) {
	topo := &kpb.Topology{
		Nodes: []*kpb.Node{{
			Name:     "node1",
			Type:     kpb.Node_ARISTA_CEOS,
			Services: map[uint32]*kpb.Service{1234: {Name: "gnmi"}},
		}, {
			Name:     "node2",
			Type:     kpb.Node_CISCO_CXR,
			Services: map[uint32]*kpb.Service{2345: {Name: "gnmi"}},
		}, {
			Name:     "node3",
			Type:     kpb.Node_JUNIPER_CEVO,
			Services: map[uint32]*kpb.Service{3456: {Name: "gnmi"}},
		}, {
			Name: "node4",
			Type: kpb.Node_IXIA_TG,
		}},
		Links: []*kpb.Link{
			{ANode: "node1", AInt: "intf1", ZNode: "node2", ZInt: "intf1"},
			{ANode: "node2", AInt: "intf2", ZNode: "node3", ZInt: "intf2"},
			{ANode: "node1", AInt: "intf2", ZNode: "node4", ZInt: "intf1"},
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
	fetchTopo = func(*Config) (*kpb.Topology, error) {
		return topo, nil
	}
	wantDUT1 := &reservation.DUT{&reservation.Dims{
		Name:            "node1",
		Vendor:          opb.Device_ARISTA,
		HardwareModel:   "ARISTA_CEOS",
		SoftwareVersion: "ARISTA_CEOS",
		Ports: map[string]*reservation.Port{
			"port1": {Name: "intf1"},
			"port2": {Name: "intf2"},
		},
	}}
	wantDUT2 := &reservation.DUT{&reservation.Dims{
		Name:            "node2",
		Vendor:          opb.Device_CISCO,
		HardwareModel:   "CISCO_CXR",
		SoftwareVersion: "CISCO_CXR",
		Ports: map[string]*reservation.Port{
			"port1": {Name: "intf1"},
			"port2": {Name: "intf2"},
		},
	}}
	wantDUT3 := &reservation.DUT{&reservation.Dims{
		Name:            "node3",
		Vendor:          opb.Device_JUNIPER,
		HardwareModel:   "JUNIPER_CEVO",
		SoftwareVersion: "JUNIPER_CEVO",
		Ports: map[string]*reservation.Port{
			"port1": {Name: "intf2"},
		},
	}}
	wantATE := &reservation.ATE{&reservation.Dims{
		Name:            "node4",
		Vendor:          opb.Device_IXIA,
		HardwareModel:   "IXIA_TG",
		SoftwareVersion: "IXIA_TG",
		Ports: map[string]*reservation.Port{
			"port1": {Name: "intf1"},
		},
	}}

	tests := []struct {
		name    string
		tb      *opb.Testbed
		wantRes *reservation.Reservation
	}{{
		name: "one dut",
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
		name: "one ate",
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
		name: "two duts",
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
		name: "dut and ate",
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
		name: "three duts",
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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b, err := New(&Config{})
			if err != nil {
				t.Fatalf("New failed: %v", err)
			}
			gotRes, err := b.Reserve(context.Background(), test.tb, time.Minute, time.Minute)
			if err != nil {
				t.Fatalf("Reserve() got error: %v", err)
			}
			if gotRes.ID == "" {
				t.Errorf("Reserve() got reservation missing ID: %v", gotRes)
			}
			gotRes.ID = ""
			if diff := cmp.Diff(test.wantRes, gotRes); diff != "" {
				t.Errorf("Reserve() got unexpected diff in reservation (-want,+got): %s", diff)
			}
		})
	}
}

func TestReserveErrors(t *testing.T) {
	tests := []struct {
		name    string
		tb      *opb.Testbed
		topo    *kpb.Topology
		wantErr string
	}{{
		name:    "too few nodes",
		tb:      &opb.Testbed{Duts: []*opb.Device{{Id: "dut1"}}},
		wantErr: "Not enough nodes",
	}, {
		name:    "too few links",
		tb:      &opb.Testbed{Links: []*opb.Link{{A: "dut1:port1", B: "dut2:port1"}}},
		wantErr: "Not enough links",
	}, {
		name: "missing gnmi",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{Id: "dut1"}},
		},
		topo: &kpb.Topology{
			Nodes: []*kpb.Node{{
				Name: "node1",
				Type: kpb.Node_ARISTA_CEOS,
			}},
		},
		wantErr: "GNMI",
	}, {
		name: "no match for DUT",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
			}},
		},
		topo: &kpb.Topology{
			Nodes: []*kpb.Node{{
				Name: "node1",
				Type: kpb.Node_CISCO_CXR,
			}},
		},
		wantErr: "No node in KNE topology to match testbed",
	}, {
		name: "no match for ATE",
		tb: &opb.Testbed{
			Ates: []*opb.Device{{
				Id: "ate1",
			}},
		},
		topo: &kpb.Topology{
			Nodes: []*kpb.Node{{
				Name: "node1",
				Type: kpb.Node_CISCO_CXR,
			}},
		},
		wantErr: "No node in KNE topology to match testbed",
	}, {
		name: "no node combination",
		tb: &opb.Testbed{
			Duts: []*opb.Device{{
				Id:     "dut1",
				Vendor: opb.Device_ARISTA,
			}, {
				Id:     "dut2",
				Vendor: opb.Device_ARISTA,
			}},
		},
		topo: &kpb.Topology{
			Nodes: []*kpb.Node{
				{Name: "node1", Type: kpb.Node_ARISTA_CEOS},
				{Name: "node2", Type: kpb.Node_CISCO_CXR},
			},
		},
		wantErr: "No combination of nodes",
	}, {
		name: "no link combination",
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
		topo: &kpb.Topology{
			Nodes: []*kpb.Node{
				{Name: "node1", Type: kpb.Node_ARISTA_CEOS},
				{Name: "node2", Type: kpb.Node_ARISTA_CEOS},
				{Name: "node3", Type: kpb.Node_JUNIPER_VMX},
				{Name: "node4", Type: kpb.Node_JUNIPER_VMX},
			},
			Links: []*kpb.Link{
				{ANode: "node1", AInt: "eth1", ZNode: "node3", ZInt: "eth1"},
				{ANode: "node2", AInt: "eth1", ZNode: "node4", ZInt: "eth1"},
			},
		},
		wantErr: "No KNE topology",
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fetchTopo = func(*Config) (*kpb.Topology, error) {
				return test.topo, nil
			}
			b, err := New(&Config{})
			if err != nil {
				t.Fatalf("New failed: %v", err)
			}
			_, gotErr := b.Reserve(context.Background(), test.tb, time.Minute, time.Minute)
			if diff := errdiff.Substring(gotErr, test.wantErr); diff != "" {
				t.Errorf("Reserve() got unexpected error diff: %s", diff)
			}
		})
	}
}
