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

var (
	node101port1 = &ConcretePort{Desc: "node101port1"}
	node101port2 = &ConcretePort{Desc: "node101port2"}
	node102port1 = &ConcretePort{Desc: "node102port1"}
	node102port2 = &ConcretePort{Desc: "node102port2"}
	node103port1 = &ConcretePort{Desc: "node103port1"}
	node103port2 = &ConcretePort{Desc: "node103port2"}
	node103port3 = &ConcretePort{Desc: "node103port3"}
	node104port1 = &ConcretePort{Desc: "node104port1"}
	node104port2 = &ConcretePort{Desc: "node104port2"}
	node105port1 = &ConcretePort{Desc: "node105port1"}
	node105port2 = &ConcretePort{Desc: "node105port2"}
	node105port3 = &ConcretePort{Desc: "node105port3"}
	node106port1 = &ConcretePort{Desc: "node106port1"}
	node106port2 = &ConcretePort{Desc: "node106port2"}
	node107port1 = &ConcretePort{Desc: "node107port1"}
	node107port2 = &ConcretePort{Desc: "node107port2"}
	node107port3 = &ConcretePort{Desc: "node107port3"}
	node108port1 = &ConcretePort{Desc: "node108port1"}
	node108port2 = &ConcretePort{Desc: "node108port2"}
	node108port3 = &ConcretePort{Desc: "node108port3"}
	node109port1 = &ConcretePort{Desc: "node109port1"}
	node109port2 = &ConcretePort{Desc: "node109port2"}
	node109port3 = &ConcretePort{Desc: "node109port3"}
	node110port1 = &ConcretePort{Desc: "node110port1"}
	node110port2 = &ConcretePort{Desc: "node110port2"}
	node110port3 = &ConcretePort{Desc: "node110port3"}
	node111port1 = &ConcretePort{Desc: "node111port1"}
	node111port2 = &ConcretePort{Desc: "node111port2"}
	node112port1 = &ConcretePort{Desc: "node112port1"}
	node112port2 = &ConcretePort{Desc: "node112port2"}
	node112port3 = &ConcretePort{Desc: "node112port3"}
	node113port1 = &ConcretePort{Desc: "node113port1"}
	node113port2 = &ConcretePort{Desc: "node113port2"}
	node113port3 = &ConcretePort{Desc: "node113port3"}
	node114port1 = &ConcretePort{Desc: "node114port1"}
	node114port2 = &ConcretePort{Desc: "node114port2"}
	node115port1 = &ConcretePort{Desc: "node115port1"}
	node115port2 = &ConcretePort{Desc: "node115port2"}
	node115port3 = &ConcretePort{Desc: "node115port3"}
	node116port1 = &ConcretePort{Desc: "node116port1"}
	node116port2 = &ConcretePort{Desc: "node116port2"}
	node116port3 = &ConcretePort{Desc: "node116port3"}
	node117port1 = &ConcretePort{Desc: "node117port1"}
	node117port2 = &ConcretePort{Desc: "node117port2"}
	node117port3 = &ConcretePort{Desc: "node117port3"}
	node117port4 = &ConcretePort{Desc: "node117port4"}
	node118port1 = &ConcretePort{Desc: "node118port1"}
	node118port2 = &ConcretePort{Desc: "node118port2"}
	node118port3 = &ConcretePort{Desc: "node118port3"}
	node119port1 = &ConcretePort{Desc: "node119port1"}
	node119port2 = &ConcretePort{Desc: "node119port2"}
	node119port3 = &ConcretePort{Desc: "node119port3"}

	node101 = &ConcreteNode{Desc: "node101", Ports: []*ConcretePort{node101port1, node101port2}, Attrs: map[string]string{"type": "NOKIA_SRL"}}
	node102 = &ConcreteNode{Desc: "node102", Ports: []*ConcretePort{node102port1, node102port2}, Attrs: map[string]string{"type": "NOKIA_SRL"}}
	node103 = &ConcreteNode{Desc: "node103", Ports: []*ConcretePort{node103port1, node103port2, node103port3}, Attrs: map[string]string{"type": "NOKIA_SRL"}}
	node104 = &ConcreteNode{Desc: "node104", Ports: []*ConcretePort{node104port1, node104port2}, Attrs: map[string]string{"type": "NOKIA_SRL"}}
	node105 = &ConcreteNode{Desc: "node105", Ports: []*ConcretePort{node105port1, node105port2, node105port3}, Attrs: map[string]string{"type": "NOKIA_SRL"}}
	node106 = &ConcreteNode{Desc: "node016", Ports: []*ConcretePort{node106port1, node106port2}, Attrs: map[string]string{"type": "NOKIA_SRL"}}
	node107 = &ConcreteNode{Desc: "node107", Ports: []*ConcretePort{node107port1, node107port2, node107port3}, Attrs: map[string]string{"type": "NOKIA_SRL"}}
	node108 = &ConcreteNode{Desc: "node108", Ports: []*ConcretePort{node108port1, node108port2, node108port3}, Attrs: map[string]string{"type": "NOKIA_SRL"}}
	node109 = &ConcreteNode{Desc: "node109", Ports: []*ConcretePort{node109port1, node109port2, node109port3}, Attrs: map[string]string{"type": "CISCO_CXR"}}
	node110 = &ConcreteNode{Desc: "node110", Ports: []*ConcretePort{node110port1, node110port2, node110port3}, Attrs: map[string]string{"type": "CISCO_CXR"}}
	node111 = &ConcreteNode{Desc: "node111", Ports: []*ConcretePort{node111port1, node111port2}, Attrs: map[string]string{"type": "NOKIA_SRL"}}
	node112 = &ConcreteNode{Desc: "node112", Ports: []*ConcretePort{node112port1, node112port2, node112port3}, Attrs: map[string]string{"type": "JUNIPER_CEVO"}}
	node113 = &ConcreteNode{Desc: "node113", Ports: []*ConcretePort{node113port1, node113port2, node113port3}, Attrs: map[string]string{"type": "JUNIPER_CEVO"}}
	node114 = &ConcreteNode{Desc: "node114", Ports: []*ConcretePort{node114port1, node114port2}, Attrs: map[string]string{"type": "JUNIPER_CEVO"}}
	node115 = &ConcreteNode{Desc: "node115", Ports: []*ConcretePort{node115port1, node115port2, node115port3}, Attrs: map[string]string{"type": "JUNIPER_CEVO"}}
	node116 = &ConcreteNode{Desc: "node116", Ports: []*ConcretePort{node116port1, node116port2, node116port3}, Attrs: map[string]string{"type": "ARISTA_CEOS"}}
	node117 = &ConcreteNode{Desc: "node117", Ports: []*ConcretePort{node117port1, node117port2, node117port3, node117port4}, Attrs: map[string]string{"type": "ARISTA_CEOS"}}
	node118 = &ConcreteNode{Desc: "node118", Ports: []*ConcretePort{node118port1, node118port2, node118port3}, Attrs: map[string]string{"type": "ARISTA_CEOS"}}
	node119 = &ConcreteNode{Desc: "node119", Ports: []*ConcretePort{node119port1, node119port2, node119port3}, Attrs: map[string]string{"type": "IXIA_TG"}}
)

