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

package report

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jstemmer/go-junit-report/v2/junit"
)

func TestReadXML(t *testing.T) {
	const text = `<?xml version="1.0" encoding="UTF-8"?>
<testsuites tests="1" failures="0" skipped="0">
	<testsuite name="example_test" tests="1" failures="0" errors="0" id="0" skipped="0" time="1.000" timestamp="2022-01-01T00:00:00Z">
		<properties>
			<property name="propName" value="propVal"></property>
		</properties>
		<testcase name="TestCase" classname="example_test" time="1.000"></testcase>
	</testsuite>
</testsuites>
`
	r := strings.NewReader(text)
	got, err := ReadXML(r)
	if err != nil {
		t.Errorf("ReadXML got error: %v", err)
	}
	wantProps := []junit.Property{{Name: "propName", Value: "propVal"}}
	want := junit.Testsuites{
		XMLName: xml.Name{Local: "testsuites"},
		Tests:   1,
		Suites: []junit.Testsuite{{
			Name:       "example_test",
			Tests:      1,
			Time:       "1.000",
			Timestamp:  "2022-01-01T00:00:00Z",
			Properties: &wantProps,
			Testcases: []junit.Testcase{{
				Name:      "TestCase",
				Classname: "example_test",
				Time:      "1.000",
			}},
		}},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("ReadXML got unexpected diff (-want,+got): %s", diff)
	}
}

func TestReadXMLError(t *testing.T) {
	const text = "<>"
	r := strings.NewReader(text)
	suites, err := ReadXML(r)
	if err == nil {
		t.Fatalf("ReadXML unexpectedly succeeded: %v", suites)
	}
	if wantErr := "error reading XML"; !strings.Contains(err.Error(), wantErr) {
		t.Errorf("ReadXML got unexpected error got %q, want %q", err.Error(), wantErr)
	}
}

func TestExtractProperties(t *testing.T) {
	props := []junit.Property{
		{Name: "foo", Value: "bar"},
		{Name: "TestExample/foo", Value: "bar"},
		{Name: "TestExample/subtest/foo", Value: "bar"},
		{Name: "foo", Value: "baz"},
		{Name: "TestExample/bar", Value: "baz"},
	}
	want := map[string]map[string]string{
		"":                    {"foo": "baz"},
		"TestExample":         {"foo": "bar", "bar": "baz"},
		"TestExample/subtest": {"foo": "bar"},
	}
	suites := junit.Testsuites{Suites: []junit.Testsuite{{Properties: &props}}}
	got := ExtractProperties(suites)
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("ExtractProperties got unexpected diff: %s", diff)
	}
}

func TestExtractPropertiesNil(t *testing.T) {
	suites := junit.Testsuites{Suites: []junit.Testsuite{{}}}
	if got := ExtractProperties(suites); len(got) > 0 {
		t.Errorf("ExtractProperties got %v, want none", got)
	}
}
