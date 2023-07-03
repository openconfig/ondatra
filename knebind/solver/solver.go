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

// Package solver creates solutions from devices and topologies.
package solver

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/binding/portgraph"
	"github.com/openconfig/ondatra/internal/orderedmap"
	"github.com/pborman/uuid"
	"google.golang.org/protobuf/proto"

	tpb "github.com/openconfig/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

const (
	// KNEServiceMapKey is the key to look up the service map in the custom data of a ServiceDUT.
	KNEServiceMapKey = "$KEY_SERVICE_MAP"

	roleLabel = "ondatra-role"
	roleDUT   = "DUT"
	roleATE   = "ATE"
)

var (
	ateVendors = map[tpb.Vendor]bool{
		tpb.Vendor_KEYSIGHT: true,
	}

	deviceVendors = map[tpb.Vendor]opb.Device_Vendor{
		tpb.Vendor_ARISTA:     opb.Device_ARISTA,
		tpb.Vendor_CISCO:      opb.Device_CISCO,
		tpb.Vendor_KEYSIGHT:   opb.Device_IXIA,
		tpb.Vendor_JUNIPER:    opb.Device_JUNIPER,
		tpb.Vendor_NOKIA:      opb.Device_NOKIA,
		tpb.Vendor_OPENCONFIG: opb.Device_OPENCONFIG,
		tpb.Vendor_HOST:       opb.Device_VENDOR_UNSPECIFIED,
	}
)

func role(node *tpb.Node) string {
	if role, ok := node.GetLabels()[roleLabel]; ok {
		return role
	}
	if _, ok := ateVendors[node.GetVendor()]; ok {
		return roleATE
	}
	return roleDUT
}

func filterTopology(topo *tpb.Topology) *tpb.Topology {
	t := &tpb.Topology{Name: topo.GetName()}
	for _, node := range topo.GetNodes() {
		// Only include nodes with known vendor.
		if _, ok := deviceVendors[node.GetVendor()]; ok {
			t.Nodes = append(t.Nodes, node)
		} else {
			log.Infof("No known device vendor for node %q (vendor %v), ignoring node", node.GetName(), node.GetVendor())
		}
	}
	for _, link := range topo.GetLinks() {
		foundA := false
		foundZ := false
		for _, node := range t.GetNodes() {
			if link.GetANode() == node.GetName() {
				foundA = true
			}
			if link.GetZNode() == node.GetName() {
				foundZ = true
			}
		}
		// Only include links with nodes passing filter.
		if foundA && foundZ {
			t.Links = append(t.Links, link)
		}
	}
	return t
}

const (
	// Attribute names mapping message fields to graph attributes/constraints.
	vendorAttr = "vendor"
	roleAttr   = "role"
	hwAttr     = "hardware_model"
	swAttr     = "software_version"
	speedAttr  = "speed"
	cardAttr   = "card_model"
	pmdAttr    = "pmd"
	groupAttr  = "group"
	mtuAttr    = "mtu"
	nameAttr   = "name"
)

var (
	reAny = portgraph.Regex(regexp.MustCompile(".*"))
)

