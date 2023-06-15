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
	"testing"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	closer "github.com/openconfig/gocloser"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/debug"
	"github.com/openconfig/ondatra/eventlis"
	"github.com/openconfig/ondatra/internal/ate"
	"github.com/openconfig/ondatra/internal/events"
	"github.com/openconfig/ondatra/internal/flags"
	"github.com/openconfig/ondatra/internal/junitxml"
	"github.com/openconfig/ondatra/internal/rawapis"
	"github.com/openconfig/ondatra/internal/testbed"
	"github.com/openconfig/ondatra/report"
	"golang.org/x/sys/unix"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

var sigc = make(chan os.Signal, 1)

// RunTests acquires the testbed of devices and runs the tests. Every device is
// initialized with a baseline configuration that allows it to be managed.
func RunTests(m *testing.M, newBindFn func() (binding.Binding, error)) {
	// Careful to only exit at the very end, because exiting skips all pending defers.
	if err := runTests(m.Run, newBindFn); err != nil {
		// If runTests returns an error, no test cases will be executed and the XML
		// result file will be empty. To avoid user confusion over empty XML
		// results, print the output of a fake TestMain test case, so that the
		// failure will be elevated to the XML, not merely hidden in the test log.
		// TODO(team): Consider unconditionally printing a TestMain test case
		// with the time until reservation to enable the Ondatra test timing report.
		fmt.Println("=== RUN   TestMain")
		fmt.Printf("    %v\n", err)
		fmt.Println("--- FAIL: TestMain (0.00s)")
		log.Exit()
	}
}

func runTests(runFn func() int, newBindFn func() (binding.Binding, error)) (rerr error) {
	flagVals, err := flags.Parse()
	if err != nil {
		return err
	}
	bind, err := newBindFn()
	if err != nil {
		return fmt.Errorf("failed to create binding: %w", err)
	}
	testbed.SetBinding(bind)

	if flagVals.XMLPath != "" {
		if err := junitxml.StartConverting(flagVals.XMLPath); err != nil {
			return fmt.Errorf("error starting JUnit XML converter: %w", err)
		}
	}

	events.TestStarted(flagVals.Debug)
	ctx := context.Background()
	if err := testbed.Reserve(ctx, flagVals); err != nil {
		return err
	}
	go releaseOnSignal(ctx)
	defer closer.Close(&rerr, func() error {
		return testbed.Release(ctx)
	}, "error releasing testbed")

	var exitCode *int
	defer closer.Close(&rerr, func() error {
		return events.TestsDone(exitCode)
	}, "error notifying tests are done")

	if err := events.ReservationDone(); err != nil {
		return err
	}
	if flagVals.RunTime > 0 {
		go func() {
			time.Sleep(flagVals.RunTime)
			log.Exitf("Ondatra test timed out after %v", flagVals.RunTime)
		}()
	}
	code := runFn()
	exitCode = &code

	if flagVals.XMLPath != "" {
		if err := junitxml.StopConverting(); err != nil {
			return fmt.Errorf("error stopping JUnit XML converter: %w", err)
		}
	}
	return nil
}

func releaseOnSignal(ctx context.Context) {
	signal.Notify(sigc, unix.SIGINT, unix.SIGTERM)
	s := <-sigc
	log.Infof("Caught signal %s, releasing testbed", s)
	if err := testbed.Release(ctx); err != nil {
		log.Errorf("Error releasing testbed: %v", err)
	}
}

// EventListener returns the Ondatra Event Listener API.
func EventListener() *eventlis.EventListener {
	return new(eventlis.EventListener)
}

// Report returns the Ondatra Report API.
func Report() *report.Report {
	return new(report.Report)
}

// Debug returns the Ondatra Debug API.
func Debug() *debug.Debug {
	return new(debug.Debug)
}

func checkRes(t testing.TB) *binding.Reservation {
	t.Helper()
	res, err := testbed.Reservation()
	if err != nil {
		t.Fatal(err)
	}
	return res
}

// DUT returns the DUT in the testbed with a given id.
func DUT(t testing.TB, id string) *DUTDevice {
	t.Helper()
	rd, err := testbed.DUT(checkRes(t), id)
	if err != nil {
		t.Fatalf("DUT(t, %s): %v", id, err)
	}
	return newDUT(id, rd)
}

// DUTs returns a map of DUT id to DUT in the testbed.
func DUTs(t testing.TB) map[string]*DUTDevice {
	t.Helper()
	rm := checkRes(t).DUTs
	m := make(map[string]*DUTDevice)
	for id, rd := range rm {
		m[id] = newDUT(id, rd)
	}
	return m
}

func newDUT(id string, res binding.DUT) *DUTDevice {
	return &DUTDevice{&Device{
		id:  id,
		res: res,
		clientFn: func(ctx context.Context) (gpb.GNMIClient, error) {
			return rawapis.FetchGNMI(ctx, res)
		},
	}}
}

// ATE returns the ATE in the testbed with a given id.
func ATE(t testing.TB, id string) *ATEDevice {
	t.Helper()
	ra, err := testbed.ATE(checkRes(t), id)
	if err != nil {
		t.Fatalf("ATE(t, %s): %v", id, err)
	}
	return newATE(id, ra)
}

// ATEs returns a map of ATE id to ATE in the testbed.
func ATEs(t testing.TB) map[string]*ATEDevice {
	t.Helper()
	rm := checkRes(t).ATEs
	m := make(map[string]*ATEDevice)
	for id, ra := range rm {
		m[id] = newATE(id, ra)
	}
	return m
}

func newATE(id string, res binding.ATE) *ATEDevice {
	return &ATEDevice{Device: &Device{
		id:  id,
		res: res,
		clientFn: func(ctx context.Context) (gpb.GNMIClient, error) {
			return ate.FetchGNMI(ctx, res)
		},
	}}
}

// ReservationID returns the reservation ID for the testbed.
func ReservationID(t testing.TB) string {
	t.Helper()
	return checkRes(t).ID
}
