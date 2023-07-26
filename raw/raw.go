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

// Package raw provides raw client APIs.
package raw

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/internal/events"
	"github.com/openconfig/ondatra/internal/rawapis"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

// NewDUTAPIs returns a new instance of raw DUT APIs.
// Tests must not call this directly.
func NewDUTAPIs(dut binding.DUT) *DUTAPIs {
	return &DUTAPIs{dut}
}

// DUTAPIs provides access to raw DUT protocol APIs.
type DUTAPIs struct {
	dut binding.DUT
}

// GNMI provides access to creating raw gNMI clients for the dut.
func (r *DUTAPIs) GNMI() *GNMIAPI {
	return &GNMIAPI{r.dut}
}

// TODO(greg-dennis): Should DialGNMI be in the binding.Device interface.
type gnmiDevice interface {
	binding.Device
	rawapis.GNMIDialer
}

// GNMIAPI provides access for creating a default or new gNMI client on the DUT.
type GNMIAPI struct {
	dev gnmiDevice
}

// New returns a new gNMI client for the dut.
func (g *GNMIAPI) New(t testing.TB) gpb.GNMIClient {
	t.Helper()
	t = events.ActionStarted(t, "Creating gNMI client for %s", g.dev)
	gnmi, err := rawapis.NewGNMI(context.Background(), g.dev)
	if err != nil {
		t.Fatalf("Failed to create gNMI client for %v: %v", g.dev, err)
	}
	return gnmi
}

// Default returns the default gNMI client for the DUT.
func (g *GNMIAPI) Default(t testing.TB) gpb.GNMIClient {
	t.Helper()
	t = events.ActionStarted(t, "Fetching gNMI client for %s", g.dev)
	gnmi, err := rawapis.FetchGNMI(context.Background(), g.dev)
	if err != nil {
		t.Fatalf("Failed to fetch gNMI client for %v: %v", g.dev, err)
	}
	return gnmi
}

// GNOI provides access to creating raw gNOI clients for the dut.
func (r *DUTAPIs) GNOI() *GNOIAPI {
	return &GNOIAPI{r.dut}
}

// GNOIAPI provides access to creating raw gNOI clients for the dut.
type GNOIAPI struct {
	dut binding.DUT
}

// GNOI stores APIs to GNOI services.
type GNOI interface {
	// Embed an unexported interface that wraps binding.GNOIClients,
	// so as to not expose the binding.GNOIClients instance directly.
	privateGNOI
}

type privateGNOI interface {
	binding.GNOIClients
}

// New returns a new gNOI client for the dut.
func (g *GNOIAPI) New(t testing.TB) GNOI {
	t.Helper()
	t = events.ActionStarted(t, "Creating gNOI client for %s", g.dut)
	bgnoi, err := rawapis.NewGNOI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("Failed to create gNOI client for %v: %v", g.dut, err)
	}
	return bgnoi
}

// Default returns the default gNOI client for the dut.
func (g *GNOIAPI) Default(t testing.TB) GNOI {
	t.Helper()
	t = events.ActionStarted(t, "Fetching gNOI client for %s", g.dut)
	bgnoi, err := rawapis.FetchGNOI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("Failed to fetch gNOI client for %v: %v", g.dut, err)
	}
	return bgnoi
}

// GNSIAPI provides access to creating raw gNSI client for the DUT.
type GNSIAPI struct {
	dut binding.DUT
}

// GNSI stores APIs to GNSI services.
type GNSI interface {
	// Embed an unexported interface that wraps binding.GNSIClients,
	// so as to not expose the binding.GNSIClients instance directly.
	privateGNSI
}

type privateGNSI interface {
	binding.GNSIClients
}

// GNSI provides access to creating raw gNSI clients for the dut.
func (r *DUTAPIs) GNSI() *GNSIAPI {
	return &GNSIAPI{r.dut}
}

// New returns a new gNSI client for the DUT.
func (g *GNSIAPI) New(t testing.TB) GNSI {
	t.Helper()
	t = events.ActionStarted(t, "Creating gNSI  client for %s", g.dut)
	bgnsi, err := rawapis.NewGNSI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("Failed to create gNSI client for %v: %v", g.dut, err)
	}
	return bgnsi
}

// Default returns the default gNSI client for the dut.
func (g *GNSIAPI) Default(t testing.TB) GNSI {
	t.Helper()
	t = events.ActionStarted(t, "Fetching gNSI client for %s", g.dut)
	bgnsi, err := rawapis.FetchGNSI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("Failed to fetch gNSI client for %v: %v", g.dut, err)
	}
	return bgnsi
}

