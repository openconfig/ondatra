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

package jsonl

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	opb "github.com/openconfig/ondatra/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/testing/protocmp"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func resetPlanID() {
	AddProperty("", "test.plan_id", "")
}

func TestAddProperty(t *testing.T) {
	tests := []struct {
		name      string
		testName  string
		propName  string
		propValue string
		wantPlan  string
	}{
		{
			name:      "irrelevant_property_ignored",
			testName:  "",
			propName:  "dut.vendor",
			propValue: "ARISTA",
			wantPlan:  "",
		},
		{
			name:      "test_plan_id_property_sets_plan_id",
			testName:  "",
			propName:  "test.plan_id",
			propValue: "RT-1.25",
			wantPlan:  "RT-1.25",
		},
		{
			name:      "empty_property_value_sets_empty",
			testName:  "",
			propName:  "test.plan_id",
			propValue: "",
			wantPlan:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetPlanID()
			t.Cleanup(resetPlanID)

			AddProperty(tt.testName, tt.propName, tt.propValue)
			got := ResolvePlanID()
			if tt.wantPlan != "" && got != tt.wantPlan {
				t.Errorf("ResolvePlanID() after AddProperty(%q, %q, %q) = %q, want %q", tt.testName, tt.propName, tt.propValue, got, tt.wantPlan)
			}
		})
	}
}

func TestResolvePlanID(t *testing.T) {
	tests := []struct {
		name       string
		planIDProp string
		osArgs     []string
		wantPlanID string
	}{
		{
			name:       "explicit_plan_id_override",
			planIDProp: "EXPLICIT-1.0",
			wantPlanID: "EXPLICIT-1.0",
		},
		{
			name:       "fallback_to_executable_binary_name",
			osArgs:     []string{"/path/to/my_feature_test"},
			wantPlanID: "my_feature_test",
		},
		{
			name:       "fallback_trims_exe_suffix",
			osArgs:     []string{"/path/to/my_feature_test.exe"},
			wantPlanID: "my_feature_test",
		},
		{
			name:       "fallback_trims_test_suffix",
			osArgs:     []string{"/path/to/my_feature_test.test"},
			wantPlanID: "my_feature_test",
		},
		{
			name:       "fallback_empty_args_returns_unknown",
			osArgs:     []string{""},
			wantPlanID: "UNKNOWN_TEST_ID",
		},
		{
			name:       "fallback_dot_arg_returns_unknown",
			osArgs:     []string{"."},
			wantPlanID: "UNKNOWN_TEST_ID",
		},
		{
			name:       "fallback_nil_args_returns_unknown",
			osArgs:     nil,
			wantPlanID: "UNKNOWN_TEST_ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetPlanID()
			t.Cleanup(resetPlanID)

			if tt.planIDProp != "" {
				AddProperty("", "test.plan_id", tt.planIDProp)
			}

			if tt.osArgs != nil || tt.name == "fallback_nil_args_returns_unknown" {
				origArgs := os.Args
				os.Args = tt.osArgs
				t.Cleanup(func() { os.Args = origArgs })
			}

			got := ResolvePlanID()
			if got != tt.wantPlanID {
				t.Errorf("ResolvePlanID() = %q, want %q", got, tt.wantPlanID)
			}
		})
	}
}

