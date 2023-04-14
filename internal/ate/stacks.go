// Copyright 2021 Google LLC
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

package ate

import (
	"errors"
	"fmt"
	"net"

	"github.com/openconfig/ondatra/internal/ixconfig"

	opb "github.com/openconfig/ondatra/proto"
)

func headerStacks(hdr *opb.Header, idx int, hasSrcVLAN bool) ([]*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	switch v := hdr.Type.(type) {
	case *opb.Header_Eth:
		s, err := ethStack(v.Eth, idx)
		if err != nil {
			return nil, err
		}
		stacks := []*ixconfig.TrafficTrafficItemConfigElementStack{s}
		if hasSrcVLAN || v.Eth.GetVlanId() != 0 {
			s, err := vlanStack(v.Eth.GetVlanId(), idx+1)
			if err != nil {
				return nil, err
			}
			stacks = append(stacks, s)
		}
		return stacks, nil
	case *opb.Header_Gre:
		s, err := greStack(v.Gre, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Ipv4:
		s, err := ipv4Stack(v.Ipv4, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Ipv6:
		s, err := ipv6Stack(v.Ipv6, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Mpls:
		s, err := mplsStack(v.Mpls, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Tcp:
		s, err := tcpStack(v.Tcp, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Udp:
		s, err := udpStack(v.Udp, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Http:
		s, err := httpStack(v.Http, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Icmp:
		s, err := icmpStack(v.Icmp, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Ospf:
		s, err := ospfStack(v.Ospf, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Rsvp:
		s, err := rsvpStack(v.Rsvp, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Pim:
		s, err := pimStack(v.Pim, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Ldp:
		s, err := ldpStack(v.Ldp, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_Macsec:
		s, err := macsecStack(v.Macsec, idx)
		if err != nil {
			return nil, err
		}
		ppts, err := payloadProtocolTypeStack(idx + 1)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s, ppts}, nil
	case *opb.Header_Esp:
		s, err := espStack(v.Esp, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	case *opb.Header_EspOverMacsec:
		s, err := espOverMacsecStack(v.EspOverMacsec, idx)
		if err != nil {
			return nil, err
		}
		return []*ixconfig.TrafficTrafficItemConfigElementStack{s}, nil
	default:
		return nil, fmt.Errorf("unrecognized header type: %v", hdr)
	}
}

func ethStack(eth *opb.EthernetHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	if eth.GetBadCrc() && idx != 0 {
		return nil, fmt.Errorf("can only set a bad CRC on the initial ethernet header: %v", eth)
	}
	stack := ixconfig.NewEthernetStack(idx)
	if eth.GetSrcAddr() != nil {
		if err := setAddrRangeField(stack.SourceAddress(), mac48AddrType, eth.GetSrcAddr()); err != nil {
			return nil, fmt.Errorf("could not set source MAC address: %w", err)
		}
	}
	if eth.GetDstAddr() != nil {
		if err := setAddrRangeField(stack.DestinationAddress(), mac48AddrType, eth.GetDstAddr()); err != nil {
			return nil, fmt.Errorf("could not set destination MAC address: %w", err)
		}
	}
	if eth.GetEtherType() > 0 {
		setSingleValue(stack.EtherType(), uintToHexStr(eth.GetEtherType()))
	}
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func vlanStack(vlanID uint32, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewVlanStack(idx)
	if vlanID > 0 {
		setSingleValue(stack.VlanTagVlanID(), uintToStr(vlanID))
	}
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func greStack(gre *opb.GreHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewGreStack(idx)
	setSingleValue(stack.KeyHolderKey(), uintToStr(gre.GetKey()))
	setSingleValue(stack.SequenceHolderSequenceNum(), uintToStr(gre.GetSeq()))
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func ipv4Stack(ipv4 *opb.Ipv4Header, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewIpv4Stack(idx)
	if ds := (ipv4.GetDscp() << 2) | ipv4.GetEcn(); ds > 0 {
		tc := stack.PriorityRaw()
		setSingleValue(tc, uintToHexStr(ds))
		tc.ActiveFieldChoice = ixconfig.Bool(true)
	}
	if id := ipv4.GetIdentification(); id > 0 {
		setSingleValue(stack.Identification(), uintToStr(id))
	}
	if ipv4.GetDontFragment() {
		setSingleValue(stack.FlagsFragment(), ixconfig.String("1"))
	}
	if ipv4.GetMoreFragments() {
		setSingleValue(stack.FlagsLastFragment(), ixconfig.String("1"))
	}
	if offset := ipv4.GetFragmentOffset(); offset > 0 {
		setSingleValue(stack.FragmentOffset(), uintToStr(offset))
	}
	if ttl := ipv4.Ttl; ttl != nil {
		setSingleValue(stack.Ttl(), uintToStr(*ttl))
	}
	if protocol := ipv4.Protocol; protocol != nil {
		setSingleValue(stack.Protocol(), uintToStr(*protocol))
	}
	if checksum := ipv4.GetChecksum(); checksum > 0 {
		setSingleValue(stack.Checksum(), uintToHexStr(checksum))
	}
	if ipv4.GetSrcAddr() != nil {
		if err := setAddrRangeField(stack.SrcIp(), ipv4AddrType, ipv4.GetSrcAddr()); err != nil {
			return nil, fmt.Errorf("could not set IPv4 source address: %w", err)
		}
	}
	if ipv4.GetDstAddr() != nil {
		if err := setAddrRangeField(stack.DstIp(), ipv4AddrType, ipv4.GetDstAddr()); err != nil {
			return nil, fmt.Errorf("could not set IPv4 destination address: %w", err)
		}
	}
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func ipv6Stack(ipv6 *opb.Ipv6Header, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewIpv6Stack(idx)
	ds := (ipv6.GetDscp() << 2) | ipv6.GetEcn()
	tc := stack.VersionTrafficClassFlowLabelTrafficClass()
	setSingleValue(tc, uintToStr(ds))
	if ipv6.GetFlowLabel() != nil {
		if err := setUintRangeField(stack.VersionTrafficClassFlowLabelFlowLabel(), ipv6.GetFlowLabel()); err != nil {
			return nil, fmt.Errorf("could not set IPv6 flow label: %w", err)
		}
	}
	if hopLimit := ipv6.HopLimit; hopLimit != nil {
		setSingleValue(stack.HopLimit(), uintToStr(*hopLimit))
	}
	if ipv6.GetSrcAddr() != nil {
		if err := setAddrRangeField(stack.SrcIP(), ipv6AddrType, ipv6.GetSrcAddr()); err != nil {
			return nil, fmt.Errorf("could not set IPv6 source address: %w", err)
		}
	}
	if ipv6.GetDstAddr() != nil {
		if err := setAddrRangeField(stack.DstIP(), ipv6AddrType, ipv6.GetDstAddr()); err != nil {
			return nil, fmt.Errorf("could not set IPv6 destination address: %w", err)
		}
	}
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func mplsStack(mpls *opb.MplsHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewMplsStack(idx)
	if err := setUintRangeField(stack.Value(), mpls.GetLabel()); err != nil {
		return nil, fmt.Errorf("could not set MPLS label: %w", err)
	}
	setSingleValue(stack.Experimental(), uintToStr(mpls.GetExp()))
	setSingleValue(stack.Ttl(), uintToStr(mpls.GetTtl()))
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func tcpStack(tcp *opb.TcpHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewTcpStack(idx)
	if tcp.GetSrcPort() != nil {
		if err := setUintRangeField(stack.SrcPort(), tcp.GetSrcPort()); err != nil {
			return nil, fmt.Errorf("could not set TCP source port: %w", err)
		}
	}
	if tcp.GetDstPort() != nil {
		if err := setUintRangeField(stack.DstPort(), tcp.GetDstPort()); err != nil {
			return nil, fmt.Errorf("could not set TCP destination port: %w", err)
		}
	}
	setSingleValue(stack.SequenceNumber(), uintToStr(tcp.GetSeq()))
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func udpStack(udp *opb.UdpHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewUdpStack(idx)
	if udp.GetSrcPort() != nil {
		if err := setUintRangeField(stack.SrcPort(), udp.GetSrcPort()); err != nil {
			return nil, fmt.Errorf("could not set UDP source port: %w", err)
		}
	}
	if udp.GetDstPort() != nil {
		if err := setUintRangeField(stack.DstPort(), udp.GetDstPort()); err != nil {
			return nil, fmt.Errorf("could not set UDP destination port: %w", err)
		}
	}
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func httpStack(http *opb.HttpHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	return ixconfig.NewHTTP_GETStack(idx).TrafficTrafficItemConfigElementStack(), nil
}

func icmpStack(icmp *opb.IcmpHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	switch icmp.Type.(type) {
	case *opb.IcmpHeader_DestinationUnreachable_,
		*opb.IcmpHeader_RedirectMessage_,
		*opb.IcmpHeader_TimeExceeded_,
		*opb.IcmpHeader_ParameterProblem_:
		return icmpV1Stack(icmp, idx)
	case *opb.IcmpHeader_EchoReply_,
		*opb.IcmpHeader_EchoRequest_,
		*opb.IcmpHeader_Timestamp_,
		*opb.IcmpHeader_TimestampReply_:
		return icmpV2Stack(icmp, idx)
	default:
		return nil, fmt.Errorf("unrecognized ICMP header type: %v", icmp)
	}
}

var dstUnreachableToCode = map[opb.IcmpHeader_DestinationUnreachable_Code]uint32{
	opb.IcmpHeader_DestinationUnreachable_NETWORK_UNREACHABLE:    0,
	opb.IcmpHeader_DestinationUnreachable_HOST_UNREACHABLE:       1,
	opb.IcmpHeader_DestinationUnreachable_PROTOCOL_UNREACHABLE:   2,
	opb.IcmpHeader_DestinationUnreachable_PORT_UNREACHABLE:       3,
	opb.IcmpHeader_DestinationUnreachable_FRAGMENTATION_REQUIRED: 4,
	opb.IcmpHeader_DestinationUnreachable_SOURCE_ROUTE_FAILED:    5,
}

var timeExceededToCode = map[opb.IcmpHeader_TimeExceeded_Code]uint32{
	opb.IcmpHeader_TimeExceeded_TRANSIT:             0,
	opb.IcmpHeader_TimeExceeded_FRAGMENT_REASSEMBLY: 1,
}

var redirectToCode = map[opb.IcmpHeader_RedirectMessage_Code]uint32{
	opb.IcmpHeader_RedirectMessage_NETWORK:         0,
	opb.IcmpHeader_RedirectMessage_HOST:            1,
	opb.IcmpHeader_RedirectMessage_TOS_AND_NETWORK: 2,
	opb.IcmpHeader_RedirectMessage_TOS_AND_HOST:    3,
}

func icmpV1Stack(icmp *opb.IcmpHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewIcmpv1Stack(idx)
	switch ih := icmp.Type.(type) {
	case *opb.IcmpHeader_DestinationUnreachable_:
		setSingleValue(stack.MessageType(), uintToStr(3))

		code := ih.DestinationUnreachable.GetCode()
		codeVal, ok := dstUnreachableToCode[code]
		if !ok {
			return nil, fmt.Errorf("unrecognized ICMP Destination Unreachable code: %v", code)
		}
		setSingleValue(stack.CodeOptionsDestUnreachableCodeOptions(), uintToStr(codeVal))
	case *opb.IcmpHeader_RedirectMessage_:
		setSingleValue(stack.MessageType(), uintToStr(5))

		code := ih.RedirectMessage.GetCode()
		codeVal, ok := redirectToCode[code]
		if !ok {
			return nil, fmt.Errorf("unrecognized ICMP Redirect code: %v", code)
		}
		setSingleValue(stack.CodeOptionsRedirectMessageOptions(), uintToStr(codeVal))

		ip, isV6 := parseIP(ih.RedirectMessage.GetIpAddr())
		if ip == nil || isV6 {
			return nil, fmt.Errorf("provided gateway IP %q for ICMP Redirect is not a V4 address", ih.RedirectMessage.GetIpAddr())
		}
		setSingleValue(stack.Next4BytesGatewayInternetAddress(), ixconfig.String(ih.RedirectMessage.GetIpAddr()))
	case *opb.IcmpHeader_TimeExceeded_:
		setSingleValue(stack.MessageType(), uintToStr(11))

		code := ih.TimeExceeded.GetCode()
		codeVal, ok := timeExceededToCode[code]
		if !ok {
			return nil, fmt.Errorf("unrecognized ICMP TimeExceeded code: %v", code)
		}
		setSingleValue(stack.CodeOptionsTimeExceededOptions(), uintToStr(codeVal))
	case *opb.IcmpHeader_ParameterProblem_:
		setSingleValue(stack.MessageType(), uintToStr(12))
		setSingleValue(stack.Next4BytesNextFieldsForParameterProblemPointer(), uintToStr(ih.ParameterProblem.GetPointer()))
	default:
		// Should not happen as header validity is checked in icmpStack function
		return nil, fmt.Errorf("invalid ICMP header type for IxNetwork icmpV1 stack: %v", icmp)
	}
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func icmpV2Stack(icmp *opb.IcmpHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewIcmpv2Stack(idx)
	switch ih := icmp.Type.(type) {
	case *opb.IcmpHeader_EchoReply_:
		setSingleValue(stack.MessageType(), uintToStr(0))
	case *opb.IcmpHeader_EchoRequest_:
		setSingleValue(stack.MessageType(), uintToStr(8))
	case *opb.IcmpHeader_Timestamp_:
		setSingleValue(stack.MessageType(), uintToStr(13))
		setSingleValue(stack.Identifier(), uintToStr(ih.Timestamp.GetId()))
		setSingleValue(stack.SequenceNumber(), uintToStr(ih.Timestamp.GetSeq()))
		setSingleValue(stack.NextFieldsFieldsForTimeStampMsgOrigTmpStmp1(), uintToStr(ih.Timestamp.GetOriginateTs()))
	case *opb.IcmpHeader_TimestampReply_:
		setSingleValue(stack.MessageType(), uintToStr(14))
		setSingleValue(stack.Identifier(), uintToStr(ih.TimestampReply.GetId()))
		setSingleValue(stack.SequenceNumber(), uintToStr(ih.TimestampReply.GetSeq()))
		setSingleValue(stack.NextFieldsFieldsForTimeStampReplyOrigTmpStmp2(), uintToStr(ih.TimestampReply.GetOriginateTs()))
		setSingleValue(stack.NextFieldsFieldsForTimeStampReplyRcvTmpStmp2(), uintToStr(ih.TimestampReply.GetReceiveTs()))
		setSingleValue(stack.NextFieldsFieldsForTimeStampReplyTransTmpStmp2(), uintToStr(ih.TimestampReply.GetTransmitTs()))
	default:
		// Should not happen as header validity is checked in icmpStack function
		return nil, fmt.Errorf("invalid ICMP header type for IxNetwork icmpV2 stack: %v", icmp)
	}
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

type ospfTrafficTrafficItemConfigElementStack interface {
	Ospfv2PacketHeaderRouterID() *ixconfig.TrafficTrafficItemConfigElementStackField
	Ospfv2PacketHeaderAreaID() *ixconfig.TrafficTrafficItemConfigElementStackField
	TrafficTrafficItemConfigElementStack() *ixconfig.TrafficTrafficItemConfigElementStack
}

func ospfStack(ospf *opb.OspfHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	var stack ospfTrafficTrafficItemConfigElementStack
	var err error
	switch o := ospf.Type.(type) {
	case *opb.OspfHeader_Hello_:
		stack, err = ospfHello(o.Hello, idx)
	case *opb.OspfHeader_Dbd:
		stack = ospfDBD(o.Dbd, idx)
	case *opb.OspfHeader_Lsr:
		stack, err = ospfLSR(o.Lsr, idx)
	case *opb.OspfHeader_Lsu:
		stack, err = ospfLSU(o.Lsu, idx)
	case *opb.OspfHeader_Lsa:
		stack, err = ospfLSA(o.Lsa, idx)
	default:
		return nil, fmt.Errorf("unrecognized OSPF header type: %v", ospf)
	}
	if err != nil {
		return nil, err
	}
	if rtrID := ospf.GetRouterId(); rtrID != "" {
		ip, isV6 := parseIP(rtrID)
		if ip == nil || isV6 {
			return nil, fmt.Errorf("provided Router ID %q for OSPF header is not a V4 address", rtrID)
		}
		setSingleValue(stack.Ospfv2PacketHeaderRouterID(), ixconfig.String(rtrID))
	}
	if aID := ospf.GetAreaId(); aID != "" {
		ip, isV6 := parseIP(aID)
		if ip == nil || isV6 {
			return nil, fmt.Errorf("provided Area ID %q for OSPF header is not a V4 address", aID)
		}
		setSingleValue(stack.Ospfv2PacketHeaderAreaID(), ixconfig.String(aID))
	}
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func ospfHello(hello *opb.OspfHeader_Hello, idx int) (ospfTrafficTrafficItemConfigElementStack, error) {
	const ipv4BitLen = 8 * net.IPv4len
	stack := ixconfig.NewOspfv2HelloStack(idx)
	if hello.GetNetworkMaskLength() > ipv4BitLen {
		return nil, fmt.Errorf("provided network mask length %d for OSPF Hello exceeds V4 address length", hello.GetNetworkMaskLength())
	}
	mask := net.IPv4bcast.Mask(net.CIDRMask(int(hello.GetNetworkMaskLength()), ipv4BitLen))
	setSingleValue(stack.NetworkMask(), ixconfig.String(mask.String()))
	setSingleValue(stack.HelloInterval(), uintToStr(hello.GetHelloIntervalSec()))
	setSingleValue(stack.RouterPriority(), uintToStr(hello.GetRouterPriority()))
	setSingleValue(stack.RouterDeadInterval(), uintToStr(hello.GetRouterDeadIntervalSec()))
	if designatedRtr := hello.GetDesignatedRouter(); designatedRtr != "" {
		ip, isV6 := parseIP(designatedRtr)
		if ip == nil || isV6 {
			return nil, fmt.Errorf("provided Designated Router %q for OSPF Hello is not a V4 address", designatedRtr)
		}
		setSingleValue(stack.DesignatedRouterID(), ixconfig.String(designatedRtr))
	}
	if backupRtr := hello.GetBackupDesignatedRouter(); backupRtr != "" {
		ip, isV6 := parseIP(backupRtr)
		if ip == nil || isV6 {
			return nil, fmt.Errorf("provided Designated Backup Router %q for OSPF Hello is not a V4 address", backupRtr)
		}
		setSingleValue(stack.BackupDesignatedRouterID(), ixconfig.String(backupRtr))
	}
	if neighbors := hello.GetNeighbors(); neighbors != nil {
		setList(stack.HelloNeighborListNeighborRouterID(), neighbors)
	}

	return stack, nil
}

func ospfDBD(dbd *opb.OspfHeader_DatabaseDescription, idx int) ospfTrafficTrafficItemConfigElementStack {
	stack := ixconfig.NewOspfv2DatabaseDescriptionStack(idx)
	setSingleValue(stack.DatabaseDescriptionBodyInterfaceMTU(), uintToStr(dbd.GetMtu()))
	var flags uint32
	if dbd.GetInitial() {
		flags = flags | 1<<2
	}
	if dbd.GetMore() {
		flags = flags | 1<<1
	}
	if dbd.GetMaster() {
		flags = flags | 1
	}
	setSingleValue(stack.DatabaseDescriptionBodyDatabaseDescriptionFlags(), uintToStr(flags))
	setSingleValue(stack.DatabaseDescriptionBodyDdSequenceNumber(), uintToStr(dbd.GetSeq()))
	return stack
}

var linkStateTypeToNum = map[opb.OspfHeader_LinkStateType]uint32{
	opb.OspfHeader_LINK_STATE_TYPE_ROUTER:          1,
	opb.OspfHeader_LINK_STATE_TYPE_NETWORK:         2,
	opb.OspfHeader_LINK_STATE_TYPE_SUMMARY_NETWORK: 3,
	opb.OspfHeader_LINK_STATE_TYPE_SUMMARY_ASBR:    4,
	opb.OspfHeader_LINK_STATE_TYPE_AS_EXTERNAL:     5,
}

func ospfLSR(lsr *opb.OspfHeader_LinkStateRequest, idx int) (ospfTrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewOspfv2LSARequestStack(idx)
	lst := lsr.GetType()
	typeNum, ok := linkStateTypeToNum[lst]
	if !ok {
		return nil, fmt.Errorf("unrecognized OSPF link state type: %v", lst)
	}
	setSingleValue(stack.LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateType(), uintToStr(typeNum))
	if lsid := lsr.GetLinkStateId(); lsid != "" {
		ip, isV6 := parseIP(lsid)
		if ip == nil || isV6 {
			return nil, fmt.Errorf("provided Link State ID %q for OSPF Link State Request is not a V4 address", lsid)
		}
		setSingleValue(stack.LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateID(), ixconfig.String(lsid))
	}
	if lsar := lsr.GetAdvertisingRouter(); lsar != "" {
		ip, isV6 := parseIP(lsar)
		if ip == nil || isV6 {
			return nil, fmt.Errorf("provided Advertising Router ID %q for OSPF Link State Request is not a V4 address", lsar)
		}
		setSingleValue(stack.LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateAdvertisingRouter(), ixconfig.String(lsar))
	}
	return stack, nil
}

func ospfLSU(lsu *opb.OspfHeader_LinkStateUpdate, idx int) (ospfTrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewOspfv2LSAUpdateStack(idx)
	var ages, types, idRtrs, advRtrs, seqs []string
	for _, adv := range lsu.GetAdvertisements() {
		hdr := adv.GetHeader()
		ages = append(ages, *uintToStr(hdr.GetAgeSeconds()))
		lut := hdr.GetType()
		typeNum, ok := linkStateTypeToNum[lut]
		if !ok {
			return nil, fmt.Errorf("unrecognized OSPF link state update type: %v", lut)
		}
		types = append(types, *uintToStr(typeNum))
		if luid := hdr.GetLinkStateId(); luid != "" {
			ip, isV6 := parseIP(luid)
			if ip == nil || isV6 {
				return nil, fmt.Errorf("provided Link State ID %q for OSPF Link State Update is not a V4 address", luid)
			}
			idRtrs = append(idRtrs, ip.String())
		} else {
			return nil, errors.New("link state ID must be provided for Link State Update advertisement")
		}
		if luar := hdr.GetAdvertisingRouter(); luar != "" {
			ip, isV6 := parseIP(luar)
			if ip == nil || isV6 {
				return nil, fmt.Errorf("provided advertision router %q for OSPF Link State Update is not a V4 address", luar)
			}
			advRtrs = append(advRtrs, ip.String())
		} else {
			return nil, errors.New("advertising router must be provided for Link State Update advertisement")
		}
		seqs = append(seqs, *uintToStr(hdr.GetSeq()))
	}
	setList(stack.LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAge(), ages)
	setList(stack.LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateType(), types)
	setList(stack.LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateID(), idRtrs)
	setList(stack.LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisingRouter(), advRtrs)
	setList(stack.LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateSequenceNumber(), seqs)

	return stack, nil
}

func ospfLSA(lsa *opb.OspfHeader_LinkStateAck, idx int) (ospfTrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewOspfv2LSAAcknowledgementStack(idx)
	var ages, types, idRtrs, advRtrs, seqs []string
	for _, hdr := range lsa.GetHeaders() {
		ages = append(ages, *uintToStr(hdr.GetAgeSeconds()))
		lut := hdr.GetType()
		typeNum, ok := linkStateTypeToNum[lut]
		if !ok {
			return nil, fmt.Errorf("unrecognized OSPF link state ack type: %v", lut)
		}
		types = append(types, *uintToStr(typeNum))
		if luid := hdr.GetLinkStateId(); luid != "" {
			ip, isV6 := parseIP(luid)
			if ip == nil || isV6 {
				return nil, fmt.Errorf("provided Link State ID %q for OSPF Link State Ack is not a V4 address", luid)
			}
			idRtrs = append(idRtrs, ip.String())
		} else {
			return nil, errors.New("link state ID must be provided for Link State Ack advertisement header")
		}
		if luar := hdr.GetAdvertisingRouter(); luar != "" {
			ip, isV6 := parseIP(luar)
			if ip == nil || isV6 {
				return nil, fmt.Errorf("provided advertising router %q for OSPF Link State Ack is not a V4 address", luar)
			}
			advRtrs = append(advRtrs, ip.String())
		} else {
			return nil, errors.New("advertising router must be provided for Link State Ack advertisement header")
		}
		seqs = append(seqs, *uintToStr(hdr.GetSeq()))
	}
	setList(stack.LinkStateAdvertisementHeaderVariableHeaderLinkStateAge(), ages)
	setList(stack.LinkStateAdvertisementHeaderVariableHeaderLinkStateType(), types)
	setList(stack.LinkStateAdvertisementHeaderVariableHeaderLinkStateID(), idRtrs)
	setList(stack.LinkStateAdvertisementHeaderVariableHeaderLinkStateAdvertisingRouter(), advRtrs)
	setList(stack.LinkStateAdvertisementHeaderVariableHeaderLinkStateSequenceNumber(), seqs)
	return stack, nil
}

var rsvpMessageTypeToNum = map[opb.RsvpHeader_MessageType]uint32{
	opb.RsvpHeader_PATH: 1,
	opb.RsvpHeader_RESV: 2,
}

func rsvpStack(rsvp *opb.RsvpHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewRsvpStack(idx)
	setSingleValue(stack.Version(), uintToStr(rsvp.GetVersion()))
	var rrc uint32
	if rsvp.GetRefreshReductionCapable() {
		rrc = 1
	}
	setSingleValue(stack.Flag(), uintToStr(rrc))
	typeNum, ok := rsvpMessageTypeToNum[rsvp.GetMessageType()]
	if !ok {
		return nil, fmt.Errorf("unrecognized RSVP message type: %v", rsvp.GetMessageType())
	}
	setSingleValue(stack.MessegeType(), uintToStr(typeNum))
	setSingleValue(stack.Ttl(), uintToStr(rsvp.GetTtl()))
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

type pimTrafficTrafficItemConfigElementStack interface {
	Version() *ixconfig.TrafficTrafficItemConfigElementStackField
	Type() *ixconfig.TrafficTrafficItemConfigElementStackField
	TrafficTrafficItemConfigElementStack() *ixconfig.TrafficTrafficItemConfigElementStack
}

func pimStack(pim *opb.PimHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	var stack pimTrafficTrafficItemConfigElementStack
	var msgType uint32
	switch pim.Type.(type) {
	case *opb.PimHeader_Hello_:
		stack = ixconfig.NewPimHelloMessageStack(idx)
		msgType = 0
	default:
		return nil, fmt.Errorf("unrecognized PIM header type: %v", pim)
	}
	setSingleValue(stack.Version(), uintToStr(2))
	setSingleValue(stack.Type(), uintToStr(msgType))
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

type ldpTrafficTrafficItemConfigElementStack interface {
	LsrID() *ixconfig.TrafficTrafficItemConfigElementStackField
	LabelSpace() *ixconfig.TrafficTrafficItemConfigElementStackField
	MessageID() *ixconfig.TrafficTrafficItemConfigElementStackField
	TrafficTrafficItemConfigElementStack() *ixconfig.TrafficTrafficItemConfigElementStack
}

func ldpStack(ldp *opb.LdpHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	var stack ldpTrafficTrafficItemConfigElementStack
	switch l := ldp.Type.(type) {
	case *opb.LdpHeader_Hello_:
		var targeted, reqTargeted uint32
		if l.Hello.GetTargeted() {
			targeted = 1
		}
		if l.Hello.GetRequestTargeted() {
			reqTargeted = 1
		}
		s := ixconfig.NewLdpHelloStack(idx)
		setSingleValue(s.CommonHelloParametersTLVHoldTime(), uintToStr(l.Hello.GetHoldTimeSec()))
		setSingleValue(s.CommonHelloParametersTLVTBit(), uintToStr(targeted))
		setSingleValue(s.CommonHelloParametersTLVRBit(), uintToStr(reqTargeted))
		stack = s
	default:
		return nil, fmt.Errorf("unrecognized LDP header type: %v", ldp)
	}
	ip, isV6 := parseIP(ldp.GetLsrId())
	if ip == nil || isV6 {
		return nil, fmt.Errorf("provided LSR ID %q for LDP message is not a V4 address", ldp.GetLsrId())
	}
	setSingleValue(stack.LsrID(), ixconfig.String(ldp.GetLsrId()))
	setSingleValue(stack.LabelSpace(), uintToStr(ldp.GetLabelSpace()))
	setSingleValue(stack.MessageID(), uintToStr(ldp.GetMessageId()))
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func macsecStack(macsec *opb.MacsecHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	return ixconfig.NewMacsecStack(idx).TrafficTrafficItemConfigElementStack(), nil
}

func payloadProtocolTypeStack(idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	return ixconfig.NewPayloadProtocolTypeStack(idx).TrafficTrafficItemConfigElementStack(), nil
}

func espStack(esp *opb.EspHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewIpEncapsulatingSecurityPayloadStack(idx)
	setSingleValue(stack.Spi(), uintToStr(esp.GetSecurityParametersIndex()))
	if esp.GetSequenceNumber() != nil {
		if err := setUintRangeField(stack.Sn(), esp.GetSequenceNumber()); err != nil {
			return nil, fmt.Errorf("could not set sequence numbers: %w", err)
		}
	}
	return stack.TrafficTrafficItemConfigElementStack(), nil
}

func espOverMacsecStack(esp *opb.EspOverMacSecHeader, idx int) (*ixconfig.TrafficTrafficItemConfigElementStack, error) {
	stack := ixconfig.NewIpEspOverMACsecStack(idx)
	setSingleValue(stack.Spi(), uintToStr(esp.GetSecurityParametersIndex()))
	if esp.GetSequenceNumber() != nil {
		if err := setUintRangeField(stack.Sn(), esp.GetSequenceNumber()); err != nil {
			return nil, fmt.Errorf("could not set sequence numbers: %w", err)
		}
	}
	return stack.TrafficTrafficItemConfigElementStack(), nil
}
