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
	"google.golang.org/protobuf/proto"

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

// WithEtherType sets the EtherType of the Ethernet header to the specified value.
func (h *EthernetHeader) WithEtherType(etherType uint32) *EthernetHeader {
	h.pb.EtherType = etherType
	return h
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
func (h *EthernetHeader) SrcAddressRange() *ixnet.AddressRange {
	if h.pb.SrcAddr == nil {
		h.pb.SrcAddr = newMACAddrRange()
	}
	return ixnet.NewAddressRange(h.pb.SrcAddr)
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
func (h *EthernetHeader) DstAddressRange() *ixnet.AddressRange {
	if h.pb.DstAddr == nil {
		h.pb.DstAddr = newMACAddrRange()
	}
	return ixnet.NewAddressRange(h.pb.DstAddr)
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
	return &IPv4Header{&opb.Ipv4Header{Ttl: proto.Uint32(64), Dscp: 0, Ecn: 0}}
}

// IPv4Header is an IPv4 packet header.
type IPv4Header struct {
	pb *opb.Ipv4Header
}

// WithDSCP sets the DSCP field of the IPv4 header.
func (h *IPv4Header) WithDSCP(dscp uint8) *IPv4Header {
	h.pb.Dscp = uint32(dscp)
	return h
}

// WithECN sets the ECN field of the IPv4 header.
func (h *IPv4Header) WithECN(ecn uint8) *IPv4Header {
	h.pb.Ecn = uint32(ecn)
	return h
}

// WithIdentification set identification field of IPv4 Header
func (h *IPv4Header) WithIdentification(identification int) *IPv4Header {
	h.pb.Identification = uint32(identification)
	return h
}

// WithDontFragment sets the "don't fragment" bit of the IPv4 header.
func (h *IPv4Header) WithDontFragment(dontFragment bool) *IPv4Header {
	h.pb.DontFragment = dontFragment
	return h
}

// WithMoreFragments sets the "more fragments" bit of the IPv4 Header
func (h *IPv4Header) WithMoreFragments(moreFragments bool) *IPv4Header {
	h.pb.MoreFragments = moreFragments
	return h
}

// WithFragmentOffset sets the fragment offset field of the IPv4 header.
func (h *IPv4Header) WithFragmentOffset(fragmentOffset int) *IPv4Header {
	h.pb.FragmentOffset = uint32(fragmentOffset)
	return h
}

// WithTTL sets the TTL of the IPv4 header.
func (h *IPv4Header) WithTTL(ttl uint8) *IPv4Header {
	h.pb.Ttl = proto.Uint32(uint32(ttl))
	return h
}

// WithProtocol sets the protocol field of the IPv4 header.
// If left unspecified, it will be inferred from the next header in the flow.
func (h *IPv4Header) WithProtocol(protocol int) *IPv4Header {
	h.pb.Protocol = proto.Uint32(uint32(protocol))
	return h
}

// WithHeaderChecksum sets the header checksum field of the IPv4 header.
func (h *IPv4Header) WithHeaderChecksum(checksum uint16) *IPv4Header {
	h.pb.Checksum = uint32(checksum)
	return h
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
func (h *IPv4Header) SrcAddressRange() *ixnet.AddressRange {
	if h.pb.SrcAddr == nil {
		h.pb.SrcAddr = newIPv4AddrRange()
	}
	return ixnet.NewAddressRange(h.pb.SrcAddr)
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
func (h *IPv4Header) DstAddressRange() *ixnet.AddressRange {
	if h.pb.DstAddr == nil {
		h.pb.DstAddr = newIPv4AddrRange()
	}
	return ixnet.NewAddressRange(h.pb.DstAddr)
}

func (h *IPv4Header) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Ipv4{h.pb}}
}

// NewIPv6Header returns a new IPv6 header.
// The header is initialized with a hop limit of 64, DSCP of 0 (best effort), and ECN of 0 (not ECN capable).
// Unless specified, a best effort will be made to infer the src and dst addresses from the topology.
func NewIPv6Header() *IPv6Header {
	return &IPv6Header{&opb.Ipv6Header{HopLimit: proto.Uint32(64), Dscp: 0, Ecn: 0}}
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
func (h *IPv6Header) SrcAddressRange() *ixnet.AddressRange {
	if h.pb.SrcAddr == nil {
		h.pb.SrcAddr = newIPv6AddrRange()
	}
	return ixnet.NewAddressRange(h.pb.SrcAddr)
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
func (h *IPv6Header) DstAddressRange() *ixnet.AddressRange {
	if h.pb.DstAddr == nil {
		h.pb.DstAddr = newIPv6AddrRange()
	}
	return ixnet.NewAddressRange(h.pb.DstAddr)
}

// WithFlowLabel sets the flow label of the IPv6 header.
func (h *IPv6Header) WithFlowLabel(flowLabel uint32) *IPv6Header {
	h.pb.FlowLabel = intRangeSingle(flowLabel)
	return h
}

// FlowLabelRange sets the flow label of the IPv6 header to a range of values and returns the range.
// By default, the range will be nonrandom values in the interval [0, 2^20).
// The count of values in the range is not set by default; the user must set it explicitly.
func (h *IPv6Header) FlowLabelRange() *ixnet.UIntRange {
	if h.pb.FlowLabel == nil {
		h.pb.FlowLabel = newFlowLabelRange()
	}
	return ixnet.NewUIntRange(h.pb.FlowLabel)
}

// WithHopLimit sets the hop limit of the IPv6 header.
func (h *IPv6Header) WithHopLimit(hopLimit uint8) *IPv6Header {
	h.pb.HopLimit = proto.Uint32(uint32(hopLimit))
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
func (h *MPLSHeader) LabelRange() *ixnet.UIntRange {
	if h.pb.Label == nil {
		h.pb.Label = newFlowLabelRange()
	}
	return ixnet.NewUIntRange(h.pb.Label)
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
func (h *TCPHeader) SrcPortRange() *ixnet.UIntRange {
	if h.pb.SrcPort == nil {
		h.pb.SrcPort = newPortRange()
	}
	return ixnet.NewUIntRange(h.pb.SrcPort)
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
func (h *TCPHeader) DstPortRange() *ixnet.UIntRange {
	if h.pb.DstPort == nil {
		h.pb.DstPort = newPortRange()
	}
	return ixnet.NewUIntRange(h.pb.DstPort)
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
func (h *UDPHeader) SrcPortRange() *ixnet.UIntRange {
	if h.pb.SrcPort == nil {
		h.pb.SrcPort = newPortRange()
	}
	return ixnet.NewUIntRange(h.pb.SrcPort)
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
func (h *UDPHeader) DstPortRange() *ixnet.UIntRange {
	if h.pb.DstPort == nil {
		h.pb.DstPort = newPortRange()
	}
	return ixnet.NewUIntRange(h.pb.DstPort)
}

func (h *UDPHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Udp{h.pb}}
}

// HTTPHeader is an HTTP packet header.
type HTTPHeader struct {
	pb *opb.HttpHeader
}

// NewHTTPHeader returns a new HTTP header.
// The header is initialized with none of its properties specified.
func NewHTTPHeader() *HTTPHeader {
	return &HTTPHeader{&opb.HttpHeader{}}
}

func (h *HTTPHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Http{h.pb}}
}

// ICMPHeader is an ICMP packet header.
type ICMPHeader struct {
	pb *opb.IcmpHeader
}

// NewICMPHeader returns a new ICMP header.
// The header is initialized as an Echo Reply.
func NewICMPHeader() *ICMPHeader {
	return &ICMPHeader{&opb.IcmpHeader{
		Type: &opb.IcmpHeader_EchoReply_{&opb.IcmpHeader_EchoReply{}},
	}}
}

func (h *ICMPHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Icmp{h.pb}}
}

// ICMPEchoReply is an ICMP echo reply message.
type ICMPEchoReply struct {
	pb *opb.IcmpHeader_EchoReply
}

// EchoReply sets the ICMPHeader message type to echo reply.
func (h *ICMPHeader) EchoReply() *ICMPEchoReply {
	epb := &opb.IcmpHeader_EchoReply{}
	h.pb.Type = &opb.IcmpHeader_EchoReply_{epb}
	return &ICMPEchoReply{epb}
}

// ICMPDestinationUnreachable is an ICMP destination unreachable message.
type ICMPDestinationUnreachable struct {
	pb *opb.IcmpHeader_DestinationUnreachable
}

// DestinationUnreachable sets the ICMPHeader message type to destination unreachable.
func (h *ICMPHeader) DestinationUnreachable() *ICMPDestinationUnreachable {
	dupb := &opb.IcmpHeader_DestinationUnreachable{}
	h.pb.Type = &opb.IcmpHeader_DestinationUnreachable_{dupb}
	return &ICMPDestinationUnreachable{dupb}
}

// WithNetworkUnreachable sets the destination unreachable code to network unreachable.
func (h *ICMPDestinationUnreachable) WithNetworkUnreachable() *ICMPDestinationUnreachable {
	h.pb.Code = opb.IcmpHeader_DestinationUnreachable_NETWORK_UNREACHABLE
	return h
}

// WithHostUnreachable sets the destination unreachable code to host unreachable.
func (h *ICMPDestinationUnreachable) WithHostUnreachable() *ICMPDestinationUnreachable {
	h.pb.Code = opb.IcmpHeader_DestinationUnreachable_HOST_UNREACHABLE
	return h
}

// WithProtocolUnreachable sets the destination unreachable code to protocol unreachable.
func (h *ICMPDestinationUnreachable) WithProtocolUnreachable() *ICMPDestinationUnreachable {
	h.pb.Code = opb.IcmpHeader_DestinationUnreachable_PROTOCOL_UNREACHABLE
	return h
}

// WithPortUnreachable sets the destination unreachable code to port unreachable.
func (h *ICMPDestinationUnreachable) WithPortUnreachable() *ICMPDestinationUnreachable {
	h.pb.Code = opb.IcmpHeader_DestinationUnreachable_PORT_UNREACHABLE
	return h
}

// WithFragmentationRequired sets the destination unreachable code to fragmentation required.
func (h *ICMPDestinationUnreachable) WithFragmentationRequired() *ICMPDestinationUnreachable {
	h.pb.Code = opb.IcmpHeader_DestinationUnreachable_FRAGMENTATION_REQUIRED
	return h
}

// WithSourceRouteFailed sets the destination unreachable code to source route failed.
func (h *ICMPDestinationUnreachable) WithSourceRouteFailed() *ICMPDestinationUnreachable {
	h.pb.Code = opb.IcmpHeader_DestinationUnreachable_SOURCE_ROUTE_FAILED
	return h
}

// ICMPRedirectMessage is an ICMP redirect message.
type ICMPRedirectMessage struct {
	pb *opb.IcmpHeader_RedirectMessage
}

// RedirectMessage sets the ICMPHeader message type to redirect message.
func (h *ICMPHeader) RedirectMessage() *ICMPRedirectMessage {
	rmpb := &opb.IcmpHeader_RedirectMessage{}
	h.pb.Type = &opb.IcmpHeader_RedirectMessage_{rmpb}
	return &ICMPRedirectMessage{rmpb}
}

// WithForNetwork sets the redirect message code to redirect for network.
func (h *ICMPRedirectMessage) WithForNetwork() *ICMPRedirectMessage {
	h.pb.Code = opb.IcmpHeader_RedirectMessage_NETWORK
	return h
}

// WithForHost sets the redirect message code to redirect for host.
func (h *ICMPRedirectMessage) WithForHost() *ICMPRedirectMessage {
	h.pb.Code = opb.IcmpHeader_RedirectMessage_HOST
	return h
}

// WithForToSAndNetwork sets the redirect message code to redirect for ToS and network.
func (h *ICMPRedirectMessage) WithForToSAndNetwork() *ICMPRedirectMessage {
	h.pb.Code = opb.IcmpHeader_RedirectMessage_TOS_AND_NETWORK
	return h
}

// WithForToSAndHost sets the redirect message code to redirect for ToS and host.
func (h *ICMPRedirectMessage) WithForToSAndHost() *ICMPRedirectMessage {
	h.pb.Code = opb.IcmpHeader_RedirectMessage_TOS_AND_HOST
	return h
}

// WithIP sets the redirect gateway IP address.
func (h *ICMPRedirectMessage) WithIP(ip string) *ICMPRedirectMessage {
	h.pb.IpAddr = ip
	return h
}

// ICMPEchoRequest is an ICMP echo request message.
type ICMPEchoRequest struct {
	pb *opb.IcmpHeader_EchoRequest
}

// EchoRequest sets the ICMPHeader message type to echo request.
func (h *ICMPHeader) EchoRequest() *ICMPEchoRequest {
	epb := &opb.IcmpHeader_EchoRequest{}
	h.pb.Type = &opb.IcmpHeader_EchoRequest_{epb}
	return &ICMPEchoRequest{epb}
}

// ICMPTimeExceeded is an ICMP time exceeded message.
type ICMPTimeExceeded struct {
	pb *opb.IcmpHeader_TimeExceeded
}

// TimeExceeded sets the ICMPHeader message type to time exceeded.
func (h *ICMPHeader) TimeExceeded() *ICMPTimeExceeded {
	tepb := &opb.IcmpHeader_TimeExceeded{}
	h.pb.Type = &opb.IcmpHeader_TimeExceeded_{tepb}
	return &ICMPTimeExceeded{tepb}
}

// WithTransit sets the time exceeded message code to time exceeded in transit.
func (h *ICMPTimeExceeded) WithTransit() *ICMPTimeExceeded {
	h.pb.Code = opb.IcmpHeader_TimeExceeded_TRANSIT
	return h
}

// WithFragmentReassembly sets the time exceeded message code to time exceeded in fragment reassembly.
func (h *ICMPTimeExceeded) WithFragmentReassembly() *ICMPTimeExceeded {
	h.pb.Code = opb.IcmpHeader_TimeExceeded_FRAGMENT_REASSEMBLY
	return h
}

// ICMPParameterProblem is an ICMP parameter problem message.
type ICMPParameterProblem struct {
	pb *opb.IcmpHeader_ParameterProblem
}

// ParameterProblem sets the ICMPHeader message type to parameter problem.
func (h *ICMPHeader) ParameterProblem() *ICMPParameterProblem {
	pppb := &opb.IcmpHeader_ParameterProblem{}
	h.pb.Type = &opb.IcmpHeader_ParameterProblem_{pppb}
	return &ICMPParameterProblem{pppb}
}

// WithPointer sets the pointer for the parameter problem message.
func (h *ICMPParameterProblem) WithPointer(pointer uint32) *ICMPParameterProblem {
	h.pb.Pointer = pointer
	return h
}

// ICMPTimestamp is an ICMP timestamp message.
type ICMPTimestamp struct {
	pb *opb.IcmpHeader_Timestamp
}

// Timestamp sets the ICMPHeader message type to timestamp message.
func (h *ICMPHeader) Timestamp() *ICMPTimestamp {
	tpb := &opb.IcmpHeader_Timestamp{}
	h.pb.Type = &opb.IcmpHeader_Timestamp_{tpb}
	return &ICMPTimestamp{tpb}
}

// WithID sets the ID for the timestamp message.
func (h *ICMPTimestamp) WithID(id uint32) *ICMPTimestamp {
	h.pb.Id = id
	return h
}

// WithSeq sets the sequence number for the timestamp message.
func (h *ICMPTimestamp) WithSeq(seq uint32) *ICMPTimestamp {
	h.pb.Seq = seq
	return h
}

// WithOriginateTimestamp sets the originate timestamp for the timestamp message.
func (h *ICMPTimestamp) WithOriginateTimestamp(ts uint32) *ICMPTimestamp {
	h.pb.OriginateTs = ts
	return h
}

// ICMPTimestampReply is an ICMP timestamp reply message.
type ICMPTimestampReply struct {
	pb *opb.IcmpHeader_TimestampReply
}

// TimestampReply sets the ICMPHeader message type to timestamp reply.
func (h *ICMPHeader) TimestampReply() *ICMPTimestampReply {
	trpb := &opb.IcmpHeader_TimestampReply{}
	h.pb.Type = &opb.IcmpHeader_TimestampReply_{trpb}
	return &ICMPTimestampReply{trpb}
}

// WithID sets the ID for the timestamp reply message.
func (h *ICMPTimestampReply) WithID(id uint32) *ICMPTimestampReply {
	h.pb.Id = id
	return h
}

// WithSeq sets the sequence number for the timestamp reply message.
func (h *ICMPTimestampReply) WithSeq(seq uint32) *ICMPTimestampReply {
	h.pb.Seq = seq
	return h
}

// WithOriginateTimestamp sets the originate timestamp for the timestamp reply message.
func (h *ICMPTimestampReply) WithOriginateTimestamp(ts uint32) *ICMPTimestampReply {
	h.pb.OriginateTs = ts
	return h
}

// WithReceiveTimestamp sets the receive timestamp for the timestamp reply message.
func (h *ICMPTimestampReply) WithReceiveTimestamp(ts uint32) *ICMPTimestampReply {
	h.pb.ReceiveTs = ts
	return h
}

// WithTransmitTimestamp sets the transmit timestamp for the timestamp reply message.
func (h *ICMPTimestampReply) WithTransmitTimestamp(ts uint32) *ICMPTimestampReply {
	h.pb.TransmitTs = ts
	return h
}

// OSPFHeader is an OSPF packet header.
type OSPFHeader struct {
	pb *opb.OspfHeader
}

// NewOSPFHeader returns a new OSPF header.
// The header is initialized as a default Hello packet.
func NewOSPFHeader() *OSPFHeader {
	hdr := &OSPFHeader{&opb.OspfHeader{}}
	hdr.Hello()
	return hdr
}

func (h *OSPFHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Ospf{h.pb}}
}

// OSPFHello is an OSPF hello message.
type OSPFHello struct {
	pb *opb.OspfHeader_Hello
}

// Hello sets the OSPFHeader message type to a hello message with the following
// default values:
// - Network Mask: 24
// - Hello Interval: 10
// - Router Dead Interval: 40
func (h *OSPFHeader) Hello() *OSPFHello {
	hpb := &opb.OspfHeader_Hello{
		NetworkMaskLength:     24,
		HelloIntervalSec:      10,
		RouterDeadIntervalSec: 40,
	}
	h.pb.Type = &opb.OspfHeader_Hello_{hpb}
	return &OSPFHello{hpb}
}

// WithNetworkMaskLength sets the network mask length for the hello message.
func (h *OSPFHello) WithNetworkMaskLength(length uint32) *OSPFHello {
	h.pb.NetworkMaskLength = length
	return h
}

// WithHelloIntervalSec sets the interval for the hello message.
func (h *OSPFHello) WithHelloIntervalSec(secs uint32) *OSPFHello {
	h.pb.HelloIntervalSec = secs
	return h
}

// WithRouterPriority sets the router priority for the hello message.
func (h *OSPFHello) WithRouterPriority(prio uint32) *OSPFHello {
	h.pb.RouterPriority = prio
	return h
}

// WithRouterDeadIntervalSec sets the dead interval for the hello message.
func (h *OSPFHello) WithRouterDeadIntervalSec(secs uint32) *OSPFHello {
	h.pb.RouterDeadIntervalSec = secs
	return h
}

// WithDesignatedRouter sets the designated router for the hello message.
func (h *OSPFHello) WithDesignatedRouter(ip string) *OSPFHello {
	h.pb.DesignatedRouter = ip
	return h
}

// WithBackupDesignatedRouter sets the backup designated router for the hello message.
func (h *OSPFHello) WithBackupDesignatedRouter(ip string) *OSPFHello {
	h.pb.BackupDesignatedRouter = ip
	return h
}

// WithNeighbors sets the neighbor routers for the hello message.
func (h *OSPFHello) WithNeighbors(ips ...string) *OSPFHello {
	h.pb.Neighbors = ips
	return h
}

// OSPFDatabaseDescription is an OSPF database description message.
type OSPFDatabaseDescription struct {
	pb *opb.OspfHeader_DatabaseDescription
}

// DatabaseDescription sets the OSPFHeader message type to a DBD message with
// MTU set to 1500.
func (h *OSPFHeader) DatabaseDescription() *OSPFDatabaseDescription {
	dpb := &opb.OspfHeader_DatabaseDescription{Mtu: 1500}
	h.pb.Type = &opb.OspfHeader_Dbd{dpb}
	return &OSPFDatabaseDescription{dpb}
}

// WithMTU sets the MTU field of the database description message.
func (h *OSPFDatabaseDescription) WithMTU(mtu uint32) *OSPFDatabaseDescription {
	h.pb.Mtu = mtu
	return h
}

// WithInitial sets the 'initial' bit of the database description message.
func (h *OSPFDatabaseDescription) WithInitial(initial bool) *OSPFDatabaseDescription {
	h.pb.Initial = initial
	return h
}

// WithMore sets the 'more' bit of the database description message.
func (h *OSPFDatabaseDescription) WithMore(more bool) *OSPFDatabaseDescription {
	h.pb.More = more
	return h
}

// WithMaster sets the 'master' bit of the database description message.
func (h *OSPFDatabaseDescription) WithMaster(master bool) *OSPFDatabaseDescription {
	h.pb.Master = master
	return h
}

// WithSeq sets the sequence number of the database description message.
func (h *OSPFDatabaseDescription) WithSeq(seq uint32) *OSPFDatabaseDescription {
	h.pb.Seq = seq
	return h
}

// OSPFLinkStateRequest is an OSPF link state request message.
type OSPFLinkStateRequest struct {
	pb *opb.OspfHeader_LinkStateRequest
}

// LinkStateRequest sets the OSPFHeader message type to a link state request,
// defaulting to asking for a Router link.
func (h *OSPFHeader) LinkStateRequest() *OSPFLinkStateRequest {
	lpb := &opb.OspfHeader_LinkStateRequest{Type: opb.OspfHeader_LINK_STATE_TYPE_ROUTER}
	h.pb.Type = &opb.OspfHeader_Lsr{lpb}
	return &OSPFLinkStateRequest{lpb}
}

// WithLinkStateTypeRouter sets the link state request type to router.
func (h *OSPFLinkStateRequest) WithLinkStateTypeRouter() *OSPFLinkStateRequest {
	h.pb.Type = opb.OspfHeader_LINK_STATE_TYPE_ROUTER
	return h
}

// WithLinkStateTypeNetwork sets the link state request type to network.
func (h *OSPFLinkStateRequest) WithLinkStateTypeNetwork() *OSPFLinkStateRequest {
	h.pb.Type = opb.OspfHeader_LINK_STATE_TYPE_NETWORK
	return h
}

// WithLinkStateTypeSummaryNetwork sets the link state request type to a network summary.
func (h *OSPFLinkStateRequest) WithLinkStateTypeSummaryNetwork() *OSPFLinkStateRequest {
	h.pb.Type = opb.OspfHeader_LINK_STATE_TYPE_SUMMARY_NETWORK
	return h
}

// WithLinkStateTypeSummaryASBR sets the link state request type to an AS boundary router summary.
func (h *OSPFLinkStateRequest) WithLinkStateTypeSummaryASBR() *OSPFLinkStateRequest {
	h.pb.Type = opb.OspfHeader_LINK_STATE_TYPE_SUMMARY_ASBR
	return h
}

// WithLinkStateTypeASExternal sets the link state request type to AS external.
func (h *OSPFLinkStateRequest) WithLinkStateTypeASExternal() *OSPFLinkStateRequest {
	h.pb.Type = opb.OspfHeader_LINK_STATE_TYPE_AS_EXTERNAL
	return h
}

// WithLinkStateID sets the link state ID of the link state request.
func (h *OSPFLinkStateRequest) WithLinkStateID(id string) *OSPFLinkStateRequest {
	h.pb.LinkStateId = id
	return h
}

// WithAdvertisingRouter sets the advertising router ID of the link state request.
func (h *OSPFLinkStateRequest) WithAdvertisingRouter(id string) *OSPFLinkStateRequest {
	h.pb.AdvertisingRouter = id
	return h
}

// OSPFLinkStateAdvertisementHeader is an OSPF link state advertisement header.
type OSPFLinkStateAdvertisementHeader struct {
	pb *opb.OspfHeader_LinkStateAdvertisementHeader
}

// WithAge sets the age in seconds of the link state advertisement.
func (h *OSPFLinkStateAdvertisementHeader) WithAge(sec uint32) *OSPFLinkStateAdvertisementHeader {
	h.pb.AgeSeconds = sec
	return h
}

// WithLinkStateTypeRouter sets the link state type to router.
func (h *OSPFLinkStateAdvertisementHeader) WithLinkStateTypeRouter() *OSPFLinkStateAdvertisementHeader {
	h.pb.Type = opb.OspfHeader_LINK_STATE_TYPE_ROUTER
	return h
}

// WithLinkStateTypeNetwork sets the link state type to network.
func (h *OSPFLinkStateAdvertisementHeader) WithLinkStateTypeNetwork() *OSPFLinkStateAdvertisementHeader {
	h.pb.Type = opb.OspfHeader_LINK_STATE_TYPE_NETWORK
	return h
}

// WithLinkStateTypeSummaryNetwork sets the link state type to a network summary.
func (h *OSPFLinkStateAdvertisementHeader) WithLinkStateTypeSummaryNetwork() *OSPFLinkStateAdvertisementHeader {
	h.pb.Type = opb.OspfHeader_LINK_STATE_TYPE_SUMMARY_NETWORK
	return h
}

// WithLinkStateTypeSummaryASBR sets the link state type to an AS boundary router summary.
func (h *OSPFLinkStateAdvertisementHeader) WithLinkStateTypeSummaryASBR() *OSPFLinkStateAdvertisementHeader {
	h.pb.Type = opb.OspfHeader_LINK_STATE_TYPE_SUMMARY_ASBR
	return h
}

// WithLinkStateTypeASExternal sets the link state type to AS external.
func (h *OSPFLinkStateAdvertisementHeader) WithLinkStateTypeASExternal() *OSPFLinkStateAdvertisementHeader {
	h.pb.Type = opb.OspfHeader_LINK_STATE_TYPE_AS_EXTERNAL
	return h
}

// WithLinkStateID sets the ID of the link state advertisement.
func (h *OSPFLinkStateAdvertisementHeader) WithLinkStateID(id string) *OSPFLinkStateAdvertisementHeader {
	h.pb.LinkStateId = id
	return h
}

// WithAdvertisingRouter sets the advertising router ID of the link state advertisement.
func (h *OSPFLinkStateAdvertisementHeader) WithAdvertisingRouter(id string) *OSPFLinkStateAdvertisementHeader {
	h.pb.AdvertisingRouter = id
	return h
}

// WithSeq sets the sequence number of the link state advertisement.
func (h *OSPFLinkStateAdvertisementHeader) WithSeq(seq uint32) *OSPFLinkStateAdvertisementHeader {
	h.pb.Seq = seq
	return h
}

// OSPFLinkStateUpdate is an OSPF link state update message.
type OSPFLinkStateUpdate struct {
	pb *opb.OspfHeader_LinkStateUpdate
}

// LinkStateUpdate sets the OSPFHeader message type to a link state update.
func (h *OSPFHeader) LinkStateUpdate() *OSPFLinkStateUpdate {
	lpb := &opb.OspfHeader_LinkStateUpdate{}
	h.pb.Type = &opb.OspfHeader_Lsu{lpb}
	return &OSPFLinkStateUpdate{lpb}
}

// OSPFAdvertisement respresents an OSPF link state update advertisement.
type OSPFAdvertisement struct {
	pb *opb.OspfHeader_LinkStateUpdate_Advertisement
}

// AddAdvertisement adds a link state advertisement to the link state update message.
// Defaults to setting the link state type to 'router'.
func (h *OSPFLinkStateUpdate) AddAdvertisement() *OSPFAdvertisement {
	apb := &opb.OspfHeader_LinkStateUpdate_Advertisement{
		Header: &opb.OspfHeader_LinkStateAdvertisementHeader{
			Type: opb.OspfHeader_LINK_STATE_TYPE_ROUTER,
		},
	}
	h.pb.Advertisements = append(h.pb.Advertisements, apb)
	return &OSPFAdvertisement{apb}
}

// AdvertisementHeader returns the advertisement header for the link state advertisement.
func (h *OSPFAdvertisement) AdvertisementHeader() *OSPFLinkStateAdvertisementHeader {
	return &OSPFLinkStateAdvertisementHeader{h.pb.Header}
}

// OSPFLinkStateAck is an OSPF link state ack message.
type OSPFLinkStateAck struct {
	pb *opb.OspfHeader_LinkStateAck
}

// LinkStateAck sets the OSPFHeader message type to a link state ack.
func (h *OSPFHeader) LinkStateAck() *OSPFLinkStateAck {
	lpb := &opb.OspfHeader_LinkStateAck{}
	h.pb.Type = &opb.OspfHeader_Lsa{lpb}
	return &OSPFLinkStateAck{lpb}
}

// AddAdvertisementHeader adds a link state advertisement header to the link
// state ack message.
// Defaults to setting the link state type to 'router'.
func (h *OSPFLinkStateAck) AddAdvertisementHeader() *OSPFLinkStateAdvertisementHeader {
	hpb := &opb.OspfHeader_LinkStateAdvertisementHeader{
		Type: opb.OspfHeader_LINK_STATE_TYPE_ROUTER,
	}
	h.pb.Headers = append(h.pb.Headers, hpb)
	return &OSPFLinkStateAdvertisementHeader{hpb}
}

// RSVPHeader is an RSVP packet header.
type RSVPHeader struct {
	pb *opb.RsvpHeader
}

// NewRSVPHeader returns a new RSVP header.
// The header is initialized with a default message type of 'Path'.
func NewRSVPHeader() *RSVPHeader {
	hdr := &RSVPHeader{&opb.RsvpHeader{
		MessageType: opb.RsvpHeader_PATH,
	}}
	return hdr
}

func (h *RSVPHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Rsvp{h.pb}}
}

// WithVersion sets the version for the RSVP header.
func (h *RSVPHeader) WithVersion(v uint32) *RSVPHeader {
	h.pb.Version = v
	return h
}

// WithRefreshReductionCapable sets the refresh reduction bit of the RSVP
// header.
func (h *RSVPHeader) WithRefreshReductionCapable(rr bool) *RSVPHeader {
	h.pb.RefreshReductionCapable = rr
	return h
}

// WithMessageTypePath sets the message type of the RSVP header to 'Path'.
func (h *RSVPHeader) WithMessageTypePath() *RSVPHeader {
	h.pb.MessageType = opb.RsvpHeader_PATH
	return h
}

// WithMessageTypeResv sets the message type of the RSVP header to 'Resv'.
func (h *RSVPHeader) WithMessageTypeResv() *RSVPHeader {
	h.pb.MessageType = opb.RsvpHeader_RESV
	return h
}

// WithTTL sets the TTL of the RSVP header.
func (h *RSVPHeader) WithTTL(ttl uint32) *RSVPHeader {
	h.pb.Ttl = ttl
	return h
}

// PIMHeader is a PIM packet header.
type PIMHeader struct {
	pb *opb.PimHeader
}

// NewPIMHeader returns a new PIM header.
// The header is initialized with a default message type of 'Hello'.
func NewPIMHeader() *PIMHeader {
	hdr := &PIMHeader{&opb.PimHeader{
		Type: &opb.PimHeader_Hello_{Hello: &opb.PimHeader_Hello{}},
	}}
	return hdr
}

func (h *PIMHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Pim{h.pb}}
}

// Hello sets the PIMHeader message type to a hello message.
func (h *PIMHeader) Hello() *PIMHello {
	hpb := &opb.PimHeader_Hello{}
	h.pb.Type = &opb.PimHeader_Hello_{hpb}
	return &PIMHello{hpb}
}

// PIMHello is a PIM hello message.
type PIMHello struct {
	pb *opb.PimHeader_Hello
}

// LDPHeader is an LDP packet header.
type LDPHeader struct {
	pb *opb.LdpHeader
}

// NewLDPHeader returns a new LDP header.
// The header is initialized as a default Hello message.
func NewLDPHeader() *LDPHeader {
	hdr := &LDPHeader{&opb.LdpHeader{}}
	hdr.Hello()
	return hdr
}

func (h *LDPHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Ldp{h.pb}}
}

// WithLSRID sets the LSR ID.
func (h *LDPHeader) WithLSRID(id string) *LDPHeader {
	h.pb.LsrId = id
	return h
}

// WithLabelSpace sets the label space.
func (h *LDPHeader) WithLabelSpace(label uint16) *LDPHeader {
	h.pb.LabelSpace = uint32(label)
	return h
}

// LDPHello is an LDP hello message.
type LDPHello struct {
	pb *opb.LdpHeader_Hello
}

// Hello sets the LDPHeader message type to a hello message with the following
// default values:
// - Hold time: 15
// - Targeted: false
// - Request Targeted: false
func (h *LDPHeader) Hello() *LDPHello {
	hpb := &opb.LdpHeader_Hello{HoldTimeSec: 15}
	h.pb.Type = &opb.LdpHeader_Hello_{hpb}
	return &LDPHello{hpb}
}

// WithHoldTimeSec sets the hold time in the Hello message.
func (h *LDPHello) WithHoldTimeSec(secs uint16) *LDPHello {
	h.pb.HoldTimeSec = uint32(secs)
	return h
}

// WithTargeted specifies if the message is a targeted hello.
func (h *LDPHello) WithTargeted(targeted bool) *LDPHello {
	h.pb.Targeted = targeted
	return h
}

// WithRequestTargeted specifies if the message is requesting targeted Hello replies.
func (h *LDPHello) WithRequestTargeted(targeted bool) *LDPHello {
	h.pb.RequestTargeted = targeted
	return h
}

// NewMACsecHeader returns a new MACsec header.
func NewMACsecHeader() *MACsecHeader {
	return &MACsecHeader{&opb.MacsecHeader{}}
}

// MACsecHeader is a MACsec packet header.
type MACsecHeader struct {
	pb *opb.MacsecHeader
}

func (h *MACsecHeader) asPB() *opb.Header {
	return &opb.Header{Type: &opb.Header_Macsec{h.pb}}
}

func intRangeSingle(i uint32) *opb.UIntRange {
	return &opb.UIntRange{Min: i, Max: i, Count: 1}
}

func addrRangeSingle(a string) *opb.AddressRange {
	return &opb.AddressRange{Min: a, Max: a, Count: 1}
}

func newPortRange() *opb.UIntRange {
	const maxPort uint32 = (1 << 16) - 1
	return &opb.UIntRange{Min: 1, Max: maxPort}
}

func newFlowLabelRange() *opb.UIntRange {
	const maxFlowLabel uint32 = (1 << 20) - 1
	return &opb.UIntRange{Max: maxFlowLabel}
}

func newMACAddrRange() *opb.AddressRange {
	return &opb.AddressRange{Min: "00:00:00:00:00:01", Max: "FF:FF:FF:FF:FF:FE"}
}

func newIPv4AddrRange() *opb.AddressRange {
	return &opb.AddressRange{Min: "0.0.0.1", Max: "255.255.255.254"}
}

func newIPv6AddrRange() *opb.AddressRange {
	return &opb.AddressRange{Min: "::1", Max: "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"}
}
