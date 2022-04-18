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

// Package fakebind implements a fake testbed binding, backed by a fake SSH server.
package fakebind

import (
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"google.golang.org/grpc"

	gnmiclient "github.com/openconfig/gnmi/client"
	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

var _ binding.Binding = &Binding{}

// Binding is a fake testbed binding comprised of stub implementations.
type Binding struct {
	Reservation     *binding.Reservation
	ResvFetcher     func(context.Context, string) (*binding.Reservation, error)
	ConfigPusher    func(context.Context, *binding.DUT, string, bool) error
	CLIDialer       func(context.Context, *binding.DUT, ...grpc.DialOption) (binding.StreamClient, error)
	ConsoleDialer   func(context.Context, *binding.DUT, ...grpc.DialOption) (binding.StreamClient, error)
	GNMIDialer      func(context.Context, *binding.DUT, ...grpc.DialOption) (gpb.GNMIClient, error)
	GNOIDialer      func(context.Context, *binding.DUT, ...grpc.DialOption) (binding.GNOIClients, error)
	GRIBIDialer     func(context.Context, *binding.DUT, ...grpc.DialOption) (grpb.GRIBIClient, error)
	P4RTDialer      func(context.Context, *binding.DUT, ...grpc.DialOption) (p4pb.P4RuntimeClient, error)
	IxNetworkDialer func(context.Context, *binding.ATE) (*binding.IxNetwork, error)
	OTGDialer       func(context.Context, *binding.ATE) (gosnappi.GosnappiApi, error)
	OTGGNMIDialer   func(*binding.ATE) (*gnmiclient.Query, error)
}

// Reset zeros out all the stub implementations.
func (b *Binding) Reset() {
	b.Reservation = nil
	b.ResvFetcher = nil
	b.ConfigPusher = nil
	b.CLIDialer = nil
	b.ConsoleDialer = nil
	b.GNMIDialer = nil
	b.GNOIDialer = nil
	b.P4RTDialer = nil
	b.IxNetworkDialer = nil
	b.OTGDialer = nil
	b.OTGGNMIDialer = nil
}

// Reserve returns b.Reservation.
func (b *Binding) Reserve(context.Context, *opb.Testbed, time.Duration, time.Duration, map[string]string) (*binding.Reservation, error) {
	return b.Reservation, nil
}

// FetchReservation delegates to b.ResvFetcher.
func (b *Binding) FetchReservation(ctx context.Context, id string) (*binding.Reservation, error) {
	return b.ResvFetcher(ctx, id)
}

// Release is a noop.
func (b *Binding) Release(context.Context) (rerr error) {
	return nil
}

// DialATEGNMI is a noop.
func (b *Binding) DialATEGNMI(ctx context.Context, ate *binding.ATE, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	return nil, nil
}

// PushConfig delegates to b.ConfigPusher.
func (b *Binding) PushConfig(ctx context.Context, dut *binding.DUT, config string, reset bool) error {
	return b.ConfigPusher(ctx, dut, config, reset)
}

// DialGNMI creates a client connection to the fake GNMI server.
func (b *Binding) DialGNMI(ctx context.Context, dut *binding.DUT, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	return b.GNMIDialer(ctx, dut, opts...)
}

// DialGNOI creates a client connection to the fake GNOI server.
func (b *Binding) DialGNOI(ctx context.Context, dut *binding.DUT, opts ...grpc.DialOption) (binding.GNOIClients, error) {
	return b.GNOIDialer(ctx, dut, opts...)
}

// DialGRIBI creates a client connection to the fake GRIBI server.
func (b *Binding) DialGRIBI(ctx context.Context, dut *binding.DUT, opts ...grpc.DialOption) (grpb.GRIBIClient, error) {
	return b.GRIBIDialer(ctx, dut, opts...)
}

// DialP4RT creates a client connection to the fake P4RT server.
func (b *Binding) DialP4RT(ctx context.Context, dut *binding.DUT, opts ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
	return b.P4RTDialer(ctx, dut, opts...)
}

// DialCLI creates a client connection to the fake CLI server.
func (b *Binding) DialCLI(ctx context.Context, dut *binding.DUT, opts ...grpc.DialOption) (binding.StreamClient, error) {
	return b.CLIDialer(ctx, dut, opts...)
}

// DialConsole creates a client connection to the fake Console server.
func (b *Binding) DialConsole(ctx context.Context, dut *binding.DUT, opts ...grpc.DialOption) (binding.StreamClient, error) {
	return b.ConsoleDialer(ctx, dut, opts...)
}

// DialIxNetwork delegates to b.IxNetworkDialer.
func (b *Binding) DialIxNetwork(ctx context.Context, ate *binding.ATE) (*binding.IxNetwork, error) {
	return b.IxNetworkDialer(ctx, ate)
}

// DialOTG delegates to b.OTGDialer.
func (b *Binding) DialOTG(ctx context.Context, ate *binding.ATE) (gosnappi.GosnappiApi, error) {
	return b.OTGDialer(ctx, ate)
}

// DialOTGGNMI delegates to b.OTGGNMIDialer.
func (b *Binding) DialOTGGNMI(ate *binding.ATE) (*gnmiclient.Query, error) {
	return b.OTGGNMIDialer(ate)
}

// HandleInfraFail logs the error and returns it unchanged.
func (b *Binding) HandleInfraFail(err error) error {
	log.Errorf("Infrastructure failure: %v", err)
	return err
}

// SetTestMetadata is a noop.
func (b *Binding) SetTestMetadata(md *binding.TestMetadata) error {
	return nil
}
