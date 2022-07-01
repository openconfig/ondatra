// Copyright 2020 Google LLC
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
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/testing/protocmp"

	frpb "github.com/openconfig/gnoi/factory_reset"
	ospb "github.com/openconfig/gnoi/os"
	spb "github.com/openconfig/gnoi/system"
	opb "github.com/openconfig/ondatra/proto"
)

var fakeGNOI = func() *fakeGNOIClient {
	fg := &fakeGNOIClient{}
	fg.SystemClient = fg
	fg.OSClient = fg
	return fg
}()

func initOperationFakes(t *testing.T) {
	t.Helper()
	initFakeBinding(t)
	reserveFakeTestbed(t)
	for _, dut := range fakeRes.DUTs {
		dut.(*fakebind.DUT).DialGNOIFn = func(context.Context, ...grpc.DialOption) (binding.GNOIClients, error) {
			return fakeGNOI, nil
		}
	}
}

type fakeGNOIClient struct {
	binding.GNOIClients
	spb.SystemClient
	ospb.OSClient
	Pinger           func(context.Context, *spb.PingRequest, ...grpc.CallOption) (spb.System_PingClient, error)
	Rebooter         func(context.Context, *spb.RebootRequest, ...grpc.CallOption) (*spb.RebootResponse, error)
	RebootStatuser   func(context.Context, *spb.RebootStatusRequest, ...grpc.CallOption) (*spb.RebootStatusResponse, error)
	KillProcessor    func(context.Context, *spb.KillProcessRequest, ...grpc.CallOption) (*spb.KillProcessResponse, error)
	Installer        func(context.Context, ...grpc.CallOption) (ospb.OS_InstallClient, error)
	SystemTime       func(context.Context, *spb.TimeRequest, ...grpc.CallOption) (*spb.TimeResponse, error)
	SwitchController func(ctx context.Context, in *spb.SwitchControlProcessorRequest, opts ...grpc.CallOption) (*spb.SwitchControlProcessorResponse, error)
}

func (fg *fakeGNOIClient) System() spb.SystemClient {
	return fg
}

func (fg *fakeGNOIClient) FactoryReset() frpb.FactoryResetClient {
	return fakeFactoryResetClient{}
}

func (fg *fakeGNOIClient) SwitchControlProcessor(ctx context.Context, in *spb.SwitchControlProcessorRequest, opts ...grpc.CallOption) (*spb.SwitchControlProcessorResponse, error) {
	return fg.SwitchController(ctx, in, opts...)
}

func (fg *fakeGNOIClient) OS() ospb.OSClient {
	return fg
}

func (fg *fakeGNOIClient) Ping(ctx context.Context, in *spb.PingRequest, opts ...grpc.CallOption) (spb.System_PingClient, error) {
	return fg.Pinger(ctx, in, opts...)
}

func (fg *fakeGNOIClient) Reboot(ctx context.Context, req *spb.RebootRequest, opts ...grpc.CallOption) (*spb.RebootResponse, error) {
	return fg.Rebooter(ctx, req, opts...)
}

func (fg *fakeGNOIClient) RebootStatus(ctx context.Context, req *spb.RebootStatusRequest, opts ...grpc.CallOption) (*spb.RebootStatusResponse, error) {
	return fg.RebootStatuser(ctx, req, opts...)
}

func (fg *fakeGNOIClient) KillProcess(ctx context.Context, req *spb.KillProcessRequest, opts ...grpc.CallOption) (*spb.KillProcessResponse, error) {
	return fg.KillProcessor(ctx, req, opts...)
}

func (fg *fakeGNOIClient) Install(ctx context.Context, opts ...grpc.CallOption) (ospb.OS_InstallClient, error) {
	return fg.Installer(ctx, opts...)
}

func (fg *fakeGNOIClient) Time(ctx context.Context, req *spb.TimeRequest, opts ...grpc.CallOption) (*spb.TimeResponse, error) {
	return fg.SystemTime(ctx, req, opts...)
}

type fakeFactoryResetClient struct {
	frpb.FactoryResetClient
}

