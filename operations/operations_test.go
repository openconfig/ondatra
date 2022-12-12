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

package operations

import (
	"bytes"
	"errors"
	"io"
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
)

var (
	fakeGNOI = new(fakeGNOIClient)
	dut      = &fakebind.DUT{
		AbstractDUT: &binding.AbstractDUT{&binding.Dims{Name: "fakeDUT"}},
		DialGNOIFn: func(context.Context, ...grpc.DialOption) (binding.GNOIClients, error) {
			return fakeGNOI, nil
		},
	}
)

type fakeGNOIClient struct {
	binding.GNOIClients
	spb.SystemClient
	ospb.OSClient
	frpb.FactoryResetClient
	PingFn          func(context.Context, *spb.PingRequest, ...grpc.CallOption) (spb.System_PingClient, error)
	RebootFn        func(context.Context, *spb.RebootRequest, ...grpc.CallOption) (*spb.RebootResponse, error)
	RebootStatusFn  func(context.Context, *spb.RebootStatusRequest, ...grpc.CallOption) (*spb.RebootStatusResponse, error)
	KillProcessFn   func(context.Context, *spb.KillProcessRequest, ...grpc.CallOption) (*spb.KillProcessResponse, error)
	InstallFn       func(context.Context, ...grpc.CallOption) (ospb.OS_InstallClient, error)
	StartFn         func(context.Context, *frpb.StartRequest, ...grpc.CallOption) (*frpb.StartResponse, error)
	TimeFn          func(context.Context, *spb.TimeRequest, ...grpc.CallOption) (*spb.TimeResponse, error)
	SwitchControlFn func(ctx context.Context, in *spb.SwitchControlProcessorRequest, opts ...grpc.CallOption) (*spb.SwitchControlProcessorResponse, error)
}

func (fg *fakeGNOIClient) System() spb.SystemClient {
	return fg
}

func (fg *fakeGNOIClient) FactoryReset() frpb.FactoryResetClient {
	return fg
}

func (fg *fakeGNOIClient) Start(ctx context.Context, req *frpb.StartRequest, opts ...grpc.CallOption) (*frpb.StartResponse, error) {
	return fg.StartFn(ctx, req, opts...)
}

func (fg *fakeGNOIClient) SwitchControlProcessor(ctx context.Context, in *spb.SwitchControlProcessorRequest, opts ...grpc.CallOption) (*spb.SwitchControlProcessorResponse, error) {
	return fg.SwitchControlFn(ctx, in, opts...)
}

func (fg *fakeGNOIClient) OS() ospb.OSClient {
	return fg
}

func (fg *fakeGNOIClient) Ping(ctx context.Context, in *spb.PingRequest, opts ...grpc.CallOption) (spb.System_PingClient, error) {
	return fg.PingFn(ctx, in, opts...)
}

func (fg *fakeGNOIClient) Reboot(ctx context.Context, req *spb.RebootRequest, opts ...grpc.CallOption) (*spb.RebootResponse, error) {
	return fg.RebootFn(ctx, req, opts...)
}

func (fg *fakeGNOIClient) RebootStatus(ctx context.Context, req *spb.RebootStatusRequest, opts ...grpc.CallOption) (*spb.RebootStatusResponse, error) {
	return fg.RebootStatusFn(ctx, req, opts...)
}

func (fg *fakeGNOIClient) KillProcess(ctx context.Context, req *spb.KillProcessRequest, opts ...grpc.CallOption) (*spb.KillProcessResponse, error) {
	return fg.KillProcessFn(ctx, req, opts...)
}

func (fg *fakeGNOIClient) Install(ctx context.Context, opts ...grpc.CallOption) (ospb.OS_InstallClient, error) {
	return fg.InstallFn(ctx, opts...)
}

