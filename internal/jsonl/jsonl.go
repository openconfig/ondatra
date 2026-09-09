// Copyright 2026 Google LLC
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

// Package jsonl provides functionality for atomically appending test execution
// results to a JSON-Lines ledger file.
package jsonl

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	log "github.com/golang/glog"
	opb "github.com/openconfig/ondatra/proto"
	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/encoding/protojson"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var (
	planIDMu sync.RWMutex
	planID   string

	timeNowFn = time.Now
)

// AddProperty processes a name-value property and sets the plan ID if the key matches.
// Note: test parameter is ignored as test.plan_id is suite-scoped (last write wins).
func AddProperty(test, name, value string) {
	if name == "test.plan_id" {
		planIDMu.Lock()
		defer planIDMu.Unlock()
		planID = value
	}
}

// ResolvePlanID resolves the test plan ID from explicitly configured values
// (such as ondatra.Report().AddSuiteProperty("test.plan_id", ...)), or falls back
// to the executable base name.
func ResolvePlanID() string {
	planIDMu.RLock()
	id := planID
	planIDMu.RUnlock()

	if id != "" {
		return id
	}

	if len(os.Args) > 0 && os.Args[0] != "" {
		base := filepath.Base(os.Args[0])
		base = strings.TrimSuffix(base, ".exe")
		base = strings.TrimSuffix(base, ".test")
		if base != "" && base != "." {
			return base
		}
	}

	return "UNKNOWN_TEST_ID"
}

// AppendResult creates a TestResult with the resolved plan ID, current timestamp, status,
// and details, and atomically appends it to the specified JSON-Lines file.
func AppendResult(path string, status opb.TestStatus, details string) error {
	res := &opb.TestResult{
		TestId:             ResolvePlanID(),
		Status:             status,
		StatusDetails:      details,
		ExecutionTimestamp: timestamppb.New(timeNowFn()),
	}
	return appendRecord(path, res)
}

// appendRecord atomically appends a TestResult object to the JSON-Lines ledger at path
// using an exclusive flock to prevent interleaving from concurrent test runs.
func appendRecord(path string, res *opb.TestResult) (rerr error) {
	if path == "" {
		return errors.New("jsonl: path cannot be empty")
	}
	if res == nil {
		return errors.New("jsonl: result record cannot be nil")
	}

	data, err := protojson.Marshal(res)
	if err != nil {
		return fmt.Errorf("failed to marshal jsonl record: %w", err)
	}

	if dir := filepath.Dir(path); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory for jsonl file %q: %w", path, err)
		}
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open jsonl file %q: %w", path, err)
	}
	defer f.Close()

	if err := unix.Flock(int(f.Fd()), unix.LOCK_EX); err != nil {
		return fmt.Errorf("failed to acquire flock on %q: %w", path, err)
	}
	defer func() {
		if unlockErr := unix.Flock(int(f.Fd()), unix.LOCK_UN); unlockErr != nil && rerr == nil {
			rerr = fmt.Errorf("failed to release flock on %q: %w", path, unlockErr)
		}
	}()

	if _, err := f.Write(append(data, '\n')); err != nil {
		return fmt.Errorf("failed to write record to jsonl file %q: %w", path, err)
	}
	return nil
}

// RecordResult evaluates the raw test execution outcome and writes the record to disk.
func RecordResult(path string, rerr error, exitCode *int) error {
	if path == "" {
		return nil
	}
	status, details := evalOutcome(rerr, exitCode)
	if err := AppendResult(path, status, details); err != nil {
		log.Errorf("Error writing JSONL test result: %v", err)
		return err
	}
	return nil
}

func evalOutcome(rerr error, exitCode *int) (opb.TestStatus, string) {
	switch {
	case rerr != nil:
		return opb.TestStatus_TEST_STATUS_FAIL, rerr.Error()
	case exitCode == nil:
		return opb.TestStatus_TEST_STATUS_FAIL, "test execution did not complete"
	case *exitCode != 0:
		return opb.TestStatus_TEST_STATUS_FAIL, fmt.Sprintf("test exited with code %d", *exitCode)
	default:
		return opb.TestStatus_TEST_STATUS_PASS, ""
	}
}
