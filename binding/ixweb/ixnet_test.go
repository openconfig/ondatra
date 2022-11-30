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
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestOpArgs(t *testing.T) {
	const want1, want2 = "a1", "a2"
	input := OpArgs{want1, want2}
	bytes, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("OpArgs got error marshaling: %v", err)
	}
	output := struct {
		Arg1 string `json:"arg1"`
		Arg2 string `json:"arg2"`
	}{}
	if err := json.Unmarshal(bytes, &output); err != nil {
		t.Fatalf("OpArgs got error unmarshaling: %v", err)
	}
	if got1, got2 := output.Arg1, output.Arg2; got1 != want1 || got2 != want2 {
		t.Errorf("OpArgs unmarshal got [%s, %s], want [%s, %s]", got1, got2, want1, want2)
	}
}

func TestFetchSessions(t *testing.T) {
	tests := []struct {
		desc         string
		doResp       *http.Response
		wantSessions map[int]string
		wantErr      string
	}{{
		desc:   "success",
		doResp: fakeResponse(200, `[{"id": 1, "sessionName": "session1"}, {"id": 100, "sessionName": "session100"}]`),
		wantSessions: map[int]string{
			1:   "session1",
			100: "session100",
		},
	}, {
		desc:    "error - 500 on fetch",
		doResp:  fakeResponse(500, ""),
		wantErr: "500",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ixn := &IxNetwork{ixweb: &IxWeb{
				client: &fakeHTTPClient{doResps: []*http.Response{test.doResp}},
			}}
			gotSessions, err := ixn.FetchSessions(context.Background())
			if got, want := err != nil, test.wantErr != ""; got != want {
				t.Fatalf("FetchSessions() got err %v, want err %v", err, want)
			}

			if err != nil {
				if !strings.Contains(err.Error(), test.wantErr) {
					t.Errorf("FetchSessions() got err %v, want contains %q", err, test.wantErr)
				}
			}

			got := map[int]string{}
			for _, s := range gotSessions {
				got[s.ID()] = s.Name()
			}
			if diff := cmp.Diff(test.wantSessions, got, cmpopts.EquateEmpty()); diff != "" {
				t.Errorf("FetchSessions() unexpected sessions returned (+want,-got): %s", diff)
			}
		})
	}
}

func TestDeleteSession(t *testing.T) {
	alreadyStoppedBody := `{"error": "Stopping a session is not permitted in state 'Stopped'"}`
	tests := []struct {
		desc    string
		doResps []*http.Response
		wantErr string
	}{{
		desc:    "success",
		doResps: []*http.Response{fakeResponse(200, ""), fakeResponse(200, "")},
	}, {
		desc:    "success - 404 on stop",
		doResps: []*http.Response{fakeResponse(404, ""), fakeResponse(200, "")},
	}, {
		desc:    "success - session already stopped",
		doResps: []*http.Response{fakeResponse(200, alreadyStoppedBody), fakeResponse(200, "")},
	}, {
		desc:    "error - 500 on stop",
		doResps: []*http.Response{fakeResponse(500, "")},
		wantErr: "500",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ixn := &IxNetwork{ixweb: &IxWeb{
				client: &fakeHTTPClient{doResps: test.doResps},
			}}
			err := ixn.DeleteSession(context.Background(), 1)
			if got, want := err != nil, test.wantErr != ""; got != want {
				t.Fatalf("DeleteSession() got err %v, want err %v", err, want)
			}
			if err != nil {
				if !strings.Contains(err.Error(), test.wantErr) {
					t.Errorf("DeleteSession() got err %v, want contains %q", err, test.wantErr)
				}
				return
			}
		})
	}
}

func TestErrors(t *testing.T) {
	errorsBody := `[{
		"id": 3,
		"lastModified": "05/11/2022 11:01:31",
		"errorLevel": "kError",
		"sourceColumnsDisplayName": [
			"Description",
			"Time"
		],
		"description": "some description",
		"provider": "ResourceManagerMiddleware",
		"errorCode": 0,
		"name": "JSON Import Issues"
	}]`
	errorInstancesBody := `[{
		"sourceValues": [
			"invalid IP address",
			"05/11/2022 11:01:31"
		]
	}]`
	invalidValueCountInstancesBody := `[{
		"sourceValues": [
			"invalid IP address"
		]
	}]`
	tests := []struct {
		desc       string
		doResps    []*http.Response
		wantErr    string
		wantErrors []*Error
	}{{
		desc:    "error fetching errors",
		doResps: []*http.Response{fakeResponse(400, "")},
		wantErr: "400",
	}, {
		desc:    "error fetching error instances",
		doResps: []*http.Response{fakeResponse(200, errorsBody), fakeResponse(500, "")},
		wantErr: "500",
	}, {
		desc:    "success",
		doResps: []*http.Response{fakeResponse(200, errorsBody), fakeResponse(200, invalidValueCountInstancesBody)},
		wantErr: "incorrect number of data values",
	}, {
		desc:    "success",
		doResps: []*http.Response{fakeResponse(200, errorsBody), fakeResponse(200, errorInstancesBody)},
		wantErrors: []*Error{{
			ID:                  3,
			LastModified:        "05/11/2022 11:01:31",
			Level:               "kError",
			Code:                0,
			Description:         "some description",
			Name:                "JSON Import Issues",
			InstanceColumnNames: []string{"Description", "Time"},
			Provider:            "ResourceManagerMiddleware",
			InstanceRowValues: []map[string]string{{
				"Description": "invalid IP address",
				"Time":        "05/11/2022 11:01:31",
			}},
		}},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			sess := &Session{ixweb: &IxWeb{
				hostname: "fakeHost",
				client: &fakeHTTPClient{
					doResps: test.doResps,
				},
			}}
			gotErrors, err := sess.Errors(context.Background())
			if got, want := err != nil, test.wantErr != ""; got != want {
				t.Fatalf("Errors() got err %v, want err %v", err, want)
			}
			if err != nil {
				if !strings.Contains(err.Error(), test.wantErr) {
					t.Errorf("Errors() got err %v, want contains %q", err, test.wantErr)
				}
				return
			}
			if diff := cmp.Diff(test.wantErrors, gotErrors); diff != "" {
				t.Errorf("Errors() got unexpected diff (-want + got): %s", diff)
			}
		})
	}
}
