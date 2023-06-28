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
	"net"
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/internal/fakegnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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
