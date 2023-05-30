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
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/openconfig/ondatra/internal/ixconfig"

	opb "github.com/openconfig/ondatra/proto"
)

func TestAddNetworks(t *testing.T) {
	const (
		ifName   = "someIntf"
		net1Name = "someNet"
		net2Name = "anotherNet"
	)
	baseClient := func(withISIS, withBGPv4Peer, withBGPv6Peer bool) *ixATE {
		cfg := &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
					Ethernet: []*ixconfig.TopologyEthernet{{}},
				}},
			}},
		}
		ifc := &intf{deviceGroup: cfg.Topology[0].DeviceGroup[0]}
		if withISIS {
			ifc.deviceGroup.Ethernet[0].IsisL3 = []*ixconfig.TopologyIsisL3{{}}
		}
		if withBGPv4Peer {
			ifc.deviceGroup.Ethernet[0].Ipv4 = []*ixconfig.TopologyIpv4{{
				BgpIpv4Peer: []*ixconfig.TopologyBgpIpv4Peer{{}},
			}}
			ifc.ipv4 = ifc.deviceGroup.Ethernet[0].Ipv4[0]
		}
		if withBGPv6Peer {
			ifc.deviceGroup.Ethernet[0].Ipv6 = []*ixconfig.TopologyIpv6{{
				BgpIpv6Peer: []*ixconfig.TopologyBgpIpv6Peer{{}},
			}}
			ifc.ipv6 = ifc.deviceGroup.Ethernet[0].Ipv6[0]
		}
		return &ixATE{
			cfg:   cfg,
			intfs: map[string]*intf{ifName: ifc},
		}
	}
	tests := []struct {
		desc                                   string
		withISIS, withBGPv4Peer, withBGPv6Peer bool
		ifc                                    *opb.InterfaceConfig
		wantErr                                bool
		wantNgs                                []*ixconfig.TopologyNetworkGroup
		wantRouteTables                        map[string]*routeTables
	}{{
		desc: "MAC Pool",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				Eth: &opb.NetworkEth{
					MacAddress: "11:22:33:44:55:66",
					Count:      1,
					VlanId:     0,
				},
			}},
		},
		wantNgs: []*ixconfig.TopologyNetworkGroup{{
			Name: ixconfig.String(net1Name),
			MacPools: []*ixconfig.TopologyMacPools{{
				EnableVlans:          ixconfig.MultivalueFalse(),
				Mac:                  ixconfig.MultivalueStr("11:22:33:44:55:66"),
				NumberOfAddressesAsy: ixconfig.MultivalueUint32(1),
				Vlan: []*ixconfig.TopologyVlan{{
					VlanId: ixconfig.MultivalueUint32(0),
				}},
			}},
		}},
	}, {
		desc: "Bad Ipv4 config",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				Ipv4: &opb.NetworkIp{
					Count: 1,
				},
			}},
		},
		wantErr: true,
	}, {
		desc: "Bad Ipv6 config",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				Ipv6: &opb.NetworkIp{
					Count: 1,
				},
			}},
		},
		wantErr: true,
	}, {
		desc: "Ipv4 Pool",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				Ipv4: &opb.NetworkIp{
					AddressCidr: "10.0.0.0/8",
					Count:       1,
				},
			}},
		},
		wantNgs: []*ixconfig.TopologyNetworkGroup{{
			Name: ixconfig.String(net1Name),
			Ipv4PrefixPools: []*ixconfig.TopologyIpv4PrefixPools{{
				NetworkAddress:       ixconfig.MultivalueStr("10.0.0.0"),
				PrefixLength:         ixconfig.MultivalueUint32(8),
				NumberOfAddressesAsy: ixconfig.MultivalueUint32(1),
			}},
		}},
	}, {
		desc: "Ipv6 Pool",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				Ipv6: &opb.NetworkIp{
					AddressCidr: "2001::4860:10:0:0:0/104",
					Count:       1,
				},
			}},
		},
		wantNgs: []*ixconfig.TopologyNetworkGroup{{
			Name: ixconfig.String(net1Name),
			Ipv6PrefixPools: []*ixconfig.TopologyIpv6PrefixPools{{
				NetworkAddress:       ixconfig.MultivalueStr("2001:0:0:4860:10::"),
				PrefixLength:         ixconfig.MultivalueUint32(104),
				NumberOfAddressesAsy: ixconfig.MultivalueUint32(1),
			}},
		}},
	}, {
		desc: "IS-IS and BGP attributes on the same IPv4 network",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				Ipv4: &opb.NetworkIp{
					AddressCidr: "10.0.0.0/8",
					Count:       1,
				},
				Isis: &opb.IPReachability{
					RouteOrigin: opb.IPReachability_INTERNAL,
				},
				BgpAttributes: &opb.BgpAttributes{
					Active:     true,
					Origin:     opb.BgpAttributes_ORIGIN_IGP,
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
				},
			}},
		},
		wantErr: true,
	}, {
		desc: "IS-IS and BGP attributes on the same IPv6 network",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				Ipv4: &opb.NetworkIp{
					AddressCidr: "10.0.0.0/8",
					Count:       1,
				},
				Isis: &opb.IPReachability{
					RouteOrigin: opb.IPReachability_INTERNAL,
				},
				BgpAttributes: &opb.BgpAttributes{
					Active:     true,
					Origin:     opb.BgpAttributes_ORIGIN_IGP,
					AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
				},
			}},
		},
		wantErr: true,
	}, {
		desc: "Advertise routes via BGP peer on different protocol",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				Ipv4: &opb.NetworkIp{
					AddressCidr: "10.0.0.0/8",
					Count:       1,
				},
				BgpAttributes: &opb.BgpAttributes{
					Active:                true,
					Origin:                opb.BgpAttributes_ORIGIN_IGP,
					AsnSetMode:            opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
					AdvertisementProtocol: opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_V6,
				},
			}, {
				Name:          net2Name,
				InterfaceName: ifName,
				Ipv6: &opb.NetworkIp{
					AddressCidr: "2002::4860:10:0:0:0/127",
					Count:       1,
				},
				BgpAttributes: &opb.BgpAttributes{
					Active:                true,
					Origin:                opb.BgpAttributes_ORIGIN_IGP,
					AsnSetMode:            opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
					AdvertisementProtocol: opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_V4,
				},
			}},
		},
		wantNgs: []*ixconfig.TopologyNetworkGroup{{
			Name: ixconfig.String(net1Name),
			Ipv4PrefixPools: []*ixconfig.TopologyIpv4PrefixPools{{
				NetworkAddress:       ixconfig.MultivalueStr("10.0.0.0"),
				PrefixLength:         ixconfig.MultivalueUint32(8),
				NumberOfAddressesAsy: ixconfig.MultivalueUint32(1),
				BgpV6IPRouteProperty: []*ixconfig.TopologyBgpV6IpRouteProperty{{
					Active:                          ixconfig.MultivalueTrue(),
					EnableNextHop:                   ixconfig.MultivalueTrue(),
					EnableOrigin:                    ixconfig.MultivalueTrue(),
					EnableLocalPreference:           ixconfig.MultivalueTrue(),
					LocalPreference:                 ixconfig.MultivalueUint32(0),
					NoOfLargeCommunities:            ixconfig.NumberUint32(0),
					NextHopType:                     ixconfig.MultivalueStr("sameaslocalip"),
					Origin:                          ixconfig.MultivalueStr("igp"),
					EnableCommunity:                 ixconfig.MultivalueFalse(),
					NoOfCommunities:                 ixconfig.NumberInt(0),
					EnableExtendedCommunity:         ixconfig.MultivalueFalse(),
					NoOfExternalCommunities:         ixconfig.NumberInt(0),
					AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
					EnableAsPathSegments:            ixconfig.MultivalueFalse(),
					NoOfASPathSegmentsPerRouteRange: ixconfig.NumberInt(0),
				}},
			}},
		}, {
			Name: ixconfig.String(net2Name),
			Ipv6PrefixPools: []*ixconfig.TopologyIpv6PrefixPools{{
				NetworkAddress:       ixconfig.MultivalueStr("2002:0:0:4860:10::"),
				PrefixLength:         ixconfig.MultivalueUint32(127),
				NumberOfAddressesAsy: ixconfig.MultivalueUint32(1),
				BgpIPRouteProperty: []*ixconfig.TopologyBgpIpRouteProperty{{
					Active:                          ixconfig.MultivalueTrue(),
					EnableNextHop:                   ixconfig.MultivalueTrue(),
					EnableOrigin:                    ixconfig.MultivalueTrue(),
					EnableLocalPreference:           ixconfig.MultivalueTrue(),
					LocalPreference:                 ixconfig.MultivalueUint32(0),
					NoOfLargeCommunities:            ixconfig.NumberUint32(0),
					NextHopType:                     ixconfig.MultivalueStr("sameaslocalip"),
					Origin:                          ixconfig.MultivalueStr("igp"),
					EnableCommunity:                 ixconfig.MultivalueFalse(),
					NoOfCommunities:                 ixconfig.NumberInt(0),
					EnableExtendedCommunity:         ixconfig.MultivalueFalse(),
					NoOfExternalCommunities:         ixconfig.NumberInt(0),
					AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
					EnableAsPathSegments:            ixconfig.MultivalueFalse(),
					NoOfASPathSegmentsPerRouteRange: ixconfig.NumberInt(0),
				}},
			}},
		}},
	}, {
		desc: "IS-IS and BGP pools",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				Ipv4: &opb.NetworkIp{
					AddressCidr: "10.0.0.0/8",
					Count:       1,
				},
				Ipv6: &opb.NetworkIp{
					AddressCidr: "2001::4860:10:0:0:0/104",
					Count:       1,
				},
				Isis: &opb.IPReachability{
					Active:      true,
					RouteOrigin: opb.IPReachability_INTERNAL,
				},
			}, {
				Name:          net2Name,
				InterfaceName: ifName,
				Ipv4: &opb.NetworkIp{
					AddressCidr: "11.0.0.0/8",
					Count:       1,
				},
				Ipv6: &opb.NetworkIp{
					AddressCidr: "2002::4860:10:0:0:0/127",
					Count:       1,
				},
				BgpAttributes: &opb.BgpAttributes{
					Active:                true,
					Origin:                opb.BgpAttributes_ORIGIN_IGP,
					AsnSetMode:            opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
					AdvertisementProtocol: opb.BgpAttributes_ADVERTISEMENT_PROTOCOL_SAME_AS_ROUTE,
				},
			}},
		},
		withISIS: true,
		wantNgs: []*ixconfig.TopologyNetworkGroup{{
			Name: ixconfig.String(net1Name),
			Ipv4PrefixPools: []*ixconfig.TopologyIpv4PrefixPools{{
				Connector:            &ixconfig.TopologyConnector{},
				NetworkAddress:       ixconfig.MultivalueStr("10.0.0.0"),
				PrefixLength:         ixconfig.MultivalueUint32(8),
				NumberOfAddressesAsy: ixconfig.MultivalueUint32(1),
				BgpIPRouteProperty: []*ixconfig.TopologyBgpIpRouteProperty{{
					Name:   ixconfig.String(fmt.Sprintf("%s BGP Inactive", net1Name)),
					Active: ixconfig.MultivalueFalse(),
				}},
				IsisL3RouteProperty: []*ixconfig.TopologyIsisL3RouteProperty{{
					Metric:                 ixconfig.MultivalueUint32(0),
					Algorithm:              ixconfig.MultivalueUint32(0),
					RouteOrigin:            ixconfig.MultivalueStr("internal"),
					Active:                 ixconfig.MultivalueTrue(),
					ConfigureSIDIndexLabel: ixconfig.MultivalueFalse(),
					SIDIndexLabel:          ixconfig.MultivalueUint32(0),
					RFlag:                  ixconfig.MultivalueFalse(),
					NFlag:                  ixconfig.MultivalueFalse(),
					PFlag:                  ixconfig.MultivalueFalse(),
					EFlag:                  ixconfig.MultivalueFalse(),
					VFlag:                  ixconfig.MultivalueFalse(),
					LFlag:                  ixconfig.MultivalueFalse(),
					IsisL3PrefixesSrSid: &ixconfig.TopologyIsisL3PrefixesSrSid{
						Active: ixconfig.MultivalueFalse(),
					},
				}},
			}},
			Ipv6PrefixPools: []*ixconfig.TopologyIpv6PrefixPools{{
				Connector:            &ixconfig.TopologyConnector{},
				NetworkAddress:       ixconfig.MultivalueStr("2001:0:0:4860:10::"),
				PrefixLength:         ixconfig.MultivalueUint32(104),
				NumberOfAddressesAsy: ixconfig.MultivalueUint32(1),
				BgpV6IPRouteProperty: []*ixconfig.TopologyBgpV6IpRouteProperty{{
					Name:   ixconfig.String(fmt.Sprintf("%s BGP V6 Inactive", net1Name)),
					Active: ixconfig.MultivalueFalse(),
				}},
				IsisL3RouteProperty: []*ixconfig.TopologyIsisL3RouteProperty{{
					Metric:                 ixconfig.MultivalueUint32(0),
					Algorithm:              ixconfig.MultivalueUint32(0),
					RouteOrigin:            ixconfig.MultivalueStr("internal"),
					Active:                 ixconfig.MultivalueTrue(),
					ConfigureSIDIndexLabel: ixconfig.MultivalueFalse(),
					SIDIndexLabel:          ixconfig.MultivalueUint32(0),
					RFlag:                  ixconfig.MultivalueFalse(),
					NFlag:                  ixconfig.MultivalueFalse(),
					PFlag:                  ixconfig.MultivalueFalse(),
					EFlag:                  ixconfig.MultivalueFalse(),
					VFlag:                  ixconfig.MultivalueFalse(),
					LFlag:                  ixconfig.MultivalueFalse(),
					IsisL3PrefixesSrSid: &ixconfig.TopologyIsisL3PrefixesSrSid{
						Active: ixconfig.MultivalueFalse(),
					},
				}},
			}},
		}, {
			Name: ixconfig.String(net2Name),
			Ipv4PrefixPools: []*ixconfig.TopologyIpv4PrefixPools{{
				NetworkAddress:       ixconfig.MultivalueStr("11.0.0.0"),
				PrefixLength:         ixconfig.MultivalueUint32(8),
				NumberOfAddressesAsy: ixconfig.MultivalueUint32(1),
				BgpIPRouteProperty: []*ixconfig.TopologyBgpIpRouteProperty{{
					Active:                          ixconfig.MultivalueTrue(),
					EnableNextHop:                   ixconfig.MultivalueTrue(),
					EnableOrigin:                    ixconfig.MultivalueTrue(),
					EnableLocalPreference:           ixconfig.MultivalueTrue(),
					LocalPreference:                 ixconfig.MultivalueUint32(0),
					NoOfLargeCommunities:            ixconfig.NumberUint32(0),
					NextHopType:                     ixconfig.MultivalueStr("sameaslocalip"),
					Origin:                          ixconfig.MultivalueStr("igp"),
					EnableCommunity:                 ixconfig.MultivalueFalse(),
					NoOfCommunities:                 ixconfig.NumberInt(0),
					EnableExtendedCommunity:         ixconfig.MultivalueFalse(),
					NoOfExternalCommunities:         ixconfig.NumberInt(0),
					AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
					EnableAsPathSegments:            ixconfig.MultivalueFalse(),
					NoOfASPathSegmentsPerRouteRange: ixconfig.NumberInt(0),
				}},
				IsisL3RouteProperty: []*ixconfig.TopologyIsisL3RouteProperty{{
					Name:   ixconfig.String(fmt.Sprintf("%s IS-IS Inactive", net2Name)),
					Active: ixconfig.MultivalueFalse(),
				}},
			}},
			Ipv6PrefixPools: []*ixconfig.TopologyIpv6PrefixPools{{
				NetworkAddress:       ixconfig.MultivalueStr("2002:0:0:4860:10::"),
				PrefixLength:         ixconfig.MultivalueUint32(127),
				NumberOfAddressesAsy: ixconfig.MultivalueUint32(1),
				BgpV6IPRouteProperty: []*ixconfig.TopologyBgpV6IpRouteProperty{{
					Active:                          ixconfig.MultivalueTrue(),
					EnableNextHop:                   ixconfig.MultivalueTrue(),
					EnableOrigin:                    ixconfig.MultivalueTrue(),
					EnableLocalPreference:           ixconfig.MultivalueTrue(),
					LocalPreference:                 ixconfig.MultivalueUint32(0),
					NoOfLargeCommunities:            ixconfig.NumberUint32(0),
					NextHopType:                     ixconfig.MultivalueStr("sameaslocalip"),
					Origin:                          ixconfig.MultivalueStr("igp"),
					EnableCommunity:                 ixconfig.MultivalueFalse(),
					NoOfCommunities:                 ixconfig.NumberInt(0),
					EnableExtendedCommunity:         ixconfig.MultivalueFalse(),
					NoOfExternalCommunities:         ixconfig.NumberInt(0),
					AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
					EnableAsPathSegments:            ixconfig.MultivalueFalse(),
					NoOfASPathSegmentsPerRouteRange: ixconfig.NumberInt(0),
				}},
				IsisL3RouteProperty: []*ixconfig.TopologyIsisL3RouteProperty{{
					Name:   ixconfig.String(fmt.Sprintf("%s IS-IS Inactive", net2Name)),
					Active: ixconfig.MultivalueFalse(),
				}},
			}},
		}},
	}, {
		desc: "Importing routes with conflicting config",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:              net1Name,
				InterfaceName:     ifName,
				BgpAttributes:     &opb.BgpAttributes{},
				ImportedBgpRoutes: &opb.Network_ImportedBgpRoutes{},
			}},
		},
		wantErr: true,
	}, {
		desc: "Unspecified imported route format",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:              net1Name,
				InterfaceName:     ifName,
				ImportedBgpRoutes: &opb.Network_ImportedBgpRoutes{},
			}},
		},
		wantErr: true,
	}, {
		desc: "Imported routes without BGP peers",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				ImportedBgpRoutes: &opb.Network_ImportedBgpRoutes{
					RouteTableFormat: opb.Network_ImportedBgpRoutes_ROUTE_TABLE_FORMAT_CISCO,
					Ipv4RoutesPath:   "ipv4",
					Ipv6RoutesPath:   "ipv6",
				},
			}},
		},
		wantErr: true,
	}, {
		desc:          "Imported IPv4 routes w/o v6 peer",
		withBGPv4Peer: true,
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				ImportedBgpRoutes: &opb.Network_ImportedBgpRoutes{
					RouteTableFormat: opb.Network_ImportedBgpRoutes_ROUTE_TABLE_FORMAT_CISCO,
					Ipv4RoutesPath:   "ipv4",
				},
			}},
		},
		wantNgs: []*ixconfig.TopologyNetworkGroup{{
			Name: ixconfig.String(net1Name),
			Ipv4PrefixPools: []*ixconfig.TopologyIpv4PrefixPools{{
				BgpIPRouteProperty: []*ixconfig.TopologyBgpIpRouteProperty{{
					Name:   ixconfig.String(fmt.Sprintf("Imported IPv4 BGP Routes")),
					Active: ixconfig.MultivalueTrue(),
				}},
			}},
		}},
		wantRouteTables: map[string]*routeTables{
			net1Name: &routeTables{
				format: routeTableFormatCisco,
				ipv4:   "ipv4",
			},
		},
	}, {
		desc:          "Imported IPv6 routes w/o v4 peer",
		withBGPv6Peer: true,
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				ImportedBgpRoutes: &opb.Network_ImportedBgpRoutes{
					RouteTableFormat: opb.Network_ImportedBgpRoutes_ROUTE_TABLE_FORMAT_CSV,
					Ipv6RoutesPath:   "ipv6",
				},
			}},
		},
		wantNgs: []*ixconfig.TopologyNetworkGroup{{
			Name: ixconfig.String(net1Name),
			Ipv6PrefixPools: []*ixconfig.TopologyIpv6PrefixPools{{
				BgpV6IPRouteProperty: []*ixconfig.TopologyBgpV6IpRouteProperty{{
					Name:   ixconfig.String(fmt.Sprintf("Imported IPv6 BGP V6 Routes")),
					Active: ixconfig.MultivalueTrue(),
				}},
			}},
		}},
		wantRouteTables: map[string]*routeTables{
			net1Name: &routeTables{
				format: routeTableFormatCSV,
				ipv6:   "ipv6",
			},
		},
	}, {
		desc:          "Full imported route config",
		withBGPv4Peer: true,
		withBGPv6Peer: true,
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Networks: []*opb.Network{{
				Name:          net1Name,
				InterfaceName: ifName,
				ImportedBgpRoutes: &opb.Network_ImportedBgpRoutes{
					RouteTableFormat: opb.Network_ImportedBgpRoutes_ROUTE_TABLE_FORMAT_CISCO,
					Ipv4RoutesPath:   "ipv4",
					Ipv6RoutesPath:   "ipv6",
				},
			}},
		},
		wantNgs: []*ixconfig.TopologyNetworkGroup{{
			Name: ixconfig.String(net1Name),
			Ipv4PrefixPools: []*ixconfig.TopologyIpv4PrefixPools{{
				BgpIPRouteProperty: []*ixconfig.TopologyBgpIpRouteProperty{{
					Name:   ixconfig.String(fmt.Sprintf("Imported IPv4 BGP Routes")),
					Active: ixconfig.MultivalueTrue(),
				}},
				BgpV6IPRouteProperty: []*ixconfig.TopologyBgpV6IpRouteProperty{{
					Name:   ixconfig.String(fmt.Sprintf("Imported IPv4 BGP V6 Routes")),
					Active: ixconfig.MultivalueTrue(),
				}},
			}},
			Ipv6PrefixPools: []*ixconfig.TopologyIpv6PrefixPools{{
				BgpIPRouteProperty: []*ixconfig.TopologyBgpIpRouteProperty{{
					Name:   ixconfig.String(fmt.Sprintf("Imported IPv6 BGP Routes")),
					Active: ixconfig.MultivalueTrue(),
				}},
				BgpV6IPRouteProperty: []*ixconfig.TopologyBgpV6IpRouteProperty{{
					Name:   ixconfig.String(fmt.Sprintf("Imported IPv6 BGP V6 Routes")),
					Active: ixconfig.MultivalueTrue(),
				}},
			}},
		}},
		wantRouteTables: map[string]*routeTables{
			net1Name: &routeTables{
				format: routeTableFormatCisco,
				ipv4:   "ipv4",
				ipv6:   "ipv6",
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := baseClient(test.withISIS, test.withBGPv4Peer, test.withBGPv6Peer)
			gotErr := c.addNetworks(test.ifc)
			if (gotErr != nil) != test.wantErr {
				t.Errorf("addNetworks: unexpected error result, got err: %v, want err? %t", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}

			gotNgs := make(map[string]*ixconfig.TopologyNetworkGroup)
			for _, ng := range c.cfg.Topology[0].DeviceGroup[0].NetworkGroup {
				gotNgs[*(ng.Name)] = ng
			}

			if len(gotNgs) != len(test.wantNgs) {
				t.Fatalf("addNetworks: saw %d configured network groups, wanted %d", len(gotNgs), len(test.wantNgs))
			}

			for _, wantNg := range test.wantNgs {
				ngName := *(wantNg.Name)
				gotNg, ok := gotNgs[ngName]
				if !ok {
					t.Errorf("addNetworks: did not find configured network group for network %q", ngName)
				}
				if diff := jsonCfgDiff(t, wantNg, gotNg); diff != "" {
					t.Errorf("addNetworks: unexpected network group config diff for network %q (-want/+got): %s", ngName, diff)
				}

				if _, ok := c.intfs[ifName].netToNetworkGroup[ngName]; !ok {
					t.Errorf("addNetworks: did not find mapped network group for network %q", ngName)
				}
			}

			if diff := cmp.Diff(test.wantRouteTables, c.intfs[ifName].netToRouteTables, cmp.AllowUnexported(routeTables{}), cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("addNetworks: unexpected diff in route table import map (-want/+got): %s", diff)
			}
		})
	}
}

