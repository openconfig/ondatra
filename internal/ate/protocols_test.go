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

	"github.com/openconfig/ondatra/internal/ixconfig"

	opb "github.com/openconfig/ondatra/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func clientWithTopoCfg(ifName string) *ixATE {
	cfg := &ixconfig.Ixnetwork{
		Topology: []*ixconfig.Topology{{
			DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
				Ethernet: []*ixconfig.TopologyEthernet{{}},
			}},
		}},
	}
	return &ixATE{
		cfg: cfg,
		intfs: map[string]*intf{
			ifName: &intf{deviceGroup: cfg.Topology[0].DeviceGroup[0]},
		},
	}
}

func TestMKA(t *testing.T) {
	tests := []struct {
		desc    string
		mka     *opb.MacSec_MKA
		wantMKA *ixconfig.TopologyMka
		wantErr string
	}{{
		desc: "cipher suite not specified",
		mka: &opb.MacSec_MKA{
			Version: 1,
		},
		wantErr: "cipher suite not specified",
	}, {
		desc: "confidentiality offset not specified",
		mka: &opb.MacSec_MKA{
			Version:     1,
			CipherSuite: opb.MacSec_AES_128,
		},
		wantErr: "confidentiality offset not specified",
	}, {
		desc: "capability not specified",
		mka: &opb.MacSec_MKA{
			Version:               1,
			CipherSuite:           opb.MacSec_AES_128,
			ConfidentialityOffset: opb.MacSec_MKA_OFFSET_0,
		},
		wantErr: "capability not specified",
	}, {
		desc: "CAK name is not hex",
		mka: &opb.MacSec_MKA{
			Version:               1,
			CipherSuite:           opb.MacSec_AES_128,
			ConfidentialityOffset: opb.MacSec_MKA_OFFSET_0,
			Capability:            opb.MacSec_MKA_NOT_IMPLEMENTED,
			ConnectivityAssociation: &opb.MacSec_MKA_ConnectivityAssociation{
				Ckn: "invalid",
			},
		},
		wantErr: "could not parse",
	}, {
		desc: "CAK name is too long",
		mka: &opb.MacSec_MKA{
			Version:               1,
			CipherSuite:           opb.MacSec_AES_128,
			ConfidentialityOffset: opb.MacSec_MKA_OFFSET_0,
			Capability:            opb.MacSec_MKA_NOT_IMPLEMENTED,
			ConnectivityAssociation: &opb.MacSec_MKA_ConnectivityAssociation{
				Ckn: "0123456789abcdef0123456789abcdef0000",
			},
		},
		wantErr: "exceeds limit of 32 characters",
	}, {
		desc: "CAK value > 128 bits for 128 bit cipher",
		mka: &opb.MacSec_MKA{
			Version:               1,
			CipherSuite:           opb.MacSec_AES_128,
			ConfidentialityOffset: opb.MacSec_MKA_OFFSET_0,
			Capability:            opb.MacSec_MKA_NOT_IMPLEMENTED,
			ConnectivityAssociation: &opb.MacSec_MKA_ConnectivityAssociation{
				Ckn: "aaaa",
				Cak: "0123456789abcdef0123456789abcdef0000",
			},
		},
		wantErr: "exceeds limit of 32 characters",
	}, {
		desc: "CAK value > 256 bits for 256 bit cipher",
		mka: &opb.MacSec_MKA{
			Version:               1,
			CipherSuite:           opb.MacSec_AES_256,
			ConfidentialityOffset: opb.MacSec_MKA_OFFSET_0,
			Capability:            opb.MacSec_MKA_NOT_IMPLEMENTED,
			ConnectivityAssociation: &opb.MacSec_MKA_ConnectivityAssociation{
				Ckn: "aaaa",
				Cak: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0000",
			},
		},
		wantErr: "exceeds limit of 64 characters",
	}, {
		desc: "valid MKA config, 128 bit cipher",
		mka: &opb.MacSec_MKA{
			Version:               1,
			CipherSuite:           opb.MacSec_AES_128,
			ConfidentialityOffset: opb.MacSec_MKA_OFFSET_0,
			Capability:            opb.MacSec_MKA_NOT_IMPLEMENTED,
			ConnectivityAssociation: &opb.MacSec_MKA_ConnectivityAssociation{
				Ckn: "aaaa",
				Cak: "bbbb",
			},
		},
		wantMKA: &ixconfig.TopologyMka{
			MkaVersion:            ixconfig.MultivalueUint32(1),
			CipherSuite:           ixconfig.MultivalueStr("aes128"),
			ConfidentialityOffset: ixconfig.MultivalueStr("noconfidentialityoffset"),
			MacsecCapability:      ixconfig.MultivalueStr("macsecnotimplemented"),
			CakCache: &ixconfig.TopologyCakCache{
				CakName:     ixconfig.MultivalueStr("aaaa"),
				CakValue128: ixconfig.MultivalueStr("bbbb"),
			},
		},
	}, {
		desc: "valid MKA config, 256 bit cipher",
		mka: &opb.MacSec_MKA{
			Version:               1,
			CipherSuite:           opb.MacSec_AES_256,
			ConfidentialityOffset: opb.MacSec_MKA_OFFSET_0,
			Capability:            opb.MacSec_MKA_NOT_IMPLEMENTED,
			ConnectivityAssociation: &opb.MacSec_MKA_ConnectivityAssociation{
				Ckn: "aaaa",
				Cak: "bbbb",
			},
		},
		wantMKA: &ixconfig.TopologyMka{
			MkaVersion:            ixconfig.MultivalueUint32(1),
			CipherSuite:           ixconfig.MultivalueStr("aes256"),
			ConfidentialityOffset: ixconfig.MultivalueStr("noconfidentialityoffset"),
			MacsecCapability:      ixconfig.MultivalueStr("macsecnotimplemented"),
			CakCache: &ixconfig.TopologyCakCache{
				CakName:     ixconfig.MultivalueStr("aaaa"),
				CakValue256: ixconfig.MultivalueStr("bbbb"),
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotMKA, gotErr := mka(test.mka)
			if (gotErr == nil && test.wantErr != "") || (gotErr != nil && test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("mka got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}
			if diff := jsonCfgDiff(t, test.wantMKA, gotMKA); diff != "" {
				t.Fatalf("unexpected MKA diff (-want/+got): %s", diff)
			}
		})
	}
}

func TestAddMACsecProtocol(t *testing.T) {
	const ifName = "someIntf"
	tests := []struct {
		desc       string
		macsec     *opb.MacSec
		wantMACsec *ixconfig.TopologyMacsec
		wantMKA    *ixconfig.TopologyMka
		wantErr    string
	}{{
		desc: "No macsec",
	}, {
		desc: "No SAK pool",
		macsec: &opb.MacSec{
			CipherSuite:           opb.MacSec_AES_256,
			EncryptedVlansEnabled: true,
		},
		wantMACsec: &ixconfig.TopologyMacsec{
			Name:                 ixconfig.String("MACsec on someIntf"),
			CipherSuite:          ixconfig.MultivalueStr("aes256"),
			EnableEncryptedVlans: ixconfig.MultivalueBool(true),
		},
	}, {
		desc: "With SAK pool",
		macsec: &opb.MacSec{
			CipherSuite: opb.MacSec_AES_XPN_256,
			RxSakPool: &opb.RxSakPool{
				Sak128: "0123456789abcdef0123456789abcdef",
				Sak256: "0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF",
			},
		},
		wantMACsec: &ixconfig.TopologyMacsec{
			Name:                 ixconfig.String("MACsec on someIntf"),
			CipherSuite:          ixconfig.MultivalueStr("aesxpn256"),
			EnableEncryptedVlans: ixconfig.MultivalueBool(false),
			RxSakPool: &ixconfig.TopologyRxSakPool{
				RxSak128: ixconfig.MultivalueStr("0123456789abcdef0123456789abcdef"),
				RxSak256: ixconfig.MultivalueStr("0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF"),
			},
		},
	}, {
		desc: "With MKA",
		macsec: &opb.MacSec{
			CipherSuite: opb.MacSec_AES_XPN_256,
			Mka: &opb.MacSec_MKA{
				Version:               1,
				Capability:            opb.MacSec_MKA_NOT_IMPLEMENTED,
				CipherSuite:           opb.MacSec_AES_XPN_256,
				ConfidentialityOffset: opb.MacSec_MKA_OFFSET_0,
			},
		},
		wantMACsec: &ixconfig.TopologyMacsec{
			Name:                 ixconfig.String("MACsec on someIntf"),
			CipherSuite:          ixconfig.MultivalueStr("aesxpn256"),
			EnableEncryptedVlans: ixconfig.MultivalueBool(false),
		},
		wantMKA: &ixconfig.TopologyMka{
			MkaVersion:            ixconfig.MultivalueUint32(1),
			CipherSuite:           ixconfig.MultivalueStr("aesxpn256"),
			ConfidentialityOffset: ixconfig.MultivalueStr("noconfidentialityoffset"),
			MacsecCapability:      ixconfig.MultivalueStr("macsecnotimplemented"),
			CakCache:              &ixconfig.TopologyCakCache{},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := clientWithTopoCfg(ifName)
			ifc := &opb.InterfaceConfig{
				Name:     ifName,
				Ethernet: &opb.EthernetConfig{Macsec: test.macsec},
			}
			gotErr := c.addMACsecProtocol(ifc)
			if (gotErr == nil && test.wantErr != "") || (gotErr != nil && test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("addMACsecProtocol got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}
			gotMACsecs := c.cfg.Topology[0].DeviceGroup[0].Ethernet[0].Macsec
			if test.wantMACsec == nil {
				if len(gotMACsecs) != 0 {
					t.Fatalf("addMACsecProtocol want no MACsec configurations, got %v", len(gotMACsecs))
				}
				return
			}
			if len(gotMACsecs) != 1 {
				t.Fatalf("addMACsecProtocol want one MACsec configuration, got %v", len(gotMACsecs))
			}
			gotMACsec := gotMACsecs[0]
			if diff := jsonCfgDiff(t, test.wantMACsec, gotMACsec); diff != "" {
				t.Fatalf("addMACsecProtocol unexpected MACsec diff (-want/+got): %s", diff)
			}

			gotMKAs := c.cfg.Topology[0].DeviceGroup[0].Ethernet[0].Mka
			if test.wantMKA == nil {
				if len(gotMKAs) != 0 {
					t.Fatalf("addMACsecProtocol want no MKA configurations, got %v", len(gotMKAs))
				}
				return
			}
			if len(gotMKAs) != 1 {
				t.Fatalf("addMACsecProtocol want one MKA configuration, got %v", len(gotMKAs))
			}
			gotMKA := gotMKAs[0]
			if diff := jsonCfgDiff(t, test.wantMKA, gotMKA); diff != "" {
				t.Fatalf("addMACsecProtocol unexpected MKA diff (-want/+got): %s", diff)
			}
		})
	}
}

func TestAddIpProtocols(t *testing.T) {
	const ifName = "someIntf"

	clientWithTopo := func(ifName string, hasMacSec bool) *ixATE {
		eth := &ixconfig.TopologyEthernet{}
		if hasMacSec {
			eth.Macsec = []*ixconfig.TopologyMacsec{{}}
		}
		cfg := &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{Ethernet: []*ixconfig.TopologyEthernet{eth}}},
			}},
		}
		return &ixATE{
			cfg: cfg,
			intfs: map[string]*intf{
				ifName: &intf{deviceGroup: cfg.Topology[0].DeviceGroup[0]},
			},
		}
	}

	tests := []struct {
		desc                       string
		ifcfg                      *opb.InterfaceConfig
		hasMacSec                  bool
		wantErr                    string
		wantCfg                    *ixconfig.TopologyEthernet
		wantIpv4Name, wantIpv6Name string
	}{{
		desc: "Incomplete IPv4 information",
		ifcfg: &opb.InterfaceConfig{
			Name: ifName,
			Ipv4: &opb.IpConfig{},
		},
		wantErr: "need both address and default gateway defined for IP V4 layer",
	}, {
		desc: "Incomplete IPv6 information",
		ifcfg: &opb.InterfaceConfig{
			Name: ifName,
			Ipv6: &opb.IpConfig{},
		},
		wantErr: "need both address and default gateway defined for IP V6 layer",
	}, {
		desc: "Valid config",
		ifcfg: &opb.InterfaceConfig{
			Name: ifName,
			Ipv4: &opb.IpConfig{
				AddressCidr:    "192.168.1.1/30",
				DefaultGateway: "192.168.1.2",
			},
			Ipv6: &opb.IpConfig{
				AddressCidr:    "2001::4860:193:168:1:1/126",
				DefaultGateway: "2001::4860:193:168:1:2",
			},
		},
		wantCfg: &ixconfig.TopologyEthernet{
			Ipv4: []*ixconfig.TopologyIpv4{{
				Name:           ixconfig.String(fmt.Sprintf("IPv4 on %s", ifName)),
				Address:        ixconfig.MultivalueStr("192.168.1.1"),
				GatewayIp:      ixconfig.MultivalueStr("192.168.1.2"),
				Prefix:         ixconfig.MultivalueUint32(30),
				ResolveGateway: ixconfig.MultivalueTrue(),
			}},
			Ipv6: []*ixconfig.TopologyIpv6{{
				Name:           ixconfig.String(fmt.Sprintf("IPv6 on %s", ifName)),
				Address:        ixconfig.MultivalueStr("2001::4860:193:168:1:1"),
				GatewayIp:      ixconfig.MultivalueStr("2001::4860:193:168:1:2"),
				Prefix:         ixconfig.MultivalueUint32(126),
				ResolveGateway: ixconfig.MultivalueTrue(),
			}},
		},
		wantIpv4Name: fmt.Sprintf("IPv4 on %s", ifName),
		wantIpv6Name: fmt.Sprintf("IPv6 on %s", ifName),
	}, {
		desc: "Valid config - unspecified gateway",
		ifcfg: &opb.InterfaceConfig{
			Name: ifName,
			Ipv4: &opb.IpConfig{
				AddressCidr:    "192.168.1.1/30",
				DefaultGateway: "0.0.0.0",
			},
			Ipv6: &opb.IpConfig{
				AddressCidr:    "2001::4860:193:168:1:1/126",
				DefaultGateway: "::",
			},
		},
		wantCfg: &ixconfig.TopologyEthernet{
			Ipv4: []*ixconfig.TopologyIpv4{{
				Name:           ixconfig.String(fmt.Sprintf("IPv4 on %s", ifName)),
				Address:        ixconfig.MultivalueStr("192.168.1.1"),
				GatewayIp:      ixconfig.MultivalueStr("0.0.0.0"),
				Prefix:         ixconfig.MultivalueUint32(30),
				ResolveGateway: ixconfig.MultivalueTrue(),
			}},
			Ipv6: []*ixconfig.TopologyIpv6{{
				Name:           ixconfig.String(fmt.Sprintf("IPv6 on %s", ifName)),
				Address:        ixconfig.MultivalueStr("2001::4860:193:168:1:1"),
				GatewayIp:      ixconfig.MultivalueStr("::"),
				Prefix:         ixconfig.MultivalueUint32(126),
				ResolveGateway: ixconfig.MultivalueTrue(),
			}},
		},
		wantIpv4Name: fmt.Sprintf("IPv4 on %s", ifName),
		wantIpv6Name: fmt.Sprintf("IPv6 on %s", ifName),
	}, {
		desc: "Valid config on Macsec",
		ifcfg: &opb.InterfaceConfig{
			Name: ifName,
			Ipv4: &opb.IpConfig{
				AddressCidr:    "192.168.1.1/30",
				DefaultGateway: "192.168.1.2",
			},
			Ipv6: &opb.IpConfig{
				AddressCidr:    "2001::4860:193:168:1:1/126",
				DefaultGateway: "2001::4860:193:168:1:2",
			},
		},
		hasMacSec: true,
		wantCfg: &ixconfig.TopologyEthernet{
			Macsec: []*ixconfig.TopologyMacsec{{
				Ipv4: []*ixconfig.TopologyIpv4{{
					Name:           ixconfig.String(fmt.Sprintf("IPv4 on %s", ifName)),
					Address:        ixconfig.MultivalueStr("192.168.1.1"),
					GatewayIp:      ixconfig.MultivalueStr("192.168.1.2"),
					Prefix:         ixconfig.MultivalueUint32(30),
					ResolveGateway: ixconfig.MultivalueTrue(),
				}},
				Ipv6: []*ixconfig.TopologyIpv6{{
					Name:           ixconfig.String(fmt.Sprintf("IPv6 on %s", ifName)),
					Address:        ixconfig.MultivalueStr("2001::4860:193:168:1:1"),
					GatewayIp:      ixconfig.MultivalueStr("2001::4860:193:168:1:2"),
					Prefix:         ixconfig.MultivalueUint32(126),
					ResolveGateway: ixconfig.MultivalueTrue(),
				}},
			}},
		},
		wantIpv4Name: fmt.Sprintf("IPv4 on %s", ifName),
		wantIpv6Name: fmt.Sprintf("IPv6 on %s", ifName),
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := clientWithTopo(ifName, test.hasMacSec)
			gotErr := c.addIPProtocols(test.ifcfg)
			if (gotErr == nil && test.wantErr != "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("addIPProtocols got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}

			ifc := c.intfs[ifName]
			if diff := jsonCfgDiff(t, test.wantCfg, ifc.deviceGroup.Ethernet[0]); diff != "" {
				t.Fatalf("addIPProtocols unexpected config diff (-want/+got): %s", diff)
			}

			var gotIpv4Name, gotIpv6Name string
			hasIpv4 := ifc.ipv4 != nil
			if hasIpv4 {
				gotIpv4Name = *(ifc.ipv4.Name)
			}
			hasIpv6 := ifc.ipv6 != nil
			if hasIpv6 {
				gotIpv6Name = *(ifc.ipv6.Name)
			}

			wantIpv4 := test.wantIpv4Name != ""
			if hasIpv4 != wantIpv4 {
				t.Errorf("addIPProtocols: unexpected IPv4 config, was configured: %t, wanted configured? %t", hasIpv4, wantIpv4)
			}
			wantIpv6 := test.wantIpv6Name != ""
			if hasIpv6 != wantIpv6 {
				t.Errorf("addIPProtocols: unexpected IPv6 config, was configured: %t, wanted configured? %t", hasIpv6, wantIpv6)
			}

			if test.wantIpv4Name != gotIpv4Name {
				t.Errorf("addIPProtocols: did not find expected IPv4 config: got name %q, wanted name %q", gotIpv4Name, test.wantIpv4Name)
			}
			if test.wantIpv6Name != gotIpv6Name {
				t.Errorf("addIPProtocols: did not find expected IPv6 config: got name %q, wanted name %q", gotIpv6Name, test.wantIpv6Name)
			}
		})
	}
}

