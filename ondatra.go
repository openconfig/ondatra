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
	"runtime/debug"
	"sync"
	"testing"
	"time"
	"unsafe"

	"flag"
	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/internal/closer"
	"golang.org/x/sys/unix"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/internal/reservemain"
)

var (
	sigc       = make(chan os.Signal, 1)
	reserveFn  = reserve
	releaseFn  = release
	runTestsFn = (*fixture).runTests
)

// Binder is the generator for providing binding to the test.
type Binder func() (binding.Binding, error)

// RunTests acquires the testbed of devices and runs the tests. Every device is
// initialized with a baseline configuration that allows it to be managed.
func RunTests(m *testing.M, binders ...Binder) {
	// Careful to only exit at the very end, because exiting skips all pending defers.
	if err := doRun(m, binders...); err != nil {
		log.Exit(err)
	}
}

func doRun(m *testing.M, binders ...Binder) (rerr error) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if len(binders) > 1 {
		return fmt.Errorf("found %d bindings, one or zero(deprecated) required", len(binders))
	}
	if len(binders) == 1 {
		b, err := binders[0]()
		if err != nil {
			return fmt.Errorf("failed to create binding: %w", err)
		}
		binding.Init(b)
	}
	if !binding.IsSet() {
		log.Warning("Binding is not set, this will likely cause a panic during test.")
	}
	fmt.Println(actionMsg("Reserving the testbed"))
	if err := reserveFn(*reservemain.TestbedPath, *reservemain.RunTime, *reservemain.WaitTime); err != nil {
		return err
	}
	go fnAfterSignal(releaseFn, unix.SIGINT, unix.SIGTERM)
	defer closer.Close(&rerr, func() error {
		fmt.Println(actionMsg("Releasing the testbed"))
		return releaseFn()
	}, "error releasing testbed")
	runTestsFn(new(fixture), m, *reservemain.RunTime)
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

type fixture struct {
	mu        sync.Mutex
	earlyFail bool
	fatalFn   func(args ...interface{})
}

func (f *fixture) runTests(m *testing.M, timeout time.Duration) {
	tests := reflect.ValueOf(m).Elem().FieldByName("tests")
	for i := 0; i < tests.Len(); i++ {
		fnVal := tests.Index(i).FieldByName("F")
		fnPtr := (*func(*testing.T))(unsafe.Pointer(fnVal.UnsafeAddr()))
		fn := *fnPtr
		*fnPtr = func(t *testing.T) {
			f.testStarted(t, timeout)
			binding.Get().SetTestMetadata(&binding.TestMetadata{TestName: t.Name()})
			defer func() {
				if r := recover(); r != nil {
					f.failEarly(fmt.Sprintf("Ondatra test panicked: %v, stack :%s", r, debug.Stack()))
				}
			}()
			fn(t)
		}
	}
	m.Run()
}

func (f *fixture) testStarted(t *testing.T, timeout time.Duration) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.earlyFail {
		t.SkipNow()
	}
	// Wait until the first testcase to start the timer, so fatalFn is definitely set.
	if timeout > 0 && f.fatalFn == nil {
		go func() {
			time.Sleep(timeout)
			f.failEarly(fmt.Sprintf("Ondatra test timed out after %v", timeout))
		}()
	}
	f.fatalFn = t.Fatal
}

func (f *fixture) failEarly(msg string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.earlyFail = true
	f.fatalFn(msg)
}

func logAction(t testing.TB, format string, dev reservation.Device) {
	t.Helper()
	t.Log(actionMsg(fmt.Sprintf(format, dev.Dimensions().Name)))
}

func actionMsg(msg string) string {
	return fmt.Sprintf("*** %s", msg)
}
