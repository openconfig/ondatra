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

package ondatra

import (
	"fmt"
	"testing"

	"github.com/openconfig/ondatra/internal/debugger"
)

// Debug returns the debugger API.
func Debug() *Debugger {
	return &Debugger{}
}

// Debugger is the debugger API.
type Debugger struct{}

// Breakpoint inserts a breakpoint in the test and prints fmt.Sprint(a...) in
// the breakpoint message.
func (d *Debugger) Breakpoint(t *testing.T, a ...any) {
	t.Helper()
	if err := debugger.Breakpoint(t, fmt.Sprint(a...)); err != nil {
		t.Fatalf("Breakpoint(t): %v", err)
	}
}

// Breakpointf inserts a breakpoint in the test and prints fmt.Sprintf(a...) in
// the breakpoint message.
func (d *Debugger) Breakpointf(t *testing.T, format string, a ...any) {
	t.Helper()
	if err := debugger.Breakpoint(t, fmt.Sprintf(format, a...)); err != nil {
		t.Fatalf("Breakpointf(t): %v", err)
	}
}
