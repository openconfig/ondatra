// Copyright 2021 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ixgnmi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ygot/ygot"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/telemetry"
)

func TestLearnedBGPToOC(t *testing.T) {
	tests := []struct {
		desc    string
		info    []bgpLearnedInfo
		v4      bool
		wantRIB *telemetry.NetworkInstance_Protocol_Bgp_Rib
		wantErr string
	}{{
		desc: "invalid attr set index",
		info: []bgpLearnedInfo{{
			IPV4NextHop: "localhost",
		}},
		wantErr: "failed to append details for elem 0",
	}, {
		desc: "invalid community index",
		info: []bgpLearnedInfo{{
			Community: "65532 : 10200",
		}},
		wantErr: "failed to append details for elem 0",
	}, {
		desc: "invalid origin type",
		info: []bgpLearnedInfo{{
			Origin: "foo",
		}},
		wantErr: "unknown origin type",
	}, {
		desc: "duplicate v4 route",
		v4:   true,
		info: []bgpLearnedInfo{{
			IPV4Prefix: "127.0.0.1",
			PathID:     0,
			Origin:     "IGP",
		}, {
			IPV4Prefix: "127.0.0.1",
			PathID:     0,
			Origin:     "IGP",
		}},
		wantErr: "failed to append route for elem 1",
	}, {
		desc: "duplicate v6 route",
		info: []bgpLearnedInfo{{
			IPV6Prefix: "::1",
			PathID:     0,
			Origin:     "IGP",
		}, {
			IPV6Prefix: "::1",
			PathID:     0,
			Origin:     "IGP",
		}},
		wantErr: "failed to append route for elem 1",
	}, {
		desc: "empty learned info",
		info: []bgpLearnedInfo{{}},
		wantRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{
			AfiSafi: map[telemetry.E_BgpTypes_AFI_SAFI_TYPE]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi{
				telemetry.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST: {
					AfiSafiName: telemetry.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST,
					Ipv6Unicast: &telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast{
						Neighbor: map[string]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor{
							"neighbor": {
								NeighborAddress: ygot.String("neighbor"),
								AdjRibInPre:     &telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre{},
							},
						},
					},
				},
			},
		},
	}, {
		desc: "v4 routes",
		v4:   true,
		info: []bgpLearnedInfo{{
			IPV4Prefix:  "127.0.0.1",
			PathID:      0,
			IPV4NextHop: "127.0.0.2",
			MED:         100,
			LocalPref:   1000,
			Origin:      "IGP",
			Community:   "65532 : 10200, 65533 : 10100",
			AIGP:        200,
		}, {
			IPV4Prefix:  "127.0.0.3",
			IPV4NextHop: "127.0.0.4",
			Origin:      "EGP",
		}},
		wantRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{
			AttrSet: map[uint64]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AttrSet{
				0: {
					Aigp:      ygot.Uint64(200),
					Index:     ygot.Uint64(0),
					LocalPref: ygot.Uint32(1000),
					Med:       ygot.Uint32(100),
					NextHop:   ygot.String("127.0.0.2"),
					Origin:    telemetry.BgpTypes_BgpOriginAttrType_IGP,
				},
				1: {
					Index:     ygot.Uint64(1),
					Aigp:      ygot.Uint64(0),
					LocalPref: ygot.Uint32(0),
					Med:       ygot.Uint32(0),
					NextHop:   ygot.String("127.0.0.4"),
					Origin:    telemetry.BgpTypes_BgpOriginAttrType_EGP,
				},
			},
			Community: map[uint64]*telemetry.NetworkInstance_Protocol_Bgp_Rib_Community{
				0: {
					Community: []telemetry.NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union{telemetry.UnionString("65532 : 10200"), telemetry.UnionString(" 65533 : 10100")},
					Index:     ygot.Uint64(0),
				},
				1: {Index: ygot.Uint64(1)},
			},
			AfiSafi: map[telemetry.E_BgpTypes_AFI_SAFI_TYPE]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi{
				telemetry.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST: {
					AfiSafiName: telemetry.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST,
					Ipv4Unicast: &telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast{
						Neighbor: map[string]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor{
							"neighbor": {
								NeighborAddress: ygot.String("neighbor"),
								AdjRibInPre: &telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre{
									Route: map[telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_Key]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route{
										{Prefix: "127.0.0.1/0"}: {
											AttrIndex:      ygot.Uint64(0),
											CommunityIndex: ygot.Uint64(0),
											PathId:         ygot.Uint32(0),
											Prefix:         ygot.String("127.0.0.1/0"),
										},
										{Prefix: "127.0.0.3/0"}: {
											AttrIndex:      ygot.Uint64(1),
											CommunityIndex: ygot.Uint64(1),
											PathId:         ygot.Uint32(0),
											Prefix:         ygot.String("127.0.0.3/0"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}, {
		desc: "v6 routes",
		info: []bgpLearnedInfo{{
			IPV6Prefix:  "::1",
			IPV6NextHop: "::2",
			MED:         100,
			LocalPref:   1000,
			Origin:      "Incomplete",
			AIGP:        200,
		}, {
			IPV6Prefix:  "::3",
			IPV6NextHop: "::4",
			Origin:      "IGP",
		}},
		wantRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{
			AttrSet: map[uint64]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AttrSet{
				0: {
					Aigp:      ygot.Uint64(200),
					Index:     ygot.Uint64(0),
					LocalPref: ygot.Uint32(1000),
					Med:       ygot.Uint32(100),
					NextHop:   ygot.String("::2"),
					Origin:    telemetry.BgpTypes_BgpOriginAttrType_INCOMPLETE,
				},
				1: {
					Index:     ygot.Uint64(1),
					Aigp:      ygot.Uint64(0),
					LocalPref: ygot.Uint32(0),
					Med:       ygot.Uint32(0),
					NextHop:   ygot.String("::4"),
					Origin:    telemetry.BgpTypes_BgpOriginAttrType_IGP,
				},
			},
			Community: map[uint64]*telemetry.NetworkInstance_Protocol_Bgp_Rib_Community{
				0: {
					Index: ygot.Uint64(0),
				},
				1: {Index: ygot.Uint64(1)},
			},
			AfiSafi: map[telemetry.E_BgpTypes_AFI_SAFI_TYPE]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi{
				telemetry.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST: {
					AfiSafiName: telemetry.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST,
					Ipv6Unicast: &telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast{
						Neighbor: map[string]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor{
							"neighbor": {
								NeighborAddress: ygot.String("neighbor"),
								AdjRibInPre: &telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre{
									Route: map[telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_Key]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route{
										{Prefix: "::1/0"}: {
											AttrIndex:      ygot.Uint64(0),
											CommunityIndex: ygot.Uint64(0),
											PathId:         ygot.Uint32(0),
											Prefix:         ygot.String("::1/0"),
										},
										{Prefix: "::3/0"}: {
											AttrIndex:      ygot.Uint64(1),
											CommunityIndex: ygot.Uint64(1),
											PathId:         ygot.Uint32(0),
											Prefix:         ygot.String("::3/0"),
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := learnedBGPToOC(test.info, "interface", "protocol", "neighbor", test.v4)
			if d := errdiff.Substring(err, test.wantErr); d != "" {
				t.Fatalf("learnedBGPToOC got unexpected error diff\n%s", d)
			}
			if err != nil {
				return
			}
			gotRIB := got.GetNetworkInstance("interface").
				GetProtocol(telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP, "protocol").GetBgp().GetRib()
			if d := cmp.Diff(test.wantRIB, gotRIB); d != "" {
				t.Errorf("learnedBGPToOC got unexpected RIB diff (-want +got)\n%s", d)
			}
		})
	}
}

func TestLearnedISISToOC(t *testing.T) {
	tests := []struct {
		desc     string
		info     []isisLearnedInfo
		wantISIS *telemetry.NetworkInstance_Protocol_Isis
		wantErr  string
	}{{
		desc:     "empty learned info",
		info:     []isisLearnedInfo{{}},
		wantISIS: &telemetry.NetworkInstance_Protocol_Isis{},
	}, {
		desc: "success",
		info: []isisLearnedInfo{{
			LearnedVia:      "L1",
			SystemID:        "01 23 45 67 89 AB",
			PseudoNodeIndex: 2,
			LSPIndex:        34,
			SeqNumber:       5,
		}, {
			LearnedVia:      "L2",
			SystemID:        "CD EF 01 23 45 67",
			PseudoNodeIndex: 98,
			LSPIndex:        7,
			SeqNumber:       0,
		}},
		wantISIS: &telemetry.NetworkInstance_Protocol_Isis{
			Level: map[uint8]*telemetry.NetworkInstance_Protocol_Isis_Level{
				1: &telemetry.NetworkInstance_Protocol_Isis_Level{
					LevelNumber: ygot.Uint8(1),
					Lsp: map[string]*telemetry.NetworkInstance_Protocol_Isis_Level_Lsp{
						"0123.4567.89AB.02-34": &telemetry.NetworkInstance_Protocol_Isis_Level_Lsp{
							IsType:         ygot.Uint8(1),
							LspId:          ygot.String("0123.4567.89AB.02-34"),
							SequenceNumber: ygot.Uint32(5),
						},
					},
				},
				2: &telemetry.NetworkInstance_Protocol_Isis_Level{
					LevelNumber: ygot.Uint8(2),
					Lsp: map[string]*telemetry.NetworkInstance_Protocol_Isis_Level_Lsp{
						"CDEF.0123.4567.98-07": &telemetry.NetworkInstance_Protocol_Isis_Level_Lsp{
							IsType:         ygot.Uint8(2),
							LspId:          ygot.String("CDEF.0123.4567.98-07"),
							SequenceNumber: ygot.Uint32(0),
						},
					},
				},
			},
		},
	}, {
		desc:    "unknown level",
		info:    []isisLearnedInfo{{LearnedFrom: "L3"}},
		wantErr: "level",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := learnedISISToOC(test.info, "interface", "protocol")
			if d := errdiff.Substring(err, test.wantErr); d != "" {
				t.Fatalf("learnedISISToOC got unexpected error diff\n%s", d)
			}
			if err != nil {
				return
			}
			gotISIS := got.GetNetworkInstance("interface").
				GetProtocol(telemetry.PolicyTypes_INSTALL_PROTOCOL_TYPE_ISIS, "protocol").GetIsis()
			if d := cmp.Diff(test.wantISIS, gotISIS); d != "" {
				t.Errorf("learnedISISToOC got unexpected ISIS diff (-want +got)\n%s", d)
			}
		})
	}
}
