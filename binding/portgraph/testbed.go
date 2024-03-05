// Copyright 2024 Google LLC
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

package portgraph

import (
	"fmt"
	"regexp"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/internal/orderedmap"
	opb "github.com/openconfig/ondatra/proto"
)

const (
	// HWAttr is the name of the node constraint on the hardware model attribute.
	HWAttr = "hardware_model"
	// NameAttr is the name of the node constraint on the name attribute.
	NameAttr = "name"
	// RoleAttr is the name of the node constraint on the role attribute.
	RoleAttr = "role"
	// SWAttr is the name of the node constraint on the software version attribute.
	SWAttr = "software_version"
	// VendorAttr is the name of the node constraint on the vendor attribute.
	VendorAttr = "vendor"

	// SpeedAttr is the name of the port constraint on the speed attribute.
	SpeedAttr = "speed"
	// CardAttr is the name of the port constraint on the card model attribute.
	CardAttr = "card_model"
	// PMDAttr is the name of the port constraint on the pmd attribute.
	PMDAttr = "pmd"
	// GroupAttr is the name of the port constraint on the group attribute.
	GroupAttr = "group"

	// RoleDUT is the name of the value of the role attribute for DUTs.
	RoleDUT = "DUT"
	// RoleATE is the name of the value of the role attribute for ATEs.
	RoleATE = "ATE"
)

// TestbedToAbstractGraph translates an Ondatra testbed to an AbstractGraph for solving.
func TestbedToAbstractGraph(tb *opb.Testbed, partial map[string]string) (*AbstractGraph, map[*AbstractNode]*opb.Device, map[*AbstractPort]*opb.Port, error) {
	var nodes []*AbstractNode
	var edges []*AbstractEdge
	name2Port := make(map[string]*AbstractPort) // Name of the port to the AbstractPort.
	node2Dev := make(map[*AbstractNode]*opb.Device)
	port2Port := make(map[*AbstractPort]*opb.Port)

	// addDevice creates an AbstractNode from a Device.
	// If the field is empty, there is not a Constraint on that field.
	addDevice := func(dev *opb.Device, isATE bool) {
		nodeConstraints := make(map[string]NodeConstraint)
		if isATE {
			nodeConstraints[RoleAttr] = Equal(RoleATE)
		} else {
			nodeConstraints[RoleAttr] = Equal(RoleDUT)
		}
		if v := dev.GetVendor(); v != opb.Device_VENDOR_UNSPECIFIED {
			nodeConstraints[VendorAttr] = Equal(v.String())
		}
		if hw, ok := modelConstraint(dev); ok {
			nodeConstraints[HWAttr] = hw
		}
		if sw, ok := versionConstraint(dev); ok {
			nodeConstraints[SWAttr] = sw
		}
		if name, ok := partial[dev.GetId()]; ok {
			nodeConstraints[NameAttr] = Equal(name)
		}
		for k, v := range dev.GetExtraDimensions() {
			nodeConstraints[k] = Equal(v)
		}
		// Process each port on the device.
		var ports []*AbstractPort
		group2PortNames := orderedmap.NewOrderedMap[string, []string]()
		absPortName2DevPort := make(map[string]*opb.Port)
		for _, port := range dev.GetPorts() {
			portConstraints := make(map[string]PortConstraint)
			if s := port.GetSpeed(); s != opb.Port_SPEED_UNSPECIFIED {
				portConstraints[SpeedAttr] = Equal(s.String())
			}
			if cm, ok := cardModelConstraint(port); ok {
				portConstraints[CardAttr] = cm
			}
			if pmd, ok := pmdConstraint(port); ok {
				portConstraints[PMDAttr] = pmd
			}
			portName := fmt.Sprintf("%s:%s", dev.GetId(), port.GetId())
			if name, ok := partial[portName]; ok {
				portConstraints[NameAttr] = Equal(name)
			}
			// Create the AbstractPort, but do not add group constraints until all ports are processed.
			p := &AbstractPort{Desc: portName, Constraints: portConstraints}
			name2Port[portName] = p

			// Check if the port has a group. If it does, additional constraint(s) need to be added.
			if port.GetGroup() != "" {
				pns, _ := group2PortNames.Get(port.GetGroup())
				group2PortNames.Set(port.GetGroup(), append(pns, portName))
				absPortName2DevPort[portName] = port
			} else { // This port is not part of a group; no further processing.
				port2Port[p] = port
				ports = append(ports, p)
			}
		}
		processedGroups := map[string]struct{}{}
		groups := group2PortNames.Keys()
		for _, group := range groups {
			portNames, _ := group2PortNames.Get(group)
			processedGroups[group] = struct{}{}
			var constraints []LeafPortConstraint
			pn := portNames[0]
			for _, pn2 := range portNames[1:] {
				p2 := name2Port[pn2]
				constraints = append(constraints, SameAsPort(p2))
				name2Port[pn2].Constraints[GroupAttr] = reAny
				port2Port[p2] = absPortName2DevPort[pn2]
				ports = append(ports, p2)
			}
			for _, g2 := range groups {
				if _, ok := processedGroups[g2]; !ok {
					p, _ := group2PortNames.Get(g2)
					constraints = append(constraints, NotSameAsPort(name2Port[p[0]]))
				} else if len(constraints) == 0 {
					constraints = append(constraints, reAny)
				}
			}
			p := name2Port[pn]
			if len(constraints) == 1 {
				p.Constraints[GroupAttr] = constraints[0]
			} else {
				p.Constraints[GroupAttr] = AndPort(constraints...)
			}
			name2Port[pn] = p
			ports = append(ports, p)
			port2Port[p] = absPortName2DevPort[pn]
		}
		n := &AbstractNode{Desc: dev.GetId(), Constraints: nodeConstraints, Ports: ports}
		node2Dev[n] = dev
		nodes = append(nodes, n)
	}

	for _, dev := range tb.GetDuts() {
		addDevice(dev, false)
	}
	for _, dev := range tb.GetAtes() {
		addDevice(dev, true)
	}
	for _, link := range tb.GetLinks() {
		src, ok := name2Port[link.GetA()]
		if !ok {
			return nil, nil, nil, fmt.Errorf("port %q in links not present in specified testbed", link.GetA())
		}
		dst, ok := name2Port[link.GetB()]
		if !ok {
			return nil, nil, nil, fmt.Errorf("port %q in links not present in specified testbed", link.GetB())
		}
		edges = append(edges, &AbstractEdge{Src: src, Dst: dst})
	}
	return &AbstractGraph{Desc: "KNE testbed", Nodes: nodes, Edges: edges}, node2Dev, port2Port, nil
}