func (ic fakeFactoryResetClient) Start(ctx context.Context, req *frpb.StartRequest, opts ...grpc.CallOption) (*frpb.StartResponse, error) {
	resp := &frpb.StartResponse{Response: &frpb.StartResponse_ResetSuccess{
		ResetSuccess: &frpb.ResetSuccess{},
	}}
	return resp, nil
}

type fakeInstallClient struct {
	ospb.OS_InstallClient
	gotSent  []*ospb.InstallRequest
	stubRecv []*ospb.InstallResponse
}

func (ic *fakeInstallClient) Send(req *ospb.InstallRequest) error {
	ic.gotSent = append(ic.gotSent, req)
	return nil
}

func (ic *fakeInstallClient) Recv() (*ospb.InstallResponse, error) {
	if len(ic.stubRecv) == 0 {
		return nil, errors.New("no more stub responses")
	}
	resp := ic.stubRecv[0]
	ic.stubRecv[0] = nil
	ic.stubRecv = ic.stubRecv[1:]
	return resp, nil
}

func (*fakeInstallClient) CloseSend() error {
	return nil
}

func TestInstall(t *testing.T) {
	initOperationFakes(t)
	const version = "1.2.3"
	dut := DUT(t, "dut_arista")

	// Make a temp file to test specifying a file by file path.
	file, err := ioutil.TempFile("", "package")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()
	if err := ioutil.WriteFile(file.Name(), []byte{0}, os.ModePerm); err != nil {
		t.Fatalf("error writing temp file: %v", err)
	}

	tests := []struct {
		desc     string
		op       *InstallOp
		resps    []*ospb.InstallResponse
		wantSent []*ospb.InstallRequest
	}{{
		desc: "target has package",
		op:   dut.Operations().NewInstall().WithVersion(version),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_Validated{&ospb.Validated{Version: version}}},
		},
		wantSent: []*ospb.InstallRequest{
			{Request: &ospb.InstallRequest_TransferRequest{&ospb.TransferRequest{Version: version}}},
		},
	}, {
		desc: "standby supervisor",
		op:   dut.Operations().NewInstall().WithVersion(version).WithStandbySupervisor(true),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_Validated{&ospb.Validated{Version: version}}},
		},
		wantSent: []*ospb.InstallRequest{
			{Request: &ospb.InstallRequest_TransferRequest{&ospb.TransferRequest{Version: version, StandbySupervisor: true}}},
		},
	}, {
		desc: "supervisor has package",
		op:   dut.Operations().NewInstall().WithVersion(version),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_SyncProgress{&ospb.SyncProgress{}}},
			{Response: &ospb.InstallResponse_SyncProgress{&ospb.SyncProgress{}}},
			{Response: &ospb.InstallResponse_Validated{&ospb.Validated{Version: version}}},
		},
		wantSent: []*ospb.InstallRequest{
			{Request: &ospb.InstallRequest_TransferRequest{&ospb.TransferRequest{Version: version}}},
		},
	}, {
		desc: "transfer content",
		op:   dut.Operations().NewInstall().WithVersion(version).WithPackageReader(bytes.NewReader([]byte{0})),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_TransferReady{&ospb.TransferReady{}}},
			{Response: &ospb.InstallResponse_TransferProgress{&ospb.TransferProgress{}}},
			{Response: &ospb.InstallResponse_TransferProgress{&ospb.TransferProgress{}}},
			{Response: &ospb.InstallResponse_Validated{&ospb.Validated{Version: version}}},
		},
		wantSent: []*ospb.InstallRequest{
			{Request: &ospb.InstallRequest_TransferRequest{&ospb.TransferRequest{Version: version}}},
			{Request: &ospb.InstallRequest_TransferContent{[]byte{0}}},
			{Request: &ospb.InstallRequest_TransferEnd{&ospb.TransferEnd{}}},
		},
	}, {
		desc: "transfer content by file",
		op:   dut.Operations().NewInstall().WithVersion(version).WithPackageFile(file.Name()),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_TransferReady{&ospb.TransferReady{}}},
			{Response: &ospb.InstallResponse_TransferProgress{&ospb.TransferProgress{}}},
			{Response: &ospb.InstallResponse_TransferProgress{&ospb.TransferProgress{}}},
			{Response: &ospb.InstallResponse_Validated{&ospb.Validated{Version: version}}},
		},
		wantSent: []*ospb.InstallRequest{
			{Request: &ospb.InstallRequest_TransferRequest{&ospb.TransferRequest{Version: version}}},
			{Request: &ospb.InstallRequest_TransferContent{[]byte{0}}},
			{Request: &ospb.InstallRequest_TransferEnd{&ospb.TransferEnd{}}},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ic := &fakeInstallClient{stubRecv: tt.resps}
			fakeGNOI.Installer = func(context.Context, ...grpc.CallOption) (ospb.OS_InstallClient, error) {
				return ic, nil
			}
			tt.op.Operate(t)
			if diff := cmp.Diff(tt.wantSent, ic.gotSent, protocmp.Transform()); diff != "" {
				t.Errorf("Operate(t) unexpected sent requests diff (-want,+got): %s", diff)
			}
			if len(ic.stubRecv) > 0 {
				t.Errorf("Operate(t) has unused stub responses: %v", ic.stubRecv)
			}
		})
	}
}

