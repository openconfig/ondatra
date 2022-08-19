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

package junitxml

import (
	"bytes"
	"errors"
	"io"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/gnmi/errdiff"
)

const (
	testLog = `=== RUN   TestPass
--- PASS: TestPass (0.06s)
=== RUN   TestPassLog
    pass_test.go:9: log line
    pass_test.go:10: log
        multi
        line
--- PASS: TestPassLog (0.10s)
PASS
ok  	package/pass	0.160s
=== RUN   TestOne
    fail_test.go:6: Error message
    fail_test.go:7: Longer
        error
        message.
--- FAIL: TestOne (0.151s)
FAIL
FAIL	package/fail	0.151s
=== RUN   TestSkip
    skip_test.go:6: skip message
--- SKIP: TestSkip (0.02s)
=== RUN   TestSkipNow
    skip_test.go:10: log message
--- SKIP: TestSkipNow (0.13s)
PASS
ok  	package/skip	0.150s
FAIL`
	testXML = `<?xml version="1.0" encoding="UTF-8"?>
<testsuites tests="5" failures="1" skipped="2">
	<testsuite name="package/pass" tests="2" failures="0" errors="0" id="0" time="0.160" timestamp="2022-01-01T00:00:00Z">
		<testcase name="TestPass" classname="package/pass" time="0.060"></testcase>
		<testcase name="TestPassLog" classname="package/pass" time="0.100">
			<system-out><![CDATA[    pass_test.go:9: log line
    pass_test.go:10: log
        multi
        line]]></system-out>
		</testcase>
	</testsuite>
	<testsuite name="package/fail" tests="1" failures="1" errors="0" id="1" time="0.151" timestamp="2022-01-01T00:00:00Z">
		<testcase name="TestOne" classname="package/fail" time="0.151">
			<failure message="Failed"><![CDATA[    fail_test.go:6: Error message
    fail_test.go:7: Longer
        error
        message.]]></failure>
		</testcase>
	</testsuite>
	<testsuite name="package/skip" tests="2" failures="0" errors="0" id="2" skipped="2" time="0.150" timestamp="2022-01-01T00:00:00Z">
		<testcase name="TestSkip" classname="package/skip" time="0.020">
			<skipped message="Skipped"><![CDATA[    skip_test.go:6: skip message]]></skipped>
		</testcase>
		<testcase name="TestSkipNow" classname="package/skip" time="0.130">
			<skipped message="Skipped"><![CDATA[    skip_test.go:10: log message]]></skipped>
		</testcase>
	</testsuite>
</testsuites>
`
)

func TestConverter(t *testing.T) {
	timeNowFn = func() time.Time {
		return time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	}
	buff := new(bytes.Buffer)
	if err := mustStart(t, buff).Stop(); err != nil {
		t.Errorf("Stop failed: %v", err)
	}
	if diff := cmp.Diff(testXML, buff.String()); diff != "" {
		t.Errorf("Converter wrote unexpected output (-want, +got): %s", diff)
	}
}

func TestConverterError(t *testing.T) {
	wantErr := "write error"
	errDest := &errorWriteCloser{wantErr}
	err := mustStart(t, errDest).Stop()
	if s := errdiff.Substring(err, wantErr); s != "" {
		t.Errorf("Stop got unexpected error: %s", s)
	}
}

func mustStart(t *testing.T, dest io.Writer) *Converter {
	conv, err := startConverter(dest)
	if err != nil {
		t.Fatalf("startConverting failed: %v", err)
	}
	if _, err := conv.file.Write([]byte(testLog)); err != nil {
		t.Fatalf("Write failed: %v", err)
	}
	return conv
}

// errorWriteCloser is a WriteCloser whose Write call always returns an error.
type errorWriteCloser struct {
	msg string
}

func (w *errorWriteCloser) Write(p []byte) (int, error) {
	return 0, errors.New(w.msg)
}

func (*errorWriteCloser) Close() error {
	return nil
}
