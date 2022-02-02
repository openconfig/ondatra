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
	opb "github.com/openconfig/ondatra/proto"
)

// Interface represents an ATE interface configuration.
type Interface struct {
	pb *opb.InterfaceConfig
}

// Implement the Endpoint marker interface.
func (*Interface) isEndpoint() {}

// IsLACPEnabled returns whether the interface has LACP enabled.
func (i *Interface) IsLACPEnabled() bool {
	return i.pb.EnableLacp
}

// AddNetwork adds and returns a network with the specified name.
func (i *Interface) AddNetwork(name string) *Network {
	npb := &opb.Network{Name: name, InterfaceName: i.pb.Name}
	i.pb.Networks = append(i.pb.Networks, npb)
	return &Network{pb: npb}
}

// ClearNetworks clears networks from the interface.
func (i *Interface) ClearNetworks() *Interface {
	i.pb.Networks = nil
	return i
}

// Networks returns a map of network names to networks.
func (i *Interface) Networks() map[string]*Network {
	nets := make(map[string]*Network)
	for _, pb := range i.pb.GetNetworks() {
		nets[pb.GetName()] = &Network{pb}
	}
	return nets
}

// WithPort specifies that the interface will be configured on the given port.
func (i *Interface) WithPort(p *Port) *Interface {
	i.pb.Link = &opb.InterfaceConfig_Port{p.Name()}
	return i
}

// WithLAG specifies that the interface will be configured on the given LAG.
// TODO: Add Ondatra test for this feature.
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
func (i *Interface) Ethernet() *Ethernet {
	return &Ethernet{pb: i.pb.Ethernet}
}

// IPv4 creates an IPv4 config for the interface or returns the existing config.
func (i *Interface) IPv4() *IP {
	if i.pb.Ipv4 == nil {
		i.pb.Ipv4 = &opb.IpConfig{}
	}
	return &IP{pb: i.pb.Ipv4}
}

// IPv6 creates an IPv6 config for the interface or returns the existing config.
func (i *Interface) IPv6() *IP {
	if i.pb.Ipv6 == nil {
		i.pb.Ipv6 = &opb.IpConfig{}
	}
	return &IP{pb: i.pb.Ipv6}
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
func (i *Interface) ISIS() *ISIS {
	if i.pb.Isis == nil {
		i.pb.Isis = &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		}
	}
	return &ISIS{pb: i.pb.Isis}
}

// BGP creates a BGP config for the interface or returns the existing config.
func (i *Interface) BGP() *BGP {
	if i.pb.Bgp == nil {
		i.pb.Bgp = &opb.BgpConfig{}
	}
	return &BGP{pb: i.pb.Bgp}
}

// AddRSVP adds an RSVP config to the interface.
func (i *Interface) AddRSVP(name string) *RSVP {
	rpb := &opb.RsvpConfig{Name: name, InterfaceName: i.pb.Name}
	i.pb.Rsvp = append(i.pb.Rsvp, rpb)
	return &RSVP{pb: rpb}
}
