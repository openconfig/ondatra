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
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func init() {
	sleepFn = func(time.Duration) {}
}

func TestConnect(t *testing.T) {
	const (
		hostname = "fakeHost"
		apiKey   = "fakeAPIKey"
	)
	successBody := fmt.Sprintf(`{"apiKey": "%s"}`, apiKey)

	tests := []struct {
		desc    string
		doResps []*http.Response
		doErrs  []error
		wantErr string
	}{{
		desc:    "success",
		doResps: []*http.Response{fakeResponse(200, successBody)},
	}, {
		desc:    "success after retry",
		doResps: []*http.Response{fakeResponse(200, "blah"), fakeResponse(200, successBody)},
		doErrs:  []error{errors.New("unknown error"), nil},
	}, {
		desc:    "bad status code",
		doResps: []*http.Response{fakeResponse(500, successBody)},
		wantErr: "status code",
	}, {
		desc:    "malformed JSON",
		doResps: []*http.Response{fakeResponse(200, "blah")},
		wantErr: "unmarshalling JSON",
	}, {
		desc:    "persistent error",
		doErrs:  repeatErrors(errors.New("unknown error"), RetryLimit+1),
		wantErr: "unknown error",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			fakeClient := &fakeHTTPClient{doResps: test.doResps, doErrs: test.doErrs}
			ix, err := Connect(context.Background(), hostname, WithHTTPClient(fakeClient))
			if got, want := err != nil, test.wantErr != ""; got != want {
				t.Fatalf("Connect() got err %v, want err %v", err, want)
			}
			if err != nil {
				if !strings.Contains(err.Error(), test.wantErr) {
					t.Errorf("Connect() got err %v, want contains %q", err, test.wantErr)
				}
				return
			}
			if ix.hostname != hostname {
				t.Errorf("Connect() got host = %s, want %s", ix.hostname, hostname)
			}
			if ix.apiKey != apiKey {
				t.Errorf("Connect() got apiKey = %s, want %s", ix.apiKey, apiKey)
			}
		})
	}
}

func TestAsyncOperation(t *testing.T) {
	successBody := `{"state":"SUCCESS"}`
	completedBody := `{"state":"COMPLETED"}`
	progressBody := `{"state":"IN_PROGRESS"}`
	errorBody := `{"state":"ERROR", "result":"error response"}`
	exceptionBody := `{"state":"EXCEPTION", "message":"exception response"}`

	tests := []struct {
		desc    string
		doResps []*http.Response
		wantErr string
	}{{
		desc:    "success",
		doResps: []*http.Response{fakeResponse(200, successBody)},
	}, {
		desc:    "completed",
		doResps: []*http.Response{fakeResponse(200, completedBody)},
	}, {
		desc:    "success after progress",
		doResps: []*http.Response{fakeResponse(202, progressBody), fakeResponse(200, successBody)},
	}, {
		desc:    "completed after progress",
		doResps: []*http.Response{fakeResponse(202, progressBody), fakeResponse(200, completedBody)},
	}, {
		desc: "progress timeout",
		doResps: append(append([]*http.Response{},
			fakeResponse(202, progressBody)),
			repeatResponses(200, progressBody, pollLimit)...),
		wantErr: "timeout",
	}, {
		desc:    "error",
		doResps: []*http.Response{fakeResponse(202, errorBody)},
		wantErr: "error response",
	}, {
		desc:    "exception",
		doResps: []*http.Response{fakeResponse(202, exceptionBody)},
		wantErr: "exception response",
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ix := &IxWeb{client: &fakeHTTPClient{doResps: test.doResps}}
			err := ix.jsonReq(context.Background(), post, "/my/path", nil, nil)
			if got, want := err != nil, test.wantErr != ""; got != want {
				t.Fatalf("request() got err %v, want err %v", err, want)
			}
			if err != nil && !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("request() got err %v, want contains %q", err, test.wantErr)
			}
		})
	}
}

type fakeHTTPClient struct {
	doResps []*http.Response
	doErrs  []error
}

func (c *fakeHTTPClient) Do(req *http.Request) (*http.Response, error) {
	var doResp *http.Response
	if len(c.doResps) > 0 {
		doResp, c.doResps = c.doResps[0], c.doResps[1:]
	}
	var doErr error
	if len(c.doErrs) > 0 {
		doErr, c.doErrs = c.doErrs[0], c.doErrs[1:]
	}
	if doResp == nil && doErr == nil {
		return nil, fmt.Errorf("No stub response available for req: %v", req)
	}
	return doResp, doErr
}

func repeatResponses(status int, content string, n int) []*http.Response {
	resps := make([]*http.Response, n)
	for i := 0; i < n; i++ {
		resps[i] = fakeResponse(status, content)
	}
	return resps
}

func repeatErrors(err error, n int) []error {
	errs := make([]error, n)
	for i := 0; i < n; i++ {
		errs[i] = err
	}
	return errs
}

func fakeResponse(status int, content string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body: &fakeBody{
			Reader: bytes.NewReader([]byte(content)),
		},
	}
}

type fakeBody struct {
	io.Reader
	closed bool
}

func (b *fakeBody) Read(p []byte) (n int, err error) {
	if b.closed {
		return 0, errors.New("closed")
	}
	return b.Reader.Read(p)
}

func (b *fakeBody) Close() error {
	b.closed = true
	return nil
}
