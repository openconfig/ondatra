// Copyright 2019 Google LLC
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

// Package ondatra represents the public Ondatra testing API.
package ondatra

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"testing"
	"time"

	log "github.com/golang/glog"
	"golang.org/x/sys/unix"
	"github.com/openconfig/gocloser"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/debugger"
	"github.com/openconfig/ondatra/internal/flags"
	"github.com/openconfig/ondatra/internal/junitxml"
	"github.com/openconfig/ondatra/internal/testbed"
)

var (
	sigc         = make(chan os.Signal, 1)
	reserveFn    = reserve
	releaseFn    = release
	runFn        = (*testing.M).Run
	flagParseFn  = flags.Parse
	setBindingFn = testbed.SetBinding
)

// Binder creates the binding for the test.
type Binder func() (binding.Binding, error)

// RunTests acquires the testbed of devices and runs the tests. Every device is
// initialized with a baseline configuration that allows it to be managed.
func RunTests(m *testing.M, binder Binder) {
	// Careful to only exit at the very end, because exiting skips all pending defers.
	if err := doRun(m, binder); err != nil {
		log.Exit(err)
	}
}

func doRun(m *testing.M, binder Binder) (rerr error) {
	fv, err := flagParseFn()
	if err != nil {
		return err
	}
	b, err := binder()
	if err != nil {
		return fmt.Errorf("failed to create binding: %w", err)
	}
	setBindingFn(b)

	if fv.XMLPath != "" {
		if err := junitxml.StartConverting(fv.XMLPath); err != nil {
			return fmt.Errorf("error starting JUnit XML converter: %w", err)
		}
	}

	debugger.TestStarted(fv.Debug)
	if err := reserveFn(fv); err != nil {
		return err
	}
	go fnAfterSignal(releaseFn, unix.SIGINT, unix.SIGTERM)
	defer closer.Close(&rerr, func() error {
		debugger.TestCasesDone()
		return releaseFn()
	}, "error releasing testbed")
	debugger.ReservationDone()
	if fv.RunTime > 0 {
		go func() {
			time.Sleep(fv.RunTime)
			log.Exitf("Ondatra test timed out after %v", fv.RunTime)
		}()
	}

	runFn(m)

	if fv.XMLPath != "" {
		if err := junitxml.StopConverting(); err != nil {
			return fmt.Errorf("error stopping JUnit XML converter: %w", err)
		}
	}
	return nil
}

// fnAfterSignal waits for one of `signals` then calls `fn`.
func fnAfterSignal(fn func() error, signals ...os.Signal) {
	signal.Notify(sigc, signals...)
	s := <-sigc
	fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	log.Infof("caught signal %s, calling %s\n", s, fnName)
	if err := fn(); err != nil {
		log.Errorf("error calling %s: %v", fnName, err)
	}
}
