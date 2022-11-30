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

// Package ixweb provides a connection to the Ixia Web Platform.
package ixweb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
)

var (
	// to be stubbed out by tests
	sleepFn = time.Sleep
)

type httpMethod string

const (
	get    = httpMethod(http.MethodGet)
	delete = httpMethod(http.MethodDelete)
	patch  = httpMethod(http.MethodPatch)
	post   = httpMethod(http.MethodPost)

	// RetryLimit is the maximum number of times an HTTP request will be retried.
	RetryLimit = 4
	// RetryDelay is the time to wait between retrying HTTP requests.
	RetryDelay = 15 * time.Second
	// pollLimit is the maximum number of times an operation will be polled.
	pollLimit = 500
)

// IxWeb is a connection to the Ixia Web Platform.
type IxWeb struct {
	hostname, apiKey string
	client           HTTPClient
}

// HTTPClient makes HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Option for connecting to IxNetwork.
type Option interface {
	set(cfg *config)
}

type option func(*config)

func (o option) set(cfg *config) { o(cfg) }

type config struct {
	username, password string
	httpClient         HTTPClient
}

func defaultConfig() *config {
	return &config{
		username:   "admin",
		password:   "admin",
		httpClient: http.DefaultClient,
	}
}

// WithLogin configures the IxWeb to use a custom username/password.
func WithLogin(username, password string) Option {
	return option(func(cfg *config) {
		cfg.username = username
		cfg.password = password
	})
}

// WithHTTPClient configures IxWeb to use a custom HTTP client.
func WithHTTPClient(client HTTPClient) Option {
	return option(func(cfg *config) {
		cfg.httpClient = client
	})
}

// Connect creates a connection to Ixia Web Platform on the specified hostname.
// Unless the given options specify otherwise, IxWeb is configured as follows:
//   - login: Ixia's default admin/admin login
//   - http client: http.DefaultClient
func Connect(ctx context.Context, hostname string, opts ...Option) (*IxWeb, error) {
	if hostname == "" {
		return nil, errors.New("no hostname specified")
	}
	if u, err := url.Parse(fmt.Sprintf("http://%s:443", hostname)); err != nil || u.Hostname() != hostname {
		return nil, fmt.Errorf("invalid hostname %q: %w", hostname, err)
	}
	cfg := defaultConfig()
	for _, o := range opts {
		o.set(cfg)
	}
	ix := &IxWeb{
		hostname: hostname,
		client:   cfg.httpClient,
	}
	return ix, ix.setAPIKey(ctx, cfg.username, cfg.password)
}

// IxNetwork returns a handle to the IxNetwork application.
func (ix *IxWeb) IxNetwork() *IxNetwork {
	return &IxNetwork{ixweb: ix}
}

// Chassis returns a handle to the Chassis application.
func (ix *IxWeb) Chassis() *Chassis {
	return &Chassis{ixweb: ix}
}

func (ix *IxWeb) setAPIKey(ctx context.Context, username, password string) error {
	in := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: username,
		Password: password,
	}
	out := struct {
		APIKey string `json:"apiKey"`
	}{}
	if err := ix.jsonReq(ctx, post, "/api/v1/auth/session", in, &out); err != nil {
		return fmt.Errorf("error requesting API key: %w", err)
	}
	ix.apiKey = out.APIKey
	return nil
}

// jsonReq issues a JSON HTTP request. The "in" JSON object is marshalled to
// the request body, and the response is unmarshalled to "out" JSON object.
func (ix *IxWeb) jsonReq(ctx context.Context, method httpMethod, path string, in, out any) error {
	var body []byte
	if in != nil {
		var err error
		body, err = json.Marshal(in)
		if err != nil {
			return fmt.Errorf("error marshalling JSON data: %w", err)
		}
	}
	status, data, err := ix.request(ctx, method, path, "application/json", body)
	if err != nil {
		return err
	}
	if status == 202 {
		return ix.waitForAsync(ctx, data, out)
	}
	return unmarshal(data, out)
}