func TestInstallErrors(t *testing.T) {
	initOperationFakes(t)
	const version = "1.2.3"
	dut := DUT(t, "dut_cisco")

	tests := []struct {
		op      *InstallOp
		resps   []*ospb.InstallResponse
		wantErr string
	}{{
		op:      dut.Operations().NewInstall(),
		wantErr: "version not set",
	}, {
		op: dut.Operations().NewInstall().WithVersion(version),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_InstallError{&ospb.InstallError{Detail: "install error on initial"}}},
		},
		wantErr: "install error on initial",
	}, {
		op: dut.Operations().NewInstall().WithVersion(version),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_TransferReady{&ospb.TransferReady{}}},
		},
		wantErr: "no package specified",
	}, {
		op: dut.Operations().NewInstall().WithVersion(version).WithPackageReader(bytes.NewReader([]byte{0})),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_TransferReady{&ospb.TransferReady{}}},
			{Response: &ospb.InstallResponse_InstallError{&ospb.InstallError{Detail: "install error on transfer"}}},
		},
		wantErr: "install error on transfer",
	}, {
		op: dut.Operations().NewInstall().WithVersion(version),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_Validated{Validated: &ospb.Validated{Version: version + "a"}}},
		},
		wantErr: "does not match requested version",
	}}

	for _, tt := range tests {
		t.Run(tt.wantErr, func(t *testing.T) {
			ic := &fakeInstallClient{stubRecv: tt.resps}
			fakeGNOI.Installer = func(context.Context, ...grpc.CallOption) (ospb.OS_InstallClient, error) {
				return ic, nil
			}
			gotErr := testt.ExpectFatal(t, func(t testing.TB) {
				tt.op.Operate(t)
			})
			if !strings.Contains(gotErr, tt.wantErr) {
				t.Errorf("Operate(t) got %q, want %q", gotErr, tt.wantErr)
			}
			if len(ic.stubRecv) > 0 {
				t.Errorf("Operate(t) has unused stub responses: %v", ic.stubRecv)
			}
		})
	}
}

type fakePingClient struct {
	spb.System_PingClient
	resp *spb.PingResponse
	err  error
}

func (pc *fakePingClient) Recv() (*spb.PingResponse, error) {
	if pc.resp == nil && pc.err == nil {
		return nil, io.EOF
	}
	resp := pc.resp
	pc.resp = nil
	return resp, pc.err
}

func (*fakePingClient) CloseSend() error {
	return nil
}

func TestPing(t *testing.T) {
	initOperationFakes(t)
	tests := []struct {
		desc, dest string
		count      int32
	}{
		{desc: "zero count", dest: "1.2.3.4"},
		{desc: "non-zero count", dest: "1.2.3.4", count: 7},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			var got string
			fakeGNOI.Pinger = func(_ context.Context, req *spb.PingRequest, _ ...grpc.CallOption) (spb.System_PingClient, error) {
				got = req.GetDestination()
				return &fakePingClient{resp: &spb.PingResponse{Sent: tt.count, Received: tt.count}}, nil
			}
			want := tt.dest
			DUT(t, "dut_juniper").Operations().NewPing().WithDestination(tt.dest).WithCount(tt.count).Operate(t)
			if got != want {
				t.Errorf("Operate(t) got %s, want %s", got, want)
			}

		})
	}

}

