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
	"golang.org/x/net/context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials/local"
	"google.golang.org/grpc"
	"github.com/openconfig/gnmi/testing/fake/gnmi"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	fpb "github.com/openconfig/gnmi/testing/fake/proto"
)

// FakeGNMI is a running fake GNMI server.
type FakeGNMI struct {
	agent *gnmi.Agent
	stub  *Stubber
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
	return &FakeGNMI{
		agent: agent,
		stub:  &Stubber{gen: gen},
	}, nil
}

// Dial dials the fake gNMI client and returns a gNMI client stub.
func (g *FakeGNMI) Dial(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	opts = append(opts, grpc.WithTransportCredentials(local.NewCredentials()))
	conn, err := grpc.DialContext(ctx, g.agent.Address(), opts...)
	if err != nil {
		return nil, errors.Wrapf(err, "DialContext(%s, %v)", g.agent.Address(), opts)
	}
	return gpb.NewGNMIClient(conn), nil
}

// Stub reset the stubbed responses to empty and returns a handle to add new ones.
func (g *FakeGNMI) Stub() *Stubber {
	g.stub.gen.Reset()
	return g.stub
}

// Requests returns the set of SubscribeRequests sent to the gNMI server.
func (g *FakeGNMI) Requests() []*gpb.SubscribeRequest {
	return g.agent.Requests()
}

// Stubber is a handle to add stubbed responses.
type Stubber struct {
	gen *fpb.FixedGenerator
}

// Notification appends the given notification as a stub response.
func (s *Stubber) Notification(n *gpb.Notification) *Stubber {
	s.gen.Responses = append(s.gen.Responses, &gpb.SubscribeResponse{
		Response: &gpb.SubscribeResponse_Update{Update: n},
	})
	return s
}

// Sync appends a sync stub response.
func (s *Stubber) Sync() *Stubber {
	s.gen.Responses = append(s.gen.Responses, &gpb.SubscribeResponse{
		Response: &gpb.SubscribeResponse_SyncResponse{},
	})
	return s
}
