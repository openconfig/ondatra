// Package ate controls automated test equipment (ATE) for ONDATRA tests.
package ate

import (
	"golang.org/x/net/context"
	"net"

	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/internal/usererr"

	opb "github.com/openconfig/ondatra/proto"
)

// StartTraffic starts traffic flows on an ATE.
func StartTraffic(ctx context.Context, ate *reservation.ATE, flows []*opb.Flow) error {
	if err := validateFlows(ate, flows); err != nil {
		return err
	}
	return binding.Get().StartTraffic(ate, flows)
}

// UpdateTraffic updates traffic flows an an ATE.
func UpdateTraffic(ctx context.Context, ate *reservation.ATE, flows []*opb.Flow) error {
	if err := validateFlows(ate, flows); err != nil {
		return err
	}
	return binding.Get().UpdateTraffic(ate, flows)
}

// StopTraffic stops traffic flows on an ATE>
func StopTraffic(ctx context.Context, ate *reservation.ATE) error {
	return binding.Get().StopTraffic(ate)
}

// PushTopology pushes a topology to an ATE.
func PushTopology(ctx context.Context, ate *reservation.ATE, top *opb.Topology) error {
	if err := validateInterfaces(ate, top.GetInterfaces()); err != nil {
		return err
	}
	return binding.Get().PushTopology(ate, top)
}

// UpdateTopology updates a topology on an ATE.
func UpdateTopology(ctx context.Context, ate *reservation.ATE, top *opb.Topology, bgpPeerStateOnly bool) error {
	if err := validateInterfaces(ate, top.GetInterfaces()); err != nil {
		return err
	}
	// TODO: Remove this branching once new Ixia config binding is used.
	if bgpPeerStateOnly {
		return binding.Get().UpdateBGPPeerStates(ate, top.GetInterfaces())
	}
	return binding.Get().UpdateTopology(ate, top)
}

// StartProtocols starts control plane protocols on an ATE.
func StartProtocols(ctx context.Context, ate *reservation.ATE) error {
	return binding.Get().StartProtocols(ate)
}

// StopProtocols stops control protocols on an ATE.
func StopProtocols(ctx context.Context, ate *reservation.ATE) error {
	return binding.Get().StopProtocols(ate)
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
