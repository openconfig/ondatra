package ondatra

import (
	opb "github.com/openconfig/ondatra/proto"
)

// Header is a packet header.
type Header interface {
	asPB() *opb.Header
}

// NewEthernetHeader returns a new Ethernet header.
// The header is initialized with none of its properties specified.
// Unless specified, a best effort will be made to infer the VLAN ID and src and dst addresses from the topology.
func NewEthernetHeader() *EthernetHeader {
	return &EthernetHeader{&opb.EthernetHeader{}}
}

// EthernetHeader is an Ethernet packet header.
type EthernetHeader struct {
	pb *opb.EthernetHeader
}

// WithSrcAddress sets the source MAC address of the Ethernet header to the specified value.
// To generate a range of source addresses, use SrcAddressRange() instead.
func (h *EthernetHeader) WithSrcAddress(addr string) *EthernetHeader {
	h.pb.SrcAddr = addrRangeSingle(addr)
	return h
}

// SrcAddressRange returns the range of source addresses of the Ethernet header for further configuration.
// By default, the range will be nonrandom values in the interval ["00:00:00:00:00:01", "FF:FF:FF:FF:FF:FF").
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *EthernetHeader) SrcAddressRange() *AddressRange {
	if h.pb.SrcAddr == nil {
		h.pb.SrcAddr = newMACAddrRange()
	}
	return &AddressRange{AddressIncRange{pb: h.pb.SrcAddr}}
}

// WithDstAddress sets the destination MAC address of the Ethernet header to the specified value.
// To generate a range of destination addresses, use DstAddressRange() instead.
func (h *EthernetHeader) WithDstAddress(addr string) *EthernetHeader {
	h.pb.DstAddr = addrRangeSingle(addr)
	return h
}

// DstAddressRange returns the range of destination addresses of the Ethernet header for further configuration.
// By default, the range will be nonrandom values in the interval ["00:00:00:00:00:01", "FF:FF:FF:FF:FF:FF").
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *EthernetHeader) DstAddressRange() *AddressRange {
	if h.pb.DstAddr == nil {
		h.pb.DstAddr = newMACAddrRange()
	}
	return &AddressRange{AddressIncRange{pb: h.pb.DstAddr}}
}

// WithVLANID sets the 12-bit VLAN ID of the Ethernet header to the specified value.
func (h *EthernetHeader) WithVLANID(vid uint16) *EthernetHeader {
	h.pb.VlanId = uint32(vid)
	return h
}

// WithBadCRC set whether the Ethernet header has an incorrect CRC in the frame
// check sequence.
func (h *EthernetHeader) WithBadCRC(bad bool) *EthernetHeader {
	h.pb.BadCrc = true
	return h
}

func (h *EthernetHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Eth{h.pb}}
}

// NewGREHeader returns a new GRE header.
// The header is initialized with none of its properties specified.
func NewGREHeader() *GREHeader {
	return &GREHeader{&opb.GreHeader{}}
}

// GREHeader is a Generic Route Encapsulation packet header.
type GREHeader struct {
	pb *opb.GreHeader
}

// WithKey sets the key of the GRE header.
func (h *GREHeader) WithKey(key uint32) *GREHeader {
	h.pb.Key = key
	return h
}

// WithSequenceNumber sets sequence number of the GRE header.
func (h *GREHeader) WithSequenceNumber(seq uint32) *GREHeader {
	h.pb.Seq = seq
	return h
}

func (h *GREHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Gre{h.pb}}
}

// NewIPv4Header returns a new IPv4 header.
// The header is initialized with a hop limit of 64, DSCP of 0 (best effort), and ECN of 0 (not ECN capable).
// Unless specified, a best effort will be made to infer the src and dst addresses from the topology.
func NewIPv4Header() *IPv4Header {
	return &IPv4Header{&opb.Ipv4Header{Ttl: 64, Dscp: 0, Ecn: 0}}
}

// IPv4Header is an IPv4 packet header.
type IPv4Header struct {
	pb *opb.Ipv4Header
}

// WithSrcAddress sets the source IP address of the IPv4 header to the specified value.
// To generate a range of source addresses, use SrcAddressRange() instead.
func (h *IPv4Header) WithSrcAddress(addr string) *IPv4Header {
	h.pb.SrcAddr = addrRangeSingle(addr)
	return h
}

// SrcAddressRange returns the range of source addresses of the IPv4 header for further configuration.
// By default, the range will be nonrandom values in the interval ["0.0.0.1", "255.255.255.255").
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *IPv4Header) SrcAddressRange() *AddressRange {
	if h.pb.SrcAddr == nil {
		h.pb.SrcAddr = newIPv4AddrRange()
	}
	return &AddressRange{AddressIncRange{pb: h.pb.SrcAddr}}
}

// WithDstAddress sets the destination IP addresses of the IPv4 header to the specified value.
// To generate a range of destination addresses, use DstAddressRange() instead.
func (h *IPv4Header) WithDstAddress(addr string) *IPv4Header {
	h.pb.DstAddr = addrRangeSingle(addr)
	return h
}