var benchGraph = &ConcreteGraph{
	Desc: "bench supergraph",
	Nodes: []*ConcreteNode{
		node101, node102, node103, node104, node105, node106, node107, node108, node109, node110, node111, node112, node113, node114, node115, node116, node117, node118, node119,
	},
	Edges: []*ConcreteEdge{
		{Src: node101port1, Dst: node102port1},
		{Src: node103port1, Dst: node104port1},
		{Src: node105port1, Dst: node106port1},
		{Src: node107port1, Dst: node108port1},
		{Src: node101port2, Dst: node109port1},
		{Src: node102port2, Dst: node109port2},
		{Src: node109port3, Dst: node116port1},
		{Src: node110port1, Dst: node107port2},
		{Src: node110port2, Dst: node108port2},
		{Src: node110port3, Dst: node117port1},
		{Src: node111port2, Dst: node105port2},
		{Src: node112port1, Dst: node103port2},
		{Src: node112port2, Dst: node104port2},
		{Src: node112port3, Dst: node116port2},
		{Src: node113port1, Dst: node103port3},
		{Src: node113port2, Dst: node105port3},
		{Src: node113port3, Dst: node117port2},
		{Src: node114port1, Dst: node106port2},
		{Src: node114port2, Dst: node117port3},
		{Src: node115port1, Dst: node107port3},
		{Src: node115port2, Dst: node108port3},
		{Src: node115port3, Dst: node118port2},
		{Src: node116port3, Dst: node119port1},
		{Src: node117port4, Dst: node119port2},
		{Src: node118port3, Dst: node119port3},
	},
}