func TestPingErrors(t *testing.T) {
	initOperationFakes(t)
	tests := []struct {
		wantErr, dest string
		pinger        func(context.Context, *spb.PingRequest, ...grpc.CallOption) (spb.System_PingClient, error)
	}{{
		wantErr: "no destination",
	}, {
		wantErr: "ping error",
		dest:    "1.2.3.4",
		pinger: func(_ context.Context, req *spb.PingRequest, _ ...grpc.CallOption) (spb.System_PingClient, error) {
			return nil, errors.New("ping error")
		},
	}, {
		wantErr: "recv error",
		dest:    "1.2.3.4",
		pinger: func(_ context.Context, req *spb.PingRequest, _ ...grpc.CallOption) (spb.System_PingClient, error) {
			return &fakePingClient{
				err: errors.New("recv error"),
			}, nil
		},
	}, {
		wantErr: "ping sent 5 packets, received 3",
		dest:    "1.2.3.4",
		pinger: func(_ context.Context, req *spb.PingRequest, _ ...grpc.CallOption) (spb.System_PingClient, error) {
			return &fakePingClient{
				resp: &spb.PingResponse{Sent: 5, Received: 3},
			}, nil
		},
	},
	}

	for _, tt := range tests {
		t.Run(tt.wantErr, func(t *testing.T) {
			fakeGNOI.Pinger = tt.pinger
			op := DUT(t, "dut_cisco").Operations().NewPing().WithDestination(tt.dest)
			gotErr := testt.ExpectFatal(t, func(t testing.TB) {
				op.Operate(t)
			})
			if !strings.Contains(gotErr, tt.wantErr) {
				t.Errorf("Operate(t) on ping got %q, want %q", gotErr, tt.wantErr)
			}
		})
	}
}

func TestSetInterfaceState(t *testing.T) {
	initOperationFakes(t)
	var gotConfig string
	for _, dut := range fakeRes.DUTs {
		dut.(*fakebind.DUT).PushConfigFn = func(_ context.Context, config string, _ bool) error {
			gotConfig = config
			return nil
		}
	}

	dutArista := DUT(t, "dut_arista")
	dutCisco := DUT(t, "dut_cisco")
	dutJuniper := DUT(t, "dut_juniper")
	portArista := dutArista.Port(t, "port1")
	portCisco := dutCisco.Port(t, "port1")
	portJuniper := dutJuniper.Port(t, "port1")

	tests := []struct {
		desc       string
		op         *SetInterfaceStateOp
		wantIntf   string
		wantConfig string
	}{{
		desc: "Arista physical intf enable",
		op: dutArista.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(portArista).
			WithStateEnabled(true),
		wantIntf:   "Et1/2/3",
		wantConfig: "no shutdown",
	}, {
		desc: "Cisco physical intf enable",
		op: dutCisco.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(portCisco).
			WithStateEnabled(true),
		wantIntf:   "Et1/2/3",
		wantConfig: "no shutdown",
	}, {
		desc: "Juniper physical intf enable",
		op: dutJuniper.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(portJuniper).
			WithStateEnabled(true),
		wantIntf:   "Et1/2/3",
		wantConfig: "delete: disable",
	}, {
		desc: "Arista physical intf disable",
		op: dutArista.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(portArista).
			WithStateEnabled(false),
		wantIntf:   "Et1/2/3",
		wantConfig: "  shutdown",
	}, {
		desc: "Cisco physical intf disable",
		op: dutCisco.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(portCisco).
			WithStateEnabled(false),
		wantIntf:   "Et1/2/3",
		wantConfig: "  shutdown",
	}, {
		desc: "Juniper physical intf disable",
		op: dutJuniper.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(portJuniper).
			WithStateEnabled(false),
		wantIntf:   "Et1/2/3",
		wantConfig: "  disable",
	}, {
		desc: "logical intf enable",
		op: dutArista.Operations().
			NewSetInterfaceState().
			WithLogicalInterface("lintf").
			WithStateEnabled(true),
		wantIntf:   "lintf",
		wantConfig: "no shutdown",
	}, {
		desc: "logical intf overwrite",
		op: dutArista.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(portArista).
			WithLogicalInterface("lintf").
			WithStateEnabled(true),
		wantIntf:   "lintf",
		wantConfig: "no shutdown",
	}, {
		desc: "state overwrite",
		op: dutArista.Operations().
			NewSetInterfaceState().
			WithLogicalInterface("lintf").
			WithStateEnabled(false).
			WithStateEnabled(true),
		wantIntf:   "lintf",
		wantConfig: "no shutdown",
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotConfig = ""
			tt.op.Operate(t)
			if !strings.Contains(gotConfig, tt.wantIntf) {
				t.Errorf("Operate() got config %q, want interface %q", gotConfig, tt.wantIntf)
			}
			if !strings.Contains(gotConfig, tt.wantConfig) {
				t.Errorf("Operate() got config %q, want %q", gotConfig, tt.wantConfig)
			}
		})
	}
}

