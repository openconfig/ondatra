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

package rawapis

import (
	"golang.org/x/net/context"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/local"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	tgrpcpb "github.com/openconfig/ondatra/internal/rawapis/testservice"
	tpb "github.com/openconfig/ondatra/internal/rawapis/testservice"
)

func TestAnnotateErrors(t *testing.T) {
	interceptSrv := new(interceptServer)
	srv := grpc.NewServer()
	defer srv.Stop()
	tgrpcpb.RegisterTestServer(srv, interceptSrv)
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("error listening on localhost: %v", err)
	}
	defer lis.Close()
	go srv.Serve(lis)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, lis.Addr().String(),
		grpc.WithTransportCredentials(local.NewCredentials()),
		withUnaryAnnotateErrors(),
		withStreamAnnotateErrors(),
	)
	if err != nil {
		t.Fatalf("error dialing server: %v", err)
	}
	interceptClient := tgrpcpb.NewTestClient(conn)

	t.Run("unary success", func(t *testing.T) {
		interceptSrv.unaryErr = nil
		if _, gotErr := interceptClient.SendUnary(ctx, new(tpb.TestRequest)); gotErr != nil {
			t.Errorf("unaryInterceptErr got error %v, want no error", gotErr)
		}
	})

	t.Run("unary error", func(t *testing.T) {
		const (
			wantCode   = codes.NotFound
			wantErrMsg = "unary error"
			wantReqMsg = "unary msg"
		)
		interceptSrv.unaryErr = status.Error(wantCode, wantErrMsg)
		_, gotErr := interceptClient.SendUnary(ctx, &tpb.TestRequest{Message: wantReqMsg})
		if gotErr == nil {
			t.Fatalf("unaryInterceptErr got no error, want error, %v", gotErr)
		}
		if diff := grpcErrDiff(t, gotErr, wantCode, wantErrMsg, wantReqMsg); diff != "" {
			t.Errorf("unaryInterceptErr got grpc error diff:\n%s", diff)
		}
	})

	t.Run("unary unknown error", func(t *testing.T) {
		const (
			wantCode   = codes.Unknown
			wantErrMsg = "unknown error"
			wantReqMsg = "unknown msg"
		)
		interceptSrv.unaryErr = fmt.Errorf(wantErrMsg)
		_, gotErr := interceptClient.SendUnary(ctx, &tpb.TestRequest{Message: wantReqMsg})
		if gotErr == nil {
			t.Fatalf("unaryInterceptErr got no error, want error, %v", gotErr)
		}
		if diff := grpcErrDiff(t, gotErr, wantCode, wantErrMsg, wantReqMsg); diff != "" {
			t.Errorf("unaryInterceptErr got grpc error diff:\n%s", diff)
		}
	})

	t.Run("stream recv eof", func(t *testing.T) {
		interceptSrv.streamErr = nil
		streamClient := mustStartStream(t, interceptClient)
		mustSendMessage(t, streamClient, "msg")
		if _, gotErr := streamClient.Recv(); gotErr != io.EOF {
			t.Errorf("streamInterceptErr got error %v, want %v", gotErr, io.EOF)
		}
	})

	t.Run("stream recv no send", func(t *testing.T) {
		const (
			wantCode   = codes.PermissionDenied
			wantErrMsg = "stream error"
		)
		interceptSrv.streamErr = status.Error(wantCode, wantErrMsg)
		streamClient := mustStartStream(t, interceptClient)
		_, gotErr := streamClient.Recv()
		if gotErr == nil {
			t.Fatalf("streamInterceptErr got no error, want error, %v", gotErr)
		}
		if diff := grpcErrDiff(t, gotErr, wantCode, wantErrMsg); diff != "" {
			t.Errorf("streamInterceptErr got grpc error diff:\n%s", diff)
		}
	})

	t.Run("stream recv after send", func(t *testing.T) {
		const (
			wantCode   = codes.PermissionDenied
			wantErrMsg = "stream error"
			wantReqMsg = "stream msg"
		)
		interceptSrv.streamErr = status.Error(wantCode, wantErrMsg)
		streamClient := mustStartStream(t, interceptClient)
		mustSendMessage(t, streamClient, wantReqMsg)
		_, gotErr := streamClient.Recv()
		if gotErr == nil {
			t.Fatalf("streamInterceptErr got no error, want error, %v", gotErr)
		}
		if diff := grpcErrDiff(t, gotErr, wantCode, wantErrMsg, wantReqMsg); diff != "" {
			t.Errorf("streamInterceptErr got grpc error diff:\n%s", diff)
		}
	})
}

type interceptServer struct {
	tgrpcpb.UnimplementedTestServer
	unaryErr  error
	streamErr error
}

func (s *interceptServer) SendUnary(context.Context, *tpb.TestRequest) (*tpb.TestResponse, error) {
	return new(tpb.TestResponse), s.unaryErr
}

func (s *interceptServer) SendStream(srv tgrpcpb.Test_SendStreamServer) error {
	return s.streamErr
}

func mustStartStream(t *testing.T, client tgrpcpb.TestClient) tgrpcpb.Test_SendStreamClient {
	streamClient, err := client.SendStream(context.Background())
	if err != nil {
		t.Fatalf("SendStream failed with error: %v", err)
	}
	return streamClient
}

func mustSendMessage(t *testing.T, streamClient tgrpcpb.Test_SendStreamClient, msg string) {
	if err := streamClient.Send(&tpb.TestRequest{Message: msg}); err != nil {
		t.Fatalf("Send failed with error: %v", err)
	}
}

func grpcErrDiff(t *testing.T, err error, wantCode codes.Code, wantSubstrs ...string) string {
	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("status.FromError failed to convert error: %v", err)
	}
	gotCode, gotMsg := st.Code(), st.Message()
	var lines []string
	if gotCode != wantCode {
		lines = append(lines, fmt.Sprintf("want code %v, got %v", wantCode, gotCode))
	}
	for _, wantSub := range wantSubstrs {
		if !strings.Contains(gotMsg, wantSub) {
			lines = append(lines, fmt.Sprintf("want message containing %q, got %q", wantSub, gotMsg))
		}
	}
	return strings.Join(lines, "\n")
}
