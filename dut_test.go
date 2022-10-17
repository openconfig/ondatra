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
	"errors"
	"strings"
	"testing"

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
	plqpb "github.com/openconfig/gnoi/packet_link_qualification"
	spb "github.com/openconfig/gnoi/system"
	wpb "github.com/openconfig/gnoi/wavelength_router"
	grpb "github.com/openconfig/gribi/v1/proto/service"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"

)

var (
	gotConfig string
	gotReset  bool
)

func initDUTFakes(t *testing.T) {
	t.Helper()
	initFakeBinding(t)
	reserveFakeTestbed(t)
	fakeArista.PushConfigFn = func(_ context.Context, config string, reset bool) error {
		gotConfig = config
		gotReset = reset
		return nil
	}
}

func TestGNMI(t *testing.T) {
	initDUTFakes(t)
	want := struct{ gpb.GNMIClient }{}
	fakeCisco.DialGNMIFn = func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
		return want, nil
	}
	if got := DUT(t, "dut_cisco").RawAPIs().GNMI().New(t); got != want {
		t.Errorf("GNMI().New(t) got %v, want %v", got, want)
	}
}

func TestGNMIError(t *testing.T) {
	initDUTFakes(t)
	wantErr := "bad gnmi"
	fakeJuniper.DialGNMIFn = func(context.Context, ...grpc.DialOption) (gpb.GNMIClient, error) {
		return nil, errors.New(wantErr)
	}
	raw := DUT(t, "dut_juniper").RawAPIs()
	gotErr := testt.ExpectFatal(t, func(t testing.TB) {
		raw.GNMI().New(t)
	})
	if !strings.Contains(gotErr, wantErr) {
		t.Errorf("GNMI().New(t) got err %v, want %v", gotErr, wantErr)
	}
}

type gnoiClients struct {
	binding.GNOIClients
	bgp               bpb.BGPClient
	certMgmt          cpb.CertificateManagementClient
	diag              dpb.DiagClient
	factoryReset      frpb.FactoryResetClient
	file              fpb.FileClient
	healthz           hpb.HealthzClient
	intface           ipb.InterfaceClient
	layer2            lpb.Layer2Client
	linkQualification plqpb.LinkQualificationClient
	mpls              mpb.MPLSClient
	os                ospb.OSClient
	otdr              otpb.OTDRClient
	system            spb.SystemClient
	waveRtr           wpb.WavelengthRouterClient
	custom            any
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

func (g *gnoiClients) LinkQualification() plqpb.LinkQualificationClient {
	return g.linkQualification
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
		diag:              struct{ dpb.DiagClient }{},
		factoryReset:      struct{ frpb.FactoryResetClient }{},
		file:              struct{ fpb.FileClient }{},
		healthz:           struct{ hpb.HealthzClient }{},
		intface:           struct{ ipb.InterfaceClient }{},
		layer2:            struct{ lpb.Layer2Client }{},
		linkQualification: struct{ plqpb.LinkQualificationClient }{},
		mpls:              struct{ mpb.MPLSClient }{},
		os:                struct{ ospb.OSClient }{},
		otdr:              struct{ otpb.OTDRClient }{},
		system:            struct{ spb.SystemClient }{},
		waveRtr:           struct{ wpb.WavelengthRouterClient }{},
		custom:            struct{}{},
	}
	fakeJuniper.DialGNOIFn = func(context.Context, ...grpc.DialOption) (binding.GNOIClients, error) {
		return bgnoi, nil
	}
	gnoi := DUT(t, "dut_juniper").RawAPIs().GNOI().New(t)
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
	if got, want := gnoi.LinkQualification(), bgnoi.LinkQualification(); got != want {
		t.Errorf("GNOI(t) got LinkQualification client %v, want %v", got, want)
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
	fakeCisco.DialGNOIFn = func(context.Context, ...grpc.DialOption) (binding.GNOIClients, error) {
		return nil, errors.New(wantErr)
	}
	raw := DUT(t, "dut_cisco").RawAPIs()
	gotErr := testt.ExpectFatal(t, func(t testing.TB) {
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
			fakeCisco.DialGRIBIFn = func(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error) {
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
		f:       DUT(t, "dut_arista").RawAPIs().GRIBI().New,
	}, {
		desc:    "Default GRIBI Client",
		wantErr: "bad gribi",
		f:       DUT(t, "dut_arista").RawAPIs().GRIBI().Default,
	}}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			fakeArista.DialGRIBIFn = func(context.Context, ...grpc.DialOption) (grpb.GRIBIClient, error) {
				return nil, errors.New(test.wantErr)
			}

			gotErr := testt.ExpectFatal(t, func(t testing.TB) {
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
	fCLI := fakebind.NewStreamClient()
	fConsole := fakebind.NewStreamClient()
	fakeJuniper.DialCLIFn = func(context.Context, ...grpc.DialOption) (binding.StreamClient, error) {
		return fCLI, nil
	}
	fakeJuniper.DialConsoleFn = func(context.Context, ...grpc.DialOption) (binding.StreamClient, error) {
		return fConsole, nil
	}
	cliClient := DUT(t, "dut_juniper").RawAPIs().CLI(t)
	consoleClient := DUT(t, "dut_juniper").RawAPIs().Console(t)
	tests := []struct {
		desc string
		c    StreamClient
		f    *fakebind.StreamClient
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
	fCLI := &fakebind.StreamClient{}
	fConsole := &fakebind.StreamClient{}
	fakeArista.DialCLIFn = func(context.Context, ...grpc.DialOption) (binding.StreamClient, error) {
		return fCLI, nil
	}
	fakeArista.DialConsoleFn = func(context.Context, ...grpc.DialOption) (binding.StreamClient, error) {
		return fConsole, nil
	}
	cliClient := DUT(t, "dut_arista").RawAPIs().CLI(t)
	consoleClient := DUT(t, "dut_arista").RawAPIs().Console(t)
	tests := []struct {
		desc    string
		c       StreamClient
		f       *fakebind.StreamClient
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