func TestAddIPLoopbackProtocols(t *testing.T) {
	const ifName = "someIntf"

	clientWithTopo := func() *ixATE {
		cfg := &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{}},
			}},
		}
		return &ixATE{
			cfg: cfg,
			intfs: map[string]*intf{
				ifName: &intf{deviceGroup: cfg.Topology[0].DeviceGroup[0]},
			},
		}
	}

	tests := []struct {
		desc    string
		ifcfg   *opb.InterfaceConfig
		wantErr string
		wantCfg *ixconfig.TopologyDeviceGroup
	}{{
		desc: "Invalid IPv4 address.",
		ifcfg: &opb.InterfaceConfig{
			Name:             ifName,
			Ipv4LoopbackCidr: "aa:bb:cc::",
		},
		wantErr: "invalid IPv4 CIDR",
	}, {
		desc: "Invalid IPv6 address.",
		ifcfg: &opb.InterfaceConfig{
			Name:             ifName,
			Ipv6LoopbackCidr: "1.2.3.4",
		},
		wantErr: "invalid IPv6 CIDR",
	}, {
		desc: "Valid IPv4 address.",
		ifcfg: &opb.InterfaceConfig{
			Name:             ifName,
			Ipv4LoopbackCidr: "1.2.3.4/30",
		},
		wantCfg: &ixconfig.TopologyDeviceGroup{
			Ipv4Loopback: []*ixconfig.TopologyIpv4Loopback{{
				Name:    ixconfig.String(fmt.Sprintf("IPv4 loopback on %s", ifName)),
				Address: ixconfig.MultivalueStr("1.2.3.4"),
				Prefix:  ixconfig.MultivalueUint32(30),
			}},
		},
	}, {
		desc: "Valid IPv6 address.",
		ifcfg: &opb.InterfaceConfig{
			Name:             ifName,
			Ipv6LoopbackCidr: "cafe:beef::1/128",
		},
		wantCfg: &ixconfig.TopologyDeviceGroup{
			Ipv6Loopback: []*ixconfig.TopologyIpv6Loopback{{
				Name:    ixconfig.String(fmt.Sprintf("IPv6 loopback on %s", ifName)),
				Address: ixconfig.MultivalueStr("cafe:beef::1"),
				Prefix:  ixconfig.MultivalueUint32(128),
			}},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := clientWithTopo()
			gotErr := c.addIPLoopbackProtocols(test.ifcfg)
			if (gotErr == nil && test.wantErr != "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("addIPLoopbackProtocols got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}

			ifc := c.intfs[ifName]
			if diff := jsonCfgDiff(t, test.wantCfg, ifc.deviceGroup); diff != "" {
				t.Fatalf("addIPLoopbackProtocols unexpected config diff (-want/+got): %s", diff)
			}
		})
	}
}

func TestAddISISProtocols(t *testing.T) {
	const ifName = "someIntf"
	tests := []struct {
		desc                                              string
		ifc                                               *opb.InterfaceConfig
		wantIsisIntfName, wantIsisRtrName, wantRtrCapID   string
		wantEnable3WayHandshake, wantSR, wantAdjacencySid bool
		wantNetwGrpCount                                  int
	}{{
		desc: "No IS-IS config",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
		},
	}, {
		desc: "Simple IS-IS config",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Isis: &opb.ISISConfig{
				Level:       opb.ISISConfig_L1,
				NetworkType: opb.ISISConfig_BROADCAST,
			},
		},
		wantIsisIntfName: fmt.Sprintf("IS-IS on %s", ifName),
		wantIsisRtrName:  fmt.Sprintf("IS-IS Router on %s", ifName),
	}, {
		desc: "Simple IS-IS config, three-way handshake enabled",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Isis: &opb.ISISConfig{
				Level:       opb.ISISConfig_L1,
				NetworkType: opb.ISISConfig_POINT_TO_POINT,
			},
		},
		wantIsisIntfName:        fmt.Sprintf("IS-IS on %s", ifName),
		wantIsisRtrName:         fmt.Sprintf("IS-IS Router on %s", ifName),
		wantEnable3WayHandshake: true,
	}, {
		desc: "IS-IS config with segment routing",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Isis: &opb.ISISConfig{
				Level:       opb.ISISConfig_L1,
				NetworkType: opb.ISISConfig_BROADCAST,
				SegmentRouting: &opb.ISISSegmentRouting{
					Enable: true,
				},
			},
		},
		wantIsisIntfName: fmt.Sprintf("IS-IS on %s", ifName),
		wantIsisRtrName:  fmt.Sprintf("IS-IS Router on %s", ifName),
		wantSR:           true,
	}, {
		desc: "IS-IS config with segment routing and adjacency",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Isis: &opb.ISISConfig{
				Level:       opb.ISISConfig_L1,
				NetworkType: opb.ISISConfig_BROADCAST,
				SegmentRouting: &opb.ISISSegmentRouting{
					Enable:       true,
					AdjacencySid: &opb.ISISSegmentRouting_AdjacencySID{Sid: "11111"},
				},
			},
		},
		wantIsisIntfName: fmt.Sprintf("IS-IS on %s", ifName),
		wantIsisRtrName:  fmt.Sprintf("IS-IS Router on %s", ifName),
		wantSR:           true,
		wantAdjacencySid: true,
	}, {
		desc: "IS-IS config with reachability",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Isis: &opb.ISISConfig{
				Level:       opb.ISISConfig_L1,
				NetworkType: opb.ISISConfig_BROADCAST,
				IsReachabilities: []*opb.ISReachability{{
					Nodes: []*opb.ISReachability_Node{{
						SegmentRouting: &opb.ISISSegmentRouting{
							Enable:       true,
							AdjacencySid: &opb.ISISSegmentRouting_AdjacencySID{Sid: "11111"},
							PrefixSid:    "1.2.3.4/32",
						},
					}},
				}, {
					Nodes: []*opb.ISReachability_Node{{}},
				}},
			},
		},
		wantIsisIntfName: fmt.Sprintf("IS-IS on %s", ifName),
		wantIsisRtrName:  fmt.Sprintf("IS-IS Router on %s", ifName),
		wantNetwGrpCount: 2,
	}, {
		desc: "IS-IS config with router capability ID",
		ifc: &opb.InterfaceConfig{
			Name: ifName,
			Isis: &opb.ISISConfig{
				Level:              opb.ISISConfig_L1,
				NetworkType:        opb.ISISConfig_BROADCAST,
				CapabilityRouterId: "1.2.3.4",
			},
		},
		wantIsisIntfName: fmt.Sprintf("IS-IS on %s", ifName),
		wantIsisRtrName:  fmt.Sprintf("IS-IS Router on %s", ifName),
		wantRtrCapID:     "1.2.3.4",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := clientWithTopoCfg(ifName)
			gotErr := c.addISISProtocols(test.ifc)
			if gotErr != nil {
				t.Fatalf("addISISProtocols: unexpected error: %v", gotErr)
			}

			var gotIsisIntfName, gotIsisRtrName, gotRtrCapID string
			var gotEnable3WayHandshake, gotSR, gotAdjacencySid bool
			dg := c.cfg.Topology[0].DeviceGroup[0]
			eth := dg.Ethernet[0]
			hasIsis := len(eth.IsisL3) > 0
			if hasIsis {
				isisL3 := eth.IsisL3[0]
				gotIsisIntfName = *(isisL3.Name)
				gotEnable3WayHandshake = *(isisL3.Enable3WayHandshake.SingleValue.Value) == "true"
				gotAdjacencySid = *(isisL3.EnableAdjSID.SingleValue.Value) == "true"
			}
			hasIsisRtr := len(dg.IsisL3Router) > 0
			if hasIsisRtr {
				rtr := dg.IsisL3Router[0]
				gotIsisRtrName = *(rtr.Name)
				gotSR = *(rtr.EnableSR)
				if rtr.RtrcapId != nil {
					gotRtrCapID = *(rtr.RtrcapId.SingleValue.Value)
				}
			}
			gotNetwGrpCount := len(dg.NetworkGroup)

			wantIsis := test.wantIsisIntfName != ""
			if hasIsis != wantIsis {
				t.Errorf("addISISProtocols: unexpected IS-IS config, was configured: %t, wanted configured? %t", hasIsis, wantIsis)
			}
			wantIsisRtr := test.wantIsisRtrName != ""
			if hasIsisRtr != wantIsisRtr {
				t.Errorf("addISISProtocols: unexpected IS-IS router config, was configured: %t, wanted configured? %t", hasIsis, wantIsis)
			}

			if test.wantIsisIntfName != gotIsisIntfName {
				t.Errorf("addISISProtocols: did not find expected IS-IS config: got name %q, wanted name %q", gotIsisIntfName, test.wantIsisIntfName)
			}
			if test.wantIsisRtrName != gotIsisRtrName {
				t.Errorf("addISISProtocols: did not find expected IS-IS router config: got name %q, wanted name %q", gotIsisRtrName, test.wantIsisRtrName)
			}
			if test.wantEnable3WayHandshake != gotEnable3WayHandshake {
				t.Errorf("addISISProtocols: incorrect three-way handshake config: enabled? %t, wanted enabled? %t", gotEnable3WayHandshake, test.wantEnable3WayHandshake)
			}
			if test.wantSR != gotSR {
				t.Errorf("addISISProtocols: incorrect SR config: enabled? %t, wanted enabled? %t", gotSR, test.wantSR)
			}
			if test.wantAdjacencySid != gotAdjacencySid {
				t.Errorf("addISISProtocols: incorrect adjacency config: enabled SID? %t, wanted enabled SID? %t", gotAdjacencySid, test.wantAdjacencySid)
			}
			if test.wantRtrCapID != gotRtrCapID {
				t.Errorf("addISISProtocols: incorrect router capability ID: got %q, want %q", gotRtrCapID, test.wantRtrCapID)
			}
			if test.wantNetwGrpCount != gotNetwGrpCount {
				t.Errorf("addISISProtocols: incorrect network group config: got %d network group, wanted %d", gotNetwGrpCount, test.wantNetwGrpCount)
			}
		})
	}
}

