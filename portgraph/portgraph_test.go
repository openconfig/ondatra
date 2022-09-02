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
	"testing"
)

// Initialize ConcretePorts and ConcreteNodes for the super graph.
// These are the ConcreteNodes and ConcretePorts that should be matched to in tests.
var (
	node1port1   = &ConcretePort{Desc: "node1:port1"}
	node1port2   = &ConcretePort{Desc: "node1:port2"}
	node2port1   = &ConcretePort{Desc: "node2:port1"}
	node2port2   = &ConcretePort{Desc: "node2:port2"}
	node3port1   = &ConcretePort{Desc: "node3:port1"}
	node3port2   = &ConcretePort{Desc: "node3:port2"}
	node4port1   = &ConcretePort{Desc: "node4:port1"}
	node4port2   = &ConcretePort{Desc: "node4:port2"}
	node8port1   = &ConcretePort{Desc: "node8:port1"}
	node9port1   = &ConcretePort{Desc: "node9:port1"}
	node11port1  = &ConcretePort{Desc: "node11:port1"}
	node11port2  = &ConcretePort{Desc: "node11:port2"}
	node11port3  = &ConcretePort{Desc: "node11:port3"}
	node12port1  = &ConcretePort{Desc: "node12:port1"}
	node12port2  = &ConcretePort{Desc: "node12:port2"}
	node13port1  = &ConcretePort{Desc: "node13:port1"}
	node13port2  = &ConcretePort{Desc: "node13:port2"}
	node14port1  = &ConcretePort{Desc: "node14:port1"}
	node14port2  = &ConcretePort{Desc: "node14:port2"}
	node20port1  = &ConcretePort{Desc: "node20:port1", Attrs: map[string]string{"attr": "FOO"}}
	node20port2  = &ConcretePort{Desc: "node20:port2", Attrs: map[string]string{"attr": "BAR"}}
	node21port1  = &ConcretePort{Desc: "node21:port1", Attrs: map[string]string{"attr": "FOO"}}
	node21port2  = &ConcretePort{Desc: "node21:port2", Attrs: map[string]string{"attr": "BAR"}}
	node22port1  = &ConcretePort{Desc: "node22:port1", Attrs: map[string]string{"attr": "A"}}
	node22port2  = &ConcretePort{Desc: "node22:port2", Attrs: map[string]string{"attr": "B"}}
	node22port3  = &ConcretePort{Desc: "node23:port1", Attrs: map[string]string{"attr": "C"}}
	node22port4  = &ConcretePort{Desc: "node23:port2", Attrs: map[string]string{"attr": "D"}}
	node23port1  = &ConcretePort{Desc: "node23:port1", Attrs: map[string]string{"attr": "A2"}}
	node23port2  = &ConcretePort{Desc: "node23:port2", Attrs: map[string]string{"attr": "B2"}}
	node23port3  = &ConcretePort{Desc: "node23:port3", Attrs: map[string]string{"attr": "C2"}}
	node23port4  = &ConcretePort{Desc: "node24:port4", Attrs: map[string]string{"attr": "D2"}}
	node60port1  = &ConcretePort{Desc: "node60:port1", Attrs: map[string]string{"speed": "S_100G"}}
	node60port2  = &ConcretePort{Desc: "node60:port2", Attrs: map[string]string{"speed": "S_101G"}}
	node60port3  = &ConcretePort{Desc: "node60:port3", Attrs: map[string]string{"speed": "S_102G"}}
	node60port4  = &ConcretePort{Desc: "node60:port4", Attrs: map[string]string{"speed": "S_103G"}}
	node60port5  = &ConcretePort{Desc: "node60:port5", Attrs: map[string]string{"speed": "S_104G"}}
	node60port6  = &ConcretePort{Desc: "node60:port6", Attrs: map[string]string{"speed": "S_105G"}}
	node60port7  = &ConcretePort{Desc: "node60:port7", Attrs: map[string]string{"speed": "S_106G"}}
	node60port8  = &ConcretePort{Desc: "node60:port8", Attrs: map[string]string{"speed": "S_107G"}}
	node61port1  = &ConcretePort{Desc: "node61:port1", Attrs: map[string]string{"speed": "S_100G"}}
	node61port2  = &ConcretePort{Desc: "node61:port2", Attrs: map[string]string{"speed": "S_101G"}}
	node61port3  = &ConcretePort{Desc: "node61:port3", Attrs: map[string]string{"speed": "S_102G"}}
	node61port4  = &ConcretePort{Desc: "node61:port4", Attrs: map[string]string{"speed": "S_103G"}}
	node61port5  = &ConcretePort{Desc: "node61:port5", Attrs: map[string]string{"speed": "S_104G"}}
	node61port6  = &ConcretePort{Desc: "node61:port6", Attrs: map[string]string{"speed": "S_105G"}}
	node61port7  = &ConcretePort{Desc: "node61:port7", Attrs: map[string]string{"speed": "S_106G"}}
	node61port8  = &ConcretePort{Desc: "node61:port8", Attrs: map[string]string{"speed": "S_107G"}}
	node61port9  = &ConcretePort{Desc: "node61:port9", Attrs: map[string]string{"attr": "FOO"}}
	node61port10 = &ConcretePort{Desc: "node61:port10", Attrs: map[string]string{"attr": "BAR"}}

	node1  = &ConcreteNode{Desc: "node1", Ports: []*ConcretePort{node1port1, node1port2}, Attrs: map[string]string{"vendor": "CISCO1"}}
	node2  = &ConcreteNode{Desc: "node2", Ports: []*ConcretePort{node2port1, node2port2}, Attrs: map[string]string{"vendor": "CISCO2"}}
	node3  = &ConcreteNode{Desc: "node3", Ports: []*ConcretePort{node3port1, node3port2}, Attrs: map[string]string{"vendor": "CISCO3"}}
	node4  = &ConcreteNode{Desc: "node4", Ports: []*ConcretePort{node4port1, node4port2}, Attrs: map[string]string{"vendor": "CISCO4"}}
	node5  = &ConcreteNode{Desc: "node5", Attrs: map[string]string{"vendor": "CISCO1"}}
	node6  = &ConcreteNode{Desc: "node6", Attrs: map[string]string{"vendor": "CISCO2"}}
	node7  = &ConcreteNode{Desc: "node7", Attrs: map[string]string{"vendor": "CISCO3"}}
	node8  = &ConcreteNode{Desc: "node8", Ports: []*ConcretePort{node8port1}, Attrs: map[string]string{"vendor": "UNIQUE8"}}
	node9  = &ConcreteNode{Desc: "node9", Ports: []*ConcretePort{node9port1}, Attrs: map[string]string{"vendor": "UNIQUE9"}}
	node10 = &ConcreteNode{Desc: "node10", Attrs: map[string]string{"vendor": "CISCO1"}}
	node11 = &ConcreteNode{Desc: "node11", Ports: []*ConcretePort{node11port1, node11port2, node11port3}, Attrs: map[string]string{"vendor": "CISCO1"}}
	node12 = &ConcreteNode{Desc: "node12", Ports: []*ConcretePort{node12port1, node12port2}, Attrs: map[string]string{"vendor": "CISCO2"}}
	node13 = &ConcreteNode{Desc: "node13", Ports: []*ConcretePort{node13port1, node13port2}, Attrs: map[string]string{"vendor": "CISCO3"}}
	node14 = &ConcreteNode{Desc: "node14", Ports: []*ConcretePort{node14port1, node14port2}, Attrs: map[string]string{"vendor": "CISCO4"}}
	node20 = &ConcreteNode{Desc: "node20", Ports: []*ConcretePort{node20port1, node20port2}, Attrs: map[string]string{"attr": "multi1"}}
	node21 = &ConcreteNode{Desc: "node21", Ports: []*ConcretePort{node21port1, node21port2}, Attrs: map[string]string{"attr": "multi2"}}
	node22 = &ConcreteNode{Desc: "node22", Ports: []*ConcretePort{node22port1, node22port2, node22port3, node22port4}}
	node23 = &ConcreteNode{Desc: "node23", Ports: []*ConcretePort{node23port1, node23port2, node23port3, node23port4}}
	node60 = &ConcreteNode{Desc: "node60", Ports: []*ConcretePort{node60port1, node60port2, node60port3, node60port4, node60port5, node60port6, node60port7, node60port8}, Attrs: map[string]string{"vendor": "CISCO1"}}
	node61 = &ConcreteNode{Desc: "node61", Ports: []*ConcretePort{node61port1, node61port2, node61port3, node61port4, node61port5, node61port6, node61port7, node61port8, node61port9, node61port10}, Attrs: map[string]string{"vendor": "CISCO1"}}
)

