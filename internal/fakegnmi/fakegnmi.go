// Copyright 2019 Google LLC
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

// Package fakegnmi implements a fake GNMI server with the ability to stub responses.
package fakegnmi

import (
	"fmt"
	"io"
	"net"
	"strings"

	"golang.org/x/net/context"

	"github.com/openconfig/gnmi/testing/fake/gnmi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/local"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	fpb "github.com/openconfig/gnmi/testing/fake/proto"
)

// FakeGNMI is a running fake GNMI server.
type FakeGNMI struct {
	agent      *gnmi.Agent
	stub       *Stubber
	getWrapper *getWrapper
}

// Start launches a new fake GNMI server on the given port
func Start(port int) (*FakeGNMI, error) {
	gen := &fpb.FixedGenerator{}
	config := &fpb.Config{
		Port:        int32(port),
		Generator:   &fpb.Config_Fixed{Fixed: gen},
		DisableSync: true, // Require every sync to be stubbed explicitly.
		EnableDelay: true, // Respect timestamps if present.
	}
	agent, err := gnmi.New(config, nil)
	if err != nil {
		return nil, err
	}
	stub := &Stubber{gen: gen}
	return &FakeGNMI{
		agent:      agent,
		stub:       stub,
		getWrapper: &getWrapper{stub: stub},
	}, nil
}

// Dial dials the fake gNMI client and returns a gNMI client stub.
func (g *FakeGNMI) Dial(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	opts = append(opts, grpc.WithTransportCredentials(local.NewCredentials()))
	conn, err := grpc.DialContext(ctx, g.agent.Address(), opts...)
	if err != nil {
		return nil, fmt.Errorf("DialContext(%s, %v): %w", g.agent.Address(), opts, err)
	}
	g.getWrapper.GNMIClient = gpb.NewGNMIClient(conn)
	return g.getWrapper, nil
}

// Stub reset the stubbed responses to empty and returns a handle to add new ones.
func (g *FakeGNMI) Stub() *Stubber {
	g.stub.gen.Reset()
	g.stub.getResponses = nil
	return g.stub
}

// Requests returns the set of SubscribeRequests sent to the gNMI server.
func (g *FakeGNMI) Requests() []*gpb.SubscribeRequest {
	return g.agent.Requests()
}

// GetRequests returns the set of GetRequests sent to the gNMI server.
func (g *FakeGNMI) GetRequests() []*gpb.GetRequest {
	return g.getWrapper.getRequests
}

// getWrapper adds gNMI Get functionality to a GNMI client.
type getWrapper struct {
	gpb.GNMIClient
	stub        *Stubber
	getRequests []*gpb.GetRequest
}

// Get is fake implement of gnmi Get.
func (g *getWrapper) Get(ctx context.Context, req *gpb.GetRequest, opts ...grpc.CallOption) (*gpb.GetResponse, error) {
	g.getRequests = append(g.getRequests, req)
	if len(g.stub.getResponses) == 0 {
		return nil, io.EOF
	}
	resp := g.stub.getResponses[0]
	g.stub.getResponses = g.stub.getResponses[1:]
	return resp, nil
}

// Stubber is a handle to add stubbed responses.
type Stubber struct {
	gen          *fpb.FixedGenerator
	getResponses []*gpb.GetResponse
}

// Notification appends the given notification as a stub response.
func (s *Stubber) Notification(n *gpb.Notification) *Stubber {
	s.gen.Responses = append(s.gen.Responses, &gpb.SubscribeResponse{
		Response: &gpb.SubscribeResponse_Update{Update: n},
	})
	return s
}

// GetResponse appends the given GetResponse as a stub response.
func (s *Stubber) GetResponse(gr *gpb.GetResponse) *Stubber {
	s.getResponses = append(s.getResponses, gr)
	return s
}

// Sync appends a sync stub response.
func (s *Stubber) Sync() *Stubber {
	s.gen.Responses = append(s.gen.Responses, &gpb.SubscribeResponse{
		Response: &gpb.SubscribeResponse_SyncResponse{},
	})
	return s
}

// StubGNMI is a running stub GNMI server.
// It doesn't implement correct subscribe mechanics, instead use Stub to insert subscribe responses.
type StubGNMI struct {
	gpb.UnimplementedGNMIServer
	stub        *Stubber
	getRequests []*gpb.GetRequest
	subRequests []*gpb.SubscribeRequest
	lis         net.Listener
	srv         *grpc.Server
}

// StartStubGNMI starts serving a gNMI server at the given port.
func StartStubGNMI(port int) (*StubGNMI, error) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	sg := NewStubGNMI(grpc.NewServer(), l)
	go sg.srv.Serve(l)
	return sg, nil
}

// NewStubGNMI creates a new stub gNMI server.
func NewStubGNMI(srv *grpc.Server, lis net.Listener) *StubGNMI {
	sg := &StubGNMI{lis: lis, srv: srv}
	gpb.RegisterGNMIServer(srv, sg)
	return sg
}

// Addr returns the listener address of the server.
func (g *StubGNMI) Addr() string {
	return strings.Replace(g.lis.Addr().String(), "[::]", "localhost", 1)
}

// Stop gracefully stops the server and closes the listener.
func (g *StubGNMI) Stop() error {
	g.srv.GracefulStop()
	return g.lis.Close()
}

// Stub reset the stubbed responses to empty and returns a handle to add new ones.
func (g *StubGNMI) Stub() *Stubber {
	if g.stub == nil {
		gen := &fpb.FixedGenerator{}
		g.stub = &Stubber{gen: gen}
	}

	g.stub.gen.Reset()
	g.stub.getResponses = nil
	return g.stub
}

// Get is stub implement of gnmi Get.
func (g *StubGNMI) Get(ctx context.Context, req *gpb.GetRequest) (*gpb.GetResponse, error) {
	g.getRequests = append(g.getRequests, req)
	if len(g.stub.getResponses) == 0 {
		return nil, io.EOF
	}
	resp := g.stub.getResponses[0]
	g.stub.getResponses = g.stub.getResponses[1:]
	return resp, nil
}

// Get is stub implement of gnmi Get.
func (g *StubGNMI) Subscribe(stream gpb.GNMI_SubscribeServer) error {
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	g.subRequests = append(g.subRequests, req)
	for _, resp := range g.stub.gen.Responses {
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

// Requests returns the set of SubscribeRequests sent to the gNMI server.
func (g *StubGNMI) Requests() []*gpb.SubscribeRequest {
	return g.subRequests
}

// GetRequests returns the set of GetRequests sent to the gNMI server.
func (g *StubGNMI) GetRequests() []*gpb.GetRequest {
	return g.getRequests
}
