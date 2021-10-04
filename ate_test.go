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
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/negtest"

	opb "github.com/openconfig/ondatra/proto"
)

var got struct {
	ate   string
	ifs   []*opb.InterfaceConfig
	flows []*opb.Flow
}

func initATEFakes(t *testing.T) {
	t.Helper()
	initFakeBinding(t)
	fakeBind.TopologyPusher = func(ate *reservation.ATE, top *opb.Topology) error {
		got.ate = ate.Name
		got.ifs = top.GetInterfaces()
		return nil
	}
	fakeBind.TrafficStarter = func(ate *reservation.ATE, flows []*opb.Flow) error {
		got.ate = ate.Name
		got.flows = flows
		return nil
	}
	reserveFakeTestbed(t)
}

func TestPushTopology(t *testing.T) {
	initATEFakes(t)
	ate := ATE(t, "ate")
	p1 := ate.Port(t, "port1")
	top := ate.Topology().New()
	intf := top.AddInterface("intf").WithPort(p1)
	intf.IPv4().
		WithAddress("192.168.1.1/30").
		WithDefaultGateway("192.168.1.2")
	net4 := intf.AddNetwork("net4")
	net4.IPv4().
		WithAddress("30.0.0.0/28").
		WithCount(10)
	net4.BGP().
		WithActive(false).
		WithNextHopAddress("1.0.0.0").
		WithOriginIGP().
		WithLocalPreference(1).
		WithCommunities("1:1")
	net4.BGP().AddExtendedCommunityColor().
		WithCOBits10().
		WithReservedBits(1).
		WithValue(11)
	net4.BGP().WithASNSetModeSET()
	net4.BGP().AddASPathSegment(1, 2, 3).
		WithTypeSET()
	intf.IPv6().
		WithAddress("2001::4860:193:168:22:1/127").
		WithDefaultGateway("2001::4860:193:168:22:0")
	net6 := intf.AddNetwork("net6")
	net6.IPv6().
		WithAddress("3000::/126").
		WithCount(20)
	intf.ISIS().
		WithLevelL1().
		WithNetworkTypeBroadcast().
		WithWideMetricEnabled(true).
		WithMetric(10).
		WithAreaID("areaId")
	intf.BGP().AddPeer().
		WithPeerAddress("192.168.1.1").
		WithLocalASN(100)
	intf.BGP().AddPeer().
		WithActive(false).
		WithTypeInternal().
		WithPeerAddress("2001::4860:193:168:22:1").
		WithLocalASN(200000).
		WithHoldTime(300).
		WithKeepaliveTime(100).
		WithMD5Key("md5Key").
		Capabilities().WithRouteRefreshEnabled(false)
	top.Push(t)

	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	want := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Ipv4:     &opb.IpConfig{AddressCidr: "192.168.1.1/30", DefaultGateway: "192.168.1.2"},
		Ipv6:     &opb.IpConfig{AddressCidr: "2001::4860:193:168:22:1/127", DefaultGateway: "2001::4860:193:168:22:0"},
		Isis: &opb.ISISConfig{
			Level:  opb.ISISConfig_L1,
			Metric: 10, AreaId: "areaId",
			EnableWideMetric: true,
			NetworkType:      opb.ISISConfig_BROADCAST,
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			PeerAddress:      "192.168.1.1",
			LocalAsn:         100,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
		}, {
			Active:           false,
			Type:             opb.BgpPeer_TYPE_INTERNAL,
			PeerAddress:      "2001::4860:193:168:22:1",
			LocalAsn:         200000,
			HoldTimeSec:      300,
			KeepaliveTimeSec: 100,
			Md5Key:           "md5Key",
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  false,
			},
		}}},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net4",
				InterfaceName: "intf",
				Ipv4:          &opb.NetworkIp{AddressCidr: "30.0.0.0/28", Count: 10},
				BgpAttributes: &opb.BgpAttributes{
					Active:          false,
					NextHopAddress:  "1.0.0.0",
					Origin:          opb.BgpAttributes_ORIGIN_IGP,
					LocalPreference: 1,
					Communities:     []string{"1:1"},
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits:       opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_10,
								ReservedBits: 1,
								Value:        11,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SET,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SET,
						Asns: []uint32{1, 2, 3},
					}},
				},
			},
			&opb.Network{
				Name:          "net6",
				InterfaceName: "intf",
				Ipv6:          &opb.NetworkIp{AddressCidr: "3000::/126", Count: 20},
			},
		},
	}
	if diff := cmp.Diff(got.ifs[0], want, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}
}