var superGraph = &ConcreteGraph{
	Desc: "super",
	Nodes: []*ConcreteNode{
		node1, node2, node3, node4, node5, node6, node7, node8, node9, node10, node11, node12, node13, node14, node20, node21, node22, node23, node60, node61,
		// Additional entries for benchmark purposes.
		{Desc: "node15", Attrs: map[string]string{"vendor": "CISCO2"}},
		{Desc: "node16", Attrs: map[string]string{"vendor": "CISCO3"}},
		{Desc: "node17", Attrs: map[string]string{"vendor": "CISCO4"}},
		{Desc: "node18", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node19", Attrs: map[string]string{"vendor": "CISCO2"}},
		{Desc: "node24", Attrs: map[string]string{"vendor": "CISCO3"}},
		{Desc: "node25", Attrs: map[string]string{"vendor": "CISCO4"}},
		{Desc: "node26", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node27", Attrs: map[string]string{"vendor": "CISCO2"}},
		{Desc: "node28", Attrs: map[string]string{"vendor": "CISCO3"}},
		{Desc: "node29", Attrs: map[string]string{"vendor": "CISCO4"}},
		{Desc: "node30", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node31", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node32", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node33", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node34", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node35", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node36", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node37", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node38", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node39", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node40", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node41", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node42", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node43", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node44", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node45", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node46", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node47", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node48", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node49", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node50", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node51", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node52", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node53", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node54", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node55", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node56", Attrs: map[string]string{"vendor": "CISCO1"}},
	},
	Edges: []*ConcreteEdge{
		{Src: node1port1, Dst: node2port1},
		{Src: node1port2, Dst: node4port1},
		{Src: node2port2, Dst: node3port1},
		{Src: node3port2, Dst: node4port2},
		{Src: node8port1, Dst: node9port1},
		{Src: node11port1, Dst: node12port1},
		{Src: node11port2, Dst: node14port1},
		{Src: node12port2, Dst: node13port1},
		{Src: node13port2, Dst: node14port2},
		{Src: node20port1, Dst: node21port1}, // Simulate OCS; multiple links from port
		{Src: node20port1, Dst: node21port2},
		{Src: node20port2, Dst: node21port1},
		{Src: node22port1, Dst: node23port1},
		{Src: node22port1, Dst: node23port2},
		{Src: node22port2, Dst: node23port2},
		{Src: node22port2, Dst: node23port3},
		{Src: node22port3, Dst: node23port3},
		{Src: node22port3, Dst: node23port4},
		{Src: node22port4, Dst: node23port4},
		{Src: node22port4, Dst: node23port1},
		{Src: node60port1, Dst: node61port1},
		{Src: node60port2, Dst: node61port2},
		{Src: node60port3, Dst: node61port3},
		{Src: node60port4, Dst: node61port4},
		{Src: node60port5, Dst: node61port5},
		{Src: node60port6, Dst: node61port6},
		{Src: node60port7, Dst: node61port7},
		{Src: node60port8, Dst: node61port8},
	},
}

// Setup abstract Nodes and Ports for testing.
var (
	// One Node
	abs1 = &AbstractNode{Desc: "one node", Attrs: map[string]Constraint{"vendor": Equal("UNIQUE8")}}

	// One Node, one Port
	abs2      = &AbstractNode{Desc: "one node, one port", Ports: []*AbstractPort{abs2port1}, Attrs: map[string]Constraint{"vendor": Equal("UNIQUE8")}}
	abs2port1 = &AbstractPort{Desc: "one port"}

	// Four Nodes, interconnected
	abs3      = &AbstractNode{Desc: "abs3", Ports: []*AbstractPort{abs3port1, abs3port2, abs3port3}, Attrs: map[string]Constraint{"vendor": Equal("CISCO1")}}
	abs4      = &AbstractNode{Desc: "abs4", Ports: []*AbstractPort{abs4port1, abs4port2}, Attrs: map[string]Constraint{"vendor": Equal("CISCO2")}}
	abs5      = &AbstractNode{Desc: "abs5", Ports: []*AbstractPort{abs5port1, abs5port2}, Attrs: map[string]Constraint{"vendor": Equal("CISCO3")}}
	abs6      = &AbstractNode{Desc: "abs6", Ports: []*AbstractPort{abs6port1, abs6port2}, Attrs: map[string]Constraint{"vendor": Equal("CISCO4")}}
	abs3port1 = &AbstractPort{Desc: "abs3port1"}
	abs3port2 = &AbstractPort{Desc: "abs3port2"}
	abs3port3 = &AbstractPort{Desc: "abs3port3"}
	abs4port1 = &AbstractPort{Desc: "abs4port1"}
	abs4port2 = &AbstractPort{Desc: "abs4port2"}
	abs5port1 = &AbstractPort{Desc: "abs5port1"}
	abs5port2 = &AbstractPort{Desc: "abs5port2"}
	abs6port1 = &AbstractPort{Desc: "abs6port1"}
	abs6port2 = &AbstractPort{Desc: "abs6port2"}

	// One Node, 8 Ports
	abs7      = &AbstractNode{Desc: "abs3", Ports: []*AbstractPort{abs7port1, abs7port2, abs7port3, abs7port4, abs7port5, abs7port6, abs7port7, abs7port8}}
	abs7port1 = &AbstractPort{Desc: "abs7port1", Attrs: map[string]Constraint{"speed": Equal("S_100G")}}
	abs7port2 = &AbstractPort{Desc: "abs7port2", Attrs: map[string]Constraint{"speed": Equal("S_101G")}}
	abs7port3 = &AbstractPort{Desc: "abs7port3", Attrs: map[string]Constraint{"speed": Equal("S_102G")}}
	abs7port4 = &AbstractPort{Desc: "abs7port4", Attrs: map[string]Constraint{"speed": Equal("S_103G")}}
	abs7port5 = &AbstractPort{Desc: "abs7port5", Attrs: map[string]Constraint{"speed": Equal("S_104G")}}
	abs7port6 = &AbstractPort{Desc: "abs7port6", Attrs: map[string]Constraint{"speed": Equal("S_105G")}}
	abs7port7 = &AbstractPort{Desc: "abs7port7", Attrs: map[string]Constraint{"speed": Equal("S_106G")}}
	abs7port8 = &AbstractPort{Desc: "abs7port8", Attrs: map[string]Constraint{"speed": Equal("S_107G")}}

	// Two Nodes, Many Links
	abs8       = &AbstractNode{Desc: "abs8", Ports: []*AbstractPort{abs8port1, abs8port2, abs8port3, abs8port4, abs8port5, abs8port6, abs8port9, abs8port10}}
	abs9       = &AbstractNode{Desc: "abs9", Ports: []*AbstractPort{abs9port1, abs9port2, abs9port3, abs9port4}}
	abs8port1  = &AbstractPort{Desc: "abs8port1", Attrs: map[string]Constraint{"speed": Equal("S_100G")}}
	abs8port2  = &AbstractPort{Desc: "abs8port2", Attrs: map[string]Constraint{"speed": Equal("S_101G")}}
	abs8port3  = &AbstractPort{Desc: "abs8port3", Attrs: map[string]Constraint{"speed": Equal("S_102G")}}
	abs8port4  = &AbstractPort{Desc: "abs8port4", Attrs: map[string]Constraint{"speed": Equal("S_103G")}}
	abs8port5  = &AbstractPort{Desc: "abs8port5", Attrs: map[string]Constraint{"speed": Equal("S_104G")}}
	abs8port6  = &AbstractPort{Desc: "abs8port6", Attrs: map[string]Constraint{"speed": Equal("S_105G")}}
	abs8port9  = &AbstractPort{Desc: "abs8port9", Attrs: map[string]Constraint{"attr": Equal("FOO")}}
	abs8port10 = &AbstractPort{Desc: "abs8port10", Attrs: map[string]Constraint{"attr": Equal("BAR")}}
	abs9port1  = &AbstractPort{Desc: "abs9port1", Attrs: map[string]Constraint{"speed": Equal("S_100G")}}
	abs9port2  = &AbstractPort{Desc: "abs9port2", Attrs: map[string]Constraint{"speed": Equal("S_101G")}}
	abs9port3  = &AbstractPort{Desc: "abs9port3", Attrs: map[string]Constraint{"speed": Equal("S_102G")}}
	abs9port4  = &AbstractPort{Desc: "abs9port4", Attrs: map[string]Constraint{"speed": Equal("S_103G")}}

	// Two Nodes, links via shared ports
	abs10      = &AbstractNode{Desc: "abs10", Ports: []*AbstractPort{abs10port1, abs10port2}, Attrs: map[string]Constraint{"attr": Equal("multi1")}}
	abs11      = &AbstractNode{Desc: "abs11", Ports: []*AbstractPort{abs11port1, abs11port2}}
	abs10port1 = &AbstractPort{Desc: "abs10port1", Attrs: map[string]Constraint{"attr": Equal("FOO")}}
	abs10port2 = &AbstractPort{Desc: "abs10port2", Attrs: map[string]Constraint{"attr": Equal("BAR")}}
	abs11port1 = &AbstractPort{Desc: "abs11port1"}
	abs11port2 = &AbstractPort{Desc: "abs11port2"}

	// 2 nodes, 4 interconnected ports
	abs12      = &AbstractNode{Desc: "abs12", Ports: []*AbstractPort{abs12port1, abs12port2, abs12port3, abs12port4}}
	abs13      = &AbstractNode{Desc: "abs13", Ports: []*AbstractPort{abs13port1, abs13port2, abs13port3, abs13port4}}
	abs12port1 = &AbstractPort{Desc: "abs12port1"}
	abs12port2 = &AbstractPort{Desc: "abs12port2", Attrs: map[string]Constraint{"attr": Equal("B")}}
	abs12port3 = &AbstractPort{Desc: "abs12port3"}
	abs12port4 = &AbstractPort{Desc: "abs12port4", Attrs: map[string]Constraint{"attr": Equal("D")}}
	abs13port1 = &AbstractPort{Desc: "abs13port1", Attrs: map[string]Constraint{"attr": Equal("A2")}}
	abs13port2 = &AbstractPort{Desc: "abs13port2"}
	abs13port3 = &AbstractPort{Desc: "abs13port3", Attrs: map[string]Constraint{"attr": Equal("C2")}}
	abs13port4 = &AbstractPort{Desc: "abs13port4"}
)

func TestSolve(t *testing.T) {
	tests := []struct {
		desc      string
		graph     *AbstractGraph
		wantNodes map[*AbstractNode]*ConcreteNode
		wantPorts map[*AbstractPort]*ConcretePort
	}{{
		desc: "one node",
		graph: &AbstractGraph{
			Desc:  "one node",
			Nodes: []*AbstractNode{abs1},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs1: node8},
	}, {
		desc: "one node, one port",
		graph: &AbstractGraph{
			Desc:  "one node, one port",
			Nodes: []*AbstractNode{abs2},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs2: node8},
		wantPorts: map[*AbstractPort]*ConcretePort{abs2port1: node8port1},
	}, {
		desc: "four nodes, interconnected",
		graph: &AbstractGraph{
			Desc:  "four nodes, interconnected",
			Nodes: []*AbstractNode{abs3, abs4, abs5, abs6},
			Edges: []*AbstractEdge{{abs3port1, abs4port1}, {abs3port2, abs6port1}, {abs4port2, abs5port1}, {abs5port2, abs6port2}},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs3: node11, abs4: node12, abs5: node13, abs6: node14},
		wantPorts: map[*AbstractPort]*ConcretePort{
			abs3port1: node11port1,
			abs3port2: node11port2,
			abs3port3: node11port3,
			abs4port1: node12port1,
			abs4port2: node12port2,
			abs5port1: node13port1,
			abs5port2: node13port2,
			abs6port1: node14port1,
			abs6port2: node14port2,
		},
	}, {
		desc: "10 unlinked ports from linked ports",
		graph: &AbstractGraph{
			Desc:  "10 unlinked ports from linked ports",
			Nodes: []*AbstractNode{abs7},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs7: node60},
		wantPorts: map[*AbstractPort]*ConcretePort{
			abs7port1: node60port1,
			abs7port2: node60port2,
			abs7port3: node60port3,
			abs7port4: node60port4,
			abs7port5: node60port5,
			abs7port6: node60port6,
			abs7port7: node60port7,
			abs7port8: node60port8,
		},
	}, {
		desc: "2 nodes, 4 linked ports, 4 unlinked ports",
		graph: &AbstractGraph{
			Desc:  "2 nodes, 4 linked ports, 4 unlinked ports",
			Nodes: []*AbstractNode{abs8, abs9},
			Edges: []*AbstractEdge{{abs8port1, abs9port1}, {abs8port2, abs9port2}, {abs8port3, abs9port3}, {abs8port4, abs9port4}},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs8: node61, abs9: node60},
		wantPorts: map[*AbstractPort]*ConcretePort{
			abs8port1:  node61port1,
			abs8port2:  node61port2,
			abs8port3:  node61port3,
			abs8port4:  node61port4,
			abs8port5:  node61port5,
			abs8port6:  node61port6,
			abs8port9:  node61port9,
			abs8port10: node61port10,
			abs9port1:  node60port1,
			abs9port2:  node60port2,
			abs9port3:  node60port3,
			abs9port4:  node60port4,
		},
	}, {
		desc: "2 nodes, links via shared ports",
		graph: &AbstractGraph{
			Desc:  "2 nodes, links via shared ports",
			Nodes: []*AbstractNode{abs10, abs11},
			Edges: []*AbstractEdge{{abs10port1, abs11port1}, {abs10port2, abs11port2}},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs10: node20, abs11: node21},
		wantPorts: map[*AbstractPort]*ConcretePort{abs10port1: node20port1, abs10port2: node20port2, abs11port1: node21port2, abs11port2: node21port1},
	}, {
		desc: "2 nodes, 4 interconnected ports each",
		graph: &AbstractGraph{
			Desc:  "2 nodes, 4 interconnected ports each",
			Nodes: []*AbstractNode{abs12, abs13},
			Edges: []*AbstractEdge{{abs12port1, abs13port1}, {abs12port2, abs13port2}, {abs12port3, abs13port3}, {abs12port4, abs13port4}},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs12: node22, abs13: node23},
		wantPorts: map[*AbstractPort]*ConcretePort{
			abs12port1: node22port1,
			abs12port2: node22port2,
			abs12port3: node22port3,
			abs12port4: node22port4,
			abs13port1: node23port1,
			abs13port2: node23port2,
			abs13port3: node23port3,
			abs13port4: node23port4,
		},
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			a, err := Solve(tc.graph, superGraph)
			if err != nil {
				t.Fatalf("Solve got error %v, want nil", err)
			}
			if len(a.Node2Node) != len(tc.wantNodes) {
				t.Fatalf("Solve assigned %d nodes, want %d nodes", len(a.Node2Node), len(tc.wantNodes))
			}
			for abs, got := range a.Node2Node {
				if want, ok := tc.wantNodes[abs]; !ok {
					t.Fatalf("Solve assiged node abstract %q to %q; name does not exist", abs.Desc, got.Desc)
				} else if got != want {
					t.Errorf("Solve for node %q got %q, want %q", abs.Desc, got.Desc, want.Desc)
				}
			}
			if len(a.Port2Port) != len(tc.wantPorts) {
				t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), len(tc.wantPorts))
			}
			for port, got := range a.Port2Port {
				if want, ok := tc.wantPorts[port]; !ok {
					t.Fatalf("Solve assiged port  %q to %q; port does not exist", port.Desc, got.Desc)
				} else if got != want {
					t.Errorf("Solve for port %q got %q, want %q", port.Desc, got.Desc, want.Desc)
				}
			}
		})
	}
}

