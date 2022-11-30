// Copyright 2022 Google LLC
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

package otg

import (
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"
	"google.golang.org/grpc"
)

var (
	fakeSnappi = new(fakeGosnappi)
	fakeATE    = &fakebind.ATE{
		AbstractATE: &binding.AbstractATE{&binding.Dims{
			Ports: map[string]*binding.Port{"port1": {}},
		}},
		DialOTGFn: func(context.Context, ...grpc.DialOption) (gosnappi.GosnappiApi, error) {
			return fakeSnappi, nil
		},
	}
	otgAPI = &OTG{ate: fakeATE}
)

func TestNewConfig(t *testing.T) {
	wantCfg := gosnappi.NewConfig()
	fakeSnappi.cfg = wantCfg
	gotCfg := otgAPI.NewConfig(t)
	if wantCfg != gotCfg {
		t.Errorf("NewConfig got unexpected config %v, want %v", gotCfg, wantCfg)
	}
}

func TestFetchConfig(t *testing.T) {
	wantCfg := gosnappi.NewConfig()
	fakeSnappi.cfg = wantCfg
	gotCfg := otgAPI.FetchConfig(t)
	if wantCfg != gotCfg {
		t.Errorf("FetchConfig got unexpected config %v, want %v", gotCfg, wantCfg)
	}
}

func TestPushConfig(t *testing.T) {
	wantCfg := gosnappi.NewConfig()
	wantCfg.Ports().Add().SetName("port1")
	otgAPI.PushConfig(t, wantCfg)
	if gotCfg := fakeSnappi.cfg; wantCfg != gotCfg {
		t.Errorf("PushConfig got unexpected config %v, want %v", gotCfg, wantCfg)
	}
}

func TestPushConfigBadPortName(t *testing.T) {
	cfg := gosnappi.NewConfig()
	cfg.Ports().Add().SetName("port3")
	gotErr := testt.ExpectFatal(t, func(t testing.TB) {
		otgAPI.PushConfig(t, cfg)
	})
	if wantErr := "not one of the ATE ports"; !strings.Contains(gotErr, wantErr) {
		t.Errorf("PushConfig got unexpected err %s, want err containing %s", gotErr, wantErr)
	}
}

func TestStartProtocols(t *testing.T) {
	fakeSnappi.protocol = nil
	otgAPI.StartProtocols(t)
	if got, want := fakeSnappi.protocol.State(), gosnappi.ProtocolStateState.START; got != want {
		t.Errorf("StartProtocols got unexpected protocol state %v, want %v", got, want)
	}
}

func TestStopProtocols(t *testing.T) {
	fakeSnappi.protocol = nil
	otgAPI.StopProtocols(t)
	if got, want := fakeSnappi.protocol.State(), gosnappi.ProtocolStateState.STOP; got != want {
		t.Errorf("StopProtocols got unexpected protocol state %v, want %v", got, want)
	}
}

func TestStartTraffic(t *testing.T) {
	fakeSnappi.transmit = nil
	otgAPI.StartTraffic(t)
	if got, want := fakeSnappi.transmit.State(), gosnappi.TransmitStateState.START; got != want {
		t.Errorf("StartTraffic got unexpected transmit state %v, want %v", got, want)
	}
}

func TestStopTraffic(t *testing.T) {
	fakeSnappi.transmit = nil
	otgAPI.StopTraffic(t)
	if got, want := fakeSnappi.transmit.State(), gosnappi.TransmitStateState.STOP; got != want {
		t.Errorf("StopTraffic got unexpected transmit state %v, want %v", got, want)
	}
}

func TestAdvertiseRoutes(t *testing.T) {
	fakeSnappi.route = nil
	wantNames := []string{"peer1", "peer2"}
	otgAPI.AdvertiseRoutes(t, wantNames)
	got := fakeSnappi.route
	if wantState := gosnappi.RouteStateState.ADVERTISE; got.State() != wantState {
		t.Errorf("AdvertiseRoutes got unexpected route state %v, want %v", got.State(), wantState)
	}
	if !cmp.Equal(got.Names(), wantNames) {
		t.Errorf("AdvertiseRoutes got unexpected route names %v, want %v", got.Names(), wantNames)
	}
}