// testbedToAbstractGraph translates an Ondatra testbed to an AbstractGraph for solving.
func testbedToAbstractGraph(tb *opb.Testbed, partial map[string]string) (*portgraph.AbstractGraph, map[*portgraph.AbstractNode]*opb.Device, map[*portgraph.AbstractPort]*opb.Port, error) {
	var nodes []*portgraph.AbstractNode
	var edges []*portgraph.AbstractEdge
	name2Port := make(map[string]*portgraph.AbstractPort) // Name of the port to the AbstractPort.
	node2Dev := make(map[*portgraph.AbstractNode]*opb.Device)
	port2Port := make(map[*portgraph.AbstractPort]*opb.Port)

	// addDevice creates an AbstractNode from a Device.
	// If the field is empty, there is not a Constraint on that field.
	addDevice := func(dev *opb.Device, isATE bool) {
		nodeConstraints := make(map[string]portgraph.NodeConstraint)
		if isATE {
			nodeConstraints[roleAttr] = portgraph.Equal(roleATE)
		} else {
			nodeConstraints[roleAttr] = portgraph.Equal(roleDUT)
		}
		if v := dev.GetVendor(); v != opb.Device_VENDOR_UNSPECIFIED {
			nodeConstraints[vendorAttr] = portgraph.Equal(v.String())
		}
		if hw, ok := modelConstraint(dev); ok {
			nodeConstraints[hwAttr] = hw
		}
		if sw, ok := versionConstraint(dev); ok {
			nodeConstraints[swAttr] = sw
		}
		if name, ok := partial[dev.GetId()]; ok {
			nodeConstraints[nameAttr] = portgraph.Equal(name)
		}
		for k, v := range dev.GetExtraDimensions() {
			nodeConstraints[k] = portgraph.Equal(v)
		}
		// Process each port on the device.
		var ports []*portgraph.AbstractPort
		group2PortNames := orderedmap.NewOrderedMap[string, []string]()
		absPortName2DevPort := make(map[string]*opb.Port)
		for _, port := range dev.GetPorts() {
			portConstraints := make(map[string]portgraph.PortConstraint)
			if s := port.GetSpeed(); s != opb.Port_SPEED_UNSPECIFIED {
				portConstraints[speedAttr] = portgraph.Equal(s.String())
			}
			if cm, ok := cardModelConstraint(port); ok {
				portConstraints[cardAttr] = cm
			}
			if pmd, ok := pmdConstraint(port); ok {
				portConstraints[pmdAttr] = pmd
			}
			portName := fmt.Sprintf("%s:%s", dev.GetId(), port.GetId())
			if name, ok := partial[portName]; ok {
				portConstraints[nameAttr] = portgraph.Equal(name)
			}
			// Create the AbstractPort, but do not add group constraints until all ports are processed.
			p := &portgraph.AbstractPort{Desc: portName, Constraints: portConstraints}
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
			var constraints []portgraph.LeafPortConstraint
			pn := portNames[0]
			for _, pn2 := range portNames[1:] {
				p2 := name2Port[pn2]
				constraints = append(constraints, portgraph.SameAsPort(p2))
				name2Port[pn2].Constraints[groupAttr] = reAny
				port2Port[p2] = absPortName2DevPort[pn2]
				ports = append(ports, p2)
			}
			for _, g2 := range groups {
				if _, ok := processedGroups[g2]; !ok {
					p, _ := group2PortNames.Get(g2)
					constraints = append(constraints, portgraph.NotSameAsPort(name2Port[p[0]]))
				} else if len(constraints) == 0 {
					constraints = append(constraints, reAny)
				}
			}
			p := name2Port[pn]
			if len(constraints) == 1 {
				p.Constraints[groupAttr] = constraints[0]
			} else {
				p.Constraints[groupAttr] = portgraph.AndPort(constraints...)
			}
			name2Port[pn] = p
			ports = append(ports, p)
			port2Port[p] = absPortName2DevPort[pn]
		}
		n := &portgraph.AbstractNode{Desc: dev.GetId(), Constraints: nodeConstraints, Ports: ports}
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
		edges = append(edges, &portgraph.AbstractEdge{Src: src, Dst: dst})
	}
	return &portgraph.AbstractGraph{Desc: "KNE testbed", Nodes: nodes, Edges: edges}, node2Dev, port2Port, nil
}

func modelConstraint(d *opb.Device) (portgraph.NodeConstraint, bool) {
	switch v := d.GetHardwareModelValue().(type) {
	case nil:
		// handled after switch
	case *opb.Device_HardwareModel:
		if v.HardwareModel != "" {
			return portgraph.Equal(v.HardwareModel), true
		}
	case *opb.Device_HardwareModelRegex:
		if v.HardwareModelRegex != "" {
			return portgraph.Regex(regexp.MustCompile(v.HardwareModelRegex)), true
		}
	default:
		log.Fatalf("unknown hardware model type: %v(%T)", v, v)
	}
	return nil, false
}

func versionConstraint(d *opb.Device) (portgraph.NodeConstraint, bool) {
	switch v := d.GetSoftwareVersionValue().(type) {
	case nil:
		// handled after switch
	case *opb.Device_SoftwareVersion:
		if v.SoftwareVersion != "" {
			return portgraph.Equal(v.SoftwareVersion), true
		}
	case *opb.Device_SoftwareVersionRegex:
		if v.SoftwareVersionRegex != "" {
			return portgraph.Regex(regexp.MustCompile(v.SoftwareVersionRegex)), true
		}
	default:
		log.Fatalf("unknown software version type: %v(%T)", v, v)
	}
	return nil, false
}

