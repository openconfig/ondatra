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

// Package events prints and responds to test events.
package events

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"

	log "github.com/golang/glog"
	closer "github.com/openconfig/gocloser"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/display"
	"github.com/openconfig/ondatra/internal/testbed"
)

var (
	reservePause bool
	running      atomic.Bool
	beforeTests  []func(*binding.Reservation) error
	afterTests   []func(*int) error

	// To be stubbed out by tests.
	startReaderFn   = display.StartReader
	readerStartedFn = display.ReaderStarted
	menuFn          = display.Menu
	readLineFn      = display.ReadLine
	reservationFn   = testbed.Reservation
)

// AddBeforeTests adds a callback to run before the tests start executing.
func AddBeforeTests(cb func(*binding.Reservation) error) {
	checkNotRunning()
	beforeTests = append(beforeTests, cb)
}

// AddAfterTests adds a callback to run after the tests finish executing.
func AddAfterTests(cb func(*int) error) {
	checkNotRunning()
	afterTests = append(afterTests, cb)
}

func runBeforeTestsCallbacks(res *binding.Reservation) []error {
	checkRunning()
	var errs []error
	for _, cb := range beforeTests {
		if err := cb(res); err != nil {
			errs = append(errs, callbackErr(err, cb))
		}
	}
	return errs
}

func runAfterTestsCallbacks(exitCode *int) []error {
	checkRunning()
	var errs []error
	for _, cb := range afterTests {
		if err := cb(exitCode); err != nil {
			errs = append(errs, callbackErr(err, cb))
		}
	}
	return errs
}

func checkRunning() {
	if !running.Load() { // impossibility check
		log.Fatalf("Callbacks can only be run while test is running.")
	}
}

func checkNotRunning() {
	if running.Load() { // precondition check
		log.Exitf("All callbacks must be added before test is running.")
	}
}

func callbackErr(err error, cb any) error {
	name := runtime.FuncForPC(reflect.ValueOf(cb).Pointer()).Name()
	if strings.HasSuffix(name, "-fm") {
		name = name[:len(name)-3]
	}
	return fmt.Errorf("error on callback %q: %v", name, err)
}

// TestStarted notifies that the test has started, whether it was started in
// debug mode, and that the testbed is about to be reserved.
func TestStarted(debugMode bool) {
	running.Store(true)
	if debugMode {
		if err := startReaderFn(); err != nil {
			log.Exitf("Error starting stdin reader: %v", err)
		}
		choice := menuFn("Welcome to Ondatra Debug Mode!",
			"Run the full test with breakpoints enabled",
			"Just reserve the testbed for now")
		if choice == 2 {
			reservePause = true
		}
	}
	display.Action(nil, "Reserving the testbed")
}

// ReservationDone notifies that the reservation is complete and that the tests
// are about to begin execution and runs all the BeforeTestsCallbacks.
func ReservationDone() error {
	res, err := reservationFn()
	if err != nil {
		return fmt.Errorf("failed to fetch the testbed reservation: %v", err)
	}

	lines := []string{
		"Testbed Reservation Complete",
		fmt.Sprintf("ID: %s\n", res.ID),
	}
	addLine := func(format string, args ...any) {
		lines = append(lines, fmt.Sprintf(format, args...))
	}
	addAssign := func(id, name string) {
		addLine("  %-17s %s", id+":", name)
	}
	for id, d := range res.DUTs {
		addAssign(id, d.Name())
		for pid, p := range d.Ports() {
			addAssign(pid, p.Name)
		}
	}
	for id, a := range res.ATEs {
		addAssign(id, a.Name())
		for pid, p := range a.Ports() {
			addAssign(pid, p.Name)
		}
	}

	if reservePause {
		lines = append(lines,
			"",
			"To reuse this reservation for another test execution, run",
			fmt.Sprintf("  go test <TEST_NAME> --reserve=%s", res.ID),
			"",
			"Press CTRL-C to release the reservation or ENTER to run the test cases.",
		)
	}

	display.Banner(nil, lines...)
	if reservePause {
		readLineFn()
	}

	if errs := runBeforeTestsCallbacks(res); len(errs) > 0 {
		return fmt.Errorf("errors running BeforeTestsCallbacks: %v", errs)
	}
	return nil
}

// TestsDone notifies that the test cases are complete and that the testbed is
// about to be released and runs all the AfterTestsCallbacks.
func TestsDone(exitCode *int) (rerr error) {
	defer closer.Close(&rerr, display.StopReader, "error stopping display reader")
	if errs := runAfterTestsCallbacks(exitCode); len(errs) > 0 {
		return fmt.Errorf("errors running AfterTestsCallbacks: %v", errs)
	}
	display.Action(nil, "Releasing the testbed")
	return nil
}

// ActionStarted notifies that the specified action has started.
// Used to restrict the library to calling t.Helper and t.Log only.
func ActionStarted(t testing.TB, format string, dev binding.Device) testing.TB {
	t.Helper()
	display.Action(t, fmt.Sprintf(format, dev.Name()))
	if readerStartedFn() {
		return &breakpointT{t}
	}
	return t
}

// Breakpoint notifies a breakpoint has been reached, which suspends test
// execution until the user indicates test execution should be resumed.
// Returns an error if the test is not in debug mode.
func Breakpoint(t testing.TB, msg string) error {
	if !readerStartedFn() {
		return errors.New("Breakpoints are only allowed in debug mode")
	}
	t.Helper()
	firstLine := "BREAKPOINT"
	if msg != "" {
		firstLine += ": " + msg
	}
	display.Banner(t, firstLine, "", "Press ENTER to continue.")
	readLineFn()
	return nil
}

type breakpointT struct {
	testing.TB
}

func (t *breakpointT) Error(args ...any) {
	Breakpoint(t, "TEST FAILED\n"+fmt.Sprint(args...))
	t.TB.Error(args...)
}

func (t *breakpointT) Errorf(format string, args ...any) {
	Breakpoint(t, "TEST FAILED\n"+fmt.Sprintf(format, args...))
	t.TB.Errorf(format, args...)
}

func (t *breakpointT) Fail() {
	Breakpoint(t, "TEST FAILED")
	t.TB.Fail()
}

func (t *breakpointT) FailNow() {
	Breakpoint(t, "TEST FAILED")
	t.TB.FailNow()
}

func (t *breakpointT) Fatal(args ...any) {
	Breakpoint(t, "TEST FAILED\n"+fmt.Sprint(args...))
	t.TB.Fatal(args...)
}

func (t *breakpointT) Fatalf(format string, args ...any) {
	Breakpoint(t, "TEST FAILED\n"+fmt.Sprintf(format, args...))
	t.TB.Fatalf(format, args...)
}
