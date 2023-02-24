// Copyright 2022 Google LLC
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

// Package portgraph searches for a graph that satisfies a set of constraints.
package portgraph

import (
	"fmt"
	"regexp"
	"sort"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/internal/orderedmap"
	"github.com/pkg/errors"
)

var (
	reAny = Regex(regexp.MustCompile(".*"))
)

// AbstractGraph is a representation of abstract nodes, ports, and edges.
type AbstractGraph struct {
	Desc  string // Description for the AbstractGraph for logging.
	Nodes []*AbstractNode
	Edges []*AbstractEdge
}

func (g *AbstractGraph) String() string {
	ret := fmt.Sprintf("Abstract Graph: %q\n", g.Desc)
	for _, n := range g.Nodes {
		ret = fmt.Sprintf("%sNode: %q\n", ret, n.Desc)
		for k, c := range n.Constraints {
			ret = fmt.Sprintf("%s  Constraint %q: %v\n", ret, k, c)
		}
		for _, p := range n.Ports {
			ret = fmt.Sprintf("%s  Port: %q\n", ret, p.Desc)
			for k, c := range p.Constraints {
				ret = fmt.Sprintf("%s    Constraint %q: %v\n", ret, k, c)
			}
		}
	}
	for _, e := range g.Edges {
		ret = fmt.Sprintf("%sEdge: %q -> %q\n", ret, e.Src.Desc, e.Dst.Desc)
	}
	return ret
}

// fetchPort2Port2EdgeMap() returns the mapping of a src AbstractPort to a dst AbstractPort to an AbstractEdge.
func (g *AbstractGraph) fetchPort2Port2EdgeMap() (map[*AbstractPort]map[*AbstractPort]*AbstractEdge, error) {
	ret := make(map[*AbstractPort]map[*AbstractPort]*AbstractEdge)
	for _, e := range g.Edges {
		if _, ok := ret[e.Src]; ok {
			return nil, fmt.Errorf("port %q is attached to more than one other port; can only be attached to one", e.Src.Desc)
		}
		if _, ok := ret[e.Dst]; ok {
			return nil, fmt.Errorf("port %q is attached to more than one other port; can only be attached to one", e.Dst.Desc)
		}
		ret[e.Src] = map[*AbstractPort]*AbstractEdge{e.Dst: e}
		ret[e.Dst] = map[*AbstractPort]*AbstractEdge{e.Src: &AbstractEdge{e.Dst, e.Src}}
	}
	return ret, nil
}

// ConcreteGraph is a representation of concrete nodes, ports, and edges.
type ConcreteGraph struct {
	Desc  string // Description for the AbstractGraph for logging.
	Nodes []*ConcreteNode
	Edges []*ConcreteEdge
}

func (g *ConcreteGraph) String() string {
	ret := fmt.Sprintf("Concrete Graph: %q\n", g.Desc)
	for _, n := range g.Nodes {
		ret = fmt.Sprintf("%sNode: %q\n", ret, n.Desc)
		for k, a := range n.Attrs {
			ret = fmt.Sprintf("%s  Attribute %q: %v\n", ret, k, a)
		}
		for _, p := range n.Ports {
			ret = fmt.Sprintf("%s  Port: %q\n", ret, p.Desc)
			for k, a := range p.Attrs {
				ret = fmt.Sprintf("%s    Attribute %q: %v\n", ret, k, a)
			}
		}
	}
	for _, e := range g.Edges {
		ret = fmt.Sprintf("%sEdge: %q -> %q\n", ret, e.Src.Desc, e.Dst.Desc)
	}
	return ret
}

// fetchPort2Port2EdgeMap() returns the mapping of a src ConcretePort to a dst ConcretePort to a ConcreteEdge.
func (g *ConcreteGraph) fetchPort2Port2EdgeMap() map[*ConcretePort]map[*ConcretePort]*ConcreteEdge {
	ret := make(map[*ConcretePort]map[*ConcretePort]*ConcreteEdge)
	for _, e := range g.Edges {
		if m, ok := ret[e.Src]; ok {
			m[e.Dst] = e
		} else {
			ret[e.Src] = map[*ConcretePort]*ConcreteEdge{e.Dst: e}
		}
		if m, ok := ret[e.Dst]; ok {
			m[e.Src] = &ConcreteEdge{e.Dst, e.Src}
		} else {
			ret[e.Dst] = map[*ConcretePort]*ConcreteEdge{e.Src: &ConcreteEdge{e.Dst, e.Src}}
		}
	}
	return ret
}

