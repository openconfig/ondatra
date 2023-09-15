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

// BindingDUT returns the underlying binding.DUT.
func (r *DUTAPIs) BindingDUT() binding.DUT {
	return r.dut
}

// GNMI returns the default gNMI client for the dut.
func (r *DUTAPIs) GNMI(t testing.TB) gpb.GNMIClient {
	t.Helper()
	t = events.ActionStarted(t, "Fetching gNMI client for %s", r.dut)
	gnmi, err := rawapis.FetchGNMI(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to fetch gNMI client for %v: %v", r.dut, err)
	}
	return gnmi
}

// GNOI returns the default gNOI clients for the dut.
func (r *DUTAPIs) GNOI(t testing.TB) binding.GNOIClients {
	t.Helper()
	t = events.ActionStarted(t, "Fetching gNOI clients for %s", r.dut)
	bgnoi, err := rawapis.FetchGNOI(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to fetch gNOI clients for %v: %v", r.dut, err)
	}
	return bgnoi
}

// GNSI returns the default gNSI clients for the dut.
func (r *DUTAPIs) GNSI(t testing.TB) binding.GNSIClients {
	t.Helper()
	t = events.ActionStarted(t, "Fetching gNSI clients for %s", r.dut)
	bgnsi, err := rawapis.FetchGNSI(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to fetch gNSI clients for %v: %v", r.dut, err)
	}
	return bgnsi
}

// GRIBI returns the default gRIBI client for the dut.
func (r *DUTAPIs) GRIBI(t testing.TB) grpb.GRIBIClient {
	t.Helper()
	t = events.ActionStarted(t, "Fetching gRIBI client for %s", r.dut)
	grc, err := rawapis.FetchGRIBI(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to fetch gRIBI client for %v: %v", r.dut, err)
	}
	return grc
}

// P4RT returns the default P4RT client for the dut.
func (r *DUTAPIs) P4RT(t testing.TB) p4pb.P4RuntimeClient {
	t.Helper()
	t = events.ActionStarted(t, "Fetching P4RT client for %s", r.dut)
	p4rtClient, err := rawapis.FetchP4RT(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to fetch P4RT client for %v: %v", r.dut, err)
	}
	return p4rtClient
}

// CLI returns a new streaming CLI client for the DUT.
func (r *DUTAPIs) CLI(t testing.TB) binding.CLIClient {
	t.Helper()
	t = events.ActionStarted(t, "Creating CLI client for %s", r.dut)
	c, err := rawapis.NewCLI(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to create CLI client for %v: %v", r.dut, err)
	}
	return c
}

// Console returns a new Console client for the DUT.
func (r *DUTAPIs) Console(t testing.TB) binding.ConsoleClient {
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

// ATEAPIs provides access to raw ATE protocol APIs.
type ATEAPIs struct {
	ate binding.ATE
}

// BindingATE returns the underlying binding.ATE.
func (r *ATEAPIs) BindingATE() binding.ATE {
	return r.ate
}

// GNMI returns the default gNMI client for the ATE.
func (r *ATEAPIs) GNMI(t testing.TB) gpb.GNMIClient {
	t.Helper()
	t = events.ActionStarted(t, "Fetching gNMI client for %s", r.ate)
	gnmi, err := rawapis.FetchGNMI(context.Background(), r.ate)
	if err != nil {
		t.Fatalf("Failed to fetch gNMI client for %v: %v", r.ate, err)
	}
	return gnmi
}
