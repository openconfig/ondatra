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
	"flag"
)

var (
	// TestbedPath is a flag for the testbed path.
	TestbedPath = flag.String("testbed", "", "Path to the Ondatra testbed file")
	// RunTime is a flag for the run time duration of the reservation.
	RunTime = flag.Duration("run_time", 0, "Timeout of the test run, excluding the wait time for the testbed to be ready. "+
		"A zero value means here is no time limit. Must be a non-negative value.")
	// WaitTime is a flag for the wait time duration of the reservation.
	WaitTime = flag.Duration("wait_time", 0, "Maximum amount of time the test should wait until the testbed is ready. "+
		"A zero value lets the binding implementation choose an appropriate wait time. Must be a non-negative value.")
)
