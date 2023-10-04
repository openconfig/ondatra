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

// Package debug provides an API to add breakpoints to the test to debug its
// execution.
//
// The functions defined in this package must only be used for local debugging
// and require that the test be run in [Debug Mode]. If a breakpoint is reached
// when not in debug mode, the test will fail there and then. Therefore, no
// tests with breakpoints should be committed to any persistent repository
//
// To insert a simple breakpoint in your code, use:
//
//	ondatra.Debug().Breakpoint(t)
//
// For more informative breakpoints, the [Debug.Breakpoint] and
// [Debug.Breakpoint] functions allow you to include custom text in the
// breakpoint message:
//
//	ondatra.Debug().Breakpoint(t, "this should be unreachable")
//	ondatra.Debug().Breakpointf(t, "myVar has value %v", myVar)
//
// [Debug Mode]: https://github.com/openconfig/ondatra#debugging-an-ondatra-test
package debug

import (
	"fmt"
	"testing"

	"github.com/openconfig/ondatra/internal/events"
)

// Debug is the Ondatra debug API.
type Debug struct{}

// Breakpoint inserts a breakpoint in the test and prints fmt.Sprint(a...) in
// the breakpoint message.
func (d *Debug) Breakpoint(t *testing.T, a ...any) {
	t.Helper()
	if err := events.Breakpoint(t, fmt.Sprint(a...)); err != nil {
		t.Fatalf("Breakpoint(t): %v", err)
	}
}

// Breakpointf inserts a breakpoint in the test and prints fmt.Sprintf(a...) in
// the breakpoint message.
func (d *Debug) Breakpointf(t *testing.T, format string, a ...any) {
	t.Helper()
	if err := events.Breakpoint(t, fmt.Sprintf(format, a...)); err != nil {
		t.Fatalf("Breakpointf(t): %v", err)
	}
}