func TestClearAndPushTopology(t *testing.T) {
	initATEFakes(t)
	wantOriginal := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:     "490001",
			TeRouterId: "0.0.0.0",
			IsReachability: []*opb.ISReachability{{
				Nodes: []*opb.ISReachability_Node{{
					IngressMetric: 10,
					EgressMetric:  10,
					SegmentRouting: &opb.ISISSegmentRouting{
						SrgbRange: []*opb.ISISSegmentRouting_SIDRange{{}},
						SrlbRange: []*opb.ISISSegmentRouting_SIDRange{{}},
					},
				}},
			}},
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
			SrtePolicyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
				Count:         1,
				Active:        true,
				Distinguisher: 1,
				PolicyColor: &opb.UInt32IncRange{
					Start: 100,
					Step:  0,
				},
				Communities: []*opb.BgpPeer_SrtePolicyGroup_Community{{
					Type:    opb.BgpPeer_SrtePolicyGroup_Community_TYPE_NO_ADVERTISE,
					Asn:     0,
					Pattern: 0,
				}},
				SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{
					Active: true,
					Segments: []*opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment{{
						Active: true,
					}},
				}},
			}},
		}}},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
						Asns: []uint32{1},
					}},
				},
			},
		},
	}
	wantClearSRLBRanges := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:     "490001",
			TeRouterId: "0.0.0.0",
			IsReachability: []*opb.ISReachability{{
				Nodes: []*opb.ISReachability_Node{{
					IngressMetric: 10,
					EgressMetric:  10,
					SegmentRouting: &opb.ISISSegmentRouting{
						SrgbRange: []*opb.ISISSegmentRouting_SIDRange{{}},
					},
				}},
			}},
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
			SrtePolicyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
				Count:         1,
				Active:        true,
				Distinguisher: 1,
				PolicyColor: &opb.UInt32IncRange{
					Start: 100,
					Step:  0,
				},
				Communities: []*opb.BgpPeer_SrtePolicyGroup_Community{{
					Type:    opb.BgpPeer_SrtePolicyGroup_Community_TYPE_NO_ADVERTISE,
					Asn:     0,
					Pattern: 0,
				}},
				SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{
					Active: true,
					Segments: []*opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment{{
						Active: true,
					}},
				}},
			}},
		}}},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
						Asns: []uint32{1},
					}},
				},
			},
		},
	}
	wantClearSRGBRanges := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:     "490001",
			TeRouterId: "0.0.0.0",
			IsReachability: []*opb.ISReachability{{
				Nodes: []*opb.ISReachability_Node{{
					IngressMetric:  10,
					EgressMetric:   10,
					SegmentRouting: &opb.ISISSegmentRouting{},
				}},
			}},
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
			SrtePolicyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
				Count:         1,
				Active:        true,
				Distinguisher: 1,
				PolicyColor: &opb.UInt32IncRange{
					Start: 100,
					Step:  0,
				},
				Communities: []*opb.BgpPeer_SrtePolicyGroup_Community{{
					Type:    opb.BgpPeer_SrtePolicyGroup_Community_TYPE_NO_ADVERTISE,
					Asn:     0,
					Pattern: 0,
				}},
				SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{
					Active: true,
					Segments: []*opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment{{
						Active: true,
					}},
				}},
			}},
		}}},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
						Asns: []uint32{1},
					}},
				},
			},
		},
	}
	wantClearISISNodes := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			IsReachability:   []*opb.ISReachability{{}},
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
			SrtePolicyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
				Count:         1,
				Active:        true,
				Distinguisher: 1,
				PolicyColor: &opb.UInt32IncRange{
					Start: 100,
					Step:  0,
				},
				Communities: []*opb.BgpPeer_SrtePolicyGroup_Community{{
					Type:    opb.BgpPeer_SrtePolicyGroup_Community_TYPE_NO_ADVERTISE,
					Asn:     0,
					Pattern: 0,
				}},
				SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{
					Active: true,
					Segments: []*opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment{{
						Active: true,
					}},
				}},
			}},
		}}},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
						Asns: []uint32{1},
					}},
				},
			},
		},
	}
	wantClearISReachabilities := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
			SrtePolicyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
				Count:         1,
				Active:        true,
				Distinguisher: 1,
				PolicyColor: &opb.UInt32IncRange{
					Start: 100,
					Step:  0,
				},
				Communities: []*opb.BgpPeer_SrtePolicyGroup_Community{{
					Type:    opb.BgpPeer_SrtePolicyGroup_Community_TYPE_NO_ADVERTISE,
					Asn:     0,
					Pattern: 0,
				}},
				SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{
					Active: true,
					Segments: []*opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment{{
						Active: true,
					}},
				}},
			}},
		}}},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
						Asns: []uint32{1},
					}},
				},
			},
		},
	}
	wantClearSegments := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
			SrtePolicyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
				Count:         1,
				Active:        true,
				Distinguisher: 1,
				PolicyColor: &opb.UInt32IncRange{
					Start: 100,
					Step:  0,
				},
				Communities: []*opb.BgpPeer_SrtePolicyGroup_Community{{
					Type:    opb.BgpPeer_SrtePolicyGroup_Community_TYPE_NO_ADVERTISE,
					Asn:     0,
					Pattern: 0,
				}},
				SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{
					Active: true,
				}},
			}},
		}}},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
						Asns: []uint32{1},
					}},
				},
			},
		},
	}
	wantClearSegmentLists := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
			SrtePolicyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
				Count:         1,
				Active:        true,
				Distinguisher: 1,
				PolicyColor: &opb.UInt32IncRange{
					Start: 100,
					Step:  0,
				},
				Communities: []*opb.BgpPeer_SrtePolicyGroup_Community{{
					Type:    opb.BgpPeer_SrtePolicyGroup_Community_TYPE_NO_ADVERTISE,
					Asn:     0,
					Pattern: 0,
				}},
			}},
		}}},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
						Asns: []uint32{1},
					}},
				},
			},
		},
	}
	wantClearCommunities := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
			SrtePolicyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
				Count:         1,
				Active:        true,
				Distinguisher: 1,
				PolicyColor: &opb.UInt32IncRange{
					Start: 100,
					Step:  0,
				},
			}},
		}}},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
						Asns: []uint32{1},
					}},
				},
			},
		},
	}
	wantClearSRTEPolicyGroups := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
		}}},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
						Asns: []uint32{1},
					}},
				},
			},
		},
	}
	wantClearPeers := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
					AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
						Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SEQ,
						Asns: []uint32{1},
					}},
				},
			},
		},
	}
	wantClearASPathSegments := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active: true,
					Origin: opb.BgpAttributes_ORIGIN_IGP,
					ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
						Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
							Color: &opb.BgpAttributes_ExtendedCommunity_Color{
								CoBits: opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_00,
							},
						},
					}},
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
				},
			},
		},
	}
	wantClearExtendedCommunityColors := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{},
		Networks: []*opb.Network{
			&opb.Network{
				Name:          "net",
				InterfaceName: "intf",
				BgpAttributes: &opb.BgpAttributes{
					Active:     true,
					Origin:     opb.BgpAttributes_ORIGIN_IGP,
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
				},
			},
		},
	}
	wantClearNetworks := &opb.InterfaceConfig{
		Name:     "intf",
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Isis: &opb.ISISConfig{
			AreaId:           "490001",
			TeRouterId:       "0.0.0.0",
			HelloIntervalSec: 10,
			DeadIntervalSec:  30,
		},
		Bgp: &opb.BgpConfig{},
	}

	ate := ATE(t, "ate")
	port := ate.Port(t, "port1")
	top := ate.Topology().New()
	intf := top.AddInterface("intf").WithPort(port)
	net := intf.AddNetwork("net")
	net.BGP().AddExtendedCommunityColor()
	net.BGP().AddASPathSegment(1)
	peer := intf.BGP().AddPeer()
	policies := peer.AddSRTEPolicyGroup()
	policies.AddCommunity()
	segmentList := policies.AddSegmentList()
	segmentList.AddSegment()
	isr := intf.ISIS().AddISReachability()
	node := isr.AddISISNode()
	node.SegmentRouting().AddSRGBRange()
	node.SegmentRouting().AddSRLBRange()
	top.Push(t)

	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantOriginal, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	node.SegmentRouting().ClearSRLBRanges()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearSRLBRanges, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	node.SegmentRouting().ClearSRGBRanges()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearSRGBRanges, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	isr.ClearISISNodes()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearISISNodes, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	intf.ISIS().ClearISReachabilities()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearISReachabilities, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	segmentList.ClearSegments()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearSegments, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	policies.ClearSegmentLists()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearSegmentLists, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	policies.ClearCommunities()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearCommunities, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	peer.ClearSRTEPolicyGroups()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearSRTEPolicyGroups, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	intf.BGP().ClearPeers()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearPeers, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	net.BGP().ClearASPathSegments()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearASPathSegments, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	net.BGP().ClearExtendedCommunityColors()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearExtendedCommunityColors, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	intf.ClearNetworks()
	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantClearNetworks, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	top.ClearInterfaces()
	got := negtest.ExpectFatal(t, func(t testing.TB) {
		top.Push(t)
	})
	if want := "zero interfaces"; !strings.Contains(strings.ToLower(got), want) {
		t.Errorf("Configure failed with message %q, want %q", got, want)
	}
}