func TestISISReachability(t *testing.T) {
	tests := []struct {
		desc         string
		isr          *opb.ISReachability
		wantName     string
		wantNetwTopo *ixconfig.TopologyNetworkTopology
		wantErr      string
	}{{
		desc: "invalid system ID",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "xyz",
			}},
		},
		wantErr: "could not decode",
	}, {
		desc: "invalid TE Router ID",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				TeRouterId: "invalid",
			}},
		},
		wantErr: "invalid TE router ID",
	}, {
		desc: "invalid capability Router ID",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				CapabilityRouterId: "invalid",
			}},
		},
		wantErr: "invalid capability router ID",
	}, {
		desc: "missing IPv4 from address",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links:    []*opb.ISReachability_Node_Link{{ToIpv4: "2.2.2.2/31"}},
			}},
		},
		wantErr: "either both or neither IS-IS node link to/from addresses must be set",
	}, {
		desc: "missing IPv4 to address",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links:    []*opb.ISReachability_Node_Link{{FromIpv4: "1.1.1.1/31"}},
			}},
		},
		wantErr: "either both or neither IS-IS node link to/from addresses must be set",
	}, {
		desc: "invalid IPv4 from address",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "invalid",
					ToIpv4:   "2.2.2.2/31",
				}},
			}},
		},
		wantErr: "could not parse",
	}, {
		desc: "invalid IPv4 to address",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "1.1.1.1/31",
					ToIpv4:   "invalid",
				}},
			}},
		},
		wantErr: "could not parse",
	}, {
		desc: "IPv4 mask mismatch",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "1.1.1.1/31",
					ToIpv4:   "2.2.2.2/32",
				}},
			}},
		},
		wantErr: "unequal masks",
	}, {
		desc: "missing IPv6 from address",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links:    []*opb.ISReachability_Node_Link{{ToIpv6: "bb::/127"}},
			}},
		},
		wantErr: "either both or neither IS-IS node link to/from addresses must be set",
	}, {
		desc: "missing IPv6 to address",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links:    []*opb.ISReachability_Node_Link{{FromIpv6: "aa::/127"}},
			}},
		},
		wantErr: "either both or neither IS-IS node link to/from addresses must be set",
	}, {
		desc: "invalid IPv6 from address",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv6: "invalid",
					ToIpv6:   "bb::/127",
				}},
			}},
		},
		wantErr: "could not parse",
	}, {
		desc: "invalid IPv6 to address",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv6: "aa::/127",
					ToIpv6:   "invalid",
				}},
			}},
		},
		wantErr: "could not parse",
	}, {
		desc: "IPv6 mask mismatch",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv6: "aa::/127",
					ToIpv6:   "bb::/128",
				}},
			}},
		},
		wantErr: "unequal masks",
	}, {
		desc: "invalid ingress metric",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "1.1.1.1/31",
					ToIpv4:   "2.2.2.2/31",
				}},
				IngressMetric: 0,
				EgressMetric:  10,
			}},
		},
		wantErr: "invalid value",
	}, {
		desc: "invalid egress metric",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "1.1.1.1/31",
					ToIpv4:   "2.2.2.2/31",
				}},
				IngressMetric: 10,
				EgressMetric:  70,
			}},
		},
		wantErr: "invalid value",
	}, {
		desc: "IPv4 link w/ exported routes",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "1.1.1.1/31",
					ToIpv4:   "2.2.2.2/31",
				}},
				IngressMetric: 10,
				EgressMetric:  10,
				RoutesIpv4: &opb.ISReachability_Node_Routes{
					Prefix:    "3.3.3.0/25",
					NumRoutes: 4,
					Reachability: &opb.IPReachability{
						Active:              true,
						RouteOrigin:         opb.IPReachability_INTERNAL,
						EnableSidIndexLabel: true,
						SidIndexLabel:       65555,
					},
				},
			}},
		},
		wantNetwTopo: &ixconfig.TopologyNetworkTopology{
			NetTopologyLinear: &ixconfig.TopologyNetTopologyLinear{
				Nodes:          ixconfig.NumberUint32(1),
				LinkMultiplier: ixconfig.NumberUint32(1),
			},
			SimInterface: []*ixconfig.TopologySimInterface{
				{
					IsisL3PseudoInterface: []*ixconfig.TopologyIsisL3PseudoInterface{{
						NoOfTeProfile: ixconfig.NumberUint32(1),
						EnableAdjSID:  ixconfig.MultivalueBoolList(false),
						AdjSID:        ixconfig.MultivalueUintList(1),
						OverrideFFlag: ixconfig.MultivalueBoolList(false),
						FFlag:         ixconfig.MultivalueBoolList(false),
						BFlag:         ixconfig.MultivalueBoolList(false),
						VFlag:         ixconfig.MultivalueBoolList(false),
						LFlag:         ixconfig.MultivalueBoolList(false),
						SFlag:         ixconfig.MultivalueBoolList(false),
						PFlag:         ixconfig.MultivalueBoolList(false),
						IsisL3PseudoIfaceAttPoint1Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint1Config{{
							LinkMetric: ixconfig.MultivalueUintList(10),
						}},
						IsisL3PseudoIfaceAttPoint2Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint2Config{{
							LinkMetric: ixconfig.MultivalueUintList(10),
						}},
					}},
					SimInterfaceIPv4Config: []*ixconfig.TopologySimInterfaceIPv4Config{{
						EnableIp:           ixconfig.MultivalueBoolList(true),
						FromIP:             ixconfig.MultivalueStrList("1.1.1.1"),
						ToIP:               ixconfig.MultivalueStrList("2.2.2.2"),
						SubnetPrefixLength: ixconfig.MultivalueUintList(31),
					}},
					SimInterfaceIPv6Config: []*ixconfig.TopologySimInterfaceIPv6Config{{
						EnableIp:           ixconfig.MultivalueBoolList(false),
						FromIP:             ixconfig.MultivalueStrList("::"),
						ToIP:               ixconfig.MultivalueStrList("::"),
						SubnetPrefixLength: ixconfig.MultivalueUintList(1),
					}},
				},
			},
			SimRouter: []*ixconfig.TopologySimRouter{{
				SystemId: ixconfig.MultivalueStrList("00 00 00 00 00 6b"),
				IsisL3PseudoRouter: []*ixconfig.TopologyIsisL3PseudoRouter{{
					TERouterId:                  ixconfig.MultivalueStrList("0.0.0.0"),
					RtrcapId:                    ixconfig.MultivalueStrList("0.0.0.0"),
					SRAlgorithmCount:            ixconfig.NumberUint32(1),
					SRGBRangeCount:              ixconfig.NumberUint32(1),
					SrlbDescriptorCount:         ixconfig.NumberUint32(1),
					Enable:                      ixconfig.MultivalueBoolList(false),
					EnableIpV6TE:                ixconfig.MultivalueBoolList(false),
					EnableWideMetric:            ixconfig.MultivalueBoolList(false),
					EnableWMforTEinNetworkGroup: ixconfig.MultivalueBoolList(false),
					SIDIndexLabel:               ixconfig.MultivalueUintList(0),
					RFlag:                       ixconfig.MultivalueBoolList(false),
					NFlag:                       ixconfig.MultivalueBoolList(false),
					PFlag:                       ixconfig.MultivalueBoolList(false),
					EFlag:                       ixconfig.MultivalueBoolList(false),
					VFlag:                       ixconfig.MultivalueBoolList(false),
					LFlag:                       ixconfig.MultivalueBoolList(false),
					NodePrefix:                  ixconfig.MultivalueStrList("0.0.0.0"),
					Mask:                        ixconfig.MultivalueUintList(1),
					IPv4PseudoNodeRoutes: []*ixconfig.TopologyIPv4PseudoNodeRoutes{{
						Active:                 ixconfig.MultivalueBoolList(true),
						NetworkAddress:         ixconfig.MultivalueStrList("3.3.3.0"),
						PrefixLength:           ixconfig.MultivalueUintList(25),
						RangeSize:              ixconfig.MultivalueUintList(4),
						Ipv4Metric:             ixconfig.MultivalueUintList(0),
						Algorithm:              ixconfig.MultivalueUintList(0),
						Ipv4RouteOrigin:        ixconfig.MultivalueStrList("internal"),
						ConfigureSIDIndexLabel: ixconfig.MultivalueBoolList(true),
						SIDIndexLabel:          ixconfig.MultivalueUintList(65555),
						Ipv4RFlag:              ixconfig.MultivalueBoolList(false),
						Ipv4NFlag:              ixconfig.MultivalueBoolList(false),
						Ipv4PFlag:              ixconfig.MultivalueBoolList(false),
						Ipv4EFlag:              ixconfig.MultivalueBoolList(false),
						Ipv4VFlag:              ixconfig.MultivalueBoolList(false),
						Ipv4LFlag:              ixconfig.MultivalueBoolList(false),
					}},
					IPv6PseudoNodeRoutes: []*ixconfig.TopologyIPv6PseudoNodeRoutes{&ixconfig.TopologyIPv6PseudoNodeRoutes{}},
				}},
			}},
		},
	}, {
		desc: "IPv6 link w/ exported routes",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv6: "cafe::/127",
					ToIpv6:   "cafe::1/127",
				}},
				IngressMetric: 10,
				EgressMetric:  10,
				RoutesIpv6: &opb.ISReachability_Node_Routes{
					Prefix:    "cafe:beef::/120",
					NumRoutes: 4,
					Reachability: &opb.IPReachability{
						Active:      true,
						RouteOrigin: opb.IPReachability_INTERNAL,
					},
				},
			}},
		},
		wantNetwTopo: &ixconfig.TopologyNetworkTopology{
			NetTopologyLinear: &ixconfig.TopologyNetTopologyLinear{
				Nodes:          ixconfig.NumberUint32(1),
				LinkMultiplier: ixconfig.NumberUint32(1),
			},
			SimInterface: []*ixconfig.TopologySimInterface{
				{
					IsisL3PseudoInterface: []*ixconfig.TopologyIsisL3PseudoInterface{{
						NoOfTeProfile: ixconfig.NumberUint32(1),
						EnableAdjSID:  ixconfig.MultivalueBoolList(false),
						AdjSID:        ixconfig.MultivalueUintList(1),
						OverrideFFlag: ixconfig.MultivalueBoolList(false),
						FFlag:         ixconfig.MultivalueBoolList(false),
						BFlag:         ixconfig.MultivalueBoolList(false),
						VFlag:         ixconfig.MultivalueBoolList(false),
						LFlag:         ixconfig.MultivalueBoolList(false),
						SFlag:         ixconfig.MultivalueBoolList(false),
						PFlag:         ixconfig.MultivalueBoolList(false),
						IsisL3PseudoIfaceAttPoint1Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint1Config{{
							LinkMetric: ixconfig.MultivalueUintList(10),
						}},
						IsisL3PseudoIfaceAttPoint2Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint2Config{{
							LinkMetric: ixconfig.MultivalueUintList(10),
						}},
					}},
					SimInterfaceIPv4Config: []*ixconfig.TopologySimInterfaceIPv4Config{{
						EnableIp:           ixconfig.MultivalueBoolList(false),
						FromIP:             ixconfig.MultivalueStrList("0.0.0.0"),
						ToIP:               ixconfig.MultivalueStrList("0.0.0.0"),
						SubnetPrefixLength: ixconfig.MultivalueUintList(1),
					}},
					SimInterfaceIPv6Config: []*ixconfig.TopologySimInterfaceIPv6Config{{
						EnableIp:           ixconfig.MultivalueBoolList(true),
						FromIP:             ixconfig.MultivalueStrList("cafe::"),
						ToIP:               ixconfig.MultivalueStrList("cafe::1"),
						SubnetPrefixLength: ixconfig.MultivalueUintList(127),
					}},
				},
			},
			SimRouter: []*ixconfig.TopologySimRouter{{
				SystemId: ixconfig.MultivalueStrList("00 00 00 00 00 6b"),
				IsisL3PseudoRouter: []*ixconfig.TopologyIsisL3PseudoRouter{{
					TERouterId:                  ixconfig.MultivalueStrList("0.0.0.0"),
					RtrcapId:                    ixconfig.MultivalueStrList("0.0.0.0"),
					SRAlgorithmCount:            ixconfig.NumberUint32(1),
					SRGBRangeCount:              ixconfig.NumberUint32(1),
					SrlbDescriptorCount:         ixconfig.NumberUint32(1),
					Enable:                      ixconfig.MultivalueBoolList(false),
					EnableIpV6TE:                ixconfig.MultivalueBoolList(false),
					EnableWideMetric:            ixconfig.MultivalueBoolList(false),
					EnableWMforTEinNetworkGroup: ixconfig.MultivalueBoolList(false),
					SIDIndexLabel:               ixconfig.MultivalueUintList(0),
					RFlag:                       ixconfig.MultivalueBoolList(false),
					NFlag:                       ixconfig.MultivalueBoolList(false),
					PFlag:                       ixconfig.MultivalueBoolList(false),
					EFlag:                       ixconfig.MultivalueBoolList(false),
					VFlag:                       ixconfig.MultivalueBoolList(false),
					LFlag:                       ixconfig.MultivalueBoolList(false),
					NodePrefix:                  ixconfig.MultivalueStrList("0.0.0.0"),
					Mask:                        ixconfig.MultivalueUintList(1),
					IPv4PseudoNodeRoutes:        []*ixconfig.TopologyIPv4PseudoNodeRoutes{&ixconfig.TopologyIPv4PseudoNodeRoutes{}},
					IPv6PseudoNodeRoutes: []*ixconfig.TopologyIPv6PseudoNodeRoutes{{
						Active:                 ixconfig.MultivalueBoolList(true),
						NetworkAddress:         ixconfig.MultivalueStrList("cafe:beef::"),
						Prefix:                 ixconfig.MultivalueUintList(120),
						RangeSize:              ixconfig.MultivalueUintList(4),
						Ipv6Metric:             ixconfig.MultivalueUintList(0),
						Algorithm:              ixconfig.MultivalueUintList(0),
						Ipv6RouteOrigin:        ixconfig.MultivalueStrList("internal"),
						ConfigureSIDIndexLabel: ixconfig.MultivalueBoolList(false),
						SIDIndexLabel:          ixconfig.MultivalueUintList(0),
						Ipv6RFlag:              ixconfig.MultivalueBoolList(false),
						Ipv6NFlag:              ixconfig.MultivalueBoolList(false),
						Ipv6PFlag:              ixconfig.MultivalueBoolList(false),
						Ipv6EFlag:              ixconfig.MultivalueBoolList(false),
						Ipv6VFlag:              ixconfig.MultivalueBoolList(false),
						Ipv6LFlag:              ixconfig.MultivalueBoolList(false),
					}},
				}},
			}},
		},
	}, {
		desc: "Routes w/o reachability",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv6: "cafe::/127",
					ToIpv6:   "cafe::1/127",
					FromIpv4: "1.1.1.0/31",
					ToIpv4:   "1.1.1.1/31",
				}},
				IngressMetric: 10,
				EgressMetric:  10,
				RoutesIpv6: &opb.ISReachability_Node_Routes{
					Prefix:    "cafe:beef::/120",
					NumRoutes: 4,
				},
				RoutesIpv4: &opb.ISReachability_Node_Routes{
					Prefix:    "10.0.0.0/8",
					NumRoutes: 4,
				},
			}},
		},
		wantNetwTopo: &ixconfig.TopologyNetworkTopology{
			NetTopologyLinear: &ixconfig.TopologyNetTopologyLinear{
				Nodes:          ixconfig.NumberUint32(1),
				LinkMultiplier: ixconfig.NumberUint32(1),
			},
			SimInterface: []*ixconfig.TopologySimInterface{
				{
					IsisL3PseudoInterface: []*ixconfig.TopologyIsisL3PseudoInterface{{
						NoOfTeProfile: ixconfig.NumberUint32(1),
						EnableAdjSID:  ixconfig.MultivalueBoolList(false),
						AdjSID:        ixconfig.MultivalueUintList(1),
						OverrideFFlag: ixconfig.MultivalueBoolList(false),
						FFlag:         ixconfig.MultivalueBoolList(false),
						BFlag:         ixconfig.MultivalueBoolList(false),
						VFlag:         ixconfig.MultivalueBoolList(false),
						LFlag:         ixconfig.MultivalueBoolList(false),
						SFlag:         ixconfig.MultivalueBoolList(false),
						PFlag:         ixconfig.MultivalueBoolList(false),
						IsisL3PseudoIfaceAttPoint1Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint1Config{{
							LinkMetric: ixconfig.MultivalueUintList(10),
						}},
						IsisL3PseudoIfaceAttPoint2Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint2Config{{
							LinkMetric: ixconfig.MultivalueUintList(10),
						}},
					}},
					SimInterfaceIPv4Config: []*ixconfig.TopologySimInterfaceIPv4Config{{
						EnableIp:           ixconfig.MultivalueBoolList(true),
						FromIP:             ixconfig.MultivalueStrList("1.1.1.0"),
						ToIP:               ixconfig.MultivalueStrList("1.1.1.1"),
						SubnetPrefixLength: ixconfig.MultivalueUintList(31),
					}},
					SimInterfaceIPv6Config: []*ixconfig.TopologySimInterfaceIPv6Config{{
						EnableIp:           ixconfig.MultivalueBoolList(true),
						FromIP:             ixconfig.MultivalueStrList("cafe::"),
						ToIP:               ixconfig.MultivalueStrList("cafe::1"),
						SubnetPrefixLength: ixconfig.MultivalueUintList(127),
					}},
				},
			},
			SimRouter: []*ixconfig.TopologySimRouter{{
				SystemId: ixconfig.MultivalueStrList("00 00 00 00 00 6b"),
				IsisL3PseudoRouter: []*ixconfig.TopologyIsisL3PseudoRouter{{
					TERouterId:                  ixconfig.MultivalueStrList("0.0.0.0"),
					RtrcapId:                    ixconfig.MultivalueStrList("0.0.0.0"),
					SRAlgorithmCount:            ixconfig.NumberUint32(1),
					SRGBRangeCount:              ixconfig.NumberUint32(1),
					SrlbDescriptorCount:         ixconfig.NumberUint32(1),
					Enable:                      ixconfig.MultivalueBoolList(false),
					EnableIpV6TE:                ixconfig.MultivalueBoolList(false),
					EnableWideMetric:            ixconfig.MultivalueBoolList(false),
					EnableWMforTEinNetworkGroup: ixconfig.MultivalueBoolList(false),
					SIDIndexLabel:               ixconfig.MultivalueUintList(0),
					RFlag:                       ixconfig.MultivalueBoolList(false),
					NFlag:                       ixconfig.MultivalueBoolList(false),
					PFlag:                       ixconfig.MultivalueBoolList(false),
					EFlag:                       ixconfig.MultivalueBoolList(false),
					VFlag:                       ixconfig.MultivalueBoolList(false),
					LFlag:                       ixconfig.MultivalueBoolList(false),
					NodePrefix:                  ixconfig.MultivalueStrList("0.0.0.0"),
					Mask:                        ixconfig.MultivalueUintList(1),
					IPv4PseudoNodeRoutes: []*ixconfig.TopologyIPv4PseudoNodeRoutes{{
						Active:                 ixconfig.MultivalueBoolList(true),
						NetworkAddress:         ixconfig.MultivalueStrList("10.0.0.0"),
						PrefixLength:           ixconfig.MultivalueUintList(8),
						RangeSize:              ixconfig.MultivalueUintList(4),
						Ipv4Metric:             ixconfig.MultivalueUintList(0),
						Algorithm:              ixconfig.MultivalueUintList(0),
						Ipv4RouteOrigin:        ixconfig.MultivalueStrList("internal"),
						ConfigureSIDIndexLabel: ixconfig.MultivalueBoolList(false),
						SIDIndexLabel:          ixconfig.MultivalueUintList(0),
						Ipv4RFlag:              ixconfig.MultivalueBoolList(false),
						Ipv4NFlag:              ixconfig.MultivalueBoolList(false),
						Ipv4PFlag:              ixconfig.MultivalueBoolList(false),
						Ipv4EFlag:              ixconfig.MultivalueBoolList(false),
						Ipv4VFlag:              ixconfig.MultivalueBoolList(false),
						Ipv4LFlag:              ixconfig.MultivalueBoolList(false),
					}},
					IPv6PseudoNodeRoutes: []*ixconfig.TopologyIPv6PseudoNodeRoutes{{
						Active:                 ixconfig.MultivalueBoolList(true),
						NetworkAddress:         ixconfig.MultivalueStrList("cafe:beef::"),
						Prefix:                 ixconfig.MultivalueUintList(120),
						RangeSize:              ixconfig.MultivalueUintList(4),
						Ipv6Metric:             ixconfig.MultivalueUintList(0),
						Algorithm:              ixconfig.MultivalueUintList(0),
						Ipv6RouteOrigin:        ixconfig.MultivalueStrList("internal"),
						ConfigureSIDIndexLabel: ixconfig.MultivalueBoolList(false),
						SIDIndexLabel:          ixconfig.MultivalueUintList(0),
						Ipv6RFlag:              ixconfig.MultivalueBoolList(false),
						Ipv6NFlag:              ixconfig.MultivalueBoolList(false),
						Ipv6PFlag:              ixconfig.MultivalueBoolList(false),
						Ipv6EFlag:              ixconfig.MultivalueBoolList(false),
						Ipv6VFlag:              ixconfig.MultivalueBoolList(false),
						Ipv6LFlag:              ixconfig.MultivalueBoolList(false),
					}},
				}},
			}},
		},
	}, {
		desc: "IPv6 link w/ wide metric and custom name",
		isr: &opb.ISReachability{
			Name: "isrName",
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv6: "aa::/127",
					ToIpv6:   "bb::/127",
				}},
				IngressMetric:    5,
				EgressMetric:     5,
				EnableWideMetric: true,
			}},
		},
		wantName: "isrName",
		wantNetwTopo: &ixconfig.TopologyNetworkTopology{
			NetTopologyLinear: &ixconfig.TopologyNetTopologyLinear{
				Nodes:          ixconfig.NumberUint32(1),
				LinkMultiplier: ixconfig.NumberUint32(1),
			},
			SimInterface: []*ixconfig.TopologySimInterface{
				{
					IsisL3PseudoInterface: []*ixconfig.TopologyIsisL3PseudoInterface{{
						NoOfTeProfile: ixconfig.NumberUint32(1),
						EnableAdjSID:  ixconfig.MultivalueBoolList(false),
						AdjSID:        ixconfig.MultivalueUintList(1),
						OverrideFFlag: ixconfig.MultivalueBoolList(false),
						FFlag:         ixconfig.MultivalueBoolList(false),
						BFlag:         ixconfig.MultivalueBoolList(false),
						VFlag:         ixconfig.MultivalueBoolList(false),
						LFlag:         ixconfig.MultivalueBoolList(false),
						SFlag:         ixconfig.MultivalueBoolList(false),
						PFlag:         ixconfig.MultivalueBoolList(false),
						IsisL3PseudoIfaceAttPoint1Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint1Config{{
							LinkMetric: ixconfig.MultivalueUintList(5),
						}},
						IsisL3PseudoIfaceAttPoint2Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint2Config{{
							LinkMetric: ixconfig.MultivalueUintList(5),
						}},
					}},
					SimInterfaceIPv4Config: []*ixconfig.TopologySimInterfaceIPv4Config{{
						EnableIp:           ixconfig.MultivalueBoolList(false),
						FromIP:             ixconfig.MultivalueStrList("0.0.0.0"),
						ToIP:               ixconfig.MultivalueStrList("0.0.0.0"),
						SubnetPrefixLength: ixconfig.MultivalueUintList(1),
					}},
					SimInterfaceIPv6Config: []*ixconfig.TopologySimInterfaceIPv6Config{{
						EnableIp:           ixconfig.MultivalueBoolList(true),
						FromIP:             ixconfig.MultivalueStrList("aa::"),
						ToIP:               ixconfig.MultivalueStrList("bb::"),
						SubnetPrefixLength: ixconfig.MultivalueUintList(127),
					}},
				},
			},
			SimRouter: []*ixconfig.TopologySimRouter{{
				SystemId: ixconfig.MultivalueStrList("00 00 00 00 00 6b"),
				IsisL3PseudoRouter: []*ixconfig.TopologyIsisL3PseudoRouter{{
					TERouterId:                  ixconfig.MultivalueStrList("0.0.0.0"),
					RtrcapId:                    ixconfig.MultivalueStrList("0.0.0.0"),
					SRAlgorithmCount:            ixconfig.NumberUint32(1),
					SRGBRangeCount:              ixconfig.NumberUint32(1),
					SrlbDescriptorCount:         ixconfig.NumberUint32(1),
					Enable:                      ixconfig.MultivalueBoolList(false),
					EnableIpV6TE:                ixconfig.MultivalueBoolList(false),
					EnableWideMetric:            ixconfig.MultivalueBoolList(true),
					EnableWMforTEinNetworkGroup: ixconfig.MultivalueBoolList(true),
					SIDIndexLabel:               ixconfig.MultivalueUintList(0),
					RFlag:                       ixconfig.MultivalueBoolList(false),
					NFlag:                       ixconfig.MultivalueBoolList(false),
					PFlag:                       ixconfig.MultivalueBoolList(false),
					EFlag:                       ixconfig.MultivalueBoolList(false),
					VFlag:                       ixconfig.MultivalueBoolList(false),
					LFlag:                       ixconfig.MultivalueBoolList(false),
					NodePrefix:                  ixconfig.MultivalueStrList("0.0.0.0"),
					Mask:                        ixconfig.MultivalueUintList(1),
					IPv4PseudoNodeRoutes:        []*ixconfig.TopologyIPv4PseudoNodeRoutes{&ixconfig.TopologyIPv4PseudoNodeRoutes{}},
					IPv6PseudoNodeRoutes:        []*ixconfig.TopologyIPv6PseudoNodeRoutes{&ixconfig.TopologyIPv6PseudoNodeRoutes{}},
				}},
			}},
		},
	}, {
		desc: "Multiple nodes with varying link counts",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "1.1.1.1/32",
					ToIpv4:   "2.2.2.2/32",
					FromIpv6: "aa::/127",
					ToIpv6:   "bb::/127",
				}},
				IngressMetric: 10,
				EgressMetric:  10,
			}, {
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "3.3.3.3/31",
					ToIpv4:   "4.4.4.4/31",
				}, {
					FromIpv6: "cc::/128",
					ToIpv6:   "dd::/128",
				}},
				IngressMetric: 10,
				EgressMetric:  10,
			}},
		},
		wantNetwTopo: &ixconfig.TopologyNetworkTopology{
			NetTopologyLinear: &ixconfig.TopologyNetTopologyLinear{
				Nodes:          ixconfig.NumberUint32(2),
				LinkMultiplier: ixconfig.NumberUint32(2),
			},
			SimInterface: []*ixconfig.TopologySimInterface{{
				IsisL3PseudoInterface: []*ixconfig.TopologyIsisL3PseudoInterface{{
					NoOfTeProfile: ixconfig.NumberUint32(1),
					EnableAdjSID:  ixconfig.MultivalueBoolList(false, false, false, false),
					AdjSID:        ixconfig.MultivalueUintList(1, 1, 1, 1),
					OverrideFFlag: ixconfig.MultivalueBoolList(false, false, false, false),
					FFlag:         ixconfig.MultivalueBoolList(false, false, false, false),
					BFlag:         ixconfig.MultivalueBoolList(false, false, false, false),
					VFlag:         ixconfig.MultivalueBoolList(false, false, false, false),
					LFlag:         ixconfig.MultivalueBoolList(false, false, false, false),
					SFlag:         ixconfig.MultivalueBoolList(false, false, false, false),
					PFlag:         ixconfig.MultivalueBoolList(false, false, false, false),
					IsisL3PseudoIfaceAttPoint1Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint1Config{{
						LinkMetric: ixconfig.MultivalueUintList(10, 10, 10, 10),
					}},
					IsisL3PseudoIfaceAttPoint2Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint2Config{{
						LinkMetric: ixconfig.MultivalueUintList(10, 10, 10, 10),
					}},
				}},
				SimInterfaceIPv4Config: []*ixconfig.TopologySimInterfaceIPv4Config{{
					EnableIp:           ixconfig.MultivalueBoolList(true, false, true, false),
					FromIP:             ixconfig.MultivalueStrList("1.1.1.1", "0.0.0.0", "3.3.3.3", "0.0.0.0"),
					ToIP:               ixconfig.MultivalueStrList("2.2.2.2", "0.0.0.0", "4.4.4.4", "0.0.0.0"),
					SubnetPrefixLength: ixconfig.MultivalueUintList(32, 1, 31, 1),
				}},
				SimInterfaceIPv6Config: []*ixconfig.TopologySimInterfaceIPv6Config{{
					EnableIp:           ixconfig.MultivalueBoolList(true, false, false, true),
					FromIP:             ixconfig.MultivalueStrList("aa::", "::", "::", "cc::"),
					ToIP:               ixconfig.MultivalueStrList("bb::", "::", "::", "dd::"),
					SubnetPrefixLength: ixconfig.MultivalueUintList(127, 1, 1, 128),
				}},
			}},
			SimRouter: []*ixconfig.TopologySimRouter{{
				SystemId: ixconfig.MultivalueStrList("00 00 00 00 00 6b", "00 00 00 00 00 6b"),
				IsisL3PseudoRouter: []*ixconfig.TopologyIsisL3PseudoRouter{{
					TERouterId:                  ixconfig.MultivalueStrList("0.0.0.0", "0.0.0.0"),
					RtrcapId:                    ixconfig.MultivalueStrList("0.0.0.0", "0.0.0.0"),
					SRAlgorithmCount:            ixconfig.NumberUint32(1),
					SRGBRangeCount:              ixconfig.NumberUint32(1),
					SrlbDescriptorCount:         ixconfig.NumberUint32(1),
					Enable:                      ixconfig.MultivalueBoolList(false, false),
					EnableIpV6TE:                ixconfig.MultivalueBoolList(false, false),
					EnableWideMetric:            ixconfig.MultivalueBoolList(false, false),
					EnableWMforTEinNetworkGroup: ixconfig.MultivalueBoolList(false, false),
					SIDIndexLabel:               ixconfig.MultivalueUintList(0, 0),
					RFlag:                       ixconfig.MultivalueBoolList(false, false),
					NFlag:                       ixconfig.MultivalueBoolList(false, false),
					PFlag:                       ixconfig.MultivalueBoolList(false, false),
					EFlag:                       ixconfig.MultivalueBoolList(false, false),
					VFlag:                       ixconfig.MultivalueBoolList(false, false),
					LFlag:                       ixconfig.MultivalueBoolList(false, false),
					NodePrefix:                  ixconfig.MultivalueStrList("0.0.0.0", "0.0.0.0"),
					Mask:                        ixconfig.MultivalueUintList(1, 1),
					IPv4PseudoNodeRoutes:        []*ixconfig.TopologyIPv4PseudoNodeRoutes{&ixconfig.TopologyIPv4PseudoNodeRoutes{}},
					IPv6PseudoNodeRoutes:        []*ixconfig.TopologyIPv6PseudoNodeRoutes{&ixconfig.TopologyIPv6PseudoNodeRoutes{}},
				}},
			}},
		},
	}, {
		desc: "inconsistent SRGB range counts",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "1.1.1.1/32",
					ToIpv4:   "2.2.2.2/32",
				}},
				SegmentRouting: &opb.ISISSegmentRouting{
					SrgbRanges: []*opb.ISISSegmentRouting_SIDRange{{
						SidStartLabel: 16000,
						SidCount:      8000,
					}, {
						SidStartLabel: 400000,
						SidCount:      65001,
					}},
				},
			}, {
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "3.3.3.3/32",
					ToIpv4:   "4.4.4.4/32",
				}},
				SegmentRouting: &opb.ISISSegmentRouting{
					SrgbRanges: []*opb.ISISSegmentRouting_SIDRange{{
						SidStartLabel: 400000,
						SidCount:      65001,
					}},
				},
			}},
		},
		wantErr: "need the same number of SRGB ranges",
	}, {
		desc: "inconsistent SRLB range counts",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "1.1.1.1/32",
					ToIpv4:   "2.2.2.2/32",
				}},
				SegmentRouting: &opb.ISISSegmentRouting{
					SrlbRanges: []*opb.ISISSegmentRouting_SIDRange{{
						SidStartLabel: 16000,
						SidCount:      8000,
					}, {
						SidStartLabel: 400000,
						SidCount:      65001,
					}},
				},
			}, {
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "3.3.3.3/32",
					ToIpv4:   "4.4.4.4/32",
				}},
				SegmentRouting: &opb.ISISSegmentRouting{
					SrlbRanges: []*opb.ISISSegmentRouting_SIDRange{{
						SidStartLabel: 400000,
						SidCount:      65001,
					}},
				},
			}},
		},
		wantErr: "need the same number of SRLB ranges",
	}, {
		desc: "SR configuration",
		isr: &opb.ISReachability{
			Nodes: []*opb.ISReachability_Node{{
				SystemId: "6b",
				Links: []*opb.ISReachability_Node_Link{{
					FromIpv4: "1.1.1.1/32",
					ToIpv4:   "2.2.2.2/32",
				}},
				IngressMetric:    10,
				EgressMetric:     10,
				EnableTe:         true,
				EnableWideMetric: true,
				SegmentRouting: &opb.ISISSegmentRouting{
					Enable:           true,
					FlagReadvertise:  true,
					FlagNodeSid:      true,
					FlagNoPhp:        true,
					FlagExplicitNull: true,
					FlagValue:        true,
					FlagLocal:        true,
					SidIndexLabel:    2,
					AdjacencySid: &opb.ISISSegmentRouting_AdjacencySID{
						Sid:                       "2",
						OverrideFlagAddressFamily: true,
						FlagAddressFamily:         true,
						FlagBackup:                true,
						FlagValue:                 true,
						FlagLocal:                 true,
						FlagSet:                   true,
						FlagPersistent:            true,
					},
					SrgbRanges: []*opb.ISISSegmentRouting_SIDRange{{
						SidStartLabel: 400000,
						SidCount:      65001,
					}},
					SrlbRanges: []*opb.ISISSegmentRouting_SIDRange{{
						SidStartLabel: 300000,
						SidCount:      65001,
					}},
					Algorithms: []uint32{1, 2},
					PrefixSid:  "10.10.10.0/24",
				},
			}},
		},
		wantNetwTopo: &ixconfig.TopologyNetworkTopology{
			NetTopologyLinear: &ixconfig.TopologyNetTopologyLinear{
				Nodes:          ixconfig.NumberUint32(1),
				LinkMultiplier: ixconfig.NumberUint32(1),
			},
			SimInterface: []*ixconfig.TopologySimInterface{{
				IsisL3PseudoInterface: []*ixconfig.TopologyIsisL3PseudoInterface{{
					NoOfTeProfile: ixconfig.NumberUint32(1),
					EnableAdjSID:  ixconfig.MultivalueBoolList(true),
					AdjSID:        ixconfig.MultivalueUintList(2),
					OverrideFFlag: ixconfig.MultivalueBoolList(true),
					FFlag:         ixconfig.MultivalueBoolList(true),
					BFlag:         ixconfig.MultivalueBoolList(true),
					VFlag:         ixconfig.MultivalueBoolList(true),
					LFlag:         ixconfig.MultivalueBoolList(true),
					SFlag:         ixconfig.MultivalueBoolList(true),
					PFlag:         ixconfig.MultivalueBoolList(true),
					IsisL3PseudoIfaceAttPoint1Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint1Config{{
						LinkMetric: ixconfig.MultivalueUintList(10),
					}},
					IsisL3PseudoIfaceAttPoint2Config: []*ixconfig.TopologyIsisL3PseudoIfaceAttPoint2Config{{
						LinkMetric: ixconfig.MultivalueUintList(10),
					}},
				}},
				SimInterfaceIPv4Config: []*ixconfig.TopologySimInterfaceIPv4Config{{
					EnableIp:           ixconfig.MultivalueBoolList(true),
					FromIP:             ixconfig.MultivalueStrList("1.1.1.1"),
					ToIP:               ixconfig.MultivalueStrList("2.2.2.2"),
					SubnetPrefixLength: ixconfig.MultivalueUintList(32),
				}},
				SimInterfaceIPv6Config: []*ixconfig.TopologySimInterfaceIPv6Config{{
					EnableIp:           ixconfig.MultivalueBoolList(false),
					FromIP:             ixconfig.MultivalueStrList("::"),
					ToIP:               ixconfig.MultivalueStrList("::"),
					SubnetPrefixLength: ixconfig.MultivalueUintList(1),
				}},
			}},
			SimRouter: []*ixconfig.TopologySimRouter{{
				SystemId: ixconfig.MultivalueStrList("00 00 00 00 00 6b"),
				IsisL3PseudoRouter: []*ixconfig.TopologyIsisL3PseudoRouter{{
					TERouterId:       ixconfig.MultivalueStrList("0.0.0.0"),
					RtrcapId:         ixconfig.MultivalueStrList("0.0.0.0"),
					SRAlgorithmCount: ixconfig.NumberUint32(2),
					IsisSRAlgorithmList: []*ixconfig.TopologyIsisSrAlgorithmList{{
						IsisSrAlgorithm: ixconfig.MultivalueUintList(1),
					}, {
						IsisSrAlgorithm: ixconfig.MultivalueUintList(2),
					}},
					SRGBRangeCount: ixconfig.NumberUint32(1),
					IsisSRGBRangeSubObjectsList: []*ixconfig.TopologyIsisSrgbRangeSubObjectsList{{
						SIDCount:      ixconfig.MultivalueUintList(65001),
						StartSIDLabel: ixconfig.MultivalueUintList(400000),
					}},
					SrlbDescriptorCount: ixconfig.NumberUint32(1),
					IsisSRLBDescriptorList: []*ixconfig.TopologyIsisSrlbDescriptorList{{
						SIDCount:      ixconfig.MultivalueUintList(65001),
						StartSIDLabel: ixconfig.MultivalueUintList(300000),
					}},
					Enable:                      ixconfig.MultivalueBoolList(true),
					EnableSR:                    ixconfig.Bool(true),
					EnableIpV6TE:                ixconfig.MultivalueBoolList(true),
					EnableWideMetric:            ixconfig.MultivalueBoolList(true),
					EnableWMforTEinNetworkGroup: ixconfig.MultivalueBoolList(true),
					SIDIndexLabel:               ixconfig.MultivalueUintList(2),
					RFlag:                       ixconfig.MultivalueBoolList(true),
					NFlag:                       ixconfig.MultivalueBoolList(true),
					PFlag:                       ixconfig.MultivalueBoolList(true),
					EFlag:                       ixconfig.MultivalueBoolList(true),
					VFlag:                       ixconfig.MultivalueBoolList(true),
					LFlag:                       ixconfig.MultivalueBoolList(true),
					NodePrefix:                  ixconfig.MultivalueStrList("10.10.10.0"),
					Mask:                        ixconfig.MultivalueUintList(24),
					IPv4PseudoNodeRoutes:        []*ixconfig.TopologyIPv4PseudoNodeRoutes{&ixconfig.TopologyIPv4PseudoNodeRoutes{}},
					IPv6PseudoNodeRoutes:        []*ixconfig.TopologyIPv6PseudoNodeRoutes{&ixconfig.TopologyIPv6PseudoNodeRoutes{}},
				}},
			}},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotNetwGrps, gotErr := isisReachability("someIntf", []*opb.ISReachability{test.isr})
			if ((gotErr == nil) != (test.wantErr == "")) || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("isIsReachabilities: got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}
			wantNetwGrpCount := 0
			if test.wantNetwTopo != nil {
				wantNetwGrpCount = 1
			}
			if wantNetwGrpCount != len(gotNetwGrps) {
				t.Fatalf("isIsReachabilities: unexpected number of network groups, got %d, want %d", len(gotNetwGrps), wantNetwGrpCount)
			}
			wantName := test.wantName
			if wantName == "" {
				wantName = fmt.Sprintf("IS-IS Topology 0 for someIntf")
			}
			if test.wantNetwTopo != nil {
				if diff := jsonCfgDiff(t, test.wantNetwTopo, gotNetwGrps[wantName].NetworkTopology); diff != "" {
					t.Fatalf("isIsReachabilities: network group diff at (-want/+got): %s", diff)
				}
			}
		})
	}
}

