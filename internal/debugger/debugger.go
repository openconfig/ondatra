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

// Package debugger provides debugging utilities.
package debugger

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/testbed"
)

var (
	writer       = bufio.NewWriter(os.Stdout)
	reader       *bufio.Reader
	reservePause bool

	// To be stubbed out by tests.
	readStringFn  = (*bufio.Reader).ReadString
	reservationFn = testbed.Reservation
)

// TestStarted notifies the debugger that the test has started and whether debug
// mode should be enabled.
func TestStarted(debugMode bool) {
	if debugMode {
		reader = bufio.NewReader(os.Stdin)
		showMenu("Welcome to Ondatra Debug Mode!")
	}
	logMain(actionMsg("Reserving the testbed"))
}

func showMenu(msg string) {
	logMain(bannerMsg(
		msg,
		"",
		"Choose from one of the following options:",
		"[1] Run the full test with breakpoints enabled",
		"[2] Just reserve the testbed for now",
		"",
		"Then press ENTER to continue or CTRL-C to quit.",
	))
	readMenuOption()
}

func readMenuOption() {
	option := strings.TrimSpace(readLine())
	switch option {
	case "1":
	case "2":
		reservePause = true
	default:
		showMenu(fmt.Sprintf("Invalid menu option: %q", option))
	}
}

// TestCasesDone notifies the debugger that the test cases are completed, and
// the testbed is about to be released.
func TestCasesDone() {
	logMain(actionMsg("Releasing the testbed"))
}

// ReservationComplete notifes the debugger that the reservation is complete.
func ReservationComplete() {
	res, err := reservationFn()
	if err != nil {
		log.Fatalf("failed to fetch the testbed reservation: %v", err)
		return
	}
	lines := []string{
		"Testbed Reservation Complete",
		fmt.Sprintf("ID: %s\n", res.ID),
	}
	addLine := func(format string, args ...interface{}) {
		lines = append(lines, fmt.Sprintf(format, args...))
	}
	addAssign := func(id, name string) {
		addLine("  %-17s %s", id+":", name)
	}
	for id, d := range res.DUTs {
		addAssign(id, d.Name)
		for pid, p := range d.Ports {
			addAssign(pid, p.Name)
		}
	}
	for id, a := range res.ATEs {
		addAssign(id, a.Name)
		for pid, p := range a.Ports {
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

	logMain(bannerMsg(lines...))
	if reservePause {
		readLine()
	}
}

// LoggerT is a minimal subset of the testing.T API for logging only.
type LoggerT interface {
	Helper()
	Log(...interface{})
}

// ActionStarted notifies the debugger that the specific action has started.
func ActionStarted(t LoggerT, format string, dev binding.Device) {
	t.Helper()
	t.Log(actionMsg(fmt.Sprintf(format, dev.Dimensions().Name)))
}

// Breakpoint suspends the execution until the user presses enter.
// Returns an error when the test is not run in debug mode.
func Breakpoint(t LoggerT) error {
	if reader == nil {
		return errors.New("Breakpoints are only allowed in debug mode")
	}
	t.Log(bannerMsg("Breakpoint: Press ENTER to continue."))
	readLine()
	return nil
}

func logMain(msg string) {
	writer.WriteString(msg)
	writer.Flush()
}

func bannerMsg(lines ...string) string {
	const (
		format = `
********************************************************************************

%s
********************************************************************************

`
		indent = "  "
	)
	b := new(strings.Builder)
	for _, ln := range lines {
		b.WriteString(indent + ln + "\n")
	}
	return fmt.Sprintf(format, b.String())
}

func actionMsg(action string) string {
	return fmt.Sprintf("\n*** %s...\n\n", action)
}

func readLine() string {
	ln, err := readStringFn(reader, '\n')
	if err != nil {
		// An error when reading from stdin is limited in practice to a user
		// pressing CTRL-C at a prompt. Logging this error would just be noise.
		return ""
	}
	return ln
}