// AbstractNode represent a vertex on an AbstractGraph.
// The AbstractNode has AbstractPorts and AbstractEdges that originate and terminate at AbstractPorts.
type AbstractNode struct {
	Desc        string                    // Description for the AbstractNode for logging.
	Ports       []*AbstractPort           // A list AbstractPorts that may be connected to another AbstractPort.
	Constraints map[string]NodeConstraint // A map of key value constraints of the AbstractNode.
}

// ConcreteNode represent a vertex on a ConcreteGraph.
// The ConcreteNode has ConcretePorts and ConcreteEdges that originate and terminate at ConcretePorts.
type ConcreteNode struct {
	Desc  string            // Description for the ConcreteNode for logging.
	Ports []*ConcretePort   // A list ConcretePorts that may be connected to another ConcretePort.
	Attrs map[string]string // A map of key value attributes of the ConctreteNode.
}

func (n *ConcreteNode) containsPort(p *ConcretePort) bool {
	for _, p2 := range n.Ports {
		if p == p2 {
			return true
		}
	}
	return false
}

// AbstractEdge represents a link from an source AbstractPort to a destination AbstractPort.
type AbstractEdge struct {
	Src, Dst *AbstractPort
}

func (e *AbstractEdge) String() string {
	if e.Dst == nil {
		return fmt.Sprintf("%q -> nil", e.Src.Desc)
	}
	return fmt.Sprintf("%q -> %q", e.Src.Desc, e.Dst.Desc)
}

// ConcreteEdge represents a link from an source ConcretePort to a destination ConcretePort.
type ConcreteEdge struct {
	Src, Dst *ConcretePort
}

// AbstractPort is a is point on an AbstractNode where an AbstractEdge may be attached.
type AbstractPort struct {
	Desc        string                    // Description for the AbstractPort for logging.
	Constraints map[string]PortConstraint // A map of key value constraints of the AbstractPort.
}

// ConcretePort is a is point on a ConcreteNode where a ConcreteEdge may be attached.
type ConcretePort struct {
	Desc  string            // Description for the ConcretePort for logging.
	Attrs map[string]string // A map of key value attributes of the ConcretePort.
}

// Assignment contains the AbstractNode -> ConcreteNode and AbstractPort -> ConcretePort mappings.
type Assignment struct {
	Node2Node map[*AbstractNode]*ConcreteNode
	Port2Port map[*AbstractPort]*ConcretePort
}

type maxAssignment struct {
	assignment         *Assignment
	numNodes, numPorts int
}

type deferredNodeConstraint func(map[*AbstractNode]*ConcreteNode) bool
type deferredPortConstraint func(map[*AbstractPort]*ConcretePort) (bool, *AbstractPort, error)

type solver struct {
	abstractGraph         *AbstractGraph
	superGraph            *ConcreteGraph
	absPort2Node          map[*AbstractPort]*AbstractNode                   // Map Port to Node it is part of.
	absNode2Node2NumEdges map[*AbstractNode]map[*AbstractNode]int           // Map Node to Node and how many edges there are.
	conNode2Node2NumEdges map[*ConcreteNode]map[*ConcreteNode]int           // Map Node to Node and how many edges there are.
	conPort2Port2Edge     map[*ConcretePort]map[*ConcretePort]*ConcreteEdge // Cache the linked concrete ports to edge.
	absPort2Port2Edge     map[*AbstractPort]map[*AbstractPort]*AbstractEdge // Cache the linked abstract ports to edge.

	maxAssign *maxAssignment // The "best" Assignment for reporting if the solve failed.

	// Constraint mappings. Deferred constraint can only be checked after all abstract elements are assigned.
	// Constraints are ordered for deterministic iteration.
	nodeConstraints         *orderedmap.OrderedMap[*AbstractNode, *orderedmap.OrderedMap[string, LeafConstraint]]
	portConstraints         *orderedmap.OrderedMap[*AbstractPort, *orderedmap.OrderedMap[string, LeafConstraint]]
	deferredNodeConstraints *orderedmap.OrderedMap[*AbstractNode, []deferredNodeConstraint]
	deferredPortConstraints *orderedmap.OrderedMap[*AbstractPort, []deferredPortConstraint]
}

