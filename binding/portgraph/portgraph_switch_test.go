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

	"golang.org/x/net/context"
)

// Initialize ConcretePorts and ConcreteNodes for the super graph.
// These are the ConcreteNodes and ConcretePorts that should be matched to in tests.
var (
	dutnode1port1 = &ConcretePort{Desc: "dutnode1:port1"}
	dutnode1port2 = &ConcretePort{Desc: "dutnode1:port2"}
	atenode2port1 = &ConcretePort{Desc: "atenode2:port1"}
	atenode2port2 = &ConcretePort{Desc: "atenode2:port2"}
	dutnode3port1 = &ConcretePort{Desc: "dutnode3:port1"}
	dutnode3port2 = &ConcretePort{Desc: "dutnode3:port2"}
	atenode4port1 = &ConcretePort{Desc: "atenode4:port1"}
	atenode4port2 = &ConcretePort{Desc: "atenode4:port2"}
	dutnode5port1 = &ConcretePort{Desc: "dutnode5:port1"}
	atenode6port1 = &ConcretePort{Desc: "atenode6:port1"}
	dutnode7port1 = &ConcretePort{Desc: "dutnode7:port1"}
	atenode8port1 = &ConcretePort{Desc: "atenode8:port1"}
	dutnode9port1  = &ConcretePort{Desc: "dutnode9:port1", Attrs: map[string]string{"attr": "FOO"}}
	dutnode9port2  = &ConcretePort{Desc: "dutnode9:port2", Attrs: map[string]string{"attr": "BAR", "attr2": "1"}}
	dutnode9port3  = &ConcretePort{Desc: "dutnode9:port3", Attrs: map[string]string{"attr": "BAR", "attr2": "2"}}
	atenode10port1  = &ConcretePort{Desc: "atenode10:port1", Attrs: map[string]string{"attr": "FOO"}}
	atenode10port2  = &ConcretePort{Desc: "atenode10:port2", Attrs: map[string]string{"attr": "BAR"}}

	switchnode9port1  = &ConcretePort{Desc: "switchnode9:port1"}
	switchnode9port2  = &ConcretePort{Desc: "switchnode9:port2"}
	switchnode9port3  = &ConcretePort{Desc: "switchnode9:port3"}
	switchnode9port4  = &ConcretePort{Desc: "switchnode9:port4"}
	switchnode9port5  = &ConcretePort{Desc: "switchnode9:port5"}
	switchnode9port6  = &ConcretePort{Desc: "switchnode9:port6"}
	switchnode9port7  = &ConcretePort{Desc: "switchnode9:port7"}
	switchnode9port8  = &ConcretePort{Desc: "switchnode9:port8"}
	switchnode10port1 = &ConcretePort{Desc: "switchnode10:port1"}
	switchnode10port2 = &ConcretePort{Desc: "switchnode10:port2"}
	switchnode10port3 = &ConcretePort{Desc: "switchnode10:port3"}
	switchnode10port4 = &ConcretePort{Desc: "switchnode10:port4"}
	switchnode10port5 = &ConcretePort{Desc: "switchnode10:port5"}
	switchnode10port6 = &ConcretePort{Desc: "switchnode10:port6"}
	switchnode10port7 = &ConcretePort{Desc: "switchnode10:port7"}
	switchnode10port8 = &ConcretePort{Desc: "switchnode10:port8"}
	
	dutnode1  = &ConcreteNode{Desc: "dutnode1", Ports: []*ConcretePort{dutnode1port1, dutnode1port2}, Attrs: map[string]string{"vendor": "CISCO1"}}
	atenode2  = &ConcreteNode{Desc: "atenode2", Ports: []*ConcretePort{atenode2port1, atenode2port2}, Attrs: map[string]string{"vendor": "KEYSIGHT"}}
	dutnode3  = &ConcreteNode{Desc: "dutnode3", Ports: []*ConcretePort{dutnode3port1, dutnode3port2}, Attrs: map[string]string{"vendor": "CISCO3"}}
	atenode4  = &ConcreteNode{Desc: "atenode4", Ports: []*ConcretePort{atenode4port1, atenode4port2}, Attrs: map[string]string{"vendor": "KEYSIGHT"}}
	dutnode5  = &ConcreteNode{Desc: "dutnode5", Ports: []*ConcretePort{dutnode5port1}, Attrs: map[string]string{"vendor": "CISCO1"}}
	atenode6  = &ConcreteNode{Desc: "atenode6", Ports: []*ConcretePort{atenode6port1}, Attrs: map[string]string{"vendor": "KEYSIGHT"}}
	dutnode7  = &ConcreteNode{Desc: "dutnode7", Ports: []*ConcretePort{dutnode7port1}, Attrs: map[string]string{"vendor": "CISCO3"}}
	atenode8  = &ConcreteNode{Desc: "atenode8", Ports: []*ConcretePort{atenode8port1}, Attrs: map[string]string{"vendor": "KEYSIGHT"}}
	dutnode9 = &ConcreteNode{Desc: "dutnode9", Ports: []*ConcretePort{dutnode9port1, dutnode9port2, dutnode9port3}, Attrs: map[string]string{"attr": "multi1"}}
	atenode10 = &ConcreteNode{Desc: "atenode10", Ports: []*ConcretePort{atenode10port1, atenode10port2}, Attrs: map[string]string{"attr": "multi2"}}
	switchnode9  = &ConcreteNode{Desc: "switchnode9", Ports: []*ConcretePort{switchnode9port1, switchnode9port2, switchnode9port3, switchnode9port4, switchnode9port5, switchnode9port6, switchnode9port7, switchnode9port8}, Attrs: map[string]string{"role": "Switch"}}
	switchnode10 = &ConcreteNode{Desc: "switchnode10", Ports: []*ConcretePort{switchnode10port1, switchnode10port2, switchnode10port3, switchnode10port4, switchnode10port5, switchnode10port6, switchnode10port7, switchnode10port8}, Attrs: map[string]string{"role": "Switch"}}
)

