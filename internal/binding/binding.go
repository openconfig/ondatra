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

// Package binding holds the server binding interface.
package binding

import (
	"golang.org/x/net/context"
	"io"
	"time"

	log "github.com/golang/glog"
	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/internal/ixweb"
	"github.com/openconfig/ondatra/internal/reservation"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	bpb "github.com/openconfig/gnoi/bgp"
	cpb "github.com/openconfig/gnoi/cert"
	dpb "github.com/openconfig/gnoi/diag"
	frpb "github.com/openconfig/gnoi/factory_reset"
	fpb "github.com/openconfig/gnoi/file"
	hpb "github.com/openconfig/gnoi/healthz"
	ipb "github.com/openconfig/gnoi/interface"
	lpb "github.com/openconfig/gnoi/layer2"
	mpb "github.com/openconfig/gnoi/mpls"
	ospb "github.com/openconfig/gnoi/os"
	otpb "github.com/openconfig/gnoi/otdr"
	spb "github.com/openconfig/gnoi/system"
	wpb "github.com/openconfig/gnoi/wavelength_router"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"

)

var (
	bind Binding
)

// Init initializes the Ondatra binding.
func Init(b Binding) {
	if bind != nil {
		log.Fatalf("Binding already initialized")
	}
	bind = b
}

// Get gets the Ondatra binding.
func Get() Binding {
	if bind == nil {
		log.Exit("Binding not initialized. Did you forget to call ondatra.RunTests in TestMain?")
	}
	return bind
}

// Binding is a strategy interface for Ondatra server implementations.
//
// The framework enforces that at most testbed is reserved at a time, so
// implementations can assume that these methods are never called out of order,
// e.g. Release() is never be called without a prior Reserve().
//
// By default, errors returned by binding methods will be reported as
// "infrastructure failures," meaning the binding implementation itself is
// responsible for the error, not the user. To generate a user error instead,
// the binding should create or wrap an error using the "usererr" package.
// All infrastructure failure errors are passed to ReportInfraFail.
type Binding interface {

	// Reserve reserves resources matching the criteria in the specified testbed.
	// The framework has already verified that the testbed is valid, and will
	// validate that the returned reservation matches the testbed criteria.
	//
	// The testbed resources must be reserved for the specified runTime, or
	// until Release is called. A runTime of zero means the caller requested an
	// unlimited runTime. The implementation is free to impose a maximum limit on
	// the runTime and return an error if it exceeds that limit. The framework has
	// already checked that the runTime is not negative.
	//
	// This method may block for up to the specified waitTime for all testbed
	// resources to become available. Given a zero waitTime, the implementation
	// must choose a reasonable duration. The framework has already checked that
	// the waitTime is not negative.
	Reserve(ctx context.Context, tb *opb.Testbed, runTime, waitTime time.Duration) (*reservation.Reservation, error)

	// Release releases the reserved testbed.
	Release(ctx context.Context) error

	// PushConfig sets config on the specified device.
	// The following Go template functions are allowed in config:
	// - {{ port "<portID>" }}: replaced with the physical port name
	// - {{ secrets "<arg1>" "<arg2>" }}: left untouched, returned as-is
	// If the openconfig option is true, the config is in openconfig JSON syntax.
	// If the append option is true, the config is appended to the existing config;
	// otherwise the existing config is replaced with the provided config.
	PushConfig(ctx context.Context, dut *reservation.DUT, config string, opts *ConfigOptions) error

	// DialGNMI creates a client connection to the specified DUT's gNMI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialGNMI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (gpb.GNMIClient, error)

	// DialGNOI creates a client connection to the specified DUT's gNOI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialGNOI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (GNOIClients, error)

	// DialP4RT creates a client connection to the specified DUT's P4RT endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialP4RT(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (p4pb.P4RuntimeClient, error)

	// DialConsole creates a client connection to the specified DUT's Console endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialConsole(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (StreamClient, error)

	// DialCLI creates a client connection to the specified DUT's CLI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialCLI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (StreamClient, error)

	// DialIxNetwork creates a client connection to the specified ATE's IxNetwork endpoint.
	DialIxNetwork(ctx context.Context, ate *reservation.ATE) (*IxNetwork, error)

	// HandleInfraFail handles the given error as an infrastructure failure.
	// If an error is a failure of the Ondatra server or binding implementation
	// rather than user error, it will be passed to HandleInfraFail, which can
	// classify the error as such to distinguish it from a genuine test failure.
	HandleInfraFail(err error) error

	// SetTestMetadata sets the metadata for the currently running test.
	SetTestMetadata(md *TestMetadata) error
}

// ConfigOptions is a set of options for the config push.
type ConfigOptions struct {
	OpenConfig, Append bool
}

// IxNetwork provides information for an IxNetwork session.
type IxNetwork struct {
	// Session is an IxNetwork session for an ATE.
	Session *ixweb.Session
	// ChassisHost is an optional hostname or IP address to use when assigning
	// Ixia VPorts. If empty, defaults to the name of the reserved ATE.
	ChassisHost string
	// SyslogHost is an optional hostname or IP address to which IxNetwork should
	// stream logs. If empty, syslog streaming is not enabled.
	SyslogHost string
}

// GNOIClients stores APIs to GNOI services.
type GNOIClients interface {
	BGP() bpb.BGPClient
	CertificateManagement() cpb.CertificateManagementClient
	Diag() dpb.DiagClient
	FactoryReset() frpb.FactoryResetClient
	File() fpb.FileClient
	Healthz() hpb.HealthzClient
	Interface() ipb.InterfaceClient
	Layer2() lpb.Layer2Client
	MPLS() mpb.MPLSClient
	OS() ospb.OSClient
	OTDR() otpb.OTDRClient
	System() spb.SystemClient
	WavelengthRouter() wpb.WavelengthRouterClient
}

// TestMetadata is metadata about a test.
type TestMetadata struct {
	TestName string
}

// StreamClient provides the interface for streaming IO to DUT.
type StreamClient interface {
	// SendCommand always expects a clean "prompt" on the underlying
	// device. If the device is interleaving Stdin writes with SendCommand
	// the underlying channel must be at a clean prompt to allow SendCommand
	// to complete the operation. Additionaly, the result buffer will be fully
	// consumed by SendCommand, ensuring it leaves a clean prompt behind.
	SendCommand(context.Context, string) (string, error)
	Stdin() io.WriteCloser
	Stdout() io.ReadCloser
	Stderr() io.ReadCloser
	Close() error
}
