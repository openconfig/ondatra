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
	"golang.org/x/net/context"
	"time"

	log "github.com/golang/glog"
	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

var _ binding.Binding = &Binding{}

// Binding is a fake testbed binding comprised of stub implementations.
type Binding struct {
	Reservation   *reservation.Reservation
	ConfigPusher  func(context.Context, *reservation.DUT, string, *binding.ConfigOptions) error
	CLIDialer     func(context.Context, *reservation.DUT, ...grpc.DialOption) (binding.StreamClient, error)
	ConsoleDialer func(context.Context, *reservation.DUT, ...grpc.DialOption) (binding.StreamClient, error)
	GNMIDialer    func(context.Context, *reservation.DUT, ...grpc.DialOption) (gpb.GNMIClient, error)
	GNOIDialer    func(context.Context, *reservation.DUT, ...grpc.DialOption) (binding.GNOIClients, error)
	P4RTDialer      func(context.Context, *reservation.DUT, ...grpc.DialOption) (p4pb.P4RuntimeClient, error)
	IxNetworkDialer func(context.Context, *reservation.ATE) (*binding.IxNetwork, error)
}

// Reset zeros out all the stub implementations.
func (b *Binding) Reset() {
	b.Reservation = nil
	b.ConfigPusher = nil
	b.CLIDialer = nil
	b.ConsoleDialer = nil
	b.GNMIDialer = nil
	b.GNOIDialer = nil
	b.P4RTDialer = nil
	b.IxNetworkDialer = nil
}

// Reserve reserves a new fake testbed, reading the definition from the given path.
// If the path is a plain filename, interprets it relative to the target directory.
func (b *Binding) Reserve(_ context.Context, _ *opb.Testbed, _, _ time.Duration) (*reservation.Reservation, error) {
	return b.Reservation, nil
}

// Release is a noop.
func (b *Binding) Release(context.Context) (rerr error) {
	return nil
}

// DialATEGNMI is a noop.
func (b *Binding) DialATEGNMI(ctx context.Context, ate *reservation.ATE, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	return nil, nil
}

// PushConfig delegates to b.ConfigPusher.
func (b *Binding) PushConfig(ctx context.Context, dut *reservation.DUT, config string, opts *binding.ConfigOptions) error {
	return b.ConfigPusher(ctx, dut, config, opts)
}

// DialGNMI creates a client connection to the fake GNMI server.
func (b *Binding) DialGNMI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	return b.GNMIDialer(ctx, dut, opts...)
}

// DialGNOI creates a client connection to the fake GNOI server.
func (b *Binding) DialGNOI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (binding.GNOIClients, error) {
	return b.GNOIDialer(ctx, dut, opts...)
}

// DialP4RT creates a client connection to the fake P4RT server.
func (b *Binding) DialP4RT(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
	return b.P4RTDialer(ctx, dut, opts...)
}

// DialCLI creates a client connection to the fake CLI server.
func (b *Binding) DialCLI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (binding.StreamClient, error) {
	return b.CLIDialer(ctx, dut, opts...)
}

// DialConsole creates a client connection to the fake Console server.
func (b *Binding) DialConsole(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (binding.StreamClient, error) {
	return b.ConsoleDialer(ctx, dut, opts...)
}

// DialIxNetwork delegates to b.IxNetworkDialer.
func (b *Binding) DialIxNetwork(ctx context.Context, ate *reservation.ATE) (*binding.IxNetwork, error) {
	return b.IxNetworkDialer(ctx, ate)
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
