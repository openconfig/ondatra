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
	"regexp"
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
	node5port1   = &ConcretePort{Desc: "node5:port1"}
	node6port1   = &ConcretePort{Desc: "node6:port1"}
	node7port1   = &ConcretePort{Desc: "node7:port1"}
	node8port1   = &ConcretePort{Desc: "node8:port1"}
	node9port1   = &ConcretePort{Desc: "node9:port1", Attrs: map[string]string{"attr": "node9port1"}}
	node10port1  = &ConcretePort{Desc: "node10:port1"}
	node10port2  = &ConcretePort{Desc: "node10:port2"}
	node11port1  = &ConcretePort{Desc: "node11:port1"}
	node11port2  = &ConcretePort{Desc: "node11:port2"}
	node11port3  = &ConcretePort{Desc: "node11:port3"}
	node12port1  = &ConcretePort{Desc: "node12:port1"}
	node12port2  = &ConcretePort{Desc: "node12:port2"}
	node13port1  = &ConcretePort{Desc: "node13:port1"}
	node13port2  = &ConcretePort{Desc: "node13:port2"}
	node14port1  = &ConcretePort{Desc: "node14:port1"}
	node14port2  = &ConcretePort{Desc: "node14:port2"}
	node15port1  = &ConcretePort{Desc: "node15:port1"}
	node15port2  = &ConcretePort{Desc: "node15:port2"}
	node15port3  = &ConcretePort{Desc: "node15:port3"}
	node16port1  = &ConcretePort{Desc: "node16:port1"}
	node16port2  = &ConcretePort{Desc: "node16:port2"}
	node17port1  = &ConcretePort{Desc: "node17:port1"}
	node17port2  = &ConcretePort{Desc: "node17:port2"}
	node18port1  = &ConcretePort{Desc: "node18:port1"}
	node18port2  = &ConcretePort{Desc: "node18:port2"}
	node20port1  = &ConcretePort{Desc: "node20:port1", Attrs: map[string]string{"attr": "FOO"}}
	node20port2  = &ConcretePort{Desc: "node20:port2", Attrs: map[string]string{"attr": "BAR", "attr2": "1"}}
	node20port3  = &ConcretePort{Desc: "node20:port3", Attrs: map[string]string{"attr": "BAR", "attr2": "2"}}
	node21port1  = &ConcretePort{Desc: "node21:port1", Attrs: map[string]string{"attr": "FOO"}}
	node21port2  = &ConcretePort{Desc: "node21:port2", Attrs: map[string]string{"attr": "BAR"}}
	node22port1  = &ConcretePort{Desc: "node22:port1", Attrs: map[string]string{"attr": "A"}}
	node22port2  = &ConcretePort{Desc: "node22:port2", Attrs: map[string]string{"attr": "B"}}
	node22port3  = &ConcretePort{Desc: "node22:port3", Attrs: map[string]string{"attr": "C"}}
	node22port4  = &ConcretePort{Desc: "node22:port4", Attrs: map[string]string{"attr": "D"}}
	node23port1  = &ConcretePort{Desc: "node23:port1", Attrs: map[string]string{"attr": "A2"}}
	node23port2  = &ConcretePort{Desc: "node23:port2", Attrs: map[string]string{"attr": "B2"}}
	node23port3  = &ConcretePort{Desc: "node23:port3", Attrs: map[string]string{"attr": "C2"}}
	node23port4  = &ConcretePort{Desc: "node23:port4", Attrs: map[string]string{"attr": "D2"}}
	node32port1  = &ConcretePort{Desc: "node32:port1", Attrs: map[string]string{"group": "G1"}}
	node32port2  = &ConcretePort{Desc: "node32:port2", Attrs: map[string]string{"group": "G1"}}
	node32port3  = &ConcretePort{Desc: "node32:port3", Attrs: map[string]string{"group": "G2"}}
	node32port4  = &ConcretePort{Desc: "node32:port4", Attrs: map[string]string{"group": "G2"}}
	node32port5  = &ConcretePort{Desc: "node32:port5", Attrs: map[string]string{"group": "G3"}}
	node32port6  = &ConcretePort{Desc: "node32:port6", Attrs: map[string]string{"group": "G3"}}
	node32port7  = &ConcretePort{Desc: "node32:port7"}
	node32port8  = &ConcretePort{Desc: "node32:port8"}
	node32port9  = &ConcretePort{Desc: "node32:port9"}
	node32port10 = &ConcretePort{Desc: "node32:port10"}
	node32port11 = &ConcretePort{Desc: "node32:port11"}
	node32port12 = &ConcretePort{Desc: "node32:port12"}
	node32port13 = &ConcretePort{Desc: "node32:port13"}
	node32port14 = &ConcretePort{Desc: "node32:port14"}
	node32port15 = &ConcretePort{Desc: "node32:port15"}
	node33port1  = &ConcretePort{Desc: "node33:port1", Attrs: map[string]string{"attr": "A"}}
	node33port2  = &ConcretePort{Desc: "node33:port2"}
	node33port3  = &ConcretePort{Desc: "node33:port3", Attrs: map[string]string{"attr": "B"}}
	node33port4  = &ConcretePort{Desc: "node33:port4"}
	node33port5  = &ConcretePort{Desc: "node33:port5", Attrs: map[string]string{"attr": "C"}}
	node33port6  = &ConcretePort{Desc: "node33:port6"}
	node33port7  = &ConcretePort{Desc: "node33:port7"}
	node33port8  = &ConcretePort{Desc: "node33:port8"}
	node33port9  = &ConcretePort{Desc: "node33:port9"}
	node33port10 = &ConcretePort{Desc: "node33:port10"}
	node33port11 = &ConcretePort{Desc: "node33:port11"}
	node33port12 = &ConcretePort{Desc: "node33:port12"}
	node33port13 = &ConcretePort{Desc: "node33:port13"}
	node33port14 = &ConcretePort{Desc: "node33:port14"}
	node33port15 = &ConcretePort{Desc: "node33:port15"}
	node40port1  = &ConcretePort{Desc: "node40:port1", Attrs: map[string]string{"attr": "A"}}
	node40port2  = &ConcretePort{Desc: "node40:port2"}
	node41port1  = &ConcretePort{Desc: "node41:port1", Attrs: map[string]string{"attr": "A"}}
	node41port2  = &ConcretePort{Desc: "node41:port2"}
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
	node70port1  = &ConcretePort{Desc: "node70:port1", Attrs: map[string]string{"speed": "S_100G", "attr": "FOO"}}
	node70port2  = &ConcretePort{Desc: "node70:port2", Attrs: map[string]string{"speed": "S_100G"}}
	node70port3  = &ConcretePort{Desc: "node70:port3", Attrs: map[string]string{"speed": "S_101G"}}
	node80port1  = &ConcretePort{Desc: "node80:port1", Attrs: map[string]string{"attr1": "1", "attr2": "2", "attr3": "3", "attr4": "4"}}
	node80port2  = &ConcretePort{Desc: "node80:port2", Attrs: map[string]string{"attr1": "1", "attr2": "1", "attr3": "1", "attr4": "1"}}
	node80port3  = &ConcretePort{Desc: "node80:port3", Attrs: map[string]string{"attr1": "2", "attr2": "2", "attr3": "2", "attr4": "2"}}
	node80port4  = &ConcretePort{Desc: "node80:port4", Attrs: map[string]string{"attr1": "3", "attr2": "3", "attr3": "3", "attr4": "3"}}
	node80port5  = &ConcretePort{Desc: "node80:port5", Attrs: map[string]string{"attr1": "4", "attr2": "4", "attr3": "4", "attr4": "4"}}
	node81port1  = &ConcretePort{Desc: "node81:port1"}
	node81port2  = &ConcretePort{Desc: "node81:port2", Attrs: map[string]string{"attr1": "1", "attr2": "1", "attr3": "1", "attr4": "1"}}
	node81port3  = &ConcretePort{Desc: "node81:port3", Attrs: map[string]string{"attr1": "2", "attr2": "2", "attr3": "2", "attr4": "2"}}
	node81port4  = &ConcretePort{Desc: "node81:port4", Attrs: map[string]string{"attr1": "3", "attr2": "3", "attr3": "3", "attr4": "3"}}
	node81port5  = &ConcretePort{Desc: "node81:port5", Attrs: map[string]string{"attr1": "4", "attr2": "4", "attr3": "4", "attr4": "4"}}
	node90port1  = &ConcretePort{Desc: "node90:port1"}
	node90port2  = &ConcretePort{Desc: "node90:port2"}
	node90port3  = &ConcretePort{Desc: "node90:port3"}
	node90port4  = &ConcretePort{Desc: "node90:port4"}
	node90port5  = &ConcretePort{Desc: "node90:port5"}
	node91port1  = &ConcretePort{Desc: "node91:port1"}
	node91port2  = &ConcretePort{Desc: "node91:port2"}
	node92port1  = &ConcretePort{Desc: "node92:port1"}
	node92port2  = &ConcretePort{Desc: "node92:port2"}

	node1 = &ConcreteNode{Desc: "node1", Ports: []*ConcretePort{node1port1, node1port2}, Attrs: map[string]string{"vendor": "CISCO1"}}
	node2 = &ConcreteNode{Desc: "node2", Ports: []*ConcretePort{node2port1, node2port2}, Attrs: map[string]string{"vendor": "CISCO2"}}
	node3 = &ConcreteNode{Desc: "node3", Ports: []*ConcretePort{node3port1, node3port2}, Attrs: map[string]string{"vendor": "CISCO3"}}
	node4 = &ConcreteNode{Desc: "node4", Ports: []*ConcretePort{node4port1, node4port2}, Attrs: map[string]string{"vendor": "CISCO4"}}
	node5 = &ConcreteNode{Desc: "node5", Ports: []*ConcretePort{node5port1}, Attrs: map[string]string{"vendor": "CISCO1"}}
	node6 = &ConcreteNode{Desc: "node6", Ports: []*ConcretePort{node6port1}, Attrs: map[string]string{"vendor": "CISCO2"}}
	node7 = &ConcreteNode{Desc: "node7", Ports: []*ConcretePort{node7port1}, Attrs: map[string]string{"vendor": "CISCO3"}}
	node8 = &ConcreteNode{Desc: "node8", Ports: []*ConcretePort{node8port1}, Attrs: map[string]string{"attr": ""}}
	node9 = &ConcreteNode{Desc: "node9", Ports: []*ConcretePort{node9port1}, Attrs: map[string]string{"vendor": "UNIQUE9"}}
	// Four Nodes, interconnected; include additional Nodes with enough ports but not enough edges.
	node10 = &ConcreteNode{Desc: "node10", Ports: []*ConcretePort{node10port1, node10port2}, Attrs: map[string]string{"vendor": "CISCO2"}}
	node11 = &ConcreteNode{Desc: "node11", Ports: []*ConcretePort{node11port1, node11port2, node11port3}, Attrs: map[string]string{"vendor": "CISCO1", "attr": "1"}}
	node12 = &ConcreteNode{Desc: "node12", Ports: []*ConcretePort{node12port1, node12port2}, Attrs: map[string]string{"vendor": "CISCO2"}}
	node13 = &ConcreteNode{Desc: "node13", Ports: []*ConcretePort{node13port1, node13port2}, Attrs: map[string]string{"vendor": "CISCO3"}}
	node14 = &ConcreteNode{Desc: "node14", Ports: []*ConcretePort{node14port1, node14port2}, Attrs: map[string]string{"vendor": "CISCO4"}}
	node15 = &ConcreteNode{Desc: "node15", Ports: []*ConcretePort{node15port1, node15port2, node15port3}, Attrs: map[string]string{"vendor": "CISCO1", "attr": "1"}}
	node16 = &ConcreteNode{Desc: "node16", Ports: []*ConcretePort{node16port1, node16port2}, Attrs: map[string]string{"vendor": "CISCO2"}}
	node17 = &ConcreteNode{Desc: "node17", Ports: []*ConcretePort{node17port1, node17port2}, Attrs: map[string]string{"vendor": "CISCO3"}}
	node18 = &ConcreteNode{Desc: "node18", Ports: []*ConcretePort{node18port1, node18port2}, Attrs: map[string]string{"vendor": "CISCO4"}}
	node20 = &ConcreteNode{Desc: "node20", Ports: []*ConcretePort{node20port1, node20port2, node20port3}, Attrs: map[string]string{"attr": "multi1"}}
	node21 = &ConcreteNode{Desc: "node21", Ports: []*ConcretePort{node21port1, node21port2}, Attrs: map[string]string{"attr": "multi2"}}
	node22 = &ConcreteNode{Desc: "node22", Ports: []*ConcretePort{node22port1, node22port2, node22port3, node22port4}}
	node23 = &ConcreteNode{Desc: "node23", Ports: []*ConcretePort{node23port1, node23port2, node23port3, node23port4}}
	node30 = &ConcreteNode{Desc: "node30", Attrs: map[string]string{"attr": "PAIRED"}}
	node31 = &ConcreteNode{Desc: "node31", Attrs: map[string]string{"attr": "PAIRED", "vendor": "CISCO1"}}
	node32 = &ConcreteNode{Desc: "node32", Ports: []*ConcretePort{node32port1, node32port2, node32port3, node32port4, node32port5, node32port6, node32port7, node32port8, node32port9, node32port10, node32port11, node32port12, node32port13, node32port14, node32port15}, Attrs: map[string]string{"attr": "group"}}
	node33 = &ConcreteNode{Desc: "node33", Ports: []*ConcretePort{node33port1, node33port2, node33port3, node33port4, node33port5, node33port6, node33port7, node33port8, node33port9, node33port10, node33port11, node33port12, node33port13, node33port14, node33port15}, Attrs: map[string]string{"attr": "group"}}
	node40 = &ConcreteNode{Desc: "node40", Ports: []*ConcretePort{node40port1, node40port2}, Attrs: map[string]string{"vendor": "CISCO1", "attr": "40"}}
	node41 = &ConcreteNode{Desc: "node41", Ports: []*ConcretePort{node41port1, node41port2}, Attrs: map[string]string{"vendor": "CISCO1"}}
	node60 = &ConcreteNode{Desc: "node60", Ports: []*ConcretePort{node60port1, node60port2, node60port3, node60port4, node60port5, node60port6, node60port7, node60port8}, Attrs: map[string]string{"vendor": "CISCO1"}}
	node61 = &ConcreteNode{Desc: "node61", Ports: []*ConcretePort{node61port1, node61port2, node61port3, node61port4, node61port5, node61port6, node61port7, node61port8, node61port9, node61port10}, Attrs: map[string]string{"vendor": "CISCO1"}}
	node70 = &ConcreteNode{Desc: "node70", Ports: []*ConcretePort{node70port1, node70port2, node70port3}}
	node80 = &ConcreteNode{Desc: "node80", Ports: []*ConcretePort{node80port1, node80port2, node80port3, node80port4, node80port5}}
	node81 = &ConcreteNode{Desc: "node81", Ports: []*ConcretePort{node81port1, node81port2, node81port3, node81port4, node81port5}}
	node90 = &ConcreteNode{Desc: "node90", Ports: []*ConcretePort{node90port1, node90port2, node90port3, node90port4, node90port5}, Attrs: map[string]string{"vendor": "UNIQUE90"}}
	node91 = &ConcreteNode{Desc: "node91", Ports: []*ConcretePort{node91port1, node91port2}, Attrs: map[string]string{"vendor": "UNIQUE91"}}
	node92 = &ConcreteNode{Desc: "node92", Ports: []*ConcretePort{node92port1, node92port2}, Attrs: map[string]string{"vendor": "UNIQUE92"}}
)