// Solve returns an assignment from superGraph that satisfies abstractGraph.
func Solve(abstractGraph *AbstractGraph, superGraph *ConcreteGraph) (*Assignment, error) {
	solveErr := &SolveErr{absGraphDesc: abstractGraph.Desc, conGraphDesc: superGraph.Desc}
	if len(abstractGraph.Nodes) > len(superGraph.Nodes) {
		return nil, solveErr

	}
	if len(abstractGraph.Edges) > len(superGraph.Edges) {
		return nil, solveErr
	}

	// Preprocess data for the solve.
	// Map the AbstractPort to AbstractNode and initialize maps for maxSolve.
	absPort2Node := make(map[*AbstractPort]*AbstractNode)
	node2Node := make(map[*AbstractNode]*ConcreteNode)
	port2Port := make(map[*AbstractPort]*ConcretePort)
	for _, n := range abstractGraph.Nodes {
		node2Node[n] = nil
		for _, p := range n.Ports {
			absPort2Node[p] = n
			port2Port[p] = nil
		}
	}
	// Map how many links there are between each AbstractNode to calculate matches.
	absNode2Node2NumEdges := make(map[*AbstractNode]map[*AbstractNode]int)
	for _, e := range abstractGraph.Edges {
		srcNode, dstNode := absPort2Node[e.Src], absPort2Node[e.Dst]
		// Count the edges from src -> dst
		if _, ok := absNode2Node2NumEdges[srcNode]; !ok {
			absNode2Node2NumEdges[srcNode] = make(map[*AbstractNode]int)
		}
		absNode2Node2NumEdges[srcNode][dstNode]++
		// Count the edges from dst -> src
		if _, ok := absNode2Node2NumEdges[dstNode]; !ok {
			absNode2Node2NumEdges[dstNode] = make(map[*AbstractNode]int)
		}
		absNode2Node2NumEdges[dstNode][srcNode]++
	}

	absPort2Port2Edge, err := abstractGraph.fetchPort2Port2EdgeMap()
	if err != nil {
		return nil, solveErr
	}

	conPort2Node := make(map[*ConcretePort]*ConcreteNode)
	for _, n := range superGraph.Nodes {
		for _, p := range n.Ports {
			conPort2Node[p] = n
		}
	}

	// Map how many links there are between each ConcreteNode to calculate matches.
	conNode2Node2NumEdges := make(map[*ConcreteNode]map[*ConcreteNode]int)
	for _, e := range superGraph.Edges {
		srcNode, dstNode := conPort2Node[e.Src], conPort2Node[e.Dst]
		// Count the edges from src -> dst
		if _, ok := conNode2Node2NumEdges[srcNode]; !ok {
			conNode2Node2NumEdges[srcNode] = make(map[*ConcreteNode]int)
		}
		conNode2Node2NumEdges[srcNode][dstNode]++
		// Count the edges from dst -> src
		if _, ok := conNode2Node2NumEdges[dstNode]; !ok {
			conNode2Node2NumEdges[dstNode] = make(map[*ConcreteNode]int)
		}
		conNode2Node2NumEdges[dstNode][srcNode]++
	}

	s := &solver{
		abstractGraph:           abstractGraph,
		superGraph:              superGraph,
		absNode2Node2NumEdges:   absNode2Node2NumEdges,
		absPort2Node:            absPort2Node,
		absPort2Port2Edge:       absPort2Port2Edge,
		conNode2Node2NumEdges:   conNode2Node2NumEdges,
		maxAssign:               &maxAssignment{&Assignment{Node2Node: node2Node, Port2Port: port2Port}, 0, 0},
		nodeConstraints:         orderedmap.NewOrderedMap[*AbstractNode, *orderedmap.OrderedMap[string, LeafConstraint]](),
		deferredNodeConstraints: orderedmap.NewOrderedMap[*AbstractNode, []deferredNodeConstraint](),
		portConstraints:         orderedmap.NewOrderedMap[*AbstractPort, *orderedmap.OrderedMap[string, LeafConstraint]](),
		deferredPortConstraints: orderedmap.NewOrderedMap[*AbstractPort, []deferredPortConstraint](),
	}

	a, ok := s.solve()
	if !ok {
		solveErr.maxAssign = s.maxAssign.assignment
		return nil, solveErr
	}
	return a, nil
}

