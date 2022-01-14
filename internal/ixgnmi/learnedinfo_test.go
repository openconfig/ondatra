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

func TestLearnedInfoToRIB(t *testing.T) {
	tests := []struct {
		name          string
		inLearnedInfo []bgpLearnedInfo
		inV4          bool
		inRIB         *telemetry.NetworkInstance_Protocol_Bgp_Rib
		wantRIB       *telemetry.NetworkInstance_Protocol_Bgp_Rib
		wantErr       string
	}{{
		name: "invalid attr set index",
		inRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{
			AttrSet: map[uint64]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AttrSet{
				1: &telemetry.NetworkInstance_Protocol_Bgp_Rib_AttrSet{},
			},
		},
		inLearnedInfo: []bgpLearnedInfo{{
			IPV4NextHop: "localhost",
		}},
		wantErr: "failed to append details for elem 0",
	}, {
		name: "invalid community index",
		inRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{
			Community: map[uint64]*telemetry.NetworkInstance_Protocol_Bgp_Rib_Community{
				1: &telemetry.NetworkInstance_Protocol_Bgp_Rib_Community{},
			},
		},
		inLearnedInfo: []bgpLearnedInfo{{
			Community: "65532 : 10200",
		}},
		wantErr: "failed to append details for elem 0",
	}, {
		name:  "invalid origin type",
		inRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{},
		inLearnedInfo: []bgpLearnedInfo{{
			Origin: "foo",
		}},
		wantErr: "unknown origin type",
	}, {
		name:  "duplicate v4 route",
		inV4:  true,
		inRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{},
		inLearnedInfo: []bgpLearnedInfo{{
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
		name:  "duplicate v6 route",
		inRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{},
		inLearnedInfo: []bgpLearnedInfo{{
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
		name:          "empty learned info",
		inRIB:         &telemetry.NetworkInstance_Protocol_Bgp_Rib{},
		inLearnedInfo: []bgpLearnedInfo{{}},
		wantRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{
			AfiSafi: map[telemetry.E_BgpTypes_AFI_SAFI_TYPE]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi{
				telemetry.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST: {
					AfiSafiName: telemetry.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST,
					Ipv6Unicast: &telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast{
						Neighbor: map[string]*telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor{
							"localhost": {
								NeighborAddress: ygot.String("localhost"),
								AdjRibInPre:     &telemetry.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre{},
							},
						},
					},
				},
			},
		},
	}, {
		name:  "v4 routes",
		inV4:  true,
		inRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{},
		inLearnedInfo: []bgpLearnedInfo{{
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
							"localhost": {
								NeighborAddress: ygot.String("localhost"),
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
		name:  "v6 routes",
		inRIB: &telemetry.NetworkInstance_Protocol_Bgp_Rib{},
		inLearnedInfo: []bgpLearnedInfo{{
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
							"localhost": {
								NeighborAddress: ygot.String("localhost"),
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := learnedInfoToRIB(tt.inLearnedInfo, "localhost", tt.inV4, tt.inRIB)
			if d := errdiff.Substring(err, tt.wantErr); d != "" {
				t.Fatalf("unexpected error diff\n%s", d)
			}
			if err != nil {
				return
			}
			if d := cmp.Diff(tt.wantRIB, tt.inRIB); d != "" {
				t.Errorf("unexpected RIB diff (-want +got)\n%s", d)
			}
		})
	}
}