func TestAddBGPProtocols(t *testing.T) {
	const ifName = "someIntf"
	baseClient := func(withIPv4, withIPv6, withIPv4Loopback, withIPv6Loopback bool) *ixATE {
		cfg := &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
					Ethernet: []*ixconfig.TopologyEthernet{{}},
				}},
			}},
		}
		ifc := &intf{deviceGroup: cfg.Topology[0].DeviceGroup[0]}
		dg := cfg.Topology[0].DeviceGroup[0]
		eth := dg.Ethernet[0]
		if withIPv4 {
			eth.Ipv4 = []*ixconfig.TopologyIpv4{{}}
			ifc.ipv4 = eth.Ipv4[0]
		}
		if withIPv6 {
			eth.Ipv6 = []*ixconfig.TopologyIpv6{{}}
			ifc.ipv6 = eth.Ipv6[0]
		}
		if withIPv4Loopback {
			dg.Ipv4Loopback = []*ixconfig.TopologyIpv4Loopback{{}}
			ifc.ipv4Loopback = dg.Ipv4Loopback[0]
		}
		if withIPv6Loopback {
			dg.Ipv6Loopback = []*ixconfig.TopologyIpv6Loopback{{}}
			ifc.ipv6Loopback = dg.Ipv6Loopback[0]
		}
		return &ixATE{
			cfg:   cfg,
			intfs: map[string]*intf{ifName: ifc},
		}
	}

	tests := []struct {
		desc                                                           string
		hasV4Intf, hasV6Intf, hasV4Loopback, hasV6Loopback             bool
		peers                                                          []*opb.BgpPeer
		wantV4Peer, wantV6Peer, wantV4LoopbackPeer, wantV6LoopbackPeer bool
		wantErr                                                        string
	}{{
		desc: "invalid peer address",
		peers: []*opb.BgpPeer{{
			PeerAddress: "invalid",
		}},
		wantErr: "invalid peer address",
	}, {
		desc: "too many peers",
		peers: []*opb.BgpPeer{{
			PeerAddress: "1.1.1.1",
		}, {
			PeerAddress: "aa::",
		}, {
			PeerAddress: "2.2.2.2",
		}},
		wantErr: "exceeds maximum",
	}, {
		desc: "v4 peer without intf",
		peers: []*opb.BgpPeer{{
			PeerAddress: "1.1.1.1",
		}},
		wantErr: "without IPv4 configured on interface",
	}, {
		desc: "v6 peer without intf",
		peers: []*opb.BgpPeer{{
			PeerAddress: "aa::",
		}},
		wantErr: "without IPv6 configured on interface",
	}, {
		desc: "v4 loopback peer without intf",
		peers: []*opb.BgpPeer{{
			PeerAddress: "1.1.1.1",
			OnLoopback:  true,
		}},
		wantErr: "without IPv4 loopback configured on interface",
	}, {
		desc: "v6 loopback peer without intf",
		peers: []*opb.BgpPeer{{
			PeerAddress: "cafe::1",
			OnLoopback:  true,
		}},
		wantErr: "without IPv6 loopback configured on interface",
	}, {
		desc:      "valid peers",
		hasV4Intf: true,
		hasV6Intf: true,
		peers: []*opb.BgpPeer{{
			PeerAddress: "1.1.1.1",
		}, {
			PeerAddress: "aa::",
		}},
		wantV4Peer: true,
		wantV6Peer: true,
	}, {
		desc:          "valid loopback peer",
		hasV4Loopback: true,
		hasV6Loopback: true,
		peers: []*opb.BgpPeer{{
			PeerAddress: "1.1.1.1",
			OnLoopback:  true,
		}, {
			PeerAddress: "cafe::1",
			OnLoopback:  true,
		}},
		wantV4LoopbackPeer: true,
		wantV6LoopbackPeer: true,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ifc := &opb.InterfaceConfig{
				Name: ifName,
				Bgp: &opb.BgpConfig{
					BgpPeers: test.peers,
				},
			}
			c := baseClient(test.hasV4Intf, test.hasV6Intf, test.hasV4Loopback, test.hasV6Loopback)
			gotErr := c.addBGPProtocols(ifc)
			if ((gotErr == nil) != (test.wantErr == "")) || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("addBGPProtocols: got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}
			dg := c.cfg.Topology[0].DeviceGroup[0]
			eth := dg.Ethernet[0]
			hasV4Peer := len(eth.Ipv4) > 0 && len(eth.Ipv4[0].BgpIpv4Peer) > 0
			if hasV4Peer != test.wantV4Peer {
				t.Errorf("addBGPProtocols: has V4 peer? %t, wanted V4 peer? %t", hasV4Peer, test.wantV4Peer)
			}
			hasV6Peer := len(eth.Ipv6) > 0 && len(eth.Ipv6[0].BgpIpv6Peer) > 0
			if hasV6Peer != test.wantV6Peer {
				t.Errorf("addBGPProtocols: has V6 peer? %t, wanted V6 peer? %t", hasV6Peer, test.wantV6Peer)
			}
			hasV4LoopbackPeer := len(dg.Ipv4Loopback) > 0 && len(dg.Ipv4Loopback[0].BgpIpv4Peer) > 0
			if hasV4LoopbackPeer != test.wantV4LoopbackPeer {
				t.Errorf("addBGPProtocols: has V4 loopback peer? %t, wanted V4 loopback peer? %t", hasV4LoopbackPeer, test.wantV4LoopbackPeer)
			}
			hasV6LoopbackPeer := len(dg.Ipv6Loopback) > 0 && len(dg.Ipv6Loopback[0].BgpIpv6Peer) > 0
			if hasV6LoopbackPeer != test.wantV6LoopbackPeer {
				t.Errorf("addBGPProtocols: has V6 loopback peer? %t, wanted V6 loopback peer? %t", hasV6LoopbackPeer, test.wantV6LoopbackPeer)
			}
		})
	}
}

