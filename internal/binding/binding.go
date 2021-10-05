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
		log.Fatal("binding already initialized")
	}
	bind = b
}

// IsSet returns if binding is set.
func IsSet() bool {
	return bind != nil
}

// Get gets the Ondatra binding.
func Get() Binding {
	if bind == nil {
		log.Fatal("binding not set; did you pass a binding to ondatra.RunTests")
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

	// DialATEGNMI creates a client connection to the GNMI endpoint for the specified ATE.
	// Implementations must append transport security options necessary to reach the server.
	DialATEGNMI(ctx context.Context, ate *reservation.ATE, opts ...grpc.DialOption) (gpb.GNMIClient, error)

	// DialConsole creates a client connection to the specified DUT's Console endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialConsole(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (StreamClient, error)

	// DialCLI creates a client connection to the specified DUT's CLI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialCLI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (StreamClient, error)

	// PushTopology pushes a topology to the ATE.
	// The framework has already verified that there is at least one interface
	// and that each port belongs to at most one port bundle.
	PushTopology(ate *reservation.ATE, topology *opb.Topology) error

	// UpdateTopology updates a topology on the ATE.
	UpdateTopology(ate *reservation.ATE, topology *opb.Topology) error

	// UpdateBGPPeerStates updates the states of BGP peers on interfaces on the ATE.
	// TODO: Remove this method once new Ixia config binding is used.
	UpdateBGPPeerStates(ate *reservation.ATE, interfaces []*opb.InterfaceConfig) error

	// StartProtocols starts control plane protocols on the ATE.
	StartProtocols(ate *reservation.ATE) error

	// StopProtocols stops control plane protocols on the ATE.
	StopProtocols(ate *reservation.ATE) error

	// StartTraffic starts traffic flows on the ATE.
	// All the flows are already verified to belong to the ATE.
	StartTraffic(ate *reservation.ATE, flows []*opb.Flow) error

	// UpdateTraffic updates traffic flows on the ATE.
	// Implementations must return an error if any of the flows have not been started.
	// All the flows are already verified to belong to the ATE.
	UpdateTraffic(ate *reservation.ATE, flows []*opb.Flow) error

	// StopTraffic stops all traffic flows on the ATE.
	StopTraffic(ate *reservation.ATE) error

	// HandleInfraFail handles the given error as an infrastructure failure.
	// If an error is a failure of the Ondatra server or binding implementation
	// rather than user error, it will be passed to HandleInfraFail, which can
	// classify the error as such to distinguish it from a genuine test failure.
	HandleInfraFail(err error) error

	// SetATEPortState sets the enabled state of a physical port on the ATE.
	SetATEPortState(ate *reservation.ATE, port string, enabled bool) error

	// SetTestMetadata sets the metadata for the currently running test.
	SetTestMetadata(md *TestMetadata) error

	// RestartRouting restarts routing on the specified target.
	RestartRouting(dut *reservation.DUT) error
}

// ConfigOptions is a set of options for the config push.
type ConfigOptions struct {
	OpenConfig, Append bool
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