func TestSolveNotSolvable(t *testing.T) {
	tests := []struct {
		desc  string
		graph *AbstractGraph
	}{{
		desc:  "Node does not exist in super",
		graph: &AbstractGraph{Desc: "Node does not exist", Nodes: []*AbstractNode{&AbstractNode{Desc: "dut1", Attrs: map[string]Constraint{"DOES NOT EXIST": Equal("???")}}}},
	}, {
		desc: "Port does not exist in super",
		graph: &AbstractGraph{
			Desc: "Port does not exist",
			Nodes: []*AbstractNode{
				&AbstractNode{Desc: "dut1", Ports: []*AbstractPort{&AbstractPort{Desc: "dut1:port1", Attrs: map[string]Constraint{"DOES NOT EXIST": Equal("???")}}}},
			},
		},
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			a, err := Solve(tc.graph, superGraph)
			if a != nil {
				t.Errorf("Solve got assignment %v, want nil", a)
			}
			if err == nil {
				t.Errorf("Solve got nil error, want error")
			}
		})
	}
}

// Benchmark solve for 4 DUT
// Run via blaze test with --test_arg=--test.bench=. --nocache_test_results
func Benchmark4DutSolve(b *testing.B) {
	b.ReportAllocs()
	requestGraph := &AbstractGraph{
		Desc:  "four nodes, interconnected",
		Nodes: []*AbstractNode{abs3, abs4, abs5, abs6},
		Edges: []*AbstractEdge{{abs3port1, abs4port1}, {abs3port2, abs6port1}, {abs4port2, abs5port1}, {abs5port2, abs6port2}},
	}
	for i := 0; i < b.N; i++ {
		Solve(requestGraph, superGraph)
	}
}
