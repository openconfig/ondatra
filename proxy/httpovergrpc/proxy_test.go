// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package httpovergrpc

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/openconfig/gnmi/errdiff"
	hpb "github.com/openconfig/ondatra/proxy/proto/httpovergrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/local"
	"google.golang.org/protobuf/testing/protocmp"
)

type fakeHTTPServer struct {
	req          *http.Request
	responseCode int
	responseBody string
	err          error
}

func (s *fakeHTTPServer) RoundTrip(req *http.Request) (*http.Response, error) {
	s.req = req

	if s.err != nil {
		return nil, s.err
	}
	resp := &http.Response{
		Status:     fmt.Sprintf("%d %s", s.responseCode, http.StatusText(s.responseCode)),
		StatusCode: s.responseCode,
		Body:       io.NopCloser(strings.NewReader(s.responseBody)),
		Header:     req.Header,
	}
	return resp, nil
}

func TestSendRequest(t *testing.T) {
	tests := []struct {
		desc            string
		req             *hpb.Request
		expectedMethod  string
		expectedHeaders map[string]string
		responseCode    int
		responseBody    string
		responseError   error
		wantErr         string
	}{{
		desc: "simple get request",
		req: &hpb.Request{
			Method: "GET",
			Url:    "www.google.com",
		},
		expectedMethod:  "GET",
		expectedHeaders: map[string]string{},
		responseCode:    200,
		responseBody:    "content",
		responseError:   nil,
	}, {
		desc: "simple post request",
		req: &hpb.Request{
			Method: "POST",
			Url:    "www.google.com",
			Body:   []byte("text"),
			Headers: []*hpb.Header{
				{Key: "Content-Type", Values: []string{"application/json"}},
			},
		},
		expectedMethod:  "POST",
		expectedHeaders: map[string]string{"Content-Type": "application/json"},
		responseCode:    400,
		responseBody:    "content",
		responseError:   nil,
	}, {
		desc: "simple patch request",
		req: &hpb.Request{
			Method: "PATCH",
			Url:    "www.google.com",
			Body:   []byte("text"),
			Headers: []*hpb.Header{
				{Key: "X-Api-Key", Values: []string{"key"}},
				{Key: "Content-Type", Values: []string{"application/json"}},
			},
		},
		expectedMethod: "PATCH",
		expectedHeaders: map[string]string{
			"Content-Type": "application/json",
			"x-api-key":    "key",
		},
		responseCode:  500,
		responseBody:  "content",
		responseError: nil,
	}, {
		desc: "simple delete request",
		req: &hpb.Request{
			Method: "DELETE",
			Url:    "www.google.com",
			Body:   nil,
			Headers: []*hpb.Header{
				{Key: "X-Api-Key", Values: []string{"key"}},
			},
		},
		expectedMethod: "DELETE",
		expectedHeaders: map[string]string{
			"x-api-key": "key",
		},
		responseCode:  200,
		responseBody:  "",
		responseError: nil,
	}, {
		desc: "unknown method",
		req: &hpb.Request{
			Method: "FAKEPOST",
			Url:    "www.google.com",
			Body:   nil,
		},
		expectedMethod:  "FAKEPOST",
		expectedHeaders: map[string]string{},
		responseCode:    0,
		responseBody:    "",
		responseError:   nil,
	}, {
		desc: "response error",
		req: &hpb.Request{
			Method: "GET",
			Url:    "www.google.com",
			Body:   nil,
		},
		expectedMethod:  "GET",
		expectedHeaders: map[string]string{},
		responseCode:    0,
		responseBody:    "",
		responseError:   fmt.Errorf("error"),
		wantErr:         "failed to send HTTP request",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			fakeServer := &fakeHTTPServer{
				responseCode: tt.responseCode,
				responseBody: tt.responseBody,
				err:          tt.responseError,
			}
			service := New(WithClient(&http.Client{Transport: fakeServer}))
			srv := grpc.NewServer(grpc.Creds(local.NewCredentials()))
			hpb.RegisterHTTPOverGRPCServer(srv, service)
			addr := startServer(t, srv)
			conn, err := grpc.DialContext(context.Background(), addr, grpc.WithTransportCredentials(local.NewCredentials()))
			if err != nil {
				t.Fatal(err)
			}
			c := hpb.NewHTTPOverGRPCClient(conn)
			resp, err := c.HTTPRequest(context.Background(), tt.req)
			if s := errdiff.Check(err, tt.wantErr); s != "" {
				t.Fatalf("HTTPRequest(%s) unexpected error: %s", tt.req, s)
			}
			if tt.wantErr != "" {
				return
			}
			req := fakeServer.req
			if req.Method != tt.expectedMethod {
				t.Errorf("req.method got %s, want %s", req.Method, tt.expectedMethod)
			}
			if req.URL.String() != tt.req.Url {
				t.Errorf("req.url got %s, want %s", req.URL.String(), tt.req.Url)
			}
			for key, expectedValue := range tt.expectedHeaders {
				if value := req.Header.Get(key); value != expectedValue {
					t.Errorf("req.headers[%s]: got %s, want %s", key, value, expectedValue)
				}
			}
			if resp.GetStatus() != int32(tt.responseCode) {
				t.Errorf("HTTPRequest(%v) resp.status got %d, want %d", tt.req, resp.GetStatus(), tt.responseCode)
			}
			if string(resp.GetBody()) != tt.responseBody {
				t.Errorf("HTTPRequest(%v) resp.body got %s, want %s", tt.req, resp.GetBody(), tt.responseBody)
			}
			if s := cmp.Diff(resp.GetHeaders(), tt.req.GetHeaders(), protocmp.Transform(), cmpopts.SortSlices(func(a, b *hpb.Header) bool { return a.Key < b.Key })); s != "" {
				t.Fatalf("Unexpected headers: %s", s)
			}
		})
	}
}

func startServer(tb testing.TB, s *grpc.Server) string {
	tb.Helper()
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		tb.Fatalf("failed to listen: %v", err)
	}
	tb.Cleanup(func() { lis.Close() })
	tb.Cleanup(s.Stop)
	go func() {
		if err := s.Serve(lis); err != nil {
			tb.Logf("server error: %v", err)
		}
	}()
	return lis.Addr().String()
}