func cardModelConstraint(p *opb.Port) (portgraph.PortConstraint, bool) {
	switch v := p.GetCardModelValue().(type) {
	case nil:
		// handled after switch
	case *opb.Port_CardModel:
		if v.CardModel != "" {
			return portgraph.Equal(v.CardModel), true
		}
	case *opb.Port_CardModelRegex:
		if v.CardModelRegex != "" {
			return portgraph.Regex(regexp.MustCompile(v.CardModelRegex)), true
		}
	default:
		log.Fatalf("unknown card model type: %v(%T)", v, v)
	}
	return nil, false
}

func pmdConstraint(p *opb.Port) (portgraph.PortConstraint, bool) {
	switch v := p.GetPmdValue().(type) {
	case nil:
		// handled after switch
	case *opb.Port_Pmd_:
		if v.Pmd != opb.Port_PMD_UNSPECIFIED {
			return portgraph.Equal(v.Pmd.String()), true
		}
	case *opb.Port_PmdRegex:
		if v.PmdRegex != "" {
			return portgraph.Regex(regexp.MustCompile(v.PmdRegex)), true
		}
	default:
		log.Fatalf("unknown PMD type: %v(%T)", v, v)
	}
	return nil, false
}

// topoToConcreteGraph translates a Topology to a ConcreteGraph.
func topoToConcreteGraph(topo *tpb.Topology) (*portgraph.ConcreteGraph, map[*portgraph.ConcreteNode]*tpb.Node, map[*portgraph.ConcretePort]*intf, error) {
	var nodes []*portgraph.ConcreteNode
	var edges []*portgraph.ConcreteEdge
	nodeName2Node := make(map[string]*portgraph.ConcreteNode)
	portName2Port := make(map[string]*portgraph.ConcretePort)
	node2Node := make(map[*portgraph.ConcreteNode]*tpb.Node)
	port2Intf := make(map[*portgraph.ConcretePort]*intf)

	makePortAndIntf := func(nodeName, intfName, vendorFromNode string, attrs map[string]string) (*portgraph.ConcretePort, *intf) {
		pn := fmt.Sprintf("%s:%s", nodeName, intfName)
		p := &portgraph.ConcretePort{Desc: pn, Attrs: attrs}
		i := &intf{name: intfName}
		if vendorFromNode != "" {
			i.vendorName = vendorFromNode
		} else {
			i.vendorName = intfName
		}
		return p, i
	}
	for _, node := range topo.GetNodes() {
		attrs := make(map[string]string)
		if vendor, ok := deviceVendors[node.GetVendor()]; ok {
			attrs[vendorAttr] = vendor.String()
		}
		if r := role(node); r != "" {
			attrs[roleAttr] = r
		}
		if m := node.GetModel(); m != "" {
			attrs[hwAttr] = m
		}
		if os := node.GetOs(); os != "" {
			attrs[swAttr] = os
		}
		if name := node.GetName(); name != "" {
			attrs[nameAttr] = name
		}
		var ports []*portgraph.ConcretePort
		for intfName, nodeIntf := range node.GetInterfaces() {
			portAttrs := make(map[string]string)
			if mtu := nodeIntf.GetMtu(); mtu != 0 {
				portAttrs[mtuAttr] = strconv.FormatInt(int64(mtu), 10)
			}
			if group := nodeIntf.GetGroup(); group != "" {
				portAttrs[groupAttr] = group
			}
			switch {
			case nodeIntf.GetName() != "": // use the vendor name if specified
				portAttrs[nameAttr] = nodeIntf.GetName()
			case intfName != "": // else use the interface name
				portAttrs[nameAttr] = intfName
			}
			p, i := makePortAndIntf(node.GetName(), intfName, nodeIntf.GetName(), portAttrs)
			port2Intf[p] = i
			portName2Port[p.Desc] = p
			ports = append(ports, p)
		}
		n := &portgraph.ConcreteNode{Desc: node.GetName(), Attrs: attrs, Ports: ports}
		node2Node[n] = node
		nodeName2Node[n.Desc] = n
	}
	createAndAddPort := func(srcNode, srcPort, intfName string) *portgraph.ConcretePort {
		p := &portgraph.ConcretePort{Desc: srcPort}
		i := &intf{name: intfName, vendorName: intfName}
		port2Intf[p] = i
		portName2Port[srcPort] = p
		nodeName2Node[srcNode].Ports = append(nodeName2Node[srcNode].Ports, p)
		return p
	}
	for _, link := range topo.GetLinks() {
		srcName := fmt.Sprintf("%s:%s", link.GetANode(), link.GetAInt())
		dstName := fmt.Sprintf("%s:%s", link.GetZNode(), link.GetZInt())
		src, ok := portName2Port[srcName]
		if !ok {
			src = createAndAddPort(link.GetANode(), srcName, link.GetAInt())
		}
		dst, ok := portName2Port[dstName]
		if !ok {
			dst = createAndAddPort(link.GetZNode(), dstName, link.GetZInt())
		}
		edges = append(edges, &portgraph.ConcreteEdge{Src: src, Dst: dst})
	}
	for _, node := range nodeName2Node {
		nodes = append(nodes, node)
	}
	return &portgraph.ConcreteGraph{Desc: topo.GetName(), Nodes: nodes, Edges: edges}, node2Node, port2Intf, nil
}

