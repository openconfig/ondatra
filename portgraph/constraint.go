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

// Constraint is a constraint checker that checks an input matches based on some matching criteria.
type Constraint interface {
	match(s string) bool
}

type equal struct {
	s string
}

func (c *equal) match(s string) bool {
	return s == c.s
}

type notEqual struct {
	s string
}

func (c *notEqual) match(s string) bool {
	return s != c.s
}

type regex struct {
	re *regexp.Regexp
}

func (c *regex) match(s string) bool {
	return c.re.MatchString(s)
}

type notRegex struct {
	re *regexp.Regexp
}

func (c *notRegex) match(s string) bool {
	return !c.re.MatchString(s)
}

// Equal returns a constraint that an attribute is equal to the specified string.
func Equal(s string) Constraint {
	return &equal{s}
}

// NotEqual returns a constraint that an attribute is not equal to the specified string.
func NotEqual(s string) Constraint {
	return &notEqual{s}
}

// Regex returns a constraint that an attribute matches the specified regex.
func Regex(re *regexp.Regexp) Constraint {
	return &regex{re}
}

// NotRegex returns a constraint that an attribute does not match the specified regex.
func NotRegex(re *regexp.Regexp) Constraint {
	return &notRegex{re}
}
