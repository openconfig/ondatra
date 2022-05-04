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
	"golang.org/x/net/context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
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