func TestWithdrawRoutes(t *testing.T) {
	fakeSnappi.route = nil
	wantNames := []string{"peer1", "peer2"}
	otgAPI.WithdrawRoutes(t, wantNames)
	got := fakeSnappi.route
	if wantState := gosnappi.RouteStateState.WITHDRAW; got.State() != wantState {
		t.Errorf("WithdrawRoutes got unexpected route state %v, want %v", got.State(), wantState)
	}
	if !cmp.Equal(got.Names(), wantNames) {
		t.Errorf("WithdrawRoutes got unexpected route names %v, want %v", got.Names(), wantNames)
	}
}

func TestEnableLACPMembers(t *testing.T) {
	wantPorts := []string{"port1", "port2"}
	otgAPI.EnableLACPMembers(t, wantPorts...)
	gotLACP := fakeSnappi.device.LacpMemberState()
	if wantState := gosnappi.LacpMemberStateState.UP; gotLACP.State() != wantState {
		t.Errorf("EnableLACPMembers got unexpected LACP member state %v, want %v", gotLACP.State(), wantState)
	}
	if !cmp.Equal(gotLACP.LagMemberPortNames(), wantPorts) {
		t.Errorf("EnableLACPMembers got unexpected LACP member ports %v, want %v", gotLACP.LagMemberPortNames(), wantPorts)
	}
}

func TestDisableLACPMembers(t *testing.T) {
	wantPorts := []string{"port1", "port2"}
	otgAPI.DisableLACPMembers(t, wantPorts...)
	gotLACP := fakeSnappi.device.LacpMemberState()
	if wantState := gosnappi.LacpMemberStateState.DOWN; gotLACP.State() != wantState {
		t.Errorf("DisableLACPMembers got unexpected LACP member state %v, want %v", gotLACP.State(), wantState)
	}
	if !cmp.Equal(gotLACP.LagMemberPortNames(), wantPorts) {
		t.Errorf("DisableLACPMembers got unexpected LACP member ports %v, want %v", gotLACP.LagMemberPortNames(), wantPorts)
	}
}

func TestStartCapture(t *testing.T) {
	wantPorts := []string{"port1", "port2"}
	otgAPI.StartCapture(t, wantPorts...)
	gotCapture := fakeSnappi.capture
	if wantState := gosnappi.CaptureStateState.START; gotCapture.State() != wantState {
		t.Errorf("StartCapture got unexpected capture state %v, want %v", gotCapture.State(), wantState)
	}
	if !cmp.Equal(gotCapture.PortNames(), wantPorts) {
		t.Errorf("StartCapture got unexpected capture ports %v, want %v", gotCapture.PortNames(), wantPorts)
	}
}

func TestStopCapture(t *testing.T) {
	wantPorts := []string{"port1", "port2"}
	otgAPI.StopCapture(t, wantPorts...)
	gotCapture := fakeSnappi.capture
	if wantState := gosnappi.CaptureStateState.STOP; gotCapture.State() != wantState {
		t.Errorf("StopCapture got unexpected capture state %v, want %v", gotCapture.State(), wantState)
	}
	if !cmp.Equal(gotCapture.PortNames(), wantPorts) {
		t.Errorf("StopCapture got unexpected capture ports %v, want %v", gotCapture.PortNames(), wantPorts)
	}
}

func TestFetchCapture(t *testing.T) {
	wantPort := "port1"
	otgAPI.FetchCapture(t, wantPort)
	gotCaptureReq := fakeSnappi.captureReq
	if !cmp.Equal(gotCaptureReq.PortName(), wantPort) {
		t.Errorf("FetchCapture got unexpected capture ports %v, want %v", gotCaptureReq.PortName(), wantPort)
	}
}

type fakeGosnappi struct {
	gosnappi.GosnappiApi
	cfg        gosnappi.Config
	protocol   gosnappi.ProtocolState
	transmit   gosnappi.TransmitState
	route      gosnappi.RouteState
	device     gosnappi.DeviceState
	capture    gosnappi.CaptureState
	captureReq gosnappi.CaptureRequest
}

func (fg *fakeGosnappi) NewConfig() gosnappi.Config {
	return fg.cfg
}

func (fg *fakeGosnappi) GetConfig() (gosnappi.Config, error) {
	return fg.cfg, nil
}

