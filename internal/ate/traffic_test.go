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
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ondatra/internal/ixconfig"
	"google.golang.org/protobuf/testing/protocmp"

	opb "github.com/openconfig/ondatra/proto"
)

func singleAddr(addr string) *opb.AddressRange {
	return &opb.AddressRange{Min: addr, Max: addr, Count: 1}
}

func TestInferAddresses(t *testing.T) {
	const (
		ifName   = "someIntf"
		mac      = "aa:bb:cc:dd:ee:ff"
		ipv4Addr = "10.10.10.10"
		ipv6Addr = "aa::"
	)
	intfEPs := []*opb.Flow_Endpoint{{InterfaceName: ifName}}
	intfs := func(ethMac, ipv4Mac, ipv6Mac, ipv4, ipv6 string) map[string]*intf {
		ifc := &intf{
			deviceGroup: &ixconfig.TopologyDeviceGroup{Ethernet: []*ixconfig.TopologyEthernet{{}}},
		}
		if ipv4 != "" {
			ifc.ipv4 = &ixconfig.TopologyIpv4{Address: ixconfig.MultivalueStr(ipv4)}
			ifc.deviceGroup.Ethernet[0].Ipv4 = []*ixconfig.TopologyIpv4{ifc.ipv4}
		}
		if ipv6 != "" {
			ifc.ipv6 = &ixconfig.TopologyIpv6{Address: ixconfig.MultivalueStr(ipv6)}
			ifc.deviceGroup.Ethernet[0].Ipv6 = []*ixconfig.TopologyIpv6{ifc.ipv6}
		}
		ifc.ethMac = ethMac
		ifc.resolvedIpv4Mac = ipv4Mac
		ifc.resolvedIpv6Mac = ipv6Mac
		return map[string]*intf{ifName: ifc}
	}
	tests := []struct {
		desc                                 string
		ethMac, ipv4Mac, ipv6Mac, ipv4, ipv6 string
		hdrs, wantHdrs                       *headers
	}{{
		desc:   "eth src mac",
		ethMac: mac,
		hdrs:   &headers{eth: &opb.EthernetHeader{}},
		wantHdrs: &headers{eth: &opb.EthernetHeader{
			SrcAddr: singleAddr(mac),
		}},
	}, {
		desc:    "resolved ipv4 dst mac for eth",
		ipv4Mac: mac,
		hdrs:    &headers{eth: &opb.EthernetHeader{}},
		wantHdrs: &headers{eth: &opb.EthernetHeader{
			DstAddr: singleAddr(mac),
		}},
	}, {
		desc:    "resolved ipv6 dst mac for eth",
		ipv6Mac: mac,
		hdrs:    &headers{eth: &opb.EthernetHeader{}},
		wantHdrs: &headers{eth: &opb.EthernetHeader{
			DstAddr: singleAddr(mac),
		}},
	}, {
		desc: "ipv4 addrs (same source/destination EP)",
		ipv4: ipv4Addr,
		hdrs: &headers{
			eth:  &opb.EthernetHeader{},
			ipv4: &opb.Ipv4Header{},
		},
		wantHdrs: &headers{
			eth: &opb.EthernetHeader{},
			ipv4: &opb.Ipv4Header{
				SrcAddr: singleAddr(ipv4Addr),
				DstAddr: singleAddr(ipv4Addr),
			},
		},
	}, {
		desc: "ipv6 addrs (same source/destination EP)",
		ipv6: ipv6Addr,
		hdrs: &headers{
			eth:  &opb.EthernetHeader{},
			ipv6: &opb.Ipv6Header{},
		},
		wantHdrs: &headers{
			eth: &opb.EthernetHeader{},
			ipv6: &opb.Ipv6Header{
				SrcAddr: singleAddr(ipv6Addr),
				DstAddr: singleAddr(ipv6Addr),
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			intfs := intfs(test.ethMac, test.ipv4Mac, test.ipv6Mac, test.ipv4, test.ipv6)
			inferAddresses(test.hdrs, intfEPs, intfEPs, intfs)
			if diff := cmp.Diff(test.wantHdrs, test.hdrs, protocmp.Transform(), cmp.AllowUnexported(headers{})); diff != "" {
				t.Errorf("inferAddresses: headers not modified as expected (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAddTraffic(t *testing.T) {
	const (
		ifName     = "someIntf"
		lagIfName  = "lagIntf"
		netGrpName = "someNetGrp"
		rsvpName   = "someRSVPCfg"
		flowName   = "someFlow"
	)
	intfEPs := []*opb.Flow_Endpoint{{InterfaceName: ifName}}
	lagIntfEPs := []*opb.Flow_Endpoint{{InterfaceName: lagIfName}}
	netwEPs := []*opb.Flow_Endpoint{{
		InterfaceName: ifName,
		Generated:     &opb.Flow_Endpoint_NetworkName{NetworkName: netGrpName},
	}}
	lspEPs := []*opb.Flow_Endpoint{{
		InterfaceName: ifName,
		Generated:     &opb.Flow_Endpoint_RsvpName{RsvpName: rsvpName},
	}}
	vportEPs := []string{"/vport[1]/protocols"}
	lagEPs := []string{"/lag[1]"}
	devGrpEPs := []string{"/topology[1]/deviceGroup[1]"}
	netGrpEPs := []string{"/topology[1]/deviceGroup[1]/networkGroup[1]"}
	egressLSPEPs := []string{"/topology[1]/deviceGroup[1]/networkGroup[2]/deviceGroup[1]/ipv4Loopback[1]/rsvpteLsps[1]/rsvpP2PEgressLsps"}
	ingressLSPEPs := []string{"/topology[1]/deviceGroup[1]/networkGroup[2]/deviceGroup[1]/ipv4Loopback[1]/rsvpteLsps[1]/rsvpP2PIngressLsps"}
	baseClient := func() *ixATE {
		cfg := &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
					Ethernet: []*ixconfig.TopologyEthernet{{}},
				}},
			}, {
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
					Ethernet: []*ixconfig.TopologyEthernet{{}},
				}},
			}},
			Vport: []*ixconfig.Vport{{
				Name:      ixconfig.String("vport"),
				Protocols: &ixconfig.VportProtocols{},
			}},
			Lag: []*ixconfig.Lag{{Name: ixconfig.String("lag")}},
		}
		dg := cfg.Topology[0].DeviceGroup[0]
		// Normal network group
		ng := &ixconfig.TopologyNetworkGroup{Name: ixconfig.String(netGrpName)}
		dg.NetworkGroup = append(dg.NetworkGroup, ng)
		// RSVP configuration
		rsvpLSP := &ixconfig.TopologyRsvpteLsps{
			RsvpP2PEgressLsps:  &ixconfig.TopologyRsvpP2PEgressLsps{},
			RsvpP2PIngressLsps: &ixconfig.TopologyRsvpP2PIngressLsps{},
		}
		rsvpNg := &ixconfig.TopologyNetworkGroup{
			DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
				Ipv4Loopback: []*ixconfig.TopologyIpv4Loopback{{
					RsvpteLsps: []*ixconfig.TopologyRsvpteLsps{rsvpLSP},
				}},
			}},
		}
		dg.NetworkGroup = append(dg.NetworkGroup, rsvpNg)
		return &ixATE{
			cfg: cfg,
			intfs: map[string]*intf{
				ifName: &intf{
					deviceGroup: dg,
					netToNetworkGroup: map[string]*ixconfig.TopologyNetworkGroup{
						netGrpName: ng,
					},
					rsvpLSPs: map[string]*ixconfig.TopologyRsvpteLsps{
						rsvpName: rsvpLSP,
					},
					link: cfg.Vport[0],
				},
				lagIfName: &intf{
					deviceGroup: cfg.Topology[1].DeviceGroup[0],
					link:        cfg.Lag[0],
				},
			},
			flowToTrafficItem:     map[string]*ixconfig.TrafficTrafficItem{},
			egressTrackFlowCounts: map[string]int{},
		}
	}
	tests := []struct {
		desc                                    string
		flow                                    *opb.Flow
		wantTrafficType                         trafficType
		wantSrcEPs, wantDstEPs                  []string
		wantStackCount                          int
		wantIngressTracking, wantEgressTracking bool
		wantFrameRateType, wantFrameSizeType    string
		wantCRC                                 string
		wantErr                                 string
	}{{
		desc: "no flow headers specified",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers:      []*opb.Header{},
		},
		wantErr: "at least one header",
	}, {
		desc: "first header is not an ethernet header",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{
				{Type: &opb.Header_Ipv4{&opb.Ipv4Header{}}},
			},
		},
		wantErr: "must be an ethernet header",
	}, {
		desc: "endpoints not set",
		flow: &opb.Flow{
			Name: flowName,
			Headers: []*opb.Header{
				{Type: &opb.Header_Eth{&opb.EthernetHeader{}}},
			},
		},
		wantErr: "endpoints defined",
	}, {
		desc: "bad interface for source endpoint",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: []*opb.Flow_Endpoint{{InterfaceName: "NO SUCH INTF"}},
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{
				{Type: &opb.Header_Eth{&opb.EthernetHeader{}}},
			},
		},
		wantErr: "no interface configured for source",
	}, {
		desc: "bad interface for dest endpoint",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: []*opb.Flow_Endpoint{{InterfaceName: "NO SUCH INTF"}},
			Headers: []*opb.Header{
				{Type: &opb.Header_Eth{&opb.EthernetHeader{}}},
			},
		},
		wantErr: "no interface configured for dest",
	}, {
		desc: "raw traffic flow",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: lagIntfEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
		},
		wantTrafficType: rawTraffic,
		wantSrcEPs:      vportEPs,
		wantDstEPs:      lagEPs,
		wantStackCount:  1,
	}, {
		desc: "ethernet address set for network endpoint",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: netwEPs,
			DstEndpoints: netwEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
		},
		wantErr: "Ethernet header should not be set",
	}, {
		desc: "ethernet traffic flow",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: netwEPs,
			DstEndpoints: netwEPs,
			Headers:      []*opb.Header{{Type: &opb.Header_Eth{&opb.EthernetHeader{}}}},
		},
		wantTrafficType: ethTraffic,
		wantSrcEPs:      netGrpEPs,
		wantDstEPs:      netGrpEPs,
		wantStackCount:  1,
	}, {
		desc: "ipv4 traffic flow",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{
				Type: &opb.Header_Eth{&opb.EthernetHeader{}},
			}, {
				Type: &opb.Header_Ipv4{&opb.Ipv4Header{}},
			}},
		},
		wantTrafficType: ipv4Traffic,
		wantSrcEPs:      devGrpEPs,
		wantDstEPs:      devGrpEPs,
		wantStackCount:  2,
	}, {
		desc: "ipv6 traffic flow",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{
				Type: &opb.Header_Eth{&opb.EthernetHeader{}},
			}, {
				Type: &opb.Header_Ipv6{&opb.Ipv6Header{}},
			}},
		},
		wantTrafficType: ipv6Traffic,
		wantSrcEPs:      devGrpEPs,
		wantDstEPs:      devGrpEPs,
		wantStackCount:  2,
	}, {
		desc: "non-IP lsp traffic flow",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: lspEPs,
			DstEndpoints: lspEPs,
			Headers:      []*opb.Header{{Type: &opb.Header_Eth{&opb.EthernetHeader{}}}},
		},
		wantErr: "cannot use RSVP endpoint",
	}, {
		desc: "lsp traffic flow",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: lspEPs,
			DstEndpoints: lspEPs,
			Headers: []*opb.Header{{
				Type: &opb.Header_Eth{&opb.EthernetHeader{}},
			}, {
				Type: &opb.Header_Mpls{&opb.MplsHeader{}},
			}, {
				Type: &opb.Header_Ipv4{&opb.Ipv4Header{}},
			}},
		},
		wantTrafficType: ipv4Traffic,
		wantSrcEPs:      ingressLSPEPs,
		wantDstEPs:      egressLSPEPs,
		wantStackCount:  3,
	}, {
		desc: "macsec traffic flow",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{
				Type: &opb.Header_Eth{&opb.EthernetHeader{}},
			}, {
				Type: &opb.Header_Macsec{&opb.MacsecHeader{}},
			}, {
				Type: &opb.Header_Ipv4{&opb.Ipv4Header{}},
			}},
		},
		wantTrafficType: ipv4Traffic,
		wantSrcEPs:      devGrpEPs,
		wantDstEPs:      devGrpEPs,
		wantStackCount:  4,
	}, {
		desc: "attempted ingress tracking by endpoint for raw traffic flow",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
			IngressTrackingFilters: &opb.Flow_IngressTrackingFilters{
				DstEndpoint: true,
			},
		},
		wantErr: "ingress tracking by src/dst",
	}, {
		desc: "ingress tracking",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
			IngressTrackingFilters: &opb.Flow_IngressTrackingFilters{
				MplsLabel: true,
			},
		},
		wantTrafficType:     rawTraffic,
		wantSrcEPs:          vportEPs,
		wantDstEPs:          vportEPs,
		wantIngressTracking: true,
		wantStackCount:      1,
	}, {
		desc: "attempted convergence measurements for raw traffic",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
			ConvergenceTracking: true,
		},
		wantErr: "convergence tracking for raw traffic",
	}, {
		desc: "attempted configuration of both ingress tracking and convergence measurements",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers:      []*opb.Header{{Type: &opb.Header_Eth{&opb.EthernetHeader{}}}},
			IngressTrackingFilters: &opb.Flow_IngressTrackingFilters{
				MplsLabel: true,
			},
			ConvergenceTracking: true,
		},
		wantErr: "ingress and convergence tracking tracking for the same flow",
	}, {
		desc: "ingress tracking for convergence measurements",
		flow: &opb.Flow{
			Name:                flowName,
			SrcEndpoints:        netwEPs,
			DstEndpoints:        netwEPs,
			Headers:             []*opb.Header{{Type: &opb.Header_Eth{&opb.EthernetHeader{}}}},
			ConvergenceTracking: true,
		},
		wantTrafficType:     ethTraffic,
		wantSrcEPs:          netGrpEPs,
		wantDstEPs:          netGrpEPs,
		wantIngressTracking: true,
		wantStackCount:      1,
	}, {
		desc: "egress tracking",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
			EgressTracking: &opb.EgressTracking{
				Enabled: true,
				Offset:  128,
				Width:   4,
			},
		},
		wantTrafficType:    rawTraffic,
		wantSrcEPs:         vportEPs,
		wantDstEPs:         vportEPs,
		wantEgressTracking: true,
		wantStackCount:     1,
	}, {
		desc: "egress tracking bad offset",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
			EgressTracking: &opb.EgressTracking{
				Enabled: true,
				Offset:  30,
				Width:   4,
			},
		},
		wantErr: "egress tracking offset",
	}, {
		desc: "egress tracking bad width",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
			EgressTracking: &opb.EgressTracking{
				Enabled: true,
				Offset:  32,
				Width:   23,
			},
		},
		wantErr: "egress tracking width",
	}, {
		desc: "frames by preset distribution and percentage of line rate",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
			FrameRate: &opb.FrameRate{
				Type: &opb.FrameRate_Percent{Percent: 50},
			},
			FrameSize: &opb.FrameSize{
				Type: &opb.FrameSize_ImixPreset_{
					ImixPreset: opb.FrameSize_IMIX_DEFAULT,
				},
			},
		},
		wantTrafficType:   rawTraffic,
		wantSrcEPs:        vportEPs,
		wantDstEPs:        vportEPs,
		wantFrameRateType: "percentLineRate",
		wantFrameSizeType: "presetDistribution",
		wantStackCount:    1,
	}, {
		desc: "fixed size frames by bit rate",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
			FrameRate: &opb.FrameRate{
				Type: &opb.FrameRate_Bps{Bps: 1500},
			},
			FrameSize: &opb.FrameSize{
				Type: &opb.FrameSize_Fixed{
					Fixed: 1500,
				},
			},
		},
		wantTrafficType:   rawTraffic,
		wantSrcEPs:        vportEPs,
		wantDstEPs:        vportEPs,
		wantFrameRateType: "bitsPerSecond",
		wantFrameSizeType: "fixed",
		wantStackCount:    1,
	}, {
		desc: "fixed size frames by frame rate",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{Type: &opb.Header_Eth{
				&opb.EthernetHeader{
					SrcAddr: singleAddr("01:02:03:04:05:06"),
				},
			}}},
			FrameRate: &opb.FrameRate{
				Type: &opb.FrameRate_Fps{Fps: 100},
			},
			FrameSize: &opb.FrameSize{
				Type: &opb.FrameSize_Random_{
					Random: &opb.FrameSize_Random{
						Min: 128,
						Max: 768,
					},
				},
			},
		},
		wantTrafficType:   rawTraffic,
		wantSrcEPs:        vportEPs,
		wantDstEPs:        vportEPs,
		wantFrameRateType: "framesPerSecond",
		wantFrameSizeType: "random",
		wantStackCount:    1,
	}, {
		desc: "bad ethernet CRC",
		flow: &opb.Flow{
			Name:         flowName,
			SrcEndpoints: intfEPs,
			DstEndpoints: intfEPs,
			Headers: []*opb.Header{{
				Type: &opb.Header_Eth{&opb.EthernetHeader{
					BadCrc: true,
				}},
			}},
		},
		wantTrafficType: rawTraffic,
		wantSrcEPs:      vportEPs,
		wantDstEPs:      vportEPs,
		wantCRC:         "badCrc",
		wantStackCount:  1,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := baseClient()
			gotErr := c.addTraffic([]*opb.Flow{test.flow})
			if (gotErr != nil) != (test.wantErr != "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Fatalf("addTraffic: unexpected error got %v, want %s", gotErr, test.wantErr)
			}
			if test.wantErr != "" {
				return
			}

			tiCfg, err := c.cfg.ResolvedConfig(c.cfg.Traffic.TrafficItem[0])
			if err != nil {
				t.Fatalf("could not resolve xpaths/references for traffic config: %v", err)
			}
			ti := tiCfg.(*ixconfig.TrafficTrafficItem)
			if gotTrafficType := *(ti.TrafficType); test.wantTrafficType != trafficType(gotTrafficType) {
				t.Errorf("addTraffic: unexpected traffic type for traffic item: got %q, want %q",
					gotTrafficType, test.wantTrafficType)
			}

			gotSrcEPs := ti.EndpointSet[0].Sources
			if diff := cmp.Diff(test.wantSrcEPs, gotSrcEPs); diff != "" {
				t.Errorf("addTraffic: unexpected source endpoints diff: %s", diff)
			}

			gotDstEPs := ti.EndpointSet[0].Destinations
			if diff := cmp.Diff(test.wantDstEPs, gotDstEPs); diff != "" {
				t.Errorf("addTraffic: unexpected destination endpoints diff: %s", diff)
			}

			if test.wantTrafficType == rawTraffic && (ti.RawTrafficRxPortsBehavior == nil || *(ti.RawTrafficRxPortsBehavior) != "loadBalanced") {
				t.Error("addTraffic: unexpected loadBalanced RxPortsBehavior setting for rawTraffic: got: false, want: true")
			}

			// The len check is not '> 0' because TrackBy always includes at least one element to enable tracking.
			if gotIngressTracking := len(ti.Tracking[0].TrackBy) > 1; test.wantIngressTracking != gotIngressTracking {
				t.Errorf("addTraffic: unexpected ingress tracking state: got enabled: %t, want enabled: %t",
					gotIngressTracking, test.wantIngressTracking)
			}

			if gotEgressTracking := (ti.EgressEnabled != nil && *(ti.EgressEnabled)) && len(ti.EgressTracking) > 0; test.wantEgressTracking != gotEgressTracking {
				t.Errorf("addTraffic: unexpected egress tracking state: got enabled: %t, want enabled: %t",
					gotEgressTracking, test.wantEgressTracking)
			}

			if gotStackCount := len(ti.ConfigElement[0].Stack); test.wantStackCount != gotStackCount {
				t.Errorf("addTraffic: unexpected stack count for traffic item: got %d, want %d",
					gotStackCount, test.wantStackCount)
			}

			var gotFrameRateType string
			fr := ti.ConfigElement[0].FrameRate
			if fr != nil {
				gotFrameRateType = *(fr.Type_)
			}
			if test.wantFrameRateType != gotFrameRateType {
				t.Errorf("addTraffic: unexpected frame rate type: got %q, want %q",
					gotFrameRateType, test.wantFrameRateType)
			}

			var gotFrameSizeType string
			fs := ti.ConfigElement[0].FrameSize
			if fs != nil {
				gotFrameSizeType = *(fs.Type_)
			}
			if test.wantFrameSizeType != gotFrameSizeType {
				t.Errorf("addTraffic: unexpected frame size type: got %q, want %q",
					gotFrameSizeType, test.wantFrameSizeType)
			}

			var gotCRC string
			crc := ti.ConfigElement[0].Crc
			if crc != nil {
				gotCRC = *crc
			}
			if test.wantCRC != gotCRC {
				t.Errorf("addTraffic: unexpected CRC: got %q, want %q", gotCRC, test.wantCRC)
			}
		})
	}
}

