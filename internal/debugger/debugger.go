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
	"runtime"
	"strings"
	"sync"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/testbed"
)

var (
	writer       = bufio.NewWriter(os.Stdout)
	reader       *ttyReader
	reservePause bool

	// To be stubbed out by tests.
	openTTYFn     = openTTY
	reservationFn = testbed.Reservation
)

// TestStarted notifies the debugger that the test has started and whether debug
// mode should be enabled.
func TestStarted(debugMode bool) {
	if debugMode {
		readFn, closeFn, err := openTTYFn()
		if err != nil {
			log.Exitf("No controlling terminal available for debug mode: %v", err)
		}
		reader = &ttyReader{readFn: readFn, closeFn: closeFn}
		reader.start()
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
	option := strings.TrimSpace(reader.readLine())
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
	if reader != nil {
		reader.stop()
	}
	logMain(actionMsg("Releasing the testbed"))
}

// ReservationDone notifes the debugger that the reservation is complete.
func ReservationDone() {
	res, err := reservationFn()
	if err != nil {
		log.Fatalf("failed to fetch the testbed reservation: %v", err)
		return
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

	logMain(bannerMsg(lines...))
	if reservePause {
		reader.readLine()
	}
}

// LoggerT is a minimal subset of the testing.T API for logging only.
type LoggerT interface {
	Helper()
	Log(...any)
}

// ActionStarted notifies the debugger that the specific action has started.
func ActionStarted(t LoggerT, format string, dev binding.Device) {
	t.Helper()
	t.Log(actionMsg(fmt.Sprintf(format, dev.Name())))
}

// Breakpoint suspends the execution until the user presses enter.
// Returns an error when the test is not run in debug mode.
func Breakpoint(t LoggerT, msg string) error {
	if reader == nil {
		return errors.New("Breakpoints are only allowed in debug mode")
	}
	t.Helper()
	firstLine := "BREAKPOINT"
	if msg != "" {
		firstLine += ": " + msg
	}
	t.Log(bannerMsg(firstLine, "", "Press ENTER to continue."))
	reader.readLine()
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

type readStringFn func(byte) (string, error)
type closeFn func() error

func openTTY() (readStringFn, closeFn, error) {
	path := "/dev/tty"
	if runtime.GOOS == "windows" {
		path = "CONIN$"
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	buf := bufio.NewReader(file)
	return buf.ReadString, file.Close, nil
}

// ttyReader continuously reads lines from the controlling terminal. It ensures
// that user input entered prior to a user prompt is _not_ interpreted as a
// response to that prompt. As there is no easy way to clear prior user input,
// it reads asynchrously and is signalled when a prompt has been displayed.
type ttyReader struct {
	readFn  readStringFn
	closeFn closeFn

	mu   sync.Mutex
	keep bool
	done bool
	ch   chan string
}

func (r *ttyReader) start() {
	r.ch = make(chan string)
	go func() {
		for done := false; !done; {
			line, err := r.readFn('\n')
			r.mu.Lock()
			done = r.done
			if err != nil {
				if done {
					r.mu.Unlock()
					break
				}
				log.Fatalf("Error reading from terminal: %v", err)
			}
			if r.keep {
				r.mu.Unlock()
				r.ch <- line
				r.mu.Lock()
			}
			r.mu.Unlock()
		}
	}()
}

func (r *ttyReader) readLine() string {
	r.mu.Lock()
	r.keep = true
	r.mu.Unlock()
	line := <-r.ch
	r.mu.Lock()
	r.keep = false
	r.mu.Unlock()
	return line
}

func (r *ttyReader) stop() {
	r.mu.Lock()
	r.done = true
	r.mu.Unlock()
	r.closeFn()
}
