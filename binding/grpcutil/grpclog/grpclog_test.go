// Copyright 2026 Google LLC
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

package grpclog

import (
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding/grpcutil/testservice"
	tgrpcpb "github.com/openconfig/ondatra/binding/grpcutil/testservice/gen"
	tpb "github.com/openconfig/ondatra/binding/grpcutil/testservice/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestSanitizeTarget(t *testing.T) {
	tests := []struct {
		name   string
		target string
		want   string
	}{
		{
			name:   "valid target",
			target: "localhost:1234",
			want:   "localhost_1234",
		},
		{
			name:   "target with special chars",
			target: "unix:///tmp/grpc.sock",
			want:   "unix____tmp_grpc.sock",
		},
		{
			name:   "empty target",
			target: "",
			want:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sanitizeTarget(tt.target); got != tt.want {
				t.Errorf("sanitizeTarget(%q) = %q, want %q", tt.target, got, tt.want)
			}
		})
	}
}

type testServer struct {
	tgrpcpb.UnimplementedTestServer
}

func (*testServer) SendUnary(context.Context, *tpb.TestRequest) (*tpb.TestResponse, error) {
	return &tpb.TestResponse{}, nil
}

func (*testServer) SendStream(stream tgrpcpb.Test_SendStreamServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.Send(&tpb.TestResponse{})
		}
		if err != nil {
			return err
		}
	}
}

func TestInterceptors(t *testing.T) {
	ctx := context.Background()
	opts := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(UnaryClientInterceptor()),
		grpc.WithChainStreamInterceptor(StreamClientInterceptor()),
	}
	ts := testservice.Start(ctx, t, &testServer{}, opts...)
	defer ts.Stop()

	md := metadata.Pairs("key", "value")
	ctx = metadata.NewOutgoingContext(ctx, md)

	ts.MustSendUnary(ctx, t, "unary")
	stream := ts.MustSendStream(ctx, t)
	stream.MustSend(t, "stream")
	stream.CloseSend()
	if !stream.MustRecv(t) {
		t.Errorf("Expected one message from stream, got none")
	}
	if stream.MustRecv(t) {
		t.Errorf("Expected EOF from stream, but received a message")
	}

	logDir := testOutputsDir
	ents, err := os.ReadDir(logDir)
	if err != nil {
		t.Fatalf("Failed to read log dir: %v", err)
	}
	var logFiles []string
	for _, ent := range ents {
		if strings.HasPrefix(ent.Name(), "grpc-target-") && strings.HasSuffix(ent.Name(), ".log") {
			logFiles = append(logFiles, ent.Name())
		}
	}
	if len(logFiles) == 0 {
		t.Fatalf("Got 0 log files, want at least 1")
	}
	sort.Strings(logFiles)

	var gotLogs string
	for _, f := range logFiles {
		logFile := filepath.Join(logDir, f)
		data, err := os.ReadFile(logFile)
		if err != nil {
			t.Fatalf("Failed to read log file %s: %v", f, err)
		}
		gotLogs += string(data)
	}

	wantSubstrings := []string{
		// Unary call logs
		`CALL START: Method: /testservice.Test/SendUnary, Metadata: map[key:[value]], Request: message:"unary"`,
		`CALL END: Method: /testservice.Test/SendUnary`,
		// Stream call logs
		`CALL START: Method: /testservice.Test/SendStream, Metadata: map[key:[value]]`,
		`(Stream /testservice.Test/SendStream) SendMsg success: message:"stream"`,
		`(Stream /testservice.Test/SendStream) RecvMsg success:`,
		`(Stream /testservice.Test/SendStream) RecvMsg error: EOF`,
	}
	for _, want := range wantSubstrings {
		if !strings.Contains(gotLogs, want) {
			t.Errorf("Log files %v do not contain expected substring %q; combined content:\n%s", logFiles, want, gotLogs)
		}
	}
}
