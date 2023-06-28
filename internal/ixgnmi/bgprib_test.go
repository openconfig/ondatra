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
	"errors"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gnmi/errdiff"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"github.com/openconfig/ygot/ygot"
)

func TestBGPRIBFromIxia(t *testing.T) {
	const (
		bgp4ID = "/fake/bgp/v4/id"
		bgp6ID = "/fake/bgp/v6/id"
	)
	bgp4XP := parseXPath(t, "/xpath/to/bgpv4")
	bgp6XP := parseXPath(t, "/xpath/to/bgpv6")
	nodes := &cachedNodes{
		bgp4Peers: []*ixconfig.TopologyBgpIpv4Peer{{Xpath: bgp4XP, DutIp: ixconfig.MultivalueStr("1.2.3.4")}},
		bgp6Peers: []*ixconfig.TopologyBgpIpv6Peer{{Xpath: bgp6XP, DutIp: ixconfig.MultivalueStr("::1")}},
	}

	tests := []struct {
		desc    string
		postErr map[string]error
		getErr  map[string]error
		getRsps map[string]string
		want    *oc.NetworkInstance
		wantErr string
	}{{
		desc: "run op v4 error",
		postErr: map[string]error{
			"topology/deviceGroup/ethernet/ipv4/bgpIpv4Peer/operations/getAllLearnedInfo": errors.New("op v4 fail"),
		},
		wantErr: "op v4 fail",
	}, {
		desc: "run op v6 error",
		postErr: map[string]error{
			"topology/deviceGroup/ethernet/ipv6/bgpIpv6Peer/operations/getAllLearnedInfo": errors.New("op v6 fail"),
		},
		wantErr: "op v6 fail",
	}, {
		desc: "get v4 error",
		getErr: map[string]error{
			bgp4ID + "/learnedInfo/1/table/1": errors.New("get v4 fail"),
		},
		wantErr: "get v4 fail",
	}, {
		desc: "get v6 error",
		getErr: map[string]error{
			bgp6ID + "/learnedInfo/1/table/1": errors.New("get v6 fail"),
		},
		wantErr: "get v6 fail",
	}, {
		desc: "unknown origin",
		getRsps: map[string]string{
			bgp4ID + "/learnedInfo/1/table/1": `{
				"columns": ["Origin"],
				"values": [["Bad Origin"]]
			}`,
		},
		wantErr: "unknown origin",
	}, {
		desc: "duplicate v4 prefix",
		getRsps: map[string]string{
			bgp4ID + "/learnedInfo/1/table/1": `{
				"columns": ["IPv4 Prefix ", "Origin"],
				"values": [["127.0.0.1", "EGP"], ["127.0.0.1", "EGP"]]
			}`,
		},
		wantErr: "duplicate key",
	}, {
		desc: "duplicate v6 prefix",
		getRsps: map[string]string{
			bgp4ID + "/learnedInfo/1/table/1": `{
				"columns": ["IPv6 Prefix", "Origin"],
				"values": [["::1", "EGP"], ["::1", "EGP"]]
			}`,
		},
		wantErr: "duplicate key",
	}, {
		desc: "invalid AS segment",
		getRsps: map[string]string{
			bgp4ID + "/learnedInfo/1/table/1": `{
					"columns": ["Origin", "AS Path"],
					"values": [["EGP", "<foo bar>"]]
				}`,
		},
		wantErr: "invalid AS segment",
	}, {
		desc: "no data",
		want: &oc.NetworkInstance{
			Protocol: map[oc.NetworkInstance_Protocol_Key]*oc.NetworkInstance_Protocol{
				oc.NetworkInstance_Protocol_Key{
					Identifier: oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP,
					Name:       "0",
				}: {
					Identifier: oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP,
					Name:       ygot.String("0"),
					Bgp: &oc.NetworkInstance_Protocol_Bgp{
						Rib: &oc.NetworkInstance_Protocol_Bgp_Rib{
							AfiSafi: map[oc.E_BgpTypes_AFI_SAFI_TYPE]*oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi{
								oc.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST: {
									AfiSafiName: oc.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST,
									Ipv4Unicast: &oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast{
										Neighbor: map[string]*oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor{
											"1.2.3.4": {
												NeighborAddress: ygot.String("1.2.3.4"),
												AdjRibInPre:     &oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre{},
											},
										},
									},
								},
								oc.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST: {
									AfiSafiName: oc.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST,
									Ipv6Unicast: &oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast{
										Neighbor: map[string]*oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor{
											"::1": {
												NeighborAddress: ygot.String("::1"),
												AdjRibInPre:     &oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre{},
											},
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
		desc: "full data",
		getRsps: map[string]string{
			bgp4ID + "/learnedInfo/1/table/1": `{
				"id": 1,
				"columns": ["IPv4 Prefix ", "Prefix Length", "Path ID", "IPv4 Next Hop", "Origin", "AIGP", "Local Preference", "MED", "Community", "Color", "AS Path"],
				"values": [
					["127.0.0.1", "30", "1", "127.0.0.2", "IGP", "200", "1000", "100", "65532 : 10200, 65533 : 10100", "foo, bar", "<65532 65533>"],
					["127.0.0.3", "24", "2", "127.0.0.4", "EGP",    "",     "",    "",                             "",         "",              ""]
				]
			}`,
			bgp6ID + "/learnedInfo/1/table/1": `{
				"id": 2,
				"columns": ["IPv6 Prefix", "Prefix Length", "Path ID", "IPv6 Next Hop", "Origin", "AIGP", "Local Preference", "MED", "Community", "Color", "AS Path"],
				"values": [
					["::1", "28", "3", "::2", "Incomplete", "200", "1000", "100", "65534 : 10400, 65535 : 10300", "baz", "<65534 65535>"],
					["::3", "26", "4", "::4",        "IGP",   "",     "",    "",                              "",    "",              ""]
				]
			}`,
		},
		want: &oc.NetworkInstance{
			Protocol: map[oc.NetworkInstance_Protocol_Key]*oc.NetworkInstance_Protocol{
				oc.NetworkInstance_Protocol_Key{
					Identifier: oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP,
					Name:       "0",
				}: {
					Identifier: oc.PolicyTypes_INSTALL_PROTOCOL_TYPE_BGP,
					Name:       ygot.String("0"),
					Bgp: &oc.NetworkInstance_Protocol_Bgp{
						Rib: &oc.NetworkInstance_Protocol_Bgp_Rib{
							AfiSafi: map[oc.E_BgpTypes_AFI_SAFI_TYPE]*oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi{
								oc.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST: {
									AfiSafiName: oc.BgpTypes_AFI_SAFI_TYPE_IPV4_UNICAST,
									Ipv4Unicast: &oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast{
										Neighbor: map[string]*oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor{
											"1.2.3.4": {
												NeighborAddress: ygot.String("1.2.3.4"),
												AdjRibInPre: &oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre{
													Route: map[oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_Key]*oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route{
														{Prefix: "127.0.0.1/30", PathId: 1}: {
															AttrIndex:      ygot.Uint64(0),
															CommunityIndex: ygot.Uint64(0),
															PathId:         ygot.Uint32(1),
															Prefix:         ygot.String("127.0.0.1/30"),
														},
														{Prefix: "127.0.0.3/24", PathId: 2}: {
															AttrIndex:      ygot.Uint64(1),
															CommunityIndex: ygot.Uint64(1),
															PathId:         ygot.Uint32(2),
															Prefix:         ygot.String("127.0.0.3/24"),
														},
													},
												},
											},
										},
									},
								},
								oc.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST: {
									AfiSafiName: oc.BgpTypes_AFI_SAFI_TYPE_IPV6_UNICAST,
									Ipv6Unicast: &oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast{
										Neighbor: map[string]*oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor{
											"::1": {
												NeighborAddress: ygot.String("::1"),
												AdjRibInPre: &oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre{
													Route: map[oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_Key]*oc.NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route{
														{Prefix: "::1/28", PathId: 3}: {
															AttrIndex:      ygot.Uint64(2),
															CommunityIndex: ygot.Uint64(2),
															PathId:         ygot.Uint32(3),
															Prefix:         ygot.String("::1/28"),
														},
														{Prefix: "::3/26", PathId: 4}: {
															AttrIndex:      ygot.Uint64(3),
															CommunityIndex: ygot.Uint64(3),
															PathId:         ygot.Uint32(4),
															Prefix:         ygot.String("::3/26"),
														},
													},
												},
											},
										},
									},
								},
							},
							AttrSet: map[uint64]*oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet{
								0: &oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet{
									Aigp:      ygot.Uint64(200),
									Index:     ygot.Uint64(0),
									LocalPref: ygot.Uint32(1000),
									Med:       ygot.Uint32(100),
									NextHop:   ygot.String("127.0.0.2"),
									Origin:    oc.RibBgp_BgpOriginAttrType_IGP,
									AsSegment: map[uint32]*oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment{0: {
										Index:  ygot.Uint32(0),
										Member: []uint32{65532, 65533},
										Type:   oc.RibBgp_AsPathSegmentType_AS_SEQ,
									}},
								},
								1: &oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet{
									Aigp:      ygot.Uint64(0),
									Index:     ygot.Uint64(1),
									LocalPref: ygot.Uint32(0),
									Med:       ygot.Uint32(0),
									NextHop:   ygot.String("127.0.0.4"),
									Origin:    oc.RibBgp_BgpOriginAttrType_EGP,
								},
								2: &oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet{
									Aigp:      ygot.Uint64(200),
									Index:     ygot.Uint64(2),
									LocalPref: ygot.Uint32(1000),
									Med:       ygot.Uint32(100),
									NextHop:   ygot.String("::2"),
									Origin:    oc.RibBgp_BgpOriginAttrType_INCOMPLETE,
									AsSegment: map[uint32]*oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment{0: {
										Index:  ygot.Uint32(0),
										Member: []uint32{65534, 65535},
										Type:   oc.RibBgp_AsPathSegmentType_AS_SEQ,
									}},
								},
								3: &oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet{
									Aigp:      ygot.Uint64(0),
									Index:     ygot.Uint64(3),
									LocalPref: ygot.Uint32(0),
									Med:       ygot.Uint32(0),
									NextHop:   ygot.String("::4"),
									Origin:    oc.RibBgp_BgpOriginAttrType_IGP,
								},
							},
							Community: map[uint64]*oc.NetworkInstance_Protocol_Bgp_Rib_Community{
								0: &oc.NetworkInstance_Protocol_Bgp_Rib_Community{
									Index: ygot.Uint64(0),
									Community: []oc.NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union{
										oc.UnionString("65532 : 10200"),
										oc.UnionString("65533 : 10100"),
									},
								},
								1: &oc.NetworkInstance_Protocol_Bgp_Rib_Community{
									Index: ygot.Uint64(1),
								},
								2: &oc.NetworkInstance_Protocol_Bgp_Rib_Community{
									Index: ygot.Uint64(2),
									Community: []oc.NetworkInstance_Protocol_Bgp_Rib_Community_Community_Union{
										oc.UnionString("65534 : 10400"),
										oc.UnionString("65535 : 10300"),
									},
								},
								3: &oc.NetworkInstance_Protocol_Bgp_Rib_Community{
									Index: ygot.Uint64(3),
								},
							},
							ExtCommunity: map[uint64]*oc.NetworkInstance_Protocol_Bgp_Rib_ExtCommunity{
								0: &oc.NetworkInstance_Protocol_Bgp_Rib_ExtCommunity{
									Index: ygot.Uint64(0),
									ExtCommunity: []oc.NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_Union{
										oc.UnionString("foo"),
										oc.UnionString("bar"),
									},
								},
								1: &oc.NetworkInstance_Protocol_Bgp_Rib_ExtCommunity{
									Index: ygot.Uint64(1),
								},
								2: &oc.NetworkInstance_Protocol_Bgp_Rib_ExtCommunity{
									Index: ygot.Uint64(2),
									ExtCommunity: []oc.NetworkInstance_Protocol_Bgp_Rib_ExtCommunity_ExtCommunity_Union{
										oc.UnionString("baz"),
									},
								},
								3: &oc.NetworkInstance_Protocol_Bgp_Rib_ExtCommunity{
									Index: ygot.Uint64(3),
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
			getRsps := make(map[string][]string)
			for p, r := range test.getRsps {
				getRsps[p] = []string{r}
			}
			client := &fakeCfgClient{
				sess: &fakeSession{
					postErrs: test.postErr,
					getErrs:  test.getErr,
					getRsps:  getRsps,
				},
				xpathToID: map[string]string{
					bgp4XP.String(): bgp4ID,
					bgp6XP.String(): bgp6ID,
				},
			}

			got := new(oc.NetworkInstance)
			err := bgpRIBFromIxia(context.Background(), client, got, nodes)
			if d := errdiff.Substring(err, test.wantErr); d != "" {
				t.Fatalf("bgpRIBFromIxia() got unexpected error diff\n%s", d)
			}
			if err != nil {
				return
			}
			if d := cmp.Diff(test.want, got); d != "" {
				t.Errorf("bgpRIBFromIxia() got unexpected diff (-want +got)\n%s", d)
			}
		})
	}
}

func TestParseASSegments(t *testing.T) {
	tests := []struct {
		desc    string
		asPath  string
		want    []*oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment
		wantErr string
	}{{
		desc: "empty",
	}, {
		desc:   "single seq",
		asPath: "<1 2 3>",
		want: []*oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment{{
			Index:  ygot.Uint32(0),
			Type:   oc.RibBgp_AsPathSegmentType_AS_SEQ,
			Member: []uint32{1, 2, 3},
		}},
	}, {
		desc:   "single set",
		asPath: "{1,2,3}",
		want: []*oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment{{
			Index:  ygot.Uint32(0),
			Type:   oc.RibBgp_AsPathSegmentType_AS_SET,
			Member: []uint32{1, 2, 3},
		}},
	}, {
		desc:   "complicated",
		asPath: " < 1 2 > {3 , 4,		5 } <6 >",
		want: []*oc.NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment{{
			Index:  ygot.Uint32(0),
			Type:   oc.RibBgp_AsPathSegmentType_AS_SEQ,
			Member: []uint32{1, 2},
		}, {
			Index:  ygot.Uint32(1),
			Type:   oc.RibBgp_AsPathSegmentType_AS_SET,
			Member: []uint32{3, 4, 5},
		}, {
			Index:  ygot.Uint32(2),
			Type:   oc.RibBgp_AsPathSegmentType_AS_SEQ,
			Member: []uint32{6},
		}},
	}, {
		desc:    "invalid",
		asPath:  "<foo>",
		wantErr: "invalid",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			segs, err := parseASSegments(test.asPath)
			if (err == nil) != (test.wantErr == "") || (err != nil && !strings.Contains(err.Error(), test.wantErr)) {
				t.Errorf("parseASSegments(%q) got error %v, want error containing %q", test.asPath, err, test.wantErr)
			}
			if diff := cmp.Diff(test.want, segs); diff != "" {
				t.Errorf("parseASSegments(%q) got unexpected diff (-want +got): %s", test.asPath, diff)
			}
		})
	}
}