// solve provides a mapping of abstract nodes and ports to concrete nodes and ports.
// 1. Find all concrete nodes that match the abstract node (check attributes and # of ports).
// 2. For each mapping, check the concrete nodes and satisfy the edges of the abstract graph.
// 3. Assign concrete ports.
// solve always returns the max assignment and true if the solve completed.
func (s *solver) solve() (*Assignment, bool) {
	abs2ConNodes := make(map[*AbstractNode][]*ConcreteNode)
	s.processConstraints()
	for _, n := range s.abstractGraph.Nodes {
		nodes := s.matchNodes(n, s.superGraph)
		if len(nodes) == 0 {
			return nil, false
		}
		for _, conNode := range nodes {
			abs2ConNodes[n] = append(abs2ConNodes[n], conNode)
		}
	}

	s.conPort2Port2Edge = s.superGraph.fetchPort2Port2EdgeMap()

	// Generate all AbstractNode -> ConcreteNode mappings.
	abs2ConNodeChan, stop := genNodeCombos(abs2ConNodes, s.absNode2Node2NumEdges, s.checkEdge)
	defer stop()
	// Iterate through each mapping.
	// For each mapping, evaluate deferred constraints and attempt to match edges and assign ports.
	for abs2ConNode := range abs2ConNodeChan {
		if !s.matchDeferredNodes(abs2ConNode) {
			continue
		}
		if len(abs2ConNode) > s.maxAssign.numNodes {
			s.maxAssign.assignment.Node2Node = abs2ConNode
			s.maxAssign.numNodes = len(abs2ConNode)
		}

		// Since the edges can be satisfied, try to assign matching ports.
		abs2ConPort := s.assignEdges(abs2ConNode)
		if abs2ConPort == nil {
			continue
		}

		return &Assignment{abs2ConNode, abs2ConPort}, true
	}
	return nil, false
}

// checkEdge validates there are enough Ports to fulfill the Edge between the given Nodes and
// returns all Ports that could fulfill the Edge.
func (s *solver) checkEdge(conSrcNode, conDstNode *ConcreteNode, edgesNeeded int) (bool, []*ConcretePort, []*ConcretePort) {
	var srcPorts, conPorts []*ConcretePort
	// Verify whether the concrete Edges satisfy the abstract Edges.
	for _, conSrcPort := range conSrcNode.Ports {
		for p := range s.conPort2Port2Edge[conSrcPort] {
			if conDstNode.containsPort(p) {
				srcPorts = append(srcPorts, conSrcPort)
				conPorts = append(conPorts, p)
			}
		}
	}
	if len(srcPorts) >= edgesNeeded {
		return true, srcPorts, conPorts
	}
	return false, nil, nil
}

type portAssigner struct {
	// Pre-processed data for this port solve. These maps should not change.
	absPort2Edge            map[*AbstractPort]*AbstractEdge
	abs2ConEdges            map[*AbstractEdge][]*ConcreteEdge
	deferredPortConstraints *orderedmap.OrderedMap[*AbstractPort, []deferredPortConstraint]
	orderedAbsEdges         []*AbstractEdge

	// Stored data between recursions.
	deferredUntilPortConstraints *orderedmap.OrderedMap[*AbstractPort, []deferredPortConstraint]
	abs2ConPorts                 map[*AbstractPort]*ConcretePort // Ports that have been assigned.
	assignedConPorts             map[*ConcretePort]struct{}
	assignQueue                  *assignQueue

	// maxAbs2ConPorts stores the best solve (that is not complete) for error reporting.
	maxAbs2ConPorts map[*AbstractPort]*ConcretePort
}

