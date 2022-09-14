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
		} else if _, ok := ret[e.Dst]; ok {
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
	Desc  string                // Description for the AbstractNode for logging.
	Ports []*AbstractPort       // A list AbstractPorts that may be connected to another AbstractPort.
	Attrs map[string]Constraint // A map of key value attributes of the AbstractNode.
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
	Desc  string                // Description for the AbstractPort for logging.
	Attrs map[string]Constraint // A map of key value attributes of the AbstractPort.
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

type solver struct {
	abstractGraph *AbstractGraph
	superGraph    *ConcreteGraph
	absPort2Node  map[*AbstractPort]*AbstractNode         // Map Port to Node it is part of.
	absNode2Node  map[*AbstractNode]map[*AbstractNode]int // Map Node to Node and how many links there are.
	conPort2Port  map[*ConcretePort][]*ConcretePort       // Cache the concrete ports that are linked.
	absPort2Port  map[*AbstractPort]*AbstractPort         // Cache of abstract ports that are linked.
}

// Solve returns as assignment from superGraph that satisfies abstractGraph.
func Solve(abstractGraph *AbstractGraph, superGraph *ConcreteGraph) (*Assignment, error) {
	if len(abstractGraph.Nodes) > len(superGraph.Nodes) {
		return nil, fmt.Errorf("not enough nodes in %q to satisfy %q", superGraph.Desc, abstractGraph.Desc)
	}
	if len(abstractGraph.Edges) > len(superGraph.Edges) {
		return nil, fmt.Errorf("not enough edges in %q to satisfy %q", superGraph.Desc, abstractGraph.Desc)
	}

	// Preprocess data for the solve.
	// Map the AbstractPort to AbstractNode.
	absPort2Node := make(map[*AbstractPort]*AbstractNode)
	for _, n := range abstractGraph.Nodes {
		for _, p := range n.Ports {
			absPort2Node[p] = n
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
		return nil, errors.Wrapf(err, "could not solve for %q from %q", abstractGraph.Desc, superGraph.Desc)
	}

	s := &solver{
		abstractGraph: abstractGraph,
		superGraph:    superGraph,
		absNode2Node:  absNode2Node,
		absPort2Node:  absPort2Node,
		absPort2Port:  absPort2Port,
	}

	a, err := s.solve()
	if err != nil {
		return nil, errors.Wrapf(err, "could not solve for %q from %q", abstractGraph.Desc, superGraph.Desc)
	}
	return a, nil
}

// solve provides a mapping of abstract nodes and ports to concrete nodes and ports.
// 1. Find all concrete nodes that match the abstract node (check attributes and # of ports).
// 2. For each mapping, check the concrete nodes and satisfy the edges of the abstract graph.
// 3. Assign concrete ports.
func (s *solver) solve() (*Assignment, error) {
	// TODO: Add a "report" that prints out info about the solve.
	abs2ConNodes := make(map[*AbstractNode][]*ConcreteNode)
	for _, n := range s.abstractGraph.Nodes {
		nodes := matchNodes(n, s.superGraph)
		if len(nodes) == 0 {
			// TODO: Migrate this to something like a "report" that prints out info about the solve.
			return nil, errors.Errorf("node %q could not be assigned to any node in %q", n.Desc, s.superGraph.Desc)
		}
		for _, conNode := range nodes {
			abs2ConNodes[n] = append(abs2ConNodes[n], conNode)
		}
	}

	s.conPort2Port = s.superGraph.fetchPort2PortMap()

	// Generate all AbstractNode -> ConcreteNode mappings.
	abs2ConNodeChan, stop := genCombos(abs2ConNodes)
	defer stop()
	// Iterate through each mapping. For each mapping, attempt to match edges and assign ports.
	for abs2ConNode := range abs2ConNodeChan {
		// For this mapping, check that the ConcreteNodes can satisfy the AbstractEdges.
		if !s.checkEdges(abs2ConNode) {
			continue
		}

		// Since the edges can be satisfied, try to assign matching ports.
		abs2ConPort, err := s.assignPorts(abs2ConNode, make(map[*AbstractPort]*ConcretePort), make(map[*AbstractNode]bool))
		if err != nil {
			return nil, err
		} else if abs2ConPort == nil {
			continue
		}

		return &Assignment{abs2ConNode, abs2ConPort}, nil
	}
	return nil, errors.Errorf("no assignment found for %q", s.abstractGraph.Desc)
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
func (s *solver) assignPorts(abs2ConNode map[*AbstractNode]*ConcreteNode, abs2ConPort map[*AbstractPort]*ConcretePort, assignedNodes map[*AbstractNode]bool) (map[*AbstractPort]*ConcretePort, error) {
	if len(s.absPort2Node) == len(abs2ConPort) {
		// Done.
		return abs2ConPort, nil
	}
	abs2ConPortCombos := make(map[*AbstractPort][]*ConcretePort)
	for absSrcNode, conSrcNode := range abs2ConNode {
		if _, ok := assignedNodes[absSrcNode]; ok {
			continue
		}
		for _, absSrcPort := range absSrcNode.Ports {
			// Check attributes match, then check if matched ports link correctly.
			matchedConPorts := matchPorts(absSrcPort, conSrcNode.Ports)
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
			canAssign := true
			for absPort, conPort := range port2Port {
				if absDstPort, ok := s.absPort2Port[absPort]; ok {
					if assignedDstPort, ok := abs2ConPort[absDstPort]; ok {
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
				// Try to assign the port.
				// If the part is part of an edge, check if the other port has been assigned. If it has, check that this port is linked to the dst port.
				abs2ConPort[absPort] = conPort
			}
			// Since the combo of Ports for this Node failed, try the next combo of Ports.
			if !canAssign {
				continue
			}
			assignedNodes[absSrcNode] = true
			if ret, err := s.assignPorts(abs2ConNode, abs2ConPort, assignedNodes); err != nil {
				return nil, err
			} else if ret != nil {
				return ret, nil
			}
			// Assignment of Ports for the next Node failed; unassign this Node.
			assignedNodes[absSrcNode] = false
		}
	}
	log.V(1).Infof("Only %d of %d ports were assigned; unassigning all ports", len(abs2ConPort), len(s.absPort2Node))
	return nil, nil
}

// Matching code.
// matchNodes returns all ConcreteNodes the ConcreteGraph that match the AbstractNode.
func matchNodes(abs *AbstractNode, superGraph *ConcreteGraph) []*ConcreteNode {
	match := func(n *ConcreteNode) bool {
		// Check that there are at least enough Ports to satisfy potential Edges.
		if len(abs.Ports) > len(n.Ports) {
			return false
		}
		for k, c := range abs.Attrs {
			if !c.match(n.Attrs[k]) {
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

// matchPorts returns all ConcretePorts that match the AbstractPort.
func matchPorts(abs *AbstractPort, con []*ConcretePort) []*ConcretePort {
	match := func(p *ConcretePort) bool {
		// for now, just do a equality match
		for k, c := range abs.Attrs {
			if !c.match(p.Attrs[k]) {
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