func TestBgpV4RouteProp(t *testing.T) {
	tests := []struct {
		desc          string
		bgpAttr       *opb.BgpAttributes
		wantRouteProp *ixconfig.TopologyBgpIpRouteProperty
		wantErr       string
	}{{
		desc: "bad community specification",
		bgpAttr: &opb.BgpAttributes{
			Origin: opb.BgpAttributes_ORIGIN_IGP,
			Communities: &opb.BgpCommunities{
				PrivateCommunities: []string{"invalid"},
			},
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
		},
		wantErr: "invalid format for BGP community",
	}, {
		desc: "basic configuration",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
		},
		wantRouteProp: &ixconfig.TopologyBgpIpRouteProperty{
			Active:        ixconfig.MultivalueFalse(),
			EnableNextHop: ixconfig.MultivalueTrue(),
			NextHopType:   ixconfig.MultivalueStr("sameaslocalip"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),
		},
	}, {
		desc: "custom next hop",
		bgpAttr: &opb.BgpAttributes{
			Origin:         opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode:     opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			NextHopAddress: "1.1.1.1",
		},
		wantRouteProp: &ixconfig.TopologyBgpIpRouteProperty{
			Active:               ixconfig.MultivalueFalse(),
			EnableNextHop:        ixconfig.MultivalueTrue(),
			NextHopType:          ixconfig.MultivalueStr("manual"),
			AdvertiseNexthopAsV4: ixconfig.MultivalueTrue(),
			NextHopIPType:        ixconfig.MultivalueStr("ipv4"),
			Ipv4NextHop:          ixconfig.MultivalueStr("1.1.1.1"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),
		},
	}, {
		desc: "custom V6 next hop",
		bgpAttr: &opb.BgpAttributes{
			Origin:         opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode:     opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			NextHopAddress: "aa::1",
		},
		wantRouteProp: &ixconfig.TopologyBgpIpRouteProperty{
			Active:               ixconfig.MultivalueFalse(),
			EnableNextHop:        ixconfig.MultivalueTrue(),
			NextHopType:          ixconfig.MultivalueStr("manual"),
			AdvertiseNexthopAsV4: ixconfig.MultivalueFalse(),
			NextHopIPType:        ixconfig.MultivalueStr("ipv6"),
			Ipv6NextHop:          ixconfig.MultivalueStr("aa::1"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),
		},
	}, {
		desc: "with communities and extended communities",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			Communities: &opb.BgpCommunities{
				PrivateCommunities: []string{"1:2", "3:4"},
			},
			ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
				Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
					Color: &opb.BgpAttributes_ExtendedCommunity_Color{
						CoBits:       opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_10,
						ReservedBits: 1,
						Value:        11,
					},
				},
			}},
		},
		wantRouteProp: &ixconfig.TopologyBgpIpRouteProperty{
			Active:        ixconfig.MultivalueFalse(),
			EnableNextHop: ixconfig.MultivalueTrue(),
			NextHopType:   ixconfig.MultivalueStr("sameaslocalip"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity: ixconfig.MultivalueTrue(),
			NoOfCommunities: ixconfig.NumberUint32(2),
			BgpCommunitiesList: []*ixconfig.TopologyBgpCommunitiesList{{
				AsNumber:      ixconfig.MultivalueUint32(1),
				LastTwoOctets: ixconfig.MultivalueUint32(2),
				Type_:         ixconfig.MultivalueStr("manual"),
			}, {
				AsNumber:      ixconfig.MultivalueUint32(3),
				LastTwoOctets: ixconfig.MultivalueUint32(4),
				Type_:         ixconfig.MultivalueStr("manual"),
			}},
			EnableExtendedCommunity: ixconfig.MultivalueTrue(),
			BgpExtendedCommunitiesList: []*ixconfig.TopologyBgpExtendedCommunitiesList{{
				Type_:             ixconfig.MultivalueStr("opaque"),
				SubType:           ixconfig.MultivalueStr("color"),
				ColorCOBits:       ixconfig.MultivalueStr("10"),
				ColorReservedBits: ixconfig.MultivalueUint32(1),
				ColorValue:        ixconfig.MultivalueUint32(11),
			}},
			NoOfExternalCommunities: ixconfig.NumberUint32(1),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),
		},
	}, {
		desc: "with AS path segments",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
			AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
				Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SET,
				Asns: []uint32{1, 2},
			}},
		},
		wantRouteProp: &ixconfig.TopologyBgpIpRouteProperty{
			Active:        ixconfig.MultivalueFalse(),
			EnableNextHop: ixconfig.MultivalueTrue(),
			NextHopType:   ixconfig.MultivalueStr("sameaslocalip"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("includelocalasasasseq"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(1),
			EnableAsPathSegments:            ixconfig.MultivalueTrue(),
			BgpAsPathSegmentList: []*ixconfig.TopologyBgpAsPathSegmentList{{
				EnableASPathSegment:       ixconfig.MultivalueTrue(),
				SegmentType:               ixconfig.MultivalueStr("asset"),
				NumberOfAsNumberInSegment: ixconfig.NumberUint32(2),
				BgpAsNumberList: []*ixconfig.TopologyBgpAsNumberList{{
					EnableASNumber: ixconfig.MultivalueTrue(),
					AsNumber:       ixconfig.MultivalueUint32(1),
				}, {
					EnableASNumber: ixconfig.MultivalueTrue(),
					AsNumber:       ixconfig.MultivalueUint32(2),
				}},
			}},

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),
		},
	}, {
		desc: "bad originator ID",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			OriginatorId: &opb.StringIncRange{
				Start: "aa::",
				Step:  "0.0.0.1",
			},
		},
		wantErr: "not a valid IPv4 address",
	}, {
		desc: "with originator ID",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			OriginatorId: &opb.StringIncRange{
				Start: "201.0.0.1",
				Step:  "0.0.0.1",
			},
		},
		wantRouteProp: &ixconfig.TopologyBgpIpRouteProperty{
			Active:        ixconfig.MultivalueFalse(),
			EnableNextHop: ixconfig.MultivalueTrue(),
			NextHopType:   ixconfig.MultivalueStr("sameaslocalip"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),

			EnableOriginatorId: ixconfig.MultivalueTrue(),
			OriginatorId:       ixconfig.MultivalueStrIncCounter("201.0.0.1", "0.0.0.1"),
		},
	}, {
		desc: "bad cluster ID",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			ClusterIds: []string{"1.1.1.1", "aa::"},
		},
		wantErr: "not a valid IPv4 address",
	}, {
		desc: "with cluster IDs",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			ClusterIds: []string{"1.1.1.1", "2.2.2.2"},
		},
		wantRouteProp: &ixconfig.TopologyBgpIpRouteProperty{
			Active:        ixconfig.MultivalueFalse(),
			EnableNextHop: ixconfig.MultivalueTrue(),
			NextHopType:   ixconfig.MultivalueStr("sameaslocalip"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),

			EnableCluster: ixconfig.MultivalueTrue(),
			NoOfClusters:  ixconfig.NumberUint32(2),
			BgpClusterIdList: []*ixconfig.TopologyBgpClusterIdList{{
				ClusterId: ixconfig.MultivalueStr("1.1.1.1"),
			}, {
				ClusterId: ixconfig.MultivalueStr("2.2.2.2"),
			}},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotRouteProp, gotErr := bgpV4RouteProp(test.bgpAttr)
			if ((gotErr == nil) != (test.wantErr == "")) || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("bgpV4RouteProp: got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}
			if diff := jsonCfgDiff(t, test.wantRouteProp, gotRouteProp); diff != "" {
				t.Fatalf("bgpV4RouteProp: route property diff at (-want/+got): %s", diff)
			}
		})
	}
}

