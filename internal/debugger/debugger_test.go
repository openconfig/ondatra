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
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/openconfig/ondatra/binding"
)

var stubLine string

func init() {
	readStringFn = func(_ *bufio.Reader, b byte) (string, error) {
		if i := strings.IndexByte(stubLine, b); i < 0 {
			return "", fmt.Errorf("cannot read line %q", stubLine)
		}
		line := stubLine
		stubLine = ""
		return line, nil
	}
	reservationFn = func() (*binding.Reservation, error) {
		return &binding.Reservation{}, nil
	}
}

func TestNoDebugMode(t *testing.T) {
	TestStarted(false)

	if err := Breakpoint(t); err == nil {
		t.Fatal("Breakpoint unexpectedly succeeded")
	}
}

func TestDebugModeRunTests(t *testing.T) {
	stubLine = "1\n"
	TestStarted(true)

	stubLine = "\n"
	if err := Breakpoint(t); err != nil {
		t.Fatalf("Breakpoint got error %v", err)
	}
}

func TestDebugModeJustReserve(t *testing.T) {
	stubLine = "2\n"
	TestStarted(true)

	stubLine = "\n"
	ReservationComplete()

	stubLine = "\n"
	if err := Breakpoint(t); err != nil {
		t.Fatalf("Breakpoint got error %v", err)
	}
}
