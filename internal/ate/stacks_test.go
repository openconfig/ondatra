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
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"google.golang.org/protobuf/proto"

	opb "github.com/openconfig/ondatra/proto"
)

type wantField struct {
	name        string
	wantVal     *string
	wantValList []string
	toField     func(*ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField
}

func TestHeaderStacks(t *testing.T) {
	const (
		srcMacAddr  = "01:02:03:04:05:06"
		dstMacAddr  = "06:05:04:03:02:01"
		vlanID      = 123
		etherType   = 33024
		key         = 1
		seq         = 2
		srcIpv4Addr = "1.2.3.4"
		dstIpv4Addr = "4.3.2.1"
		ttl         = 64
		checksum    = 65376
		srcIpv6Addr = "1:2:3:4:5:6:7:8"
		dstIpv6Addr = "8:7:6:5:4:3:2:1"
		hopLimit    = 64
		label       = 100
		exp         = 1
		srcPort     = 443
		dstPort     = 80
	)
	tests := []struct {
		desc       string
		hdr        *opb.Header
		srcVLAN    bool
		wantFields [][]wantField
	}{{
		desc: "ethernet header",
		hdr: &opb.Header{
			Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr:   &opb.AddressRange{Min: srcMacAddr, Max: srcMacAddr, Count: 1},
					DstAddr:   &opb.AddressRange{Min: dstMacAddr, Max: dstMacAddr, Count: 1},
					EtherType: etherType,
					VlanId:    vlanID,
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "src addr",
			wantVal: ixconfig.String(srcMacAddr),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				e := ixconfig.EthernetStack(*s)
				return (&e).SourceAddress()
			},
		}, {
			name:    "dst addr",
			wantVal: ixconfig.String(dstMacAddr),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				e := ixconfig.EthernetStack(*s)
				return (&e).DestinationAddress()
			},
		}, {
			name:    "ether type",
			wantVal: uintToHexStr(etherType),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				v := ixconfig.EthernetStack(*s)
				return (&v).EtherType()
			},
		}}, {{
			name:    "vlanId",
			wantVal: ixconfig.String(strconv.Itoa(vlanID)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				v := ixconfig.VlanStack(*s)
				return (&v).VlanTagVlanID()
			},
		}}},
	}, {
		desc: "ethernet header w/ src vlan",
		hdr: &opb.Header{
			Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: &opb.AddressRange{Min: srcMacAddr, Max: srcMacAddr, Count: 1},
					DstAddr: &opb.AddressRange{Min: dstMacAddr, Max: dstMacAddr, Count: 1},
				},
			},
		},
		srcVLAN: true,
		wantFields: [][]wantField{{{
			name:    "src addr",
			wantVal: ixconfig.String(srcMacAddr),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				e := ixconfig.EthernetStack(*s)
				return (&e).SourceAddress()
			},
		}, {
			name:    "dst addr",
			wantVal: ixconfig.String(dstMacAddr),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				e := ixconfig.EthernetStack(*s)
				return (&e).DestinationAddress()
			},
		}}, {{
			name: "vlanId",
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				v := ixconfig.VlanStack(*s)
				return (&v).VlanTagVlanID()
			},
		}}},
	}, {
		desc: "gre header",
		hdr: &opb.Header{
			Type: &opb.Header_Gre{
				&opb.GreHeader{
					Key: key,
					Seq: seq,
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "key",
			wantVal: ixconfig.String(strconv.Itoa(key)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				g := ixconfig.GreStack(*s)
				return (&g).KeyHolderKey()
			},
		}, {
			name:    "sequence",
			wantVal: ixconfig.String(strconv.Itoa(seq)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				g := ixconfig.GreStack(*s)
				return (&g).SequenceHolderSequenceNum()
			},
		}}},
	}, {
		desc: "ipv4 header",
		hdr: &opb.Header{
			Type: &opb.Header_Ipv4{
				&opb.Ipv4Header{
					Dscp:           2,
					Ecn:            1,
					Identification: 3,
					DontFragment:   true,
					MoreFragments:  true,
					FragmentOffset: 4,
					Ttl:            proto.Uint32(ttl),
					Protocol: func() *uint32 {
						protocol := uint32(5)
						return &protocol
					}(),
					Checksum: checksum,
					SrcAddr:  &opb.AddressRange{Min: srcIpv4Addr, Max: srcIpv4Addr, Count: 1},
					DstAddr:  &opb.AddressRange{Min: dstIpv4Addr, Max: dstIpv4Addr, Count: 1},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "traffic class",
			wantVal: ixconfig.String("9"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv4Stack(*s)
				return (&ip).PriorityRaw()
			},
		}, {
			name:    "identification",
			wantVal: ixconfig.String("3"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv4Stack(*s)
				return (&ip).Identification()
			},
		}, {
			name:    "don't fragment",
			wantVal: ixconfig.String("1"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv4Stack(*s)
				return (&ip).FlagsFragment()
			},
		}, {
			name:    "more fragments",
			wantVal: ixconfig.String("1"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv4Stack(*s)
				return (&ip).FlagsFragment()
			},
		}, {
			name:    "fragment offset",
			wantVal: ixconfig.String("4"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv4Stack(*s)
				return (&ip).FragmentOffset()
			},
		}, {
			name:    "ttl",
			wantVal: ixconfig.String(strconv.Itoa(ttl)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv4Stack(*s)
				return (&ip).Ttl()
			},
		}, {
			name:    "protocol",
			wantVal: ixconfig.String("5"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv4Stack(*s)
				return (&ip).Protocol()
			},
		}, {
			name:    "checksum",
			wantVal: ixconfig.String("ff60"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv4Stack(*s)
				return (&ip).Checksum()
			},
		}, {
			name:    "src addr",
			wantVal: ixconfig.String(srcIpv4Addr),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv4Stack(*s)
				return (&ip).SrcIp()
			},
		}, {
			name:    "dst addr",
			wantVal: ixconfig.String(dstIpv4Addr),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv4Stack(*s)
				return (&ip).DstIp()
			},
		}}},
	}, {
		desc: "ipv6 header",
		hdr: &opb.Header{
			Type: &opb.Header_Ipv6{
				&opb.Ipv6Header{
					SrcAddr:   &opb.AddressRange{Min: srcIpv6Addr, Max: srcIpv6Addr, Count: 1},
					DstAddr:   &opb.AddressRange{Min: dstIpv6Addr, Max: dstIpv6Addr, Count: 1},
					FlowLabel: &opb.UIntRange{Min: label, Max: label, Count: 1},
					Dscp:      1,
					HopLimit:  proto.Uint32(hopLimit),
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "traffic class",
			wantVal: ixconfig.String("4"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv6Stack(*s)
				return (&ip).VersionTrafficClassFlowLabelTrafficClass()
			},
		}, {
			name:    "flow label",
			wantVal: ixconfig.String(strconv.Itoa(label)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv6Stack(*s)
				return (&ip).VersionTrafficClassFlowLabelFlowLabel()
			},
		}, {
			name:    "hop limit",
			wantVal: ixconfig.String(strconv.Itoa(hopLimit)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv6Stack(*s)
				return (&ip).HopLimit()
			},
		}, {
			name:    "src addr",
			wantVal: ixconfig.String(srcIpv6Addr),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv6Stack(*s)
				return (&ip).SrcIP()
			},
		}, {
			name:    "dst addr",
			wantVal: ixconfig.String(dstIpv6Addr),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				ip := ixconfig.Ipv6Stack(*s)
				return (&ip).DstIP()
			},
		}}},
	}, {
		desc: "mpls header",
		hdr: &opb.Header{
			Type: &opb.Header_Mpls{
				&opb.MplsHeader{
					Label: &opb.UIntRange{Min: label, Max: label, Count: 1},
					Exp:   exp,
					Ttl:   ttl,
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "label",
			wantVal: ixconfig.String(strconv.Itoa(label)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				m := ixconfig.MplsStack(*s)
				return (&m).Value()
			},
		}, {
			name:    "experimental",
			wantVal: ixconfig.String(strconv.Itoa(exp)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				m := ixconfig.MplsStack(*s)
				return (&m).Experimental()
			},
		}, {
			name:    "ttl",
			wantVal: ixconfig.String(strconv.Itoa(ttl)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				m := ixconfig.MplsStack(*s)
				return (&m).Ttl()
			},
		}}},
	}, {
		desc: "tcp header",
		hdr: &opb.Header{
			Type: &opb.Header_Tcp{
				&opb.TcpHeader{
					SrcPort: &opb.UIntRange{Min: srcPort, Max: srcPort, Count: 1},
					DstPort: &opb.UIntRange{Min: dstPort, Max: dstPort, Count: 1},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "src port",
			wantVal: ixconfig.String(strconv.Itoa(srcPort)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				t := ixconfig.TcpStack(*s)
				return (&t).SrcPort()
			},
		}, {
			name:    "dst port",
			wantVal: ixconfig.String(strconv.Itoa(dstPort)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				t := ixconfig.TcpStack(*s)
				return (&t).DstPort()
			},
		}}},
	}, {
		desc: "udp header",
		hdr: &opb.Header{
			Type: &opb.Header_Udp{
				&opb.UdpHeader{
					SrcPort: &opb.UIntRange{Min: srcPort, Max: srcPort, Count: 1},
					DstPort: &opb.UIntRange{Min: dstPort, Max: dstPort, Count: 1},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "src port",
			wantVal: ixconfig.String(strconv.Itoa(srcPort)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				u := ixconfig.UdpStack(*s)
				return (&u).SrcPort()
			},
		}, {
			name:    "dst port",
			wantVal: ixconfig.String(strconv.Itoa(dstPort)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				u := ixconfig.UdpStack(*s)
				return (&u).DstPort()
			},
		}}},
	}, {
		desc: "icmp header, IxNetwork ICMPv1",
		hdr: &opb.Header{
			Type: &opb.Header_Icmp{
				&opb.IcmpHeader{
					Type: &opb.IcmpHeader_RedirectMessage_{
						&opb.IcmpHeader_RedirectMessage{
							Code:   opb.IcmpHeader_RedirectMessage_TOS_AND_NETWORK,
							IpAddr: "192.168.1.2",
						},
					},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "message type",
			wantVal: ixconfig.String(strconv.Itoa(5)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Icmpv1Stack(*s)
				return (&i).MessageType()
			},
		}, {
			name:    "code",
			wantVal: ixconfig.String(strconv.Itoa(2)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Icmpv1Stack(*s)
				return (&i).CodeOptionsRedirectMessageOptions()
			},
		}, {
			name:    "ip address",
			wantVal: ixconfig.String("192.168.1.2"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Icmpv1Stack(*s)
				return (&i).Next4BytesGatewayInternetAddress()
			},
		}}},
	}, {
		desc: "icmp header, IxNetwork ICMPv2",
		hdr: &opb.Header{
			Type: &opb.Header_Icmp{
				&opb.IcmpHeader{
					Type: &opb.IcmpHeader_Timestamp_{
						&opb.IcmpHeader_Timestamp{
							Id:          1,
							Seq:         2,
							OriginateTs: 3,
						},
					},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "message type",
			wantVal: ixconfig.String(strconv.Itoa(13)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Icmpv2Stack(*s)
				return (&i).MessageType()
			},
		}, {
			name:    "id",
			wantVal: ixconfig.String(strconv.Itoa(1)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Icmpv2Stack(*s)
				return (&i).Identifier()
			},
		}, {
			name:    "sequence",
			wantVal: ixconfig.String(strconv.Itoa(2)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Icmpv2Stack(*s)
				return (&i).SequenceNumber()
			},
		}, {
			name:    "originate timestamp",
			wantVal: ixconfig.String(strconv.Itoa(3)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Icmpv2Stack(*s)
				return (&i).NextFieldsFieldsForTimeStampMsgOrigTmpStmp1()
			},
		}}},
	}, {
		desc: "ospf header, Hello message",
		hdr: &opb.Header{
			Type: &opb.Header_Ospf{
				&opb.OspfHeader{
					RouterId: "10.0.0.0",
					AreaId:   "11.0.0.0",
					Type: &opb.OspfHeader_Hello_{
						&opb.OspfHeader_Hello{
							NetworkMaskLength:      16,
							HelloIntervalSec:       10,
							RouterPriority:         0,
							RouterDeadIntervalSec:  40,
							DesignatedRouter:       "1.1.1.1",
							BackupDesignatedRouter: "2.2.2.2",
							Neighbors:              []string{"3.3.3.3", "4.4.4.4"},
						},
					},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "router ID",
			wantVal: ixconfig.String("10.0.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2HelloStack(*s)
				return (&i).Ospfv2PacketHeaderRouterID()
			},
		}, {
			name:    "area ID",
			wantVal: ixconfig.String("11.0.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2HelloStack(*s)
				return (&i).Ospfv2PacketHeaderAreaID()
			},
		}, {
			name:    "mask",
			wantVal: ixconfig.String("255.255.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2HelloStack(*s)
				return (&i).NetworkMask()
			},
		}, {
			name:    "hello interval",
			wantVal: ixconfig.String(strconv.Itoa(10)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2HelloStack(*s)
				return (&i).HelloInterval()
			},
		}, {
			name:    "router priority",
			wantVal: ixconfig.String(strconv.Itoa(0)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2HelloStack(*s)
				return (&i).RouterPriority()
			},
		}, {
			name:    "router dead interval",
			wantVal: ixconfig.String(strconv.Itoa(40)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2HelloStack(*s)
				return (&i).RouterDeadInterval()
			},
		}, {
			name:    "designated router",
			wantVal: ixconfig.String("1.1.1.1"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2HelloStack(*s)
				return (&i).DesignatedRouterID()
			},
		}, {
			name:    "backup designated router",
			wantVal: ixconfig.String("2.2.2.2"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2HelloStack(*s)
				return (&i).BackupDesignatedRouterID()
			},
		}, {
			name:        "neighbors",
			wantValList: []string{"3.3.3.3", "4.4.4.4"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2HelloStack(*s)
				return (&i).HelloNeighborListNeighborRouterID()
			},
		}}},
	}, {
		desc: "ospf header, DBD message",
		hdr: &opb.Header{
			Type: &opb.Header_Ospf{
				&opb.OspfHeader{
					RouterId: "10.0.0.0",
					AreaId:   "11.0.0.0",
					Type: &opb.OspfHeader_Dbd{
						&opb.OspfHeader_DatabaseDescription{
							Mtu:     1500,
							Initial: true,
							More:    true,
							Master:  true,
							Seq:     1,
						},
					},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "router ID",
			wantVal: ixconfig.String("10.0.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2DatabaseDescriptionStack(*s)
				return (&i).Ospfv2PacketHeaderRouterID()
			},
		}, {
			name:    "area ID",
			wantVal: ixconfig.String("11.0.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2DatabaseDescriptionStack(*s)
				return (&i).Ospfv2PacketHeaderAreaID()
			},
		}, {
			name:    "mtu",
			wantVal: ixconfig.String(strconv.Itoa(1500)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2DatabaseDescriptionStack(*s)
				return (&i).DatabaseDescriptionBodyInterfaceMTU()
			},
		}, {
			name:    "flags",
			wantVal: ixconfig.String("7"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2DatabaseDescriptionStack(*s)
				return (&i).DatabaseDescriptionBodyDatabaseDescriptionFlags()
			},
		}, {
			name:    "sequence",
			wantVal: ixconfig.String(strconv.Itoa(1)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2DatabaseDescriptionStack(*s)
				return (&i).DatabaseDescriptionBodyDdSequenceNumber()
			},
		}}},
	}, {
		desc: "ospf header, link state request message",
		hdr: &opb.Header{
			Type: &opb.Header_Ospf{
				&opb.OspfHeader{
					RouterId: "10.0.0.0",
					AreaId:   "11.0.0.0",
					Type: &opb.OspfHeader_Lsr{
						&opb.OspfHeader_LinkStateRequest{
							Type:              opb.OspfHeader_LINK_STATE_TYPE_NETWORK,
							LinkStateId:       "1.1.1.1",
							AdvertisingRouter: "2.2.2.2",
						},
					},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "router ID",
			wantVal: ixconfig.String("10.0.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSARequestStack(*s)
				return (&i).Ospfv2PacketHeaderRouterID()
			},
		}, {
			name:    "area ID",
			wantVal: ixconfig.String("11.0.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSARequestStack(*s)
				return (&i).Ospfv2PacketHeaderAreaID()
			},
		}, {
			name:    "link state type",
			wantVal: ixconfig.String(strconv.Itoa(2)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSARequestStack(*s)
				return (&i).LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateType()
			},
		}, {
			name:    "link state ID",
			wantVal: ixconfig.String("1.1.1.1"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSARequestStack(*s)
				return (&i).LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateID()
			},
		}, {
			name:    "advertising router",
			wantVal: ixconfig.String("2.2.2.2"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSARequestStack(*s)
				return (&i).LinkStateRequestBodyRequestedLSAsListRequestedLSADescriptionLinkStateAdvertisingRouter()
			},
		}}},
	}, {
		desc: "ospf header, link state update message",
		hdr: &opb.Header{
			Type: &opb.Header_Ospf{
				&opb.OspfHeader{
					RouterId: "10.0.0.0",
					AreaId:   "11.0.0.0",
					Type: &opb.OspfHeader_Lsu{
						&opb.OspfHeader_LinkStateUpdate{
							Advertisements: []*opb.OspfHeader_LinkStateUpdate_Advertisement{{
								Header: &opb.OspfHeader_LinkStateAdvertisementHeader{
									AgeSeconds:        1,
									Type:              opb.OspfHeader_LINK_STATE_TYPE_ROUTER,
									LinkStateId:       "1.1.1.1",
									AdvertisingRouter: "2.2.2.2",
									Seq:               2,
								},
							}, {
								Header: &opb.OspfHeader_LinkStateAdvertisementHeader{
									AgeSeconds:        10,
									Type:              opb.OspfHeader_LINK_STATE_TYPE_NETWORK,
									LinkStateId:       "3.3.3.3",
									AdvertisingRouter: "4.4.4.4",
									Seq:               20,
								},
							}},
						},
					},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "router ID",
			wantVal: ixconfig.String("10.0.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAUpdateStack(*s)
				return (&i).Ospfv2PacketHeaderRouterID()
			},
		}, {
			name:    "area ID",
			wantVal: ixconfig.String("11.0.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAUpdateStack(*s)
				return (&i).Ospfv2PacketHeaderAreaID()
			},
		}, {
			name:        "ages",
			wantValList: []string{"1", "10"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAUpdateStack(*s)
				return (&i).LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAge()
			},
		}, {
			name:        "types",
			wantValList: []string{"1", "2"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAUpdateStack(*s)
				return (&i).LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateType()
			},
		}, {
			name:        "ids",
			wantValList: []string{"1.1.1.1", "3.3.3.3"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAUpdateStack(*s)
				return (&i).LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateID()
			},
		}, {
			name:        "advertising routers",
			wantValList: []string{"2.2.2.2", "4.4.4.4"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAUpdateStack(*s)
				return (&i).LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateAdvertisingRouter()
			},
		}, {
			name:        "sequence numbers",
			wantValList: []string{"2", "20"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAUpdateStack(*s)
				return (&i).LinkStateUpdateBodyLsaListLinkStateAdvertisementVariableHeaderLinkStateSequenceNumber()
			},
		}}},
	}, {
		desc: "ospf header, link state ack message",
		hdr: &opb.Header{
			Type: &opb.Header_Ospf{
				&opb.OspfHeader{
					RouterId: "10.0.0.0",
					AreaId:   "11.0.0.0",
					Type: &opb.OspfHeader_Lsa{
						&opb.OspfHeader_LinkStateAck{
							Headers: []*opb.OspfHeader_LinkStateAdvertisementHeader{{
								AgeSeconds:        1,
								Type:              opb.OspfHeader_LINK_STATE_TYPE_ROUTER,
								LinkStateId:       "1.1.1.1",
								AdvertisingRouter: "2.2.2.2",
								Seq:               2,
							}, {
								AgeSeconds:        10,
								Type:              opb.OspfHeader_LINK_STATE_TYPE_NETWORK,
								LinkStateId:       "3.3.3.3",
								AdvertisingRouter: "4.4.4.4",
								Seq:               20,
							}},
						},
					},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "router ID",
			wantVal: ixconfig.String("10.0.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAAcknowledgementStack(*s)
				return (&i).Ospfv2PacketHeaderRouterID()
			},
		}, {
			name:    "area ID",
			wantVal: ixconfig.String("11.0.0.0"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAAcknowledgementStack(*s)
				return (&i).Ospfv2PacketHeaderAreaID()
			},
		}, {
			name:        "ages",
			wantValList: []string{"1", "10"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAAcknowledgementStack(*s)
				return (&i).LinkStateAdvertisementHeaderVariableHeaderLinkStateAge()
			},
		}, {
			name:        "types",
			wantValList: []string{"1", "2"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAAcknowledgementStack(*s)
				return (&i).LinkStateAdvertisementHeaderVariableHeaderLinkStateType()
			},
		}, {
			name:        "ids",
			wantValList: []string{"1.1.1.1", "3.3.3.3"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAAcknowledgementStack(*s)
				return (&i).LinkStateAdvertisementHeaderVariableHeaderLinkStateID()
			},
		}, {
			name:        "advertising routers",
			wantValList: []string{"2.2.2.2", "4.4.4.4"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAAcknowledgementStack(*s)
				return (&i).LinkStateAdvertisementHeaderVariableHeaderLinkStateAdvertisingRouter()
			},
		}, {
			name:        "sequence numbers",
			wantValList: []string{"2", "20"},
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.Ospfv2LSAAcknowledgementStack(*s)
				return (&i).LinkStateAdvertisementHeaderVariableHeaderLinkStateSequenceNumber()
			},
		}}},
	}, {
		desc: "rsvp header",
		hdr: &opb.Header{
			Type: &opb.Header_Rsvp{
				&opb.RsvpHeader{
					Version:                 0,
					RefreshReductionCapable: true,
					MessageType:             opb.RsvpHeader_PATH,
					Ttl:                     60,
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "version",
			wantVal: ixconfig.String(strconv.Itoa(0)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.RsvpStack(*s)
				return (&i).Version()
			},
		}, {
			name:    "refresh reduction capable",
			wantVal: ixconfig.String(strconv.Itoa(1)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.RsvpStack(*s)
				return (&i).Flag()
			},
		}, {
			name:    "message type",
			wantVal: ixconfig.String(strconv.Itoa(1)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.RsvpStack(*s)
				return (&i).MessegeType()
			},
		}, {
			name:    "ttl",
			wantVal: ixconfig.String(strconv.Itoa(60)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.RsvpStack(*s)
				return (&i).Ttl()
			},
		}}},
	}, {
		desc: "PIM header, Hello message",
		hdr: &opb.Header{
			Type: &opb.Header_Pim{
				&opb.PimHeader{
					Type: &opb.PimHeader_Hello_{
						&opb.PimHeader_Hello{},
					},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "version",
			wantVal: ixconfig.String(strconv.Itoa(2)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.PimHelloMessageStack(*s)
				return (&i).Version()
			},
		}, {
			name:    "message type",
			wantVal: ixconfig.String(strconv.Itoa(0)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.PimHelloMessageStack(*s)
				return (&i).Type()
			},
		}}},
	}, {
		desc: "LDP header, Hello message",
		hdr: &opb.Header{
			Type: &opb.Header_Ldp{
				&opb.LdpHeader{
					LsrId:      "1.1.1.1",
					LabelSpace: 2,
					MessageId:  3,
					Type: &opb.LdpHeader_Hello_{
						&opb.LdpHeader_Hello{
							HoldTimeSec:     15,
							Targeted:        true,
							RequestTargeted: true,
						},
					},
				},
			},
		},
		wantFields: [][]wantField{{{
			name:    "lsr ID",
			wantVal: ixconfig.String("1.1.1.1"),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.LdpHelloStack(*s)
				return (&i).LsrID()
			},
		}, {
			name:    "label space",
			wantVal: ixconfig.String(strconv.Itoa(2)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.LdpHelloStack(*s)
				return (&i).LabelSpace()
			},
		}, {
			name:    "message ID",
			wantVal: ixconfig.String(strconv.Itoa(3)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.LdpHelloStack(*s)
				return (&i).MessageID()
			},
		}, {
			name:    "hold time secs",
			wantVal: ixconfig.String(strconv.Itoa(15)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.LdpHelloStack(*s)
				return (&i).CommonHelloParametersTLVHoldTime()
			},
		}, {
			name:    "targeted",
			wantVal: ixconfig.String(strconv.Itoa(1)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.LdpHelloStack(*s)
				return (&i).CommonHelloParametersTLVTBit()
			},
		}, {
			name:    "request targeted",
			wantVal: ixconfig.String(strconv.Itoa(1)),
			toField: func(s *ixconfig.TrafficTrafficItemConfigElementStack) *ixconfig.TrafficTrafficItemConfigElementStackField {
				i := ixconfig.LdpHelloStack(*s)
				return (&i).CommonHelloParametersTLVRBit()
			},
		}}},
	}, {
		desc: "macsec header",
		hdr: &opb.Header{
			Type: &opb.Header_Macsec{
				&opb.MacsecHeader{},
			},
		},
		wantFields: [][]wantField{{}, {}},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			stacks, err := headerStacks(test.hdr, 0, test.srcVLAN)
			if err != nil {
				t.Fatalf("headerStacks(%v): unexpected error %v", test.hdr, err)
			}
			if len(stacks) != len(test.wantFields) {
				t.Fatalf("headerStacks(%v): unexpected stack count %d, wanted %d", test.hdr, len(stacks), len(test.wantFields))
			}
			for i, s := range stacks {
				wantVals, gotVals := make(map[string]any), make(map[string]any)
				for _, wf := range test.wantFields[i] {
					field := wf.toField(s)
					if wf.wantVal != nil {
						if field.ValueType == nil || *(field.ValueType) != "singleValue" {
							t.Fatalf("headerStacks(%v): invalid value type %v for field %q, wanted 'singleValue'", test.hdr, field.ValueType, wf.name)
						}
						wantVals[wf.name] = wf.wantVal
						gotVals[wf.name] = field.SingleValue
					}

					if wf.wantValList != nil {
						if *(field.ValueType) != "valueList" {
							t.Fatalf("headerStacks(%v): invalid value type %q for field %q, wanted 'valueList'", test.hdr, *(field.ValueType), wf.name)
						}
						wantVals[wf.name] = wf.wantValList
						gotVals[wf.name] = field.ValueList
					}
				}
				if diff := cmp.Diff(wantVals, gotVals); diff != "" {
					t.Fatalf("headerStacks(%v): unexpected diff for stack: %s", test.hdr, diff)
				}
			}
		})
	}
}

func TestHeaderStacksError(t *testing.T) {
	tests := []struct {
		desc    string
		hdr     *opb.Header
		idx     int
		wantErr string
	}{{
		desc: "bad crc not in initial ethernet header",
		hdr: &opb.Header{
			Type: &opb.Header_Eth{
				&opb.EthernetHeader{BadCrc: true},
			},
		},
		idx:     1,
		wantErr: "bad CRC",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			stacks, err := headerStacks(test.hdr, test.idx, false)
			if err == nil {
				t.Fatalf("headerStacks(%v): unexpectedly succeeded %v", test.hdr, stacks)
			}
			if !strings.Contains(err.Error(), test.wantErr) {
				t.Fatalf("headerStacks(%v): got err %s, want %s", test.hdr, err, test.wantErr)
			}
		})
	}
}
