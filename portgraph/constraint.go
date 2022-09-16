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
)

// NodeConstraint is a constraint on a Node.
type NodeConstraint interface {
	isNodeConstraint()
}

// PortConstraint is a constraint on a Port.
type PortConstraint interface {
	isPortConstraint()
}

// Constraint is a constraint checker that checks an input matches based on some matching criteria.
type Constraint interface {
	universalConstraint
	match(s string) bool
}

type universalConstraint interface {
	NodeConstraint
	PortConstraint
}

type equal struct {
	universalConstraint
	s string
}

func (c *equal) match(s string) bool {
	return s == c.s
}

type notEqual struct {
	universalConstraint
	s string
}

func (c *notEqual) match(s string) bool {
	return s != c.s
}

type regex struct {
	universalConstraint
	re *regexp.Regexp
}

func (c *regex) match(s string) bool {
	return c.re.MatchString(s)
}

type notRegex struct {
	universalConstraint
	re *regexp.Regexp
}

func (c *notRegex) match(s string) bool {
	return !c.re.MatchString(s)
}

type sameAsNode struct {
	n *AbstractNode
}

func (c *sameAsNode) isNodeConstraint() {}

type sameAsPort struct {
	p *AbstractPort
}

func (c *sameAsPort) isPortConstraint() {}

type notSameAsNode struct {
	n *AbstractNode
}

func (c *notSameAsNode) isNodeConstraint() {}

type notSameAsPort struct {
	p *AbstractPort
}

func (c *notSameAsPort) isPortConstraint() {}

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