var superGraph = &ConcreteGraph{
	Desc: "super",
	Nodes: []*ConcreteNode{
		node1, node2, node3, node4, node5, node6, node7, node8, node9, node10, node11, node12, node13, node14, node15, node16, node17, node18, node20, node21, node22, node23, node30, node31, node32, node33, node40, node41, node60, node61, node70, node80, node81, node90, node91, node92,
		// Additional entries for benchmark purposes.
		{Desc: "node19", Attrs: map[string]string{"vendor": "CISCO2"}},
		{Desc: "node24", Attrs: map[string]string{"vendor": "CISCO3"}},
		{Desc: "node25", Attrs: map[string]string{"vendor": "CISCO4"}},
		{Desc: "node26", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node27", Attrs: map[string]string{"vendor": "CISCO2"}},
		{Desc: "node28", Attrs: map[string]string{"vendor": "CISCO3"}},
		{Desc: "node29", Attrs: map[string]string{"vendor": "CISCO4"}},
		{Desc: "node32", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node33", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node34", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node35", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node36", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node37", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node38", Attrs: map[string]string{"vendor": "CISCO1"}},
		{Desc: "node39", Attrs: map[string]string{"vendor": "CISCO1"}},
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
		{Src: node5port1, Dst: node8port1},
		{Src: node6port1, Dst: node8port1},
		{Src: node7port1, Dst: node8port1},
		{Src: node8port1, Dst: node9port1},
		// 4 interconnected nodes
		{Src: node10port1, Dst: node14port1},
		{Src: node11port1, Dst: node10port1},
		{Src: node11port1, Dst: node12port1},
		{Src: node11port2, Dst: node14port2},
		{Src: node11port1, Dst: node16port1},
		{Src: node12port2, Dst: node14port1},
		{Src: node13port2, Dst: node14port2},
		{Src: node15port1, Dst: node16port1},
		{Src: node15port2, Dst: node18port1},
		{Src: node16port2, Dst: node17port1},
		{Src: node17port2, Dst: node18port2},
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
		{Src: node32port1, Dst: node33port1},
		{Src: node32port2, Dst: node33port2},
		{Src: node32port3, Dst: node33port3},
		{Src: node32port4, Dst: node33port4},
		{Src: node32port5, Dst: node33port5},
		{Src: node32port6, Dst: node33port6},
		{Src: node32port7, Dst: node33port7},
		{Src: node32port8, Dst: node33port8},
		{Src: node32port9, Dst: node33port9},
		{Src: node32port10, Dst: node33port10},
		{Src: node32port11, Dst: node33port11},
		{Src: node32port12, Dst: node33port12},
		{Src: node32port13, Dst: node33port13},
		{Src: node32port14, Dst: node33port14},
		{Src: node32port15, Dst: node33port15},
		{Src: node60port1, Dst: node61port1},
		{Src: node60port2, Dst: node61port2},
		{Src: node60port3, Dst: node61port3},
		{Src: node60port4, Dst: node61port4},
		{Src: node60port5, Dst: node61port5},
		{Src: node60port6, Dst: node61port6},
		{Src: node60port7, Dst: node61port7},
		{Src: node60port8, Dst: node61port8},
		{Src: node80port1, Dst: node81port1},
		{Src: node80port2, Dst: node81port2},
		{Src: node80port3, Dst: node81port3},
		{Src: node80port4, Dst: node81port4},
		{Src: node80port5, Dst: node81port5},
		{Src: node90port1, Dst: node91port1},
		{Src: node90port2, Dst: node91port2},
		{Src: node90port3, Dst: node92port1},
	},
}

// Setup abstract Nodes and Ports for testing.
var (
	// One Node
	abs1 = &AbstractNode{Desc: "abs1", Constraints: map[string]NodeConstraint{"attr": Equal("")}}

	// One Node, one Port
	abs2      = &AbstractNode{Desc: "abs2", Ports: []*AbstractPort{abs2port1}, Constraints: map[string]NodeConstraint{"vendor": Equal("UNIQUE9")}}
	abs2port1 = &AbstractPort{Desc: "abs2:port1"}

	// Four Nodes, interconnected
	abs3      = &AbstractNode{Desc: "abs3", Ports: []*AbstractPort{abs3port1, abs3port2, abs3port3}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO1"), "attr": Equal("1")}}
	abs4      = &AbstractNode{Desc: "abs4", Ports: []*AbstractPort{abs4port1, abs4port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO2")}}
	abs5      = &AbstractNode{Desc: "abs5", Ports: []*AbstractPort{abs5port1, abs5port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO3")}}
	abs6      = &AbstractNode{Desc: "abs6", Ports: []*AbstractPort{abs6port1, abs6port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO4")}}
	abs3port1 = &AbstractPort{Desc: "abs3:port1"}
	abs3port2 = &AbstractPort{Desc: "abs3:port2"}
	abs3port3 = &AbstractPort{Desc: "abs3:port3"}
	abs4port1 = &AbstractPort{Desc: "abs4:port1"}
	abs4port2 = &AbstractPort{Desc: "abs4:port2"}
	abs5port1 = &AbstractPort{Desc: "abs5:port1"}
	abs5port2 = &AbstractPort{Desc: "abs5:port2"}
	abs6port1 = &AbstractPort{Desc: "abs6:port1"}
	abs6port2 = &AbstractPort{Desc: "abs6:port2"}

	// One Node, 8 Ports
	abs7      = &AbstractNode{Desc: "abs3", Ports: []*AbstractPort{abs7port1, abs7port2, abs7port3, abs7port4, abs7port5, abs7port6, abs7port7, abs7port8}}
	abs7port1 = &AbstractPort{Desc: "abs7:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G")}}
	abs7port2 = &AbstractPort{Desc: "abs7:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_101G")}}
	abs7port3 = &AbstractPort{Desc: "abs7:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_102G")}}
	abs7port4 = &AbstractPort{Desc: "abs7:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_103G")}}
	abs7port5 = &AbstractPort{Desc: "abs7:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_104G")}}
	abs7port6 = &AbstractPort{Desc: "abs7:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_105G")}}
	abs7port7 = &AbstractPort{Desc: "abs7:port7", Constraints: map[string]PortConstraint{"speed": Equal("S_106G")}}
	abs7port8 = &AbstractPort{Desc: "abs7:port8", Constraints: map[string]PortConstraint{"speed": Equal("S_107G")}}

	// Two Nodes, Many Links
	abs8       = &AbstractNode{Desc: "abs8", Ports: []*AbstractPort{abs8port1, abs8port2, abs8port3, abs8port4, abs8port5, abs8port6, abs8port9, abs8port10}}
	abs9       = &AbstractNode{Desc: "abs9", Ports: []*AbstractPort{abs9port1, abs9port2, abs9port3, abs9port4}}
	abs8port1  = &AbstractPort{Desc: "abs8:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G")}}
	abs8port2  = &AbstractPort{Desc: "abs8:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_101G")}}
	abs8port3  = &AbstractPort{Desc: "abs8:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_102G")}}
	abs8port4  = &AbstractPort{Desc: "abs8:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_103G")}}
	abs8port5  = &AbstractPort{Desc: "abs8:port5", Constraints: map[string]PortConstraint{"speed": Equal("S_104G")}}
	abs8port6  = &AbstractPort{Desc: "abs8:port6", Constraints: map[string]PortConstraint{"speed": Equal("S_105G")}}
	abs8port9  = &AbstractPort{Desc: "abs8:port9", Constraints: map[string]PortConstraint{"attr": Equal("FOO")}}
	abs8port10 = &AbstractPort{Desc: "abs8:port10", Constraints: map[string]PortConstraint{"attr": Equal("BAR")}}
	abs9port1  = &AbstractPort{Desc: "abs9:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_100G")}}
	abs9port2  = &AbstractPort{Desc: "abs9:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_101G")}}
	abs9port3  = &AbstractPort{Desc: "abs9:port3", Constraints: map[string]PortConstraint{"speed": Equal("S_102G")}}
	abs9port4  = &AbstractPort{Desc: "abs9:port4", Constraints: map[string]PortConstraint{"speed": Equal("S_103G")}}

	// Two Nodes, links via shared ports
	abs10      = &AbstractNode{Desc: "abs10", Ports: []*AbstractPort{abs10port1, abs10port2}, Constraints: map[string]NodeConstraint{"attr": Equal("multi1")}}
	abs11      = &AbstractNode{Desc: "abs11", Ports: []*AbstractPort{abs11port1, abs11port2}}
	abs10port1 = &AbstractPort{Desc: "abs10:port1", Constraints: map[string]PortConstraint{"attr": Equal("FOO")}}
	abs10port2 = &AbstractPort{Desc: "abs10:port2", Constraints: map[string]PortConstraint{"attr": Equal("BAR")}}
	abs11port1 = &AbstractPort{Desc: "abs11:port1"}
	abs11port2 = &AbstractPort{Desc: "abs11:port2"}

	// 2 nodes, 4 interconnected ports
	abs12      = &AbstractNode{Desc: "abs12", Ports: []*AbstractPort{abs12port1, abs12port2, abs12port3, abs12port4}}
	abs13      = &AbstractNode{Desc: "abs13", Ports: []*AbstractPort{abs13port1, abs13port2, abs13port3, abs13port4}}
	abs12port1 = &AbstractPort{Desc: "abs12:port1"}
	abs12port2 = &AbstractPort{Desc: "abs12:port2", Constraints: map[string]PortConstraint{"attr": Equal("B")}}
	abs12port3 = &AbstractPort{Desc: "abs12:port3"}
	abs12port4 = &AbstractPort{Desc: "abs12:port4", Constraints: map[string]PortConstraint{"attr": Equal("D")}}
	abs13port1 = &AbstractPort{Desc: "abs13:port1", Constraints: map[string]PortConstraint{"attr": Equal("A2")}}
	abs13port2 = &AbstractPort{Desc: "abs13:port2"}
	abs13port3 = &AbstractPort{Desc: "abs13:port3", Constraints: map[string]PortConstraint{"attr": Equal("C2")}}
	abs13port4 = &AbstractPort{Desc: "abs13:port4"}

	// 2 nodes, same as node attr
	abs14 = &AbstractNode{Desc: "abs14", Constraints: map[string]NodeConstraint{"attr": SameAsNode(abs15)}}
	abs15 = &AbstractNode{Desc: "abs15", Constraints: map[string]NodeConstraint{"attr": Equal("PAIRED"), "vendor": Equal("CISCO1")}}

	// 2 nodes, 3 port groups
	abs16       = &AbstractNode{Desc: "abs16", Ports: []*AbstractPort{abs16port1, abs16port2, abs16port3, abs16port4, abs16port5, abs16port6, abs16port7, abs16port8, abs16port9, abs16port10, abs16port11, abs16port12, abs16port13, abs16port14, abs16port15}, Constraints: map[string]NodeConstraint{"attr": Equal("group")}}
	abs17       = &AbstractNode{Desc: "abs17", Ports: []*AbstractPort{abs17port1, abs17port2, abs17port3, abs17port4, abs17port5, abs17port6, abs17port7, abs17port8, abs17port9, abs17port10, abs17port11, abs17port12, abs17port13, abs17port14, abs17port15}, Constraints: map[string]NodeConstraint{"attr": Equal("group")}}
	abs16port1  = &AbstractPort{Desc: "abs16:port1", Constraints: map[string]PortConstraint{"group": SameAsPort(abs16port2)}}
	abs16port2  = &AbstractPort{Desc: "abs16:port2", Constraints: map[string]PortConstraint{"group": reAny}}
	abs16port3  = &AbstractPort{Desc: "abs16:port3", Constraints: map[string]PortConstraint{"group": AndPort(SameAsPort(abs16port4), NotSameAsPort(abs16port1))}}
	abs16port4  = &AbstractPort{Desc: "abs16:port4", Constraints: map[string]PortConstraint{"group": reAny}}
	abs16port5  = &AbstractPort{Desc: "abs16:port5", Constraints: map[string]PortConstraint{"group": AndPort(SameAsPort(abs16port6), NotSameAsPort(abs16port1), NotSameAsPort(abs16port3))}}
	abs16port6  = &AbstractPort{Desc: "abs16:port6"}
	abs16port7  = &AbstractPort{Desc: "abs16:port7"}
	abs16port8  = &AbstractPort{Desc: "abs16:port8"}
	abs16port9  = &AbstractPort{Desc: "abs16:port9"}
	abs16port10 = &AbstractPort{Desc: "abs16:port10"}
	abs16port11 = &AbstractPort{Desc: "abs16:port11"}
	abs16port12 = &AbstractPort{Desc: "abs16:port12"}
	abs16port13 = &AbstractPort{Desc: "abs16:port13"}
	abs16port14 = &AbstractPort{Desc: "abs16:port14"}
	abs16port15 = &AbstractPort{Desc: "abs16:port15"}
	abs17port1  = &AbstractPort{Desc: "abs17:port1", Constraints: map[string]PortConstraint{"attr": Equal("A")}}
	abs17port2  = &AbstractPort{Desc: "abs17:port2"}
	abs17port3  = &AbstractPort{Desc: "abs17:port3", Constraints: map[string]PortConstraint{"attr": Equal("B")}}
	abs17port4  = &AbstractPort{Desc: "abs17:port4"}
	abs17port5  = &AbstractPort{Desc: "abs17:port5", Constraints: map[string]PortConstraint{"attr": Equal("C")}}
	abs17port6  = &AbstractPort{Desc: "abs17:port6"}
	abs17port7  = &AbstractPort{Desc: "abs17:port7"}
	abs17port8  = &AbstractPort{Desc: "abs17:port8"}
	abs17port9  = &AbstractPort{Desc: "abs17:port9"}
	abs17port10 = &AbstractPort{Desc: "abs17:port10"}
	abs17port11 = &AbstractPort{Desc: "abs17:port11"}
	abs17port12 = &AbstractPort{Desc: "abs17:port12"}
	abs17port13 = &AbstractPort{Desc: "abs17:port13"}
	abs17port14 = &AbstractPort{Desc: "abs17:port14"}
	abs17port15 = &AbstractPort{Desc: "abs17:port15"}

	// 2 nodes, 2 ports with and constraints
	abs18      = &AbstractNode{Desc: "abs18", Ports: []*AbstractPort{abs18port1}, Constraints: map[string]NodeConstraint{"vendor": AndNode(Equal("CISCO1"), SameAsNode(abs19)), "attr": Equal("40")}}
	abs19      = &AbstractNode{Desc: "abs19", Ports: []*AbstractPort{abs19port1}}
	abs18port1 = &AbstractPort{Desc: "abs18:port1", Constraints: map[string]PortConstraint{"attr": AndPort(Equal("A"), SameAsPort(abs19port1))}}
	abs19port1 = &AbstractPort{Desc: "abs19:port1"}

	// 1 node, 3 ports
	// This is lightly constrained so there are multiple nodes/ports to check with a deferred constraint.
	abs20      = &AbstractNode{Desc: "abs20", Ports: []*AbstractPort{abs20port1, abs20port2, abs20port3}}
	abs20port1 = &AbstractPort{Desc: "abs20:port1", Constraints: map[string]PortConstraint{"speed": SameAsPort(abs20port2), "attr": Equal("FOO")}}
	abs20port2 = &AbstractPort{Desc: "abs20:port2", Constraints: map[string]PortConstraint{"speed": NotSameAsPort(abs20port3)}}
	abs20port3 = &AbstractPort{Desc: "abs20:port3"}

	// 2 nodes, many ports that depend on one port
	abs21      = &AbstractNode{Desc: "abs21", Ports: []*AbstractPort{abs21port1, abs21port2, abs21port3, abs21port4, abs21port5}}
	abs22      = &AbstractNode{Desc: "abs22", Ports: []*AbstractPort{abs22port1, abs22port2, abs22port3, abs22port4, abs22port5}}
	abs21port1 = &AbstractPort{Desc: "abs21:port1", Constraints: map[string]PortConstraint{"attr1": Equal("1"), "attr2": Equal("2"), "attr3": Equal("3"), "attr4": Equal("4")}} // This should be the most constrained so it's solved first.
	abs21port2 = &AbstractPort{Desc: "abs21:port2", Constraints: map[string]PortConstraint{"attr1": SameAsPort(abs21port1)}}
	abs21port3 = &AbstractPort{Desc: "abs21:port3", Constraints: map[string]PortConstraint{"attr2": SameAsPort(abs21port1)}}
	abs21port4 = &AbstractPort{Desc: "abs21:port4"}
	abs21port5 = &AbstractPort{Desc: "abs21:port5"}
	abs22port1 = &AbstractPort{Desc: "abs22:port1"}
	abs22port2 = &AbstractPort{Desc: "abs22:port2"}
	abs22port3 = &AbstractPort{Desc: "abs22:port3"}
	abs22port4 = &AbstractPort{Desc: "abs22:port4", Constraints: map[string]PortConstraint{"attr3": SameAsPort(abs21port1)}}
	abs22port5 = &AbstractPort{Desc: "abs22:port5", Constraints: map[string]PortConstraint{"attr4": SameAsPort(abs21port1)}}

	// 2 nodes, one link. Checking that many working nodes doesn't cause match issues.
	abs23      = &AbstractNode{Desc: "abs23", Ports: []*AbstractPort{abs23port1}, Constraints: map[string]NodeConstraint{"attr": Equal("")}}
	abs24      = &AbstractNode{Desc: "abs24", Ports: []*AbstractPort{abs24port1}}
	abs23port1 = &AbstractPort{Desc: "abs23:port1"}
	abs24port1 = &AbstractPort{Desc: "abs24:port1", Constraints: map[string]PortConstraint{"attr": Equal("node9port1")}}
)