func (fg *fakeGNOIClient) Time(ctx context.Context, req *spb.TimeRequest, opts ...grpc.CallOption) (*spb.TimeResponse, error) {
	return fg.TimeFn(ctx, req, opts...)
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

	// Make a temp file to test specifying a file by file path.
	file, err := os.CreateTemp("", "package")
	if err != nil {
		t.Fatalf("error creating temp file: %v", err)
	}
	defer os.Remove(file.Name())
	defer file.Close()
	if err := os.WriteFile(file.Name(), []byte{0}, os.ModePerm); err != nil {
		t.Fatalf("error writing temp file: %v", err)
	}

	tests := []struct {
		desc     string
		op       *InstallOp
		resps    []*ospb.InstallResponse
		wantSent []*ospb.InstallRequest
	}{{
		desc: "target has package",
		op:   New(dut).NewInstall().WithVersion(version),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_Validated{&ospb.Validated{Version: version}}},
		},
		wantSent: []*ospb.InstallRequest{
			{Request: &ospb.InstallRequest_TransferRequest{&ospb.TransferRequest{Version: version}}},
		},
	}, {
		desc: "standby supervisor",
		op:   New(dut).NewInstall().WithVersion(version).WithStandbySupervisor(true),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_Validated{&ospb.Validated{Version: version}}},
		},
		wantSent: []*ospb.InstallRequest{
			{Request: &ospb.InstallRequest_TransferRequest{&ospb.TransferRequest{Version: version, StandbySupervisor: true}}},
		},
	}, {
		desc: "supervisor has package",
		op:   New(dut).NewInstall().WithVersion(version),
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
		op:   New(dut).NewInstall().WithVersion(version).WithPackageReader(bytes.NewReader([]byte{0})),
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
		op:   New(dut).NewInstall().WithVersion(version).WithPackageFile(file.Name()),
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
			fakeGNOI.InstallFn = func(context.Context, ...grpc.CallOption) (ospb.OS_InstallClient, error) {
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

	tests := []struct {
		op      *InstallOp
		resps   []*ospb.InstallResponse
		wantErr string
	}{{
		op:      New(dut).NewInstall(),
		wantErr: "version not set",
	}, {
		op: New(dut).NewInstall().WithVersion(version),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_InstallError{&ospb.InstallError{Detail: "install error on initial"}}},
		},
		wantErr: "install error on initial",
	}, {
		op: New(dut).NewInstall().WithVersion(version),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_TransferReady{&ospb.TransferReady{}}},
		},
		wantErr: "no package specified",
	}, {
		op: New(dut).NewInstall().WithVersion(version).WithPackageReader(bytes.NewReader([]byte{0})),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_TransferReady{&ospb.TransferReady{}}},
			{Response: &ospb.InstallResponse_InstallError{&ospb.InstallError{Detail: "install error on transfer"}}},
		},
		wantErr: "install error on transfer",
	}, {
		op: New(dut).NewInstall().WithVersion(version),
		resps: []*ospb.InstallResponse{
			{Response: &ospb.InstallResponse_Validated{Validated: &ospb.Validated{Version: version + "a"}}},
		},
		wantErr: "does not match requested version",
	}}
	for _, tt := range tests {
		t.Run(tt.wantErr, func(t *testing.T) {
			ic := &fakeInstallClient{stubRecv: tt.resps}
			fakeGNOI.InstallFn = func(context.Context, ...grpc.CallOption) (ospb.OS_InstallClient, error) {
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
	tests := []struct {
		desc, dest        string
		count, packetSize int32
	}{
		{desc: "zero count, zero packet size", dest: "1.2.3.4"},
		{desc: "non-zero count, zero packet size", dest: "1.2.3.4", count: 7},
		{desc: "zero count, non-zero packet size", dest: "1.2.3.4", packetSize: 1000},
		{desc: "non-zero count, non-zero packet size", dest: "1.2.3.4", count: 7, packetSize: 1000},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			var got string
			fakeGNOI.PingFn = func(_ context.Context, req *spb.PingRequest, _ ...grpc.CallOption) (spb.System_PingClient, error) {
				got = req.GetDestination()
				return &fakePingClient{resp: &spb.PingResponse{Sent: tt.count, Received: tt.count}}, nil
			}
			want := tt.dest
			New(dut).NewPing().WithDestination(tt.dest).WithCount(tt.count).Operate(t)
			if got != want {
				t.Errorf("Operate(t) got %s, want %s", got, want)
			}

		})
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
	}, {
		wantErr: "ping sent 5 packets, received 3",
		dest:    "1.2.3.4",
		pinger: func(_ context.Context, req *spb.PingRequest, _ ...grpc.CallOption) (spb.System_PingClient, error) {
			return &fakePingClient{
				resp: &spb.PingResponse{Sent: 5, Received: 3},
			}, nil
		},
	}}
	for _, tt := range tests {
		t.Run(tt.wantErr, func(t *testing.T) {
			fakeGNOI.PingFn = tt.pinger
			op := New(dut).NewPing().WithDestination(tt.dest)
			gotErr := testt.ExpectFatal(t, func(t testing.TB) {
				op.Operate(t)
			})
			if !strings.Contains(gotErr, tt.wantErr) {
				t.Errorf("Operate(t) on ping got %q, want %q", gotErr, tt.wantErr)
			}
		})
	}
}

func TestReboot(t *testing.T) {
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
			fakeGNOI.RebootFn = func(context.Context, *spb.RebootRequest, ...grpc.CallOption) (*spb.RebootResponse, error) {
				rebooted = true
				return &spb.RebootResponse{}, nil
			}
			var index uint8
			fakeGNOI.RebootStatusFn = func(context.Context, *spb.RebootStatusRequest, ...grpc.CallOption) (*spb.RebootStatusResponse, error) {
				status, err := tt.statuses[index], tt.statusErrs[index]
				index = index + 1
				return status, err
			}
			New(dut).NewReboot().Operate(t)
			if !rebooted {
				t.Fatal("Operate() on reboot failed, want success")
			}
		})
	}
}

