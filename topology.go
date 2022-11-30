// Copyright 2020 Google LLC
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

package ondatra

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/ate"
	"github.com/openconfig/ondatra/internal/events"

	opb "github.com/openconfig/ondatra/proto"
)

// Topology is the ATE Topology API.
type Topology struct {
	ate binding.ATE
}

func (tp *Topology) String() string {
	return fmt.Sprintf("Topology%+v", *tp)
}

// New returns a new ATE topology.
func (tp *Topology) New() *ATETopology {
	return &ATETopology{ate: tp.ate, top: &ate.Topology{}}
}

// ATETopology is an ATE topology.
type ATETopology struct {
	ate binding.ATE
	top *ate.Topology
}

func (at *ATETopology) String() string {
	return fmt.Sprintf("{ate: %s, topology: %s}", at.ate, at.top)
}

// AddInterface adds and returns an interface with the specified name.
// The returned interface has an Ethernet configuration with the following defaults:
// MTU: 1500
// FEC: enabled
func (at *ATETopology) AddInterface(name string) *Interface {
	ipb := &opb.InterfaceConfig{
		Name: name,
		Ethernet: &opb.EthernetConfig{
			Mtu: 1500,
			Fec: &opb.Fec{Enabled: true},
		},
	}
	at.top.Interfaces = append(at.top.Interfaces, ipb)
	return &Interface{pb: ipb}
}

// ClearInterfaces clear interfaces from the ATE topology.
func (at *ATETopology) ClearInterfaces() *ATETopology {
	at.top.Interfaces = nil
	return at
}

// Interfaces returns a map of interface names to interfaces.
func (at *ATETopology) Interfaces() map[string]*Interface {
	intfs := make(map[string]*Interface)
	for _, pb := range at.top.Interfaces {
		intfs[pb.GetName()] = &Interface{pb}
	}
	return intfs
}

// AddLAG adds and returns a new LAG.
// By default the LAG has LACP enabled.
func (at *ATETopology) AddLAG(name string) *LAG {
	lpb := &opb.Lag{
		Name: name,
		Lacp: &opb.Lag_Lacp{Enabled: true},
	}
	at.top.LAGs = append(at.top.LAGs, lpb)
	return &LAG{pb: lpb}
}

// Push replaces the topology to the ATE with this one.
// Currently running protocols will stop.
func (at *ATETopology) Push(t testing.TB) *ATETopology {
	t.Helper()
	t = events.ActionStarted(t, "Pushing topology to %s", at.ate)
	if err := ate.PushTopology(context.Background(), at.ate, at.top); err != nil {
		t.Fatalf("Push(t) on %s: %v", at, err)
	}
	return at
}

// Update updates the topology on the ATE to this one.
// Currently running protocols will continue running.
func (at *ATETopology) Update(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Updating topology to %s", at.ate)
	if err := ate.UpdateTopology(context.Background(), at.ate, at.top, false); err != nil {
		t.Fatalf("Update(t) on %s: %v", at, err)
	}
}

// UpdateBGPPeerStates is equivalent to Update() but only updates the BGP peer state.
// This is provided as a temporary workaround for the high overhead of Update().
// TODO(team): Remove this method once new Ixia config binding is used.
func (at *ATETopology) UpdateBGPPeerStates(t testing.TB) {
	t.Helper()
	if err := ate.UpdateTopology(context.Background(), at.ate, at.top, true); err != nil {
		t.Fatalf("UpdateBGPPeerState(t) on %s: %v", at, err)
	}
}

// UpdateNetworks is equivalent to Update() but only updates the Network config on the fly.
func (at *ATETopology) UpdateNetworks(t testing.TB) {
	t.Helper()
	if err := ate.UpdateNetworks(context.Background(), at.ate, at.top); err != nil {
		t.Fatalf("UpdateNetworks(t) on %s: %v", at, err)
	}
}

// StartProtocols starts the control plane protocols on the ATE.
func (at *ATETopology) StartProtocols(t testing.TB) *ATETopology {
	t.Helper()
	t = events.ActionStarted(t, "Starting protocols on %s", at.ate)
	if err := ate.StartProtocols(context.Background(), at.ate); err != nil {
		t.Fatalf("StartProtocols(t) on %s: %v", at, err)
	}
	return at
}

// StopProtocols stops the control plane protocols on the ATE.
func (at *ATETopology) StopProtocols(t testing.TB) *ATETopology {
	t.Helper()
	t = events.ActionStarted(t, "Stopping protocols to %s", at.ate)
	if err := ate.StopProtocols(context.Background(), at.ate); err != nil {
		t.Fatalf("StopProtocols(t) on %s: %v", at, err)
	}
	return at
}
