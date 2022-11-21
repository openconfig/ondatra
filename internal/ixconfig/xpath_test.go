// Copyright 2021 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ixconfig

import (
	"encoding/json"
	"strconv"
	"testing"

	// TODO(team): If open-sourced, replace with the equivalent "github.com/google/go-cmp/cmp" package.
	"github.com/google/go-cmp/cmp"
)

// TestMarshalUnmarshalXPath tests correctness of both unmarshaling an XPath as a JSON string, and then
// marshaling the XPath object back to a string (if the initial unmarshal should have succeeded.)
func TestMarshalUnmarshalXPath(t *testing.T) {
	tests := []struct {
		desc      string
		jsonXPath string
		wantXPath XPath
		wantErr   bool
	}{{
		desc:      "Root XPath",
		jsonXPath: "/",
		wantXPath: XPath{
			parentXPath: "/",
		},
	}, {
		desc:      "Single subconfig node.",
		jsonXPath: "/availableHardware",
		wantXPath: XPath{
			parentXPath: "/",
			objectName:  "availableHardware",
		},
	}, {
		desc:      "Multivalue path.",
		jsonXPath: "/multivalue[@source = '/globals/topology/ancp/portDownRate enabled']",
		wantXPath: XPath{
			parentXPath:  "/globals/topology/ancp/portDownRate",
			objectName:   "enabled",
			isMultivalue: true,
		},
	}, {
		desc:      "Array path.",
		jsonXPath: "/traffic/trafficItem[2]",
		wantXPath: XPath{
			parentXPath: "/traffic",
			objectName:  "trafficItem",
			index:       2,
		},
	}, {
		desc:      "Alias path.",
		jsonXPath: "/traffic/trafficItem[1]/configElement[1]/stack[@alias = 'ethernet-1']",
		wantXPath: XPath{
			parentXPath: "/traffic/trafficItem[1]/configElement[1]",
			objectName:  "stack",
			alias:       "ethernet-1",
		},
	}, {
		desc:      "Alias with spaces",
		jsonXPath: "/traffic/trafficItem[5]/configElement[1]/stack[@alias = 'HTTP_200_OK-4']/field[@alias = 'HTTP_200_OK.header.fieldHolder0.REQUEST VERSION-1']",
		wantXPath: XPath{
			parentXPath: "/traffic/trafficItem[5]/configElement[1]/stack[@alias = 'HTTP_200_OK-4']",
			objectName:  "field",
			alias:       "HTTP_200_OK.header.fieldHolder0.REQUEST VERSION-1",
		},
	}, {
		desc:      "Complex parent path with spaces.",
		jsonXPath: "/availableHardware/chassis[@alias = '108.177.95.165']/card[1]",
		wantXPath: XPath{
			parentXPath: "/availableHardware/chassis[@alias = '108.177.95.165']",
			objectName:  "card",
			index:       1,
		},
	}, {
		desc:      "Invalid path (has an invalid array index.).",
		jsonXPath: "/traffic/trafficItem[0]",
		wantErr:   true,
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			xPath := XPath{}
			quotedJSONXPath := strconv.Quote(tc.jsonXPath)

			err := json.Unmarshal([]byte(quotedJSONXPath), &xPath)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Unexpected error value from XPath Unmarshal, got error: %v, wanted error? %t", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}

			if diff := cmp.Diff(tc.wantXPath, xPath, cmp.AllowUnexported(XPath{})); diff != "" {
				t.Fatalf("Incorrect XPath unmarshaled from %q, diff was: %s", tc.jsonXPath, diff)
			}
			xPathStr, err := json.Marshal(&xPath)
			if err != nil {
				t.Fatalf("Unexpected error re-marshaling XPath %q: %v", tc.jsonXPath, err)
			}
			if string(xPathStr) != quotedJSONXPath {
				t.Fatalf("Unexpected marshaling result: got %v, want %v", string(xPathStr), quotedJSONXPath)
			}
		})
	}
}

// TestMarshalXPathSchemaIssues tests the marshaling of some XPaths that require special handling
// because their behavior is inconsistent with other XPaths and is underspecified in the
// published IxNetwork schema.
func TestMarshalXPathSchemaIssues(t *testing.T) {
	tests := []struct {
		desc    string
		xPath   *XPath
		wantStr string
	}{{
		desc: "Globals config, no array objects actually indexed.",
		xPath: &XPath{
			parentXPath: "/globals/topology/ancp",
			objectName:  "tlvEditor",
			index:       1,
		},
		wantStr: "/globals/topology/ancp/tlvEditor",
	}, {
		desc: "Always singular config node (transmissionDistribution).",
		xPath: &XPath{
			parentXPath: "/traffic/trafficItem[1]/configElement[1]",
			objectName:  "transmissionDistribution",
			index:       1,
		},
		wantStr: "/traffic/trafficItem[1]/configElement[1]/transmissionDistribution",
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			xPathStr := tc.xPath.String()
			if xPathStr != tc.wantStr {
				t.Fatalf("Output string did not match expected: got %q, want %s", xPathStr, tc.wantStr)
			}
		})
	}
}