// assignQueue is the queue of Edges to solve next and stores the previous queue state.
type assignQueue struct {
	prev  *assignQueue
	queue []*AbstractEdge // Assign ports as edges; this significantly reduces search space.
}

func (aq *assignQueue) newQueue() *assignQueue {
	newQueue := &assignQueue{
		prev:  aq,
		queue: aq.queue,
	}
	return newQueue
}

func (pa *portAssigner) nextEdge() *AbstractEdge {
	for _, e := range pa.orderedAbsEdges {
		if _, ok := pa.abs2ConPorts[e.Src]; !ok {
			return e
		}
	}
	return nil
}

func (pa *portAssigner) assign() bool {
	var absEdge *AbstractEdge
	// If there are enqueued edges, pop until we find one that hasn't been assigned.
	for len(pa.assignQueue.queue) > 0 {
		absEdge = pa.assignQueue.queue[len(pa.assignQueue.queue)-1]
		pa.assignQueue.queue = pa.assignQueue.queue[:len(pa.assignQueue.queue)-1]
		if _, ok := pa.abs2ConPorts[absEdge.Src]; !ok {
			break
		}
		absEdge = nil
	}
	if absEdge == nil {
		absEdge = pa.nextEdge()
		if absEdge == nil {
			// Done.
			return true
		}
	}

	// Assign the port(s) in the edge.
	absSrcPort, absDstPort := absEdge.Src, absEdge.Dst

	// The ports that were assigned/enqueued.
	var assignedAbsPorts []*AbstractPort
	var assignedConPorts []*ConcretePort
	var numEnqueuedPortConstraints map[*AbstractPort]int
	rewindState := func() {
		for _, p := range assignedAbsPorts {
			delete(pa.abs2ConPorts, p)
		}
		for _, p := range assignedConPorts {
			delete(pa.assignedConPorts, p)
		}
		for p, i := range numEnqueuedPortConstraints {
			pcs, ok := pa.deferredUntilPortConstraints.Get(p)
			if !ok {
				log.Fatalf("Impossible: cannot remove non-existant port constraints")
			}
			if len(pcs) <= i {
				pa.deferredUntilPortConstraints.Delete(p)
			} else {
				pa.deferredUntilPortConstraints.Set(p, pcs[:len(pcs)-i])
			}
		}
	}
	for _, conEdge := range pa.abs2ConEdges[absEdge] {
		conSrcPort, conDstPort := conEdge.Src, conEdge.Dst
		if _, ok := pa.assignedConPorts[conSrcPort]; ok {
			continue
		} else if _, ok := pa.assignedConPorts[conDstPort]; ok {
			continue
		}

		// Reset state for each iteration.
		queue := pa.assignQueue.newQueue()
		tryPortAssignment := func(queue *assignQueue) bool {
			assignedAbsPorts = nil
			assignedConPorts = nil
			numEnqueuedPortConstraints = make(map[*AbstractPort]int, len(pa.absPort2Edge))

			// Try assigning the ports.
			assign := func(a *AbstractPort, c *ConcretePort) {
				pa.abs2ConPorts[a] = c
				pa.assignedConPorts[c] = struct{}{}
				assignedAbsPorts = append(assignedAbsPorts, a)
				assignedConPorts = append(assignedConPorts, c)
			}
			assign(absSrcPort, conSrcPort)
			if absDstPort != nil {
				assign(absDstPort, conDstPort)
			}

			// Check whether the assigned ports work with constraints.
			checkAssignment := func(absPort *AbstractPort) bool {
				// Start with constraints from previously assigned ports that depend on the port currently being assigned.
				if dcs, ok := pa.deferredUntilPortConstraints.Get(absPort); ok {
					for _, dc := range dcs {
						ok, _, err := dc(pa.abs2ConPorts)
						if err != nil {
							// This is an impossibility and should not happen.
							log.Fatalf("could not solve: %v", err)
						} else if !ok {
							// This ConcreteEdge does not work with previous assignments.
							return false
						}
					}
				}
				// Check deferred constraints.
				// If the constraint depends on another port that has not been assigned yet, defer the constraint
				// and add associated edge to the queue.
				if constraints, ok := pa.deferredPortConstraints.Get(absPort); ok {
					for _, dc := range constraints {
						if ok, p, err := dc(pa.abs2ConPorts); err != nil {
							dpcs, _ := pa.deferredUntilPortConstraints.Get(p)
							pa.deferredUntilPortConstraints.Set(p, append(dpcs, dc))
							queue.queue = append(queue.queue, pa.absPort2Edge[p])
							numEnqueuedPortConstraints[p]++
						} else if !ok {
							return false
						}
					}
				}
				return true
			}
			if ok := checkAssignment(absSrcPort); !ok {
				return false
			}
			if absDstPort != nil {
				if ok := checkAssignment(absDstPort); !ok {
					return false
				}
			}
			return true
		}
		// Try to assign the ports.
		// Evaluate deferredPortConstraints. If this assignment does not work with other currently
		// assigned ports, rewind the state and try the next assignment.
		if !tryPortAssignment(pa.assignQueue) {
			rewindState()
			continue
		}
		// This edge assignment works; recurse for the next edge.
		pa.assignQueue = queue
		if pa.assign() {
			return true
		}
		pa.assignQueue = pa.assignQueue.prev
		// Store the best assignment to report.
		if len(pa.abs2ConPorts) > len(pa.maxAbs2ConPorts) {
			pa.maxAbs2ConPorts = make(map[*AbstractPort]*ConcretePort, len(pa.abs2ConPorts))
			for absPort, conPort := range pa.abs2ConPorts {
				pa.maxAbs2ConPorts[absPort] = conPort
			}
		}
		rewindState()
	}
	// None of these Edges work with the current state.
	return false
}

