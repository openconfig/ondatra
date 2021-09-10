package ondatra

import (
	"golang.org/x/net/context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/negtest"

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
)

var (
	gotConfig string
	gotOpts   *binding.ConfigOptions
)

func init() {
	initFakeBinding()
	reserveFakeTestbed()
	fakeBind.ConfigPusher = func(_ context.Context, _ *reservation.DUT, config string, opts *binding.ConfigOptions) error {
		gotConfig = config
		gotOpts = opts
		return nil
	}
}

func TestPushConfig(t *testing.T) {
	dutArista := DUT(t, "dut")
	testsPass := []struct {
		desc       string
		config     *DUTConfig
		wantConfig string
		wantOpts   *binding.ConfigOptions
	}{{
		desc: "correct config text",
		config: dutArista.Config().New().
			WithAristaText("Arista config").
			WithCiscoText("Cisco config").
			WithJuniperText("Juniper config"),
		wantConfig: "Arista config",
		wantOpts:   &binding.ConfigOptions{},
	}, {
		desc: "correct config file",
		config: dutArista.Config().New().
			WithAristaFile(filepath.Join("testdata", "example_config_1.txt")).
			WithCiscoText("Cisco config").
			WithJuniperFile(filepath.Join("testdata", "example_config_2.txt")),
		wantConfig: "example_config_1",
		wantOpts:   &binding.ConfigOptions{},
	}, {
		desc: "correct openconfig",
		config: dutArista.Config().New().
			WithOpenConfigText("Openconfig").
			WithCiscoText("Cisco config").
			WithJuniperText("Juniper config"),
		wantConfig: "Openconfig",
		wantOpts:   &binding.ConfigOptions{OpenConfig: true},
	}, {
		desc: "correct openconfig file",
		config: dutArista.Config().New().
			WithOpenConfigFile(filepath.Join("testdata", "example_config_1.txt")).
			WithCiscoText("Cisco config").
			WithJuniperFile(filepath.Join("testdata", "example_config_2.txt")),
		wantConfig: "example_config_1",
		wantOpts:   &binding.ConfigOptions{OpenConfig: true},
	}, {
		desc: "openconfig override",
		config: dutArista.Config().New().
			WithOpenConfigText("Openconfig").
			WithAristaText("Arista config"),
		wantConfig: "Arista config",
		wantOpts:   &binding.ConfigOptions{},
	}, {
		desc: "port template",
		config: dutArista.Config().New().
			WithAristaText(`reconfigure {{ port "port1" }} and {{ port "port2" }}`),
		wantConfig: "reconfigure Et1/2/3 and Et4/5/6",
		wantOpts:   &binding.ConfigOptions{},
	}, {
		desc: "secrets template",
		config: dutArista.Config().New().
			WithAristaText(`shh {{ secrets "hello" "there" }} wink`),
		wantConfig: `shh {{ secrets "hello" "there" }} wink`,
		wantOpts:   &binding.ConfigOptions{},
	}, {
		desc: "var template",
		config: dutArista.Config().New().
			WithAristaText(`hello {{ var "foo" }} there`).
			WithVarValue("foo", "bar"),
		wantConfig: `hello bar there`,
		wantOpts:   &binding.ConfigOptions{},
	}, {
		desc: "var map template",
		config: dutArista.Config().New().
			WithAristaText(`hello {{ var "x" }} and {{ var "y" }}`).
			WithVarMap(map[string]string{"x": "apple", "y": "orange"}),
		wantConfig: `hello apple and orange`,
		wantOpts:   &binding.ConfigOptions{},
	}}

	for _, tt := range testsPass {
		t.Run(tt.desc, func(t *testing.T) {
			gotConfig = ""
			gotOpts = nil
			tt.config.Push(t)
			if diff := cmp.Diff(tt.wantConfig, gotConfig); diff != "" {
				t.Errorf("Push(t) got unexpected config diff(-want,+got):\n %s", diff)
			}
			if diff := cmp.Diff(tt.wantOpts, gotOpts); diff != "" {
				t.Errorf("Push(t) got unexpected options diff(-want,+got):\n %s", diff)
			}
		})
	}
}