func TestSRTEPolicyCounts(t *testing.T) {
	tooManyCommunities := &opb.BgpPeer_SrtePolicyGroup{
		Communities: &opb.BgpCommunities{},
	}
	for i := 0; i < 123; i++ {
		tooManyCommunities.Communities.PrivateCommunities = append(tooManyCommunities.Communities.PrivateCommunities, "1:1")
	}
	tooManySegLists := &opb.BgpPeer_SrtePolicyGroup{}
	for i := 0; i < 123; i++ {
		tooManySegLists.SegmentLists = append(tooManySegLists.SegmentLists, &opb.BgpPeer_SrtePolicyGroup_SegmentList{})
	}
	tooManySegments := &opb.BgpPeer_SrtePolicyGroup{SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{}}}
	for i := 0; i < 123; i++ {
		tooManySegments.SegmentLists[0].Segments = append(tooManySegments.SegmentLists[0].Segments, &opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment{})
	}

	tests := []struct {
		desc             string
		policyGroups     []*opb.BgpPeer_SrtePolicyGroup
		wantPolicyCounts policyCounts
		wantErr          string
	}{{
		desc:         "exceeds max community count",
		policyGroups: []*opb.BgpPeer_SrtePolicyGroup{tooManyCommunities},
		wantErr:      "only up to 32 communities",
	}, {
		desc:         "exceeds max segment list count",
		policyGroups: []*opb.BgpPeer_SrtePolicyGroup{tooManySegLists},
		wantErr:      "only up to 64 segment lists",
	}, {
		desc:         "exceeds max segment count",
		policyGroups: []*opb.BgpPeer_SrtePolicyGroup{tooManySegments},
		wantErr:      "only up to 20 segments",
	}, {
		desc: "inconsistent community count",
		policyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
			Communities: &opb.BgpCommunities{
				PrivateCommunities: []string{"1:1"}, // 1
			},
		}, {
			Communities: &opb.BgpCommunities{
				PrivateCommunities: []string{"1:1", "2:2"}, // 2
			},
		}},
		wantErr: "must have the same number of communities",
	}, {
		desc: "inconsistent community count",
		policyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
			Count: 3,
			Communities: &opb.BgpCommunities{
				PrivateCommunities: []string{"1:1", "2:2"}, // 2
			},
			SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{}, {}, {}}, // 3
		}, {
			Count: 1,
			Communities: &opb.BgpCommunities{
				PrivateCommunities: []string{"1:1", "2:2"}, // 2
			},
			SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{
				Segments: []*opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment{{}}, // 1
			}},
		}},
		wantPolicyCounts: policyCounts{
			numPolicies:     4,
			numCommunities:  2,
			maxSegListCount: 3,
			maxSegCount:     1,
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotPolicyCounts, gotErr := srtePolicyCounts(test.policyGroups)
			if ((gotErr == nil) != (test.wantErr == "")) || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("srtePolicyCounts: got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}
			if *(gotPolicyCounts) != test.wantPolicyCounts {
				t.Errorf("srtePolicyCounts: got policy counts %v, wanted %v", gotPolicyCounts, test.wantPolicyCounts)
			}
		})
	}
}

