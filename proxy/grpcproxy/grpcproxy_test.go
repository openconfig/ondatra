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
package grpcproxy

import (
	"context"
	"errors"
	"io"
	"net"
	"strings"
	"testing"

	"github.com/openconfig/ondatra/internal/fakegnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
)

type dialer struct{}

func (d *dialer) DialGRPC(ctx context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return grpc.DialContext(ctx, target, opts...)
}

func TestProxyChain(t *testing.T) {
	fake, err := fakegnmi.StartStubGNMI(0)
	if err != nil {
		fake.Stop()
		t.Fatalf("failed to start device: %v", err)
	}
	defer fake.Stop()
	deviceAddr := fake.Addr()

	p1Lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to start proxy 1 listener: %v", err)
	}
	p1Addr := p1Lis.Addr().String()

	p2Lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to start proxy 2 listener: %v", err)
	}
	p2Addr := p2Lis.Addr().String()

	dp1 := func(_ metadata.MD, _ string) (string, metadata.MD, error) {
		return p2Addr, nil, nil
	}
	p1, err := NewServer(dp1, WithDialer(&dialer{}))
	if err != nil {
		t.Fatalf("failed to create new gRPC proxy server: %v", err)
	}
	go func() {
		if err := p1.Serve(p1Lis); err != nil {
			t.Errorf("proxy 1 failed to serve: %v", err)
		}
		defer p1.Stop()
	}()

	dp2 := func(_ metadata.MD, _ string) (string, metadata.MD, error) {
		return deviceAddr, nil, nil
	}
	p2, err := NewServer(dp2, WithDialer(&dialer{}))
	if err != nil {
		t.Fatalf("failed to create new gRPC proxy server: %v", err)
	}
	go func() {
		if err := p2.Serve(p2Lis); err != nil {
			t.Errorf("proxy 2 failed to serve: %v", err)
		}
		defer p2.Stop()
	}()

	conn, err := grpc.DialContext(context.Background(), p1Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial proxy: %v", err)
	}
	fake.Stub().GetResponse(&gnmipb.GetResponse{
		Notification: []*gnmipb.Notification{{
			Timestamp: 1,
			Update: []*gnmipb.Update{{
				Path: &gnmipb.Path{
					Elem: []*gnmipb.PathElem{{
						Name: "interfaces",
					}},
				},
				Val: &gnmipb.TypedValue{
					Value: &gnmipb.TypedValue_StringVal{
						StringVal: "test",
					},
				},
			}},
		},
		}},
	).Notification(&gnmipb.Notification{
		Timestamp: 1,
		Update: []*gnmipb.Update{{
			Path: &gnmipb.Path{
				Elem: []*gnmipb.PathElem{{
					Name: "interfaces",
				}},
			},
			Val: &gnmipb.TypedValue{
				Value: &gnmipb.TypedValue_StringVal{
					StringVal: "test",
				},
			},
		}},
	}).Sync()
	c := gnmipb.NewGNMIClient(conn)
	_, err = c.Get(context.Background(), &gnmipb.GetRequest{})
	if err != nil {
		t.Fatalf("Get error: %v", err)
	}
	stream, err := c.Subscribe(context.Background())
	stream.Send(&gnmipb.SubscribeRequest{})
	if err != nil {
		t.Fatalf("Send failed error: %v", err)
	}
	wantResps := []*gnmipb.SubscribeResponse{}
	for i := 0; i < len(wantResps); i++ {
		v, err := stream.Recv()
		if err != nil {
			t.Fatalf("Subscribe failed: %v", err)
		}
		if !proto.Equal(v, wantResps[i]) {
			t.Fatalf("invalid message: got %s, want %s", v, wantResps[i])
		}
	}
}

func TestProxyError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode codes.Code
	}{
		{
			name:     "Canceled",
			err:      status.Error(codes.Canceled, "context canceled"),
			wantCode: codes.Canceled,
		},
		{
			name:     "DeadlineExceeded",
			err:      status.Error(codes.DeadlineExceeded, "context deadline exceeded"),
			wantCode: codes.DeadlineExceeded,
		},
		{
			name:     "Unavailable",
			err:      status.Error(codes.Unavailable, "connection closed"),
			wantCode: codes.Unavailable,
		},
		{
			name:     "NonGRPCErrorDefaultsToInternal",
			err:      errors.New("unexpected failure"),
			wantCode: codes.Internal,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotErr := proxyError(tc.err)
			if gotCode := status.Code(gotErr); gotCode != tc.wantCode {
				t.Errorf("proxyError(%v) code = %v, want %v", tc.err, gotCode, tc.wantCode)
			}
		})
	}
}

