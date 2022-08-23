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
	"strings"

	log "github.com/golang/glog"
	"github.com/pborman/uuid"
	"github.com/openconfig/ondatra/binding"

	tpb "github.com/openconfig/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

var (
	ateTypes = map[tpb.Node_Type]bool{
		tpb.Node_IXIA_TG: true,
	}

	// type2VendorMap maps the KNE node type to the Ondatra vendor.
	type2VendorMap = map[tpb.Node_Type]opb.Device_Vendor{
		tpb.Node_ARISTA_CEOS: opb.Device_ARISTA,
		// TODO: when Ondatra supports the OS dimension, use it to
		// distinguish CSR from CXR and CEVO from VMX.
		tpb.Node_CISCO_CSR:    opb.Device_CISCO,
		tpb.Node_CISCO_CXR:    opb.Device_CISCO,
		tpb.Node_CISCO_E8000:  opb.Device_CISCO,
		tpb.Node_CISCO_XRD:    opb.Device_CISCO,
		tpb.Node_IXIA_TG:      opb.Device_IXIA,
		tpb.Node_JUNIPER_CEVO: opb.Device_JUNIPER,
		tpb.Node_JUNIPER_VMX:  opb.Device_JUNIPER,
		tpb.Node_NOKIA_SRL:    opb.Device_NOKIA,
	}
)

// Solve creates a new Reservation from a desired testbed and an available topology.
func Solve(tb *opb.Testbed, topo *tpb.Topology) (*binding.Reservation, error) {
	devs := append(append([]*opb.Device{}, tb.GetDuts()...), tb.GetAtes()...)
	if numDevs, numNodes := len(devs), len(topo.GetNodes()); numDevs > numNodes {
		return nil, fmt.Errorf("not enough nodes in KNE topology for specified testbed: "+
			" testbed has %d devices and topology only has %d nodes", numDevs, numNodes)
	}
	if numTBLinks, numTopoLinks := len(tb.GetLinks()), len(topo.GetLinks()); numTBLinks > numTopoLinks {
		return nil, fmt.Errorf("not enough links in KNE topology for specified testbed "+
			" testbed has %d links and topology only has %d links", numTBLinks, numTopoLinks)
	}
	s := &solver{
		testbed:    tb,
		topology:   topo,
		id2Dev:     make(map[string]*opb.Device),
		dev2Ports:  make(map[*opb.Device]map[string]*opb.Port),
		dev2Links:  make(map[*opb.Device]map[*opb.Device]*[]*link),
		name2Node:  make(map[string]*tpb.Node),
		node2Intfs: make(map[*tpb.Node]map[string]*intf),
		node2Links: make(map[*tpb.Node]map[*tpb.Node]*[]*link),
	}

	// Cache various info in the solver about the testbed and topology.
	for _, dev := range devs {
		s.id2Dev[dev.GetId()] = dev
		ports := make(map[string]*opb.Port)
		for _, port := range dev.GetPorts() {
			ports[port.GetId()] = port
		}
		s.dev2Ports[dev] = ports
		s.dev2Links[dev] = make(map[*opb.Device]*[]*link)
	}
	// Build up the dev => links map for TB
	for _, l := range tb.GetLinks() {
		devParts := strings.Split(l.GetA(), ":")
		dev := s.id2Dev[devParts[0]]
		peerParts := strings.Split(l.GetB(), ":")
		peer := s.id2Dev[peerParts[0]]
		linkPtr := &link{name: devParts[0], port: devParts[1], peerName: peerParts[0], peerPort: peerParts[1]}
		if _, ok := s.dev2Links[dev][peer]; !ok {
			linkList := []*link{}
			s.dev2Links[dev][peer] = &linkList
			s.dev2Links[peer][dev] = &linkList
		}
		*s.dev2Links[dev][peer] = append(*s.dev2Links[dev][peer], linkPtr)
	}

	for _, node := range s.topology.GetNodes() {
		s.name2Node[node.GetName()] = node
		s.node2Intfs[node] = make(map[string]*intf)
		s.node2Links[node] = make(map[*tpb.Node]*[]*link)
	}
	getIntf := func(nodeName, intfName string) *intf {
		node := s.name2Node[nodeName]
		i, ok := s.node2Intfs[node][intfName]
		if !ok {
			i = &intf{name: intfName}
			s.node2Intfs[node][intfName] = i
		}
		i.vendorName = node.Interfaces[intfName].GetName()
		if i.vendorName == "" {
			i.vendorName = intfName
		}
		return i
	}
	for _, link := range s.topology.GetLinks() {
		getIntf(link.GetANode(), link.GetAInt())
		getIntf(link.GetZNode(), link.GetZInt())
	}
	// Build up the node => links map for topology
	for _, l := range topo.GetLinks() {
		node := s.name2Node[l.GetANode()]
		peer := s.name2Node[l.GetZNode()]
		linkPtr := &link{name: l.GetANode(), port: l.GetAInt(), peerName: l.GetZNode(), peerPort: l.GetZInt()}
		if _, ok := s.node2Links[node][peer]; !ok {
			linkList := []*link{}
			s.node2Links[node][peer] = &linkList
			s.node2Links[peer][node] = &linkList
		}
		*s.node2Links[node][peer] = append(*s.node2Links[node][peer], linkPtr)
	}

	a, err := s.solve()
	if err != nil {
		return nil, err
	}

	res := &binding.Reservation{
		ID:   uuid.New(),
		DUTs: make(map[string]binding.DUT),
		ATEs: make(map[string]binding.ATE),
	}

	for _, dut := range tb.GetDuts() {
		resDUT, err := a.resolveDUT(dut)
		if err != nil {
			return nil, err
		}
		res.DUTs[dut.GetId()] = resDUT
	}
	for _, ate := range tb.GetAtes() {
		resATE, err := a.resolveATE(ate)
		if err != nil {
			return nil, err
		}
		res.ATEs[ate.GetId()] = resATE
	}
	return res, nil
}