func (s *solver) assignEdges(abs2ConNode map[*AbstractNode]*ConcreteNode) map[*AbstractPort]*ConcretePort {
	if len(s.absPort2Node) == 0 {
		return map[*AbstractPort]*ConcretePort{}
	}
	// Generate all possible edges for this node assignment.
	abs2ConEdgeCombos := make(map[*AbstractEdge][]*ConcreteEdge)
	absPort2Edge := make(map[*AbstractPort]*AbstractEdge)
	for _, absSrcNode := range s.abstractGraph.Nodes {
		conSrcNode, ok := abs2ConNode[absSrcNode]
		if !ok {
			// This should never happen; there is a node in Abstract Graph that wasn't assigned.
			return nil
		}
		for _, absSrcPort := range absSrcNode.Ports {
			// Check attributes match, then check if matched ports link correctly.
			matchedConPorts := s.matchPorts(absSrcPort, conSrcNode.Ports)
			if len(matchedConPorts) == 0 {
				// No possible assignments for this port.
				return nil
			}
			absDstPort2Edge, ok := s.absPort2Port2Edge[absSrcPort]
			if !ok {
				// The abstract port is not part of an edge.
				// NOTE: Unfortunately, this causes some hackiness where we consider unlinked Port A
				// a different port from Port A linked to Port B.
				var conEdges []*ConcreteEdge
				for _, p := range matchedConPorts {
					conEdges = append(conEdges, &ConcreteEdge{p, nil})
				}
				e := &AbstractEdge{absSrcPort, nil}
				abs2ConEdgeCombos[e] = conEdges
				absPort2Edge[absSrcPort] = e
				continue
			}

			// The AbstractPort is part of an edge, so check the ConcretePorts link to the corresponding ConcreteNode.
			var absDstPort *AbstractPort
			var absEdge *AbstractEdge
			for absDstPort, absEdge = range absDstPort2Edge {
				break
			}
			conDstNode := abs2ConNode[s.absPort2Node[absDstPort]]
			for _, conSrcPort := range matchedConPorts {
				conDstPorts2Edge := s.conPort2Port2Edge[conSrcPort]
				for p, conEdge := range conDstPorts2Edge {
					if conDstNode.containsPort(p) && s.matchPort(absDstPort, p) {
						abs2ConEdgeCombos[absEdge] = append(abs2ConEdgeCombos[absEdge], conEdge)
						absPort2Edge[absSrcPort] = absEdge
						absPort2Edge[absDstPort] = absEdge
					}
				}
			}
			if len(abs2ConEdgeCombos[absEdge]) == 0 {
				return nil
			}
		}
	}
	var orderedEdges []*AbstractEdge
	for e := range abs2ConEdgeCombos {
		orderedEdges = append(orderedEdges, e)
	}
	sort.SliceStable(orderedEdges, func(i, j int) bool {
		ei, ej := orderedEdges[i], orderedEdges[j]
		eiWeight, ejWeight := len(ei.Src.Constraints), len(ej.Src.Constraints)
		if ei.Dst != nil {
			eiWeight += len(ei.Dst.Constraints)
		}
		if ej.Dst != nil {
			ejWeight += len(ej.Dst.Constraints)
		}
		return eiWeight > ejWeight
	})

	aq := &assignQueue{
		prev:  &assignQueue{},
		queue: []*AbstractEdge{},
	}
	pa := &portAssigner{
		abs2ConEdges:                 abs2ConEdgeCombos,
		absPort2Edge:                 absPort2Edge,
		deferredPortConstraints:      s.deferredPortConstraints,
		orderedAbsEdges:              orderedEdges,
		deferredUntilPortConstraints: orderedmap.NewOrderedMap[*AbstractPort, []deferredPortConstraint](),
		abs2ConPorts:                 make(map[*AbstractPort]*ConcretePort, len(s.absPort2Node)),
		assignedConPorts:             make(map[*ConcretePort]struct{}, len(s.absPort2Node)),
		assignQueue:                  aq,
		maxAbs2ConPorts:              make(map[*AbstractPort]*ConcretePort, len(s.absPort2Node)),
	}
	if ok := pa.assign(); !ok {
		if len(pa.maxAbs2ConPorts) > s.maxAssign.numPorts {
			s.maxAssign.assignment.Node2Node = abs2ConNode
			maxPorts := make(map[*AbstractPort]*ConcretePort, len(s.absPort2Node))
			for p := range s.absPort2Node {
				if cp, ok := pa.maxAbs2ConPorts[p]; ok {
					maxPorts[p] = cp
				} else {
					maxPorts[p] = nil
				}
			}
			s.maxAssign.assignment.Port2Port = maxPorts
			s.maxAssign.numPorts = len(pa.maxAbs2ConPorts)
		}
		return nil
	}
	return pa.abs2ConPorts
}

