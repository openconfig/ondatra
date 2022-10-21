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

	log "github.com/golang/glog"
	"github.com/pkg/errors"
)

// AbstractGraph is a representation of abstract nodes, ports, and edges.
type AbstractGraph struct {
	Desc  string // Description for the AbstractGraph for logging.
	Nodes []*AbstractNode
	Edges []*AbstractEdge
}

// fetchPort2PortMap() returns the mapping of an AbstractPort and any AbstractPorts it shares an AbstractEdge wtih.
func (g *AbstractGraph) fetchPort2PortMap() (map[*AbstractPort]*AbstractPort, error) {
	ret := make(map[*AbstractPort]*AbstractPort)
	for _, e := range g.Edges {
		if _, ok := ret[e.Src]; ok {
			return nil, fmt.Errorf("port %q is attached to more than one other port; can only be attached to one", e.Src.Desc)
		}
		if _, ok := ret[e.Dst]; ok {
			return nil, fmt.Errorf("port %q is attached to more than one other port; can only be attached to one", e.Dst.Desc)
		}
		ret[e.Src] = e.Dst
		ret[e.Dst] = e.Src
	}
	return ret, nil
}

// ConcreteGraph is a representation of concrete nodes, ports, and edges.
type ConcreteGraph struct {
	Desc  string // Description for the AbstractGraph for logging.
	Nodes []*ConcreteNode
	Edges []*ConcreteEdge
}

