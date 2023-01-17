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
	"sort"
	"strings"

	"github.com/openconfig/ondatra/internal/ixconfig"

	opb "github.com/openconfig/ondatra/proto"
)

func (ix *ixATE) addPorts(top *Topology) error {
	ports := make(map[string]bool)
	portToLag := make(map[string]*opb.Lag)
	for _, ol := range top.LAGs {
		for _, p := range ol.GetPorts() {
			if ol2 := portToLag[p]; ol2 != nil {
				return fmt.Errorf("port %s belongs to two lags: %v and %v", p, ol, ol2)
			}
			portToLag[p] = ol
			ports[p] = true
		}
	}
	for _, intf := range top.Interfaces {
		if lp, ok := intf.GetLink().(*opb.InterfaceConfig_Port); ok {
			if ol := portToLag[lp.Port]; ol != nil {
				return fmt.Errorf("interface %v on port %s which already belongs to lag %v", intf, lp.Port, ol)
			}
			ports[lp.Port] = true
		}
	}

	portList := make([]string, 0, len(ports))
	for p := range ports {
		portList = append(portList, p)
	}
	// Alphabetize ports so that a particular reservation produces the same config.
	sort.Strings(portList)
	for _, port := range portList {
		vport := &ixconfig.Vport{
			Name:      ixconfig.String(port),
			Location:  ixconfig.String(fmt.Sprintf("%s;%s", ix.chassisHost, strings.ReplaceAll(port, "/", ";"))),
			L1Config:  &ixconfig.VportL1Config{},
			Protocols: &ixconfig.VportProtocols{}, // Used as the target for traffic endpoints
			ProtocolStack: &ixconfig.VportProtocolStack{
				Options: &ixconfig.VportProtocolStackOptions{
					// (b/192980845) Configure more neighbor solicitations spaced further apart to workaround IPv6 failures behind LACP.
					McastSolicit: ixconfig.NumberUint32(8),
					RetransTime:  ixconfig.NumberUint32(7000),
				},
			},
		}
		// TODO (team): Set this based on the actual hardware type of the Ixia chassis port.
		const portInstrumentation = "floating"
		aresOneFourHundredGigLan(vport.L1Config).AutoInstrumentation = ixconfig.String(portInstrumentation)
		novusHundredGigLan(vport.L1Config).AutoInstrumentation = ixconfig.String(portInstrumentation)
		novusTenGigLan(vport.L1Config).AutoInstrumentation = ixconfig.String(portInstrumentation)
		ix.cfg.Vport = append(ix.cfg.Vport, vport)
		ix.ports[port] = vport
	}
	return nil
}

func (ix *ixATE) addLAGs(lags []*opb.Lag) error {
	portToLag := make(map[string]*opb.Lag)
	for _, ol := range lags {
		lag := &ixconfig.Lag{
			Name:    ixconfig.String(ol.GetName()),
			LagMode: &ixconfig.LagLagMode{},
			ProtocolStack: &ixconfig.LagProtocolStack{
				Multiplier: ixconfig.NumberFloat64(1),
				Enabled:    ixconfig.MultivalueBool(true),
				Ethernet:   []*ixconfig.LagEthernet{{Multiplier: ixconfig.NumberFloat64(1)}},
			},
		}
		if ol.GetLacp().GetEnabled() {
			lag.LagMode.LagProtocol = ixconfig.MultivalueStr("lacp")
			lag.ProtocolStack.Ethernet[0].Lagportlacp = []*ixconfig.LagLagportlacp{{
				Multiplier: ixconfig.NumberFloat64(1),
			}}
		} else {
			lag.LagMode.LagProtocol = ixconfig.MultivalueStr("staticlag")
			lag.ProtocolStack.Ethernet[0].Lagportstaticlag = []*ixconfig.LagLagportstaticlag{{
				Multiplier: ixconfig.NumberFloat64(1),
			}}
		}

		var vports []*ixconfig.Vport
		for _, p := range ol.GetPorts() {
			if prevLag, ok := portToLag[p]; ok {
				return fmt.Errorf("port %s belongs to multiple LAGs: %s and %s", p, prevLag.GetName(), ol.GetName())
			}
			portToLag[p] = ol
			vport := ix.ports[p]
			vports = append(vports, vport)
			lag.AppendVportsRef(vport)
		}
		// Alphabetize ports so that a particular reservation produces the same config.
		sort.Strings(lag.Vports)
		ix.cfg.Lag = append(ix.cfg.Lag, lag)
		ix.lags[ol.GetName()] = lag
		ix.lagPorts[lag] = vports
	}
	// Alphabetize lags so that a particular reservation produces the same config.
	sort.Slice(ix.cfg.Lag, func(i int, j int) bool {
		return lags[i].GetName() <= lags[j].GetName()
	})
	return nil
}

