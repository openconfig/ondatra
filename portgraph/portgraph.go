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

// Graph is a representation of nodes, ports, and edges.
type Graph struct {
	Desc  string // Description for the Graph for logging.
	Nodes []*Node
	Edges []*Edge
}

// Node represent a vertex on a graph.
// The Node has Ports and Edges that originate and terminate at Ports.
type Node struct {
	Desc  string            // Description for the Node for logging.
	Ports []*Port           // A list Ports that may be connected to another Port.
	Attrs map[string]string // A map of key value attributes of the Node.
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