func TestBgpV4Peers(t *testing.T) {
	tests := []struct {
		desc     string
		peerCfg  *opb.BgpPeer
		wantPeer *ixconfig.TopologyBgpIpv4Peer
	}{{
		desc: "basic configuration",
		peerCfg: &opb.BgpPeer{
			Active:      true,
			Type:        opb.BgpPeer_TYPE_EXTERNAL,
			PeerAddress: "1.1.1.1",
			LocalAsn:    15169,
		},
		wantPeer: &ixconfig.TopologyBgpIpv4Peer{
			Active:             ixconfig.MultivalueTrue(),
			Type_:              ixconfig.MultivalueStr("external"),
			DutIp:              ixconfig.MultivalueStr("1.1.1.1"),
			HoldTimer:          ixconfig.MultivalueUint32(0),
			KeepaliveTimer:     ixconfig.MultivalueUint32(0),
			Enable4ByteAs:      ixconfig.MultivalueFalse(),
			LocalAs2Bytes:      ixconfig.MultivalueUint32(15169),
			ActAsRestarted:     ixconfig.MultivalueBool(false),
			AdvertiseEndOfRib:  ixconfig.MultivalueBool(false),
			RestartTime:        ixconfig.MultivalueUint32(0),
			StaleTime:          ixconfig.MultivalueUint32(0),
			NumberSRTEPolicies: ixconfig.NumberUint32(0),
			FilterIpV4Unicast:  ixconfig.MultivalueTrue(),
		},
	}, {
		desc: "with capabilities",
		peerCfg: &opb.BgpPeer{
			Active:      true,
			Type:        opb.BgpPeer_TYPE_EXTERNAL,
			PeerAddress: "1.1.1.1",
			LocalAsn:    15169,
			Capabilities: &opb.BgpPeer_Capabilities{
				Evpn: true,
			},
		},
		wantPeer: &ixconfig.TopologyBgpIpv4Peer{
			Active:                        ixconfig.MultivalueTrue(),
			Type_:                         ixconfig.MultivalueStr("external"),
			DutIp:                         ixconfig.MultivalueStr("1.1.1.1"),
			HoldTimer:                     ixconfig.MultivalueUint32(0),
			KeepaliveTimer:                ixconfig.MultivalueUint32(0),
			Enable4ByteAs:                 ixconfig.MultivalueFalse(),
			NumberSRTEPolicies:            ixconfig.NumberUint32(0),
			LocalAs2Bytes:                 ixconfig.MultivalueUint32(15169),
			ActAsRestarted:                ixconfig.MultivalueBool(false),
			AdvertiseEndOfRib:             ixconfig.MultivalueBool(false),
			RestartTime:                   ixconfig.MultivalueUint32(0),
			StaleTime:                     ixconfig.MultivalueUint32(0),
			CapabilityIpV4Unicast:         ixconfig.MultivalueFalse(),
			CapabilityIpV4Multicast:       ixconfig.MultivalueFalse(),
			EnableGracefulRestart:         ixconfig.MultivalueFalse(),
			CapabilityIpV4MplsVpn:         ixconfig.MultivalueFalse(),
			CapabilityIpV6Unicast:         ixconfig.MultivalueFalse(),
			CapabilityIpV6Multicast:       ixconfig.MultivalueFalse(),
			CapabilityIpV6MplsVpn:         ixconfig.MultivalueFalse(),
			CapabilityIpV4Mdt:             ixconfig.MultivalueFalse(),
			CapabilityVpls:                ixconfig.MultivalueFalse(),
			CapabilityIpV4MulticastVpn:    ixconfig.MultivalueFalse(),
			CapabilityIpV6MulticastVpn:    ixconfig.MultivalueFalse(),
			CapabilityRouteRefresh:        ixconfig.MultivalueFalse(),
			CapabilityRouteConstraint:     ixconfig.MultivalueFalse(),
			CapabilityLinkStateNonVpn:     ixconfig.MultivalueFalse(),
			Evpn:                          ixconfig.MultivalueTrue(),
			Ipv4MulticastBgpMplsVpn:       ixconfig.MultivalueFalse(),
			Ipv6MulticastBgpMplsVpn:       ixconfig.MultivalueFalse(),
			Capabilityipv4UnicastFlowSpec: ixconfig.MultivalueFalse(),
			Capabilityipv6UnicastFlowSpec: ixconfig.MultivalueFalse(),
			CapabilityIpv4UnicastAddPath:  ixconfig.MultivalueFalse(),
			CapabilityIpv6UnicastAddPath:  ixconfig.MultivalueFalse(),
			CapabilitySRTEPoliciesV4:      ixconfig.MultivalueFalse(),
			CapabilitySRTEPoliciesV6:      ixconfig.MultivalueFalse(),
			FilterIpV4Unicast:             ixconfig.MultivalueTrue(),
		},
	}, {
		desc: "with 4-byte ASN",
		peerCfg: &opb.BgpPeer{
			Active:      true,
			Type:        opb.BgpPeer_TYPE_EXTERNAL,
			PeerAddress: "1.1.1.1",
			LocalAsn:    151690,
		},
		wantPeer: &ixconfig.TopologyBgpIpv4Peer{
			Active:             ixconfig.MultivalueTrue(),
			Type_:              ixconfig.MultivalueStr("external"),
			DutIp:              ixconfig.MultivalueStr("1.1.1.1"),
			HoldTimer:          ixconfig.MultivalueUint32(0),
			KeepaliveTimer:     ixconfig.MultivalueUint32(0),
			Enable4ByteAs:      ixconfig.MultivalueTrue(),
			LocalAs4Bytes:      ixconfig.MultivalueUint32(151690),
			ActAsRestarted:     ixconfig.MultivalueBool(false),
			AdvertiseEndOfRib:  ixconfig.MultivalueBool(false),
			RestartTime:        ixconfig.MultivalueUint32(0),
			StaleTime:          ixconfig.MultivalueUint32(0),
			NumberSRTEPolicies: ixconfig.NumberUint32(0),
			FilterIpV4Unicast:  ixconfig.MultivalueTrue(),
		},
	}, {
		desc: "with MD5 key",
		peerCfg: &opb.BgpPeer{
			Active:      true,
			Type:        opb.BgpPeer_TYPE_EXTERNAL,
			PeerAddress: "1.1.1.1",
			LocalAsn:    15169,
			Md5Key:      "aaaa",
		},
		wantPeer: &ixconfig.TopologyBgpIpv4Peer{
			Active:             ixconfig.MultivalueTrue(),
			Type_:              ixconfig.MultivalueStr("external"),
			DutIp:              ixconfig.MultivalueStr("1.1.1.1"),
			HoldTimer:          ixconfig.MultivalueUint32(0),
			KeepaliveTimer:     ixconfig.MultivalueUint32(0),
			Enable4ByteAs:      ixconfig.MultivalueFalse(),
			LocalAs2Bytes:      ixconfig.MultivalueUint32(15169),
			ActAsRestarted:     ixconfig.MultivalueBool(false),
			AdvertiseEndOfRib:  ixconfig.MultivalueBool(false),
			RestartTime:        ixconfig.MultivalueUint32(0),
			StaleTime:          ixconfig.MultivalueUint32(0),
			Authentication:     ixconfig.MultivalueStr("md5"),
			Md5Key:             ixconfig.MultivalueStr("aaaa"),
			NumberSRTEPolicies: ixconfig.NumberUint32(0),
			FilterIpV4Unicast:  ixconfig.MultivalueTrue(),
		},
	}, {
		desc: "with SRTE policies",
		peerCfg: &opb.BgpPeer{
			Active:      true,
			Type:        opb.BgpPeer_TYPE_EXTERNAL,
			PeerAddress: "1.1.1.1",
			LocalAsn:    15169,
			SrtePolicyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
				Count:          2,
				Active:         true,
				NextHopAddress: "2.2.2.2",
				Endpoint:       "3.3.3.3",
				AsnSetMode:     opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
				Preference:     &opb.BgpPeer_SrtePolicyGroup_Preference{Preference: 170},
				Enlp:           &opb.BgpPeer_SrtePolicyGroup_Enlp{Enlp: 3},
				Binding: &opb.BgpPeer_SrtePolicyGroup_Binding{
					Type: &opb.BgpPeer_SrtePolicyGroup_Binding_NoBinding{
						NoBinding: &emptypb.Empty{},
					},
				},
				Distinguisher: 1,
				PolicyColor: &opb.UInt32IncRange{
					Start: 100,
					Step:  1,
				},
				Communities: &opb.BgpCommunities{
					PrivateCommunities: []string{"64512:0"},
				},
				SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{
					Active: true,
					Weight: &opb.BgpPeer_SrtePolicyGroup_SegmentList_Weight{Weight: 1},
					Segments: []*opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment{{
						Active: true,
						Type: &opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_MplsSid_{
							MplsSid: &opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_MplsSid{
								Label: 16,
								Tc:    0,
								S:     false,
								Ttl:   255,
							},
						},
					}},
				}},
			}},
		},
		wantPeer: &ixconfig.TopologyBgpIpv4Peer{
			Active:             ixconfig.MultivalueTrue(),
			Type_:              ixconfig.MultivalueStr("external"),
			DutIp:              ixconfig.MultivalueStr("1.1.1.1"),
			HoldTimer:          ixconfig.MultivalueUint32(0),
			KeepaliveTimer:     ixconfig.MultivalueUint32(0),
			Enable4ByteAs:      ixconfig.MultivalueFalse(),
			LocalAs2Bytes:      ixconfig.MultivalueUint32(15169),
			ActAsRestarted:     ixconfig.MultivalueBool(false),
			AdvertiseEndOfRib:  ixconfig.MultivalueBool(false),
			RestartTime:        ixconfig.MultivalueUint32(0),
			StaleTime:          ixconfig.MultivalueUint32(0),
			NumberSRTEPolicies: ixconfig.NumberUint32(2),
			FilterIpV4Unicast:  ixconfig.MultivalueTrue(),
			BgpSRTEPoliciesListV4: &ixconfig.TopologyBgpSrtePoliciesListV4{
				Active:        ixconfig.MultivalueBoolList(true, true),
				Distinguisher: ixconfig.MultivalueUintList(1, 1),
				PolicyColor:   ixconfig.MultivalueUintList(100, 101),
				PolicyType:    ixconfig.MultivalueStrList("ipv4", "ipv4"),
				EndPointV4:    ixconfig.MultivalueStrList("3.3.3.3", "3.3.3.3"),
				EndPointV6:    ixconfig.MultivalueStrList("::", "::"),

				EnableNextHop:    ixconfig.MultivalueBoolList(true, true),
				SetNextHop:       ixconfig.MultivalueStrList("manually", "manually"),
				SetNextHopIpType: ixconfig.MultivalueStrList("ipv4", "ipv4"),
				Ipv4NextHop:      ixconfig.MultivalueStrList("2.2.2.2", "2.2.2.2"),
				Ipv6NextHop:      ixconfig.MultivalueStrList("::", "::"),

				OverridePeerAsSetMode: ixconfig.MultivalueBoolList(true, true),
				AsSetMode:             ixconfig.MultivalueStrList("dontincludelocalas", "dontincludelocalas"),

				NoOfCommunities: ixconfig.NumberUint32(1),
				EnableCommunity: ixconfig.MultivalueBoolList(true, true),
				BgpCommunitiesList: []*ixconfig.TopologyBgpCommunitiesList{{
					Type_:         ixconfig.MultivalueStrList("manual", "manual"),
					AsNumber:      ixconfig.MultivalueUintList(64512, 64512),
					LastTwoOctets: ixconfig.MultivalueUintList(0, 0),
				}},

				EnableOriginatorId: ixconfig.MultivalueBoolList(false, false),
				OriginatorId:       ixconfig.MultivalueStrList("", ""),

				BgpSRTEPoliciesTunnelEncapsulationListV4: &ixconfig.TopologyBgpSrtePoliciesTunnelEncapsulationListV4{
					Active:    ixconfig.MultivalueBoolList(true, true),
					EnPrefTLV: ixconfig.MultivalueBoolList(true, true),
					PrefValue: ixconfig.MultivalueUintList(170, 170),
					EnENLPTLV: ixconfig.MultivalueBoolList(true, true),
					ENLPValue: ixconfig.MultivalueUintList(3, 3),

					EnBindingTLV:   ixconfig.MultivalueBoolList(true, true),
					BindingSIDType: ixconfig.MultivalueStrList("nobinding", "nobinding"),
					UseAsMPLSLabel: ixconfig.MultivalueBoolList(false, false),
					SID4Octet:      ixconfig.MultivalueUintList(0, 0),
					IPv6SID:        ixconfig.MultivalueStrList("::", "::"),

					NumberOfSegmentListV4: ixconfig.NumberUint32(1),
					BgpSRTEPoliciesSegmentListV4: &ixconfig.TopologyBgpSrtePoliciesSegmentListV4{
						Active:             ixconfig.MultivalueBoolList(true, true),
						EnWeight:           ixconfig.MultivalueBoolList(true, true),
						Weight:             ixconfig.MultivalueUintList(1, 1),
						NumberOfSegmentsV4: ixconfig.NumberUint32(1),
						BgpSRTEPoliciesSegmentsCollectionV4: &ixconfig.TopologyBgpSrtePoliciesSegmentsCollectionV4{
							Active:        ixconfig.MultivalueBoolList(true, true),
							SegmentType:   ixconfig.MultivalueStrList("mplssid", "mplssid"),
							Label:         ixconfig.MultivalueUintList(16, 16),
							TrafficClass:  ixconfig.MultivalueUintList(0, 0),
							BottomOfStack: ixconfig.MultivalueBoolList(false, false),
							TimeToLive:    ixconfig.MultivalueUintList(255, 255),
							Ipv6SID:       ixconfig.MultivalueStrList("::", "::"),
						},
					},
				},
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotPeers, err := bgpV4Peers([]*opb.BgpPeer{test.peerCfg})
			if err != nil {
				t.Fatalf("bgpV4Peers: unexpected error: %v", err)
			}
			if len(gotPeers) != 1 {
				t.Fatalf("bgpV4Peers: got %d peers, wanted 1", len(gotPeers))
			}
			if diff := jsonCfgDiff(t, test.wantPeer, gotPeers[0]); diff != "" {
				t.Fatalf("bgpV4Peers: unexpected BGP peer diff (-want/+got): %s", diff)
			}
		})
	}
}