// ServiceDUT is a DUT that contains a service map.
type ServiceDUT struct {
	*binding.AbstractDUT
	Services map[string]*tpb.Service
	Cert     *tpb.CertificateCfg
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
	Services map[string]*tpb.Service
	Cert     *tpb.CertificateCfg
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

type link struct {
	name     string
	port     string
	peerName string
	peerPort string
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
	dims, srvs, cert, err := a.resolveDevice(dev)
	if err != nil {
		return nil, err
	}
	return &ServiceDUT{
		AbstractDUT: &binding.AbstractDUT{dims},
		Services:    srvs,
		Cert:        cert,
	}, nil
}

func (a *assign) resolveATE(dev *opb.Device) (*ServiceATE, error) {
	dims, srvs, cert, err := a.resolveDevice(dev)
	if err != nil {
		return nil, err
	}
	return &ServiceATE{
		AbstractATE: &binding.AbstractATE{dims},
		Services:    srvs,
		Cert:        cert,
	}, nil
}

func (a *assign) resolveDevice(dev *opb.Device) (*binding.Dims, map[string]*tpb.Service, *tpb.CertificateCfg, error) {
	node, ok := a.dev2Node[dev]
	if !ok {
		return nil, nil, nil, fmt.Errorf("node %q not resolved", dev.GetId())
	}
	vendor, ok := type2VendorMap[node.GetType()]
	if !ok {
		return nil, nil, nil, fmt.Errorf("no known device vendor for node type: %v", node.GetType())
	}
	typeName := tpb.Node_Type_name[int32(node.GetType())]
	dims := &binding.Dims{
		Name:   node.GetName(),
		Vendor: vendor,
		// TODO: Determine appropriate hardware model and software version
		HardwareModel:   typeName,
		SoftwareVersion: typeName,
		Ports:           make(map[string]*binding.Port),
	}
	for _, p := range dev.GetPorts() {
		dims.Ports[p.GetId()] = &binding.Port{Name: a.port2Intf[p].vendorName}
	}
	sm := map[string]*tpb.Service{}
	for _, s := range node.GetServices() {
		sm[s.GetName()] = s
	}
	return dims, sm, node.GetConfig().GetCert(), nil
}

type solver struct {
	testbed    *opb.Testbed
	topology   *tpb.Topology
	id2Dev     map[string]*opb.Device
	dev2Ports  map[*opb.Device]map[string]*opb.Port
	dev2Links  map[*opb.Device]map[*opb.Device]*[]*link
	name2Node  map[string]*tpb.Node
	node2Intfs map[*tpb.Node]map[string]*intf
	node2Links map[*tpb.Node]map[*tpb.Node]*[]*link
}

// arrangeDevKeys creates a list out devices from TB. The devices are added in the order they
// appear linked to the already added mesh. All isolated devices are added at the end.
// m - is the map of all TB devices and possible node mapping
// keys - updates this list of devices in the desired order for selection search
func (s *solver) arrangeDevKeys(m map[*opb.Device][]*tpb.Node, keys *[]*opb.Device) {
	// Get the entry with max number of peers first
	processed := make(map[*opb.Device]bool)
	maxPeer := -1
	var devPtr *opb.Device
	for k := range m {
		if len(s.dev2Links[k]) > maxPeer {
			maxPeer = len(s.dev2Links[k])
			devPtr = k
		}
	}
	index := 0
	*keys = append(*keys, devPtr)
	processed[devPtr] = true
	for index < len(*keys) {
		dev := (*keys)[index]
		for peer, _ := range s.dev2Links[dev] {
			if _, found := processed[peer]; !found {
				*keys = append(*keys, peer)
				processed[peer] = true
			}
		}
		index = index + 1
	}
	if len(m) > len(*keys) {
		// Add all remaining devices
		for k := range m {
			if _, ok := processed[k]; !ok {
				*keys = append(*keys, k)
				processed[k] = true
			}
		}
	}
	return
}

// checkDev validates whether a topology node can be assigned for a testbed device by validating
// required number of links are present between node peers.
// dev - is the testbed device
// node - is potential mapped node.
// res - is working resultset based on past selections.
// It returns a boolean; true if valid selection, false otherwise.
func (s *solver) checkDev(dev *opb.Device, node *tpb.Node, res map[*opb.Device]*tpb.Node) bool {
	for devPeer, linksPtr := range s.dev2Links[dev] {
		if nodePeer, ok := res[devPeer]; ok {
			devRLinks := len(*linksPtr)
			if _, ok := s.node2Links[node][nodePeer]; !ok {
				return false
			}
			nodeRLinks := len(*s.node2Links[node][nodePeer])
			if devRLinks > nodeRLinks {
				return false
			}
		}
	}
	return true
}

// checkLink validates whether a topology link can be assigned for a testbed link by validating
// ports assigned match requirements.
// d2n - device to node mapped for this solution set.
// devLink - is the testbed link.
// nodeLink - is potential mapped node link.
// res - is working resultset based on past selections.
// It returns a boolean; true if valid selection, false otherwise.
func (s *solver) checkLink(d2n map[*opb.Device]*tpb.Node, devLink *link, nodeLink *link, res map[*link]*link) bool {
	dev := s.id2Dev[devLink.name]
	dPort := s.dev2Ports[dev][devLink.port]
	pPort := s.dev2Ports[s.id2Dev[devLink.peerName]][devLink.peerPort]
	node := s.name2Node[nodeLink.name]
	nIntf := s.node2Intfs[node][nodeLink.port]
	pIntf := s.node2Intfs[s.name2Node[nodeLink.peerName]][nodeLink.peerPort]
	if d2n[dev] == node {
		if !s.portMatch(dPort, nIntf) || !s.portMatch(pPort, pIntf) {
			return false
		}
	} else {
		if !s.portMatch(dPort, pIntf) || !s.portMatch(pPort, nIntf) {
			return false
		}
	}
	return true
}

func (s *solver) solve() (*assign, error) {
	// Find all the matching device->node assignments, and
	// for each of those, all the port->intf assignments.
	dev2Nodes := make(map[*opb.Device][]*tpb.Node)
	for _, dut := range s.testbed.GetDuts() {
		nodeList, err := s.nodeMatches(dut, false)
		if err != nil {
			return nil, err
		}
		dev2Nodes[dut] = nodeList
	}
	for _, ate := range s.testbed.GetAtes() {
		nodeList, err := s.nodeMatches(ate, true)
		if err != nil {
			return nil, err
		}
		dev2Nodes[ate] = nodeList
	}

	// Generate possible testbed device -> topology node combinations.
	dev2NodeChan := s.genDevCombos(dev2Nodes)
	var hasNodeCombo bool
	for dev2Node := range dev2NodeChan {
		hasNodeCombo = true
		link2link := make(map[*link][]*link)
		s.getLinkMap(dev2Node, link2link)
		link2LinkChan := s.genLinkCombos(dev2Node, link2link)
		for dLink2nLink := range link2LinkChan {
			a := s.newAssign(dev2Node, dLink2nLink)
			if a != nil {
				// TODO: When solver is rewritten, signal the gorouting
				// channel to exit early here and avoid leaving the goroutine hanging.
				// Not disastrous but ideally the goroutines would terminate here.
				return a, nil
			}
		}
	}
	// Give a more specific error message for the case when we didn't need to even
	// consider the links to determine that there were no matching topologies.
	if !hasNodeCombo {
		return nil, fmt.Errorf("no combination of nodes in the KNE topology matches the testbed, irrespective of links")
	}
	return nil, fmt.Errorf("no KNE topology matches the testbed")
}

// getLinkMap generates the set of links for all TB links, for a device => node selection
// d2n - is the device => node map
// l2l - is the processed TB link => topology link
func (s *solver) getLinkMap(d2n map[*opb.Device]*tpb.Node, l2l map[*link][]*link) {
	processed := make(map[*opb.Device]bool)
	for dev, node := range d2n {
		for devPeer, links := range s.dev2Links[dev] {
			if _, ok := processed[devPeer]; ok {
				nodePeer := d2n[devPeer]
				nodeLinks := s.node2Links[node][nodePeer]
				for _, link := range *links {
					l2l[link] = *nodeLinks
				}
			}
		}
		processed[dev] = true
	}
}

func (s *solver) nodeMatches(dev *opb.Device, isATE bool) ([]*tpb.Node, error) {
	nodeList := []*tpb.Node{}
	for _, node := range s.topology.GetNodes() {
		if isATE != ateTypes[node.GetType()] {
			continue
		}
		match := s.devMatch(dev, node)
		if match {
			log.V(1).Infof("Found match: testbed device %q -> KNE topology node %q", dev.GetId(), node.GetName())
			nodeList = append(nodeList, node)
		}
	}
	if len(nodeList) == 0 {
		return nil, fmt.Errorf("no node in KNE topology to match testbed device %q", dev.GetId())
	}
	return nodeList, nil
}

func (s *solver) devMatch(dev *opb.Device, node *tpb.Node) bool {
	if dev.GetHardwareModel() != "" && dev.GetHardwareModel() != hardwareModel(node) {
		return false
	}
	if dev.GetSoftwareVersion() != "" && dev.GetSoftwareVersion() != softwareVersion(node) {
		return false
	}
	if v := dev.GetVendor(); v != opb.Device_VENDOR_UNSPECIFIED && v != type2VendorMap[node.GetType()] {
		return false
	}
	log.V(1).Infof("Found node match: %q", dev.GetId())
	intfs := s.node2Intfs[node]
	log.V(1).Infof("Interfaces: %v", intfs)
	// If the device needs more ports than the node, this node cannot match.
	if len(dev.GetPorts()) > len(intfs) {
		return false
	}
	return true
}

func (s *solver) portMatch(port *opb.Port, intf *intf) bool {
	// TODO: When solver is rewritten, support more generic port matching.
	if port.GetSpeed() != opb.Port_SPEED_UNSPECIFIED {
		return false
	}
	return true
}

func hardwareModel(node *tpb.Node) string {
	return tpb.Node_Type_name[int32(node.GetType())]
}

func softwareVersion(node *tpb.Node) string {
	return hardwareModel(node)
}

// genLinkCombos yields port link->node link mappings
func (s *solver) genLinkCombos(d2n map[*opb.Device]*tpb.Node, m map[*link][]*link) <-chan map[*link]*link {
	keys := []*link{}
	for k := range m {
		keys = append(keys, k)
	}
	ch := make(chan map[*link]*link)
	go func() {
		defer close(ch)
		s.genLinkRecurse(d2n, m, keys, make(map[*link]*link), make(map[*link]bool), ch)
	}()
	return ch
}

func (s *solver) genLinkRecurse(
	d2n map[*opb.Device]*tpb.Node,
	m map[*link][]*link,
	keys []*link,
	res map[*link]*link,
	used map[*link]bool,
	ch chan<- map[*link]*link) {
	if len(keys) == 0 {
		copy := make(map[*link]*link)
		for k, v := range res {
			copy[k] = v
		}
		ch <- copy
		return
	}
	first := keys[0]
	for _, i := range m[first] {
		if !used[i] {
			if !s.checkLink(d2n, first, i, res) {
				continue
			}
			res[first] = i
			used[i] = true
			s.genLinkRecurse(d2n, m, keys[1:], res, used, ch)
			delete(used, i)
		}
	}
}

// genDutCombos yields device -> node mappings
func (s *solver) genDevCombos(m map[*opb.Device][]*tpb.Node) <-chan map[*opb.Device]*tpb.Node {
	keys := []*opb.Device{}
	s.arrangeDevKeys(m, &keys)
	ch := make(chan map[*opb.Device]*tpb.Node)
	go func() {
		defer close(ch)
		s.genDevRecurse(m, keys, make(map[*opb.Device]*tpb.Node), make(map[*tpb.Node]bool), ch)
	}()
	return ch
}

func (s *solver) genDevRecurse(
	m map[*opb.Device][]*tpb.Node,
	keys []*opb.Device,
	res map[*opb.Device]*tpb.Node,
	used map[*tpb.Node]bool,
	ch chan<- map[*opb.Device]*tpb.Node) {
	if len(keys) == 0 {
		copy := make(map[*opb.Device]*tpb.Node)
		for k, v := range res {
			copy[k] = v
		}
		ch <- copy
		return
	}
	first := keys[0]
	for _, i := range m[first] {
		if !used[i] {
			if !s.checkDev(first, i, res) {
				continue
			}
			res[first] = i
			used[i] = true
			s.genDevRecurse(m, keys[1:], res, used, ch)
			delete(used, i)
			delete(res, first)
		}
	}
}

func (s *solver) newAssign(dev2Node map[*opb.Device]*tpb.Node, link2link map[*link]*link) *assign {
	a := &assign{
		dev2Node:  make(map[*opb.Device]*tpb.Node),
		port2Intf: make(map[*opb.Port]*intf),
	}
	processedPorts := make(map[*opb.Port]bool)
	processedIntfs := make(map[*intf]bool)
	// From the mapped link => link, build up the port => intf map
	for dLink, nLink := range link2link {
		dev := s.id2Dev[dLink.name]
		dPort := s.dev2Ports[dev][dLink.port]
		pPort := s.dev2Ports[s.id2Dev[dLink.peerName]][dLink.peerPort]
		node := s.name2Node[nLink.name]
		nIntf := s.node2Intfs[node][nLink.port]
		pIntf := s.node2Intfs[s.name2Node[nLink.peerName]][nLink.peerPort]
		if dev2Node[dev] == node {
			a.port2Intf[dPort] = nIntf
			a.port2Intf[pPort] = pIntf
		} else {
			a.port2Intf[dPort] = pIntf
			a.port2Intf[pPort] = nIntf
		}
		processedPorts[dPort] = true
		processedPorts[pPort] = true
		processedIntfs[nIntf] = true
		processedIntfs[pIntf] = true
	}
	for d, n := range dev2Node {
		a.dev2Node[d] = n
		// Assign remaining ports not connected to any device
		for _, port := range s.dev2Ports[d] {
			if _, ok := processedPorts[port]; !ok {
				for _, intf := range s.node2Intfs[n] {
					if _, ok = processedIntfs[intf]; !ok {
						a.port2Intf[port] = intf
						processedPorts[port] = true
						processedIntfs[intf] = true
						break
					}
				}
			}
		}
	}
	return a
}