// assignmentToReservation maps the portgraph.Assignment from portgraph.Solve to a reservation proto.
func assignmentToReservation(
	assignment *portgraph.Assignment,
	duts, ates []*opb.Device,
	absNode2Dev map[*portgraph.AbstractNode]*opb.Device,
	absPort2Port map[*portgraph.AbstractPort]*opb.Port,
	conNode2Node map[*portgraph.ConcreteNode]*tpb.Node,
	conNode2Intf map[*portgraph.ConcretePort]*intf) (*binding.Reservation, error) {
	dev2Node := make(map[*opb.Device]*tpb.Node)
	port2Intf := make(map[*opb.Port]*intf)
	for absNode, conNode := range assignment.Node2Node {
		dev2Node[absNode2Dev[absNode]] = conNode2Node[conNode]
	}
	for absPort, conPort := range assignment.Port2Port {
		port2Intf[absPort2Port[absPort]] = conNode2Intf[conPort]
	}

	a := &assign{dev2Node, port2Intf}
	res := &binding.Reservation{
		ID:   uuid.New(),
		DUTs: make(map[string]binding.DUT),
		ATEs: make(map[string]binding.ATE),
	}
	for _, dut := range duts {
		resDUT, err := a.resolveDUT(dut)
		if err != nil {
			return nil, err
		}
		res.DUTs[dut.GetId()] = resDUT
	}
	for _, ate := range ates {
		resATE, err := a.resolveATE(ate)
		if err != nil {
			return nil, err
		}
		res.ATEs[ate.GetId()] = resATE
	}

	return res, nil
}

// runtimeProtoCheck checks the proto messages have the expected number of fields.
func runtimeProtoCheck() error {
	const (
		testbedDeviceFields = 8
		testbedPortFields   = 7
		topoNodeFields      = 11
		topoIntfFields      = 7
	)
	numFields := func(m proto.Message) int {
		return m.ProtoReflect().Descriptor().Fields().Len()
	}
	if n := numFields((*opb.Device)(nil)); n != testbedDeviceFields {
		return fmt.Errorf("Ondatra testbed proto Device has %d fields, want %d", n, testbedDeviceFields)
	}
	if n := numFields((*opb.Port)(nil)); n != testbedPortFields {
		return fmt.Errorf("Ondatra testbed proto Port has %d fields, want %d", n, testbedPortFields)
	}
	if n := numFields((*tpb.Node)(nil)); n != topoNodeFields {
		return fmt.Errorf("Ondatra topology proto Node has %d fields, want %d", n, topoNodeFields)
	}
	if n := numFields((*tpb.Interface)(nil)); n != topoIntfFields {
		return fmt.Errorf("Ondatra topology proto Interface has %d fields, want %d", n, topoIntfFields)
	}
	return nil
}