func TestPushTopologyErrors(t *testing.T) {
	initATEFakes(t)
	ate := ATE(t, "ate")
	p1 := ate.Port(t, "port1")
	p2 := ate.Port(t, "port2")

	tests := []struct {
		name         string
		top          *ATETopology
		wantFatalMsg string
	}{{
		name:         "zero interfaces",
		top:          ate.Topology().New(),
		wantFatalMsg: "zero",
	}, {
		name: "duplicate interface name",
		top: func() *ATETopology {
			top := ate.Topology().New()
			top.AddInterface("intf").WithPort(p1)
			top.AddInterface("intf").WithPort(p2)
			return top
		}(),
		wantFatalMsg: "duplicate interface name",
	}, {
		name: "duplicate network name",
		top: func() *ATETopology {
			top := ate.Topology().New()
			intf := top.AddInterface("intf1").WithPort(p1)
			intf.AddNetwork("net")
			intf.AddNetwork("net")
			return top
		}(),
		wantFatalMsg: "duplicate network name",
	}, {
		name: "no link",
		top: func() *ATETopology {
			top := ate.Topology().New()
			top.AddInterface("i")
			return top
		}(),
		wantFatalMsg: "no port or lag",
	}, {
		name: "lag and lacp",
		top: func() *ATETopology {
			top := ate.Topology().New()
			top.AddInterface("i").WithLAG(top.AddLAG().WithPorts(p1)).
				WithLACPEnabled(true)
			return top
		}(),
		wantFatalMsg: "both a lag and that lacp",
	}, {
		name: "bad address format",
		top: func() *ATETopology {
			top := ate.Topology().New()
			top.AddInterface("i").WithPort(p1).
				IPv4().WithAddress("blah")
			return top
		}(),
		wantFatalMsg: "not valid cidr notation",
	}, {
		name: "bad gateway format",
		top: func() *ATETopology {
			top := ate.Topology().New()
			top.AddInterface("i").WithPort(p1).
				IPv4().WithAddress("1.2.3.4/24").WithDefaultGateway("blah")
			return top
		}(),
		wantFatalMsg: "not valid ip notation",
	}, {
		name: "gateway not in cidr range",
		top: func() *ATETopology {
			top := ate.Topology().New()
			top.AddInterface("i").WithPort(p1).
				IPv4().WithAddress("1.2.3.4/24").WithDefaultGateway("2.3.4.5")
			return top
		}(),
		wantFatalMsg: "not in cidr range",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := negtest.ExpectFatal(t, func(t testing.TB) {
				tt.top.Push(t)
			})
			if !strings.Contains(strings.ToLower(got), tt.wantFatalMsg) {
				t.Errorf("Configure failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}
}

func TestUpdateTopology(t *testing.T) {
	initATEFakes(t)
	ate := ATE(t, "ate")
	p1 := ate.Port(t, "port1")
	top := ate.Topology().New()
	intf := top.AddInterface("intf").WithPort(p1)
	peer := intf.BGP().AddPeer().WithActive(true)

	wantPush := &opb.InterfaceConfig{
		Name:     intf.pb.GetName(),
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           true,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
		}}},
	}

	wantUpdate := &opb.InterfaceConfig{
		Name:     intf.pb.GetName(),
		Link:     &opb.InterfaceConfig_Port{"1/1"},
		Ethernet: &opb.EthernetConfig{Mtu: 1500},
		Bgp: &opb.BgpConfig{BgpPeers: []*opb.BgpPeer{{
			Active:           false,
			Type:             opb.BgpPeer_TYPE_EXTERNAL,
			HoldTimeSec:      90,
			KeepaliveTimeSec: 30,
			Capabilities: &opb.BgpPeer_Capabilities{
				Ipv4Unicast:   true,
				Ipv4Multicast: true,
				Ipv4MplsVpn:   true,
				Ipv6Unicast:   true,
				Ipv6Multicast: true,
				Ipv6MplsVpn:   true,
				Vpls:          true,
				RouteRefresh:  true,
			},
		}}},
	}

	top.Push(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("Push(): unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("Push(): expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantPush, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("Push(): did not get expected interface, diff(-got,+want):\n%s", diff)
	}

	peer.WithActive(false)
	top.Update(t)
	if want := "ix1"; got.ate != want {
		t.Errorf("Update(): unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if len(got.ifs) != 1 {
		t.Errorf("Update(): expected 1 interface, got: %v", len(got.ifs))
	}
	if diff := cmp.Diff(got.ifs[0], wantUpdate, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("Update(): did not get expected interface, diff(-got,+want):\n%s", diff)
	}
}

func TestStartTraffic(t *testing.T) {
	initATEFakes(t)
	ate := ATE(t, "ate")
	p1 := ate.Port(t, "port1")
	p2 := ate.Port(t, "port2")
	top := ate.Topology().New()
	intf1 := top.AddInterface("intf1").WithPort(p1)
	intf2 := top.AddInterface("intf2").WithPort(p2)
	srcEPs := []*opb.Flow_Endpoint{{InterfaceName: intf1.pb.GetName()}}
	dstEPs := []*opb.Flow_Endpoint{{InterfaceName: intf2.pb.GetName()}}

	tests := []struct {
		desc     string
		flow     *Flow
		wantFlow *opb.Flow
	}{{
		desc: "create ipv4 flow",
		flow: ate.Traffic().NewFlow("ipv4 flow").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithHeaders(
				NewEthernetHeader(),
				func() *IPv4Header {
					h := NewIPv4Header()
					h.SrcAddressRange().WithMin("10.0.0.0").WithCount(5)
					h.DstAddressRange().WithMin("10.1.0.0").WithMax("10.1.0.255").WithCount(10).WithRandom()
					return h
				}()),
		wantFlow: &opb.Flow{
			Name:         "ipv4 flow",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			Headers: []*opb.Header{
				{Type: &opb.Header_Eth{&opb.EthernetHeader{}}},
				{Type: &opb.Header_Ipv4{&opb.Ipv4Header{
					SrcAddr: &opb.AddressRange{
						Min:    "10.0.0.0",
						Max:    "255.255.255.254",
						Count:  5,
						Random: false,
					},
					DstAddr: &opb.AddressRange{
						Min:    "10.1.0.0",
						Max:    "10.1.0.255",
						Count:  10,
						Random: true,
					},
					Ttl: 64,
				}}},
			},
		},
	}, {
		desc: "create ipv6 flow",
		flow: ate.Traffic().NewFlow("ipv6 flow").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithHeaders(
				NewEthernetHeader(),
				func() *IPv6Header {
					h := NewIPv6Header()
					h.SrcAddressRange().WithMin("2001:db8:3c4d::").WithCount(5)
					h.DstAddressRange().WithMin("2001:db9:3c4d::").WithMax("2001:db9:3c4d:ffff:ffff:ffff:ffff:ffff").WithCount(10).WithRandom()
					return h
				}()),
		wantFlow: &opb.Flow{
			Name:         "ipv6 flow",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			Headers: []*opb.Header{
				{Type: &opb.Header_Eth{&opb.EthernetHeader{}}},
				{Type: &opb.Header_Ipv6{&opb.Ipv6Header{
					SrcAddr: &opb.AddressRange{
						Min:    "2001:db8:3c4d::",
						Max:    "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff",
						Count:  5,
						Random: false,
					},
					DstAddr: &opb.AddressRange{
						Min:    "2001:db9:3c4d::",
						Max:    "2001:db9:3c4d:ffff:ffff:ffff:ffff:ffff",
						Count:  10,
						Random: true,
					},
					HopLimit: 64,
				}}},
			},
		},
	}, {
		desc: "create tcp flow",
		flow: ate.Traffic().NewFlow("tcp flow").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithHeaders(
				NewEthernetHeader(),
				NewIPv4Header(),
				func() *TCPHeader {
					h := NewTCPHeader()
					h.SrcPortRange().WithMin(1).WithCount(5)
					h.DstPortRange().WithMin(3000).WithMax(4000).WithCount(20).WithRandom()
					return h
				}()),
		wantFlow: &opb.Flow{
			Name:         "tcp flow",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			Headers: []*opb.Header{
				{Type: &opb.Header_Eth{&opb.EthernetHeader{}}},
				{Type: &opb.Header_Ipv4{&opb.Ipv4Header{Ttl: 64}}},
				{Type: &opb.Header_Tcp{&opb.TcpHeader{
					SrcPort: &opb.UIntRange{
						Min:    1,
						Max:    65535,
						Count:  5,
						Random: false,
					},
					DstPort: &opb.UIntRange{
						Min:    3000,
						Max:    4000,
						Count:  20,
						Random: true,
					},
				}}},
			},
		},
	}, {
		desc: "create mpls flow - single label",
		flow: ate.Traffic().NewFlow("mpls flow").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithHeaders(
				NewEthernetHeader(),
				NewMPLSHeader().WithLabel(100)),
		wantFlow: &opb.Flow{
			Name:         "mpls flow",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			Headers: []*opb.Header{
				{Type: &opb.Header_Eth{&opb.EthernetHeader{}}},
				{Type: &opb.Header_Mpls{
					&opb.MplsHeader{
						Label: &opb.UIntRange{
							Min:    100,
							Max:    100,
							Count:  1,
							Random: false,
						},
						Ttl: 255,
					},
				}},
			},
		},
	}, {
		desc: "create mpls flow - label ranges",
		flow: ate.Traffic().NewFlow("mpls flow").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithHeaders(
				NewEthernetHeader(),
				func() *MPLSHeader {
					h := NewMPLSHeader()
					h.LabelRange().WithMin(100).WithCount(3)
					return h
				}()),
		wantFlow: &opb.Flow{
			Name:         "mpls flow",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			Headers: []*opb.Header{
				{Type: &opb.Header_Eth{&opb.EthernetHeader{}}},
				{Type: &opb.Header_Mpls{
					&opb.MplsHeader{
						Label: &opb.UIntRange{
							Min:   uint32(100),
							Max:   maxFlowLabel,
							Count: uint32(3),
						},
						Ttl: 255,
					},
				}},
			},
		},
	}, {
		desc: "frame size fixed",
		flow: ate.Traffic().NewFlow("fixed size flow").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithFrameSize(321),
		wantFlow: &opb.Flow{
			Name:         "fixed size flow",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			FrameSize: &opb.FrameSize{
				Type: &opb.FrameSize_Fixed{Fixed: 321},
			},
		},
	}, {
		desc: "frame size random",
		flow: ate.Traffic().NewFlow("random size flow").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithFrameSizeRandom(111, 222),
		wantFlow: &opb.Flow{
			Name:         "random size flow",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			FrameSize: &opb.FrameSize{
				Type: &opb.FrameSize_Random_{
					Random: &opb.FrameSize_Random{Min: 111, Max: 222},
				},
			},
		},
	}, {
		desc: "frame size preset IMIX",
		flow: ate.Traffic().NewFlow("std imix flow").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithFrameSizeIMIXStandard(),
		wantFlow: &opb.Flow{
			Name:         "std imix flow",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			FrameSize: &opb.FrameSize{
				Type: &opb.FrameSize_ImixPreset_{
					ImixPreset: opb.FrameSize_IMIX_STANDARD,
				},
			},
		},
	}, {
		desc: "frame rate percent",
		flow: ate.Traffic().NewFlow("frame rate percent").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithFrameRatePct(55),
		wantFlow: &opb.Flow{
			Name:         "frame rate percent",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			FrameRate: &opb.FrameRate{
				Type: &opb.FrameRate_Percent{Percent: 55},
			},
		},
	}, {
		desc: "frame rate bps",
		flow: ate.Traffic().NewFlow("frame rate bps").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithFrameRateBPS(123456789),
		wantFlow: &opb.Flow{
			Name:         "frame rate bps",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			FrameRate: &opb.FrameRate{
				Type: &opb.FrameRate_Bps{Bps: 123456789},
			},
		},
	}, {
		desc: "frame rate fps",
		flow: ate.Traffic().NewFlow("frame rate fps").
			WithSrcEndpoints(intf1).
			WithDstEndpoints(intf2).
			WithFrameRateFPS(123),
		wantFlow: &opb.Flow{
			Name:         "frame rate fps",
			SrcEndpoints: srcEPs,
			DstEndpoints: dstEPs,
			FrameRate: &opb.FrameRate{
				Type: &opb.FrameRate_Fps{Fps: 123},
			},
		},
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ate.Traffic().Start(t, tt.flow)
			if want := "ix1"; got.ate != want {
				t.Errorf("unexpected ATE name, got: %q, want: %q", got.ate, want)
			}
			if gotFlows := len(got.flows); gotFlows != 1 {
				t.Fatalf("wrong number of flows started: got: %d, want: 1", gotFlows)
			}
			gotFlow := got.flows[0]
			if diff := cmp.Diff(gotFlow, tt.wantFlow, cmp.Comparer(proto.Equal)); diff != "" {
				t.Fatalf("did not get expected flow: diff(-got,+want):\n%s", diff)
			}
		})
	}
}

func TestStartTrafficErrors(t *testing.T) {
	initATEFakes(t)
	ate := ATE(t, "ate")
	port := ate.Port(t, "port2")
	intf := ate.Topology().New().AddInterface("intf").WithPort(port)

	tests := []struct {
		desc         string
		flow         *Flow
		wantFatalMsg string
	}{{
		desc:         "missing source endpoint",
		flow:         ate.Traffic().NewFlow("gaga").WithDstEndpoints(intf),
		wantFatalMsg: "src endpoint",
	}, {
		desc:         "missing destination endpoint",
		flow:         ate.Traffic().NewFlow("gaga").WithSrcEndpoints(intf),
		wantFatalMsg: "dst endpoint",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := negtest.ExpectFatal(t, func(t testing.TB) {
				ate.Traffic().Start(t, tt.flow)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Fatalf("Start() failed with message %q, want: %q", got, tt.wantFatalMsg)
			}
		})
	}
}

func TestUpdateTraffic(t *testing.T) {
	initATEFakes(t)
	ate := ATE(t, "ate")
	p1 := ate.Port(t, "port1")
	p2 := ate.Port(t, "port2")
	top := ate.Topology().New()
	intf1 := top.AddInterface("intf1").WithPort(p1)
	intf2 := top.AddInterface("intf2").WithPort(p2)
	srcEPs := []*opb.Flow_Endpoint{{InterfaceName: intf1.pb.GetName()}}
	dstEPs := []*opb.Flow_Endpoint{{InterfaceName: intf2.pb.GetName()}}

	flow := ate.Traffic().NewFlow("gaga").
		WithSrcEndpoints(intf1).
		WithDstEndpoints(intf2).
		WithFrameSize(123)

	wantFlow := &opb.Flow{
		Name:         "gaga",
		SrcEndpoints: srcEPs,
		DstEndpoints: dstEPs,
		FrameSize: &opb.FrameSize{
			Type: &opb.FrameSize_Fixed{Fixed: 123},
		},
	}

	wantMod := &opb.Flow{
		Name:         "gaga",
		SrcEndpoints: srcEPs,
		DstEndpoints: dstEPs,
		FrameSize: &opb.FrameSize{
			Type: &opb.FrameSize_Fixed{Fixed: 456},
		},
	}

	ate.Traffic().Start(t, flow)
	if want := "ix1"; got.ate != want {
		t.Errorf("Start(): unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if gotFlows := len(got.flows); gotFlows != 1 {
		t.Fatalf("Start(): wrong number of flows started: got: %d, want: 1", gotFlows)
	}
	gotFlow := got.flows[0]
	if diff := cmp.Diff(gotFlow, wantFlow, cmp.Comparer(proto.Equal)); diff != "" {
		t.Fatalf("Start(): did not get expected flow: diff(-got,+want):\n%s", diff)
	}

	flow.WithFrameSize(456)
	ate.Traffic().Update(t, flow)
	if want := "ix1"; got.ate != want {
		t.Errorf("Update(): unexpected ATE name, got: %q, want: %q", got.ate, want)
	}
	if gotFlows := len(got.flows); gotFlows != 1 {
		t.Fatalf("Update(): wrong number of flows started: got: %d, want: 1", gotFlows)
	}
	gotFlow = got.flows[0]
	if diff := cmp.Diff(gotFlow, wantMod, cmp.Comparer(proto.Equal)); diff != "" {
		t.Fatalf("Update(): did not get expected flow: diff(-got,+want):\n%s", diff)
	}
}