func (fg *fakeGosnappi) SetConfig(cfg gosnappi.Config) (gosnappi.ResponseWarning, error) {
	fg.cfg = cfg
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewProtocolState() gosnappi.ProtocolState {
	return new(fakeProtocolState)
}

func (fg *fakeGosnappi) SetProtocolState(state gosnappi.ProtocolState) (gosnappi.ResponseWarning, error) {
	fg.protocol = state
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewTransmitState() gosnappi.TransmitState {
	return new(fakeTransmitState)
}

func (fg *fakeGosnappi) SetTransmitState(state gosnappi.TransmitState) (gosnappi.ResponseWarning, error) {
	fg.transmit = state
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewRouteState() gosnappi.RouteState {
	return new(fakeRouteState)
}

func (fg *fakeGosnappi) SetRouteState(state gosnappi.RouteState) (gosnappi.ResponseWarning, error) {
	fg.route = state
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewDeviceState() gosnappi.DeviceState {
	return new(fakeDeviceState)
}

func (fg *fakeGosnappi) SetDeviceState(state gosnappi.DeviceState) (gosnappi.ResponseWarning, error) {
	fg.device = state
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewCaptureState() gosnappi.CaptureState {
	return new(fakeCaptureState)
}

func (fg *fakeGosnappi) SetCaptureState(state gosnappi.CaptureState) (gosnappi.ResponseWarning, error) {
	fg.capture = state
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewCaptureRequest() gosnappi.CaptureRequest {
	return new(fakeCaptureRequest)
}

func (fg *fakeGosnappi) GetCapture(request gosnappi.CaptureRequest) ([]byte, error) {
	fg.captureReq = request
	return nil, nil
}

type fakeProtocolState struct {
	gosnappi.ProtocolState
	state gosnappi.ProtocolStateStateEnum
}

func (fp *fakeProtocolState) State() gosnappi.ProtocolStateStateEnum {
	return fp.state
}

func (fp *fakeProtocolState) SetState(state gosnappi.ProtocolStateStateEnum) gosnappi.ProtocolState {
	fp.state = state
	return fp
}

type fakeTransmitState struct {
	gosnappi.TransmitState
	state gosnappi.TransmitStateStateEnum
}

func (ft *fakeTransmitState) State() gosnappi.TransmitStateStateEnum {
	return ft.state
}

func (ft *fakeTransmitState) SetState(state gosnappi.TransmitStateStateEnum) gosnappi.TransmitState {
	ft.state = state
	return ft
}

type fakeRouteState struct {
	gosnappi.RouteState
	state gosnappi.RouteStateStateEnum
	names []string
}

func (fp *fakeRouteState) State() gosnappi.RouteStateStateEnum {
	return fp.state
}

func (fp *fakeRouteState) Names() []string {
	return fp.names
}

func (fp *fakeRouteState) SetState(state gosnappi.RouteStateStateEnum) gosnappi.RouteState {
	fp.state = state
	return fp
}

func (fp *fakeRouteState) SetNames(names []string) gosnappi.RouteState {
	fp.names = names
	return fp
}

type fakeDeviceState struct {
	gosnappi.DeviceState
	lacp gosnappi.LacpMemberState
}

func (fd *fakeDeviceState) LacpMemberState() gosnappi.LacpMemberState {
	return fd.lacp
}

func (fd *fakeDeviceState) SetLacpMemberState(lacp gosnappi.LacpMemberState) gosnappi.DeviceState {
	fd.lacp = lacp
	return fd
}

type fakeCaptureState struct {
	gosnappi.CaptureState
	capture   gosnappi.CaptureStateStateEnum
	portNames []string
}

func (fc *fakeCaptureState) State() gosnappi.CaptureStateStateEnum {
	return fc.capture
}

func (fc *fakeCaptureState) PortNames() []string {
	return fc.portNames
}

func (fc *fakeCaptureState) SetState(capture gosnappi.CaptureStateStateEnum) gosnappi.CaptureState {
	fc.capture = capture
	return fc
}

func (fc *fakeCaptureState) SetPortNames(portNames []string) gosnappi.CaptureState {
	fc.portNames = portNames
	return fc
}

type fakeCaptureRequest struct {
	gosnappi.CaptureRequest
	portName string
}

func (fr *fakeCaptureRequest) PortName() string {
	return fr.portName
}

func (fr *fakeCaptureRequest) SetPortName(portName string) gosnappi.CaptureRequest {
	fr.portName = portName
	return fr
}
