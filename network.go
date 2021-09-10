package ondatra

import (
	opb "github.com/openconfig/ondatra/proto"
)

// Network is a group of simulated device interfaces.
type Network struct {
	pb *opb.Network
}

// Implement the Endpoint marker interface.
func (*Network) isEndpoint() {}

// Ethernet creates an Ethernet config for the network or returns the existing config.
// The default count of MAC addresses in the network is 1.
func (n *Network) Ethernet() *NetworkEthernet {
	if n.pb.Eth == nil {
		n.pb.Eth = &opb.NetworkEth{Count: 1}
	}
	return &NetworkEthernet{n.pb.Eth}
}

// NetworkEthernet is the Ethernet config of a network.
type NetworkEthernet struct {
	pb *opb.NetworkEth
}

// WithMACAddress sets the first MAC address.
func (e *NetworkEthernet) WithMACAddress(address string) *NetworkEthernet {
	e.pb.MacAddress = address
	return e
}

// WithCount sets the count of MAC addressses.
func (e *NetworkEthernet) WithCount(count uint32) *NetworkEthernet {
	e.pb.Count = count
	return e
}

// WithVLANID enables a VLAN with a 12-bit specified ID on the interface.
func (e *NetworkEthernet) WithVLANID(id uint16) *NetworkEthernet {
	e.pb.VlanId = uint32(id)
	return e
}

// IPv4 creates an IPv4 config for the network or returns the existing config.
// The default count of IP addresses in the network is 1.
func (n *Network) IPv4() *NetworkIP {
	if n.pb.Ipv4 == nil {
		n.pb.Ipv4 = &opb.NetworkIp{Count: 1}
	}
	return &NetworkIP{n.pb.Ipv4}
}

// IPv6 creates an IPv6 config for the network or returns the existing config.
// The default count of IP addresses in the network is 1.
func (n *Network) IPv6() *NetworkIP {
	if n.pb.Ipv6 == nil {
		n.pb.Ipv6 = &opb.NetworkIp{Count: 1}
	}
	return &NetworkIP{n.pb.Ipv6}
}

// NetworkIP is the Network config of a network.
type NetworkIP struct {
	pb *opb.NetworkIp
}

// WithAddress sets the first IP address in CIDR notation.
func (i *NetworkIP) WithAddress(address string) *NetworkIP {
	i.pb.AddressCidr = address
	return i
}

// WithCount sets the count of IP addressses.
func (i *NetworkIP) WithCount(count uint32) *NetworkIP {
	i.pb.Count = count
	return i
}

// Address returns the currently configured address.
func (i *NetworkIP) Address() string {
	return i.pb.GetAddressCidr()
}

// Count returns the currently configured address count.
func (i *NetworkIP) Count() uint32 {
	return i.pb.GetCount()
}

// ISIS creates an ISIS configuration for the network or returns the existing config.
// The default config params are Route Origin: Internal and Metric: 10
func (n *Network) ISIS() *IPReachabilityConfig {
	if n.pb.Isis == nil {
		n.pb.Isis = &opb.IPReachability{Metric: 10, RouteOrigin: opb.IPReachability_INTERNAL}
	}
	return &IPReachabilityConfig{pb: n.pb.Isis}
}

// BGP creates a BGP config for the network or returns the existing config.  By
// default, the network will have the following configuration:
//   Active: true
//   Origin: IGP
//   ASN set mode: AS-SEQ
func (n *Network) BGP() *BGPAttributes {
	if n.pb.BgpAttributes == nil {
		n.pb.BgpAttributes = &opb.BgpAttributes{
			Active:     true,
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
		}
	}
	return &BGPAttributes{pb: n.pb.BgpAttributes}
}

// ImportedBGPRoutes creates a BGP route import configuration for the network or
// returns the existing config.
func (n *Network) ImportedBGPRoutes() *ImportedBGPRoutes {
	if n.pb.ImportedBgpRoutes == nil {
		n.pb.ImportedBgpRoutes = &opb.Network_ImportedBgpRoutes{}
	}
	return &ImportedBGPRoutes{pb: n.pb.ImportedBgpRoutes}
}

// ImportedBGPRoutes is the config for importing existing route tables into a
// network.
type ImportedBGPRoutes struct {
	pb *opb.Network_ImportedBgpRoutes
}

// WithRouteTableFormatCisco sets Cisco as the expected route table format of
// the specified routes.
func (i *ImportedBGPRoutes) WithRouteTableFormatCisco() *ImportedBGPRoutes {
	i.pb.RouteTableFormat = opb.Network_ImportedBgpRoutes_ROUTE_TABLE_FORMAT_CISCO
	return i
}

// WithRouteTableFormatJuniper sets Juniper as the expected route table format of
// the specified routes.
func (i *ImportedBGPRoutes) WithRouteTableFormatJuniper() *ImportedBGPRoutes {
	i.pb.RouteTableFormat = opb.Network_ImportedBgpRoutes_ROUTE_TABLE_FORMAT_JUNIPER
	return i
}

// WithIPv4RoutesPath sets the path for a file containing a BGP IPv4 route
// table.
func (i *ImportedBGPRoutes) WithIPv4RoutesPath(path string) *ImportedBGPRoutes {
	i.pb.Ipv4RoutesPath = path
	return i
}

// WithIPv6RoutesPath sets the path for a file containing a BGP IPv6 route
// table.
func (i *ImportedBGPRoutes) WithIPv6RoutesPath(path string) *ImportedBGPRoutes {
	i.pb.Ipv6RoutesPath = path
	return i
}
