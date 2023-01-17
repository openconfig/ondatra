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
)

func TestAddPorts(t *testing.T) {
	const (
		chassisHost         = "192.168.1.1"
		portInstrumentation = "floating"
		firstPort           = "1/1"
		secondPort          = "2/2"
	)
	wantCfg := &ixconfig.Ixnetwork{
		Vport: []*ixconfig.Vport{{
			Name:     ixconfig.String("1/1"),
			Location: ixconfig.String(fmt.Sprintf("%s;1;1", chassisHost)),
			L1Config: &ixconfig.VportL1Config{
				AresOneFourHundredGigLan: &ixconfig.VportL1ConfigAresOneFourHundredGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
				NovusHundredGigLan: &ixconfig.VportL1ConfigNovusHundredGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
				NovusTenGigLan: &ixconfig.VportL1ConfigNovusTenGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
			},
			Protocols: &ixconfig.VportProtocols{},
			ProtocolStack: &ixconfig.VportProtocolStack{
				Options: &ixconfig.VportProtocolStackOptions{
					McastSolicit: ixconfig.NumberUint32(8),
					RetransTime:  ixconfig.NumberUint32(7000),
				},
			},
		}, {
			Name:     ixconfig.String("2/2"),
			Location: ixconfig.String(fmt.Sprintf("%s;2;2", chassisHost)),
			L1Config: &ixconfig.VportL1Config{
				AresOneFourHundredGigLan: &ixconfig.VportL1ConfigAresOneFourHundredGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
				NovusHundredGigLan: &ixconfig.VportL1ConfigNovusHundredGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
				NovusTenGigLan: &ixconfig.VportL1ConfigNovusTenGigLan{
					AutoInstrumentation: ixconfig.String(portInstrumentation),
				},
			},
			Protocols: &ixconfig.VportProtocols{},
			ProtocolStack: &ixconfig.VportProtocolStack{
				Options: &ixconfig.VportProtocolStackOptions{
					McastSolicit: ixconfig.NumberUint32(8),
					RetransTime:  ixconfig.NumberUint32(7000),
				},
			},
		}},
	}

	c := &ixATE{
		chassisHost: chassisHost,
		cfg:         &ixconfig.Ixnetwork{},
		ports:       make(map[string]*ixconfig.Vport),
	}

	top := &Topology{
		// Invert order in function call to test alphabetization.
		Interfaces: []*opb.InterfaceConfig{
			{Link: &opb.InterfaceConfig_Port{secondPort}},
			// Test with multiple interfaces on a port.
			{Link: &opb.InterfaceConfig_Port{firstPort}},
			{Link: &opb.InterfaceConfig_Port{firstPort}},
		},
	}
	if err := c.addPorts(top); err != nil {
		t.Fatal(err)
	}
	if diff := jsonCfgDiff(t, wantCfg, c.cfg); diff != "" {
		t.Fatalf("addPorts: Incorrect config generated: diff (-want +got)\n%s", diff)
	}
	if diff := jsonCfgDiff(t, wantCfg.Vport[0], c.ports[firstPort]); diff != "" {
		t.Fatalf("addPorts: Incorrect config for port %s: diff (-want +got)\n%s", firstPort, diff)
	}
	if diff := jsonCfgDiff(t, wantCfg.Vport[1], c.ports[secondPort]); diff != "" {
		t.Fatalf("addPorts: Incorrect config for port %s: diff (-want +got)\n%s", secondPort, diff)
	}
}

func TestAddPortsErrors(t *testing.T) {
	c := &ixATE{}

	tests := []struct {
		desc    string
		top     *Topology
		wantErr string
	}{{
		desc: "port in two lags",
		top: &Topology{
			LAGs: []*opb.Lag{{
				Ports: []string{"1/1"},
			}, {
				Ports: []string{"1/1"},
			}},
		},
		wantErr: "belongs to two lags",
	}, {
		desc: "interface on port which belongs to lag",
		top: &Topology{
			LAGs:       []*opb.Lag{{Ports: []string{"1/1"}}},
			Interfaces: []*opb.InterfaceConfig{{Link: &opb.InterfaceConfig_Port{"1/1"}}},
		},
		wantErr: "already belongs to lag",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := c.addPorts(test.top)
			if err == nil || !strings.Contains(err.Error(), test.wantErr) {
				t.Fatalf("addPorts: Expected error %q but got %v", test.wantErr, err)
			}
		})
	}
}

