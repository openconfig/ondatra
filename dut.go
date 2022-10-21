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

package ondatra

import (
	"golang.org/x/net/context"
	"testing"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/config"
	"github.com/openconfig/ondatra/config/device"
	"github.com/openconfig/ondatra/internal/cli"
	"github.com/openconfig/ondatra/internal/console"
	"github.com/openconfig/ondatra/internal/debugger"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ondatra/internal/gribi"
	"github.com/openconfig/ondatra/internal/operations"
	"github.com/openconfig/ondatra/internal/p4rt"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"

)

// DUTDevice is a device under test.
type DUTDevice struct {
	*Device
}

// Config returns a handle to the DUT configuration API.
func (d *DUTDevice) Config() *Config {
	root := device.DeviceRoot(d.Name())
	if d.Vendor() == CISCO || d.Vendor() == JUNIPER {
		genutil.PutUseGetForConfig(root, true)
	}
	root.PutCustomData(genutil.DefaultClientKey, d.Device.clientFn)
	return &Config{
		dut:        d.res.(binding.DUT),
		DevicePath: root,
	}
}

// Config is the DUT configuration API.
type Config struct {
	dut binding.DUT
	*device.DevicePath
}

// New returns an empty DUT configuration.
func (c *Config) New() *config.VendorConfig {
	return config.NewVendorConfig(c.dut)
}

// Operations returns a handle to the DUT operations API.
func (d *DUTDevice) Operations() *Operations {
	return &Operations{dut: d.res.(binding.DUT)}
}

// RawAPIs returns a handle to raw protocol APIs on the DUT.
func (d *DUTDevice) RawAPIs() *RawDUTAPIs {
	return &RawDUTAPIs{dut: d.res.(binding.DUT)}
}

// RawDUTAPIs provides access to raw DUT protocol APIs.
type RawDUTAPIs struct {
	dut binding.DUT
}

// GNMI provides access to either a new or default gNMI client.
func (r *RawDUTAPIs) GNMI() *GNMIAPI {
	return &GNMIAPI{dut: r.dut}
}

// GNOI provides access to either a new or default gNOI client.
func (r *RawDUTAPIs) GNOI() *GNOIAPI {
	return &GNOIAPI{dut: r.dut}
}

// GRIBI provides access to either a new or default GRIBI client.
func (r *RawDUTAPIs) GRIBI() *GRIBIAPI {
	return &GRIBIAPI{dut: r.dut}
}

// GNMIAPI provides access for creating a default or new gNMI client on the DUT.
type GNMIAPI struct {
	dut binding.DUT
}

// GNOIAPI provides access for creating a default or new gNOI client on the DUT.
type GNOIAPI struct {
	dut binding.DUT
}

// GRIBIAPI provides access for creating a default or new GRIBI client on the DUT.
type GRIBIAPI struct {
	dut binding.DUT
}

// New returns a new gNMI client on the DUT. This client will not be cached.
func (g *GNMIAPI) New(t testing.TB) gpb.GNMIClient {
	t.Helper()
	debugger.ActionStarted(t, "Creating gNMI client for %s", g.dut)
	gnmi, err := newGNMI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("GNMI(t) on %v: %v", g.dut, err)
	}
	return gnmi
}

// Default returns the default gNMI client for the DUT.
func (g *GNMIAPI) Default(t testing.TB) gpb.GNMIClient {
	t.Helper()
	debugger.ActionStarted(t, "Fetching gNMI client for %s", g.dut)
	gnmi, err := fetchGNMI(context.Background(), g.dut, nil)
	if err != nil {
		t.Fatalf("GNMI(t) on %v: %v", g.dut, err)
	}
	return gnmi
}

// New returns a new gNOI client on the DUT.
func (g *GNOIAPI) New(t testing.TB) GNOI {
	t.Helper()
	debugger.ActionStarted(t, "Creating gNOI client for %s", g.dut)
	bgnoi, err := operations.NewGNOI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("GNOI(t) on %v: %v", g.dut, err)
	}
	return bgnoi
}

// Default returns the default gNOI client for the DUT.
func (g *GNOIAPI) Default(t testing.TB) GNOI {
	t.Helper()
	debugger.ActionStarted(t, "Fetching gNOI client for %s", g.dut)
	bgnoi, err := operations.FetchGNOI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("GNOI(t) on %v: %v", g.dut, err)
	}
	return bgnoi
}

// New returns a new gRIBI client on the DUT.
func (g *GRIBIAPI) New(t testing.TB) grpb.GRIBIClient {
	t.Helper()
	debugger.ActionStarted(t, "Creating gRIBI client for %s", g.dut)
	grc, err := gribi.NewGRIBI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("GRIBI(t) on %v: %v", g.dut, err)
	}
	return grc
}

// Default returns the default gRIBI client for the DUT.
func (g *GRIBIAPI) Default(t testing.TB) grpb.GRIBIClient {
	t.Helper()
	debugger.ActionStarted(t, "Fetching gRIBI client for %s", g.dut)
	grc, err := gribi.FetchGRIBI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("GRIBI(t) on %v: %v", g.dut, err)
	}
	return grc
}

// GNOI stores gNOI clients to a DUT.
type GNOI interface {
	// Embed an unexported interface that wraps binding.GNOIClients,
	// so as to not expose the binding.GNOIClients instance directly.
	privateGNOI
}

type privateGNOI interface {
	binding.GNOIClients
}

// P4RT returns a P4RT client on the DUT.
func (r *RawDUTAPIs) P4RT(t testing.TB) p4pb.P4RuntimeClient {
	t.Helper()
	debugger.ActionStarted(t, "Creating P4RT client for %s", r.dut)
	p4rtClient, err := p4rt.NewP4RT(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to create P4RT client on %v: %v", r.dut, err)
	}
	return p4rtClient
}

// StreamClient provides the interface for streaming IO to DUT.
type StreamClient interface {
	// Embed an unexported interface that wraps binding.GNOIClients,
	// so as to not expose the binding.GNOIClients instance directly.
	privateStreamClient
}

type privateStreamClient interface {
	binding.StreamClient
}

// CLI returns a streaming CLI client on the DUT.
func (r *RawDUTAPIs) CLI(t testing.TB) StreamClient {
	t.Helper()
	debugger.ActionStarted(t, "Creating CLI client for %s", r.dut)
	c, err := cli.NewCLI(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to create CLI client on %v: %v", r.dut, err)
	}
	return c
}

// Console returns a transactional CLI client on the DUT.
func (r *RawDUTAPIs) Console(t testing.TB) StreamClient {
	t.Helper()
	debugger.ActionStarted(t, "Creating console client for %s", r.dut)
	c, err := console.NewConsole(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to create console client on %v: %v", r.dut, err)
	}
	return c
}