func TestBgpV6RouteProp(t *testing.T) {
	tests := []struct {
		desc          string
		bgpAttr       *opb.BgpAttributes
		wantRouteProp *ixconfig.TopologyBgpV6IpRouteProperty
		wantErr       string
	}{{
		desc: "bad community specification",
		bgpAttr: &opb.BgpAttributes{
			Origin: opb.BgpAttributes_ORIGIN_IGP,
			Communities: &opb.BgpCommunities{
				PrivateCommunities: []string{"invalid"},
			},
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
		},
		wantErr: "invalid format for BGP community",
	}, {
		desc: "basic configuration",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
		},
		wantRouteProp: &ixconfig.TopologyBgpV6IpRouteProperty{
			Active:        ixconfig.MultivalueFalse(),
			EnableNextHop: ixconfig.MultivalueTrue(),
			NextHopType:   ixconfig.MultivalueStr("sameaslocalip"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),
		},
	}, {
		desc: "custom next hop",
		bgpAttr: &opb.BgpAttributes{
			Origin:         opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode:     opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			NextHopAddress: "aa::",
		},
		wantRouteProp: &ixconfig.TopologyBgpV6IpRouteProperty{
			Active:               ixconfig.MultivalueFalse(),
			EnableNextHop:        ixconfig.MultivalueTrue(),
			NextHopType:          ixconfig.MultivalueStr("manual"),
			AdvertiseNexthopAsV4: ixconfig.MultivalueFalse(),
			Ipv6NextHop:          ixconfig.MultivalueStr("aa::"),
			NextHopIPType:        ixconfig.MultivalueStr("ipv6"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),
		},
	}, {
		desc: "with communities and extended communities",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			Communities: &opb.BgpCommunities{
				PrivateCommunities: []string{"1:2", "3:4"},
			},
			ExtendedCommunities: []*opb.BgpAttributes_ExtendedCommunity{{
				Type: &opb.BgpAttributes_ExtendedCommunity_Color_{
					Color: &opb.BgpAttributes_ExtendedCommunity_Color{
						CoBits:       opb.BgpAttributes_ExtendedCommunity_Color_CO_BITS_10,
						ReservedBits: 1,
						Value:        11,
					},
				},
			}},
		},
		wantRouteProp: &ixconfig.TopologyBgpV6IpRouteProperty{
			Active:        ixconfig.MultivalueFalse(),
			EnableNextHop: ixconfig.MultivalueTrue(),
			NextHopType:   ixconfig.MultivalueStr("sameaslocalip"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity: ixconfig.MultivalueTrue(),
			NoOfCommunities: ixconfig.NumberUint32(2),
			BgpCommunitiesList: []*ixconfig.TopologyBgpCommunitiesList{{
				AsNumber:      ixconfig.MultivalueUint32(1),
				LastTwoOctets: ixconfig.MultivalueUint32(2),
				Type_:         ixconfig.MultivalueStr("manual"),
			}, {
				AsNumber:      ixconfig.MultivalueUint32(3),
				LastTwoOctets: ixconfig.MultivalueUint32(4),
				Type_:         ixconfig.MultivalueStr("manual"),
			}},
			EnableExtendedCommunity: ixconfig.MultivalueTrue(),
			BgpExtendedCommunitiesList: []*ixconfig.TopologyBgpExtendedCommunitiesList{{
				Type_:             ixconfig.MultivalueStr("opaque"),
				SubType:           ixconfig.MultivalueStr("color"),
				ColorCOBits:       ixconfig.MultivalueStr("10"),
				ColorReservedBits: ixconfig.MultivalueUint32(1),
				ColorValue:        ixconfig.MultivalueUint32(11),
			}},
			NoOfExternalCommunities: ixconfig.NumberUint32(1),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),
		},
	}, {
		desc: "with AS path segments",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_AS_SEQ,
			AsPathSegments: []*opb.BgpAttributes_AsPathSegment{{
				Type: opb.BgpAttributes_AsPathSegment_TYPE_AS_SET,
				Asns: []uint32{1, 2},
			}},
		},
		wantRouteProp: &ixconfig.TopologyBgpV6IpRouteProperty{
			Active:        ixconfig.MultivalueFalse(),
			EnableNextHop: ixconfig.MultivalueTrue(),
			NextHopType:   ixconfig.MultivalueStr("sameaslocalip"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("includelocalasasasseq"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(1),
			EnableAsPathSegments:            ixconfig.MultivalueTrue(),
			BgpAsPathSegmentList: []*ixconfig.TopologyBgpAsPathSegmentList{{
				EnableASPathSegment:       ixconfig.MultivalueTrue(),
				SegmentType:               ixconfig.MultivalueStr("asset"),
				NumberOfAsNumberInSegment: ixconfig.NumberUint32(2),
				BgpAsNumberList: []*ixconfig.TopologyBgpAsNumberList{{
					EnableASNumber: ixconfig.MultivalueTrue(),
					AsNumber:       ixconfig.MultivalueUint32(1),
				}, {
					EnableASNumber: ixconfig.MultivalueTrue(),
					AsNumber:       ixconfig.MultivalueUint32(2),
				}},
			}},

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),
		},
	}, {
		desc: "bad originator ID",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			OriginatorId: &opb.StringIncRange{
				Start: "1.1.1.1",
				Step:  "00:00:00:00:00:01",
			},
		},
		wantErr: "not a valid IPv4 address",
	}, {
		desc: "with originator ID",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			OriginatorId: &opb.StringIncRange{
				Start: "201.0.0.1",
				Step:  "0.0.0.1",
			},
		},
		wantRouteProp: &ixconfig.TopologyBgpV6IpRouteProperty{
			Active:        ixconfig.MultivalueFalse(),
			EnableNextHop: ixconfig.MultivalueTrue(),
			NextHopType:   ixconfig.MultivalueStr("sameaslocalip"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),

			EnableOriginatorId: ixconfig.MultivalueTrue(),
			OriginatorId:       ixconfig.MultivalueStrIncCounter("201.0.0.1", "0.0.0.1"),
		},
	}, {
		desc: "bad cluster ID",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			ClusterIds: []string{"1.1.1.1", "aa::"},
		},
		wantErr: "not a valid IPv4 address",
	}, {
		desc: "with cluster IDs",
		bgpAttr: &opb.BgpAttributes{
			Origin:     opb.BgpAttributes_ORIGIN_IGP,
			AsnSetMode: opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
			ClusterIds: []string{"1.1.1.1", "2.2.2.2"},
		},
		wantRouteProp: &ixconfig.TopologyBgpV6IpRouteProperty{
			Active:        ixconfig.MultivalueFalse(),
			EnableNextHop: ixconfig.MultivalueTrue(),
			NextHopType:   ixconfig.MultivalueStr("sameaslocalip"),

			EnableOrigin: ixconfig.MultivalueTrue(),
			Origin:       ixconfig.MultivalueStr("igp"),

			EnableLocalPreference: ixconfig.MultivalueTrue(),
			LocalPreference:       ixconfig.MultivalueUint32(0),

			AsSetMode:                       ixconfig.MultivalueStr("dontincludelocalas"),
			NoOfASPathSegmentsPerRouteRange: ixconfig.NumberUint32(0),
			EnableAsPathSegments:            ixconfig.MultivalueFalse(),

			EnableCommunity:         ixconfig.MultivalueFalse(),
			NoOfCommunities:         ixconfig.NumberUint32(0),
			EnableExtendedCommunity: ixconfig.MultivalueFalse(),
			NoOfExternalCommunities: ixconfig.NumberUint32(0),
			NoOfLargeCommunities:    ixconfig.NumberUint32(0),

			EnableCluster: ixconfig.MultivalueTrue(),
			NoOfClusters:  ixconfig.NumberUint32(2),
			BgpClusterIdList: []*ixconfig.TopologyBgpClusterIdList{{
				ClusterId: ixconfig.MultivalueStr("1.1.1.1"),
			}, {
				ClusterId: ixconfig.MultivalueStr("2.2.2.2"),
			}},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotRouteProp, gotErr := bgpV6RouteProp(test.bgpAttr)
			if ((gotErr == nil) != (test.wantErr == "")) || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("bgpV6RouteProp: got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}
			if diff := jsonCfgDiff(t, test.wantRouteProp, gotRouteProp); diff != "" {
				t.Fatalf("bgpV6RouteProp: route property diff at (-want/+got): %s", diff)
			}
		})
	}
}
