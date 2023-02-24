// Copyright 2023 Google LLC
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

// Package display prints and receives input from the console.
package display

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"

	log "github.com/golang/glog"
)

var (
	writer = bufio.NewWriter(os.Stdout)
	reader *ttyReader
)

// Menu displays a menu with the specified message and list of options and
// returns the 1-indexed option number that the user selected.
func Menu(msg string, options ...string) int {
	menuLines := []string{
		msg,
		"",
		"Choose from one of the following options:",
	}
	for i, option := range options {
		menuLines = append(menuLines, fmt.Sprintf("[%d] %s", i+1, option))
	}
	menuLines = append(menuLines, "", "Then press ENTER to continue or CTRL-C to quit.")

	Banner(nil, menuLines...)
	return readMenuOption(options)
}

func readMenuOption(options []string) int {
	option := strings.TrimSpace(ReadLine())
	num, err := strconv.Atoi(option)
	if err != nil || num < 0 && num > len(options) {
		return Menu(fmt.Sprintf("Invalid menu option: %q", option), options...)
	}
	return num
}

// Banner displays a banner message with the specified lines.
// If t is not nil, the message is associated with the current test case.
func Banner(t testing.TB, lines ...string) {
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
	display(t, fmt.Sprintf(format, b.String()))
}

// Action displays the specified action message.
// If t is not nil, the message is associated with the current test case.
func Action(t testing.TB, action string) {
	display(t, fmt.Sprintf("\n*** %s...\n\n", action))
}

func display(t testing.TB, msg string) {
	if t == nil {
		writer.WriteString(msg)
		writer.Flush()
	} else {
		t.Helper()
		t.Log(msg)
	}
}

// StartReader starts a stdin reader.
func StartReader() error {
	path := "/dev/tty"
	if runtime.GOOS == "windows" {
		path = "CONIN$"
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(file)
	reader = &ttyReader{readFn: buf.ReadString, closeFn: file.Close}
	reader.start()
	return nil
}

// ReaderStarted returns whether or not the reader has started.
func ReaderStarted() bool {
	return reader != nil
}

// StopReader stops the reader. Noop if it hasn't started.
func StopReader() error {
	if reader == nil {
		return nil
	}
	return reader.stop()
}

// ReadLine reads a line of user input.
// Exits fatally if the reader has not been started.
func ReadLine() string {
	if reader == nil {
		log.Fatalf("The reader has not been started")
	}
	return reader.readLine()
}

// ttyReader continuously reads lines from the controlling terminal. It ensures
// that user input entered prior to a user prompt is _not_ interpreted as a
// response to that prompt. As there is no easy way to clear prior user input,
// it reads asynchrously and is signalled when a prompt has been displayed.
type ttyReader struct {
	readFn  func(byte) (string, error)
	closeFn func() error

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

func (r *ttyReader) stop() error {
	r.mu.Lock()
	r.done = true
	r.mu.Unlock()
	return r.closeFn()
}
