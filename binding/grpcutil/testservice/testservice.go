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

// Package testservice provides an interface for testing with the Test service.
package testservice

import (
	"io"
	"net"
	"testing"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/local"

	tgrpcpb "github.com/openconfig/ondatra/binding/grpcutil/testservice/gen"
	tpb "github.com/openconfig/ondatra/binding/grpcutil/testservice/gen"
)

// Start starts the gRCP server with the specified TestServer implementation,
// dials the server, and returns a handle to the test service.
func Start(ctx context.Context, t *testing.T, testSrv tgrpcpb.TestServer, opts ...grpc.DialOption) *TestService {
	t.Helper()
	srv := grpc.NewServer()
	tgrpcpb.RegisterTestServer(srv, testSrv)
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("net.Listen() failed: %v", err)
	}
	go srv.Serve(lis)

	opts = append([]grpc.DialOption{grpc.WithTransportCredentials(local.NewCredentials())}, opts...)
	conn, err := grpc.DialContext(ctx, lis.Addr().String(), opts...)
	if err != nil {
		t.Fatalf("grpc.DialContext() failed: %v", err)
	}
	return &TestService{TestClient: tgrpcpb.NewTestClient(conn), stopFn: srv.Stop}
}

// TestService represents a running test service client and server.
type TestService struct {
	tgrpcpb.TestClient
	stopFn func()
}

// MustSendUnary calls SendUnary and fails the test fatally on error.
func (ts *TestService) MustSendUnary(ctx context.Context, t *testing.T, msg string) {
	t.Helper()
	if _, err := ts.SendUnary(ctx, &tpb.TestRequest{Message: msg}); err != nil {
		t.Fatalf("SendUnary() failed: %v", err)
	}
}

// MustSendStream calls SendStream and fails the test fatally on error.
func (ts *TestService) MustSendStream(ctx context.Context, t *testing.T) *TestStreamClient {
	t.Helper()
	streamClient, err := ts.SendStream(ctx)
	if err != nil {
		t.Fatalf("SendStream() failed: %v", err)
	}
	return &TestStreamClient{streamClient}
}

// Stop stops the TestService server.
func (ts *TestService) Stop() {
	ts.stopFn()
}

// TestStreamClient is a SendStreamClient with additional Must methods.
type TestStreamClient struct {
	tgrpcpb.Test_SendStreamClient
}

// MustSend calls Send and fails the test fatally on error.
func (sc *TestStreamClient) MustSend(t *testing.T, msg string) {
	t.Helper()
	if err := sc.Send(&tpb.TestRequest{Message: msg}); err != nil {
		t.Fatalf("Send() failed: %v", err)
	}
}

// MustRecv calls Recv and fails the test fatally on a non-EOF error.
// Returns true if a message is received (if there is no error or an EOF error).
func (sc *TestStreamClient) MustRecv(t *testing.T) bool {
	t.Helper()
	if _, err := sc.Recv(); err != nil {
		if err == io.EOF {
			return false
		}
		t.Fatalf("Recv() failed: %v", err)
	}
	return true
}