func TestFrameRate(t *testing.T) {
	tests := []struct {
		desc             string
		frameRatePB      *opb.FrameRate
		wantFrameRate    *ixconfig.TrafficTrafficItemConfigElementFrameRate
		wantFrameRateMap map[string]any
	}{{
		desc: "no frame rate set",
	}, {
		desc: "percent line rate",
		frameRatePB: &opb.FrameRate{
			Type: &opb.FrameRate_Percent{Percent: 50},
		},
		wantFrameRate: &ixconfig.TrafficTrafficItemConfigElementFrameRate{
			Type_: ixconfig.String("percentLineRate"),
			Rate:  ixconfig.NumberFloat64(50),
		},
		wantFrameRateMap: map[string]any{
			"rate": 50.0,
		},
	}, {
		desc: "bits per second",
		frameRatePB: &opb.FrameRate{
			Type: &opb.FrameRate_Bps{Bps: 1024},
		},
		wantFrameRate: &ixconfig.TrafficTrafficItemConfigElementFrameRate{
			Type_:            ixconfig.String("bitsPerSecond"),
			BitRateUnitsType: ixconfig.String("bitsPerSec"),
			Rate:             ixconfig.NumberUint64(1024),
		},
		wantFrameRateMap: map[string]any{
			"bitRateUnitsType": "bitsPerSec",
			"rate":             uint64(1024),
		},
	}, {
		desc: "frames per second",
		frameRatePB: &opb.FrameRate{
			Type: &opb.FrameRate_Fps{Fps: 1000},
		},
		wantFrameRate: &ixconfig.TrafficTrafficItemConfigElementFrameRate{
			Type_: ixconfig.String("framesPerSecond"),
			Rate:  ixconfig.NumberUint64(1000),
		},
		wantFrameRateMap: map[string]any{
			"rate": uint64(1000),
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotFrameRate, gotFrameRateMap, gotErr := frameRate(test.frameRatePB)
			if gotErr != nil {
				t.Fatalf("frameRate: unexpected err: %v", gotErr)
			}
			if diff := jsonCfgDiff(t, test.wantFrameRate, gotFrameRate); diff != "" {
				t.Errorf("frameRate: unexpected TrafficTrafficItemConfigElementFrameRate diff (-want/+got): %s", diff)
			}
			if diff := cmp.Diff(test.wantFrameRateMap, gotFrameRateMap); diff != "" {
				t.Errorf("frameRate: unexpected frame rate data map (-want/+got): %s", diff)
			}
		})
	}
}

