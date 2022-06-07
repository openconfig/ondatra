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

func (action *Actions) String() string {
	return fmt.Sprintf("{ate: %s}", action.ate)
}

// BGPPeerNotification is an actions API to send BGP notification/error codes
type BGPPeerNotification struct {
	ate                 binding.ATE
	notificationCode    int
	notificationSubCode int
	bgpPeerNames        []string
}

func (notification *BGPPeerNotification) String() string {
	return fmt.Sprintf("{notificationCode: %d, notificationSubCode: %d, bgpPeerNames: %s}", notification.notificationCode, notification.notificationSubCode, notification.bgpPeerNames)
}

// NewBGPPeerNotification returns an instance of API for sending BGP notification/error codes
func (action *Actions) NewBGPPeerNotification() *BGPPeerNotification {
	// setting the default values for the respective parameters
	return &BGPPeerNotification{
		ate:                 action.ate,
		notificationCode:    1,
		notificationSubCode: 1,
		bgpPeerNames:        []string{},
	}
}

// WithNotificationCode Sets the notification code that needs to be sent from BGP peer
func (notification *BGPPeerNotification) WithNotificationCode(code int) *BGPPeerNotification {
	notification.notificationCode = code
	return notification
}

// WithNotificationSubCode Sets the notification sub code that needs to be sent from BGP peer
func (notification *BGPPeerNotification) WithNotificationSubCode(subCode int) *BGPPeerNotification {
	notification.notificationSubCode = subCode
	return notification
}

// WithBGPPeers adds the BGP peers from which the notification/error codes is to be sent
func (notification *BGPPeerNotification) WithBGPPeers(bgpPeers ...*BGPPeer) *BGPPeerNotification {
	for _, bgpPeer := range bgpPeers {
		notification.bgpPeerNames = append(notification.bgpPeerNames, bgpPeer.FetchName())
	}
	return notification
}

// ClearBGPPeers clears the bgp peers currently assigned
func (notification *BGPPeerNotification) ClearBGPPeers() *BGPPeerNotification {
	notification.bgpPeerNames = notification.bgpPeerNames[:0]
	return notification
}

// Send executes the operation to send notification/error codes
func (notification *BGPPeerNotification) Send(t testing.TB) {
	t.Helper()

	if len(notification.bgpPeerNames) == 0 {
		t.Fatalf("Send(t) on %s: no bgp peer provided", notification)
	}

	if err := ate.SendBGPPeerNotification(context.Background(), notification.ate, notification.notificationCode, notification.notificationSubCode, notification.bgpPeerNames); err != nil {
		t.Fatalf("Send(t) failed with parameters %d, %d, %s on %s: %v", notification.notificationCode, notification.notificationSubCode, notification.bgpPeerNames, notification, err)
	}
}
