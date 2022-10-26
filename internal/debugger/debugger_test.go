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

package debugger

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/openconfig/ondatra/binding"
)

var (
	stubMu   sync.Mutex
	stubLine string
)

func resetTTYReader() {
	reader = nil
}

func writeLine(line string) {
	stubMu.Lock()
	defer stubMu.Unlock()
	stubLine = line
}

func init() {
	readString := func(b byte) (string, error) {
		stubMu.Lock()
		defer stubMu.Unlock()
		if i := strings.IndexByte(stubLine, b); i < 0 {
			return "", fmt.Errorf("cannot read line %q", stubLine)
		}
		return stubLine, nil
	}
	close := func() error { return nil }
	openTTYFn = func() (readStringFn, closeFn, error) {
		return readString, close, nil
	}
	reservationFn = func() (*binding.Reservation, error) {
		return &binding.Reservation{}, nil
	}
}
func TestNoDebugMode(t *testing.T) {
	resetTTYReader()
	TestStarted(false)
	defer TestCasesDone()

	if err := Breakpoint(t, ""); err == nil {
		t.Fatal("Breakpoint unexpectedly succeeded")
	}
	if err := Breakpoint(t, "msg"); err == nil {
		t.Fatal("Breakpoint unexpectedly succeeded")
	}
}

func TestDebugModeRunTests(t *testing.T) {
	resetTTYReader()
	writeLine("1\n")
	TestStarted(true)
	defer TestCasesDone()

	writeLine("\n")
	if err := Breakpoint(t, ""); err != nil {
		t.Fatalf("Breakpoint got error %v", err)
	}
	if err := Breakpoint(t, "msg"); err != nil {
		t.Fatalf("Breakpoint got error %v", err)
	}
}

func TestDebugModeJustReserve(t *testing.T) {
	resetTTYReader()
	writeLine("2\n")
	TestStarted(true)
	defer TestCasesDone()

	writeLine("\n")
	ReservationDone()

	writeLine("\n")
	if err := Breakpoint(t, ""); err != nil {
		t.Fatalf("Breakpoint got error %v", err)
	}
	if err := Breakpoint(t, "msg"); err != nil {
		t.Fatalf("Breakpoint got error %v", err)
	}
}
