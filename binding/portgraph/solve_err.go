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

// solveReport generates a report with all the events that happened during a solve.
type solveReport struct {
	report []string
}

func (s *solveReport) String() string {
	report := &strings.Builder{}
	fmt.Fprintln(report, "Solve report:")
	for _, r := range s.report {
		fmt.Fprintln(report, r)
	}
	return report.String()
}

func (s *solveReport) TryAssign(device string, node string) {
	s.report = append(s.report, fmt.Sprintf("[TRY_ASSIGN]: Trying to assign %s to %s", device, node))
}

func (s *solveReport) Assigned(device string, node string, constraints map[string]string) {
	desc := &strings.Builder{}
	fmt.Fprintf(desc, "[ASSIGNED]: Assigned %s to %s", device, node)
	if len(constraints) > 0 {
		fmt.Fprintf(desc, " with constraints")
		for k, v := range constraints {
			fmt.Fprintf(desc, " %s: %s", k, v)
		}
	}
	s.report = append(s.report, desc.String())
}

func (s *solveReport) Unassigned(device string, node string) {
	s.report = append(s.report, fmt.Sprintf("[UNASSIGNED]: Unassigned %s from %s", device, node))
}

func (s *solveReport) ConstraintFailed(device string, node string, constraint string, got string, want string) {
	s.report = append(s.report, fmt.Sprintf("[CONSTRAINT_FAIL]: %s does not match required constraint %s for %s got: %s want: %s", device, constraint, node, got, want))
}

// solveError implements error and contains information about a call to Solve.
type solveError struct {
	absGraphDesc string
	conGraphDesc string
	wrappedErr   error
	maxAssign    *Assignment
	report       *solveReport
}

func (s *solveError) Unwrap() error {
	return s.wrappedErr
}

// Error returns and error string. This function implements error.
func (s *solveError) Error() string { return s.String() }

// String compiles SolveErr to a string format.
func (s *solveError) String() string {
	ret := &strings.Builder{}
	fmt.Fprintf(ret, "Could not satisfy %q from %q", s.absGraphDesc, s.conGraphDesc)
	if s.wrappedErr != nil {
		fmt.Fprintf(ret, ": %v", s.wrappedErr)
	}
	if s.maxAssign != nil {
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
	}
	return ret.String()
}
