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
	plqpb "github.com/openconfig/gnoi/packet_link_qualification"
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
	//
	// Devices in the returned reservation should be initialiazed with a fixed
	// base configuration. Implementations are encouraged to make this base
	// configuration "minimal," meaning a configuration that ensures the device
	// is reachable and capable of being configured further, but little else,
	// so that tests make as few assumptions about the devices as possible.
	Reserve(ctx context.Context, tb *opb.Testbed, runTime, waitTime time.Duration, partial map[string]string) (*Reservation, error)

	// Release releases the reserved testbed.
	// It is not called if the reservation was fetched.
	Release(ctx context.Context) error

	// FetchReservation looks up and returns the pre-existing reservation with
	// the specified ID.
	FetchReservation(ctx context.Context, id string) (*Reservation, error)
}

// Reservation holds the reserved DUTs and ATEs as an id map.
type Reservation struct {
	ID   string
	DUTs map[string]DUT
	ATEs map[string]ATE
}

// Device is a reserved DUT or ATE.
type Device interface {
	Name() string
	Vendor() opb.Device_Vendor
	HardwareModel() string
	SoftwareVersion() string
	Ports() map[string]*Port
	CustomData() map[string]any
}

// Dims contains the dimensions of reserved DUT or ATE.
type Dims struct {
	Name            string
	Vendor          opb.Device_Vendor
	HardwareModel   string
	SoftwareVersion string
	Ports           map[string]*Port
	CustomData      map[string]any
}

func (d *Dims) String() string {
	return fmt.Sprintf("Dims%+v", *d)
}

// DUT is a reserved DUT.
// All implementations of this interface must embed AbstractDUT.
type DUT interface {
	Device

	// PushConfig adds config to the device. If reset is true, the device
	// is reset to its base config before the specified config is added.
	// The following Go template functions are allowed in config:
	// - {{ port "<portID>" }}: replaced with the physical port name
	// - {{ secrets "<arg1>" "<arg2>" }}: left untouched, returned as-is
	PushConfig(ctx context.Context, config string, reset bool) error

	// DialCLI creates a client connection to the DUT's CLI endpoint.
	DialCLI(context.Context) (StreamClient, error)

	// DialConsole creates a client connection to the DUT's Console endpoint.
	DialConsole(context.Context) (StreamClient, error)

	// DialGNMI creates a client connection to the DUT's gNMI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialGNMI(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error)

	// DialGNOI creates a client connection to the DUT's gNOI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialGNOI(context.Context, ...grpc.DialOption) (GNOIClients, error)

	// DialGRIBI creates a client connection to the DUT's gRIBI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialGRIBI(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error)

	// DialP4RT creates a client connection to the DUT's P4RT endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialP4RT(context.Context, ...grpc.DialOption) (p4pb.P4RuntimeClient, error)

	mustEmbedAbstractDUT()
}

// ATE is a reserved ATE.
// All implementations of this interface must embed AbstractATE.
type ATE interface {
	Device

	// DialIxNetwork creates a client connection to the ATE's IxNetwork endpoint.
	DialIxNetwork(context.Context) (*IxNetwork, error)

	// DialGNMI creates a client connection to the ATE's gNMI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	// This method must be implemented to receive gNMI from OTG but not from IxNetwork;
	// Implementing DialIxNetwork is sufficient for gNMI support for IxNetwork.
	DialGNMI(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error)

	// DialOTG creates a client connection to the ATE's OTG endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialOTG(context.Context, ...grpc.DialOption) (gosnappi.GosnappiApi, error)

	mustEmbedAbstractATE()
}

// Port is a reserved Port.
type Port struct {
	Name      string
	Speed     opb.Port_Speed
	CardModel string
	PMD       opb.Port_Pmd
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
// All implementations of this interface must embed AbstractGNOIClients.
type GNOIClients interface {
	BGP() bpb.BGPClient
	CertificateManagement() cpb.CertificateManagementClient
	Diag() dpb.DiagClient
	FactoryReset() frpb.FactoryResetClient
	File() fpb.FileClient
	Healthz() hpb.HealthzClient
	Interface() ipb.InterfaceClient
	Layer2() lpb.Layer2Client
	LinkQualification() plqpb.LinkQualificationClient
	MPLS() mpb.MPLSClient
	OS() ospb.OSClient
	OTDR() otpb.OTDRClient
	System() spb.SystemClient
	WavelengthRouter() wpb.WavelengthRouterClient
	mustEmbedAbstractGNOIClients()
}

// StreamClient provides the interface for streaming IO to DUT.
// All implementations of this interface must embed AbstractStreamClient.
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
	mustEmbedAbstractStreamClient()
}