func TestFrameSize(t *testing.T) {
	tests := []struct {
		desc             string
		frameSizePB      *opb.FrameSize
		wantFrameSize    *ixconfig.TrafficTrafficItemConfigElementFrameSize
		wantFrameSizeMap map[string]any
		wantErr          string
	}{{
		desc: "no frame size set",
	}, {
		desc: "invalid preset",
		frameSizePB: &opb.FrameSize{
			Type: &opb.FrameSize_ImixPreset_{
				ImixPreset: opb.FrameSize_IMIX_UNKNOWN,
			},
		},
		wantErr: "invalid preset",
	}, {
		desc: "valid preset",
		frameSizePB: &opb.FrameSize{
			Type: &opb.FrameSize_ImixPreset_{
				ImixPreset: opb.FrameSize_IMIX_DEFAULT,
			},
		},
		wantFrameSize: &ixconfig.TrafficTrafficItemConfigElementFrameSize{
			Type_:              ixconfig.String("presetDistribution"),
			PresetDistribution: ixconfig.String(`"imix"`),
			QuadGaussian:       []float32{},
			WeightedPairs:      []float32{},
		},
		wantFrameSizeMap: map[string]any{
			"presetDistribution": `"imix"`,
		},
	}, {
		desc:        "fixed frame size",
		frameSizePB: &opb.FrameSize{Type: &opb.FrameSize_Fixed{Fixed: 1500}},
		wantFrameSize: &ixconfig.TrafficTrafficItemConfigElementFrameSize{
			Type_:         ixconfig.String("fixed"),
			FixedSize:     ixconfig.NumberUint32(1500),
			QuadGaussian:  []float32{},
			WeightedPairs: []float32{},
		},
		wantFrameSizeMap: map[string]any{
			"fixedSize": uint32(1500),
		},
	}, {
		desc: "random frame size",
		frameSizePB: &opb.FrameSize{
			Type: &opb.FrameSize_Random_{
				Random: &opb.FrameSize_Random{Min: 64, Max: 512},
			},
		},
		wantFrameSize: &ixconfig.TrafficTrafficItemConfigElementFrameSize{
			Type_:         ixconfig.String("random"),
			RandomMin:     ixconfig.NumberUint32(64),
			RandomMax:     ixconfig.NumberUint32(512),
			QuadGaussian:  []float32{},
			WeightedPairs: []float32{},
		},
		wantFrameSizeMap: map[string]any{
			"randomMin": uint32(64),
			"randomMax": uint32(512),
		},
	}, {
		desc: "custom frame sizes as weighted pairs",
		frameSizePB: &opb.FrameSize{
			Type: &opb.FrameSize_ImixCustom_{
				ImixCustom: &opb.FrameSize_ImixCustom{
					Entries: []*opb.FrameSize_ImixCustomEntry{{
						Size:   82,
						Weight: 35,
					}, {
						Size:   125,
						Weight: 22,
					}},
				},
			},
		},
		wantFrameSize: &ixconfig.TrafficTrafficItemConfigElementFrameSize{
			Type_:         ixconfig.String("weightedPairs"),
			QuadGaussian:  []float32{},
			WeightedPairs: []float32{82, 35, 125, 22},
		},
		wantFrameSizeMap: map[string]any{
			"weightedPairs": []float32{82, 35, 125, 22},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotFrameSize, gotFrameSizeMap, gotErr := frameSize(test.frameSizePB)
			if (gotErr == nil && test.wantErr != "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("frameSize: got err: %v, want err %q", gotErr, test.wantErr)
			}
			if diff := jsonCfgDiff(t, test.wantFrameSize, gotFrameSize); diff != "" {
				t.Errorf("frameSize: unexpected TrafficTrafficItemConfigElementFrameSize diff (-want/+got): %s", diff)
			}
			if diff := cmp.Diff(test.wantFrameSizeMap, gotFrameSizeMap); diff != "" {
				t.Errorf("frameSize: unexpected frame size data map (-want/+got): %s", diff)
			}
		})
	}
}

