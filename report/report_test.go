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
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jstemmer/go-junit-report/v2/junit"
)

func TestExtractProperties(t *testing.T) {
	props := []junit.Property{
		{Name: "foo", Value: "bar"},
		{Name: "TestExample.foo", Value: "bar"},
		{Name: "TestExample/subtest.foo", Value: "bar"},
		{Name: "foo", Value: "baz"},
		{Name: "TestExample.bar", Value: "baz"},
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