func TestAddLAGs(t *testing.T) {
	vport1 := &ixconfig.Vport{
		Name:     ixconfig.String("1/1"),
		Location: ixconfig.String("192.168.1.1;1;1"),
	}
	vport2 := &ixconfig.Vport{
		Name:     ixconfig.String("2/2"),
		Location: ixconfig.String("192.168.1.1;2;2"),
	}

	tests := []struct {
		desc    string
		lags    []*opb.Lag
		wantCfg *ixconfig.Ixnetwork
		wantErr string
	}{{
		desc: "static",
		lags: []*opb.Lag{{
			Name:  "staticLAG",
			Ports: []string{"1/1", "2/2"},
		}},
		wantCfg: func() *ixconfig.Ixnetwork {
			cfg := &ixconfig.Ixnetwork{
				Vport: []*ixconfig.Vport{vport1, vport2},
				Lag: []*ixconfig.Lag{{
					Name:    ixconfig.String("staticLAG"),
					LagMode: &ixconfig.LagLagMode{LagProtocol: ixconfig.MultivalueStr("staticlag")},
					ProtocolStack: &ixconfig.LagProtocolStack{
						Enabled: ixconfig.MultivalueBool(true),
						Ethernet: []*ixconfig.LagEthernet{{
							Multiplier: ixconfig.NumberFloat64(1),
							Lagportstaticlag: []*ixconfig.LagLagportstaticlag{{
								Multiplier: ixconfig.NumberFloat64(1),
							}},
						}},
						Multiplier: ixconfig.NumberFloat64(1),
					},
				}},
			}
			cfg.Lag[0].SetVportsRefs([]ixconfig.IxiaCfgNode{vport1, vport2})
			return cfg
		}(),
	}, {
		desc: "lacp",
		lags: []*opb.Lag{{
			Name:  "lacpLAG",
			Ports: []string{"1/1", "2/2"},
			Lacp:  &opb.Lag_Lacp{Enabled: true},
		}},
		wantCfg: func() *ixconfig.Ixnetwork {
			cfg := &ixconfig.Ixnetwork{
				Vport: []*ixconfig.Vport{vport1, vport2},
				Lag: []*ixconfig.Lag{{
					Name:    ixconfig.String("lacpLAG"),
					LagMode: &ixconfig.LagLagMode{LagProtocol: ixconfig.MultivalueStr("lacp")},
					ProtocolStack: &ixconfig.LagProtocolStack{
						Enabled: ixconfig.MultivalueBool(true),
						Ethernet: []*ixconfig.LagEthernet{{
							Multiplier: ixconfig.NumberFloat64(1),
							Lagportlacp: []*ixconfig.LagLagportlacp{{
								Multiplier: ixconfig.NumberFloat64(1),
							}},
						}},
						Multiplier: ixconfig.NumberFloat64(1),
					},
				}},
			}
			cfg.Lag[0].SetVportsRefs([]ixconfig.IxiaCfgNode{vport1, vport2})
			return cfg
		}(),
	}, {
		desc: "port in multiple lags",
		lags: []*opb.Lag{{
			Name:  "lag1",
			Ports: []string{"1/1"},
		}, {
			Name:  "lag2",
			Ports: []string{"1/1", "1/2"},
		}},
		wantErr: "multiple LAGs",
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &ixATE{
				name:        "ixia1",
				chassisHost: "192.168.1.1",
				cfg: &ixconfig.Ixnetwork{
					Vport: []*ixconfig.Vport{vport1, vport2},
				},
				ports: map[string]*ixconfig.Vport{
					"1/1": vport1,
					"2/2": vport2,
				},
				lags:     make(map[string]*ixconfig.Lag),
				lagPorts: make(map[*ixconfig.Lag][]*ixconfig.Vport),
			}

			err := c.addLAGs(test.lags)
			if (err == nil) != (test.wantErr == "") || err != nil && !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("addLags got error %v want %s", err, test.wantErr)
			}
			if err != nil {
				return
			}
			wantCfg, err := test.wantCfg.ResolvedConfig(test.wantCfg)
			if err != nil {
				t.Fatalf("could not get resolved expected cfg: %v", err)
			}
			gotCfg, err := c.cfg.ResolvedConfig(c.cfg)
			if err != nil {
				t.Fatalf("could not get resolved cfg: %v", err)
			}
			if diff := jsonCfgDiff(t, wantCfg, gotCfg); diff != "" {
				t.Errorf("addLAGs: Incorrect config generated: diff (-want +got)\n%s", diff)
			}
		})
	}
}