func modelConstraint(d *opb.Device) (NodeConstraint, bool) {
	switch v := d.GetHardwareModelValue().(type) {
	case nil:
		// handled after switch
	case *opb.Device_HardwareModel:
		if v.HardwareModel != "" {
			return Equal(v.HardwareModel), true
		}
	case *opb.Device_HardwareModelRegex:
		if v.HardwareModelRegex != "" {
			return Regex(regexp.MustCompile(v.HardwareModelRegex)), true
		}
	default:
		log.Fatalf("unknown hardware model type: %v(%T)", v, v)
	}
	return nil, false
}

func versionConstraint(d *opb.Device) (NodeConstraint, bool) {
	switch v := d.GetSoftwareVersionValue().(type) {
	case nil:
		// handled after switch
	case *opb.Device_SoftwareVersion:
		if v.SoftwareVersion != "" {
			return Equal(v.SoftwareVersion), true
		}
	case *opb.Device_SoftwareVersionRegex:
		if v.SoftwareVersionRegex != "" {
			return Regex(regexp.MustCompile(v.SoftwareVersionRegex)), true
		}
	default:
		log.Fatalf("unknown software version type: %v(%T)", v, v)
	}
	return nil, false
}

func cardModelConstraint(p *opb.Port) (PortConstraint, bool) {
	switch v := p.GetCardModelValue().(type) {
	case nil:
		// handled after switch
	case *opb.Port_CardModel:
		if v.CardModel != "" {
			return Equal(v.CardModel), true
		}
	case *opb.Port_CardModelRegex:
		if v.CardModelRegex != "" {
			return Regex(regexp.MustCompile(v.CardModelRegex)), true
		}
	default:
		log.Fatalf("unknown card model type: %v(%T)", v, v)
	}
	return nil, false
}

func pmdConstraint(p *opb.Port) (PortConstraint, bool) {
	switch v := p.GetPmdValue().(type) {
	case nil:
		// handled after switch
	case *opb.Port_Pmd_:
		if v.Pmd != opb.Port_PMD_UNSPECIFIED {
			return Equal(v.Pmd.String()), true
		}
	case *opb.Port_PmdRegex:
		if v.PmdRegex != "" {
			return Regex(regexp.MustCompile(v.PmdRegex)), true
		}
	default:
		log.Fatalf("unknown PMD type: %v(%T)", v, v)
	}
	return nil, false
}