var superGraphTest = &ConcreteGraph{
	Desc: "super",
	Nodes: []*ConcreteNode{
		dutnode1, atenode2, dutnode3, atenode4, dutnode5, atenode6, dutnode7, atenode8, dutnode9, atenode10, switchnode9, switchnode10,
	},
	Edges: []*ConcreteEdge{
		{Src: dutnode1port1, Dst: switchnode9port1},
		{Src: switchnode9port2, Dst: atenode2port1},
		{Src: dutnode3port1, Dst: switchnode9port3},
		{Src: switchnode9port4, Dst: atenode4port1},
		{Src: dutnode5port1, Dst: switchnode9port5},
		{Src: switchnode9port6, Dst: atenode6port1},
		{Src: dutnode7port1, Dst: switchnode9port7},
		{Src: switchnode9port8, Dst: atenode8port1},
		{Src: dutnode1port2, Dst: switchnode10port1},
		{Src: switchnode10port2, Dst: atenode2port2},
		{Src: dutnode3port2, Dst: switchnode10port3},
		{Src: switchnode10port4, Dst: atenode4port2},
		{Src: dutnode9port1, Dst: atenode10port1}, 
		{Src: dutnode9port1, Dst: atenode10port2},
		{Src: dutnode9port2, Dst: atenode10port1},
	},
}