func TestAddTopologies(t *testing.T) {
	const (
		lagName = "testLag"
		port    = "1/1"
		ifName  = "someIntf"
	)
	vport := &ixconfig.Vport{
		Name:     ixconfig.String(port),
		Location: ixconfig.String("192.168.1.1;1;1"),
		L1Config: &ixconfig.VportL1Config{},
	}
	lag := &ixconfig.Lag{
		Name:   ixconfig.String(lagName),
		Vports: []string{port},
	}
	emptyCfgClient := func() *ixATE {
		cfg := &ixconfig.Ixnetwork{
			Vport: []*ixconfig.Vport{vport},
			Lag:   []*ixconfig.Lag{lag},
		}
		return &ixATE{
			cfg:   cfg,
			ports: map[string]*ixconfig.Vport{port: vport},
			lags:  map[string]*ixconfig.Lag{lagName: lag},
			intfs: make(map[string]*intf),
		}
	}
	tests := []struct {
		desc                 string
		ifs                  []*opb.InterfaceConfig
		wantDg, wantNestedDg string
		wantLink             ixconfig.IxiaCfgNode
		wantErr              string
	}{{
		desc: "Port does not exist",
		ifs: []*opb.InterfaceConfig{{
			Name:     ifName,
			Link:     &opb.InterfaceConfig_Port{"invalid"},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
		}},
		wantDg:   fmt.Sprintf("Device Group on %s", ifName),
		wantLink: vport,
		wantErr:  "no ATE port configured",
	}, {
		desc: "LAG does not exist",
		ifs: []*opb.InterfaceConfig{{
			Name:     ifName,
			Link:     &opb.InterfaceConfig_Lag{"invalid"},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
		}},
		wantDg:   fmt.Sprintf("Device Group on %s", ifName),
		wantLink: vport,
		wantErr:  "no ATE LAG configured",
	}, {
		desc: "No LAG",
		ifs: []*opb.InterfaceConfig{{
			Name:     ifName,
			Link:     &opb.InterfaceConfig_Port{port},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
		}},
		wantDg:   fmt.Sprintf("Device Group on %s", ifName),
		wantLink: vport,
	}, {
		desc: "With LACP",
		ifs: []*opb.InterfaceConfig{{
			Name:       ifName,
			Link:       &opb.InterfaceConfig_Port{port},
			Ethernet:   &opb.EthernetConfig{Mtu: 1500},
			EnableLacp: true,
		}},
		wantDg:       fmt.Sprintf("LACP DeviceGroup on %s", ifName),
		wantNestedDg: fmt.Sprintf("Device Group on %s", ifName),
		wantLink:     vport,
	}, {
		desc: "With LAG",
		ifs: []*opb.InterfaceConfig{{
			Name:     ifName,
			Link:     &opb.InterfaceConfig_Lag{"testLag"},
			Ethernet: &opb.EthernetConfig{Mtu: 1500},
		}},
		wantDg:   fmt.Sprintf("Device Group on %s", ifName),
		wantLink: lag,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := emptyCfgClient()
			gotErr := c.addTopology(test.ifs)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("addTopology: got err: %v, want err %q", gotErr, test.wantErr)
			}
			if gotErr != nil {
				return
			}

			var gotDg, gotNestedDg string
			dg := c.cfg.Topology[0].DeviceGroup[0]
			gotDg = *(dg.Name)
			if len(dg.DeviceGroup) > 0 {
				gotNestedDg = *(dg.DeviceGroup[0].Name)
			}
			if gotDg != test.wantDg {
				t.Errorf("addTopology: did not find expected device group: got %q, want %q", gotDg, test.wantDg)
			}
			if gotNestedDg != test.wantNestedDg {
				t.Errorf("addTopology: did not find expected nested device group: got %q, want %q", gotNestedDg, test.wantNestedDg)
			}

			if c.intfs[ifName].deviceGroup == nil {
				t.Fatalf("addTopology: no mapped device group for interface %q", ifName)
			}
			if c.intfs[ifName].link != test.wantLink {
				t.Fatalf("addTopology: did not find expected vport for interface %q", ifName)
			}
		})
	}
}
