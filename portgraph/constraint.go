// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package portgraph

import (
	"regexp"

	log "github.com/golang/glog"
)

// NodeConstraint is a constraint on a Node.
type NodeConstraint interface {
	isNodeConstraint()
}

// PortConstraint is a constraint on a Port.
type PortConstraint interface {
	isPortConstraint()
	fetchPort() *AbstractPort
}

// Constraint is a constraint checker that checks an input matches based on some matching criteria.
type Constraint interface {
	universalConstraint
	match(string, bool) bool
}

type universalConstraint interface {
	NodeConstraint
	PortConstraint
}

type equal struct {
	universalConstraint
	s string
}

func (c *equal) match(s string, ok bool) bool {
	return ok && s == c.s
}

type notEqual struct {
	universalConstraint
	s string
}

func (c *notEqual) match(s string, ok bool) bool {
	return ok && s != c.s
}

type regex struct {
	universalConstraint
	re *regexp.Regexp
}

func (c *regex) match(s string, ok bool) bool {
	return ok && c.re.MatchString(s)
}

type notRegex struct {
	universalConstraint
	re *regexp.Regexp
}

func (c *notRegex) match(s string, ok bool) bool {
	return ok && !c.re.MatchString(s)
}

func attrsNodePair(k string, n1, n2 *AbstractNode, abs2ConNode map[*AbstractNode]*ConcreteNode) (string, string, bool) {
	cn1, ok := abs2ConNode[n1]
	if !ok {
		log.Fatalf("node %q does not have an assignment", n1.Desc)
	}
	v1, ok := cn1.Attrs[k]
	if !ok {
		return "", "", false
	}
	cn2, ok := abs2ConNode[n2]
	if !ok {
		log.Fatalf("node %q does not have an assignment", n2.Desc)
	}
	v2, ok := cn2.Attrs[k]
	if !ok {
		return "", "", false
	}
	return v1, v2, ok
}

func attrsPortPair(k string, p1, p2 *AbstractPort, abs2ConPort map[*AbstractPort]*ConcretePort) (string, string, bool) {
	cp1, ok := abs2ConPort[p1]
	if !ok {
		log.Fatalf("port %q does not have an assignement", p1.Desc)
	}
	v1, ok := cp1.Attrs[k]
	if !ok {
		return "", "", false
	}
	cp2, ok := abs2ConPort[p2]
	if !ok {
		log.Fatalf("port %q does not have an assignment", p2.Desc)
	}
	v2, ok := cp2.Attrs[k]
	if !ok {
		return "", "", false
	}
	return v1, v2, ok
}

type sameAsNode struct {
	n *AbstractNode
}

func (c *sameAsNode) isNodeConstraint() {}

func (c *sameAsNode) match(k string, n *AbstractNode, abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
	v1, v2, ok := attrsNodePair(k, n, c.n, abs2ConNode)
	return ok && v1 == v2
}

type sameAsPort struct {
	p *AbstractPort
}

func (c *sameAsPort) isPortConstraint() {}

func (c *sameAsPort) fetchPort() *AbstractPort { return c.p }

func (c *sameAsPort) match(k string, p *AbstractPort, abs2ConPort map[*AbstractPort]*ConcretePort) bool {
	v1, v2, ok := attrsPortPair(k, p, c.p, abs2ConPort)
	return ok && v1 == v2
}

type notSameAsNode struct {
	n *AbstractNode
}

func (c *notSameAsNode) isNodeConstraint() {}

func (c *notSameAsNode) match(k string, n *AbstractNode, abs2ConNode map[*AbstractNode]*ConcreteNode) bool {
	v1, v2, ok := attrsNodePair(k, n, c.n, abs2ConNode)
	return ok && v1 != v2
}

type notSameAsPort struct {
	p *AbstractPort
}

func (c *notSameAsPort) isPortConstraint() {}

func (c *notSameAsPort) fetchPort() *AbstractPort { return c.p }

func (c *notSameAsPort) match(k string, p *AbstractPort, abs2ConPort map[*AbstractPort]*ConcretePort) bool {
	v1, v2, ok := attrsPortPair(k, p, c.p, abs2ConPort)
	return ok && v1 != v2
}

// Equal returns a constraint that an attribute is equal to the specified string.
func Equal(s string) Constraint {
	return &equal{s: s}
}

// NotEqual returns a constraint that an attribute is not equal to the specified string.
func NotEqual(s string) Constraint {
	return &notEqual{s: s}
}

// Regex returns a constraint that an attribute matches the specified regex.
func Regex(re *regexp.Regexp) Constraint {
	return &regex{re: re}
}

// NotRegex returns a constraint that an attribute does not match the specified regex.
func NotRegex(re *regexp.Regexp) Constraint {
	return &notRegex{re: re}
}

// SameAsNode returns a constraint that an attribute has the same value as the specified Node.
func SameAsNode(n *AbstractNode) NodeConstraint {
	return &sameAsNode{n}
}

// NotSameAsNode returns a constraint that an attribute has a different value from the specified Node.
func NotSameAsNode(n *AbstractNode) NodeConstraint {
	return &notSameAsNode{n}
}

// SameAsPort returns a constraint that an attribute has the same value as the specified Port.
func SameAsPort(p *AbstractPort) PortConstraint {
	return &sameAsPort{p}
}

// NotSameAsPort returns a constraint that an attribute has a different value from the specified Port.
func NotSameAsPort(p *AbstractPort) PortConstraint {
	return &notSameAsPort{p}
}
