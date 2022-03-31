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

// Package binding holds the Ondatra binding interface.
package binding

import (
	"fmt"
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding/ixweb"
	"google.golang.org/grpc"

	gnmiclient "github.com/openconfig/gnmi/client"
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
	grpb "github.com/openconfig/gribi/v1/proto/service"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"
)

// Binding is a strategy interface for Ondatra vendor implementations.
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
	//
	// The 'partial' map gives a partial mapping of device and port IDs in the
	// testbed to concrete names to constrain the topology that is reserved.
	Reserve(ctx context.Context, tb *opb.Testbed, runTime, waitTime time.Duration, partial map[string]string) (*Reservation, error)

	// Release releases the reserved testbed.
	// It is not called if the reservation was fetched.
	Release(ctx context.Context) error

	// FetchReservation looks up and returns the pre-existing reservation with
	// the specified ID.
	FetchReservation(ctx context.Context, id string) (*Reservation, error)

	// PushConfig adds config to the specified device. If reset is true, the
	// device is reset to its base config before the specified config is added.
	// The following Go template functions are allowed in config:
	// - {{ port "<portID>" }}: replaced with the physical port name
	// - {{ secrets "<arg1>" "<arg2>" }}: left untouched, returned as-is
	PushConfig(ctx context.Context, dut *DUT, config string, reset bool) error

	// DialGNMI creates a client connection to the specified DUT's gNMI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialGNMI(ctx context.Context, dut *DUT, opts ...grpc.DialOption) (gpb.GNMIClient, error)

	// DialGNOI creates a client connection to the specified DUT's gNOI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialGNOI(ctx context.Context, dut *DUT, opts ...grpc.DialOption) (GNOIClients, error)

	// DialGRIBI creates a client connection to the specified DUT's gRIBI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialGRIBI(ctx context.Context, dut *DUT, opts ...grpc.DialOption) (grpb.GRIBIClient, error)

	// DialP4RT creates a client connection to the specified DUT's P4RT endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialP4RT(ctx context.Context, dut *DUT, opts ...grpc.DialOption) (p4pb.P4RuntimeClient, error)

	// DialConsole creates a client connection to the specified DUT's Console endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialConsole(ctx context.Context, dut *DUT, opts ...grpc.DialOption) (StreamClient, error)

	// DialCLI creates a client connection to the specified DUT's CLI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialCLI(ctx context.Context, dut *DUT, opts ...grpc.DialOption) (StreamClient, error)

	// DialIxNetwork creates a client connection to the specified ATE's IxNetwork endpoint.
	DialIxNetwork(ctx context.Context, ate *ATE) (*IxNetwork, error)

	// DialOTG creates a client commnetion to the specified ATE's OTG endpoint.
	DialOTG(ctx context.Context, ate *ATE) (gosnappi.GosnappiApi, error)

	// DialOTGGNMI creates a client commnetion to the specified ATE's OTG gNMI endpoint.
	DialOTGGNMI(ate *ATE) (*gnmiclient.Query, error)

	// DialATEGNMI creates a client connection to the specified ATE's gNMI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialATEGNMI(ctx context.Context, ate *ATE, opts ...grpc.DialOption) (gpb.GNMIClient, error)

	// HandleInfraFail handles the given error as an infrastructure failure.
	// If an error is a failure of the Ondatra server or binding implementation
	// rather than user error, it will be passed to HandleInfraFail, which can
	// classify the error as such to distinguish it from a genuine test failure.
	HandleInfraFail(err error) error

	// SetTestMetadata sets the metadata for the currently running test.
	SetTestMetadata(md *TestMetadata) error
}

// Reservation holds the reserved DUTs and ATEs as an id map.
type Reservation struct {
	ID   string
	DUTs map[string]*DUT
	ATEs map[string]*ATE
}

// Device is a reserved DUT or ATE.
type Device interface {
	Dimensions() *Dims
}

// Dims contains the dimensions of reserved DUT or ATE.
type Dims struct {
	Name            string
	Vendor          opb.Device_Vendor
	HardwareModel   string
	SoftwareVersion string
	Ports           map[string]*Port
}

func (d *Dims) String() string {
	return fmt.Sprintf("Dims%+v", *d)
}

// DUT is a reserved DUT
type DUT struct {
	*Dims
}

// Dimensions returns the dimensions of the device.
func (d *DUT) Dimensions() *Dims {
	return d.Dims
}

func (d *DUT) String() string {
	return fmt.Sprintf("DUT%+v", *d)
}

// ATE is a reserved ATE.
type ATE struct {
	*Dims
}

// Dimensions returns the dimensions of the device.
func (a *ATE) Dimensions() *Dims {
	return a.Dims
}

func (a *ATE) String() string {
	return fmt.Sprintf("ATE%+v", *a)
}

// Port is a reserved Port.
type Port struct {
	Name  string
	Speed opb.Port_Speed
}

func (p *Port) String() string {
	return fmt.Sprintf("Port%+v", *p)
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