// Matching code.
// processConstraints iterates through the abstract graph and processes constraints.
func (s *solver) processConstraints() {
	for _, n := range s.abstractGraph.Nodes {
		n := n
		nc := orderedmap.NewOrderedMap[string, LeafConstraint]()
		var dnc []deferredNodeConstraint
		addNodeConstraint := func(k string, c NodeConstraint, n *AbstractNode) {
			switch v := c.(type) {
			case *sameAsNode:
				dnc = append(dnc, func(abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
					return v.match(k, n, abs2ConNode)
				})
				nc.Set(k, reAny)
			case LeafConstraint:
				nc.Set(k, v)
			default:
				log.Fatalf("Unrecognized constraint type %T for node matching", v)
			}
		}
		for k, c := range n.Constraints {
			if v, ok := c.(*andNode); ok {
				for _, c := range v.nlcs {
					addNodeConstraint(k, c, n)
				}
			} else {
				addNodeConstraint(k, c, n)
			}
		}
		s.nodeConstraints.Set(n, nc)
		s.deferredNodeConstraints.Set(n, dnc)
		for _, p := range n.Ports {
			p := p
			pc := orderedmap.NewOrderedMap[string, LeafConstraint]()
			var dpc []deferredPortConstraint
			addPortConstraint := func(k string, c PortConstraint, p *AbstractPort) {
				switch v := c.(type) {
				case *sameAsPort:
					dpc = append(dpc, func(abs2ConPort map[*AbstractPort]*ConcretePort) (bool, *AbstractPort, error) {
						if _, ok := abs2ConPort[v.p]; !ok {
							return false, v.p, errors.Errorf("port %q not assigned", v.p.Desc)
						}
						return v.match(k, p, abs2ConPort), nil, nil
					})
					pc.Set(k, reAny)
				case LeafConstraint:
					pc.Set(k, v)
				default:
					log.Fatalf("Unrecognized constraint type %T for port matching", v)
				}
			}
			for k, c := range p.Constraints {
				if v, ok := c.(*andPort); ok {
					for _, c := range v.plcs {
						addPortConstraint(k, c, p)
					}
				} else {
					addPortConstraint(k, c, p)
				}
			}
			s.portConstraints.Set(p, pc)
			s.deferredPortConstraints.Set(p, dpc)
		}
	}
}

