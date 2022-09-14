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

// Package report provides ways to customize the JUnit XML report.
package report

import (
	"encoding/xml"
	"fmt"
	"io"
	"testing"

	log "github.com/golang/glog"
	"github.com/jstemmer/go-junit-report/v2/junit"
	"github.com/openconfig/ondatra/internal/junitxml"
)

// Report is the API to customizing a JUnit XML report.
type Report struct{}

// AddSuiteProperty attaches a name-value property to the test suite in the
// generated XML report. This is preferred, but equivalent, to calling
// AddRawProperty("", name, value).
func (r *Report) AddSuiteProperty(name, value string) {
	r.AddRawProperty("", name, value)
}

// AddTestProperty attaches a name-value property to the current test in the
// generated XML report. This is preferred, but equivalent, to calling
// AddRawProperty(t.Name(), name, value).
func (r *Report) AddTestProperty(t *testing.T, name, value string) {
	r.AddRawProperty(t.Name(), name, value)
}

// AddRawProperty attaches a name-value property to the specified test name in
// the generated the XML report. If the test name is the empty string, the
// property is attached to the test suite. For improved test readability,
// callers should prefer using AddSuiteProperty or AddTestProperty.
func (r *Report) AddRawProperty(test, name, value string) {
	fmt.Printf("*** PROPERTY: %s -> %s\n", name, value)
	junitxml.AddProperty(test, name, value)
}

// ReadXML decodes XML bytes into a JUnit Testsuites element.
func ReadXML(r io.Reader) (junit.Testsuites, error) {
	suites := junit.Testsuites{}
	err := xml.NewDecoder(r).Decode(&suites)
	if err != nil {
		return suites, fmt.Errorf("error reading XML: %w", err)
	}
	return suites, nil
}

// ExtractProperties extracts a testName->propertyName->propertyValue map from
// the specified JUnit Testsuites. It assumes the specified Testsuites contains
// exactly one Testsuite, since the XML report always contains exactly one, and
// if the Testsuite contains multiple properties with the same name, later
// property values overwrite earlier property values.
func ExtractProperties(suites junit.Testsuites) map[string]map[string]string {
	if len(suites.Suites) != 1 {
		log.Fatalf("Testsuites must contain exactly one Testsuite")
	}
	jprops := suites.Suites[0].Properties
	if jprops == nil {
		return nil
	}
	props := make(map[string]map[string]string)
	for _, p := range *jprops {
		test, name, value := junitxml.DecodeProperty(p)
		testProps, ok := props[test]
		if !ok {
			testProps = make(map[string]string)
			props[test] = testProps
		}
		testProps[name] = value
	}
	return props
}