func TestTransmissionControl(t *testing.T) {
	tests := []struct {
		desc             string
		transmissionPB   *opb.Transmission
		wantTransmission *ixconfig.TrafficTrafficItemConfigElementTransmissionControl
		wantErr          string
	}{{
		desc: "default",
	}, {
		desc: "invalid pattern",
		transmissionPB: &opb.Transmission{
			Pattern: opb.Transmission_PATTERN_UNSPECIFIED,
		},
		wantErr: "unrecognized transmission pattern",
	}, {
		desc: "burst packet count set for continuous",
		transmissionPB: &opb.Transmission{
			Pattern:         opb.Transmission_CONTINUOUS,
			PacketsPerBurst: 1,
		},
		wantErr: "burst packet count should not be set",
	}, {
		desc: "burst gap set for continuous",
		transmissionPB: &opb.Transmission{
			Pattern:       opb.Transmission_CONTINUOUS,
			InterburstGap: &opb.Transmission_Bytes{Bytes: 64},
		},
		wantErr: "burst gap should not be set",
	}, {
		desc: "continuous",
		transmissionPB: &opb.Transmission{
			Pattern:     opb.Transmission_CONTINUOUS,
			MinGapBytes: 64,
		},
		wantTransmission: &ixconfig.TrafficTrafficItemConfigElementTransmissionControl{
			Type_:       ixconfig.String("continuous"),
			MinGapBytes: ixconfig.NumberUint32(64),
		},
	}, {
		desc: "packet count not set for burst config",
		transmissionPB: &opb.Transmission{
			Pattern: opb.Transmission_BURST,
		},
		wantErr: "burst packet count must be a positive value",
	}, {
		desc: "burst gap not set for burst config",
		transmissionPB: &opb.Transmission{
			Pattern:         opb.Transmission_BURST,
			PacketsPerBurst: 100,
		},
		wantErr: "burst gap must be set",
	}, {
		desc: "burst gap bytes",
		transmissionPB: &opb.Transmission{
			Pattern:         opb.Transmission_BURST,
			PacketsPerBurst: 100,
			InterburstGap:   &opb.Transmission_Bytes{Bytes: 64},
		},
		wantTransmission: &ixconfig.TrafficTrafficItemConfigElementTransmissionControl{
			Type_:               ixconfig.String("custom"),
			EnableInterBurstGap: ixconfig.Bool(true),
			MinGapBytes:         ixconfig.NumberUint32(0),
			BurstPacketCount:    ixconfig.NumberUint32(100),
			InterBurstGapUnits:  ixconfig.String("bytes"),
			InterBurstGap:       ixconfig.NumberUint32(64),
		},
	}, {
		desc: "burst gap nanoseconds",
		transmissionPB: &opb.Transmission{
			Pattern:         opb.Transmission_BURST,
			PacketsPerBurst: 100,
			InterburstGap:   &opb.Transmission_Nanoseconds{Nanoseconds: 30},
		},
		wantTransmission: &ixconfig.TrafficTrafficItemConfigElementTransmissionControl{
			Type_:               ixconfig.String("custom"),
			EnableInterBurstGap: ixconfig.Bool(true),
			MinGapBytes:         ixconfig.NumberUint32(0),
			BurstPacketCount:    ixconfig.NumberUint32(100),
			InterBurstGapUnits:  ixconfig.String("nanoseconds"),
			InterBurstGap:       ixconfig.NumberUint32(30),
		},
	}, {
		desc: "burst packet count set for fixed frame count",
		transmissionPB: &opb.Transmission{
			Pattern:         opb.Transmission_FIXED_FRAME_COUNT,
			PacketsPerBurst: 1,
		},
		wantErr: "burst packet count should not be set",
	}, {
		desc: "burst gap set for fixed fixed frame count",
		transmissionPB: &opb.Transmission{
			Pattern:       opb.Transmission_FIXED_FRAME_COUNT,
			InterburstGap: &opb.Transmission_Bytes{Bytes: 64},
		},
		wantErr: "burst gap should not be set",
	}, {
		desc: "burst packet count set for duration transmission",
		transmissionPB: &opb.Transmission{
			Pattern:         opb.Transmission_FIXED_DURATION,
			PacketsPerBurst: 1,
		},
		wantErr: "burst packet count should not be set",
	}, {
		desc: "burst gap set for duration transmission",
		transmissionPB: &opb.Transmission{
			Pattern:       opb.Transmission_FIXED_DURATION,
			InterburstGap: &opb.Transmission_Bytes{Bytes: 64},
		},
		wantErr: "burst gap should not be set",
	}, {
		desc: "fixed frame count",
		transmissionPB: &opb.Transmission{
			Pattern:     opb.Transmission_FIXED_FRAME_COUNT,
			MinGapBytes: 64,
			FrameCount:  1000,
		},
		wantTransmission: &ixconfig.TrafficTrafficItemConfigElementTransmissionControl{
			Type_:       ixconfig.String("fixedFrameCount"),
			MinGapBytes: ixconfig.NumberUint32(64),
			FrameCount:  ixconfig.NumberUint32(1000),
		},
	}, {
		desc: "fixed duration",
		transmissionPB: &opb.Transmission{
			Pattern:      opb.Transmission_FIXED_DURATION,
			MinGapBytes:  64,
			DurationSecs: 100,
		},
		wantTransmission: &ixconfig.TrafficTrafficItemConfigElementTransmissionControl{
			Type_:       ixconfig.String("fixedDuration"),
			MinGapBytes: ixconfig.NumberUint32(64),
			Duration:    ixconfig.NumberUint32(100),
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotTransmission, gotErr := transmissionControl(test.transmissionPB)
			if (gotErr == nil && test.wantErr != "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("transmissionControl: got err: %v, want err %q", gotErr, test.wantErr)
			}
			if diff := jsonCfgDiff(t, test.wantTransmission, gotTransmission); diff != "" {
				t.Errorf("transmissionControl: unexpected TrafficTrafficItemConfigElementTransmissionControl diff (-want/+got): %s", diff)
			}
		})
	}
}