func TestRebootErrors(t *testing.T) {
	tests := []struct {
		desc                 string
		timeout              time.Duration
		rebootErr, statusErr error
		wantErr              string
	}{{
		desc:    "non-positive duration",
		timeout: -1,
		wantErr: "positive duration",
	}, {
		desc:      "reboot error",
		rebootErr: errors.New("reboot error"),
		wantErr:   "gnoi reboot",
	}, {
		desc:      "timeout error",
		timeout:   1,
		statusErr: errors.New("status error"),
		wantErr:   "timed out",
	}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			fakeGNOI.RebootFn = func(context.Context, *spb.RebootRequest, ...grpc.CallOption) (*spb.RebootResponse, error) {
				return &spb.RebootResponse{}, tt.rebootErr
			}
			fakeGNOI.RebootStatusFn = func(context.Context, *spb.RebootStatusRequest, ...grpc.CallOption) (*spb.RebootStatusResponse, error) {
				return &spb.RebootStatusResponse{}, tt.statusErr
			}
			reboot := New(dut).NewReboot().WithTimeout(tt.timeout)
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
	var killed bool
	fakeGNOI.KillProcessFn = func(context.Context, *spb.KillProcessRequest, ...grpc.CallOption) (*spb.KillProcessResponse, error) {
		killed = true
		return &spb.KillProcessResponse{}, nil
	}
	op := New(dut).NewKillProcess().WithPID(123)
	op.Operate(t)
	if !killed {
		t.Fatalf("Operate() on op %v failed, want success", op)
	}
}

func TestKillProcessErrors(t *testing.T) {
	fakeGNOI.KillProcessFn = func(context.Context, *spb.KillProcessRequest, ...grpc.CallOption) (*spb.KillProcessResponse, error) {
		return nil, errors.New("bad bad bad :(")
	}
	op := New(dut).NewKillProcess()
	gotErr := testt.ExpectFatal(t, func(t testing.TB) {
		op.Operate(t)
	})
	if gotErr == "" {
		t.Errorf("Operate() on op %v succeeded, want error", op)
	}
}

func TestTime(t *testing.T) {
	fakeGNOI.TimeFn = func(context.Context, *spb.TimeRequest, ...grpc.CallOption) (*spb.TimeResponse, error) {
		return &spb.TimeResponse{Time: 12345}, nil
	}
	got := New(dut).Time(t)
	if want := time.Unix(0, 12345); !want.Equal(got) {
		t.Fatalf("Operation failed, want %v got %v", want, got)
	}
}

func TestFactoryReset(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		fakeGNOI.StartFn = func(context.Context, *frpb.StartRequest, ...grpc.CallOption) (*frpb.StartResponse, error) {
			return &frpb.StartResponse{Response: &frpb.StartResponse_ResetSuccess{&frpb.ResetSuccess{}}}, nil
		}
		New(dut).NewFactoryReset().WithZeroFill(true).WithFactoryOS(true).Operate(t)
	})

	t.Run("error", func(t *testing.T) {
		fakeGNOI.StartFn = func(context.Context, *frpb.StartRequest, ...grpc.CallOption) (*frpb.StartResponse, error) {
			return &frpb.StartResponse{Response: &frpb.StartResponse_ResetError{&frpb.ResetError{}}}, nil
		}
		testt.ExpectFatal(t, func(t testing.TB) {
			New(dut).NewFactoryReset().WithZeroFill(true).WithFactoryOS(true).Operate(t)
		})
	})
}

func TestSwitchControlProcessor(t *testing.T) {
	var got string
	fakeGNOI.SwitchControlFn = func(ctx context.Context, in *spb.SwitchControlProcessorRequest, opts ...grpc.CallOption) (*spb.SwitchControlProcessorResponse, error) {
		got = in.GetControlProcessor().GetElem()[1].GetKey()["name"]
		return &spb.SwitchControlProcessorResponse{ControlProcessor: in.ControlProcessor}, nil
	}
	New(dut).NewSwitchControlProcessor().WithDestination("RP0").Operate(t)
	if want := "RP0"; !strings.Contains(got, want) {
		t.Errorf("Operate(t) on switch control process got %q, want %q", got, want)
	}
}

func TestSwitchControlProcessorErrors(t *testing.T) {
	sw := New(dut).NewSwitchControlProcessor()
	fakeGNOI.SwitchControlFn = func(context.Context, *spb.SwitchControlProcessorRequest, ...grpc.CallOption) (*spb.SwitchControlProcessorResponse, error) {
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
