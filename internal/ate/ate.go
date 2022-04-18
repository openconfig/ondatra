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
	"golang.org/x/net/context"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/testbed"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	opb "github.com/openconfig/ondatra/proto"
)

var (
	mu    sync.Mutex
	ixias = make(map[*binding.ATE]*ixATE)
)

// Topology is an ATE topology.
type Topology struct {
	Interfaces []*opb.InterfaceConfig
	LAGs       []*opb.Lag
}

func ixiaForATE(ctx context.Context, ate *binding.ATE) (*ixATE, error) {
	mu.Lock()
	defer mu.Unlock()
	ix, ok := ixias[ate]
	if !ok {
		ixnet, err := testbed.Bind().DialIxNetwork(ctx, ate)
		if err != nil {
			return nil, err
		}
		ix, err = newIxATE(ctx, ate.Name, ixnet)
		if err != nil {
			return nil, err
		}
		ixias[ate] = ix
	}
	return ix, nil
}

// PushTopology pushes a topology to an ATE.
func PushTopology(ctx context.Context, ate *binding.ATE, top *Topology) error {
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
func UpdateTopology(ctx context.Context, ate *binding.ATE, top *Topology, bgpPeerStateOnly bool) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	// TODO: Remove this branching once new Ixia config binding is used.
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

// StartProtocols starts control plane protocols on an ATE.
func StartProtocols(ctx context.Context, ate *binding.ATE) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	if err := ix.StartProtocols(ctx); err != nil {
		return errors.Wrap(err, "failed to start protocols")
	}
	ix.FlushStats()
	return nil
}

// StopProtocols stops control protocols on an ATE.
func StopProtocols(ctx context.Context, ate *binding.ATE) error {
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
func StartTraffic(ctx context.Context, ate *binding.ATE, flows []*opb.Flow) error {
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
func UpdateTraffic(ctx context.Context, ate *binding.ATE, flows []*opb.Flow) error {
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
func StopTraffic(ctx context.Context, ate *binding.ATE) error {
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

// SetInterfaceState sets the state of a specified interface on the ATE.
func SetInterfaceState(ctx context.Context, ate *binding.ATE, intf string, enabled bool) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	return ix.SetPortState(ctx, intf, enabled)
}

// DialGNMI constructs and returns a GNMI client for the Ixia.
func DialGNMI(ctx context.Context, ate *binding.ATE, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return nil, err
	}
	return ix.DialGNMI(ctx, opts...)
}
