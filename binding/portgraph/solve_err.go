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
	"sort"
	"strings"
	"time"
)

var (
	// indent is the number of spaces to indent the report.
	// 4 spaces for each indent level
	indent = "    "
)

func mapToSortedString(prefix string, m map[string]string, separator string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	ret := &strings.Builder{}
	ret.WriteString(prefix)
	for i, k := range keys {
		if i > 0 {
			ret.WriteString(", ")
		}
		ret.WriteString(k + separator + m[k])
	}
	return ret.String()
}

// Event is a step that happened during the solve process.
type Event struct {
	description    string
	PrecedingEvent *Event
}

// solveReport generates a report with all the events that happened during a solve.
type solveReport struct {
	report    []*Event
	startTime time.Time
	timeTaken time.Duration
}

// appendEvent creates a process event.
func (s *solveReport) appendEvent(description string, precedingEvent *Event) *Event {
	event := &Event{
		description:    description,
		PrecedingEvent: precedingEvent,
	}
	s.report = append(s.report, event)
	return event
}

// String returns all events of the solve report as a string.
func (s *solveReport) String() string {
	report := &strings.Builder{}
	report.WriteString("Solve started at " + s.startTime.String() + "\n")
	for _, event := range s.report {
		if event.PrecedingEvent == nil {
			s.printEvent(report, event, 0)
		}
	}
	report.WriteString("\nSolve completed in " + s.timeTaken.String() + "\n")
	return report.String()
}

// printEvent prints the event and its subsequent events.
func (s *solveReport) printEvent(report *strings.Builder, event *Event, indentLevel int) {
	for i := 0; i < indentLevel; i++ {
		report.WriteString(indent)
	}
	report.WriteString(event.description + "\n")

	// Recursively print subsequent events
	for _, e := range s.report {
		if e.PrecedingEvent == event {
			s.printEvent(report, e, indentLevel+1)
		}
	}
}

// AddPhase adds a new phase to the solve report.
func (s *solveReport) AddPhase(phaseString string) *Event {
	var desc strings.Builder
	desc.WriteString("\n[PHASE] " + phaseString + "\n")
	return s.appendEvent(desc.String(), nil)
}

// AppendTryAssignGraph appends an event for trying to assign a graph.
func (s *solveReport) AppendTryAssignGraph(abs2ConNode map[*AbstractNode]*ConcreteNode, precedingEvent *Event) *Event {
	var desc strings.Builder
	graph := make(map[string]string)
	for absNode, conNode := range abs2ConNode {
		graph[absNode.Desc] = conNode.Desc
	}
	desc.WriteString("[TRY_ASSIGN_GRAPH]: Trying to assign topology" + mapToSortedString(" with nodes ", graph, ": "))
	return s.appendEvent(desc.String(), precedingEvent)
}

// AppendEvaluateNode appends an event for evaluating a node.
func (s *solveReport) AppendEvaluateNode(device string, node string, constraints map[string]string, precedingEvent *Event) *Event {
	var desc strings.Builder
	desc.WriteString("[EVALUATE_NODE]: Evaluating " + device + " for " + node)
	if len(constraints) > 0 {
		desc.WriteString(mapToSortedString(" with constraints ", constraints, " "))
	}
	return s.appendEvent(desc.String(), precedingEvent)
}

// AppendMatch appends an event for a node that matched.
func (s *solveReport) AppendMatch(device string, node string, precedingEvent *Event) *Event {
	var desc strings.Builder
	desc.WriteString("[MATCH]: Matched " + device + " to " + node)
	return s.appendEvent(desc.String(), precedingEvent)
}

// AppendConstraintFailed appends an event for a node that failed a constraint.
func (s *solveReport) AppendConstraintFailed(device string, node string, constraint string, got string, want string, precedingEvent *Event) *Event {
	var desc strings.Builder
	desc.WriteString("[CONSTRAINT_FAIL]: " + device + " cannot be assigned to " + node + "; reservation request requires " + constraint + " " + want + " but " + device + " " + constraint + " is " + got)
	return s.appendEvent(desc.String(), precedingEvent)
}

// AppendConstraintFailedString appends an event for a node that failed a constraint.
func (s *solveReport) AppendConstraintFailedString(device string, node string, reason string, precedingEvent *Event) *Event {
	var desc strings.Builder
	desc.WriteString("[CONSTRAINT_FAIL]: " + device + " cannot be assigned to " + node + "; " + reason)
	return s.appendEvent(desc.String(), precedingEvent)
}

// SolveError implements error and contains information about a call to Solve.
type SolveError struct {
	absGraphDesc string
	conGraphDesc string
	wrappedErr   error
	maxAssign    *Assignment
	report       *solveReport
}

func (s *SolveError) Unwrap() error {
	return s.wrappedErr
}

// Error returns and error string. This function implements error.
func (s *SolveError) Error() string { return s.String() }

// String compiles SolveErr to a string format.
func (s *SolveError) String() string {
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

// Report returns the solve report as a string.
func (s *SolveError) Report() string {
	return s.report.String()
}