func BenchmarkSolveScale(b *testing.B) {
	var (
		dut1port1 = &AbstractPort{Desc: "dut1:port1"}
		dut1port2 = &AbstractPort{Desc: "dut1:port2"}
		dut1port3 = &AbstractPort{Desc: "dut1:port3"}
		dut2port1 = &AbstractPort{Desc: "dut2:port1"}
		dut2port2 = &AbstractPort{Desc: "dut2:port2"}
		dut2port3 = &AbstractPort{Desc: "dut2:port3"}
		dut3port1 = &AbstractPort{Desc: "dut3:port1"}
		dut3port2 = &AbstractPort{Desc: "dut3:port2"}
		dut3port3 = &AbstractPort{Desc: "dut3:port3"}
		dut4port1 = &AbstractPort{Desc: "dut4:port1"}
		dut4port2 = &AbstractPort{Desc: "dut4:port2"}
		dut5port1 = &AbstractPort{Desc: "dut5:port1"}
		dut5port2 = &AbstractPort{Desc: "dut5:port2"}
		dut6port1 = &AbstractPort{Desc: "dut6:port1"}
		dut6port2 = &AbstractPort{Desc: "dut6:port2"}
		dut7port1 = &AbstractPort{Desc: "dut7:port1"}
		dut7port2 = &AbstractPort{Desc: "dut7:port2"}
		ate1port1 = &AbstractPort{Desc: "ate1:port1"}
	)
	absGraph := &AbstractGraph{
		Desc: "benchmark abstract graph",
		Nodes: []*AbstractNode{{
			Desc:        "dut1",
			Ports:       []*AbstractPort{dut1port1, dut1port2, dut1port3},
			Constraints: map[string]NodeConstraint{"type": Equal("ARISTA_CEOS")},
		}, {
			Desc:        "dut2",
			Ports:       []*AbstractPort{dut2port1, dut2port2, dut2port3},
			Constraints: map[string]NodeConstraint{"type": Equal("CISCO_CXR")},
		}, {
			Desc:        "dut3",
			Ports:       []*AbstractPort{dut3port1, dut3port2, dut3port3},
			Constraints: map[string]NodeConstraint{"type": Equal("JUNIPER_CEVO")},
		}, {
			Desc:        "dut4",
			Ports:       []*AbstractPort{dut4port1, dut4port2},
			Constraints: map[string]NodeConstraint{"type": Equal("NOKIA_SRL")},
		}, {
			Desc:        "dut5",
			Ports:       []*AbstractPort{dut5port1, dut5port2},
			Constraints: map[string]NodeConstraint{"type": Equal("NOKIA_SRL")},
		}, {
			Desc:        "dut6",
			Ports:       []*AbstractPort{dut6port1, dut6port2},
			Constraints: map[string]NodeConstraint{"type": Equal("NOKIA_SRL")},
		}, {
			Desc:        "dut7",
			Ports:       []*AbstractPort{dut7port1, dut7port2},
			Constraints: map[string]NodeConstraint{"type": Equal("NOKIA_SRL")},
		}, {
			Desc:        "ate1",
			Ports:       []*AbstractPort{ate1port1},
			Constraints: map[string]NodeConstraint{"type": Equal("IXIA_TG")},
		}},
		Edges: []*AbstractEdge{
			{Src: dut1port1, Dst: dut2port1},
			{Src: dut1port2, Dst: dut3port1},
			{Src: dut1port3, Dst: ate1port1},
			{Src: dut2port2, Dst: dut4port1},
			{Src: dut2port3, Dst: dut5port1},
			{Src: dut4port2, Dst: dut5port2},
			{Src: dut3port2, Dst: dut6port1},
			{Src: dut3port3, Dst: dut7port1},
			{Src: dut6port2, Dst: dut7port2},
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_, err := Solve(absGraph, benchGraph)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Solve() got unexpected error: %v", err)
		}
	}
}

func BenchmarkManyLinksSolve(b *testing.B) {
	b.ReportAllocs()
	requestGraph := &AbstractGraph{
		Desc:  "2 nodes, 3 port groups",
		Nodes: []*AbstractNode{abs16, abs17},
		Edges: []*AbstractEdge{{abs16port1, abs17port1}, {abs16port2, abs17port2}, {abs16port3, abs17port3}, {abs16port4, abs17port4}, {abs16port5, abs17port5}, {abs16port6, abs17port6}, {abs16port7, abs17port7}, {abs16port8, abs17port8}, {abs16port9, abs17port9}, {abs16port10, abs17port10}, {abs16port11, abs17port11}, {abs16port12, abs17port12}, {abs16port13, abs17port13}, {abs16port14, abs17port14}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_, err := Solve(requestGraph, superGraph)
		b.StopTimer()
		if err != nil {
			b.Fatalf("Solve() got unexpected error: %v", err)
		}
	}
}
