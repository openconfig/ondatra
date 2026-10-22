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
	"strconv"
	"strings"
	"sync"
	"time"

	log "github.com/golang/glog"
	opb "github.com/openconfig/ondatra/proto"
	"golang.org/x/sys/unix"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// Suite property names recorded in the ledger.
const (
	planIDProperty    = "test.plan_id"
	gitCommitProperty = "git.commit"
	gitCleanProperty  = "git.clean"
)

var (
	propsMu   sync.RWMutex
	planID    string
	gitCommit string
	gitDirty  *bool // nil if git.clean was not reported or is not a bool.

	timeNowFn = time.Now
)

// AddProperty records the name-value property if it is one recorded in the
// ledger: the test plan ID, the git commit, or the git clean status.
// Note: test parameter is ignored as these properties are suite-scoped (last
// write wins).
func AddProperty(test, name, value string) {
	propsMu.Lock()
	defer propsMu.Unlock()
	switch name {
	case planIDProperty:
		planID = value
	case gitCommitProperty:
		gitCommit = value
	case gitCleanProperty:
		gitDirty = nil
		if clean, err := strconv.ParseBool(value); err == nil {
			gitDirty = proto.Bool(!clean)
		}
	}
}

// ResolvePlanID resolves the test plan ID from explicitly configured values
// (such as ondatra.Report().AddSuiteProperty("test.plan_id", ...)), or falls back
// to the executable base name.
func ResolvePlanID() string {
	propsMu.RLock()
	id := planID
	propsMu.RUnlock()

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

// AppendResult creates a TestResult with the resolved plan ID, git commit and
// dirty state of the test code, current timestamp, status, and details, and
// atomically appends it to the specified JSON-Lines file.
func AppendResult(path string, status opb.TestStatus, details string) error {
	propsMu.RLock()
	commit, dirty := gitCommit, gitDirty
	propsMu.RUnlock()

	res := &opb.TestResult{
		TestId:             ResolvePlanID(),
		Status:             status,
		StatusDetails:      details,
		ExecutionTimestamp: timestamppb.New(timeNowFn()),
		GitCommit:          commit,
		GitDirty:           dirty,
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