// DstAddressRange returns the range of destination addresses of the IPv4 header for further configuration.
// By default, the range will be nonrandom values in the interval ["0.0.0.1", "255.255.255.255").
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *IPv4Header) DstAddressRange() *AddressRange {
	if h.pb.DstAddr == nil {
		h.pb.DstAddr = newIPv4AddrRange()
	}
	return &AddressRange{AddressIncRange{pb: h.pb.DstAddr}}
}

// WithDontFragment sets the "don't fragment" bit of the IPv4 header.
func (h *IPv4Header) WithDontFragment(dontFragment bool) *IPv4Header {
	h.pb.DontFragment = dontFragment
	return h
}

// WithTTL sets the TTL of the IPv4 header.
func (h *IPv4Header) WithTTL(ttl uint8) *IPv4Header {
	h.pb.Ttl = uint32(ttl)
	return h
}

// WithDSCP sets the DSCP value of the IPv4 header.
func (h *IPv4Header) WithDSCP(dscp uint8) *IPv4Header {
	h.pb.Dscp = uint32(dscp)
	return h
}

// WithECN sets the ECN value of the IPv4 header.
func (h *IPv4Header) WithECN(ecn uint8) *IPv4Header {
	h.pb.Ecn = uint32(ecn)
	return h
}

func (h *IPv4Header) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Ipv4{h.pb}}
}

// NewIPv6Header returns a new IPv6 header.
// The header is initialized with a hop limit of 64, DSCP of 0 (best effort), and ECN of 0 (not ECN capable).
// Unless specified, a best effort will be made to infer the src and dst addresses from the topology.
func NewIPv6Header() *IPv6Header {
	return &IPv6Header{&opb.Ipv6Header{HopLimit: 64, Dscp: 0, Ecn: 0}}
}

// IPv6Header is an IPv6 packet header.
type IPv6Header struct {
	pb *opb.Ipv6Header
}

// WithSrcAddress sets the source IP addresses of the IPv6 header to the specified value.
// To generate a range of source addresses, use SrcAddressRange() instead.
func (h *IPv6Header) WithSrcAddress(addr string) *IPv6Header {
	h.pb.SrcAddr = addrRangeSingle(addr)
	return h
}

// SrcAddressRange returns the range of source addresses of the IPv6 header for further configuration.
// By default, the range will be nonrandom values in the interval ["::1", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"].
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *IPv6Header) SrcAddressRange() *AddressRange {
	if h.pb.SrcAddr == nil {
		h.pb.SrcAddr = newIPv6AddrRange()
	}
	return &AddressRange{AddressIncRange{pb: h.pb.SrcAddr}}
}

// WithDstAddress sets the destination IP addresses of the IPv6 header to the specified value.
// To generate a range of destination addresses, use DstAddressRange() instead.
func (h *IPv6Header) WithDstAddress(addr string) *IPv6Header {
	h.pb.DstAddr = addrRangeSingle(addr)
	return h
}

// DstAddressRange returns the range of destination addresses of the IPv6 header for further configuration.
// By default, the range will be nonrandom values in the interval ["::1", "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"].
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *IPv6Header) DstAddressRange() *AddressRange {
	if h.pb.DstAddr == nil {
		h.pb.DstAddr = newIPv6AddrRange()
	}
	return &AddressRange{AddressIncRange{pb: h.pb.DstAddr}}
}

// WithFlowLabel sets the flow label of the IPv6 header.
func (h *IPv6Header) WithFlowLabel(flowLabel uint32) *IPv6Header {
	h.pb.FlowLabel = intRangeSingle(flowLabel)
	return h
}

// FlowLabelRange sets the flow label of the IPv6 header to a range of values and returns the range.
// By default, the range will be nonrandom values in the interval [0, 2^20).
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *IPv6Header) FlowLabelRange() *UIntRange {
	if h.pb.FlowLabel == nil {
		h.pb.FlowLabel = newFlowLabelRange()
	}
	return &UIntRange{pb: h.pb.FlowLabel}
}

// WithHopLimit sets the hop limit of the IPv6 header.
func (h *IPv6Header) WithHopLimit(hopLimit uint8) *IPv6Header {
	h.pb.HopLimit = uint32(hopLimit)
	return h
}

// WithDSCP sets the DSCP value of the IPv6 header.
func (h *IPv6Header) WithDSCP(dscp uint8) *IPv6Header {
	h.pb.Dscp = uint32(dscp)
	return h
}

// WithECN sets the ECN value of the IPv6 header.
func (h *IPv6Header) WithECN(ecn uint8) *IPv6Header {
	h.pb.Ecn = uint32(ecn)
	return h
}

func (h *IPv6Header) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Ipv6{h.pb}}
}