// matchDeferredNodes checks whether the proposed node mapping fulfills constraints.
func (s *solver) matchDeferredNodes(abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
	for _, n := range s.deferredNodeConstraints.Keys() {
		dncs, ok := s.deferredNodeConstraints.Get(n)
		if !ok {
			log.Fatalf("Impossible: node %q not present in orderedmap", n.Desc)
		}
		for _, dnc := range dncs {
			if !dnc(abs2ConNode) {
				return false
			}
		}
	}
	return true
}

// matchNodes returns all ConcreteNodes in the ConcreteGraph that match the AbstractNode.
func (s *solver) matchNodes(abs *AbstractNode, superGraph *ConcreteGraph) []*ConcreteNode {
	match := func(n *ConcreteNode) bool {
		// Check that there are at least enough Ports to satisfy potential Edges.
		if len(abs.Ports) > len(n.Ports) {
			return false
		}
		// Check that there are at least enough Edges to other Nodes.
		var absEdges, conEdges []int
		for _, i := range s.absNode2Node2NumEdges[abs] {
			absEdges = append(absEdges, i)
		}

		for _, i := range s.conNode2Node2NumEdges[n] {
			conEdges = append(conEdges, i)
		}
		// Sort the number of edges to another Node in descending order.
		sort.Slice(absEdges, func(i, j int) bool { return absEdges[i] > absEdges[j] })
		sort.Slice(conEdges, func(i, j int) bool { return conEdges[i] > conEdges[j] })
		// Check the ConcreteNode has Edges to enough other Nodes.
		if len(conEdges) < len(absEdges) {
			return false
		}
		// Check if there are at least enough Edges other Nodes.
		for i, num := range absEdges {
			if conEdges[i] < num {
				return false
			}
		}

		constraints, ok := s.nodeConstraints.Get(abs)
		if !ok {
			return true
		}
		for _, k := range constraints.Keys() {
			c, ok := constraints.Get(k)
			if !ok {
				continue
			}
			s, ok := n.Attrs[k]
			if !c.match(s, ok) {
				return false
			}
		}
		return true
	}

	var nodes []*ConcreteNode
	for _, n := range superGraph.Nodes {
		if match(n) {
			nodes = append(nodes, n)
		}
	}
	sort.SliceStable(nodes, func(i, j int) bool { return nodes[i].Desc < nodes[j].Desc })
	return nodes
}

// matchDeferredPort checks the deferred constraints against the port.
func (s *solver) matchDeferredPort(port *AbstractPort, abs2ConPort map[*AbstractPort]*ConcretePort, attr string, constraint PortConstraint) bool {
	switch v := constraint.(type) {
	case *sameAsPort:
		if !constraint.(*sameAsPort).match(attr, port, abs2ConPort) {
			return false
		}
	default:
		log.Fatalf("Unrecognized constraint type %T for port matching", v)
	}
	return true

}

func (s *solver) matchPort(abs *AbstractPort, p *ConcretePort) bool {
	constraints, ok := s.portConstraints.Get(abs)
	if !ok {
		return true
	}
	for _, k := range constraints.Keys() {
		c, ok := constraints.Get(k)
		if !ok {
			log.Fatalf("Impossible: key %q not present in orderedmap", k)
		}
		s, ok := p.Attrs[k]
		if !c.match(s, ok) {
			return false
		}
	}
	return true
}

// matchPorts returns all ConcretePorts that match the AbstractPort.
func (s *solver) matchPorts(abs *AbstractPort, con []*ConcretePort) []*ConcretePort {
	var ports []*ConcretePort
	for _, p := range con {
		if s.matchPort(abs, p) {
			ports = append(ports, p)
		}
	}
	return ports
}