// GRIBI provides access to createing raw gRIBI clients for the dut.
func (r *DUTAPIs) GRIBI() *GRIBIAPI {
	return &GRIBIAPI{r.dut}
}

// GRIBIAPI provides access to creating raw gRIBI clients for the DUT.
type GRIBIAPI struct {
	dut binding.DUT
}

// New returns a new gRIBI client for the dut.
func (g *GRIBIAPI) New(t testing.TB) grpb.GRIBIClient {
	t.Helper()
	t = events.ActionStarted(t, "Creating gRIBI client for %s", g.dut)
	grc, err := rawapis.NewGRIBI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("Failed to create gRIBI client for %v: %v", g.dut, err)
	}
	return grc
}

// Default returns the default gRIBI client for the dut.
func (g *GRIBIAPI) Default(t testing.TB) grpb.GRIBIClient {
	t.Helper()
	t = events.ActionStarted(t, "Fetching gRIBI client for %s", g.dut)
	grc, err := rawapis.FetchGRIBI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("Failed to fetch gRIBI client for %v: %v", g.dut, err)
	}
	return grc
}

// P4RT provides access to creating raw P4RT clients for the dut.
func (r *DUTAPIs) P4RT() *P4RTAPI {
	return &P4RTAPI{r.dut}
}

// P4RTAPI provides access for creating a default or new GRIBI client on the DUT.
type P4RTAPI struct {
	dut binding.DUT
}

// New returns a P4RT client on the DUT.
func (p *P4RTAPI) New(t testing.TB) p4pb.P4RuntimeClient {
	t.Helper()
	t = events.ActionStarted(t, "Creating P4RT client for %s", p.dut)
	p4rtClient, err := rawapis.NewP4RT(context.Background(), p.dut)
	if err != nil {
		t.Fatalf("Failed to create P4RT client for %v: %v", p.dut, err)
	}
	return p4rtClient
}

// Default returns the default P4RT client on the DUT.
func (p *P4RTAPI) Default(t testing.TB) p4pb.P4RuntimeClient {
	t.Helper()
	t = events.ActionStarted(t, "Fetching P4RT client for %s", p.dut)
	p4rtClient, err := rawapis.FetchP4RT(context.Background(), p.dut)
	if err != nil {
		t.Fatalf("Failed to fetch P4RT client for %v: %v", p.dut, err)
	}
	return p4rtClient
}

// StreamClient provides the interface for streaming IO to DUT.
type StreamClient interface {
	// Embed an unexported interface that wraps binding.StreamClient,
	// so as to not expose the binding.StreamClient instance directly.
	privateStreamClient
}

type privateStreamClient interface {
	binding.StreamClient
}

// CLI returns a new streaming CLI client for the DUT.
func (r *DUTAPIs) CLI(t testing.TB) StreamClient {
	t.Helper()
	t = events.ActionStarted(t, "Creating CLI client for %s", r.dut)
	c, err := rawapis.NewCLI(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to create CLI client for %v: %v", r.dut, err)
	}
	return c
}

// Console returns a new Console client for the DUT.
func (r *DUTAPIs) Console(t testing.TB) StreamClient {
	t.Helper()
	t = events.ActionStarted(t, "Creating console client for %s", r.dut)
	c, err := rawapis.NewConsole(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to create console client for %v: %v", r.dut, err)
	}
	return c
}

// NewATEAPIs returns a new instance of raw ATE APIs.
// Tests must not call this directly.
func NewATEAPIs(ate binding.ATE) *ATEAPIs {
	return &ATEAPIs{ate}
}

// ATEAPIs provides access to raw DUT protocol APIs.
type ATEAPIs struct {
	ate binding.ATE
}

// GNMI provides access to creating raw gNMI clients for the dut.
func (r *ATEAPIs) GNMI() *GNMIAPI {
	return &GNMIAPI{r.ate}
}

// IxNetwork returns the raw IxNetwork session for the ATE.
// TODO(team): Add unit tests once raw APIs is factored out into its own package.
func (r *ATEAPIs) IxNetwork(t testing.TB) *ixweb.Session {
	t.Helper()
	t = events.ActionStarted(t, "Fetching IxNetwork session for %s", r.ate)
	ixnet, err := rawapis.FetchIxNetwork(context.Background(), r.ate)
	if err != nil {
		t.Fatalf("IxNetwork(t) on %v: %v", r.ate, err)
	}
	return ixnet.Session
}