func TestBgpV6Peers(t *testing.T) {
	tests := []struct {
		desc     string
		peerCfg  *opb.BgpPeer
		wantPeer *ixconfig.TopologyBgpIpv6Peer
	}{{
		desc: "basic configuration",
		peerCfg: &opb.BgpPeer{
			Active:      true,
			Type:        opb.BgpPeer_TYPE_INTERNAL,
			PeerAddress: "aa::",
			LocalAsn:    15169,
		},
		wantPeer: &ixconfig.TopologyBgpIpv6Peer{
			Active:             ixconfig.MultivalueTrue(),
			Type_:              ixconfig.MultivalueStr("internal"),
			DutIp:              ixconfig.MultivalueStr("aa::"),
			HoldTimer:          ixconfig.MultivalueUint32(0),
			KeepaliveTimer:     ixconfig.MultivalueUint32(0),
			Enable4ByteAs:      ixconfig.MultivalueFalse(),
			LocalAs2Bytes:      ixconfig.MultivalueUint32(15169),
			ActAsRestarted:     ixconfig.MultivalueBool(false),
			AdvertiseEndOfRib:  ixconfig.MultivalueBool(false),
			RestartTime:        ixconfig.MultivalueUint32(0),
			StaleTime:          ixconfig.MultivalueUint32(0),
			NumberSRTEPolicies: ixconfig.NumberUint32(0),
			FilterIpV6Unicast:  ixconfig.MultivalueTrue(),
		},
	}, {
		desc: "with capabilities",
		peerCfg: &opb.BgpPeer{
			Active:      true,
			Type:        opb.BgpPeer_TYPE_INTERNAL,
			PeerAddress: "aa::",
			LocalAsn:    15169,
			Capabilities: &opb.BgpPeer_Capabilities{
				ExtendedNextHopEncoding: true,
			},
		},
		wantPeer: &ixconfig.TopologyBgpIpv6Peer{
			Active:                           ixconfig.MultivalueTrue(),
			Type_:                            ixconfig.MultivalueStr("internal"),
			DutIp:                            ixconfig.MultivalueStr("aa::"),
			HoldTimer:                        ixconfig.MultivalueUint32(0),
			KeepaliveTimer:                   ixconfig.MultivalueUint32(0),
			Enable4ByteAs:                    ixconfig.MultivalueFalse(),
			NumberSRTEPolicies:               ixconfig.NumberUint32(0),
			LocalAs2Bytes:                    ixconfig.MultivalueUint32(15169),
			ActAsRestarted:                   ixconfig.MultivalueBool(false),
			AdvertiseEndOfRib:                ixconfig.MultivalueBool(false),
			RestartTime:                      ixconfig.MultivalueUint32(0),
			StaleTime:                        ixconfig.MultivalueUint32(0),
			CapabilityIpV4Unicast:            ixconfig.MultivalueFalse(),
			CapabilityIpV4Multicast:          ixconfig.MultivalueFalse(),
			CapabilityIpV4MplsVpn:            ixconfig.MultivalueFalse(),
			CapabilityIpV6Unicast:            ixconfig.MultivalueFalse(),
			CapabilityIpV6Multicast:          ixconfig.MultivalueFalse(),
			CapabilityIpV6MplsVpn:            ixconfig.MultivalueFalse(),
			CapabilityIpV4Mdt:                ixconfig.MultivalueFalse(),
			CapabilityVpls:                   ixconfig.MultivalueFalse(),
			EnableGracefulRestart:            ixconfig.MultivalueFalse(),
			CapabilityIpV4MulticastVpn:       ixconfig.MultivalueFalse(),
			CapabilityIpV6MulticastVpn:       ixconfig.MultivalueFalse(),
			CapabilityRouteRefresh:           ixconfig.MultivalueFalse(),
			CapabilityRouteConstraint:        ixconfig.MultivalueFalse(),
			CapabilityLinkStateNonVpn:        ixconfig.MultivalueFalse(),
			Evpn:                             ixconfig.MultivalueFalse(),
			Ipv4MulticastBgpMplsVpn:          ixconfig.MultivalueFalse(),
			Ipv6MulticastBgpMplsVpn:          ixconfig.MultivalueFalse(),
			Capabilityipv4UnicastFlowSpec:    ixconfig.MultivalueFalse(),
			Capabilityipv6UnicastFlowSpec:    ixconfig.MultivalueFalse(),
			CapabilityIpv4UnicastAddPath:     ixconfig.MultivalueFalse(),
			CapabilityIpv6UnicastAddPath:     ixconfig.MultivalueFalse(),
			CapabilitySRTEPoliciesV4:         ixconfig.MultivalueFalse(),
			CapabilitySRTEPoliciesV6:         ixconfig.MultivalueFalse(),
			CapabilityNHEncodingCapabilities: ixconfig.MultivalueTrue(),
			FilterIpV6Unicast:                ixconfig.MultivalueTrue(),
		},
	}, {
		desc: "with 4-byte ASN",
		peerCfg: &opb.BgpPeer{
			Active:      true,
			Type:        opb.BgpPeer_TYPE_INTERNAL,
			PeerAddress: "aa::",
			LocalAsn:    151690,
		},
		wantPeer: &ixconfig.TopologyBgpIpv6Peer{
			Active:             ixconfig.MultivalueTrue(),
			Type_:              ixconfig.MultivalueStr("internal"),
			DutIp:              ixconfig.MultivalueStr("aa::"),
			HoldTimer:          ixconfig.MultivalueUint32(0),
			KeepaliveTimer:     ixconfig.MultivalueUint32(0),
			Enable4ByteAs:      ixconfig.MultivalueTrue(),
			LocalAs4Bytes:      ixconfig.MultivalueUint32(151690),
			NumberSRTEPolicies: ixconfig.NumberUint32(0),
			ActAsRestarted:     ixconfig.MultivalueBool(false),
			AdvertiseEndOfRib:  ixconfig.MultivalueBool(false),
			RestartTime:        ixconfig.MultivalueUint32(0),
			StaleTime:          ixconfig.MultivalueUint32(0),
			FilterIpV6Unicast:  ixconfig.MultivalueTrue(),
		},
	}, {
		desc: "with MD5 key",
		peerCfg: &opb.BgpPeer{
			Active:      true,
			Type:        opb.BgpPeer_TYPE_INTERNAL,
			PeerAddress: "aa::",
			LocalAsn:    15169,
			Md5Key:      "aaaa",
		},
		wantPeer: &ixconfig.TopologyBgpIpv6Peer{
			Active:             ixconfig.MultivalueTrue(),
			Type_:              ixconfig.MultivalueStr("internal"),
			DutIp:              ixconfig.MultivalueStr("aa::"),
			HoldTimer:          ixconfig.MultivalueUint32(0),
			KeepaliveTimer:     ixconfig.MultivalueUint32(0),
			Enable4ByteAs:      ixconfig.MultivalueFalse(),
			LocalAs2Bytes:      ixconfig.MultivalueUint32(15169),
			ActAsRestarted:     ixconfig.MultivalueBool(false),
			AdvertiseEndOfRib:  ixconfig.MultivalueBool(false),
			RestartTime:        ixconfig.MultivalueUint32(0),
			StaleTime:          ixconfig.MultivalueUint32(0),
			Authentication:     ixconfig.MultivalueStr("md5"),
			Md5Key:             ixconfig.MultivalueStr("aaaa"),
			NumberSRTEPolicies: ixconfig.NumberUint32(0),
			FilterIpV6Unicast:  ixconfig.MultivalueTrue(),
		},
	}, {
		desc: "with SRTE policies",
		peerCfg: &opb.BgpPeer{
			Active:      true,
			Type:        opb.BgpPeer_TYPE_INTERNAL,
			PeerAddress: "aa::",
			LocalAsn:    15169,
			SrtePolicyGroups: []*opb.BgpPeer_SrtePolicyGroup{{
				Count:          2,
				Active:         true,
				NextHopAddress: "bb::",
				Endpoint:       "cc::",
				AsnSetMode:     opb.BgpAsnSetMode_ASN_SET_MODE_DO_NOT_INCLUDE,
				OriginatorId: &opb.StringIncRange{
					Start: "1.1.1.1",
					Step:  "0.0.0.2",
				},
				Preference: &opb.BgpPeer_SrtePolicyGroup_Preference{Preference: 170},
				Enlp:       &opb.BgpPeer_SrtePolicyGroup_Enlp{Enlp: 3},
				Binding: &opb.BgpPeer_SrtePolicyGroup_Binding{
					Type: &opb.BgpPeer_SrtePolicyGroup_Binding_NoBinding{
						NoBinding: &emptypb.Empty{},
					},
				},
				Distinguisher: 1,
				PolicyColor: &opb.UInt32IncRange{
					Start: 100,
					Step:  1,
				},
				Communities: &opb.BgpCommunities{
					PrivateCommunities: []string{"64512:0"},
				},
				SegmentLists: []*opb.BgpPeer_SrtePolicyGroup_SegmentList{{
					Active: true,
					Weight: &opb.BgpPeer_SrtePolicyGroup_SegmentList_Weight{Weight: 1},
					Segments: []*opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment{{
						Active: true,
						Type: &opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_MplsSid_{
							MplsSid: &opb.BgpPeer_SrtePolicyGroup_SegmentList_Segment_MplsSid{
								Label: 16,
								Tc:    0,
								S:     false,
								Ttl:   255,
							},
						},
					}},
				}},
			}},
		},
		wantPeer: &ixconfig.TopologyBgpIpv6Peer{
			Active:             ixconfig.MultivalueTrue(),
			Type_:              ixconfig.MultivalueStr("internal"),
			DutIp:              ixconfig.MultivalueStr("aa::"),
			HoldTimer:          ixconfig.MultivalueUint32(0),
			KeepaliveTimer:     ixconfig.MultivalueUint32(0),
			Enable4ByteAs:      ixconfig.MultivalueFalse(),
			LocalAs2Bytes:      ixconfig.MultivalueUint32(15169),
			ActAsRestarted:     ixconfig.MultivalueBool(false),
			AdvertiseEndOfRib:  ixconfig.MultivalueBool(false),
			RestartTime:        ixconfig.MultivalueUint32(0),
			StaleTime:          ixconfig.MultivalueUint32(0),
			NumberSRTEPolicies: ixconfig.NumberUint32(2),
			FilterIpV6Unicast:  ixconfig.MultivalueTrue(),
			BgpSRTEPoliciesListV6: &ixconfig.TopologyBgpSrtePoliciesListV6{
				Active:        ixconfig.MultivalueBoolList(true, true),
				Distinguisher: ixconfig.MultivalueUintList(1, 1),
				PolicyColor:   ixconfig.MultivalueUintList(100, 101),
				PolicyType:    ixconfig.MultivalueStrList("ipv6", "ipv6"),
				EndPointV4:    ixconfig.MultivalueStrList("0.0.0.0", "0.0.0.0"),
				EndPointV6:    ixconfig.MultivalueStrList("cc::", "cc::"),

				EnableNextHop:    ixconfig.MultivalueBoolList(true, true),
				SetNextHop:       ixconfig.MultivalueStrList("manually", "manually"),
				SetNextHopIpType: ixconfig.MultivalueStrList("ipv6", "ipv6"),
				Ipv4NextHop:      ixconfig.MultivalueStrList("0.0.0.0", "0.0.0.0"),
				Ipv6NextHop:      ixconfig.MultivalueStrList("bb::", "bb::"),

				OverridePeerAsSetMode: ixconfig.MultivalueBoolList(true, true),
				AsSetMode:             ixconfig.MultivalueStrList("dontincludelocalas", "dontincludelocalas"),

				NoOfCommunities: ixconfig.NumberUint32(1),
				EnableCommunity: ixconfig.MultivalueBoolList(true, true),
				BgpCommunitiesList: []*ixconfig.TopologyBgpCommunitiesList{{
					Type_:         ixconfig.MultivalueStrList("manual", "manual"),
					AsNumber:      ixconfig.MultivalueUintList(64512, 64512),
					LastTwoOctets: ixconfig.MultivalueUintList(0, 0),
				}},

				EnableOriginatorId: ixconfig.MultivalueBoolList(true, true),
				OriginatorId:       ixconfig.MultivalueStrList("1.1.1.1", "1.1.1.3"),

				BgpSRTEPoliciesTunnelEncapsulationListV6: &ixconfig.TopologyBgpSrtePoliciesTunnelEncapsulationListV6{
					Active:    ixconfig.MultivalueBoolList(true, true),
					EnPrefTLV: ixconfig.MultivalueBoolList(true, true),
					PrefValue: ixconfig.MultivalueUintList(170, 170),
					EnENLPTLV: ixconfig.MultivalueBoolList(true, true),
					ENLPValue: ixconfig.MultivalueUintList(3, 3),

					EnBindingTLV:   ixconfig.MultivalueBoolList(true, true),
					BindingSIDType: ixconfig.MultivalueStrList("nobinding", "nobinding"),
					UseAsMPLSLabel: ixconfig.MultivalueBoolList(false, false),
					SID4Octet:      ixconfig.MultivalueUintList(0, 0),
					IPv6SID:        ixconfig.MultivalueStrList("::", "::"),

					NumberOfSegmentListV6: ixconfig.NumberUint32(1),
					BgpSRTEPoliciesSegmentListV6: &ixconfig.TopologyBgpSrtePoliciesSegmentListV6{
						Active:             ixconfig.MultivalueBoolList(true, true),
						EnWeight:           ixconfig.MultivalueBoolList(true, true),
						Weight:             ixconfig.MultivalueUintList(1, 1),
						NumberOfSegmentsV6: ixconfig.NumberUint32(1),
						BgpSRTEPoliciesSegmentsCollectionV6: &ixconfig.TopologyBgpSrtePoliciesSegmentsCollectionV6{
							Active:        ixconfig.MultivalueBoolList(true, true),
							SegmentType:   ixconfig.MultivalueStrList("mplssid", "mplssid"),
							Label:         ixconfig.MultivalueUintList(16, 16),
							TrafficClass:  ixconfig.MultivalueUintList(0, 0),
							BottomOfStack: ixconfig.MultivalueBoolList(false, false),
							TimeToLive:    ixconfig.MultivalueUintList(255, 255),
							Ipv6SID:       ixconfig.MultivalueStrList("::", "::"),
						},
					},
				},
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			gotPeers, err := bgpV6Peers([]*opb.BgpPeer{test.peerCfg})
			if err != nil {
				t.Fatalf("bgpV6Peers: unexpected error: %v", err)
			}
			if len(gotPeers) != 1 {
				t.Fatalf("bgpV6Peers: got %d peers, wanted 1", len(gotPeers))
			}
			if diff := jsonCfgDiff(t, test.wantPeer, gotPeers[0]); diff != "" {
				t.Fatalf("bgpV6Peers: unexpected BGP peer diff (-want/+got): %s", diff)
			}
		})
	}
}

func TestAddRSVPProtocols(t *testing.T) {
	const (
		ifName  = "someIntf"
		isrName = "isr"
	)

	clientWithTopo := func(hasIPv4, hasISISNetwGrp bool, numISISNodes int) *ixATE {
		var ipv4 *ixconfig.TopologyIpv4
		if hasIPv4 {
			ipv4 = &ixconfig.TopologyIpv4{}
		}
		var isisNetwGrp *ixconfig.TopologyNetworkGroup
		if hasISISNetwGrp {
			isisNetwGrp = &ixconfig.TopologyNetworkGroup{
				NetworkTopology: &ixconfig.TopologyNetworkTopology{
					NetTopologyLinear: &ixconfig.TopologyNetTopologyLinear{
						Nodes: ixconfig.NumberInt(numISISNodes),
					},
				},
			}
		}
		return &ixATE{
			intfs: map[string]*intf{
				ifName: &intf{
					ipv4: ipv4,
					isrToNetworkGroup: map[string]*ixconfig.TopologyNetworkGroup{
						isrName: isisNetwGrp,
					},
				},
			},
		}
	}

	tests := []struct {
		desc                    string
		hasIPv4, hasISISNetwGrp bool
		numISISNodes            int
		rsvp                    *opb.RsvpConfig
		wantErr                 string
		wantIfCfg               *ixconfig.TopologyRsvpteIf
		wantDeviceGroup         *ixconfig.TopologyDeviceGroup
	}{{
		desc: "no IPv4 interface configured",
		rsvp: &opb.RsvpConfig{
			IsReachabilityName: isrName,
			Loopbacks:          []*opb.RsvpConfig_Loopback{{LocalIpCidr: "1.1.1.1/32"}},
		},
		wantErr: "could not find IPv4",
	}, {
		desc: "no reachability configured",
		rsvp: &opb.RsvpConfig{
			IsReachabilityName: isrName,
			Loopbacks:          []*opb.RsvpConfig_Loopback{{LocalIpCidr: "1.1.1.1/32"}},
		},
		hasIPv4: true,
		wantErr: "could not find IS-IS reachability config",
	}, {
		desc: "no IS-IS nodes configured",
		rsvp: &opb.RsvpConfig{
			IsReachabilityName: isrName,
			Loopbacks:          []*opb.RsvpConfig_Loopback{{LocalIpCidr: "1.1.1.1/32"}},
		},
		hasIPv4:        true,
		hasISISNetwGrp: true,
		wantErr:        "cannot add more loopbacks",
	}, {
		desc: "bad loopback ip",
		rsvp: &opb.RsvpConfig{
			IsReachabilityName: isrName,
			Loopbacks:          []*opb.RsvpConfig_Loopback{{LocalIpCidr: "aa::/127"}},
		},
		hasIPv4:        true,
		hasISISNetwGrp: true,
		numISISNodes:   2,
		wantErr:        "could not parse",
	}, {
		desc: "valid egress-only config with bundle messages and refresh reduction",
		rsvp: &opb.RsvpConfig{
			IsReachabilityName: isrName,
			Loopbacks: []*opb.RsvpConfig_Loopback{
				{LocalIpCidr: "1.1.1.1/32"},
				{LocalIpCidr: "2.2.2.2/31"},
			},
			BundleMessageSending: true,
			RefreshReduction:     true,
		},
		hasIPv4:        true,
		hasISISNetwGrp: true,
		numISISNodes:   2,
		wantIfCfg: &ixconfig.TopologyRsvpteIf{
			Active:                     ixconfig.MultivalueTrue(),
			EnableHelloExtension:       ixconfig.MultivalueTrue(),
			EnableBundleMessageSending: ixconfig.MultivalueTrue(),
			EnableRefreshReduction:     ixconfig.MultivalueTrue(),
		},
		wantDeviceGroup: &ixconfig.TopologyDeviceGroup{
			Multiplier: ixconfig.NumberUint32(2),
			Ipv4Loopback: []*ixconfig.TopologyIpv4Loopback{{
				Address: ixconfig.MultivalueStrList("1.1.1.1", "2.2.2.2"),
				Prefix:  ixconfig.MultivalueUintList(32, 31),
				RsvpteLsps: []*ixconfig.TopologyRsvpteLsps{{
					Name:              ixconfig.String("RSVP LSPs 0 for someIntf"),
					Active:            ixconfig.MultivalueTrue(),
					EnableP2PEgress:   ixconfig.Bool(true),
					IngressP2PLsps:    ixconfig.NumberUint32(0),
					RsvpP2PEgressLsps: &ixconfig.TopologyRsvpP2PEgressLsps{},
					RsvpP2PIngressLsps: &ixconfig.TopologyRsvpP2PIngressLsps{
						Active: ixconfig.MultivalueFalse(),
					},
				}},
			}},
		},
	}, {
		desc: "mismatched ingress LSP counds",
		rsvp: &opb.RsvpConfig{
			IsReachabilityName: isrName,
			Loopbacks: []*opb.RsvpConfig_Loopback{{
				LocalIpCidr: "1.1.1.1/32",
			}, {
				LocalIpCidr: "2.2.2.2/31",
				IngressLsps: []*opb.RsvpConfig_Loopback_IngressLSP{{}},
			}},
		},
		hasIPv4:        true,
		hasISISNetwGrp: true,
		numISISNodes:   2,
		wantErr:        "ingress LSP count must match",
	}, {
		desc: "mismatched ingress LSP ERO counts",
		rsvp: &opb.RsvpConfig{
			IsReachabilityName: isrName,
			Loopbacks: []*opb.RsvpConfig_Loopback{{
				LocalIpCidr: "1.1.1.1/32",
				IngressLsps: []*opb.RsvpConfig_Loopback_IngressLSP{{
					Eros: []*opb.RsvpConfig_Loopback_IngressLSP_ERO{{}},
				}},
			}, {
				LocalIpCidr: "2.2.2.2/31",
				IngressLsps: []*opb.RsvpConfig_Loopback_IngressLSP{{
					Eros: []*opb.RsvpConfig_Loopback_IngressLSP_ERO{{}, {}},
				}},
			}},
		},
		hasIPv4:        true,
		hasISISNetwGrp: true,
		numISISNodes:   2,
		wantErr:        "ERO count must match",
	}, {
		desc: "mismatched ingress LSP RRO counts",
		rsvp: &opb.RsvpConfig{
			IsReachabilityName: isrName,
			Loopbacks: []*opb.RsvpConfig_Loopback{{
				LocalIpCidr: "1.1.1.1/32",
				IngressLsps: []*opb.RsvpConfig_Loopback_IngressLSP{{
					Rros: []*opb.RsvpConfig_Loopback_IngressLSP_RRO{{}},
				}},
			}, {
				LocalIpCidr: "2.2.2.2/31",
				IngressLsps: []*opb.RsvpConfig_Loopback_IngressLSP{{
					Rros: []*opb.RsvpConfig_Loopback_IngressLSP_RRO{{}, {}},
				}},
			}},
		},
		hasIPv4:        true,
		hasISISNetwGrp: true,
		numISISNodes:   2,
		wantErr:        "RRO count must match",
	}, {
		desc: "bad ingress LSP remote address",
		rsvp: &opb.RsvpConfig{
			IsReachabilityName: isrName,
			Loopbacks: []*opb.RsvpConfig_Loopback{{
				LocalIpCidr: "1.1.1.1/32",
				IngressLsps: []*opb.RsvpConfig_Loopback_IngressLSP{{
					RemoteIpCidr: "aa::/127",
				}},
			}},
		},
		hasIPv4:        true,
		hasISISNetwGrp: true,
		numISISNodes:   1,
		wantErr:        "could not parse",
	}, {
		desc: "valid config with ingress LSPs",
		rsvp: &opb.RsvpConfig{
			IsReachabilityName: isrName,
			Loopbacks: []*opb.RsvpConfig_Loopback{{
				LocalIpCidr: "1.1.1.1/32",
				IngressLsps: []*opb.RsvpConfig_Loopback_IngressLSP{{
					RemoteIpCidr: "3.3.3.3/31",
					Eros: []*opb.RsvpConfig_Loopback_IngressLSP_ERO{
						{Ipv4Cidr: "5.5.5.5/31"},
						{Ipv4Cidr: "7.7.7.7/32"},
					},
				}, {
					RemoteIpCidr:        "9.9.9.9/32",
					LocalProtection:     true,
					BandwidthProtection: true,
					TunnelId:            1,
					LspId:               2,
				}},
			}, {
				LocalIpCidr: "2.2.2.2/31",
				IngressLsps: []*opb.RsvpConfig_Loopback_IngressLSP{{
					RemoteIpCidr:       "4.4.4.4/32",
					FastReroute:        true,
					PathReoptimization: true,
					TunnelId:           3,
					LspId:              4,
				}, {
					RemoteIpCidr: "6.6.6.6/31",
					Rros: []*opb.RsvpConfig_Loopback_IngressLSP_RRO{
						{Ipv4: "8.8.8.8"},
						{Ipv4: "10.10.10.10"},
					},
				}},
			}},
		},
		hasIPv4:        true,
		hasISISNetwGrp: true,
		numISISNodes:   2,
		wantIfCfg: &ixconfig.TopologyRsvpteIf{
			Active:                     ixconfig.MultivalueTrue(),
			EnableHelloExtension:       ixconfig.MultivalueTrue(),
			EnableBundleMessageSending: ixconfig.MultivalueFalse(),
			EnableRefreshReduction:     ixconfig.MultivalueFalse(),
		},
		wantDeviceGroup: &ixconfig.TopologyDeviceGroup{
			Multiplier: ixconfig.NumberUint32(2),
			Ipv4Loopback: []*ixconfig.TopologyIpv4Loopback{{
				Address: ixconfig.MultivalueStrList("1.1.1.1", "2.2.2.2"),
				Prefix:  ixconfig.MultivalueUintList(32, 31),
				RsvpteLsps: []*ixconfig.TopologyRsvpteLsps{{
					Name:              ixconfig.String("RSVP LSPs 0 for someIntf"),
					Active:            ixconfig.MultivalueTrue(),
					EnableP2PEgress:   ixconfig.Bool(true),
					IngressP2PLsps:    ixconfig.NumberUint32(2),
					RsvpP2PEgressLsps: &ixconfig.TopologyRsvpP2PEgressLsps{},
					RsvpP2PIngressLsps: &ixconfig.TopologyRsvpP2PIngressLsps{
						Active:                ixconfig.MultivalueTrue(),
						NumberOfEroSubObjects: ixconfig.NumberUint32(2),
						EnableEro:             ixconfig.MultivalueBoolList(true, false, false, false),
						RsvpEROSubObjectsList: []*ixconfig.TopologyRsvpEroSubObjectsList{{
							Ip:           ixconfig.MultivalueStrList("5.5.5.5", "", "", ""),
							PrefixLength: ixconfig.MultivalueUintList(31, 0, 0, 0),
						}, {
							Ip:           ixconfig.MultivalueStrList("7.7.7.7", "", "", ""),
							PrefixLength: ixconfig.MultivalueUintList(32, 0, 0, 0),
						}},
						NumberOfRroSubObjects: ixconfig.NumberUint32(2),
						SendRro:               ixconfig.MultivalueBoolList(false, false, false, true),
						RsvpIngressRROSubObjectsList: []*ixconfig.TopologyRsvpIngressRroSubObjectsList{{
							Ip: ixconfig.MultivalueStrList("", "", "", "8.8.8.8"),
						}, {
							Ip: ixconfig.MultivalueStrList("", "", "", "10.10.10.10"),
						}},
						RemoteIp:                   ixconfig.MultivalueStrList("3.3.3.3", "9.9.9.9", "4.4.4.4", "6.6.6.6"),
						PrefixLength:               ixconfig.MultivalueUintList(31, 32, 32, 31),
						LocalProtectionDesired:     ixconfig.MultivalueBoolList(false, true, false, false),
						BandwidthProtectionDesired: ixconfig.MultivalueBoolList(false, true, false, false),
						EnableFastReroute:          ixconfig.MultivalueBoolList(false, false, true, false),
						EnablePathReOptimization:   ixconfig.MultivalueBoolList(false, false, true, false),
						TunnelId:                   ixconfig.MultivalueUintList(0, 1, 3, 0),
						LspId:                      ixconfig.MultivalueUintList(0, 2, 4, 0),
					},
				}},
			}},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := clientWithTopo(test.hasIPv4, test.hasISISNetwGrp, test.numISISNodes)
			ifc := &opb.InterfaceConfig{
				Name:  ifName,
				Rsvps: []*opb.RsvpConfig{test.rsvp},
			}
			gotErr := c.addRSVPProtocols(ifc)
			if (gotErr == nil && test.wantErr != "") || (gotErr != nil && test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("addRSVPProtocols got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}
			intf := c.intfs[ifName]
			if diff := jsonCfgDiff(t, test.wantIfCfg, intf.ipv4.RsvpteIf[0]); diff != "" {
				t.Fatalf("addRSVPProtocols: unexpected RSVP-TE interface diff (-want/+got): %s", diff)
			}
			if diff := jsonCfgDiff(t, test.wantDeviceGroup, intf.isrToNetworkGroup[isrName].DeviceGroup[0]); diff != "" {
				t.Fatalf("addRSVPProtocols: unexpected RSVP-TE device group (-want/+got): %s", diff)
			}
		})
	}
}

func TestAddDHCPProtocols(t *testing.T) {
	const intfName = "intf1"
	clientWithTopo := func() *ixATE {
		cfg := &ixconfig.Ixnetwork{
			Topology: []*ixconfig.Topology{{
				DeviceGroup: []*ixconfig.TopologyDeviceGroup{{
					Ethernet: []*ixconfig.TopologyEthernet{&ixconfig.TopologyEthernet{
						Ipv6: []*ixconfig.TopologyIpv6{&ixconfig.TopologyIpv6{}},
					}}}},
			}},
		}
		return &ixATE{
			cfg: cfg,
			intfs: map[string]*intf{
				intfName: &intf{
					deviceGroup: cfg.Topology[0].DeviceGroup[0],
					ipv6:        cfg.Topology[0].DeviceGroup[0].Ethernet[0].Ipv6[0],
				},
			},
		}
	}

	tests := []struct {
		desc         string
		intf         *opb.InterfaceConfig
		wantV6Client *ixconfig.TopologyDhcpv6client
		wantV6Server *ixconfig.TopologyDhcpv6server
	}{{
		desc: "dhcp v6 client",
		intf: &opb.InterfaceConfig{
			Name:         intfName,
			Dhcpv6Client: &opb.DhcpV6Client{},
		},
		wantV6Client: &ixconfig.TopologyDhcpv6client{},
	}, {
		desc: "dhcp v6 server",
		intf: &opb.InterfaceConfig{
			Name: intfName,
			Dhcpv6Server: &opb.DhcpV6Server{
				LeaseAddrs: &opb.AddressRange{
					Min:   "::10",
					Max:   "::80",
					Count: 5,
				},
			},
		},
		wantV6Server: &ixconfig.TopologyDhcpv6server{
			Dhcp6ServerSessions: &ixconfig.TopologyDhcp6ServerSessions{
				IpAddress:          ixconfig.MultivalueStr("::10"),
				IpAddressIncrement: ixconfig.MultivalueStr("::16"),
				PoolSize:           ixconfig.MultivalueUint32(5),
			},
		},
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := clientWithTopo()
			if err := c.addDHCPProtocols(test.intf); err != nil {
				t.Fatalf("addDHCPProtocols: unexpected error: %v", err)
			}

			intf := c.intfs[intfName]
			var gotV6Client *ixconfig.TopologyDhcpv6client
			if gotV6Clients := intf.deviceGroup.Ethernet[0].Dhcpv6client; len(gotV6Clients) > 0 {
				gotV6Client = gotV6Clients[0]
			}
			if diff := jsonCfgDiff(t, test.wantV6Client, gotV6Client); diff != "" {
				t.Fatalf("addDHCPProtocols: unexpected v6 client config (-want/+got): %s", diff)
			}
			var gotV6Server *ixconfig.TopologyDhcpv6server
			if gotV6Servers := intf.ipv6.Dhcpv6server; len(gotV6Servers) > 0 {
				gotV6Server = gotV6Servers[0]
			}
			if diff := jsonCfgDiff(t, test.wantV6Server, gotV6Server); diff != "" {
				t.Fatalf("addDHCPProtocols: unexpected v6 server config (-want/+got): %s", diff)
			}
		})
	}
}
