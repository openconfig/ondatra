package ondatra

import (
	"golang.org/x/net/context"
	"fmt"
	"testing"

	"github.com/openconfig/ondatra/internal/ate"
	"github.com/openconfig/ondatra/internal/reservation"

	opb "github.com/openconfig/ondatra/proto"
)

var lagCount uint32

// Topology is an ATE topology API.
type Topology struct {
	ate *reservation.ATE
}

func (tp *Topology) String() string {
	return fmt.Sprintf("{ate: %s}", tp.ate)
}

// New returns a new ATE topology.
func (tp *Topology) New() *ATETopology {
	return &ATETopology{ate: tp.ate, top: &opb.Topology{}}
}

// ATETopology is an ATE topology.
type ATETopology struct {
	ate *reservation.ATE
	top *opb.Topology
}

func (at *ATETopology) String() string {
	return fmt.Sprintf("{ate: %s, topology: %s}", at.ate, at.top)
}

// AddInterface adds and returns an interface with the specified name.
// The returned interface has an Ethernet configuration with an MTU of 1500.
func (at *ATETopology) AddInterface(name string) *Interface {
	ipb := &opb.InterfaceConfig{
		Name:     name,
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
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
func (at *ATETopology) AddLAG() *LAG {
	lagCount = lagCount + 1
	lpb := &opb.Lag{Name: fmt.Sprintf("lag%d", lagCount)}
	at.top.Lags = append(at.top.Lags, lpb)
	return &LAG{pb: lpb}
}

// Push replaces the topology to the ATE with this one.
// Currently running protocols will stop.
func (at *ATETopology) Push(t testing.TB) *ATETopology {
	t.Helper()
	logAction(t, "Pushing topology to %s", at.ate)
	if err := ate.PushTopology(context.Background(), at.ate, at.top); err != nil {
		t.Fatalf("Push(t) on %s: %v", at, err)
	}
	return at
}

// Update updates the topology on the ATE to this one.
// Currently running protocols will continue running.
func (at *ATETopology) Update(t testing.TB) {
	t.Helper()
	logAction(t, "Updating topology to %s", at.ate)
	if err := ate.UpdateTopology(context.Background(), at.ate, at.top, false); err != nil {
		t.Fatalf("Update(t) on %s: %v", at, err)
	}
}

// UpdateBGPPeerStates is equivalent to Update() but only updates the BGP peer state.
// This is provided as a temporary workaround for the high overhead of Update().
// TODO: Remove this method once new Ixia config binding is used.
func (at *ATETopology) UpdateBGPPeerStates(t testing.TB) {
	t.Helper()
	if err := ate.UpdateTopology(context.Background(), at.ate, at.top, true); err != nil {
		t.Fatalf("UpdateBGPPeerState(t) on %s: %v", at, err)
	}
}

// StartProtocols starts the control plane protocols on the ATE.
func (at *ATETopology) StartProtocols(t testing.TB) *ATETopology {
	t.Helper()
	logAction(t, "Starting protocols on %s", at.ate)
	if err := ate.StartProtocols(context.Background(), at.ate); err != nil {
		t.Fatalf("StartProtocols(t) on %s: %v", at, err)
	}
	return at
}

// StopProtocols stops the control plane protocols on the ATE.
func (at *ATETopology) StopProtocols(t testing.TB) *ATETopology {
	t.Helper()
	logAction(t, "Stopping protocols to %s", at.ate)
	if err := ate.StopProtocols(context.Background(), at.ate); err != nil {
		t.Fatalf("StopProtocols(t) on %s: %v", at, err)
	}
	return at
}