// binaryReq issues a binary HTTP request.
func (ix *IxWeb) binaryReq(ctx context.Context, method httpMethod, path string, content []byte) ([]byte, error) {
	_, data, err := ix.request(ctx, method, path, "application/octet-stream", content)
	return data, err
}

func (ix *IxWeb) request(ctx context.Context, method httpMethod, path, contentType string, content []byte) (int, []byte, error) {
	url := fmt.Sprintf("https://%s%s", ix.hostname, path)
	var body io.Reader
	if len(content) > 0 {
		body = bytes.NewReader(content)
	}
	req, err := http.NewRequestWithContext(ctx, string(method), url, body)
	if err != nil {
		return 0, nil, fmt.Errorf("error creating HTTP request: %w", err)
	}
	if len(content) > 0 {
		req.Header.Set("Content-Type", contentType)
	}
	if ix.apiKey != "" {
		req.Header.Set("x-api-key", ix.apiKey)
	}

	var resp *http.Response
	for i := 0; ; i++ {
		// Get copy of the body in case a retry is necessary.
		var body io.ReadCloser
		if req.Body != nil {
			body, err = req.GetBody()
			if err != nil {
				return 0, nil, fmt.Errorf("could not copy body from HTTP request for retry: %w", err)
			}
		}

		// Attempt the HTTP request.
		retryCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		req = req.Clone(retryCtx)
		resp, err = ix.client.Do(req)
		if err == nil { // if no err, do not retry
			break
		}
		// If out of retries or if the original context is no longer active (there would be error from the 'Do' in the latter case.).
		if i == RetryLimit || ctx.Err() != nil {
			return 0, nil, fmt.Errorf("error on HTTP request %v: %w", req, err)
		}

		if ctx.Err() != nil {
			return 0, nil, fmt.Errorf("context finished before successful request: %w", ctx.Err())
		}

		// Reset the body before the next attempt.
		log.Warningf("Retrying %s %s due to error %v", req.Method, req.URL.Path, err)
		req.Body = body
		sleepFn(RetryDelay)
	}

	// The contract of 'Do' ensures Body is non-nil.
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, fmt.Errorf("could not read response body: %w", err)
	}
	log.V(1).Infof("Response to %q: %q", req.URL.Path, data)
	status := resp.StatusCode
	if status < 200 || status > 299 {
		return 0, nil, fmt.Errorf("error status code %d on request %+v, response: %q", status, req, data)
	}
	return status, data, nil
}

func (ix *IxWeb) waitForAsync(ctx context.Context, data []byte, out any) error {
	const pollDelay = 5 * time.Second
	var status struct {
		State  string          `json:"state"`
		URL    string          `json:"url"`
		Result json.RawMessage `json:"result"`
	}
	if err := unmarshal(data, &status); err != nil {
		return err
	}

	for i := 0; i < pollLimit; i++ {
		switch status.State {
		case "ERROR", "EXCEPTION":
			return fmt.Errorf("operation failed: %s", string(data))
		case "IN_PROGRESS":
			pollURL, err := url.Parse(status.URL)
			if err != nil {
				return fmt.Errorf("invalid operation poll URL %q: %w", status.URL, err)
			}
			sleepFn(pollDelay)
			if err := ix.jsonReq(ctx, get, pollURL.Path, nil, &status); err != nil {
				return err
			}
		case "COMPLETED", "SUCCESS":
			return unmarshal(status.Result, out)
		default:
			return fmt.Errorf("unknown state in response: %v", out)
		}
	}

	return fmt.Errorf("operation timeout; last status: %v", status)
}

func unmarshal(data []byte, out any) error {
	// Unmarshal will fail if data is empty or out is nil; make it a noop instead.
	if out != nil && len(data) > 0 {
		if err := json.Unmarshal(data, out); err != nil {
			return fmt.Errorf("error unmarshalling JSON %q: %w", string(data), err)
		}
	}
	return nil
}