// Setup abstract Nodes and Ports for testing.
var (
	
	// Two Nodes, interconnected with Switch
	sabs1       = &AbstractNode{Desc: "sabs1", Ports: []*AbstractPort{sabs1port1, sabs1port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
	sabs2       = &AbstractNode{Desc: "sabs2", Ports: []*AbstractPort{sabs2port1, sabs2port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
	sabs1port1  = &AbstractPort{Desc: "sabs1:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
	sabs1port2  = &AbstractPort{Desc: "sabs1:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
	sabs2port1  = &AbstractPort{Desc: "sabs2:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
	sabs2port2  = &AbstractPort{Desc: "sabs2:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
	// Four Nodes, interconnected with multiple Switch
	sabs3       = &AbstractNode{Desc: "sabs3", Ports: []*AbstractPort{sabs3port1, sabs3port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
	sabs4       = &AbstractNode{Desc: "sabs4", Ports: []*AbstractPort{sabs4port1, sabs4port2}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
	sabs5       = &AbstractNode{Desc: "sabs1", Ports: []*AbstractPort{sabs5port1}, Constraints: map[string]NodeConstraint{"vendor": Equal("CISCO")}}
	sabs6       = &AbstractNode{Desc: "sabs2", Ports: []*AbstractPort{sabs6port1}, Constraints: map[string]NodeConstraint{"vendor": Equal("TGEN")}}
	sabs3port1  = &AbstractPort{Desc: "sabs3:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
	sabs3port2  = &AbstractPort{Desc: "sabs3:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
	sabs4port1  = &AbstractPort{Desc: "sabs4:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
	sabs4port2  = &AbstractPort{Desc: "sabs4:port2", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
	sabs5port1  = &AbstractPort{Desc: "sabs5:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}
	sabs6port1  = &AbstractPort{Desc: "sabs6:port1", Constraints: map[string]PortConstraint{"speed": Equal("S_400GB")}}	
)

func TestSolveTest(t *testing.T) {
	tests := []struct {
		desc            string
		graph           *AbstractGraph
		wantNodes       map[*AbstractNode]*ConcreteNode
		wantPorts       map[*AbstractPort]*ConcretePort
		wantSolvedPorts []*AbstractPort
	}{{
		desc: "Two nodes, interconnected with Switch",
		graph: &AbstractGraph{
			Desc:  "four nodes, interconnected",
			Nodes: []*AbstractNode{sabs1, sabs2},
			Edges: []*AbstractEdge{{sabs1port1, sabs2port1}, {sabs1port2, sabs2port2}},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{sabs1: dutnode1, sabs2: atenode2},
		wantPorts: map[*AbstractPort]*ConcretePort{
			sabs1port1:  dutnode1port1,
			sabs1port2:  dutnode1port2,
			sabs2port1:  atenode2port1,
			sabs2port2:  atenode2port2,
		},
	}, {
		desc: "Four nodes, interconnected with multiple Switch",
		graph: &AbstractGraph{
			Desc:  "four nodes, interconnected",
			Nodes: []*AbstractNode{sabs3, sabs4, sabs5, sabs6},
			Edges: []*AbstractEdge{{sabs3port1, sabs4port1}, {sabs3port2, sabs4port2}, {sabs5port1, sabs6port1}},
		},
		wantNodes: map[*AbstractNode]*ConcreteNode{sabs3: dutnode3, sabs4: atenode4, sabs5: dutnode5, sabs6: atenode6},
		wantPorts: map[*AbstractPort]*ConcretePort{
			sabs3port1:  dutnode3port1,
			sabs3port2:  dutnode3port2,
			sabs4port1:  atenode4port1,
			sabs4port2:  atenode4port2,
			sabs5port1:  dutnode5port1,
			sabs6port1:  atenode6port1,
		},
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			a, err := Solve(context.Background(), tc.graph, superGraphTest)
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

func TestSolveNotSolvableTest(t *testing.T) {
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
		wantAssigned:   map[string]string{nodeDesc1: "dutnode9"},
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
		wantAssigned:   map[string]string{nodeDesc1: "dutnode9", portDesc1: "dutnode9:port2", portDesc2: "dutnode9:port3"},
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
			a, err := Solve(context.Background(), tc.graph, superGraphTest)
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

func TestSolveCancelledSwitch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	requestGraph := &AbstractGraph{
		Desc:  "four nodes, interconnected with Switch",
		// Nodes: []*AbstractNode{absdut, absdut1, absate, absate1},
		// Edges: []*AbstractEdge{{absdutport1, absateport1}, {absdut1port1, absate1port1}},
		Nodes: []*AbstractNode{abs3, abs4, abs5, abs6},
		Edges: []*AbstractEdge{{abs3port1, abs4port1}, {abs3port2, abs4port2}, {abs5port1, abs6port1}},
	}
	_, err := Solve(ctx, requestGraph, superGraphTest)
	if err == nil {
		t.Error("Solve got nil error, want error due to cancel")
	}
}

// Benchmark solve for 4 DUT
// Run via blaze test with --test_arg=--test.bench=. --nocache_test_results
func Benchmark4DutSolveSwitch(b *testing.B) {
	b.ReportAllocs()
	requestGraph := &AbstractGraph{
		Desc:  "four nodes, interconnected with Switch",
		Nodes: []*AbstractNode{abs3, abs4, abs5, abs6},
		Edges: []*AbstractEdge{{abs3port1, abs4port1}, {abs3port2, abs4port2}, {abs5port1, abs6port1}},
	}
	for i := 0; i < b.N; i++ {
		Solve(context.Background(), requestGraph, superGraphTest)
	}
}
