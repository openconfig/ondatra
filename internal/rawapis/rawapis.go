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

// Package rawapis centralizes the management of Ondatra raw APIs.
//
// This module provides "New" and "Fetch" functions for raw client APIs, which
// return new and a cached API clients, respectively. For every DUT gRPC API,
// this module provides both "New" and "Fetch" functions. For every StreamClient
// API, only a "New" function is provided, since those clients are intended to
// be created and closed on each individual use. For every ATE API, only "Fetch"
// functions are provided, since there is no use case for multiple clients.
package rawapis

import (
	"fmt"
	"sync"

	"golang.org/x/net/context"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"google.golang.org/grpc"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

// CommonDialOpts to include in all gRPC dial calls.
// TODO(greg-dennis): Unexport once IxNetwork is removed.
var CommonDialOpts = []grpc.DialOption{
	grpc.WithBlock(),
	withUnaryAnnotateErrors(),
	withStreamAnnotateErrors(),
}

// NewCLI creates a CLI client for the specified DUT.
func NewCLI(ctx context.Context, dut binding.DUT) (binding.StreamClient, error) {
	return dut.DialCLI(ctx)
}

// NewConsole creates a console client for the specified DUT.
func NewConsole(ctx context.Context, dut binding.DUT) (binding.StreamClient, error) {
	return dut.DialConsole(ctx)
}

var (
	gnmiMu sync.Mutex
	gnmis  = make(map[GNMIDialer]gpb.GNMIClient)
)

// GNMIDialer is an interface for devices that can dial gNMI.
type GNMIDialer interface {
	DialGNMI(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error)
}

// NewGNMI creates a new gNMI client for the specified DUT.
func NewGNMI(ctx context.Context, dialer GNMIDialer) (gpb.GNMIClient, error) {
	return dialer.DialGNMI(ctx, CommonDialOpts...)
}

// FetchGNMI fetches the cached gNMI client for the specified DUT.
func FetchGNMI(ctx context.Context, dut GNMIDialer) (gpb.GNMIClient, error) {
	gnmiMu.Lock()
	defer gnmiMu.Unlock()
	c, ok := gnmis[dut]
	if !ok {
		var err error
		c, err = NewGNMI(ctx, dut)
		if err != nil {
			return nil, fmt.Errorf("error dialing gNMI: %w", err)
		}
		gnmis[dut] = c
	}
	return c, nil
}

var (
	gnoiMu sync.Mutex
	gnois  = make(map[binding.DUT]binding.GNOIClients)
)

// NewGNOI creates a gNOI client for the specified DUT.
func NewGNOI(ctx context.Context, dut binding.DUT) (binding.GNOIClients, error) {
	return dut.DialGNOI(ctx, CommonDialOpts...)
}

// FetchGNOI fetches the cached gNOI client for the specified DUT.
func FetchGNOI(ctx context.Context, dut binding.DUT) (binding.GNOIClients, error) {
	gnoiMu.Lock()
	defer gnoiMu.Unlock()
	c, ok := gnois[dut]
	if !ok {
		var err error
		c, err = NewGNOI(ctx, dut)
		if err != nil {
			return nil, fmt.Errorf("error dialing gNOI: %w", err)
		}
		gnois[dut] = c
	}
	return c, nil
}

var (
	gnsiMu sync.Mutex
	gnsis  = make(map[binding.DUT]binding.GNSIClients)
)

// NewGNSI creates a gNSI client for the specified DUT.
func NewGNSI(ctx context.Context, dut binding.DUT) (binding.GNSIClients, error) {
	return dut.DialGNSI(ctx, CommonDialOpts...)
}

// FetchGNSI fetches the cached gNSI client for the specified DUT.
func FetchGNSI(ctx context.Context, dut binding.DUT) (binding.GNSIClients, error) {
	gnsiMu.Lock()
	defer gnsiMu.Unlock()
	c, ok := gnsis[dut]
	if !ok {
		var err error
		c, err = NewGNSI(ctx, dut)
		if err != nil {
			return nil, fmt.Errorf("error dialing gNSI: %w", err)
		}
		gnsis[dut] = c
	}
	return c, nil
}

var (
	gribiMu sync.Mutex
	gribis  = make(map[binding.Device]grpb.GRIBIClient)
)

// NewGRIBI creates a new gRIBI client for the specified DUT.
func NewGRIBI(ctx context.Context, dut binding.DUT) (grpb.GRIBIClient, error) {
	return dut.DialGRIBI(ctx, CommonDialOpts...)
}

// FetchGRIBI fetches the cached gRIBI client for the specified DUT.
func FetchGRIBI(ctx context.Context, dut binding.DUT) (grpb.GRIBIClient, error) {
	gribiMu.Lock()
	defer gribiMu.Unlock()
	c, ok := gribis[dut]
	if !ok {
		var err error
		c, err = NewGRIBI(ctx, dut)
		if err != nil {
			return nil, fmt.Errorf("error dialing gRIBI: %w", err)
		}
		gribis[dut] = c
	}
	return c, nil
}

var (
	p4rtMu sync.Mutex
	p4rts  = make(map[binding.DUT]p4pb.P4RuntimeClient)
)

// NewP4RT creates a new P4RT client for the specified DUT.
func NewP4RT(ctx context.Context, dut binding.DUT) (p4pb.P4RuntimeClient, error) {
	return dut.DialP4RT(ctx, CommonDialOpts...)
}

// FetchP4RT fetches the cached P4RT client for the specified DUT.
func FetchP4RT(ctx context.Context, dut binding.DUT) (p4pb.P4RuntimeClient, error) {
	p4rtMu.Lock()
	defer p4rtMu.Unlock()
	c, ok := p4rts[dut]
	if !ok {
		var err error
		c, err = NewP4RT(ctx, dut)
		if err != nil {
			return nil, fmt.Errorf("error dialing P4RT: %w", err)
		}
		p4rts[dut] = c
	}
	return c, nil
}

var (
	ixnetMu sync.Mutex
	ixnets  = make(map[binding.ATE]*binding.IxNetwork)
)

// FetchIxNetwork returns the cached IxNetwork client for the specified ATE.
func FetchIxNetwork(ctx context.Context, ate binding.ATE) (*binding.IxNetwork, error) {
	ixnetMu.Lock()
	defer ixnetMu.Unlock()
	c, ok := ixnets[ate]
	if !ok {
		var err error
		c, err = ate.DialIxNetwork(ctx)
		if err != nil {
			return nil, fmt.Errorf("error dialing IxNetwork: %w", err)
		}
		ixnets[ate] = c
	}
	return c, nil
}

var (
	otgMu sync.Mutex
	otgs  = make(map[binding.ATE]gosnappi.GosnappiApi)
)

// FetchOTG fetches the cached OTG client for the specified ATE.
func FetchOTG(ctx context.Context, ate binding.ATE) (gosnappi.GosnappiApi, error) {
	otgMu.Lock()
	defer otgMu.Unlock()
	c, ok := otgs[ate]
	if !ok {
		var err error
		c, err = ate.DialOTG(ctx, CommonDialOpts...)
		if err != nil {
			return nil, fmt.Errorf("error dialing OTG: %w", err)
		}
		otgs[ate] = c
	}
	return c, nil
}

var (
	otgGNMIMu sync.Mutex
	otgGNMIs  = make(map[binding.ATE]gpb.GNMIClient)
)

// FetchOTGGNMI fetches the cached OTG GNMI client for the specified ATE.
func FetchOTGGNMI(ctx context.Context, ate binding.ATE) (gpb.GNMIClient, error) {
	otgGNMIMu.Lock()
	defer otgGNMIMu.Unlock()
	c, ok := otgGNMIs[ate]
	if !ok {
		var err error
		c, err = ate.DialGNMI(ctx, CommonDialOpts...)
		if err != nil {
			return nil, fmt.Errorf("error dialing OTG GNMI: %w", err)
		}
		otgGNMIs[ate] = c
	}
	return c, nil
}
