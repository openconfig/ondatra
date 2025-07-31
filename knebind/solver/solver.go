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

	"golang.org/x/net/context"

	"bitbucket.org/creachadair/stringset"
	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/binding/introspect"
	"github.com/openconfig/ondatra/binding/portgraph"
	"github.com/pborman/uuid"

	tpb "github.com/openconfig/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

const (
	roleLabel = "ondatra-role"
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
		tpb.Vendor_ALPINE:     opb.Device_ALPINE,
	}
)

// ServiceName returns the name of the specified service in KNE.
func ServiceName(svc introspect.Service) (string, bool) {
	switch svc {
	case introspect.GNMI:
		return "gnmi", true
	case introspect.GNOI:
		return "gnoi", true
	case introspect.GNSI:
		return "gnsi", true
	case introspect.GRIBI:
		return "gribi", true
	case introspect.OTG:
		return "grpc", true
	case introspect.P4RT:
		return "p4rt", true
	}
	return "", false
}

func role(node *tpb.Node) string {
	if role, ok := node.GetLabels()[roleLabel]; ok {
		return role
	}
	if _, ok := ateVendors[node.GetVendor()]; ok {
		return portgraph.RoleATE
	}
	return portgraph.RoleDUT
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

var (
	reAny = portgraph.Regex(regexp.MustCompile(".*"))
)

// topoToConcreteGraph translates a Topology to a ConcreteGraph.
func topoToConcreteGraph(topo *tpb.Topology) (*portgraph.ConcreteGraph, map[*portgraph.ConcreteNode]*tpb.Node, map[*portgraph.ConcretePort]*intf, error) {
	const mtuAttr = "mtu"
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
			attrs[portgraph.VendorAttr] = vendor.String()
		}
		if r := role(node); r != "" {
			attrs[portgraph.RoleAttr] = r
		}
		if m := node.GetModel(); m != "" {
			attrs[portgraph.HWAttr] = m
		}
		if os := node.GetOs(); os != "" {
			attrs[portgraph.SWAttr] = os
		}
		if name := node.GetName(); name != "" {
			attrs[portgraph.NameAttr] = name
		}
		var ports []*portgraph.ConcretePort
		for intfName, nodeIntf := range node.GetInterfaces() {
			portAttrs := make(map[string]string)
			if mtu := nodeIntf.GetMtu(); mtu != 0 {
				portAttrs[mtuAttr] = strconv.FormatInt(int64(mtu), 10)
			}
			if group := nodeIntf.GetGroup(); group != "" {
				portAttrs[portgraph.GroupAttr] = group
			}
			switch {
			case nodeIntf.GetName() != "": // use the vendor name if specified
				portAttrs[portgraph.NameAttr] = nodeIntf.GetName()
			case intfName != "": // else use the interface name
				portAttrs[portgraph.NameAttr] = intfName
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

// createExactMap finds any exact mappings from testbed device ID to topology node names. If an
// exact mapping exists, i.e. there is a testbed device ID that equals a topology node name, then
// add it to the returned map. This mapping can be used to increase the efficiency of the solver.
func createExactMap(tb *opb.Testbed, topo *tpb.Topology) map[string]string {
	names := make(map[string]bool)
	for _, node := range topo.GetNodes() {
		names[node.GetName()] = true
	}
	m := make(map[string]string)
	for _, dut := range tb.GetDuts() {
		id := dut.GetId()
		// Ondatra testbed devices cannot contain "-" and KNE topology node names cannot contain "_" so
		// allow an exact match to include the substitution.
		name := strings.ReplaceAll(id, "_", "-")
		if names[name] {
			log.Infof("Found exact mapping for device %q -> %q", id, name)
			m[id] = name
		}
	}
	return m
}

// Solve creates a new Reservation from a desired testbed and an available topology.
func Solve(ctx context.Context, tb *opb.Testbed, topo *tpb.Topology, partial map[string]string) (*binding.Reservation, error) {
	if partial == nil {
		partial = createExactMap(tb, topo)
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

	abstractGraph, absNode2Dev, absPort2Port, err := portgraph.TestbedToAbstractGraph(tb, partial)
	if err != nil {
		return nil, fmt.Errorf("could not parse specified testbed: %w", err)
	}
	superGraph, conNode2Node, conPort2Intf, err := topoToConcreteGraph(topo)
	if err != nil {
		return nil, fmt.Errorf("could not parse specified testbed: %w", err)
	}

	assignment, err := portgraph.Solve(ctx, abstractGraph, superGraph)
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
		AbstractDUT: &binding.AbstractDUT{Dims: dr.dims},
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
		AbstractATE: &binding.AbstractATE{Dims: dr.dims},
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
		names := stringset.New()
		for _, n := range append(s.GetNames(), s.GetName()) {
			if n != "" {
				names.Add(n)
			}
		}
		for _, n := range names.Elements() {
			sm[n] = s
		}
	}
	dims := &binding.Dims{
		Name:            node.GetName(),
		Vendor:          vendor,
		HardwareModel:   node.GetModel(),
		SoftwareVersion: node.GetOs(),
		Ports:           make(map[string]*binding.Port),
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
