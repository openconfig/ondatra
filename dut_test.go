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

package ondatra

import (
	"bufio"
	"golang.org/x/net/context"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"github.com/openconfig/gnmi/errdiff"
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
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakes/fakestreamclient"
	"github.com/openconfig/ondatra/negtest"

)

var (
	gotConfig string
	gotOpts   *binding.ConfigOptions
)

func initDUTFakes(t *testing.T) {
	t.Helper()
	initFakeBinding(t)
	reserveFakeTestbed(t)
	fakeBind.ConfigPusher = func(_ context.Context, _ *binding.DUT, config string, opts *binding.ConfigOptions) error {
		gotConfig = config
		gotOpts = opts
		return nil
	}
}

func TestPushConfig(t *testing.T) {
	initDUTFakes(t)
	dutArista := DUT(t, "dut")
	testsPass := []struct {
		desc       string
		config     *DUTConfig
		wantConfig string
		wantOpts   *binding.ConfigOptions
	}{{
		desc:       "correct text",
		config:     dutArista.Config().New().WithText("generated"),
		wantConfig: "generated",
		wantOpts:   &binding.ConfigOptions{},
	}, {
		desc:       "correct file",
		config:     dutArista.Config().New().WithFile(filepath.Join("testdata", "example_config_1.txt")),
		wantConfig: "example_config_1",
		wantOpts:   &binding.ConfigOptions{},
	}, {
		desc: "correct per-vendor text",
		config: dutArista.Config().New().
			WithAristaText("Arista config").
			WithCiscoText("Cisco config").
			WithJuniperText("Juniper config"),
		wantConfig: "Arista config",
		wantOpts:   &binding.ConfigOptions{},
	}, {
		desc: "correct per-vendor file",
		config: dutArista.Config().New().
			WithAristaFile(filepath.Join("testdata", "example_config_1.txt")).
			WithCiscoText("Cisco config").
			WithJuniperFile(filepath.Join("testdata", "example_config_2.txt")),
		wantConfig: "example_config_1",
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
	initDUTFakes(t)
	dutArista := DUT(t, "dut")
	testsFail := []struct {
		desc         string
		config       *DUTConfig
		wantFatalMsg string
	}{{
		desc:         "no config",
		config:       dutArista.Config().New().WithCiscoText("gaga"),
		wantFatalMsg: "no config",
	}, {
		desc: "all-vendor and per-vendor",
		config: dutArista.Config().New().
			WithText("text1").
			WithAristaText("text2"),
		wantFatalMsg: "all-vendor and per-vendor",
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
	initDUTFakes(t)
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
	initDUTFakes(t)
	want := struct{ gpb.GNMIClient }{}
	fakeBind.GNMIDialer = func(context.Context, *binding.DUT, ...grpc.DialOption) (gpb.GNMIClient, error) {
		return want, nil
	}
	if got := DUT(t, "dut").RawAPIs().GNMI().New(t); got != want {
		t.Errorf("GNMI().New(t) got %v, want %v", got, want)
	}
}

func TestGNMIError(t *testing.T) {
	initDUTFakes(t)
	wantErr := "bad gnmi"
	fakeBind.GNMIDialer = func(context.Context, *binding.DUT, ...grpc.DialOption) (gpb.GNMIClient, error) {
		return nil, errors.New(wantErr)
	}
	raw := DUT(t, "dut_cisco").RawAPIs()
	gotErr := negtest.ExpectFatal(t, func(t testing.TB) {
		raw.GNMI().New(t)
	})
	if !strings.Contains(gotErr, wantErr) {
		t.Errorf("GNMI().New(t) got err %v, want %v", gotErr, wantErr)
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
	initDUTFakes(t)
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
	fakeBind.GNOIDialer = func(context.Context, *binding.DUT, ...grpc.DialOption) (binding.GNOIClients, error) {
		return bgnoi, nil
	}
	gnoi := DUT(t, "dut").RawAPIs().GNOI().New(t)
	if got, want := gnoi.BGP(), bgnoi.BGP(); got != want {
		t.Errorf("GNOI(t) got BGP client %v, want %v", got, want)
	}
	if got, want := gnoi.CertificateManagement(), bgnoi.CertificateManagement(); got != want {
		t.Errorf("GNOI(t) got CertificateManagement client %v, want %v", got, want)
	}
	if got, want := gnoi.Diag(), bgnoi.Diag(); got != want {
		t.Errorf("GNOI(t) got Diag client %v, want %v", got, want)
	}
	if got, want := gnoi.FactoryReset(), bgnoi.FactoryReset(); got != want {
		t.Errorf("GNOI(t) got FactoryReset client %v, want %v", got, want)
	}
	if got, want := gnoi.File(), bgnoi.File(); got != want {
		t.Errorf("GNOI(t) got File client %v, want %v", got, want)
	}
	if got, want := gnoi.Healthz(), bgnoi.Healthz(); got != want {
		t.Errorf("GNOI(t) got Healthz client %v, want %v", got, want)
	}
	if got, want := gnoi.Interface(), bgnoi.Interface(); got != want {
		t.Errorf("GNOI(t) got Interface client %v, want %v", got, want)
	}
	if got, want := gnoi.Layer2(), bgnoi.Layer2(); got != want {
		t.Errorf("GNOI(t) got Layer2 client %v, want %v", got, want)
	}
	if got, want := gnoi.MPLS(), bgnoi.MPLS(); got != want {
		t.Errorf("GNOI(t) got MPLS client %v, want %v", got, want)
	}
	if got, want := gnoi.OS(), bgnoi.OS(); got != want {
		t.Errorf("GNOI(t) got OS client %v, want %v", got, want)
	}
	if got, want := gnoi.OTDR(), bgnoi.OTDR(); got != want {
		t.Errorf("GNOI(t) got OTDRS client %v, want %v", got, want)
	}
	if got, want := gnoi.System(), bgnoi.System(); got != want {
		t.Errorf("GNOI(t) got System client %v, want %v", got, want)
	}
	if got, want := gnoi.WavelengthRouter(), bgnoi.WavelengthRouter(); got != want {
		t.Errorf("GNOI(t) got WavelengthRouter client %v, want %v", got, want)
	}
	if got, want := gnoi.(*gnoiClients).custom, bgnoi.custom; got != want {
		t.Errorf("GNOI(t) got custom client %v, want %v", got, want)
	}
}

func TestGNOIError(t *testing.T) {
	initDUTFakes(t)
	wantErr := "bad gnoi"
	fakeBind.GNOIDialer = func(context.Context, *binding.DUT, ...grpc.DialOption) (binding.GNOIClients, error) {
		return nil, errors.New(wantErr)
	}
	raw := DUT(t, "dut_cisco").RawAPIs()
	gotErr := negtest.ExpectFatal(t, func(t testing.TB) {
		raw.GNOI().New(t)
	})
	if !strings.Contains(gotErr, wantErr) {
		t.Errorf("GNOI(t) got err %v, want %v", gotErr, wantErr)
	}
}

func TestGRIBI(t *testing.T) {
	initDUTFakes(t)
	tests := []struct {
		desc string
		want struct{ grpb.GRIBIClient }
		f    func(testing.TB) grpb.GRIBIClient
	}{{
		desc: "New GRIBI Client",
		want: struct{ grpb.GRIBIClient }{},
		f:    DUT(t, "dut_cisco").RawAPIs().GRIBI().New,
	}, {
		desc: "Default GRIBI Client",
		want: struct{ grpb.GRIBIClient }{},
		f:    DUT(t, "dut_cisco").RawAPIs().GRIBI().Default,
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			fakeBind.GRIBIDialer = func(context.Context, *binding.DUT, ...grpc.DialOption) (grpb.GRIBIClient, error) {
				return test.want, nil
			}

			if got := test.f(t); got != test.want {
				t.Errorf("GRIBI Client got %v, want %v", got, test.want)
			}
		})
	}
}

func TestGRIBIError(t *testing.T) {
	initDUTFakes(t)
	tests := []struct {
		desc    string
		wantErr string
		f       func(testing.TB) grpb.GRIBIClient
	}{{
		desc:    "New GRIBI Client",
		wantErr: "bad gribi",
		f:       DUT(t, "dut").RawAPIs().GRIBI().New,
	}, {
		desc:    "Default GRIBI Client",
		wantErr: "bad gribi",
		f:       DUT(t, "dut").RawAPIs().GRIBI().Default,
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			fakeBind.GRIBIDialer = func(context.Context, *binding.DUT, ...grpc.DialOption) (grpb.GRIBIClient, error) {
				return nil, errors.New(test.wantErr)
			}

			gotErr := negtest.ExpectFatal(t, func(t testing.TB) {
				test.f(t)
			})

			if !strings.Contains(gotErr, test.wantErr) {
				t.Errorf("GRIBI Client got err %v, want %v", gotErr, test.wantErr)
			}
		})
	}
}

func TestStreamingClient(t *testing.T) {
	initDUTFakes(t)
	fCLI := fakestreamclient.New()
	fConsole := fakestreamclient.New()
	fakeBind.CLIDialer = func(context.Context, *binding.DUT, ...grpc.DialOption) (binding.StreamClient, error) {
		return fCLI, nil
	}
	fakeBind.ConsoleDialer = func(context.Context, *binding.DUT, ...grpc.DialOption) (binding.StreamClient, error) {
		return fConsole, nil
	}
	cliClient := DUT(t, "dut").RawAPIs().CLI(t)
	consoleClient := DUT(t, "dut").RawAPIs().Console(t)
	tests := []struct {
		desc string
		c    StreamClient
		f    *fakestreamclient.FakeStreamClient
	}{{
		desc: "Console",
		c:    consoleClient,
		f:    fConsole,
	}, {
		desc: "CLI",
		c:    cliClient,
		f:    fCLI,
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			want := "show version\n"
			go tt.c.Stdin().Write([]byte(want))
			got, err := bufio.NewReader(tt.f.InReader).ReadString('\n')
			if err != nil {
				t.Fatalf("failed to write to test buffer: %v", err)
			}
			if got != want {
				t.Fatalf("failed to get expect stream data: got %v, want %v", got, want)
			}
			want = "some really cool output\n"
			go tt.f.OutWriter.Write([]byte(want))
			got, err = bufio.NewReader(tt.c.Stdout()).ReadString('\n')
			if err != nil {
				t.Fatalf("failed to write to test buffer: %v", err)
			}
			if got != want {
				t.Fatalf("failed to get expect stream data: got %v, want %v", got, want)
			}
			want = "some errors written to stderr\n"
			go tt.f.ErrWriter.Write([]byte(want))
			got, err = bufio.NewReader(tt.c.Stderr()).ReadString('\n')
			if err != nil {
				t.Fatalf("failed to write to test buffer: %v", err)
			}
			if got != want {
				t.Fatalf("failed to get expect stream data: got %v, want %v", got, want)
			}
		})
	}
}

func TestSendCommand(t *testing.T) {
	initDUTFakes(t)
	fCLI := &fakestreamclient.FakeStreamClient{}
	fConsole := &fakestreamclient.FakeStreamClient{}
	fakeBind.CLIDialer = func(context.Context, *binding.DUT, ...grpc.DialOption) (binding.StreamClient, error) {
		return fCLI, nil
	}
	fakeBind.ConsoleDialer = func(context.Context, *binding.DUT, ...grpc.DialOption) (binding.StreamClient, error) {
		return fConsole, nil
	}
	cliClient := DUT(t, "dut").RawAPIs().CLI(t)
	consoleClient := DUT(t, "dut").RawAPIs().Console(t)
	tests := []struct {
		desc    string
		c       StreamClient
		f       *fakestreamclient.FakeStreamClient
		cmd     string
		wantErr string
	}{{
		desc: "Console",
		c:    consoleClient,
		f:    fConsole,
		cmd:  "some cli command",
	}, {
		desc: "CLI",
		c:    cliClient,
		f:    fCLI,
		cmd:  "some cli command",
	}, {
		desc:    "Console",
		c:       consoleClient,
		f:       fConsole,
		cmd:     "error",
		wantErr: "error",
	}, {
		desc:    "CLI",
		c:       cliClient,
		f:       fCLI,
		cmd:     "error",
		wantErr: "error",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got, err := tt.c.SendCommand(context.Background(), tt.cmd)
			if s := errdiff.Substring(err, tt.wantErr); s != "" {
				t.Fatalf("unexpected error: %s", s)
			}
			if err != nil {
				return
			}
			if got != tt.cmd {
				t.Fatalf("SendCommand failed: got %v, want %v", got, tt.cmd)
			}
		})
	}
}
