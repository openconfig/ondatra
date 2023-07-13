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
	"github.com/openconfig/ondatra/internal/testbed"
	"google.golang.org/grpc"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

// Setup initializes Ondatra with a new fake binding, initializes the testbed
// to an unreserved state, and returns the fake binding for further stubbing.
func Setup() *Binding {
	bind := new(Binding)
	testbed.SetBinding(bind)
	return bind.WithReservation(nil)
}

var _ binding.Binding = (*Binding)(nil)

// Binding is a fake binding.Binding implementation comprised of stubs.
type Binding struct {
	ReserveFn          func(context.Context, *opb.Testbed, time.Duration, time.Duration, map[string]string) (*binding.Reservation, error)
	ReleaseFn          func(context.Context) error
	FetchReservationFn func(context.Context, string) (*binding.Reservation, error)
}

// WithReservation sets Ondatra to a state in which the specified reservation
// has been reserved. It does not alter the implementation of this binding's
// stub functions in any way. If a nil reservation is supplied, Ondatra is set
// to an unreserved state.
func (b *Binding) WithReservation(res *binding.Reservation) *Binding {
	testbed.SetReservationForTesting(res)
	return b
}

// Reserve delegates to b.ReserveFn.
func (b *Binding) Reserve(ctx context.Context, tb *opb.Testbed, runTime, waitTime time.Duration, partial map[string]string) (*binding.Reservation, error) {
	if b.ReserveFn == nil {
		log.Fatal("fakebind Reserve called but ReserveFn not set")
	}
	return b.ReserveFn(ctx, tb, runTime, waitTime, partial)
}

// Release delegates to b.ReleaseFN.
func (b *Binding) Release(ctx context.Context) error {
	if b.ReleaseFn == nil {
		log.Fatal("fakebind Release called but ReleaseFn not set")
	}
	return b.ReleaseFn(ctx)
}

// FetchReservation delegates to b.FetchReservationFn.
func (b *Binding) FetchReservation(ctx context.Context, id string) (*binding.Reservation, error) {
	if b.FetchReservationFn == nil {
		log.Fatal("fakebind FetchReservation called but FetchReservationFn not set")
	}
	return b.FetchReservationFn(ctx, id)
}

var _ binding.DUT = (*DUT)(nil)

// DUT is a fake implementation of binding.DUT comprised of stubs.
type DUT struct {
	*binding.AbstractDUT
	PushConfigFn  func(context.Context, string, bool) error
	DialCLIFn     func(context.Context) (binding.StreamClient, error)
	DialConsoleFn func(context.Context) (binding.StreamClient, error)
	DialGNMIFn    func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error)
	DialGNOIFn    func(context.Context, ...grpc.DialOption) (binding.GNOIClients, error)
	DialGNSIFn    func(context.Context, ...grpc.DialOption) (binding.GNSIClients, error)
	DialGRIBIFn   func(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error)
	DialP4RTFn    func(context.Context, ...grpc.DialOption) (p4pb.P4RuntimeClient, error)
}

// PushConfig delegates to d.PushConfigFn.
func (d *DUT) PushConfig(ctx context.Context, config string, reset bool) error {
	if d.PushConfigFn == nil {
		log.Fatal("fakebind PushConfig called but PushConfigFn not set")
	}
	return d.PushConfigFn(ctx, config, reset)
}

// DialCLI delegates to d.DialCLIFn.
func (d *DUT) DialCLI(ctx context.Context) (binding.StreamClient, error) {
	if d.DialCLIFn == nil {
		log.Fatal("fakebind DialCLI called but DialCLIFn not set")
	}
	return d.DialCLIFn(ctx)
}

// DialConsole delegates to d.DialConsoleFn.
func (d *DUT) DialConsole(ctx context.Context) (binding.StreamClient, error) {
	if d.DialConsoleFn == nil {
		log.Fatal("fakebind DialConsole called but DialConsoleFn not set")
	}
	return d.DialConsoleFn(ctx)
}

// DialGNMI delegates to d.DialGNMIFn.
func (d *DUT) DialGNMI(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	if d.DialGNMIFn == nil {
		log.Fatal("fakebind DialGNMI called but DialGNMIFn not set")
	}
	return d.DialGNMIFn(ctx, opts...)
}

// DialGNOI delegates to d.DialGNOIFn.
func (d *DUT) DialGNOI(ctx context.Context, opts ...grpc.DialOption) (binding.GNOIClients, error) {
	if d.DialGNOIFn == nil {
		log.Fatal("fakebind DialGNOI called but DialGNOIFn not set")
	}
	return d.DialGNOIFn(ctx, opts...)
}

// DialGNSI delegates to d.DialGNSIFn.
func (d *DUT) DialGNSI(ctx context.Context, opts ...grpc.DialOption) (binding.GNSIClients, error) {
	if d.DialGNSIFn == nil {
		log.Fatal("fakebind DialGNSI called but DialGNSIFn not set")
	}
	return d.DialGNSIFn(ctx, opts...)
}

// DialGRIBI delegates to d.DialGRIBIFn.
func (d *DUT) DialGRIBI(ctx context.Context, opts ...grpc.DialOption) (grpb.GRIBIClient, error) {
	if d.DialGRIBIFn == nil {
		log.Fatal("fakebind DialGRIBI called but DialGRIBIFn not set")
	}
	return d.DialGRIBIFn(ctx, opts...)
}

// DialP4RT delegates to d.DialP4RTFn.
func (d *DUT) DialP4RT(ctx context.Context, opts ...grpc.DialOption) (p4pb.P4RuntimeClient, error) {
	if d.DialP4RTFn == nil {
		log.Fatal("fakebind DialP4RT called but DialP4RTFn not set")
	}
	return d.DialP4RTFn(ctx, opts...)
}

var _ binding.ATE = (*ATE)(nil)

// ATE is a fake implementation of binding.ATE comprised of stubs.
type ATE struct {
	*binding.AbstractATE
	DialIxNetworkFn func(context.Context) (*binding.IxNetwork, error)
	DialGNMIFn      func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error)
	DialOTGFn       func(context.Context, ...grpc.DialOption) (gosnappi.GosnappiApi, error)
}

// DialIxNetwork delegates to a.DialIxNetworkFn.
func (a *ATE) DialIxNetwork(ctx context.Context) (*binding.IxNetwork, error) {
	if a.DialIxNetworkFn == nil {
		log.Fatal("fakebind DialIxNetwork called but DialIxNetworkFn not set")
	}
	return a.DialIxNetworkFn(ctx)
}

// DialGNMI delegates to a.DialGNMIFn.
func (a *ATE) DialGNMI(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	if a.DialGNMIFn == nil {
		log.Fatal("fakebind DialGNMI called but DialGNMIFn not set")
	}
	return a.DialGNMIFn(ctx, opts...)
}

// DialOTG delegates to a.DialOTGFn.
func (a *ATE) DialOTG(ctx context.Context, opts ...grpc.DialOption) (gosnappi.GosnappiApi, error) {
	if a.DialOTGFn == nil {
		log.Fatal("fakebind DialOTG called but DialOTGFn not set")
	}
	return a.DialOTGFn(ctx, opts...)
}
