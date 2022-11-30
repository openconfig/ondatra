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

// Package httpovergrpc implements a http/https service over grpc transport.
// The server side implementation will be responsible for dialing the proxied
// destination with the proper http method and credentials and then proxying
// the responses back to the caller.
package httpovergrpc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context"

	log "github.com/golang/glog"

	closer "github.com/openconfig/gocloser"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"

	hpb "github.com/openconfig/ondatra/proxy/proto/httpovergrpc"
)

// Service provides the HttpProxy RPC service definitions.
type Service struct {
	hpb.UnimplementedHTTPOverGRPCServer
	client HTTPDoer
}

// Option defines the options to configure the HTTP over gRPC service.
type Option func(s *Service)

// HTTPDoer provides an interface for round tripping an http request response.
type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

// WithClient allows setting the http client for the HTTP over
// gRPC service.
func WithClient(client HTTPDoer) Option {
	return func(s *Service) {
		s.client = client
	}
}

// New creates a new HTTP over gRPC service.
func New(opts ...Option) *Service {
	s := &Service{client: http.DefaultClient}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// SendRequest sends the HTTP request to the path and returns the response
// received.
func (s *Service) HTTPRequest(ctx context.Context, req *hpb.Request) (*hpb.Response, error) {
	ctx, _ = tag.New(ctx, tag.Upsert(urlKey, req.GetMethod()))
	stats.Record(ctx, mNumRequests.M(1))
	request, err := http.NewRequest(req.GetMethod(), req.GetUrl(), bytes.NewReader(req.GetBody()))
	if err != nil {
		return nil, fmt.Errorf("error creating a new HTTP request: %v", err)
	}
	// Add all header values
	for _, header := range req.GetHeaders() {
		for _, v := range header.GetValues() {
			request.Header.Add(header.GetKey(), v)
		}
	}

	log.Infof("Request method: %s URL: %s body:\n%s", req.GetMethod(), req.GetUrl(), string(req.GetBody()))
	response, err := s.client.Do(request)
	if err != nil {
		mut := []tag.Mutator{tag.Upsert(errorKey, err.Error())}
		stats.RecordWithTags(ctx, mut, mClientErrors.M(1))
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}

	defer closer.CloseAndLog(response.Body.Close, "failed to close response body")
	body, err := io.ReadAll(response.Body)
	if err != nil {
		mut := []tag.Mutator{tag.Upsert(errorKey, err.Error())}
		stats.RecordWithTags(ctx, mut, mProxyErrors.M(1))
		return nil, fmt.Errorf("failed to read HTTP response: %w", err)
	}
	log.Infof("Response status: %q body:\n%s", response.Status, string(body))

	var rHeaders []*hpb.Header
	for k, vals := range response.Header {
		h := &hpb.Header{Key: k}
		for _, val := range vals {
			h.Values = append(h.Values, val)
		}
		rHeaders = append(rHeaders, h)
	}
	return &hpb.Response{
		Status:  int32(response.StatusCode),
		Body:    body,
		Headers: rHeaders,
	}, nil
}
