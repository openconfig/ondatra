// Package binding holds the server binding interface.
package binding

import (
	"context"
	"time"

	"github.com/golang/glog"
	"github.com/golang/grpc/grpc"
	"github.com/openconfig/ondatra/internal/reservation/reservation"

	gpb "github.com/openconfig/gnmi/proto/gnmi/gnmi_go_proto"
	bpb "github.com/openconfig/gnoi/bgp/bgp_go_proto"
	cpb "github.com/openconfig/gnoi/cert/cert_go_proto"
	dpb "github.com/openconfig/gnoi/diag/diag_go_proto"
	frpb "github.com/openconfig/gnoi/factory_reset/factory_reset_go_proto"
	fpb "github.com/openconfig/gnoi/file/file_go_proto"
	hpb "github.com/openconfig/gnoi/healthz/healthz_go_proto"
	ipb "github.com/openconfig/gnoi/interface/interface_go_proto"
	lpb "github.com/openconfig/gnoi/layer2/layer2_go_proto"
	mpb "github.com/openconfig/gnoi/mpls/mpls_go_proto"
	ospb "github.com/openconfig/gnoi/os/os_go_proto"
	otpb "github.com/openconfig/gnoi/otdr/otdr_go_proto"
	spb "github.com/openconfig/gnoi/system/system_go_proto"
	wpb "github.com/openconfig/gnoi/wavelength_router/wavelength_router_go_proto"
	opb "github.com/openconfig/ondatra/proto/ondatra_go_proto"
)

var bind Binding

// Init initializes the Ondatra binding.
func Init(b Binding) {
	if bind != nil {
		glog.Fatal("binding already initialized")
	}
	bind = b
}

// Get gets the Ondatra binding.
func Get() Binding {
	if bind == nil {
		glog.Fatal("binding not initialized")
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
	// until Release is called. The framework has already checked that the
	// runTime is a positive value.
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

	// DialGNMI creates a client connection to the testbed's GNMI server.
	// Implementations must append transport security options necessary to reach the server.
	DialGNMI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (gpb.GNMIClient, error)

	// DialGNOI creates a client connection to the specified DUT's gNOI endpoint.
	// Implementations must append transport security options necessary to reach the server.
	DialGNOI(ctx context.Context, dut *reservation.DUT, opts ...grpc.DialOption) (*GNOIClient, error)

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

// GNOIClient stores APIs to GNOI services.
// Additional client fields will be added to this as needed.
type GNOIClient struct {
	BGP                   bpb.BGPClient
	CertificateManagement cpb.CertificateManagementClient
	Diag                  dpb.DiagClient
	FactoryReset          frpb.FactoryResetClient
	File                  fpb.FileClient
	Healthz               hpb.HealthzClient
	Interface             ipb.InterfaceClient
	Layer2                lpb.Layer2Client
	MPLS                  mpb.MPLSClient
	OS                    ospb.OSClient
	OTDR                  otpb.OTDRClient
	System                spb.SystemClient
	WavelengthRouter      wpb.WavelengthRouterClient
}

// TestMetadata is metadata about a test.
type TestMetadata struct {
	TestName string
}