func TestAppendRecord(t *testing.T) {
	ts1 := time.Date(2026, 8, 16, 15, 0, 0, 0, time.UTC)
	ts2 := time.Date(2026, 8, 16, 15, 5, 0, 0, time.UTC)

	tests := []struct {
		name        string
		records     []*opb.TestResult
		path        func(t *testing.T) string
		wantErr     bool
		wantRecords []*opb.TestResult
	}{
		{
			name: "empty_path_returns_error",
			records: []*opb.TestResult{
				{TestId: "RT-1.1", Status: opb.TestStatus_TEST_STATUS_PASS},
			},
			path:        func(t *testing.T) string { return "" },
			wantErr:     true,
			wantRecords: nil,
		},
		{
			name:        "nil_record_returns_error",
			records:     []*opb.TestResult{nil},
			path:        func(t *testing.T) string { return filepath.Join(t.TempDir(), "results.jsonl") },
			wantErr:     true,
			wantRecords: nil,
		},
		{
			name: "valid_single_record",
			records: []*opb.TestResult{
				{
					TestId:             "RT-1.1",
					Status:             opb.TestStatus_TEST_STATUS_PASS,
					ExecutionTimestamp: timestamppb.New(ts1),
				},
			},
			path:    func(t *testing.T) string { return filepath.Join(t.TempDir(), "single.jsonl") },
			wantErr: false,
			wantRecords: []*opb.TestResult{
				{
					TestId:             "RT-1.1",
					Status:             opb.TestStatus_TEST_STATUS_PASS,
					ExecutionTimestamp: timestamppb.New(ts1),
				},
			},
		},
		{
			name: "valid_multiple_records_with_details",
			records: []*opb.TestResult{
				{
					TestId:             "RT-1.1",
					Status:             opb.TestStatus_TEST_STATUS_PASS,
					ExecutionTimestamp: timestamppb.New(ts1),
				},
				{
					TestId:             "BGP-1.2",
					Status:             opb.TestStatus_TEST_STATUS_FAIL,
					StatusDetails:      "BGP session failed to reach ESTABLISHED",
					ExecutionTimestamp: timestamppb.New(ts2),
				},
			},
			path:    func(t *testing.T) string { return filepath.Join(t.TempDir(), "nested", "dir", "results.jsonl") },
			wantErr: false,
			wantRecords: []*opb.TestResult{
				{
					TestId:             "RT-1.1",
					Status:             opb.TestStatus_TEST_STATUS_PASS,
					ExecutionTimestamp: timestamppb.New(ts1),
				},
				{
					TestId:             "BGP-1.2",
					Status:             opb.TestStatus_TEST_STATUS_FAIL,
					StatusDetails:      "BGP session failed to reach ESTABLISHED",
					ExecutionTimestamp: timestamppb.New(ts2),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			targetPath := tt.path(t)

			for _, rec := range tt.records {
				err := appendRecord(targetPath, rec)
				if (err != nil) != tt.wantErr {
					t.Fatalf("appendRecord(%q, %+v) error = %v, wantErr %v", targetPath, rec, err, tt.wantErr)
				}
			}

			if targetPath == "" || len(tt.wantRecords) == 0 {
				return
			}

			f, err := os.Open(targetPath)
			if err != nil {
				t.Fatalf("Failed to open ledger file %q: %v", targetPath, err)
			}
			defer f.Close()

			var gotRecords []*opb.TestResult
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				res := &opb.TestResult{}
				if err := protojson.Unmarshal([]byte(line), res); err != nil {
					t.Fatalf("Failed to unmarshal JSON line %q: %v", line, err)
				}
				gotRecords = append(gotRecords, res)
			}
			if err := scanner.Err(); err != nil {
				t.Fatalf("Scanner error on %q: %v", targetPath, err)
			}

			if diff := cmp.Diff(tt.wantRecords, gotRecords, protocmp.Transform()); diff != "" {
				t.Errorf("Parsed records mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAppendResult(t *testing.T) {
	fixedTime := time.Date(2026, 8, 16, 12, 0, 0, 123456789, time.UTC)

	tests := []struct {
		name       string
		planID     string
		status     opb.TestStatus
		details    string
		path       func(t *testing.T) string
		wantErr    bool
		wantResult *opb.TestResult
	}{
		{
			name:       "empty_path_returns_error",
			planID:     "RT-1.1",
			status:     opb.TestStatus_TEST_STATUS_PASS,
			path:       func(t *testing.T) string { return "" },
			wantErr:    true,
			wantResult: nil,
		},
		{
			name:    "passing_result",
			planID:  "INTEG-1.0",
			status:  opb.TestStatus_TEST_STATUS_PASS,
			details: "",
			path:    func(t *testing.T) string { return filepath.Join(t.TempDir(), "pass.jsonl") },
			wantErr: false,
			wantResult: &opb.TestResult{
				TestId:             "INTEG-1.0",
				Status:             opb.TestStatus_TEST_STATUS_PASS,
				StatusDetails:      "",
				ExecutionTimestamp: timestamppb.New(fixedTime),
			},
		},
		{
			name:    "failing_result_with_details",
			planID:  "FAIL-2.0",
			status:  opb.TestStatus_TEST_STATUS_FAIL,
			details: "device unreachable: connection refused",
			path:    func(t *testing.T) string { return filepath.Join(t.TempDir(), "fail.jsonl") },
			wantErr: false,
			wantResult: &opb.TestResult{
				TestId:             "FAIL-2.0",
				Status:             opb.TestStatus_TEST_STATUS_FAIL,
				StatusDetails:      "device unreachable: connection refused",
				ExecutionTimestamp: timestamppb.New(fixedTime),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetPlanID()
			t.Cleanup(resetPlanID)

			if tt.planID != "" {
				AddProperty("", "test.plan_id", tt.planID)
			}

			origTimeNow := timeNowFn
			timeNowFn = func() time.Time { return fixedTime }
			t.Cleanup(func() { timeNowFn = origTimeNow })

			targetPath := tt.path(t)
			err := AppendResult(targetPath, tt.status, tt.details)
			if (err != nil) != tt.wantErr {
				t.Fatalf("AppendResult(%q, %v, %q) error = %v, wantErr %v", targetPath, tt.status, tt.details, err, tt.wantErr)
			}

			if targetPath == "" || tt.wantResult == nil {
				return
			}

			data, err := os.ReadFile(targetPath)
			if err != nil {
				t.Fatalf("Failed to read ledger file %q: %v", targetPath, err)
			}

			got := &opb.TestResult{}
			if err := protojson.Unmarshal(data, got); err != nil {
				t.Fatalf("Failed to unmarshal Result JSON: %v", err)
			}

			if diff := cmp.Diff(tt.wantResult, got, protocmp.Transform()); diff != "" {
				t.Errorf("AppendResult Result mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestConcurrentAppends(t *testing.T) {
	tempDir := t.TempDir()
	ledgerPath := filepath.Join(tempDir, "concurrent.jsonl")

	const numGoroutines = 30
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(idx int) {
			defer wg.Done()
			res := &opb.TestResult{
				TestId:             fmt.Sprintf("CONCURRENT-%02d", idx),
				Status:             opb.TestStatus_TEST_STATUS_PASS,
				ExecutionTimestamp: timestamppb.Now(),
			}
			if err := appendRecord(ledgerPath, res); err != nil {
				t.Errorf("Concurrent appendRecord(%d) failed: %v", idx, err)
			}
		}(i)
	}

	wg.Wait()

	f, err := os.Open(ledgerPath)
	if err != nil {
		t.Fatalf("Failed to open concurrent ledger: %v", err)
	}
	defer f.Close()

	seenIDs := make(map[string]bool)
	scanner := bufio.NewScanner(f)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		res := &opb.TestResult{}
		if err := protojson.Unmarshal([]byte(line), res); err != nil {
			t.Fatalf("Corrupted JSON on line %d (%q): %v", lineCount, line, err)
		}
		if res.GetTestId() == "" {
			t.Errorf("Empty TestID on line %d", lineCount)
		}
		if res.GetStatus() != opb.TestStatus_TEST_STATUS_PASS {
			t.Errorf("Line %d status = %v, want %v", lineCount, res.GetStatus(), opb.TestStatus_TEST_STATUS_PASS)
		}
		if seenIDs[res.GetTestId()] {
			t.Errorf("Duplicate record for TestID %q on line %d", res.GetTestId(), lineCount)
		}
		seenIDs[res.GetTestId()] = true
	}

	if err := scanner.Err(); err != nil {
		t.Fatalf("Scanner error on %q: %v", ledgerPath, err)
	}

	if lineCount != numGoroutines {
		t.Errorf("Total records written = %d, want %d", lineCount, numGoroutines)
	}
}

func TestEvalOutcome(t *testing.T) {
	intPtr := func(i int) *int { return &i }

	tests := []struct {
		name        string
		rerr        error
		exitCode    *int
		wantStatus  opb.TestStatus
		wantDetails string
	}{
		{
			name:        "rerr_error",
			rerr:        fmt.Errorf("reservation failed"),
			exitCode:    nil,
			wantStatus:  opb.TestStatus_TEST_STATUS_FAIL,
			wantDetails: "reservation failed",
		},
		{
			name:        "nil_exit_code_incomplete",
			rerr:        nil,
			exitCode:    nil,
			wantStatus:  opb.TestStatus_TEST_STATUS_FAIL,
			wantDetails: "test execution did not complete",
		},
		{
			name:        "non_zero_exit_code",
			rerr:        nil,
			exitCode:    intPtr(1),
			wantStatus:  opb.TestStatus_TEST_STATUS_FAIL,
			wantDetails: "test exited with code 1",
		},
		{
			name:        "successful_run",
			rerr:        nil,
			exitCode:    intPtr(0),
			wantStatus:  opb.TestStatus_TEST_STATUS_PASS,
			wantDetails: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, gotDetails := evalOutcome(tt.rerr, tt.exitCode)
			if gotStatus != tt.wantStatus {
				t.Errorf("evalOutcome() status = %v, want %v", gotStatus, tt.wantStatus)
			}
			if gotDetails != tt.wantDetails {
				t.Errorf("evalOutcome() details = %q, want %q", gotDetails, tt.wantDetails)
			}
		})
	}
}

func TestRecordResult(t *testing.T) {
	intPtr := func(i int) *int { return &i }

	tests := []struct {
		name        string
		path        func(t *testing.T) string
		rerr        error
		exitCode    *int
		wantStatus  opb.TestStatus
		wantDetails string
		wantWritten bool
		wantErr     bool
	}{
		{
			name:        "empty_path_is_noop",
			path:        func(t *testing.T) string { return "" },
			exitCode:    intPtr(0),
			wantWritten: false,
		},
		{
			name:        "success_records_pass",
			path:        func(t *testing.T) string { return filepath.Join(t.TempDir(), "pass.jsonl") },
			exitCode:    intPtr(0),
			wantStatus:  opb.TestStatus_TEST_STATUS_PASS,
			wantWritten: true,
		},
		{
			name:        "reservation_error_records_fail",
			path:        func(t *testing.T) string { return filepath.Join(t.TempDir(), "fail.jsonl") },
			rerr:        fmt.Errorf("reservation failed"),
			wantStatus:  opb.TestStatus_TEST_STATUS_FAIL,
			wantDetails: "reservation failed",
			wantWritten: true,
		},
		{
			name:        "non_zero_exit_code_records_fail",
			path:        func(t *testing.T) string { return filepath.Join(t.TempDir(), "exit_code.jsonl") },
			exitCode:    intPtr(2),
			wantStatus:  opb.TestStatus_TEST_STATUS_FAIL,
			wantDetails: "test exited with code 2",
			wantWritten: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			targetPath := tt.path(t)
			err := RecordResult(targetPath, tt.rerr, tt.exitCode)
			if (err != nil) != tt.wantErr {
				t.Fatalf("RecordResult(%q) error = %v, wantErr = %v", targetPath, err, tt.wantErr)
			}

			if !tt.wantWritten {
				return
			}

			data, err := os.ReadFile(targetPath)
			if err != nil {
				t.Fatalf("Failed to read ledger file %q: %v", targetPath, err)
			}

			got := &opb.TestResult{}
			if err := protojson.Unmarshal(data, got); err != nil {
				t.Fatalf("Failed to unmarshal Result JSON: %v", err)
			}
			if got.GetStatus() != tt.wantStatus {
				t.Errorf("Result.Status = %v, want %v", got.GetStatus(), tt.wantStatus)
			}
			if got.GetStatusDetails() != tt.wantDetails {
				t.Errorf("Result.StatusDetails = %q, want %q", got.GetStatusDetails(), tt.wantDetails)
			}
		})
	}
}