func TestSolve(t *testing.T) {
	tests := []struct {
		desc            string
		graph           *AbstractGraph
		wantNodes       map[*AbstractNode]*ConcreteNode
		wantPorts       map[*AbstractPort]*ConcretePort
		wantSolvedPorts []*AbstractPort
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
		wantNodes: map[*AbstractNode]*ConcreteNode{abs2: node9},
		wantPorts: map[*AbstractPort]*ConcretePort{abs2port1: node9port1},
	}, {
		desc: "four nodes, interconnected",
		graph: &AbstractGraph{
			Desc:  "four nodes, interconnected",
			Nodes: []*AbstractNode{abs3, abs4, abs5, abs6},
			Edges: []*AbstractEdge{{abs3port1, abs4port1}, {abs3port2, abs6port1}, {abs4port2, abs5port1}, {abs5port2, abs6port2}},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs3: node15, abs4: node16, abs5: node17, abs6: node18},
		wantPorts: map[*AbstractPort]*ConcretePort{
			abs3port1: node15port1,
			abs3port2: node15port2,
			abs3port3: node15port3,
			abs4port1: node16port1,
			abs4port2: node16port2,
			abs5port1: node17port1,
			abs5port2: node17port2,
			abs6port1: node18port1,
			abs6port2: node18port2,
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
	}, {
		desc: "2 nodes, same as",
		graph: &AbstractGraph{
			Desc:  "2 nodes, same as",
			Nodes: []*AbstractNode{abs14, abs15},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs14: node30, abs15: node31},
	}, {
		desc: "2 nodes, 3 port groups",
		graph: &AbstractGraph{
			Desc:  "2 nodes, 3 port groups",
			Nodes: []*AbstractNode{abs16, abs17},
			Edges: []*AbstractEdge{{abs16port1, abs17port1}, {abs16port2, abs17port2}, {abs16port3, abs17port3}, {abs16port4, abs17port4}, {abs16port5, abs17port5}, {abs16port6, abs17port6}, {abs16port7, abs17port7}, {abs16port8, abs17port8}, {abs16port9, abs17port9}, {abs16port10, abs17port10}, {abs16port11, abs17port11}, {abs16port12, abs17port12}, {abs16port13, abs17port13}, {abs16port14, abs17port14}, {abs16port15, abs17port15}},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs16: node32, abs17: node33},
		wantPorts: map[*AbstractPort]*ConcretePort{
			abs16port1: node32port1,
			abs16port2: node32port2,
			abs16port3: node32port3,
			abs16port4: node32port4,
			abs16port5: node32port5,
			abs16port6: node32port6,
			abs17port1: node33port1,
			abs17port2: node33port2,
			abs17port3: node33port3,
			abs17port4: node33port4,
			abs17port5: node33port5,
			abs17port6: node33port6,
		},
		wantSolvedPorts: []*AbstractPort{
			abs16port7,
			abs16port8,
			abs16port9,
			abs16port10,
			abs16port11,
			abs16port12,
			abs16port13,
			abs16port14,
			abs16port15,
			abs17port7,
			abs17port8,
			abs17port9,
			abs17port10,
			abs17port11,
			abs17port12,
			abs17port13,
			abs17port14,
			abs17port15,
		},
	}, {
		desc: "2 nodes, 2 ports with ands",
		graph: &AbstractGraph{
			Desc:  "2 nodes, 2 ports with ands",
			Nodes: []*AbstractNode{abs18, abs19},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs18: node40, abs19: node41},
		wantPorts: map[*AbstractPort]*ConcretePort{abs18port1: node40port1, abs19port1: node41port1},
	}, {
		desc: "1 node, 3 ports with SameAs and NotSameAs",
		graph: &AbstractGraph{
			Desc:  "1 node, 3 ports with SameAs and NotSameAs",
			Nodes: []*AbstractNode{abs20},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs20: node70},
		wantPorts: map[*AbstractPort]*ConcretePort{abs20port1: node70port1, abs20port2: node70port2, abs20port3: node70port3},
	}, {
		desc: "2 nodes, other ports depend on same port",
		graph: &AbstractGraph{
			Desc:  "2 nodes, other ports depend on same port",
			Nodes: []*AbstractNode{abs21, abs22},
			Edges: []*AbstractEdge{{abs21port1, abs22port1}, {abs21port2, abs22port2}, {abs21port3, abs22port3}, {abs21port4, abs22port4}, {abs21port5, abs22port5}},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs21: node80, abs22: node81},
		wantPorts: map[*AbstractPort]*ConcretePort{
			abs21port1: node80port1,
			abs21port2: node80port2,
			abs21port3: node80port3,
			abs21port4: node80port4,
			abs21port5: node80port5,
			abs22port1: node81port1,
			abs22port2: node81port2,
			abs22port3: node81port3,
			abs22port4: node81port4,
			abs22port5: node81port5,
		},
	}, {
		desc: "2 nodes, 1 link; multiple possible nodes for second node",
		graph: &AbstractGraph{
			Desc:  "2 nodes, 1 link; multiple possible nodes for second node",
			Nodes: []*AbstractNode{abs23, abs24},
			Edges: []*AbstractEdge{{abs23port1, abs24port1}},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{abs23: node8, abs24: node9},
		wantPorts: map[*AbstractPort]*ConcretePort{abs23port1: node8port1, abs24port1: node9port1},
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
			if len(a.Port2Port) != (len(tc.wantPorts) + len(tc.wantSolvedPorts)) {
				t.Fatalf("Solve assigned %d ports, want %d ports", len(a.Port2Port), (len(tc.wantPorts) + len(tc.wantSolvedPorts)))
			}
			for port, got := range a.Port2Port {
				if want, ok := tc.wantPorts[port]; !ok {
					ok2 := false
					for _, p := range tc.wantSolvedPorts {
						if p == port {
							ok2 = true
							break
						}
					}
					if !ok2 {
						t.Fatalf("Solve assigned port %q to %q; port does not exist", port.Desc, got.Desc)
					}
				} else if got != want {
					t.Errorf("Solve for port %q got %q, want %q", port.Desc, got.Desc, want.Desc)
				}
			}
		})
	}
}