func TestPushConfigErrors(t *testing.T) {
	dutArista := DUT(t, "dut")
	testsFail := []struct {
		desc         string
		config       *DUTConfig
		wantFatalMsg string
	}{{
		desc:         "no config",
		config:       dutArista.Config().New().WithCiscoText("gaga"),
		wantFatalMsg: "vendor",
	}, {
		desc:         "template function does not exist",
		config:       dutArista.Config().New().WithAristaText(`{{ qwerty "port1" }}`),
		wantFatalMsg: `function "qwerty" not defined`,
	}, {
		desc:         "port name does not exist",
		config:       dutArista.Config().New().WithAristaText(`{{ port "port10" }}`),
		wantFatalMsg: "port10 not found",
	}, {
		desc:         "template malformed",
		config:       dutArista.Config().New().WithAristaText(`{{ port"port10" }}`),
		wantFatalMsg: "bad character",
	}, {
		desc:         "var has no value",
		config:       dutArista.Config().New().WithAristaText(`{{ var "key1" }}`),
		wantFatalMsg: "No value for key",
	}}

	for _, tt := range testsFail {
		t.Run(tt.desc, func(t *testing.T) {
			got := negtest.ExpectFatal(t, func(t testing.TB) {
				tt.config.Push(t)
			})
			if !strings.Contains(got, tt.wantFatalMsg) {
				t.Errorf("Push(t) failed with message %q, want %q", got, tt.wantFatalMsg)
			}
		})
	}
}

func TestAppendConfig(t *testing.T) {
	gotConfig = ""
	gotOpts = nil
	wantConfig := "arista config"
	wantOpts := &binding.ConfigOptions{Append: true}
	DUT(t, "dut").Config().New().WithAristaText(wantConfig).Append(t)
	if gotConfig != wantConfig {
		t.Errorf("Append(t) got pushed config %v, want %v", gotConfig, wantConfig)
	}
	if !cmp.Equal(gotOpts, wantOpts) {
		t.Errorf("Append(t) got pushed options %v, want %v", gotOpts, wantOpts)
	}
}

func TestGNMI(t *testing.T) {
	want := struct{ gpb.GNMIClient }{}
	fakeBind.GNMIDialer = func(context.Context, *reservation.DUT, ...grpc.DialOption) (gpb.GNMIClient, error) {
		return want, nil
	}
	if got := DUT(t, "dut").RawAPIs().GNMI(t); got != want {
		t.Errorf("GNMI(t) got %v, want %v", got, want)
	}
}

func TestGNMIError(t *testing.T) {
	wantErr := "bad gnmi"
	fakeBind.GNMIDialer = func(context.Context, *reservation.DUT, ...grpc.DialOption) (gpb.GNMIClient, error) {
		return nil, errors.New(wantErr)
	}
	raw := DUT(t, "dut_cisco").RawAPIs()
	gotErr := negtest.ExpectFatal(t, func(t testing.TB) {
		raw.GNMI(t)
	})
	if !strings.Contains(gotErr, wantErr) {
		t.Errorf("GNMI(t) got err %v, want %v", gotErr, wantErr)
	}
}

type gnoiClients struct {
	binding.GNOIClients
	bgp          bpb.BGPClient
	certMgmt     cpb.CertificateManagementClient
	diag         dpb.DiagClient
	factoryReset frpb.FactoryResetClient
	file         fpb.FileClient
	healthz      hpb.HealthzClient
	intface      ipb.InterfaceClient
	layer2       lpb.Layer2Client
	mpls         mpb.MPLSClient
	os           ospb.OSClient
	otdr         otpb.OTDRClient
	system       spb.SystemClient
	waveRtr      wpb.WavelengthRouterClient
	custom       interface{}
}

func (g *gnoiClients) BGP() bpb.BGPClient {
	return g.bgp
}

func (g *gnoiClients) CertificateManagement() cpb.CertificateManagementClient {
	return g.certMgmt
}

func (g *gnoiClients) Diag() dpb.DiagClient {
	return g.diag
}

func (g *gnoiClients) FactoryReset() frpb.FactoryResetClient {
	return g.factoryReset
}

func (g *gnoiClients) File() fpb.FileClient {
	return g.file
}

func (g *gnoiClients) Healthz() hpb.HealthzClient {
	return g.healthz
}

func (g *gnoiClients) Interface() ipb.InterfaceClient {
	return g.intface
}

func (g *gnoiClients) Layer2() lpb.Layer2Client {
	return g.layer2
}

func (g *gnoiClients) MPLS() mpb.MPLSClient {
	return g.mpls
}

func (g *gnoiClients) OS() ospb.OSClient {
	return g.os
}

func (g *gnoiClients) OTDR() otpb.OTDRClient {
	return g.otdr
}

func (g *gnoiClients) System() spb.SystemClient {
	return g.system
}