// Solve creates a new Reservation from a desired testbed and an available topology.
func Solve(tb *opb.Testbed, topo *tpb.Topology, partial map[string]string) (*binding.Reservation, error) {
	if err := runtimeProtoCheck(); err != nil {
		return nil, err
	}
	topo = filterTopology(topo)
	devs := append(append([]*opb.Device{}, tb.GetDuts()...), tb.GetAtes()...)
	if numDevs, numNodes := len(devs), len(topo.GetNodes()); numDevs > numNodes {
		return nil, fmt.Errorf("not enough nodes in KNE topology for specified testbed: "+
			" testbed has %d devices and topology only has %d nodes", numDevs, numNodes)
	}
	if numTBLinks, numTopoLinks := len(tb.GetLinks()), len(topo.GetLinks()); numTBLinks > numTopoLinks {
		return nil, fmt.Errorf("not enough links in KNE topology for specified testbed "+
			" testbed has %d links and topology only has %d links", numTBLinks, numTopoLinks)
	}

	abstractGraph, absNode2Dev, absPort2Port, err := testbedToAbstractGraph(tb, partial)
	if err != nil {
		return nil, fmt.Errorf("could not parse specified testbed: %w", err)
	}
	superGraph, conNode2Node, conPort2Intf, err := topoToConcreteGraph(topo)
	if err != nil {
		return nil, fmt.Errorf("could not parse specified testbed: %w", err)
	}

	assignment, err := portgraph.Solve(abstractGraph, superGraph)
	if err != nil {
		return nil, fmt.Errorf("could not solve for specified testbed: %w", err)
	}

	res, err := assignmentToReservation(assignment, tb.GetDuts(), tb.GetAtes(), absNode2Dev, absPort2Port, conNode2Node, conPort2Intf)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ServiceDUT is a DUT that contains a service map.
type ServiceDUT struct {
	*binding.AbstractDUT
	Services   map[string]*tpb.Service
	Cert       *tpb.CertificateCfg
	NodeVendor tpb.Vendor
}

// Service returns the KNE service details for a given service name.
func (d *ServiceDUT) Service(service string) (*tpb.Service, error) {
	srv, ok := d.Services[service]
	if !ok {
		return nil, fmt.Errorf("service %q not found on DUT %q", service, d.Name())
	}
	return srv, nil
}

// ServiceATE is an ATE that contains a service map.
type ServiceATE struct {
	*binding.AbstractATE
	Services   map[string]*tpb.Service
	Cert       *tpb.CertificateCfg
	NodeVendor tpb.Vendor
}

// Service returns the KNE service details for a given service name.
func (a *ServiceATE) Service(service string) (*tpb.Service, error) {
	srv, ok := a.Services[service]
	if !ok {
		return nil, fmt.Errorf("service %q not found on ATE %q", service, a.Name())
	}
	return srv, nil
}

type assign struct {
	dev2Node  map[*opb.Device]*tpb.Node
	port2Intf map[*opb.Port]*intf
}

type intf struct {
	name       string
	vendorName string
}

func (a *assign) String() string {
	var sb strings.Builder
	for dut, node := range a.dev2Node {
		sb.WriteString(fmt.Sprintf("%s->%s\n", dut.GetId(), node.GetName()))
		for _, p := range dut.GetPorts() {
			sb.WriteString(fmt.Sprintf("  %s->%s\n", p.GetId(), a.port2Intf[p]))
		}
	}
	return sb.String()
}

func (a *assign) resolveDUT(dev *opb.Device) (*ServiceDUT, error) {
	dr, err := a.resolveDevice(dev)
	if err != nil {
		return nil, err
	}
	return &ServiceDUT{
		AbstractDUT: &binding.AbstractDUT{dr.dims},
		Services:    dr.services,
		Cert:        dr.cert,
		NodeVendor:  dr.nodeVendor,
	}, nil
}

func (a *assign) resolveATE(dev *opb.Device) (*ServiceATE, error) {
	dr, err := a.resolveDevice(dev)
	if err != nil {
		return nil, err
	}
	return &ServiceATE{
		AbstractATE: &binding.AbstractATE{dr.dims},
		Services:    dr.services,
		Cert:        dr.cert,
		NodeVendor:  dr.nodeVendor,
	}, nil
}

type deviceResolution struct {
	dims       *binding.Dims
	services   map[string]*tpb.Service
	cert       *tpb.CertificateCfg
	nodeVendor tpb.Vendor
}

func (a *assign) resolveDevice(dev *opb.Device) (*deviceResolution, error) {
	node, ok := a.dev2Node[dev]
	if !ok {
		return nil, fmt.Errorf("node %q not resolved", dev.GetId())
	}
	vendor, ok := deviceVendors[node.GetVendor()]
	if !ok {
		return nil, fmt.Errorf("no known device vendor for node %q (vendor %v)", node.GetName(), node.GetVendor())
	}
	sm := map[string]*tpb.Service{}
	for _, s := range node.GetServices() {
		sm[s.GetName()] = s
	}
	dims := &binding.Dims{
		Name:            node.GetName(),
		Vendor:          vendor,
		HardwareModel:   node.GetModel(),
		SoftwareVersion: node.GetOs(),
		Ports:           make(map[string]*binding.Port),
		CustomData:      map[string]any{KNEServiceMapKey: sm},
	}
	for _, p := range dev.GetPorts() {
		dims.Ports[p.GetId()] = &binding.Port{Name: a.port2Intf[p].vendorName}
	}
	return &deviceResolution{
		dims:       dims,
		services:   sm,
		cert:       node.GetConfig().GetCert(),
		nodeVendor: node.GetVendor(),
	}, nil
}
