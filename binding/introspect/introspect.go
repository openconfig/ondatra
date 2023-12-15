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

// Service is an enum of provided services.
type Service string

// Enums of provided services.
const (
	GNMI  Service = "gnmi"
	GNOI  Service = "gnoi"
	GNSI  Service = "gnsi"
	GRIBI Service = "gribi"
	P4RT  Service = "p4rt"
)

// Introspect returns the grpc connection details for the specified service.
// It fails if the device does not meet the Introspector interface or if the
// connection details cannot be retrieved.
func Introspect(t testing.TB, dut *ondatra.DUTDevice, service Service) *ConnDetails {
	t.Helper()
	var i Introspector
	if err := binding.DUTAs(dut.RawAPIs().BindingDUT(), &i); err != nil {
		t.Fatalf("DUT does not support Introspector interface: %v", err)
	}
	cd, err := i.ConnDetails(service)
	if err != nil {
		t.Fatalf("Failed to retrieve connection details: %v", err)
	}
	return cd
}

// Introspector is an interface to introspect grpc binding connections.
type Introspector interface {
	ConnDetails(Service) (*ConnDetails, error)
}

// ConnDetails provides details of a grpc binding connection.
type ConnDetails struct {
	DevicePort      int
	DefaultDial     func(context.Context, string, ...grpc.DialOption) (*grpc.ClientConn, error)
	DefaultTarget   string
	DefaultDialOpts []grpc.DialOption
}