func TestSetInterfaceStateErrors(t *testing.T) {
	initOperationFakes(t)
	dut := DUT(t, "dut_juniper")
	port := dut.Port(t, "port1")

	tests := []struct {
		desc string
		op   *SetInterfaceStateOp
	}{{
		desc: "no interface",
		op: dut.Operations().
			NewSetInterfaceState().
			WithStateEnabled(true),
	}, {
		desc: "no state",
		op: dut.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(port),
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotErr := testt.ExpectFatal(t, func(t testing.TB) {
				tt.op.Operate(t)
			})
			if gotErr == "" {
				t.Errorf("Operate() on op %v succeeded, want error", tt.op)
			}
		})
	}
}

func TestReboot(t *testing.T) {
	initOperationFakes(t)
	reboot := DUT(t, "dut_arista").Operations().NewReboot()

	tests := []struct {
		desc       string
		statuses   []*spb.RebootStatusResponse
		statusErrs []error
	}{{
		desc:       "reboot immediately",
		statuses:   []*spb.RebootStatusResponse{{}},
		statusErrs: []error{nil},
	}, {
		desc:       "reboot delayed",
		statuses:   []*spb.RebootStatusResponse{{Active: true}, {}, {}},
		statusErrs: []error{nil, errors.New("cannot reach device"), nil},
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			var rebooted bool
			fakeGNOI.Rebooter = func(context.Context, *spb.RebootRequest, ...grpc.CallOption) (*spb.RebootResponse, error) {
				rebooted = true
				return &spb.RebootResponse{}, nil
			}
			var index uint8
			fakeGNOI.RebootStatuser = func(context.Context, *spb.RebootStatusRequest, ...grpc.CallOption) (*spb.RebootStatusResponse, error) {
				status, err := tt.statuses[index], tt.statusErrs[index]
				index = index + 1
				return status, err
			}
			reboot.Operate(t)
			if !rebooted {
				t.Fatal("Operate() on reboot failed, want success")
			}
		})
	}
}

func TestRebootErrors(t *testing.T) {
	initOperationFakes(t)
	tests := []struct {
		desc                 string
		dut                  *DUTDevice
		timeout              time.Duration
		rebootErr, statusErr error
		wantErr              string
	}{{
		desc:    "non-positive duration",
		dut:     DUT(t, "dut_arista"),
		timeout: -1,
		wantErr: "positive duration",
	}, {
		desc:      "reboot error",
		dut:       DUT(t, "dut_cisco"),
		rebootErr: errors.New("reboot error"),
		wantErr:   "gnoi reboot",
	}, {
		desc:      "timeout error",
		dut:       DUT(t, "dut_juniper"),
		timeout:   1,
		statusErr: errors.New("status error"),
		wantErr:   "timed out",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			fakeGNOI.Rebooter = func(context.Context, *spb.RebootRequest, ...grpc.CallOption) (*spb.RebootResponse, error) {
				return &spb.RebootResponse{}, tt.rebootErr
			}
			fakeGNOI.RebootStatuser = func(context.Context, *spb.RebootStatusRequest, ...grpc.CallOption) (*spb.RebootStatusResponse, error) {
				return &spb.RebootStatusResponse{}, tt.statusErr
			}
			reboot := tt.dut.Operations().NewReboot().WithTimeout(tt.timeout)
			gotErr := testt.ExpectFatal(t, func(t testing.TB) {
				reboot.Operate(t)
			})
			if !strings.Contains(gotErr, tt.wantErr) {
				t.Errorf("Operate(t) on reboot got %q, want %q", gotErr, tt.wantErr)
			}
		})
	}
}

