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

// Package raw provides low-level access to the raw device APIs, to be used when
// the other higher-level APIs are not sufficient.
//
// The high-level APIs provided by other Ondatra packages like the [gNMI API]
// are not appropriate for every test scenario. For instance, the gNMI API
// would not be appropriate for low-level testing of a device's conformance to
// the gNMI specification. It would in fact be difficult, if not impossible, to
// use it for some negative test cases, and certainly for fuzz testing, since by
// design the API means to confine tests to constructing valid gNMI requests.
//
// # Accessing the Raw Client APIs
//
// For such test scenarios, tests should use the Ondatra "Raw APIs". Calling
// "dut.RawAPIs()" returns a set of APIs for low-level access to "dut". For
// example, the calls "dut.RawAPIs().GNMI(t)" and "dut.RawAPIs().GNOI(t)"
// return handles to the Go gRPC clients that Ondatra uses for gNMI and gNOI
// interactions, respectively. Similarly, calling "ate.RawAPIs()" returns a set
// of APIs for low-level access to "ate".
//
// The clients returned by these methods are those returned by the test's
// binding. As a result, if a test needs to access a proprietary method provided
// by the binding's implementation of the client, it can simply type assert the
// client to expose the method, e.g.:
//
//	dut.RawAPIs().GNOI(t).(interface {
//		MyProprietaryMethod()
//	}).MyProprietaryMethod()
//
// # Accessing the Binding Device
//
// For still other tests, access to the clients that Ondatra has created is
// insufficient. Some tests may deliberately need to create multiple clients, or
// they may need fine-grained control over how the client is dialed, such as
// special context metadata or dial options that the binding does not provide by
// default. For these situations, a test should retrieve the [binding.DUT] and
// [binding.ATE] instances underlying a DUT or ATE with the calls
// "dut.RawAPIs().BindingDUT()" and "ate.RawAPIs().BindingATE()," respectively.
// With these, they can dial the client themselves, like so:
//
//	gnmiClient, err := dut.RawAPIs().BindingDUT().DialGNMI(ctx, gnmiDialOpts...)
//	if err != nil {
//		t.Fatal("DialGNMI failed: %v", err)
//	}
//	otgClient, err := ate.RawAPIs().BindingATE().DialOTG(ctx, otgDialOpts...)
//	if err != nil {
//		t.Fatal("DialOTG failed: %v", err)
//	}
//
// If a test needs to dial proprietary clients, it can type assert the device
// returned by the [DUTAPIs.BindingDUT] or [ATEAPIs.BindingATE] functions to one
// that supports the method of interest. However, because the [binding.DUT] or
// [binding.ATE] instance with the proprietary method may be embedded inside
// a DUT or ATE wrapper type, a simple type assertion may fail. For that reason,
// the assertion on a DUT or ATE should use the utility functions [binding.DUTAs]
// or [binding.ATEAs] instead:
//
//	var target interface {
//		MyProprietaryClient()
//	}
//	if err := binding.DUTAs(dut, &target) {
//		t.Fatalf("DUT does not support MyProprietaryClient(): ", err)
//	}
//
// [gNMI API]: https://pkg.go.dev/github.com/openconfig/ondatra/gnmi
package raw

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/gnoigo"
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
func (r *DUTAPIs) GNOI(t testing.TB) gnoigo.Clients {
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
