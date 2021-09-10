package knebind

import (
	"fmt"
	"strings"

	log "github.com/golang/glog"
	"github.com/pkg/errors"

	kpb "github.com/google/kne/proto/topo"
	opb "github.com/openconfig/ondatra/proto"
)

func solve(tb *opb.Testbed, topo *kpb.Topology) (*assign, error) {
	if len(tb.GetAtes()) > 0 {
		return nil, errors.New("KNE binding does not yet support ATEs")
	}
	if numDUTs, numNodes := len(tb.GetDuts()), len(topo.GetNodes()); numDUTs > numNodes {
		return nil, errors.Errorf("Not enough nodes in KNE topology for specified testbed: "+
			" testbed has %d DUTs and topology only has %d nodes", numDUTs, numNodes)
	}
	if numTBLinks, numTopoLinks := len(tb.GetLinks()), len(topo.GetLinks()); numTBLinks > numTopoLinks {
		return nil, errors.Errorf("Not enough links in KNE topology for specified testbed "+
			" testbed has %d links and topology only has %d links", numTBLinks, numTopoLinks)
	}
	s := &solver{
		testbed:    tb,
		topology:   topo,
		id2DUT:     make(map[string]*opb.Device),
		dut2Ports:  make(map[*opb.Device]map[string]*opb.Port),
		node2Intfs: make(map[*kpb.Node]map[string]*intf),
		intf2Intf:  make(map[*intf]*intf),
	}

	// Cache various info in the solver about the testbed and topology.
	for _, dut := range s.testbed.GetDuts() {
		s.id2DUT[dut.GetId()] = dut
		ports := make(map[string]*opb.Port)
		for _, port := range dut.GetPorts() {
			ports[port.GetId()] = port
		}
		s.dut2Ports[dut] = ports
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
	dut2Node  map[*opb.Device]*kpb.Node
	port2Intf map[*opb.Port]*intf
}

type intf struct {
	name string
}

func (a *assign) String() string {
	var sb strings.Builder
	for dut, node := range a.dut2Node {
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
	id2DUT     map[string]*opb.Device
	dut2Ports  map[*opb.Device]map[string]*opb.Port
	node2Intfs map[*kpb.Node]map[string]*intf
	intf2Intf  map[*intf]*intf
}

func (s *solver) solve() (*assign, error) {
	// Find all the matching dut->node assignments, and
	// for each of those, all the port->intf assignments.
	dut2Nodes := make(map[interface{}][]interface{})
	dut2Node2Port2Infs := make(map[*opb.Device]map[*kpb.Node]map[*opb.Port][]interface{})
	for _, dut := range s.testbed.GetDuts() {
		dut2Node2Port2Infs[dut] = make(map[*kpb.Node]map[*opb.Port][]interface{})
		var nodes []interface{}
		for _, node := range s.topology.GetNodes() {
			match, port2Infs := s.dutMatch(dut, node)
			if match {
				log.V(1).Infof("Found match testbed: %q -> topology: %q", dut.GetId(), node.GetName())
				nodes = append(nodes, node)
				dut2Node2Port2Infs[dut][node] = port2Infs
			}
		}
		if len(nodes) == 0 {
			return nil, errors.Errorf("No node in KNE topology to match testbed %q", dut.GetId())
		}
		dut2Nodes[dut] = nodes
	}

	// Iterate over each of the possible testbed->topology combinations.
	dut2NodeChan := genCombos(dut2Nodes)
	var hasNodeCombo bool
	for dut2Node := range dut2NodeChan {
		hasNodeCombo = true
		port2Intfs := make(map[interface{}][]interface{})
		for dut, node := range dut2Node {
			for port, intfs := range dut2Node2Port2Infs[dut.(*opb.Device)][node.(*kpb.Node)] {
				port2Intfs[port] = intfs
			}
		}
		port2IntfChan := genCombos(port2Intfs)
		for port2Intf := range port2IntfChan {
			if a := newAssign(dut2Node, port2Intf); s.linksMatch(a) {
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

func (s *solver) linksMatch(a *assign) bool {
	getIntf := func(tbLink string) *intf {
		parts := strings.Split(tbLink, ":")
		dut := s.id2DUT[parts[0]]
		port := s.dut2Ports[dut][parts[1]]
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

func (s *solver) dutMatch(dut *opb.Device, node *kpb.Node) (bool, map[*opb.Port][]interface{}) {
	if dut.GetHardwareModel() != "" && dut.GetHardwareModel() != hardwareModel(node) {
		return false, nil
	}
	if dut.GetSoftwareVersion() != "" && dut.GetSoftwareVersion() != softwareVersion(node) {
		return false, nil
	}
	if v := dut.GetVendor(); v != opb.Device_UNKNOWN && v != type2VendorMap[node.GetType()] {
		return false, nil
	}
	log.V(1).Infof("Found node match: %q", dut.GetId())
	intfs := s.node2Intfs[node]
	log.V(1).Infof("Interfaces: %v", intfs)
	// If the DUT needs more ports than the node, this node cannot match.
	if len(dut.GetPorts()) > len(intfs) {
		return false, nil
	}
	port2Infs := make(map[*opb.Port][]interface{})
	for _, port := range dut.GetPorts() {
		var infs []interface{}
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

func newAssign(dut2Node, port2Intf map[interface{}]interface{}) *assign {
	a := &assign{
		dut2Node:  make(map[*opb.Device]*kpb.Node),
		port2Intf: make(map[*opb.Port]*intf),
	}
	for d, n := range dut2Node {
		a.dut2Node[d.(*opb.Device)] = n.(*kpb.Node)
	}
	for p, i := range port2Intf {
		a.port2Intf[p.(*opb.Port)] = i.(*intf)
	}
	return a
}
