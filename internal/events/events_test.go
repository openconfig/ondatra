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

package events

import (
	"testing"

	"github.com/openconfig/ondatra/binding"
)

var (
	exitCode = func() *int {
		code := 0
		return &code
	}()

	// To be stubbed out by tests.
	readerStartedStub bool
	menuChoiceStub    int
)

func init() {
	startReaderFn = func() error {
		readerStartedStub = true
		return nil
	}
	readerStartedFn = func() bool {
		return readerStartedStub
	}
	menuFn = func(msg string, options ...string) int {
		return menuChoiceStub
	}
	readLineFn = func() string {
		return ""
	}
	reservationFn = func() (*binding.Reservation, error) {
		return &binding.Reservation{}, nil
	}
}

func resetStubs() {
	readerStartedStub = false
	menuChoiceStub = 0
}

func TestNoDebugMode(t *testing.T) {
	resetStubs()
	TestStarted(false)
	defer TestsDone(exitCode)

	if err := Breakpoint(t, ""); err == nil {
		t.Fatal("Breakpoint unexpectedly succeeded")
	}
	if err := Breakpoint(t, "msg"); err == nil {
		t.Fatal("Breakpoint unexpectedly succeeded")
	}
}

func TestDebugModeRunTests(t *testing.T) {
	resetStubs()
	menuChoiceStub = 2
	TestStarted(true)
	defer TestsDone(exitCode)

	if err := Breakpoint(t, ""); err != nil {
		t.Fatalf("Breakpoint got error %v", err)
	}
	if err := Breakpoint(t, "msg"); err != nil {
		t.Fatalf("Breakpoint got error %v", err)
	}
}

func TestDebugModeJustReserve(t *testing.T) {
	resetStubs()
	menuChoiceStub = 1
	TestStarted(true)
	defer TestsDone(exitCode)
	ReservationDone()

	if err := Breakpoint(t, ""); err != nil {
		t.Fatalf("Breakpoint got error %v", err)
	}
	if err := Breakpoint(t, "msg"); err != nil {
		t.Fatalf("Breakpoint got error %v", err)
	}
}
