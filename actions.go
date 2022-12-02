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

package ondatra

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/ate"
	"github.com/openconfig/ondatra/internal/events"
	"github.com/openconfig/ondatra/ixnet"
)

// Actions is the ATE Actions API.
type Actions struct {
	ate binding.ATE
}

func (a *Actions) String() string {
	return fmt.Sprintf("Actions%+v", *a)
}

// NewSetPortState returns a new SetPortState action.
func (a *Actions) NewSetPortState() *SetPortState {
	return &SetPortState{ate: a.ate}
}

// SetPortState is an action to set the state of a port on an ATE.
type SetPortState struct {
	ate      binding.ATE
	portName string
	enabled  *bool
}

func (s *SetPortState) String() string {
	return fmt.Sprintf("SetPortState%+v", *s)
}

// WithPort sets the port whose state will be changed.
func (s *SetPortState) WithPort(port *Port) *SetPortState {
	s.portName = port.Name()
	return s
}

// WithEnabled sets the desired state of the port.
func (s *SetPortState) WithEnabled(enabled bool) *SetPortState {
	s.enabled = &enabled
	return s
}

// Send sends the SetPortState action to the ATE.
func (s *SetPortState) Send(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Setting port state on %s", s.ate)
	if err := ate.SetPortState(context.Background(), s.ate, s.portName, s.enabled); err != nil {
		t.Fatalf("Send(t) of %v: %v", s, err)
	}
}

// NewSetLACPState returns a new SetLACPState action.
func (a *Actions) NewSetLACPState() *SetLACPState {
	return &SetLACPState{ate: a.ate}
}

// SetLACPState is an action to set the state of LACP on an ATE.
type SetLACPState struct {
	ate      binding.ATE
	portName string
	enabled  *bool
}

func (s *SetLACPState) String() string {
	return fmt.Sprintf("SetLACPState%+v", *s)
}

// WithPort sets the port whose LACP state will be changed.
func (s *SetLACPState) WithPort(port *Port) *SetLACPState {
	s.portName = port.Name()
	return s
}

// WithEnabled sets the desired LACP state of the port.
func (s *SetLACPState) WithEnabled(enabled bool) *SetLACPState {
	s.enabled = &enabled
	return s
}

// Send sends the SetLACPState action to the ATE.
func (s *SetLACPState) Send(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Setting LACP state on %s", s.ate)
	if err := ate.SetLACPState(context.Background(), s.ate, s.portName, s.enabled); err != nil {
		t.Fatalf("Send(t) of %v: %v", s, err)
	}
}

// NewBGPPeerNotification returns a new BGPPeerNotification action.
// The action is configured by default with a code and sub code of 1.
func (a *Actions) NewBGPPeerNotification() *BGPPeerNotification {
	return &BGPPeerNotification{
		ate:     a.ate,
		code:    1,
		subCode: 1,
	}
}

// BGPPeerNotification is an action to send BGP notification and error codes.
type BGPPeerNotification struct {
	ate     binding.ATE
	peerIDs []uint32
	code    int
	subCode int
}

func (n *BGPPeerNotification) String() string {
	return fmt.Sprintf("BGPPeerNotification%+v", *n)
}

// WithPeers sets the BGP peers from which the notification is to be sent.
func (n *BGPPeerNotification) WithPeers(bgpPeers ...*ixnet.BGPPeer) *BGPPeerNotification {
	n.peerIDs = nil
	for _, bgpPeer := range bgpPeers {
		n.peerIDs = append(n.peerIDs, bgpPeer.ID())
	}
	return n
}

// WithCode sets the code of the notification to be sent from the BGP peers.
func (n *BGPPeerNotification) WithCode(code int) *BGPPeerNotification {
	n.code = code
	return n
}

// WithSubCode sets the sub code of the notification to be sent from the BGP peers.
func (n *BGPPeerNotification) WithSubCode(subCode int) *BGPPeerNotification {
	n.subCode = subCode
	return n
}

// Send executes the operation to send notification/error codes.
func (n *BGPPeerNotification) Send(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Sending BGP peer notification on %s", n.ate)
	if err := ate.SendBGPPeerNotification(context.Background(), n.ate, n.peerIDs, n.code, n.subCode); err != nil {
		t.Fatalf("Send(t) of %v: %v", n, err)
	}
}

// NewBGPGracefulRestart returns a new BGPGracefulRestart action.
func (a *Actions) NewBGPGracefulRestart() *BGPGracefulRestart {
	return &BGPGracefulRestart{
		ate: a.ate,
	}
}

// BGPGracefulRestart is an action to trigger a BGP graceful restart event.
type BGPGracefulRestart struct {
	ate     binding.ATE
	peerIDs []uint32
	delay   time.Duration
}

func (r *BGPGracefulRestart) String() string {
	return fmt.Sprintf("BGPGracefulRestart%+v", *r)
}

// WithPeers sets the BGP peers from which the graceful restart is to be sent.
func (r *BGPGracefulRestart) WithPeers(bgpPeers ...*ixnet.BGPPeer) *BGPGracefulRestart {
	r.peerIDs = nil
	for _, bgpPeer := range bgpPeers {
		r.peerIDs = append(r.peerIDs, bgpPeer.ID())
	}
	return r
}

// WithRestartTime sets the delay for the BGP sessions to remain down before restarting.
func (r *BGPGracefulRestart) WithRestartTime(delay time.Duration) *BGPGracefulRestart {
	r.delay = delay
	return r
}

// Send executes the operation to start the graceful restart event.
func (r *BGPGracefulRestart) Send(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Sending BGP graceful restart on %s", r.ate)
	if err := ate.SendBGPGracefulRestart(context.Background(), r.ate, r.peerIDs, r.delay); err != nil {
		t.Fatalf("Send(t) of %v: %v", r, err)
	}
}
