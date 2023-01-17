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

package ondatra

import (
	"github.com/openconfig/ondatra/ixnet"

	opb "github.com/openconfig/ondatra/proto"
)

// Interface represents an ATE interface configuration.
type Interface struct {
	pb *opb.InterfaceConfig
}

// EndpointPB returns the Interface config as an endpoint proto message.
func (i *Interface) EndpointPB() *opb.Flow_Endpoint {
	return &opb.Flow_Endpoint{InterfaceName: i.pb.GetName()}
}

// IsLACPEnabled returns whether the interface has LACP enabled.
func (i *Interface) IsLACPEnabled() bool {
	return i.pb.EnableLacp
}

// AddNetwork adds and returns a network with the specified name.
func (i *Interface) AddNetwork(name string) *ixnet.Network {
	npb := &opb.Network{Name: name, InterfaceName: i.pb.Name}
	i.pb.Networks = append(i.pb.Networks, npb)
	return ixnet.NewNetwork(npb)
}

// ClearNetworks clears networks from the interface.
func (i *Interface) ClearNetworks() *Interface {
	i.pb.Networks = nil
	return i
}

// Networks returns a map of network names to networks.
func (i *Interface) Networks() map[string]*ixnet.Network {
	nets := make(map[string]*ixnet.Network)
	for _, pb := range i.pb.GetNetworks() {
		nets[pb.GetName()] = ixnet.NewNetwork(pb)
	}
	return nets
}

// WithPort specifies that the interface will be configured on the given port.
func (i *Interface) WithPort(p *Port) *Interface {
	i.pb.Link = &opb.InterfaceConfig_Port{p.Name()}
	return i
}

// WithLAG specifies that the interface will be configured on the given LAG.
// TODO(team): Add Ondatra test for this feature.
func (i *Interface) WithLAG(l *LAG) *Interface {
	i.pb.Link = &opb.InterfaceConfig_Lag{l.pb.GetName()}
	return i
}

// WithLACPEnabled specifies whether the interface has LACP enabled.
func (i *Interface) WithLACPEnabled(enabled bool) *Interface {
	i.pb.EnableLacp = enabled
	return i
}

// Ethernet returns the existing Ethernet config.
func (i *Interface) Ethernet() *ixnet.Ethernet {
	return ixnet.NewEthernet(i.pb.Ethernet)
}

// IPv4 creates an IPv4 config for the interface or returns the existing config.
func (i *Interface) IPv4() *ixnet.IP {
	if i.pb.Ipv4 == nil {
		i.pb.Ipv4 = &opb.IpConfig{}
	}
	return ixnet.NewIP(i.pb.Ipv4)
}

// IPv6 creates an IPv6 config for the interface or returns the existing config.
func (i *Interface) IPv6() *ixnet.IP {
	if i.pb.Ipv6 == nil {
		i.pb.Ipv6 = &opb.IpConfig{}
	}
	return ixnet.NewIP(i.pb.Ipv6)
}

// WithIPv4Loopback configures the IPv4 loopback address for the interface.
func (i *Interface) WithIPv4Loopback(ipv4LoopbackCIDR string) *Interface {
	i.pb.Ipv4LoopbackCidr = ipv4LoopbackCIDR
	return i
}

// WithIPv6Loopback configures the IPv6 loopback address for the interface.
func (i *Interface) WithIPv6Loopback(ipv6LoopbackCIDR string) *Interface {
	i.pb.Ipv6LoopbackCidr = ipv6LoopbackCIDR
	return i
}

// ISIS creates an ISIS config for the interface or returns the existing config.
// The default config paramas are:
// Area Id: 490001
// Router Id: 0.0.0.0
// Hello Interval: 10 seconds
// Dead Interval: 30 seconds
func (i *Interface) ISIS() *ixnet.ISIS {
	if i.pb.Isis == nil {
		i.pb.Isis = &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		}
	}
	return ixnet.NewISIS(i.pb.Isis)
}

// BGP creates a BGP config for the interface or returns the existing config.
func (i *Interface) BGP() *ixnet.BGP {
	if i.pb.Bgp == nil {
		i.pb.Bgp = &opb.BgpConfig{}
	}
	return ixnet.NewBGP(i.pb.Bgp)
}

// AddRSVP adds an RSVP config to the interface.
func (i *Interface) AddRSVP(name string) *ixnet.RSVP {
	rpb := &opb.RsvpConfig{Name: name, InterfaceName: i.pb.Name}
	i.pb.Rsvps = append(i.pb.Rsvps, rpb)
	return ixnet.NewRSVP(rpb)
}

// DHCPV6Client creates a DHCP v6 Client or returns the existing config.
func (i *Interface) DHCPV6Client() *ixnet.DHCPV6Client {
	if i.pb.Dhcpv6Client == nil {
		i.pb.Dhcpv6Client = &opb.DhcpV6Client{}
	}
	return ixnet.NewDHCPV6Client(i.pb.Dhcpv6Client)
}

// DHCPV6Server creates a DHCP v6 Server or returns the existing config.
func (i *Interface) DHCPV6Server() *ixnet.DHCPV6Server {
	if i.pb.Dhcpv6Server == nil {
		i.pb.Dhcpv6Server = &opb.DhcpV6Server{}
	}
	return ixnet.NewDHCPV6Server(i.pb.Dhcpv6Server)
}
