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
package proxy

import (
	"golang.org/x/net/context"
	"net"
	"testing"

	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/lemming"
	rpb "github.com/openconfig/ondatra/proxy/proto/reservation"
)

type fakeBinding struct {
	dialGRPC func(ctx context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error)
	resolve  func() (*rpb.Reservation, error)
}

func (fb *fakeBinding) DialGRPC(ctx context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return fb.dialGRPC(ctx, target, opts...)
}

func (fb *fakeBinding) Resolve() (*rpb.Reservation, error) {
	return fb.resolve()
}

func TestProxy(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("failed to start device listener: %v", err)
	}
	fake, err := lemming.New(lis)
	if err != nil {
		t.Fatalf("failed to start device: %v", err)
	}
	defer fake.Stop()
	fAddr := fake.Addr()
	b := &fakeBinding{}
	b.dialGRPC = grpc.DialContext
	b.resolve = func() (*rpb.Reservation, error) {
		return &rpb.Reservation{
			Id: "fake reservation",
			Devices: map[string]*rpb.ResolvedDevice{
				"dut1": {
					Id:   "dut1",
					Name: "device1",
					Services: map[string]*rpb.Service{
						"grpc": &rpb.Service{
							Id: "grpc",
							Endpoint: &rpb.Service_ProxiedGrpc{
								ProxiedGrpc: &rpb.ProxiedGRPCEndpoint{
									Address: fAddr,
								},
							},
						},
					},
				},
			},
		}, nil
	}
	m, err := New(b)
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}
	defer m.Stop()
	proxies := m.Endpoints()
	conn, err := grpc.DialContext(context.Background(), proxies["dut1"].Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial proxy: %v", err)
	}
	fake.GNMI().GetResponses = []interface{}{
		&gnmipb.GetResponse{
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
	}
	fake.GNMI().Responses = [][]*gnmipb.SubscribeResponse{{{
		Response: &gnmipb.SubscribeResponse_Update{
			Update: &gnmipb.Notification{
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
		},
	}, {
		Response: &gnmipb.SubscribeResponse_SyncResponse{
			SyncResponse: true,
		},
	}}}
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
	for i := 0; i < len(fake.GNMI().Responses[0]); i++ {
		v, err := stream.Recv()
		if err != nil {
			t.Fatalf("Subscribe failed: %v", err)
		}
		if !proto.Equal(v, fake.GNMI().Responses[0][i]) {
			t.Fatalf("invalid message: got %s, want %s", v, fake.GNMI().Responses[0][i])
		}
	}
}
