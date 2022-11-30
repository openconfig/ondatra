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

// Package flags defines the flags accepted by Ondatra tests.
package flags

import (
	"fmt"
	"strings"
	"time"

	"flag"
)

var (
	testbed = flag.String("testbed", "", "Path to the Ondatra testbed file")
	runTime = flag.Duration("run_time", 0, "Timeout of the test run, excluding the wait time for the testbed to be ready. "+
		"A zero value means here is no time limit. Must be a non-negative value.")
	waitTime = flag.Duration("wait_time", 0, "Maximum amount of time the test should wait until the testbed is ready. "+
		"A zero value lets the binding implementation choose an appropriate wait time. Must be a non-negative value.")
	reserve = flag.String("reserve", "", "Reservation id or a mapping of device and port IDs to names of the form "+
		"'dut=mydevice,dut:port1=Ethernet1/1,ate=myixia,ate:port2=2/3'")
	xml   = flag.String("xml", "", "File path to write JUnit XML test results; disables normal Go test logging.")
	debug = flag.Bool("debug", false, "Whether the test is run in debug mode")
)

// Values is the set of parsed and validated flag values.
type Values struct {
	TestbedPath string
	RunTime     time.Duration
	WaitTime    time.Duration
	ResvID      string
	ResvPartial map[string]string
	XMLPath     string
	Debug       bool
}

// Parse parse and validates the flag values.
func Parse() (*Values, error) {
	if !flag.Parsed() {
		flag.Parse()
	}

	// Always stream verbose test output. This is necessary for XML support,
	// for notifying users when breakpoints are reached, and is generally
	// a better default user experience for Ondatra tests.
	flag.Set("test.v", "true")

	if *testbed == "" {
		return nil, fmt.Errorf("testbed path not specified")
	}
	if *runTime < 0 {
		return nil, fmt.Errorf("run timeout is negative: %d", *runTime)
	}
	if *waitTime < 0 {
		return nil, fmt.Errorf("wait timeout is negative: %d", *waitTime)
	}
	if *reserve != "" && !*debug {
		return nil, fmt.Errorf("reserve flag is only allowed in debug mode")
	}
	resvID, resvPartial, err := parseReserve(*reserve)
	if err != nil {
		return nil, err
	}
	return &Values{
		TestbedPath: *testbed,
		RunTime:     *runTime,
		WaitTime:    *waitTime,
		ResvID:      resvID,
		ResvPartial: resvPartial,
		XMLPath:     *xml,
		Debug:       *debug,
	}, nil
}

func parseReserve(res string) (string, map[string]string, error) {
	if res == "" {
		return "", nil, nil
	}

	// If the string does not contain an equals sign, it is a unique ID.
	if !strings.Contains(res, "=") {
		return res, nil, nil
	}

	// Otherwise, it is a partial device and port mapping.
	partial := make(map[string]string)
	for _, mapStr := range strings.Split(res, ",") {
		if mapStr == "" {
			continue
		}
		mapPair := strings.Split(mapStr, "=")
		if len(mapPair) != 2 {
			return "", nil, fmt.Errorf("cannot parse assignment %q from --reserve %s", mapStr, res)
		}
		// It may not be a link, but tolerate the syntax of a link here.
		rID := strings.TrimSpace(mapPair[0])
		if _, ok := partial[rID]; ok {
			return "", nil, fmt.Errorf("duplicate ID %q in --reserve %s", strings.TrimSpace(mapPair[0]), res)
		}
		partial[rID] = strings.TrimSpace(mapPair[1])
	}
	return "", partial, nil
}
