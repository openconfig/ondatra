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
)

// Actions is an ATE Action API.
type Actions struct {
	ate binding.ATE
}

func (a *Actions) String() string {
	return fmt.Sprintf("Actions%+v", *a)
}

// BGPPeerNotification is an actions API to send BGP notification/error codes.
type BGPPeerNotification struct {
	ate       binding.ATE
	code      int
	subCode   int
	peerNames []string
}

func (n *BGPPeerNotification) String() string {
	return fmt.Sprintf("BGPPeerNotification%+v", *n)
}

// NewBGPPeerNotification returns an instance of API for sending BGP notification/error codes.
func (a *Actions) NewBGPPeerNotification() *BGPPeerNotification {
	// setting the default values for the respective parameters
	return &BGPPeerNotification{
		ate:       a.ate,
		code:      1,
		subCode:   1,
		peerNames: []string{},
	}
}

// WithCode sets the notification code that needs to be sent from BGP peer.
func (n *BGPPeerNotification) WithCode(code int) *BGPPeerNotification {
	n.code = code
	return n
}

// WithSubCode sets the notification sub code that needs to be sent from BGP peer.
func (n *BGPPeerNotification) WithSubCode(subCode int) *BGPPeerNotification {
	n.subCode = subCode
	return n
}

// WithBGPPeers adds the BGP peers from which the notification/error codes is to be sent.
func (n *BGPPeerNotification) WithBGPPeers(bgpPeers ...*BGPPeer) *BGPPeerNotification {
	for _, bgpPeer := range bgpPeers {
		n.peerNames = append(n.peerNames, bgpPeer.Name())
	}
	return n
}

// ClearBGPPeers clears the bgp peers currently assigned.
func (n *BGPPeerNotification) ClearBGPPeers() *BGPPeerNotification {
	n.peerNames = n.peerNames[:0]
	return n
}

// Send executes the operation to send notification/error codes.
func (n *BGPPeerNotification) Send(t testing.TB) {
	t.Helper()
	if err := ate.SendBGPPeerNotification(context.Background(), n.ate, n.code, n.subCode, n.peerNames); err != nil {
		t.Fatalf("Send(t) failed on %s: %v", n, err)
	}
}
