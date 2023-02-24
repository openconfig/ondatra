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
	"fmt"
	"strings"
)

// SolveErr implements error and contains information about a call to Solve.
type SolveErr struct {
	maxAssign                  *Assignment
	absGraphDesc, conGraphDesc string
}

// Error returns and error string. This function implements error.
func (s *SolveErr) Error() string { return s.String() }

// String compiles SolveErr to a string format.
func (s *SolveErr) String() string {
	ret := &strings.Builder{}
	fmt.Fprintf(ret, "Could not satisfy %q from %q\n", s.absGraphDesc, s.conGraphDesc)
	fmt.Fprintf(ret, "\nMax assignment:\n")
	for a, c := range s.maxAssign.Node2Node {
		if c != nil {
			fmt.Fprintf(ret, "Node %q is assigned to %q\n", a.Desc, c.Desc)
		} else {
			fmt.Fprintf(ret, "Node %q was not assigned\n", a.Desc)
		}
	}
	for a, c := range s.maxAssign.Port2Port {
		if c != nil {
			fmt.Fprintf(ret, "Port %q is assigned to %q\n", a.Desc, c.Desc)
		} else {
			fmt.Fprintf(ret, "Port %q was not assigned\n", a.Desc)
		}
	}
	return ret.String()
}