type fakeClientStream struct {
	grpc.ClientStream
	headerErr    error
	sendMsgErr   error
	closeSendErr error
}

func (f *fakeClientStream) Header() (metadata.MD, error) {
	return nil, f.headerErr
}

func (f *fakeClientStream) SendMsg(any) error {
	return f.sendMsgErr
}

func (f *fakeClientStream) CloseSend() error {
	return f.closeSendErr
}

type fakeServerStream struct {
	grpc.ServerStream
	recvMsgErr    error
	sendHeaderErr error
}

func (f *fakeServerStream) RecvMsg(any) error {
	return f.recvMsgErr
}

func (f *fakeServerStream) SendHeader(metadata.MD) error {
	return f.sendHeaderErr
}

func TestForwardServerToClientErrors(t *testing.T) {
	tests := []struct {
		name     string
		dst      *fakeClientStream
		src      *fakeServerStream
		wantCode codes.Code
	}{
		{
			name:     "RecvMsgCanceled",
			dst:      &fakeClientStream{},
			src:      &fakeServerStream{recvMsgErr: status.Error(codes.Canceled, "client canceled")},
			wantCode: codes.Canceled,
		},
		{
			name:     "CloseSendError",
			dst:      &fakeClientStream{closeSendErr: status.Error(codes.Unavailable, "close send failed")},
			src:      &fakeServerStream{recvMsgErr: io.EOF},
			wantCode: codes.Unavailable,
		},
		{
			name:     "SendMsgError",
			dst:      &fakeClientStream{sendMsgErr: status.Error(codes.Canceled, "send msg failed")},
			src:      &fakeServerStream{},
			wantCode: codes.Canceled,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := forwardServerToClient(tc.dst, tc.src)
			if gotCode := status.Code(err); gotCode != tc.wantCode {
				t.Errorf("forwardServerToClient() code = %v, want %v (err: %v)", gotCode, tc.wantCode, err)
			}
			if err == nil || !strings.Contains(err.Error(), "proxy_error:") {
				t.Errorf("forwardServerToClient() err = %v, want substring %q", err, "proxy_error:")
			}
		})
	}
}

func TestForwardClientToServerErrors(t *testing.T) {
	tests := []struct {
		name     string
		dst      *fakeServerStream
		src      *fakeClientStream
		wantCode codes.Code
	}{
		{
			name:     "HeaderError",
			dst:      &fakeServerStream{},
			src:      &fakeClientStream{headerErr: status.Error(codes.Canceled, "header canceled")},
			wantCode: codes.Canceled,
		},
		{
			name:     "SendHeaderError",
			dst:      &fakeServerStream{sendHeaderErr: status.Error(codes.Internal, "send header failed")},
			src:      &fakeClientStream{},
			wantCode: codes.Internal,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := forwardClientToServer(tc.dst, tc.src)
			if gotCode := status.Code(err); gotCode != tc.wantCode {
				t.Errorf("forwardClientToServer() code = %v, want %v (err: %v)", gotCode, tc.wantCode, err)
			}
			if err == nil || !strings.Contains(err.Error(), "proxy_error:") {
				t.Errorf("forwardClientToServer() err = %v, want substring %q", err, "proxy_error:")
			}
		})
	}
}

type closedConnDialer struct{}

func (d *closedConnDialer) DialGRPC(ctx context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return nil, err
	}
	conn.Close()
	return conn, nil
}

func TestStreamHandlerNewClientStreamError(t *testing.T) {
	p := &proxy{
		destProvider: func(_ metadata.MD, _ string) (string, metadata.MD, error) {
			return "localhost:0", nil, nil
		},
		grpcDialer: &closedConnDialer{},
	}
	ctx := metadata.NewIncomingContext(context.Background(), metadata.MD{})
	err := p.streamHandler(ctx, "/gnmi.gNMI/Get", &fakeServerStream{})
	if gotCode := status.Code(err); gotCode != codes.Canceled {
		t.Errorf("streamHandler() code = %v, want %v (err: %v)", gotCode, codes.Canceled, err)
	}
	if err == nil || !strings.Contains(err.Error(), "proxy_error:") {
		t.Errorf("streamHandler() err = %v, want substring %q", err, "proxy_error:")
	}
}
