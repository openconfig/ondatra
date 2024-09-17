// Copyright 2023 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package introspect provides an API to introspect grpc binding connections.
package introspect

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/binding"
	"google.golang.org/grpc"
)

// Service is a grpc service.
// The package includes an enumeration of well-known open source services,
// but bindings may define their own services.
type Service string

// Enumeration of well-known open source services.
const (
	GNMI  Service = "gNMI"
	GNOI  Service = "gNOI"
	GNPSI Service = "gNPSI"
	GNSI  Service = "gNSI"
	GRIBI Service = "gRIBI"
	OTG   Service = "OTG"
	P4RT  Service = "P4RT"
)

// DUTDialer returns the grpc dialer for the specified DUT service.
// It fails if the device does not meet the Introspector interface or if the
// connection details cannot be retrieved.
func DUTDialer(t testing.TB, dut *ondatra.DUTDevice, service Service) *Dialer {
	t.Helper()
	var i Introspector
	if err := binding.DUTAs(dut.RawAPIs().BindingDUT(), &i); err != nil {
		t.Fatalf("DUT does not support Introspector interface: %v", err)
	}
	cd, err := i.Dialer(service)
	if err != nil {
		t.Fatalf("Failed to retrieve DUT grpc dialer: %v", err)
	}
	return cd
}

// ATEDialer returns the grpc dialer for the specified ATE service.
// It fails if the device does not meet the Introspector interface or if the
// connection details cannot be retrieved.
func ATEDialer(t testing.TB, ate *ondatra.ATEDevice, service Service) *Dialer {
	t.Helper()
	var i Introspector
	if err := binding.ATEAs(ate.RawAPIs().BindingATE(), &i); err != nil {
		t.Fatalf("ATE does not support Introspector interface: %v", err)
	}
	cd, err := i.Dialer(service)
	if err != nil {
		t.Fatalf("Failed to retrieve ATE grpc dialer: %v", err)
	}
	return cd
}

// Introspector is an interface to introspect grpc service dialers.
type Introspector interface {
	Dialer(Service) (*Dialer, error)
}

// Dialer is capable of dialing grpc endpoints.
type Dialer struct {
	DialFunc   func(context.Context, string, ...grpc.DialOption) (*grpc.ClientConn, error)
	DialTarget string
	DialOpts   []grpc.DialOption
	DevicePort int
}

// Dial dials a gRPC connection using the provided connection details.
func (d *Dialer) Dial(ctx context.Context, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(append([]grpc.DialOption{}, d.DialOpts...), opts...)
	return d.DialFunc(ctx, d.DialTarget, opts...)
}