func TestKillProcess(t *testing.T) {
	initOperationFakes(t)
	var killed bool
	fakeGNOI.KillProcessor = func(context.Context, *spb.KillProcessRequest, ...grpc.CallOption) (*spb.KillProcessResponse, error) {
		killed = true
		return &spb.KillProcessResponse{}, nil
	}

	dut := DUT(t, "dut_juniper")
	op := dut.Operations().NewKillProcess().WithPID(123)
	op.Operate(t)
	if !killed {
		t.Fatalf("Operate() on op %v failed, want success", op)
	}
}

func TestSystemTime(t *testing.T) {
	initOperationFakes(t)
	fakeGNOI.SystemTime = func(context.Context, *spb.TimeRequest, ...grpc.CallOption) (*spb.TimeResponse, error) {
		return &spb.TimeResponse{
			Time: 12345,
		}, nil
	}
	dut := DUT(t, "dut_cisco")
	got := dut.Operations().SystemTime(t)
	if want := time.Unix(0, 12345); want != got {
		t.Fatalf("Operation failed, want %v got %v", want, got)
	}
}

func TestFactoryReset(t *testing.T) {
	initOperationFakes(t)
	dut := DUT(t, "dut_cisco")
	dut.Operations().NewFactoryReset().WithZeroFill(true).WithFactoryOS(true).Operate(t)
}

func TestKillProcessErrors(t *testing.T) {
	initOperationFakes(t)
	fakeGNOI.KillProcessor = func(context.Context, *spb.KillProcessRequest, ...grpc.CallOption) (*spb.KillProcessResponse, error) {
		return nil, errors.New("bad bad bad :(")
	}

	tests := []struct {
		desc string
		dut  *DUTDevice
		err  error
	}{{
		desc: "bad DUT",
		dut: &DUTDevice{&Device{
			id:  "not in the tb",
			res: &binding.AbstractDUT{&binding.Dims{Vendor: opb.Device_JUNIPER}},
		}},
	}, {
		desc: "kill fails on good dut",
		dut:  DUT(t, "dut_juniper"),
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			var op *KillProcessOp
			gotErr := testt.ExpectFatal(t, func(t testing.TB) {
				op = tt.dut.Operations().NewKillProcess()
				op.Operate(t)
			})
			if gotErr == "" {
				t.Errorf("Operate() on op %v succeeded, want error", op)
			}
		})
	}
}

func TestSwitchControlProcessor(t *testing.T) {
	initOperationFakes(t)
	var got string
	fakeGNOI.SwitchController = func(ctx context.Context, in *spb.SwitchControlProcessorRequest, opts ...grpc.CallOption) (*spb.SwitchControlProcessorResponse, error) {
		got = in.GetControlProcessor().GetElem()[1].GetKey()["name"]
		return &spb.SwitchControlProcessorResponse{ControlProcessor: in.ControlProcessor}, nil
	}
	ciscoDUT := DUT(t, "dut_cisco")

	tests := []struct {
		desc string
		op   *SwitchControlProcessorOp
		want string
	}{{
		desc: "cisco switch control processor",
		op: ciscoDUT.Operations().
			NewSwitchControlProcessor().
			WithDestination("RP0"),
		want: "RP0",
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got = ""
			tt.op.Operate(t)
			if !strings.Contains(got, tt.want) {
				t.Errorf("Operate(t) on switch control process got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSwitchControlProcessorErrors(t *testing.T) {
	initOperationFakes(t)
	sw := DUT(t, "dut_arista").Operations().NewSwitchControlProcessor()
	fakeGNOI.SwitchController = func(context.Context, *spb.SwitchControlProcessorRequest, ...grpc.CallOption) (*spb.SwitchControlProcessorResponse, error) {
		return nil, errors.New("invalid route controller")
	}
	op := sw.WithDestination("RP0")
	gotErr := testt.ExpectFatal(t, func(t testing.TB) {
		op.Operate(t)
	})
	if gotErr == "" {
		t.Errorf("Operate() on op %v succeeded, want error", op)
	}
}
