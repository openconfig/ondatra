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
	"github.com/jstemmer/go-junit-report/v2/junit"
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
=== RUN   TestFail
    fail_test.go:6: Error message
    fail_test.go:7: Longer
        error
        message.
--- FAIL: TestFail (0.15s)
=== RUN   TestSkip
    skip_test.go:6: skip message
--- SKIP: TestSkip (0.02s)
=== RUN   TestSkipNow
    skip_test.go:10: log message
--- SKIP: TestSkipNow (0.13s)
FAIL
FAIL  	example_test	0.46s`
	testXML = `<?xml version="1.0" encoding="UTF-8"?>
<testsuites tests="5" failures="1" skipped="2">
	<testsuite name="example_test" tests="5" failures="1" errors="0" id="0" skipped="2" time="0.460" timestamp="2022-01-01T00:00:00Z">
		<properties>
			<property name="suiteKey" value="suiteVal"></property>
			<property name="TestPass/passKey" value="passVal"></property>
			<property name="TestFail/failKey" value="failVal"></property>
		</properties>
		<testcase name="TestPass" classname="example_test" time="0.060"></testcase>
		<testcase name="TestPassLog" classname="example_test" time="0.100">
			<system-out><![CDATA[    pass_test.go:9: log line
    pass_test.go:10: log
        multi
        line]]></system-out>
		</testcase>
		<testcase name="TestFail" classname="example_test" time="0.150">
			<failure message="Failed"><![CDATA[    fail_test.go:6: Error message
    fail_test.go:7: Longer
        error
        message.]]></failure>
		</testcase>
		<testcase name="TestSkip" classname="example_test" time="0.020">
			<skipped message="Skipped"><![CDATA[    skip_test.go:6: skip message]]></skipped>
		</testcase>
		<testcase name="TestSkipNow" classname="example_test" time="0.130">
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
	srcBuf := new(bytes.Buffer)
	dstBuf := new(bytes.Buffer)
	conv := mustStart(t, srcBuf, dstBuf)
	if err := conv.Stop(); err != nil {
		t.Errorf("Stop failed: %v", err)
	}
	if diff := cmp.Diff(testLog, srcBuf.String()); diff != "" {
		t.Errorf("Converter wrote unexpected tee output (-want, +got): %s", diff)
	}
	if diff := cmp.Diff(testXML, dstBuf.String()); diff != "" {
		t.Errorf("Converter wrote unexpected dest output (-want, +got): %s", diff)
	}
}

func TestSrcError(t *testing.T) {
	wantErr := "src error"
	errSrc := &errorWriteCloser{wantErr}
	err := mustStart(t, errSrc, new(bytes.Buffer)).Stop()
	if s := errdiff.Substring(err, wantErr); s != "" {
		t.Errorf("Stop got unexpected error: %s", s)
	}
}

func TestDstError(t *testing.T) {
	wantErr := "dst error"
	errDst := &errorWriteCloser{wantErr}
	err := mustStart(t, new(bytes.Buffer), errDst).Stop()
	if s := errdiff.Substring(err, wantErr); s != "" {
		t.Errorf("Stop got unexpected error: %s", s)
	}
}

func mustStart(t *testing.T, src, dst io.Writer) *converter {
	conv, err := startConverter(src, dst)
	if err != nil {
		t.Fatalf("startConverting failed: %v", err)
	}
	conv.addProperty("", "suiteKey", "suiteVal")
	conv.addProperty("TestPass", "passKey", "passVal")
	conv.addProperty("TestFail", "failKey", "failVal")
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

func TestEncodeProperty(t *testing.T) {
	const wantVal = "propValue"

	tests := []struct {
		desc, test, name, wantName string
	}{{
		desc:     "suite property",
		name:     "propName",
		wantName: "propName",
	}, {
		desc:     "test property",
		test:     "TestExample",
		name:     "propName",
		wantName: "TestExample/propName",
	}, {
		desc:     "subtest property",
		test:     "TestExample/subtest",
		name:     "propName",
		wantName: "TestExample/subtest/propName",
	}, {
		desc: "suite property name empty",
	}, {
		desc:     "test property name empty",
		test:     "TestExample",
		wantName: "TestExample/",
	}, {
		desc:     "suite property name has slashes",
		name:     "foo/bar",
		wantName: "foo//bar",
	}, {
		desc:     "test property name has slashes",
		test:     "TestExample",
		name:     "foo/bar",
		wantName: "TestExample/foo//bar",
	}, {
		desc:     "property name has multiple slashes",
		test:     "TestExample",
		name:     "foo/bar/baz",
		wantName: "TestExample/foo//bar//baz",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := encodeProperty(test.test, test.name, wantVal)
			if got.Name != test.wantName {
				t.Errorf("EncodeProperty got name %q, want %q", got.Name, test.wantName)
			}
			if got.Value != wantVal {
				t.Errorf("EncodeProperty got value %q, want %q", got.Value, wantVal)
			}
		})
	}
}

func TestDecodeProperty(t *testing.T) {
	const wantVal = "propValue"

	tests := []struct {
		desc, name, wantTest, wantName string
	}{{
		desc:     "suite property",
		name:     "propName",
		wantName: "propName",
	}, {
		desc:     "test property",
		name:     "TestExample/propName",
		wantTest: "TestExample",
		wantName: "propName",
	}, {
		desc:     "subtest property",
		name:     "TestExample/subtest/propName",
		wantTest: "TestExample/subtest",
		wantName: "propName",
	}, {
		desc: "suite property name empty",
	}, {
		desc:     "test property name empty",
		name:     "TestExample/",
		wantTest: "TestExample",
	}, {
		desc:     "suite property name has slashes",
		name:     "foo//bar",
		wantName: "foo/bar",
	}, {
		desc:     "test property name has slashes",
		name:     "TestExample/foo//bar",
		wantTest: "TestExample",
		wantName: "foo/bar",
	}, {
		desc:     "property name has multiple slashes",
		name:     "TestExample/foo//bar//baz",
		wantTest: "TestExample",
		wantName: "foo/bar/baz",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			prop := junit.Property{Name: test.name, Value: wantVal}
			gotTest, gotName, gotVal := DecodeProperty(prop)
			if gotTest != test.wantTest {
				t.Errorf("DecodeProperty got test name %q, want %q", gotTest, test.wantTest)
			}
			if gotName != test.wantName {
				t.Errorf("DecodeProperty got property name %q, want %q", gotName, test.wantName)
			}
			if gotVal != wantVal {
				t.Errorf("DecodeProperty got property value %q, want %q", gotVal, wantVal)
			}
		})
	}
}
