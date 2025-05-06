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

package portgraph

import (
	"fmt"
	"sort"
)

type genCombosOptions struct {
	m                     map[*AbstractNode][]*ConcreteNode
	absNode2Node2NumEdges map[*AbstractNode]map[*AbstractNode]int
	checkEdge             func(*ConcreteNode, *ConcreteNode, int) (bool, []*ConcretePort, []*ConcretePort)
	solveReport           *solveReport
}

// genNodeCombos yields all key->value assignment maps, given the specified range
// of possible values for each key, such that no two keys are assigned the same
// value. Returns a channel that yields the assignments and a function to stop
// the generation. Caller should immediately defer a call to the stop function
// to ensure the routine generating the assignments exits.
func genNodeCombos(opts genCombosOptions) (<-chan map[*AbstractNode]*ConcreteNode, func()) {
	genComboEvent := opts.solveReport.AddPhase("Generating node combinations without port constraints")
	gen := &nodeGenerator{
		m:                      opts.m,
		res:                    make(map[*AbstractNode]*ConcreteNode),
		used:                   make(map[*ConcreteNode]bool),
		ch:                     make(chan map[*AbstractNode]*ConcreteNode),
		stop:                   make(chan struct{}, 1),
		done:                   make(chan struct{}, 1),
		absNode2Node2NumEdges:  opts.absNode2Node2NumEdges,
		checkEdge:              opts.checkEdge,
		conNode2Node2NeedPorts: make(map[*ConcreteNode]map[*ConcreteNode]*needPorts),
		solveReport:            opts.solveReport,
	}
	var keys []*AbstractNode
	for k := range opts.m {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		iWeight, jWeight := 0, 0
		for _, w := range opts.absNode2Node2NumEdges[keys[i]] {
			iWeight += w
		}
		for _, w := range opts.absNode2Node2NumEdges[keys[j]] {
			jWeight += w
		}
		// If weight is the same, sort by Desc (alphabetically).
		if iWeight == jWeight {
			return keys[i].Desc < keys[j].Desc
		}
		return iWeight > jWeight
	})
	go func() {
		defer close(gen.ch)
		gen.recurse(keys, genComboEvent)
		gen.done <- struct{}{}
	}()
	return gen.ch, gen.stopAndWait
}

type nodeGenerator struct {
	m    map[*AbstractNode][]*ConcreteNode
	res  map[*AbstractNode]*ConcreteNode
	used map[*ConcreteNode]bool
	ch   chan map[*AbstractNode]*ConcreteNode
	stop chan struct{}
	done chan struct{}

	absNode2Node2NumEdges  map[*AbstractNode]map[*AbstractNode]int
	checkEdge              func(*ConcreteNode, *ConcreteNode, int) (bool, []*ConcretePort, []*ConcretePort)
	conNode2Node2NeedPorts map[*ConcreteNode]map[*ConcreteNode]*needPorts
	solveReport            *solveReport
}

// needPorts stores how many ports are needed and which ConcretePorts can fulfill the needed ports.
type needPorts struct {
	need       int             // Total number of ports needed.
	canFulfill []*ConcretePort // The ports that can fulfill the needed ports.
}

// checkEnoughPorts checks there are enough unique ports to fulfill all needs.
func checkEnoughPorts(nps []*needPorts) (bool, string) {
	var needTotal int
	uniquePorts := make(map[*ConcretePort]struct{})
	for _, np := range nps {
		needTotal += np.need
		for _, p := range np.canFulfill {
			uniquePorts[p] = struct{}{}
		}
	}
	if needTotal > len(uniquePorts) {
		return false, fmt.Sprintf("got %d unique ports, want %d unique ports", len(uniquePorts), needTotal)
	}
	return true, ""
}

// stopAndWait signals the generation to stop and waits for it to stop.
func (g *nodeGenerator) stopAndWait() {
	g.stop <- struct{}{}
	<-g.done
}

// recurse generates all the map combinations for the specified set of nodes
// and checks that the set of nodes have appropriate edges, returning true
// until it receives a stop signal.
func (g *nodeGenerator) recurse(nodes []*AbstractNode, precedingEvent *Event) bool {
	if len(nodes) == 0 {
		copy := make(map[*AbstractNode]*ConcreteNode, len(g.m))
		for k, v := range g.res {
			copy[k] = v
		}
		select {
		case <-g.stop:
			return false
		case g.ch <- copy:
			return true
		}
	}
	first := nodes[0]
	for _, n := range g.m[first] {
		if g.used[n] {
			continue
		}
		g.conNode2Node2NeedPorts[n] = make(map[*ConcreteNode]*needPorts)
		canUse := true
		for dstAbsNode, dstConNode := range g.res {
			// Check if first has enough Edges to already allocated Nodes.
			if numEdges, ok := g.absNode2Node2NumEdges[first][dstAbsNode]; ok && numEdges > 0 {
				ok, srcPorts, dstPorts := g.checkEdge(n, dstConNode, numEdges)
				if !ok {
					// There are no Edges, so we can't use this Node.
					g.solveReport.AppendComboFailed(first.Desc, n.Desc, dstAbsNode.Desc, dstConNode.Desc, fmt.Sprintf("there are %d edges but %d are required", len(srcPorts), numEdges), precedingEvent)
					canUse = false
					break
				}
				// Store how many ports are needed to fulfill the Edges for allocated Nodes from both sides.
				srcNeedPorts := &needPorts{numEdges, srcPorts}
				dstNeedPorts := &needPorts{numEdges, dstPorts}
				// Check that we can still fulfill all nodes with this allocation.
				srcNeedPortsList := []*needPorts{srcNeedPorts}
				dstNeedPortsList := []*needPorts{dstNeedPorts}
				for _, need := range g.conNode2Node2NeedPorts[n] {
					srcNeedPortsList = append(srcNeedPortsList, need)
				}
				if srcOk, srcReason := checkEnoughPorts(srcNeedPortsList); !srcOk {
					g.solveReport.AppendComboFailed(first.Desc, n.Desc, dstAbsNode.Desc, dstConNode.Desc, srcReason+" to satisfy edges for "+n.Desc, precedingEvent)
					canUse = false
					break
				}
				for _, need := range g.conNode2Node2NeedPorts[dstConNode] {
					dstNeedPortsList = append(dstNeedPortsList, need)
				}
				if dstOk, dstReason := checkEnoughPorts(dstNeedPortsList); !dstOk {
					g.solveReport.AppendComboFailed(first.Desc, n.Desc, dstAbsNode.Desc, dstConNode.Desc, dstReason+" to satisfy edge for "+dstConNode.Desc, precedingEvent)
					canUse = false
					break
				}
				g.conNode2Node2NeedPorts[n][dstConNode] = srcNeedPorts
				g.conNode2Node2NeedPorts[dstConNode][n] = dstNeedPorts
			}
			g.solveReport.AppendValidCombo(first.Desc, n.Desc, dstAbsNode.Desc, dstConNode.Desc, precedingEvent)
		}
		if canUse {
			g.res[first] = n
			g.used[n] = true
			if !g.recurse(nodes[1:], precedingEvent) {
				return false
			}
			delete(g.used, n)
			delete(g.res, first)
		}
		for k := range g.conNode2Node2NeedPorts[n] {
			delete(g.conNode2Node2NeedPorts[k], n)
		}
		delete(g.conNode2Node2NeedPorts, n)
	}
	return true
}
