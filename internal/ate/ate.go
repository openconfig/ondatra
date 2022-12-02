// Copyright 2019 Google LLC
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

// Package ate controls automated test equipment (ATE) for ONDATRA tests.
package ate

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/rawapis"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	opb "github.com/openconfig/ondatra/proto"
)

var (
	mu    sync.Mutex
	ixias = make(map[binding.ATE]*ixATE)
)

// Topology is an ATE topology.
type Topology struct {
	Interfaces []*opb.InterfaceConfig
	LAGs       []*opb.Lag
}

func ixiaForATE(ctx context.Context, ate binding.ATE) (*ixATE, error) {
	mu.Lock()
	defer mu.Unlock()
	ix, ok := ixias[ate]
	if !ok {
		ixnet, err := rawapis.FetchIxNetwork(ctx, ate)
		if err != nil {
			return nil, err
		}
		ix, err = newIxATE(ctx, ate.Name(), ixnet)
		if err != nil {
			return nil, err
		}
		ixias[ate] = ix
	}
	return ix, nil
}

// PushTopology pushes a topology to an ATE.
func PushTopology(ctx context.Context, ate binding.ATE, top *Topology) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	if err := ix.PushTopology(ctx, top); err != nil {
		return err
	}
	ix.FlushStats()
	return nil
}

// UpdateTopology updates a topology on an ATE.
func UpdateTopology(ctx context.Context, ate binding.ATE, top *Topology, bgpPeerStateOnly bool) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	// TODO(team): Remove this branching once new Ixia config binding is used.
	if bgpPeerStateOnly {
		err = ix.UpdateBGPPeerStates(ctx, top.Interfaces)
	} else {
		err = ix.UpdateTopology(ctx, top)
	}
	if err != nil {
		return err
	}
	ix.FlushStats()
	return nil
}

// UpdateNetworks updates network groups in a topology on an ATE on the fly.
func UpdateNetworks(ctx context.Context, ate binding.ATE, top *Topology) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	return ix.UpdateNetworkGroups(ctx, top.Interfaces)
}

// StartProtocols starts control plane protocols on an ATE.
func StartProtocols(ctx context.Context, ate binding.ATE) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	if err := ix.StartProtocols(ctx); err != nil {
		return fmt.Errorf("failed to start protocols: %w", err)
	}
	ix.FlushStats()
	return nil
}

// StopProtocols stops control protocols on an ATE.
func StopProtocols(ctx context.Context, ate binding.ATE) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	if err := ix.StopProtocols(ctx); err != nil {
		return err
	}
	ix.FlushStats()
	return nil
}

// StartTraffic starts traffic flows on an ATE.
func StartTraffic(ctx context.Context, ate binding.ATE, flows []*opb.Flow) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	if err := ix.StartTraffic(ctx, flows); err != nil {
		return err
	}
	ix.FlushStats()
	return nil
}

// UpdateTraffic updates traffic flows an an ATE.
func UpdateTraffic(ctx context.Context, ate binding.ATE, flows []*opb.Flow) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	if err := ix.UpdateTraffic(ctx, flows); err != nil {
		return err
	}
	ix.FlushStats()
	return nil
}

// StopTraffic stops traffic flows on an ATE.
func StopTraffic(ctx context.Context, ate binding.ATE) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	if err := ix.StopAllTraffic(ctx); err != nil {
		return err
	}
	ix.FlushStats()
	return nil
}

// FetchGNMI returns the GNMI client for the Ixia.
func FetchGNMI(ctx context.Context, ate binding.ATE) (gpb.GNMIClient, error) {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return nil, err
	}
	return ix.FetchGNMI(ctx)
}

// SetPortState sets the state of a specified interface on the ATE.
func SetPortState(ctx context.Context, ate binding.ATE, port string, enabled *bool) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	return ix.SetPortState(ctx, port, enabled)
}

// SetLACPState sets the LACP state of a specified interface on the ATE.
func SetLACPState(ctx context.Context, ate binding.ATE, port string, enabled *bool) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	return ix.SetLACPState(ctx, port, enabled)
}

// SendBGPPeerNotification sends a notification from BGP peers.
func SendBGPPeerNotification(ctx context.Context, ate binding.ATE, peerIDs []uint32, code int, subCode int) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	if err := ix.SendBGPPeerNotification(ctx, peerIDs, code, subCode); err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}
	return nil
}

// SendBGPGracefulRestart sends a BGP graceful restart event to BGP peers.
func SendBGPGracefulRestart(ctx context.Context, ate binding.ATE, peerIDs []uint32, delay time.Duration) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	if err := ix.SendBGPGracefulRestart(ctx, peerIDs, delay); err != nil {
		return fmt.Errorf("failed to send graceful restart: %w", err)
	}
	return nil
}