func (g *gnoiClients) WavelengthRouter() wpb.WavelengthRouterClient {
	return g.waveRtr
}

func TestGNOI(t *testing.T) {
	bgnoi := &gnoiClients{
		bgp: struct{ bpb.BGPClient }{},
		certMgmt: struct {
			cpb.CertificateManagementClient
		}{},
		diag:         struct{ dpb.DiagClient }{},
		factoryReset: struct{ frpb.FactoryResetClient }{},
		file:         struct{ fpb.FileClient }{},
		healthz:      struct{ hpb.HealthzClient }{},
		intface:      struct{ ipb.InterfaceClient }{},
		layer2:       struct{ lpb.Layer2Client }{},
		mpls:         struct{ mpb.MPLSClient }{},
		os:           struct{ ospb.OSClient }{},
		otdr:         struct{ otpb.OTDRClient }{},
		system:       struct{ spb.SystemClient }{},
		waveRtr:      struct{ wpb.WavelengthRouterClient }{},
		custom:       struct{}{},
	}
	fakeBind.GNOIDialer = func(context.Context, *reservation.DUT, ...grpc.DialOption) (binding.GNOIClients, error) {
		return bgnoi, nil
	}
	gnoi := DUT(t, "dut").RawAPIs().GNOI(t)
	if got := gnoi.BGP(); got != bgnoi.BGP() {
		t.Errorf("GNOI(t) got BGP client %v, want %v", got, bgnoi.BGP)
	}
	if got := gnoi.CertificateManagement(); got != bgnoi.CertificateManagement() {
		t.Errorf("GNOI(t) got CertificateManagement client %v, want %v", got, bgnoi.CertificateManagement)
	}
	if got := gnoi.Diag(); got != bgnoi.Diag() {
		t.Errorf("GNOI(t) got Diag client %v, want %v", got, bgnoi.Diag)
	}
	if got := gnoi.FactoryReset(); got != bgnoi.FactoryReset() {
		t.Errorf("GNOI(t) got FactoryReset client %v, want %v", got, bgnoi.FactoryReset)
	}
	if got := gnoi.File(); got != bgnoi.File() {
		t.Errorf("GNOI(t) got File client %v, want %v", got, bgnoi.File)
	}
	if got := gnoi.Healthz(); got != bgnoi.Healthz() {
		t.Errorf("GNOI(t) got Healthz client %v, want %v", got, bgnoi.Healthz)
	}
	if got := gnoi.Interface(); got != bgnoi.Interface() {
		t.Errorf("GNOI(t) got Interface client %v, want %v", got, bgnoi.Interface)
	}
	if got := gnoi.Layer2(); got != bgnoi.Layer2() {
		t.Errorf("GNOI(t) got Layer2 client %v, want %v", got, bgnoi.Layer2)
	}
	if got := gnoi.MPLS(); got != bgnoi.MPLS() {
		t.Errorf("GNOI(t) got MPLS client %v, want %v", got, bgnoi.MPLS)
	}
	if got := gnoi.OS(); got != bgnoi.OS() {
		t.Errorf("GNOI(t) got OS client %v, want %v", got, bgnoi.OS)
	}
	if got := gnoi.OTDR(); got != bgnoi.OTDR() {
		t.Errorf("GNOI(t) got OTDRS client %v, want %v", got, bgnoi.OTDR)
	}
	if got := gnoi.System(); got != bgnoi.System() {
		t.Errorf("GNOI(t) got System client %v, want %v", got, bgnoi.System)
	}
	if got := gnoi.WavelengthRouter(); got != bgnoi.WavelengthRouter() {
		t.Errorf("GNOI(t) got WavelengthRouter client %v, want %v", got, bgnoi.WavelengthRouter)
	}
	if got := gnoi.(*gnoiClients).custom; got != bgnoi.custom {
		t.Errorf("GNOI(t) got custom client %v, want %v", got, bgnoi.custom)
	}
}

func TestGNOIError(t *testing.T) {
	wantErr := "bad gnoi"
	fakeBind.GNOIDialer = func(context.Context, *reservation.DUT, ...grpc.DialOption) (binding.GNOIClients, error) {
		return nil, errors.New(wantErr)
	}
	raw := DUT(t, "dut_cisco").RawAPIs()
	gotErr := negtest.ExpectFatal(t, func(t testing.TB) {
		raw.GNOI(t)
	})
	if !strings.Contains(gotErr, wantErr) {
		t.Errorf("GNOI(t) got err %v, want %v", gotErr, wantErr)
	}
}
