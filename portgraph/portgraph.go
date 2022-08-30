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

// Graph is a representation of nodes, ports, and edges.
type Graph struct {
	Desc  string // Description for the Graph for logging.
	Nodes []*Node
	Edges []*Edge
}

// fetchPort2PortMap() returns the mapping of a Port any Ports it shares and Edge wtih.
func (g *Graph) fetchPort2PortMap() map[*Port][]*Port {
	ret := make(map[*Port][]*Port)
	for _, e := range g.Edges {
		ret[e.Src] = append(ret[e.Src], e.Dst)
		ret[e.Dst] = append(ret[e.Dst], e.Src)
	}
	return ret
}

// Node represent a vertex on a graph.
// The Node has Ports and Edges that originate and terminate at Ports.
type Node struct {
	Desc  string            // Description for the Node for logging.
	Ports []*Port           // A list Ports that may be connected to another Port.
	Attrs map[string]string // A map of key value attributes of the Node.
}

func (n *Node) containsPort(p *Port) bool {
	for _, p2 := range n.Ports {
		if p == p2 {
			return true
		}
	}
	return false
}

// Edge represents a link from an source Port to a destination Port.
type Edge struct {
	Src, Dst *Port
}

// A Port is a is point on a Node where an Edge may be attached.
type Port struct {
	Desc  string            // Description for the Port for logging.
	Attrs map[string]string // A map of key value attributes of the Port.
}

// Assignment contains the Node -> Node and Port -> Port mappings.
type Assignment struct {
	Node2Node map[*Node]*Node
	Port2Port map[*Port]*Port
}

type solver struct {
	abstractGraph, superGraph *Graph

	absNode2Node map[*Node]map[*Node]int // Map Node to Node and how many links there are.
	conPort2Port map[*Port][]*Port       // Cache the concrete ports that are linked.
}

// Solve returns as assignment from superGraph that satisfies abstractGraph.
func Solve(abstractGraph *Graph, superGraph *Graph) (*Assignment, error) {
	if len(abstractGraph.Nodes) > len(superGraph.Nodes) {
		return nil, fmt.Errorf("not enough nodes in %q to satisfy %q", superGraph.Desc, abstractGraph.Desc)
	}
	if len(abstractGraph.Edges) > len(superGraph.Edges) {
		return nil, fmt.Errorf("not enough edges in %q to satisfy %q", superGraph.Desc, abstractGraph.Desc)
	}

	// Preprocess data for the solve.
	// Map the abstract Port to abstract Node.
	absPort2Node := make(map[*Port]*Node)
	for _, n := range abstractGraph.Nodes {
		for _, p := range n.Ports {
			absPort2Node[p] = n
		}
	}
	// Map how many links there are between each abstract node to calculate matches.
	absNode2Node := make(map[*Node]map[*Node]int)
	for _, e := range abstractGraph.Edges {
		srcNode, dstNode := e.Src, e.Dst
		if _, ok := absNode2Node[absPort2Node[srcNode]]; !ok {
			absNode2Node[absPort2Node[srcNode]] = make(map[*Node]int)
		}
		absNode2Node[absPort2Node[srcNode]][absPort2Node[dstNode]]++
	}

	s := &solver{
		abstractGraph: abstractGraph,
		superGraph:    superGraph,
		absNode2Node:  absNode2Node,
	}

	a, err := s.solve()
	if err != nil {
		return nil, errors.Wrapf(err, "could not solve for %q from %q", abstractGraph.Desc, superGraph.Desc)
	}
	return a, nil
}

// solve provides a mapping of an abstract graph to a concrete graph.
// 1. Find all concrete nodes that match the abstract node (check attributes and # of ports).
// 2. For each mapping, check the concrete nodes and satisfy the edges of the abstract graph.
// 3. Assign concrete ports.
func (s *solver) solve() (*Assignment, error) {
	// TODO: Add a "report" that prints out info about the solve.
	abs2ConNodes := make(map[*Node][]*Node)
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

	// Generate all abstract node -> concrete node mappings.
	abs2ConNodeChan := genNodeCombos(abs2ConNodes)
	// Iterate through each mapping. For each mapping, attempt to match edges and assign ports.
	for abs2ConNode := range abs2ConNodeChan {
		// For this mapping, check the concrete nodes can satisfy the abstract edges.
		if !s.checkEdges(abs2ConNode) {
			continue
		}

		// Since the edges can be satisfied, try to assign matching ports.
		s.assignPorts()
	}
	return nil, nil
}

// checkEdges validates that the concrete nodes can satisfy the abstract edges.
func (s *solver) checkEdges(abs2ConNode map[*Node]*Node) bool {
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

func (s *solver) checkEdge(conSrcNode, conDstNode *Node, edgesNeeded int) bool {
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

func (s *solver) assignPorts() {
	return
}

// genNodeCombos yields every key->value mapping, where no two keys map to the same
// value, given a map of keys to their possible values.
func genNodeCombos(m map[*Node][]*Node) <-chan map[*Node]*Node {
	var keys []*Node
	for k := range m {
		keys = append(keys, k)
	}
	ch := make(chan map[*Node]*Node)
	go func() {
		// TODO: End the goroutine elegantly when the solve is done.
		defer close(ch)
		genNodeRecurse(m, keys, make(map[*Node]*Node), make(map[*Node]bool), ch)
	}()
	return ch
}

func genNodeRecurse(
	m map[*Node][]*Node,
	keys []*Node,
	res map[*Node]*Node,
	used map[*Node]bool,
	ch chan<- map[*Node]*Node) {
	if len(keys) == 0 {
		copy := make(map[*Node]*Node)
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
			genNodeRecurse(m, keys[1:], res, used, ch)
			delete(used, i)
		}
	}
}

// genPortCombos yields every key->value mapping, where no two keys map to the same
// value, given a map of keys to their possible values.
func genPortCombos(m map[*Port][]*Port) <-chan map[*Port]*Port {
	var keys []*Port
	for k := range m {
		keys = append(keys, k)
	}
	ch := make(chan map[*Port]*Port)
	go func() {
		// TODO: End the goroutine elegantly when the solve is done.
		defer close(ch)
		genPortRecurse(m, keys, make(map[*Port]*Port), make(map[*Port]bool), ch)
	}()
	return ch
}

func genPortRecurse(
	m map[*Port][]*Port,
	keys []*Port,
	res map[*Port]*Port,
	used map[*Port]bool,
	ch chan<- map[*Port]*Port) {
	if len(keys) == 0 {
		copy := make(map[*Port]*Port)
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
			genPortRecurse(m, keys[1:], res, used, ch)
			delete(used, i)
		}
	}
}

// Matching code.
// matchNodes returns all nodes in topo that match the abstract Node.
func matchNodes(abs *Node, superGraph *Graph) []*Node {
	match := func(n *Node) bool {
		// Check that there are at least enough Ports to satisfy potential Edges.
		if len(abs.Ports) > len(n.Ports) {
			return false
		}
		// for now, just do an equality match
		for k, v := range abs.Attrs {
			if v != n.Attrs[k] {
				return false
			}
		}
		return true
	}

	var nodes []*Node
	for _, n := range superGraph.Nodes {
		if match(n) {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

// matchPorts returns all ports that match the absPort.
func matchPorts(absPort *Port, conPorts []*Port) []*Port {
	match := func(p *Port) bool {
		// for now, just do a equality match
		for k, v := range absPort.Attrs {
			if v != p.Attrs[k] {
				return false
			}
		}
		return true
	}

	var ports []*Port
	for _, p := range conPorts {
		if match(p) {
			ports = append(ports, p)
		}
	}
	return ports
}
