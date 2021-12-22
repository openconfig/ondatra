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

package knebind

import (
	"fmt"
	"strings"

	log "github.com/golang/glog"
	"github.com/pkg/errors"

	kpb "github.com/google/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

var ateTypes = map[kpb.Node_Type]bool{
	kpb.Node_IXIA_TG: true,
}

func solve(tb *opb.Testbed, topo *kpb.Topology) (*assign, error) {
	devs := append(append([]*opb.Device{}, tb.GetDuts()...), tb.GetAtes()...)
	if numDevs, numNodes := len(devs), len(topo.GetNodes()); numDevs > numNodes {
		return nil, errors.Errorf("Not enough nodes in KNE topology for specified testbed: "+
			" testbed has %d devices and topology only has %d nodes", numDevs, numNodes)
	}
	if numTBLinks, numTopoLinks := len(tb.GetLinks()), len(topo.GetLinks()); numTBLinks > numTopoLinks {
		return nil, errors.Errorf("Not enough links in KNE topology for specified testbed "+
			" testbed has %d links and topology only has %d links", numTBLinks, numTopoLinks)
	}
	s := &solver{
		testbed:    tb,
		topology:   topo,
		id2Dev:     make(map[string]*opb.Device),
		dev2Ports:  make(map[*opb.Device]map[string]*opb.Port),
		node2Intfs: make(map[*kpb.Node]map[string]*intf),
		intf2Intf:  make(map[*intf]*intf),
	}

	// Cache various info in the solver about the testbed and topology.
	for _, dev := range devs {
		s.id2Dev[dev.GetId()] = dev
		ports := make(map[string]*opb.Port)
		for _, port := range dev.GetPorts() {
			ports[port.GetId()] = port
		}
		s.dev2Ports[dev] = ports
	}
	name2Node := make(map[string]*kpb.Node)
	for _, node := range s.topology.GetNodes() {
		name2Node[node.GetName()] = node
		s.node2Intfs[node] = make(map[string]*intf)
	}
	getIntf := func(nodeName, intfName string) *intf {
		node := name2Node[nodeName]
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
		intfA := getIntf(link.GetANode(), link.GetAInt())
		intfZ := getIntf(link.GetZNode(), link.GetZInt())
		s.intf2Intf[intfA] = intfZ
		s.intf2Intf[intfZ] = intfA
	}

	return s.solve()
}

type assign struct {
	dev2Node  map[*opb.Device]*kpb.Node
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

type solver struct {
	testbed    *opb.Testbed
	topology   *kpb.Topology
	id2Dev     map[string]*opb.Device
	dev2Ports  map[*opb.Device]map[string]*opb.Port
	node2Intfs map[*kpb.Node]map[string]*intf
	intf2Intf  map[*intf]*intf
}

func (s *solver) solve() (*assign, error) {
	// Find all the matching device->node assignments, and
	// for each of those, all the port->intf assignments.
	dev2Node2Port2Intfs := make(map[*opb.Device]map[*kpb.Node]map[*opb.Port][]*intf)
	for _, dut := range s.testbed.GetDuts() {
		node2Port2Intfs, err := s.nodeMatches(dut, false)
		if err != nil {
			return nil, err
		}
		dev2Node2Port2Intfs[dut] = node2Port2Intfs
	}
	for _, ate := range s.testbed.GetAtes() {
		node2Port2Intfs, err := s.nodeMatches(ate, true)
		if err != nil {
			return nil, err
		}
		dev2Node2Port2Intfs[ate] = node2Port2Intfs
	}

	// Iterate over each of the possible testbed->topology combinations.
	dev2Nodes := make(map[interface{}][]interface{})
	for dev, node2Port2Intfs := range dev2Node2Port2Intfs {
		for node := range node2Port2Intfs {
			dev2Nodes[dev] = append(dev2Nodes[dev], node)
		}
	}
	dev2NodeChan := genCombos(dev2Nodes)
	var hasNodeCombo bool
	for dev2Node := range dev2NodeChan {
		hasNodeCombo = true
		port2Intfs := make(map[interface{}][]interface{})
		for dut, node := range dev2Node {
			for port, intfs := range dev2Node2Port2Intfs[dut.(*opb.Device)][node.(*kpb.Node)] {
				for _, i := range intfs {
					port2Intfs[port] = append(port2Intfs[port], i)
				}
			}
		}
		port2IntfChan := genCombos(port2Intfs)
		for port2Intf := range port2IntfChan {
			if a := newAssign(dev2Node, port2Intf); s.linksMatch(a) {
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
		return nil, errors.Errorf("No combination of nodes in the KNE topology matches the testbed, irrespective of links")
	}
	return nil, errors.Errorf("No KNE topology matches the testbed")
}

func (s *solver) nodeMatches(dev *opb.Device, isATE bool) (map[*kpb.Node]map[*opb.Port][]*intf, error) {
	node2Port2Intfs := make(map[*kpb.Node]map[*opb.Port][]*intf)
	for _, node := range s.topology.GetNodes() {
		if isATE != ateTypes[node.GetType()] {
			continue
		}
		match, port2Intfs := s.devMatch(dev, node)
		if match {
			log.V(1).Infof("Found match: testbed device %q -> KNE topology node %q", dev.GetId(), node.GetName())
			node2Port2Intfs[node] = port2Intfs
		}
	}
	if len(node2Port2Intfs) == 0 {
		return nil, errors.Errorf("No node in KNE topology to match testbed device %q", dev.GetId())
	}
	return node2Port2Intfs, nil
}

func (s *solver) devMatch(dev *opb.Device, node *kpb.Node) (bool, map[*opb.Port][]*intf) {
	if dev.GetHardwareModel() != "" && dev.GetHardwareModel() != hardwareModel(node) {
		return false, nil
	}
	if dev.GetSoftwareVersion() != "" && dev.GetSoftwareVersion() != softwareVersion(node) {
		return false, nil
	}
	if v := dev.GetVendor(); v != opb.Device_UNKNOWN && v != type2VendorMap[node.GetType()] {
		return false, nil
	}
	log.V(1).Infof("Found node match: %q", dev.GetId())
	intfs := s.node2Intfs[node]
	log.V(1).Infof("Interfaces: %v", intfs)
	// If the device needs more ports than the node, this node cannot match.
	if len(dev.GetPorts()) > len(intfs) {
		return false, nil
	}
	port2Infs := make(map[*opb.Port][]*intf)
	for _, port := range dev.GetPorts() {
		var infs []*intf
		for _, intf := range intfs {
			if s.portMatch(port, intf) {
				log.V(1).Infof("Matched Port: %q %q", port.GetId(), intf.name)
				infs = append(infs, intf)
			}
		}
		if len(intfs) == 0 {
			return false, nil
		}
		port2Infs[port] = infs
	}
	return true, port2Infs
}

func (s *solver) portMatch(port *opb.Port, intf *intf) bool {
	// TODO: When solver is rewritten, support more generic port matching.
	if port.GetSpeed() != opb.Port_S_UNKNOWN {
		return false
	}
	return true
}

func (s *solver) linksMatch(a *assign) bool {
	getIntf := func(tbLink string) *intf {
		parts := strings.Split(tbLink, ":")
		dut := s.id2Dev[parts[0]]
		port := s.dev2Ports[dut][parts[1]]
		return a.port2Intf[port]
	}
	for _, link := range s.testbed.GetLinks() {
		intfA := getIntf(link.GetA())
		intfB := getIntf(link.GetB())
		if s.intf2Intf[intfA] != intfB {
			return false
		}
	}
	return true
}

func hardwareModel(node *kpb.Node) string {
	return kpb.Node_Type_name[int32(node.GetType())]
}

func softwareVersion(node *kpb.Node) string {
	return hardwareModel(node)
}

// genCombos yields every key->value mapping, where no two keys map to the same
// value, given a map of keys to their possible values.
func genCombos(m map[interface{}][]interface{}) <-chan map[interface{}]interface{} {
	var keys []interface{}
	for k := range m {
		keys = append(keys, k)
	}
	ch := make(chan map[interface{}]interface{})
	go func() {
		defer close(ch)
		genRecurse(m, keys, make(map[interface{}]interface{}), make(map[interface{}]bool), ch)
	}()
	return ch
}

func genRecurse(
	m map[interface{}][]interface{},
	keys []interface{},
	res map[interface{}]interface{},
	used map[interface{}]bool,
	ch chan<- map[interface{}]interface{}) {
	if len(keys) == 0 {
		copy := make(map[interface{}]interface{})
		for k, v := range res {
			copy[k] = v
		}
		ch <- copy
		return
	}
	first := keys[0]
	for _, i := range m[first] {
		if !used[i] {
			res[first] = i
			used[i] = true
			genRecurse(m, keys[1:], res, used, ch)
			delete(used, i)
		}
	}
}

func newAssign(dev2Node, port2Intf map[interface{}]interface{}) *assign {
	a := &assign{
		dev2Node:  make(map[*opb.Device]*kpb.Node),
		port2Intf: make(map[*opb.Port]*intf),
	}
	for d, n := range dev2Node {
		a.dev2Node[d.(*opb.Device)] = n.(*kpb.Node)
	}
	for p, i := range port2Intf {
		a.port2Intf[p.(*opb.Port)] = i.(*intf)
	}
	return a
}
