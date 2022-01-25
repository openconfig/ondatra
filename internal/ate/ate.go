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
	"net"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/internal/usererr"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	opb "github.com/openconfig/ondatra/proto"
)

var (
	mu    sync.Mutex
	ixias = make(map[*reservation.ATE]*ixATE)
)

func ixiaForATE(ctx context.Context, ate *reservation.ATE) (*ixATE, error) {
	mu.Lock()
	defer mu.Unlock()
	ix, ok := ixias[ate]
	if !ok {
		ixnet, err := binding.Get().DialIxNetwork(ctx, ate)
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
func PushTopology(ctx context.Context, ate *reservation.ATE, top *opb.Topology) error {
	if err := validateInterfaces(ate, top.GetInterfaces()); err != nil {
		return err
	}
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
func UpdateTopology(ctx context.Context, ate *reservation.ATE, top *opb.Topology, bgpPeerStateOnly bool) error {
	if err := validateInterfaces(ate, top.GetInterfaces()); err != nil {
		return err
	}
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	// TODO: Remove this branching once new Ixia config binding is used.
	if bgpPeerStateOnly {
		err = ix.UpdateBGPPeerStates(ctx, top.GetInterfaces())
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
func StartProtocols(ctx context.Context, ate *reservation.ATE) error {
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
func StopProtocols(ctx context.Context, ate *reservation.ATE) error {
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
func StartTraffic(ctx context.Context, ate *reservation.ATE, flows []*opb.Flow) error {
	if err := validateFlows(ate, flows); err != nil {
		return err
	}
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
func UpdateTraffic(ctx context.Context, ate *reservation.ATE, flows []*opb.Flow) error {
	if err := validateFlows(ate, flows); err != nil {
		return err
	}
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
func StopTraffic(ctx context.Context, ate *reservation.ATE) error {
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
func SetInterfaceState(ctx context.Context, ate *reservation.ATE, intf string, enabled bool) error {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return err
	}
	return ix.SetPortState(ctx, intf, enabled)
}

// DialGNMI constructs and returns a GNMI client for the Ixia.
func DialGNMI(ctx context.Context, ate *reservation.ATE, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	ix, err := ixiaForATE(ctx, ate)
	if err != nil {
		return nil, err
	}
	return ix.DialGNMI(ctx, opts...)
}

func validateFlows(ate *reservation.ATE, fs []*opb.Flow) error {
	for _, f := range fs {
		if len(f.GetSrcEndpoints()) == 0 {
			return usererr.New("flow has no src endpointd")
		}
		if len(f.GetDstEndpoints()) == 0 {
			return usererr.New("flow has no dst endpoints")
		}
	}
	return nil
}

func validateInterfaces(ate *reservation.ATE, ifs []*opb.InterfaceConfig) error {
	if len(ifs) == 0 {
		return usererr.New("zero interfaces to configure, need at least one")
	}
	intfs := make(map[string]bool)

	for _, i := range ifs {
		if i.GetPort() == "" && i.GetLag() == "" {
			return usererr.New("interface has no port or lag specified: %v", i)
		}
		if i.GetLag() != "" && i.GetEnableLacp() {
			return usererr.New("interface should not specify both a LAG and that LACP is enabled: %v", i)
		}
		if intfs[i.GetName()] {
			return usererr.New("duplicate interface name: %s", i.GetName())
		}
		intfs[i.GetName()] = true
		nets := make(map[string]bool)
		for _, n := range i.GetNetworks() {
			if nets[n.GetName()] {
				return usererr.New("duplicate network name: %s", n.GetName())
			}
			nets[n.GetName()] = true
		}
		if err := validateIP(i.GetIpv4(), "ipv4 on "+i.GetName()); err != nil {
			return err
		}
		if err := validateIP(i.GetIpv6(), "ipv6 on "+i.GetName()); err != nil {
			return err
		}
	}
	return nil
}

func validateIP(ipc *opb.IpConfig, desc string) error {
	if ipc == nil {
		return nil
	}
	addr := ipc.GetAddressCidr()
	gway := ipc.GetDefaultGateway()
	_, an, err := net.ParseCIDR(addr)
	if err != nil {
		return usererr.New("%s address is not valid CIDR notation: %s", desc, addr)
	}
	gi := net.ParseIP(gway)
	if gi == nil {
		return usererr.New("%s default gateway is not valid IP notation: %s", desc, gway)
	}
	if !an.Contains(gi) {
		return usererr.New("%s default gateway is not in CIDR range %s: %s", desc, addr, gway)
	}
	return nil
}