func TestSolveNotSolvable(t *testing.T) {
	nodeDesc1 := "dut1"
	nodeDesc2 := "dut2"
	nodeDesc3 := "dut3"
	portDesc1 := "dut:port1"
	portDesc2 := "dut:port2"
	portDesc3 := "dut:port3"
	portDesc4 := "dut:port4"
	portDesc5 := "dut:port5"
	portDesc6 := "dut:port6"
	portDesc7 := "dut:port7"
	portDesc8 := "dut:port8"
	samePort := &AbstractPort{Desc: portDesc1, Constraints: map[string]PortConstraint{"attr": Equal("BAR"), "attr2": Equal("1")}}
	aPort1 := &AbstractPort{Desc: portDesc1}
	aPort2 := &AbstractPort{Desc: portDesc2}
	aPort3 := &AbstractPort{Desc: portDesc3}
	aPort4 := &AbstractPort{Desc: portDesc4}
	zPort1 := &AbstractPort{Desc: portDesc5}
	zPort2 := &AbstractPort{Desc: portDesc6}
	zPort3 := &AbstractPort{Desc: portDesc7}
	zPort4 := &AbstractPort{Desc: portDesc8}

	tests := []struct {
		desc           string
		graph          *AbstractGraph
		wantAssigned   map[string]string
		wantUnassigned []string
	}{{
		desc:           "Node does not exist in super",
		graph:          &AbstractGraph{Desc: "Node does not exist", Nodes: []*AbstractNode{&AbstractNode{Desc: nodeDesc1, Constraints: map[string]NodeConstraint{"DOES NOT EXIST": Equal("???")}}}},
		wantUnassigned: []string{nodeDesc1},
	}, {
		desc: "Port does not exist in super",
		graph: &AbstractGraph{
			Desc: "Port does not exist",
			Nodes: []*AbstractNode{
				&AbstractNode{
					Desc:        nodeDesc1,
					Ports:       []*AbstractPort{&AbstractPort{Desc: portDesc1, Constraints: map[string]PortConstraint{"DOES NOT EXIST": Equal("???")}}},
					Constraints: map[string]NodeConstraint{"attr": Equal("multi1")},
				},
			},
		},
		wantAssigned:   map[string]string{nodeDesc1: "node20"},
		wantUnassigned: []string{portDesc1},
	}, {
		desc: "One Port does not exist in super",
		graph: &AbstractGraph{
			Desc: "One port does not exist",
			Nodes: []*AbstractNode{
				&AbstractNode{
					Desc: nodeDesc1,
					Ports: []*AbstractPort{
						samePort,
						&AbstractPort{Desc: portDesc2, Constraints: map[string]PortConstraint{"attr": SameAsPort(samePort), "attr2": Equal("2")}},
						// portDesc3 must be least constrained in order to be solved for last.
						&AbstractPort{Desc: portDesc3, Constraints: map[string]PortConstraint{"attr2": Equal("1")}},
					},
					Constraints: map[string]NodeConstraint{"attr": Equal("multi1")},
				},
			},
		},
		wantAssigned:   map[string]string{nodeDesc1: "node20", portDesc1: "node20:port2", portDesc2: "node20:port3"},
		wantUnassigned: []string{portDesc3},
	}, {
		desc: "Not enough edges to satisfy",
		graph: &AbstractGraph{
			Desc: "Not enough edges to satisfy",
			Nodes: []*AbstractNode{
				&AbstractNode{
					Desc:        nodeDesc1,
					Constraints: map[string]NodeConstraint{"vendor": Equal("UNIQUE90")},
					Ports:       []*AbstractPort{aPort1, aPort2, aPort3, aPort4},
				},
				&AbstractNode{Desc: nodeDesc2, Ports: []*AbstractPort{zPort1, zPort2}},
				&AbstractNode{Desc: nodeDesc3, Ports: []*AbstractPort{zPort3, zPort4}},
			},
			Edges: []*AbstractEdge{{aPort1, zPort1}, {aPort2, zPort2}, {aPort3, zPort3}, {aPort4, zPort4}},
		},
		wantUnassigned: []string{nodeDesc1, nodeDesc2, nodeDesc3, portDesc1, portDesc2, portDesc3, portDesc4, portDesc5, portDesc6, portDesc7, portDesc8},
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
			solveErr, ok := err.(*SolveErr)
			if !ok {
				t.Fatal("Solve got not *SolveErr type err, want *SolveErr type")
			}
			errString := solveErr.String()
			for abs, con := range tc.wantAssigned {
				re := regexp.MustCompile(fmt.Sprintf("%s.*%s.*%s", abs, "assigned", con))
				if !re.MatchString(errString) {
					t.Errorf("Solve got error %q, want error to report %q is assigned to %q", errString, abs, con)
				}
			}
			for _, unassigned := range tc.wantUnassigned {
				re := regexp.MustCompile(fmt.Sprintf("%s.*%s", unassigned, "not assigned"))
				if !re.MatchString(errString) {
					t.Errorf("Solve got error %q, want error to report %q is not assigned", errString, unassigned)
				}
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