// addTopology adds an IxNetwork topology with ports assigned and device groups created with ethernet configuration.
// Each interface config must apply to the same set of ports.
func (ix *ixATE) addTopology(ifs []*opb.InterfaceConfig) error {
	// Add an empty topology to the config.
	var name string
	var link ixconfig.IxiaCfgNode
	var linkPorts []*ixconfig.Vport
	topo := &ixconfig.Topology{}
	switch v := ifs[0].GetLink().(type) {
	case *opb.InterfaceConfig_Port:
		var ok bool
		name = v.Port
		link, ok = ix.ports[name]
		if !ok {
			return fmt.Errorf("no ATE port configured with name %q", name)
		}
		linkPorts = []*ixconfig.Vport{link.(*ixconfig.Vport)}
		topo.SetVportsRefs([]ixconfig.IxiaCfgNode{link})
	case *opb.InterfaceConfig_Lag:
		var ok bool
		name = v.Lag
		link, ok = ix.lags[name]
		if !ok {
			return fmt.Errorf("no ATE LAG configured with name %q", name)
		}
		linkPorts = ix.lagPorts[link.(*ixconfig.Lag)]
		topo.SetPortsRefs([]ixconfig.IxiaCfgNode{link})
	}
	topo.Name = ixconfig.String(fmt.Sprintf("Topology on %s", name))

	for _, ifc := range ifs {
		// Configure device group for the interface, underneath a LAG if specified.
		dg := &ixconfig.TopologyDeviceGroup{Enabled: ixconfig.MultivalueTrue()}
		topo.DeviceGroup = append(topo.DeviceGroup, dg)
		mtu := uint32(1500)
		eth := ifc.GetEthernet()
		if eth.GetMtu() > mtu {
			mtu = eth.GetMtu()
		}
		if ifc.GetEnableLacp() {
			dg.Name = ixconfig.String(fmt.Sprintf("LACP DeviceGroup on %s", ifc.GetName()))
			dg.Multiplier = ixconfig.NumberUint32(1)
			dg.Ethernet = append(dg.Ethernet, &ixconfig.TopologyEthernet{
				Name: ixconfig.String(fmt.Sprintf("LACP Ethernet on %s", ifc.GetName())),
				Mtu:  ixconfig.MultivalueUint32(mtu),
				Lacp: []*ixconfig.TopologyLacp{{}},
			})
			// Add a child device group for configuring the rest of the topology.
			dg.DeviceGroup = append(dg.DeviceGroup, &ixconfig.TopologyDeviceGroup{Enabled: ixconfig.MultivalueTrue()})
			dg = dg.DeviceGroup[0]
		}
		dg.Name = ixconfig.String(fmt.Sprintf("Device Group on %s", ifc.GetName()))
		dg.Multiplier = ixconfig.NumberUint32(1)

		// Configure ethernet
		enableVlan := eth.GetVlanId() != 0
		topoEth := &ixconfig.TopologyEthernet{
			Name:        ixconfig.String(fmt.Sprintf("Ethernet on %s", name)),
			Mtu:         ixconfig.MultivalueUint32(mtu),
			EnableVlans: ixconfig.MultivalueBool(enableVlan),
			Mac:         &ixconfig.Multivalue{}, // Include to enable MAC resolution.
		}
		if enableVlan {
			topoEth.Vlan = []*ixconfig.TopologyVlan{{VlanId: ixconfig.MultivalueUint32(eth.GetVlanId())}}
		}

		if !eth.GetFec().GetEnabled() {
			for _, p := range linkPorts {
				// Turn on force disable FEC.
				aresOneFourHundredGigLan(p.L1Config).ForceDisableFEC = ixconfig.Bool(true)
				atlasFourHundredGigLan(p.L1Config).ForceDisableFEC = ixconfig.Bool(true)
				krakenFourHundredGigLan(p.L1Config).ForceDisableFEC = ixconfig.Bool(true)
				novusHundredGigLan(p.L1Config).ForceDisableFEC = ixconfig.Bool(true)
				uhdOneHundredGigLan(p.L1Config).ForceDisableFEC = ixconfig.Bool(true)
				// Turn off IEEE defaults.
				novusHundredGigLan(p.L1Config).IeeeL1Defaults = ixconfig.Bool(false)
				uhdOneHundredGigLan(p.L1Config).IeeeL1Defaults = ixconfig.Bool(false)
				// Turn off RS FEC.
				aresOneFourHundredGigLan(p.L1Config).EnableRsFec = ixconfig.Bool(false)
				atlasFourHundredGigLan(p.L1Config).EnableRsFec = ixconfig.Bool(false)
				krakenFourHundredGigLan(p.L1Config).EnableRsFec = ixconfig.Bool(false)
				novusHundredGigLan(p.L1Config).EnableRsFec = ixconfig.Bool(false)
				uhdOneHundredGigLan(p.L1Config).EnableRsFec = ixconfig.Bool(false)
				// Turn off RS FEC Stats.
				aresOneFourHundredGigLan(p.L1Config).EnableRsFecStats = ixconfig.Bool(false)
				atlasFourHundredGigLan(p.L1Config).EnableRsFecStats = ixconfig.Bool(false)
				krakenFourHundredGigLan(p.L1Config).EnableRsFecStats = ixconfig.Bool(false)
				novusHundredGigLan(p.L1Config).EnableRsFecStats = ixconfig.Bool(false)
			}
		}
		dg.Ethernet = append(dg.Ethernet, topoEth)
		ix.intfs[ifc.GetName()] = &intf{
			deviceGroup: dg,
			link:        link,
			hasVLAN:     enableVlan,
		}
	}
	ix.cfg.Topology = append(ix.cfg.Topology, topo)
	return nil
}

