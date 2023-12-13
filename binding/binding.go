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
	"github.com/openconfig/gnoigo"
	"github.com/openconfig/ondatra/binding/ixweb"
	"google.golang.org/grpc"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	acctzpb "github.com/openconfig/gnsi/acctz"
	authzpb "github.com/openconfig/gnsi/authz"
	certzpb "github.com/openconfig/gnsi/certz"
	credzpb "github.com/openconfig/gnsi/credentialz"
	pathzpb "github.com/openconfig/gnsi/pathz"
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

// DUT is a reserved DUT.
//
// All implementations of this interface must embed AbstractDUT.
//
// For methods that dial gRPC endpoints, implementations must use a default set
// of dial options necessary to reach the endpoint, so that the dial succeeds
// when no options are provided by the caller. When the caller does specify
// options, implementations are encouraged to append those options to the
// default set, so the call can both inherit and override the default behavior.
type DUT interface {
	Device

	// PushConfig adds config to the device. If reset is true, the device
	// is reset to its base config before the specified config is added.
	// The following Go template functions are allowed in config:
	// - {{ port "<portID>" }}: replaced with the physical port name
	// - {{ secrets "<arg1>" "<arg2>" }}: left untouched, returned as-is
	PushConfig(ctx context.Context, config string, reset bool) error

	// DialCLI creates a client connection to the DUT's CLI endpoint.
	DialCLI(context.Context) (CLIClient, error)

	// DialConsole creates a client connection to the DUT's Console endpoint.
	DialConsole(context.Context) (ConsoleClient, error)

	// DialGNMI creates a client connection to the DUT's gNMI endpoint.
	// See the interface comment for proper handling of dial options.
	DialGNMI(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error)

	// DialGNOI creates a client connection to the DUT's gNOI endpoint.
	// See the interface comment for proper handling of dial options.
	DialGNOI(context.Context, ...grpc.DialOption) (gnoigo.Clients, error)

	// DialGNSI creates a client connection to the DUT's gNSI endpoint.
	// See the interface comment for proper handling of dial options.
	DialGNSI(context.Context, ...grpc.DialOption) (GNSIClients, error)

	// DialGRIBI creates a client connection to the DUT's gRIBI endpoint.
	// See the interface comment for proper handling of dial options.
	DialGRIBI(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error)

	// DialP4RT creates a client connection to the DUT's P4RT endpoint.
	// See the interface comment for proper handling of dial options.
	DialP4RT(context.Context, ...grpc.DialOption) (p4pb.P4RuntimeClient, error)

	mustEmbedAbstractDUT()
}

// ATE is a reserved ATE.
//
// All implementations of this interface must embed AbstractATE.
//
// For methods that dial gRPC endpoints, implementations must use a default set
// of dial options necessary to reach the endpoint, so that the dial succeeds
// when no options are provided by the caller. When the caller does specify
// options, implementations are encouraged to append those options to the
// default set, so the call can both inherit and override the default behavior.
type ATE interface {
	Device

	// DialIxNetwork creates a client connection to the ATE's IxNetwork endpoint.
	DialIxNetwork(context.Context) (*IxNetwork, error)

	// DialGNMI creates a client connection to the ATE's gNMI endpoint.
	// See the interface comment for proper handling of dial options.
	// This method must be implemented to receive gNMI from OTG but not from IxNetwork;
	// Implementing DialIxNetwork is sufficient for gNMI support for IxNetwork.
	DialGNMI(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error)

	// DialOTG creates a client connection to the ATE's OTG endpoint.
	// See the interface comment for proper handling of dial options.
	DialOTG(context.Context, ...grpc.DialOption) (gosnappi.Api, error)

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

// GNSIClients stores APIs to GNSI services.
// All implementations of this interface must embed AbstractGNSIClients.
type GNSIClients interface {
	Authz() authzpb.AuthzClient
	Pathz() pathzpb.PathzClient
	Certz() certzpb.CertzClient
	Credentialz() credzpb.CredentialzClient
	Acctz() acctzpb.AcctzClient
	mustEmbedAbstractGNSIClients()
}

// CLIClient provides the interface for sending CLI commands to the DUT.
// All implementations of this interface must embed AbstractCLIClient.
type CLIClient interface {
	RunCommand(context.Context, string) (CommandResult, error)
	mustEmbedAbstractCLIClient()
}

// CommandResult provides the interface for the result of a CLI command.
// All implementations of this interface must embed AbstractCommandResult.
type CommandResult interface {
	// Output returns the output of the command.
	// The return value may be non-empty even when the command fails.
	Output() string
	// Error returns an error message that occurred when running the command.
	// The return value is the empty string if and only if the command succeeds.
	Error() string
	mustEmbedAbstractCommandResult()
}

// ConsoleClient provides the interface for console access to the DUT.
// All implementations of this interface must embed AbstractConsoleClient.
type ConsoleClient interface {
	Stdin() io.WriteCloser
	Stdout() io.ReadCloser
	Stderr() io.ReadCloser
	Close() error
	mustEmbedAbstractConsoleClient()
}
