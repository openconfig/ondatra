// Copyright 2021 Google LLC
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

package ixweb

import (
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
)

func TestExport(t *testing.T) {
	tests := []struct {
		desc    string
		doResp  *http.Response
		wantCfg string
		wantErr string
	}{{
		desc:    "error",
		doResp:  fakeResponse(500, `"POST error"`),
		wantErr: "POST error",
	}, {
		desc:    "success",
		doResp:  fakeResponse(200, `"myCfg"`),
		wantCfg: "myCfg",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			config := &Config{sess: &Session{ixweb: &IxWeb{client: &fakeHTTPClient{
				doResps: []*http.Response{test.doResp},
			}}}}
			gotCfg, gotErr := config.Export(context.Background())
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("Export: unexpected error, got err %v, want err %q", gotErr, test.wantErr)
			}
			if diff := cmp.Diff(test.wantCfg, gotCfg); diff != "" {
				t.Errorf("Export: unexpected config, diff: (-want +got)\n%s.", diff)
			}
		})
	}
}

func TestImport(t *testing.T) {
	tests := []struct {
		desc    string
		doResp  *http.Response
		wantErr string
	}{{
		desc:    "error",
		doResp:  fakeResponse(500, `"POST error"`),
		wantErr: "POST error",
	}, {
		desc:   "success",
		doResp: fakeResponse(200, `"myCfg"`),
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			config := &Config{sess: &Session{ixweb: &IxWeb{client: &fakeHTTPClient{
				doResps: []*http.Response{test.doResp},
			}}}}
			gotErr := config.Import(context.Background(), "myCfg", false)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("Import: unexpected error, got err %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestQueryIDs(t *testing.T) {
	xpaths := []string{
		"/topology[1]",
		"/topology[1]/deviceGroup[1]",
	}

	tests := []struct {
		desc    string
		doResp  *http.Response
		wantIDs map[string]string
		wantErr string
	}{{
		desc:    "error",
		doResp:  fakeResponse(500, "POST error"),
		wantErr: "POST error",
	}, {
		desc:    "empty response",
		doResp:  fakeResponse(200, "[]"),
		wantErr: "expected one element",
	}, {
		desc: "missing ID",
		doResp: fakeResponse(200, `[{
		  "href":  "/id/to/root",
		  "xpath": "/xpath/to/nowhere"
		}]`),
		wantErr: "did not find",
	}, {
		desc: "success",
		doResp: fakeResponse(200, `[{
		  "href":  "/id/to/root",
		  "xpath": "/",
		  "topology": [{
		  	"href":  "/id/to/topology",
			  "xpath": "/topology[1]",
			  "deviceGroup": [{
				  "href":  "/id/to/deviceGroup",
				  "xpath": "/topology[1]/deviceGroup[1]"
			  }]
		  }]
	  }]`),
		wantIDs: map[string]string{
			"/topology[1]":                "/id/to/topology",
			"/topology[1]/deviceGroup[1]": "/id/to/deviceGroup",
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			config := &Config{sess: &Session{ixweb: &IxWeb{client: &fakeHTTPClient{
				doResps: []*http.Response{test.doResp},
			}}}}
			gotIDs, gotErr := config.QueryIDs(context.Background(), xpaths...)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("QueryIDs: unexpected error, got err %v, want err %q", gotErr, test.wantErr)
			}
			if diff := cmp.Diff(test.wantIDs, gotIDs); diff != "" {
				t.Errorf("QueryIDs: unexpected diff in IDs (-want,+got): %s\n", diff)
			}
		})
	}
}