// NewMPLSHeader returns a new MPLS header.
// The header is initialized with a traffic class of 0 (best effort) and TTL of 255.
func NewMPLSHeader() *MPLSHeader {
	return &MPLSHeader{&opb.MplsHeader{Exp: 0, Ttl: 255}}
}

// MPLSHeader is a Multiprotocol Label Switching packet header.
type MPLSHeader struct {
	pb *opb.MplsHeader
}

// WithLabel sets the label of the MPLS header to the specified value.
// To generate a range of labels, use LabelRange() instead.
func (h *MPLSHeader) WithLabel(label uint32) *MPLSHeader {
	h.pb.Label = intRangeSingle(label)
	return h
}

// LabelRange sets the label of the MPLS header to a range of values and returns the range.
// By default, the range will be nonrandom values in the interval [0, 2^20).
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *MPLSHeader) LabelRange() *UIntRange {
	if h.pb.Label == nil {
		h.pb.Label = newFlowLabelRange()
	}
	return &UIntRange{pb: h.pb.Label}
}

// WithEXP sets the EXP (aka traffic class) of the MPLS header.
func (h *MPLSHeader) WithEXP(exp uint8) *MPLSHeader {
	h.pb.Exp = uint32(exp)
	return h
}

// WithTTL sets the TTL of the MPLS header.
func (h *MPLSHeader) WithTTL(ttl uint8) *MPLSHeader {
	h.pb.Ttl = uint32(ttl)
	return h
}

func (h *MPLSHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Mpls{h.pb}}
}

// NewTCPHeader returns a new TCP header.
// The header is initialized with none of its properties specified.
func NewTCPHeader() *TCPHeader {
	return &TCPHeader{&opb.TcpHeader{}}
}

// TCPHeader is a Transmission Control Protocol packet header.
type TCPHeader struct {
	pb *opb.TcpHeader
}

// WithSrcPort sets the source port of the TCP header to the specified value.
// To generate a range of source ports, use SrcPortRange() instead.
func (h *TCPHeader) WithSrcPort(port uint16) *TCPHeader {
	h.pb.SrcPort = intRangeSingle(uint32(port))
	return h
}

// SrcPortRange returns the range of source ports of the TCP header for further configuration.
// By default, the range will be nonrandom values in the interval [1, 2^16).
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *TCPHeader) SrcPortRange() *UIntRange {
	if h.pb.SrcPort == nil {
		h.pb.SrcPort = newPortRange()
	}
	return &UIntRange{pb: h.pb.SrcPort}
}

// WithDstPort sets the destination port of the TCP header to the specified value.
// To generate a range of destination ports, use DstPortRange() instead.
func (h *TCPHeader) WithDstPort(port uint16) *TCPHeader {
	h.pb.DstPort = intRangeSingle(uint32(port))
	return h
}

// DstPortRange returns the range of destination ports of the TCP header for further configuration.
// By default, the range will be nonrandom values in the interval [1, 2^16).
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *TCPHeader) DstPortRange() *UIntRange {
	if h.pb.DstPort == nil {
		h.pb.DstPort = newPortRange()
	}
	return &UIntRange{pb: h.pb.DstPort}
}

// WithSequenceNumber sets sequence number of the TCP header.
func (h *TCPHeader) WithSequenceNumber(seq uint32) *TCPHeader {
	h.pb.Seq = seq
	return h
}

func (h *TCPHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Tcp{h.pb}}
}

// NewUDPHeader returns a new UDP header.
// The header is initialized with none of its properties specified.
func NewUDPHeader() *UDPHeader {
	return &UDPHeader{&opb.UdpHeader{}}
}

// UDPHeader is a User Datagram Protocol packet header.
type UDPHeader struct {
	pb *opb.UdpHeader
}

// WithSrcPort sets the source port of the TCP header to the specified value.
// To generate a range of source ports, use SrcPortRange() instead.
func (h *UDPHeader) WithSrcPort(port uint16) *UDPHeader {
	h.pb.SrcPort = intRangeSingle(uint32(port))
	return h
}

// SrcPortRange returns the range of source ports of the TCP header for further configuration.
// By default, the range will be nonrandom values in the interval [1, 2^16).
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *UDPHeader) SrcPortRange() *UIntRange {
	if h.pb.SrcPort == nil {
		h.pb.SrcPort = newPortRange()
	}
	return &UIntRange{pb: h.pb.SrcPort}
}

// WithDstPort sets the destination port of the TCP header to the specified value.
// To generate a range of destination ports, use DstPortRange() instead.
func (h *UDPHeader) WithDstPort(port uint16) *UDPHeader {
	h.pb.DstPort = intRangeSingle(uint32(port))
	return h
}

// DstPortRange returns the range of destination ports of the TCP header for further configuration.
// By default, the range will be nonrandom values in the interval [1, 2^16).
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *UDPHeader) DstPortRange() *UIntRange {
	if h.pb.DstPort == nil {
		h.pb.DstPort = newPortRange()
	}
	return &UIntRange{pb: h.pb.DstPort}
}

func (h *UDPHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Udp{h.pb}}
}
