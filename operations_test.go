package ondatra

import (
	"bytes"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/testing/protocmp"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/negtest"

	ospb "github.com/openconfig/gnoi/os"
	spb "github.com/openconfig/gnoi/system"
	opb "github.com/openconfig/ondatra/proto"
)

var fakeGNOI *fakeGNOIClient

func init() {
	initFakeBinding()
	reserveFakeTestbed()
	fakeGNOI = &fakeGNOIClient{}
	fakeGNOI.SystemClient = fakeGNOI
	fakeGNOI.OSClient = fakeGNOI
	fakeBind.GNOIDialer = func(context.Context, *reservation.DUT, ...grpc.DialOption) (binding.GNOIClients, error) {
		return fakeGNOI, nil
	}
}

type fakeGNOIClient struct {
	binding.GNOIClients
	spb.SystemClient
	ospb.OSClient
	Pinger         func(context.Context, *spb.PingRequest, ...grpc.CallOption) (spb.System_PingClient, error)
	Rebooter       func(context.Context, *spb.RebootRequest, ...grpc.CallOption) (*spb.RebootResponse, error)
	RebootStatuser func(context.Context, *spb.RebootStatusRequest, ...grpc.CallOption) (*spb.RebootStatusResponse, error)
	Installer      func(context.Context, ...grpc.CallOption) (ospb.OS_InstallClient, error)
}

func (fg *fakeGNOIClient) System() spb.SystemClient {
	return fg
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

func (fg *fakeGNOIClient) Install(ctx context.Context, opts ...grpc.CallOption) (ospb.OS_InstallClient, error) {
	return fg.Installer(ctx, opts...)
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
	const version = "1.2.3"
	dut := DUT(t, "dut")

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
	const version = "1.2.3"
	dut := DUT(t, "dut")

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
			gotErr := negtest.ExpectFatal(t, func(t testing.TB) {
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
	return pc.resp, pc.err
}

func (*fakePingClient) CloseSend() error {
	return nil
}

func TestPing(t *testing.T) {
	var got string
	fakeGNOI.Pinger = func(_ context.Context, req *spb.PingRequest, _ ...grpc.CallOption) (spb.System_PingClient, error) {
		got = req.GetDestination()
		return &fakePingClient{resp: &spb.PingResponse{}}, nil
	}
	want := "1.2.3.4"
	DUT(t, "dut").Operations().NewPing().WithDestination(want).Operate(t)
	if got != want {
		t.Errorf("Operate(t) got %s, want %s", got, want)
	}
}

func TestPingErrors(t *testing.T) {
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
	}}

	for _, tt := range tests {
		t.Run(tt.wantErr, func(t *testing.T) {
			fakeGNOI.Pinger = tt.pinger
			op := DUT(t, "dut").Operations().NewPing().WithDestination(tt.dest)
			gotErr := negtest.ExpectFatal(t, func(t testing.TB) {
				op.Operate(t)
			})
			if !strings.Contains(gotErr, tt.wantErr) {
				t.Errorf("Operate(t) on ping got %q, want %q", gotErr, tt.wantErr)
			}
		})
	}
}

func TestSetInterfaceState(t *testing.T) {
	var gotConfig string
	fakeBind.ConfigPusher = func(_ context.Context, _ *reservation.DUT, config string, _ *binding.ConfigOptions) error {
		gotConfig = config
		return nil
	}
	dut := DUT(t, "dut")
	port := dut.Port(t, "port1")

	tests := []struct {
		desc        string
		op          *SetInterfaceStateOp
		wantIntf    string
		wantEnabled bool
	}{{
		desc: "physical intf enable",
		op: dut.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(port).
			WithStateEnabled(true),
		wantIntf:    "Et1/2/3",
		wantEnabled: true,
	}, {
		desc: "physical intf disable",
		op: dut.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(port).
			WithStateEnabled(false),
		wantIntf:    "Et1/2/3",
		wantEnabled: false,
	}, {
		desc: "logical intf",
		op: dut.Operations().
			NewSetInterfaceState().
			WithLogicalInterface("lintf").
			WithStateEnabled(true),
		wantIntf:    "lintf",
		wantEnabled: true,
	}, {
		desc: "logical intf overwrite",
		op: dut.Operations().
			NewSetInterfaceState().
			WithPhysicalInterface(port).
			WithLogicalInterface("lintf").
			WithStateEnabled(true),
		wantIntf:    "lintf",
		wantEnabled: true,
	}, {
		desc: "state overwrite",
		op: dut.Operations().
			NewSetInterfaceState().
			WithLogicalInterface("lintf").
			WithStateEnabled(false).
			WithStateEnabled(true),
		wantIntf:    "lintf",
		wantEnabled: true,
	}}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			gotConfig = ""
			tt.op.Operate(t)
			if !strings.Contains(gotConfig, tt.wantIntf) {
				t.Errorf("Operate() got config %q, want interface %q", gotConfig, tt.wantIntf)
			}
			if wantEnabled := strconv.FormatBool(tt.wantEnabled); !strings.Contains(gotConfig, wantEnabled) {
				t.Errorf("Operate() got config %q, want enabled %q", gotConfig, wantEnabled)
			}
		})
	}
}