// fetchPort2PortMap() returns the mapping of a ConcretePort and any ConcretePorts it shares a ConcreteEdge wtih.
func (g *ConcreteGraph) fetchPort2PortMap() map[*ConcretePort][]*ConcretePort {
	ret := make(map[*ConcretePort][]*ConcretePort)
	for _, e := range g.Edges {
		ret[e.Src] = append(ret[e.Src], e.Dst)
		ret[e.Dst] = append(ret[e.Dst], e.Src)
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
	abstractGraph *AbstractGraph
	superGraph    *ConcreteGraph
	absPort2Node  map[*AbstractPort]*AbstractNode         // Map Port to Node it is part of.
	absNode2Node  map[*AbstractNode]map[*AbstractNode]int // Map Node to Node and how many links there are.
	conPort2Port  map[*ConcretePort][]*ConcretePort       // Cache the concrete ports that are linked.
	absPort2Port  map[*AbstractPort]*AbstractPort         // Cache of abstract ports that are linked.

	maxAssign *maxAssignment // The "best" Assignment for reporting if the solve failed.

	// Constraint mappings. Deferred constraint can only be checked after all abstract elements are assigned.
	nodeConstraints         map[*AbstractNode]map[string]LeafConstraint
	portConstraints         map[*AbstractPort]map[string]LeafConstraint
	deferredNodeConstraints map[*AbstractNode][]deferredNodeConstraint
	deferredPortConstraints map[*AbstractPort][]deferredPortConstraint
}

// Solve returns as assignment from superGraph that satisfies abstractGraph.
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
	absNode2Node := make(map[*AbstractNode]map[*AbstractNode]int)
	for _, e := range abstractGraph.Edges {
		srcNode, dstNode := e.Src, e.Dst
		if _, ok := absNode2Node[absPort2Node[srcNode]]; !ok {
			absNode2Node[absPort2Node[srcNode]] = make(map[*AbstractNode]int)
		}
		absNode2Node[absPort2Node[srcNode]][absPort2Node[dstNode]]++
	}

	absPort2Port, err := abstractGraph.fetchPort2PortMap()
	if err != nil {
		return nil, solveErr
	}

	s := &solver{
		abstractGraph:           abstractGraph,
		superGraph:              superGraph,
		absNode2Node:            absNode2Node,
		absPort2Node:            absPort2Node,
		absPort2Port:            absPort2Port,
		maxAssign:               &maxAssignment{&Assignment{Node2Node: node2Node, Port2Port: port2Port}, 0, 0},
		nodeConstraints:         make(map[*AbstractNode]map[string]LeafConstraint),
		deferredNodeConstraints: make(map[*AbstractNode][]deferredNodeConstraint),
		portConstraints:         make(map[*AbstractPort]map[string]LeafConstraint),
		deferredPortConstraints: make(map[*AbstractPort][]deferredPortConstraint),
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

	s.conPort2Port = s.superGraph.fetchPort2PortMap()

	// Generate all AbstractNode -> ConcreteNode mappings.
	abs2ConNodeChan, stop := genCombos(abs2ConNodes)
	defer stop()
	// Iterate through each mapping.
	// For each mapping, evaluate deferred constraints and attempt to match edges and assign ports.
	for abs2ConNode := range abs2ConNodeChan {
		if !s.matchDeferredNodes(abs2ConNode) {
			continue
		}
		// For this mapping, check that the ConcreteNodes can satisfy the AbstractEdges.
		if !s.checkEdges(abs2ConNode) {
			continue
		}
		if len(abs2ConNode) > s.maxAssign.numNodes {
			s.maxAssign.assignment.Node2Node = abs2ConNode
			s.maxAssign.numNodes = len(abs2ConNode)
		}

		// Since the edges can be satisfied, try to assign matching ports.
		abs2ConPort, _ := s.assignPorts(abs2ConNode, make(map[*AbstractPort]*ConcretePort), make(map[*AbstractNode]bool), make(map[*AbstractPort][]deferredPortConstraint))
		if abs2ConPort == nil {
			continue
		}

		return &Assignment{abs2ConNode, abs2ConPort}, true
	}
	return nil, false
}

// checkEdges validates that the concrete nodes can satisfy the abstract edges.
func (s *solver) checkEdges(abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
	// Iterate through the abstract edges.
	for absSrcNode, absDstNodes := range s.absNode2Node {
		for absDstNode, need := range absDstNodes {
			conSrcNode, conDstNode := abs2ConNode[absSrcNode], abs2ConNode[absDstNode]
			if !s.checkEdge(conSrcNode, conDstNode, need) {
				log.V(1).Infof("node %q cannot be assigned to %q because there are not enough edges to node %q assigned to %q; need %d",
					absSrcNode.Desc, conSrcNode.Desc, absDstNode.Desc, conDstNode.Desc, need)
				return false
			}
		}
	}
	return true
}

func (s *solver) checkEdge(conSrcNode, conDstNode *ConcreteNode, edgesNeeded int) bool {
	var count int
	// Verify whether the concrete Edges satisfy the abstract Edges.
	for _, conSrcPort := range conSrcNode.Ports {
		for _, p := range s.conPort2Port[conSrcPort] {
			if conDstNode.containsPort(p) {
				count++
				if count >= edgesNeeded {
					return true
				}
			}
		}
	}
	return false
}

// assignPorts tries to assign the ConcretePorts given the Node mapping.
// Returns the port mapping if the solve was successful and a boolean for whether to continue recursing with this Node mapping.
func (s *solver) assignPorts(abs2ConNode map[*AbstractNode]*ConcreteNode, abs2ConPort map[*AbstractPort]*ConcretePort, assignedNodes map[*AbstractNode]bool, deferredConstraints map[*AbstractPort][]deferredPortConstraint) (map[*AbstractPort]*ConcretePort, bool) {
	if len(s.absPort2Node) == len(abs2ConPort) {
		// Done.
		return abs2ConPort, true
	}
	abs2ConPortCombos := make(map[*AbstractPort][]*ConcretePort)
	for absSrcNode, conSrcNode := range abs2ConNode {
		if _, ok := assignedNodes[absSrcNode]; ok {
			continue
		}
		for _, absSrcPort := range absSrcNode.Ports {
			// Check attributes match, then check if matched ports link correctly.
			matchedConPorts := s.matchPorts(absSrcPort, conSrcNode.Ports)
			absDstPort, ok := s.absPort2Port[absSrcPort]
			if !ok {
				abs2ConPortCombos[absSrcPort] = matchedConPorts
				continue
			}
			// The AbstractPort is part of an edge, so check the ConcretePorts link to corresponding ConcreteNode.
			conDstNode := abs2ConNode[s.absPort2Node[absDstPort]]
			for _, conSrcPort := range matchedConPorts {
				conDstPorts := s.conPort2Port[conSrcPort]
				for _, p := range conDstPorts {
					if conDstNode.containsPort(p) {
						abs2ConPortCombos[absSrcPort] = append(abs2ConPortCombos[absSrcPort], conSrcPort)
					}
				}
			}
		}

		// Try the Port combos for this ConcreteNode.
		abs2ConPortChan, stop := genCombos(abs2ConPortCombos)
		defer stop()
		for port2Port := range abs2ConPortChan {
			deferredConstraintsCopy := make(map[*AbstractPort][]deferredPortConstraint)
			for n, cons := range deferredConstraints {
				deferredConstraintsCopy[n] = cons
			}
			abs2ConPortCopy := make(map[*AbstractPort]*ConcretePort)
			for a, c := range abs2ConPort {
				abs2ConPortCopy[a] = c
			}
			canAssign := true
			for absPort, conPort := range port2Port {
				if absDstPort, ok := s.absPort2Port[absPort]; ok {
					if assignedDstPort, ok := abs2ConPortCopy[absDstPort]; ok {
						isLinked := false
						for _, p := range s.conPort2Port[conPort] {
							if p == assignedDstPort {
								// This port is linked to an already assigned port; ok to assign.
								isLinked = true
								break
							}
						}
						if !isLinked {
							canAssign = false
						}
					}
				}
				// This port combo does not work with previously assigned ports.
				if !canAssign {
					break
				}
				// Try to assign the port, then evaluate deferred constraints.
				abs2ConPortCopy[absPort] = conPort

				// Start with constraints from previously assigned ports that depend on the port currently being assigned.
				if dcs, ok := deferredConstraintsCopy[absPort]; ok {
					for _, dc := range dcs {
						ok, _, err := dc(abs2ConPortCopy)
						if err != nil {
							// This constraint was already deferred from a previously assigned port.
							// An error means a port that should have been assigned is not.
							return nil, false
						}
						if !ok {
							canAssign = false
							break
						}
					}
				}
				if !canAssign {
					break
				}
				// Check deferred constraints.
				// If the constraint depends on another port that has not been assigned yet, defer the constraint.
				if constraints, ok := s.deferredPortConstraints[absPort]; ok {
					for _, dc := range constraints {
						if ok, p, err := dc(abs2ConPortCopy); err != nil {
							// At this stage, an error is fine. The check must be deferred until both ports are assigned.
							deferredConstraintsCopy[p] = append(deferredConstraintsCopy[p], dc)
						} else if !ok {
							canAssign = false
							break
						}
					}
					if !canAssign {
						break
					}
				}
			}
			// Since the combo of Ports for this Node failed, try the next combo of Ports.
			if !canAssign {
				continue
			}
			assignedNodes[absSrcNode] = true
			ret, ok := s.assignPorts(abs2ConNode, abs2ConPortCopy, assignedNodes, deferredConstraintsCopy)
			if !ok {
				if len(abs2ConPortCopy) > s.maxAssign.numPorts {
					s.maxAssign.assignment.Node2Node = abs2ConNode
					s.maxAssign.assignment.Port2Port = abs2ConPortCopy
					s.maxAssign.numPorts = len(abs2ConPortCopy)
				}
				return nil, false
			}
			if ret != nil {
				return ret, true
			}
			// Assignment of Ports for the next Node failed; unassign this Node.
			assignedNodes[absSrcNode] = false
		}
	}
	log.V(1).Infof("Only %d of %d ports were assigned; unassigning all ports", len(abs2ConPort), len(s.absPort2Node))
	if len(abs2ConPort) > s.maxAssign.numPorts {
		s.maxAssign.assignment.Node2Node = abs2ConNode
		s.maxAssign.assignment.Port2Port = abs2ConPort
		s.maxAssign.numPorts = len(abs2ConPort)
	}
	return nil, true
}

// Matching code.
// processConstraints iterates through the abstract graph and processes constraints.
func (s *solver) processConstraints() {
	for _, n := range s.abstractGraph.Nodes {
		n := n
		nc := make(map[string]LeafConstraint)
		dnc := []deferredNodeConstraint{}
		addNodeConstraint := func(k string, c NodeConstraint, n *AbstractNode) {
			switch v := c.(type) {
			case *sameAsNode:
				dnc = append(dnc, func(abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
					return v.match(k, n, abs2ConNode)
				})
			case *notSameAsNode:
				dnc = append(dnc, func(abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
					return v.match(k, n, abs2ConNode)
				})
			case LeafConstraint:
				nc[k] = v
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
		s.nodeConstraints[n] = nc
		s.deferredNodeConstraints[n] = dnc
		for _, p := range n.Ports {
			p := p
			pc := make(map[string]LeafConstraint)
			dpc := []deferredPortConstraint{}
			addPortConstraint := func(k string, c PortConstraint, p *AbstractPort) {
				switch v := c.(type) {
				case *sameAsPort:
					dpc = append(dpc, func(abs2ConPort map[*AbstractPort]*ConcretePort) (bool, *AbstractPort, error) {
						if _, ok := abs2ConPort[v.p]; !ok {
							return false, v.p, errors.Errorf("port %q not assigned", v.p.Desc)
						}
						return v.match(k, p, abs2ConPort), nil, nil
					})
				case *notSameAsPort:
					dpc = append(dpc, func(abs2ConPort map[*AbstractPort]*ConcretePort) (bool, *AbstractPort, error) {
						if _, ok := abs2ConPort[v.p]; !ok {
							return false, v.p, errors.Errorf("port %q not assigned", v.p.Desc)
						}
						return v.match(k, p, abs2ConPort), nil, nil
					})
				case LeafConstraint:
					pc[k] = v
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
			s.portConstraints[p] = pc
			s.deferredPortConstraints[p] = dpc
		}
	}
}

// matchDeferredNodes checks whether the proposed node mapping fulfills constraints.
func (s *solver) matchDeferredNodes(abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
	for _, dncs := range s.deferredNodeConstraints {
		for _, dnc := range dncs {
			if !dnc(abs2ConNode) {
				return false
			}
		}
	}
	return true
}

// matchNodes returns all ConcreteNodes the ConcreteGraph that match the AbstractNode.
func (s *solver) matchNodes(abs *AbstractNode, superGraph *ConcreteGraph) []*ConcreteNode {
	match := func(n *ConcreteNode) bool {
		// Check that there are at least enough Ports to satisfy potential Edges.
		if len(abs.Ports) > len(n.Ports) {
			return false
		}
		for k, c := range s.nodeConstraints[abs] {
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
	return nodes
}

// matchDeferredPort checks the deferred constraints against the port.
func (s *solver) matchDeferredPort(port *AbstractPort, abs2ConPort map[*AbstractPort]*ConcretePort, attr string, constraint PortConstraint) bool {
	switch v := constraint.(type) {
	case *sameAsPort:
		if !constraint.(*sameAsPort).match(attr, port, abs2ConPort) {
			return false
		}
	case *notSameAsPort:
		if !constraint.(*notSameAsPort).match(attr, port, abs2ConPort) {
			return false
		}
	default:
		log.Fatalf("Unrecognized constraint type %T for port matching", v)
	}
	return true

}

// matchPorts returns all ConcretePorts that match the AbstractPort.
func (s *solver) matchPorts(abs *AbstractPort, con []*ConcretePort) []*ConcretePort {
	match := func(p *ConcretePort) bool {
		for k, c := range s.portConstraints[abs] {
			s, ok := p.Attrs[k]
			if !c.match(s, ok) {
				return false
			}
		}
		return true
	}

	var ports []*ConcretePort
	for _, p := range con {
		if match(p) {
			ports = append(ports, p)
		}
	}
	return ports
}