func aresOneFourHundredGigLan(l1 *ixconfig.VportL1Config) *ixconfig.VportL1ConfigAresOneFourHundredGigLan {
	if l1.AresOneFourHundredGigLan == nil {
		l1.AresOneFourHundredGigLan = &ixconfig.VportL1ConfigAresOneFourHundredGigLan{}
	}
	return l1.AresOneFourHundredGigLan
}

func atlasFourHundredGigLan(l1 *ixconfig.VportL1Config) *ixconfig.VportL1ConfigAtlasFourHundredGigLan {
	if l1.AtlasFourHundredGigLan == nil {
		l1.AtlasFourHundredGigLan = &ixconfig.VportL1ConfigAtlasFourHundredGigLan{}
	}
	return l1.AtlasFourHundredGigLan
}

func krakenFourHundredGigLan(l1 *ixconfig.VportL1Config) *ixconfig.VportL1ConfigKrakenFourHundredGigLan {
	if l1.KrakenFourHundredGigLan == nil {
		l1.KrakenFourHundredGigLan = &ixconfig.VportL1ConfigKrakenFourHundredGigLan{}
	}
	return l1.KrakenFourHundredGigLan
}

func novusHundredGigLan(l1 *ixconfig.VportL1Config) *ixconfig.VportL1ConfigNovusHundredGigLan {
	if l1.NovusHundredGigLan == nil {
		l1.NovusHundredGigLan = &ixconfig.VportL1ConfigNovusHundredGigLan{}
	}
	return l1.NovusHundredGigLan
}

func novusTenGigLan(l1 *ixconfig.VportL1Config) *ixconfig.VportL1ConfigNovusTenGigLan {
	if l1.NovusTenGigLan == nil {
		l1.NovusTenGigLan = &ixconfig.VportL1ConfigNovusTenGigLan{}
	}
	return l1.NovusTenGigLan
}

func uhdOneHundredGigLan(l1 *ixconfig.VportL1Config) *ixconfig.VportL1ConfigUhdOneHundredGigLan {
	if l1.UhdOneHundredGigLan == nil {
		l1.UhdOneHundredGigLan = &ixconfig.VportL1ConfigUhdOneHundredGigLan{}
	}
	return l1.UhdOneHundredGigLan
}