func TestSetInterfaceStateErrors(t *testing.T) {
	dut := DUT(t, "dut")
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
			gotErr := negtest.ExpectFatal(t, func(t testing.TB) {
				tt.op.Operate(t)
			})
			if gotErr == "" {
				t.Errorf("Operate() on op %v succeeded, want error", tt.op)
			}
		})
	}
}

func TestReboot(t *testing.T) {
	reboot := DUT(t, "dut").Operations().NewReboot()

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
	tests := []struct {
		desc                 string
		dut                  *DUTDevice
		timeout              time.Duration
		rebootErr, statusErr error
		wantErr              string
	}{{
		desc:    "non-positive duration",
		dut:     DUT(t, "dut"),
		timeout: -1,
		wantErr: "positive duration",
	}, {
		desc:      "reboot error",
		dut:       DUT(t, "dut"),
		rebootErr: errors.New("reboot error"),
		wantErr:   "gnoi reboot",
	}, {
		desc:      "timeout error",
		dut:       DUT(t, "dut"),
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
			gotErr := negtest.ExpectFatal(t, func(t testing.TB) {
				reboot.Operate(t)
			})
			if !strings.Contains(gotErr, tt.wantErr) {
				t.Errorf("Operate(t) on reboot got %q, want %q", gotErr, tt.wantErr)
			}
		})
	}
}

func TestRestartRouting(t *testing.T) {
	var restarted bool
	fakeBind.RoutingRestarter = func(*reservation.DUT) error {
		restarted = true
		return nil
	}

	dut := DUT(t, "dut_juniper")
	op := dut.Operations().NewRestartRouting()
	op.Operate(t)
	if !restarted {
		t.Fatalf("Operate() on op %v failed, want success", op)
	}
}

func TestRestartRoutingErrors(t *testing.T) {
	fakeBind.RoutingRestarter = func(*reservation.DUT) error {
		return errors.New("bad bad bad :(")
	}

	tests := []struct {
		desc string
		dut  *DUTDevice
		err  error
	}{{
		desc: "bad DUT",
		dut: &DUTDevice{&Device{
			id:  "not in the tb",
			res: &reservation.DUT{&reservation.Dims{Vendor: opb.Device_JUNIPER}},
		}},
	}, {
		desc: "restart fails on good dut",
		dut:  DUT(t, "dut_juniper"),
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			var op *RestartRoutingOp
			gotErr := negtest.ExpectFatal(t, func(t testing.TB) {
				op = tt.dut.Operations().NewRestartRouting()
				op.Operate(t)
			})
			if gotErr == "" {
				t.Errorf("Operate() on op %v succeeded, want error", op)
			}
		})
	}
}
