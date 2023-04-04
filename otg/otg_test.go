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
	fakeSnappi.config = wantCfg
	gotCfg := otgAPI.NewConfig(t)
	if wantCfg != gotCfg {
		t.Errorf("NewConfig got unexpected config %v, want %v", gotCfg, wantCfg)
	}
}

func TestFetchConfig(t *testing.T) {
	wantCfg := gosnappi.NewConfig()
	fakeSnappi.config = wantCfg
	gotCfg := otgAPI.FetchConfig(t)
	if wantCfg != gotCfg {
		t.Errorf("FetchConfig got unexpected config %v, want %v", gotCfg, wantCfg)
	}
}

func TestPushConfig(t *testing.T) {
	wantCfg := gosnappi.NewConfig()
	wantCfg.Ports().Add().SetName("port1")
	otgAPI.PushConfig(t, wantCfg)
	if gotCfg := fakeSnappi.config; wantCfg != gotCfg {
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
	fakeSnappi.controlState = nil
	otgAPI.StartProtocols(t)
	if got, want := fakeSnappi.controlState.Protocol().All().State(), gosnappi.StateProtocolAllState.START; got != want {
		t.Errorf("StartProtocols got unexpected protocol state %v, want %v", got, want)
	}
}

func TestStopProtocols(t *testing.T) {
	fakeSnappi.controlState = nil
	otgAPI.StopProtocols(t)
	if got, want := fakeSnappi.controlState.Protocol().All().State(), gosnappi.StateProtocolAllState.STOP; got != want {
		t.Errorf("StopProtocols got unexpected protocol state %v, want %v", got, want)
	}
}

func TestStartTraffic(t *testing.T) {
	fakeSnappi.controlState = nil
	otgAPI.StartTraffic(t)
	if got, want := fakeSnappi.controlState.Traffic().FlowTransmit().State(), gosnappi.StateTrafficFlowTransmitState.START; got != want {
		t.Errorf("StartTraffic got unexpected transmit state %v, want %v", got, want)
	}
}

func TestStopTraffic(t *testing.T) {
	fakeSnappi.controlState = nil
	otgAPI.StopTraffic(t)
	if got, want := fakeSnappi.controlState.Traffic().FlowTransmit().State(), gosnappi.StateTrafficFlowTransmitState.STOP; got != want {
		t.Errorf("StopTraffic got unexpected transmit state %v, want %v", got, want)
	}
}

func TestAdvertiseRoutes(t *testing.T) {
	fakeSnappi.controlState = nil
	wantNames := []string{"peer1", "peer2"}
	otgAPI.AdvertiseRoutes(t, wantNames)
	got := fakeSnappi.controlState.Protocol().Route()
	if wantState := gosnappi.StateProtocolRouteState.ADVERTISE; got.State() != wantState {
		t.Errorf("AdvertiseRoutes got unexpected route state %v, want %v", got.State(), wantState)
	}
	if !cmp.Equal(got.Names(), wantNames) {
		t.Errorf("AdvertiseRoutes got unexpected route names %v, want %v", got.Names(), wantNames)
	}
}

func TestWithdrawRoutes(t *testing.T) {
	fakeSnappi.controlState = nil
	wantNames := []string{"peer1", "peer2"}
	otgAPI.WithdrawRoutes(t, wantNames)
	got := fakeSnappi.controlState.Protocol().Route()
	if wantState := gosnappi.StateProtocolRouteState.WITHDRAW; got.State() != wantState {
		t.Errorf("WithdrawRoutes got unexpected route state %v, want %v", got.State(), wantState)
	}
	if !cmp.Equal(got.Names(), wantNames) {
		t.Errorf("WithdrawRoutes got unexpected route names %v, want %v", got.Names(), wantNames)
	}
}

func TestEnableLACPMembers(t *testing.T) {
	fakeSnappi.controlState = nil
	wantNames := []string{"port1", "port2"}
	otgAPI.EnableLACPMembers(t, wantNames...)
	got := fakeSnappi.controlState.Protocol().Lacp().Admin()
	if wantState := gosnappi.StateProtocolLacpAdminState.UP; got.State() != wantState {
		t.Errorf("EnableLACPMembers got unexpected LACP member state %v, want %v", got.State(), wantState)
	}
	if !cmp.Equal(got.LagMemberNames(), wantNames) {
		t.Errorf("EnableLACPMembers got unexpected LACP member ports %v, want %v", got.LagMemberNames(), wantNames)
	}
}

func TestDisableLACPMembers(t *testing.T) {
	fakeSnappi.controlState = nil
	wantNames := []string{"port1", "port2"}
	otgAPI.DisableLACPMembers(t, wantNames...)
	got := fakeSnappi.controlState.Protocol().Lacp().Admin()
	if wantState := gosnappi.StateProtocolLacpAdminState.DOWN; got.State() != wantState {
		t.Errorf("DisableLACPMembers got unexpected LACP member state %v, want %v", got.State(), wantState)
	}
	if !cmp.Equal(got.LagMemberNames(), wantNames) {
		t.Errorf("DisableLACPMembers got unexpected LACP member ports %v, want %v", got.LagMemberNames(), wantNames)
	}
}

func TestStartCapture(t *testing.T) {
	fakeSnappi.controlState = nil
	wantPorts := []string{"port1", "port2"}
	otgAPI.StartCapture(t, wantPorts...)
	got := fakeSnappi.controlState.Port().Capture()
	if wantState := gosnappi.StatePortCaptureState.START; got.State() != wantState {
		t.Errorf("StartCapture got unexpected capture state %v, want %v", got.State(), wantState)
	}
	if !cmp.Equal(got.PortNames(), wantPorts) {
		t.Errorf("StartCapture got unexpected capture ports %v, want %v", got.PortNames(), wantPorts)
	}
}

func TestStopCapture(t *testing.T) {
	fakeSnappi.controlState = nil
	wantPorts := []string{"port1", "port2"}
	otgAPI.StopCapture(t, wantPorts...)
	got := fakeSnappi.controlState.Port().Capture()
	if wantState := gosnappi.StatePortCaptureState.STOP; got.State() != wantState {
		t.Errorf("StopCapture got unexpected capture state %v, want %v", got.State(), wantState)
	}
	if !cmp.Equal(got.PortNames(), wantPorts) {
		t.Errorf("StopCapture got unexpected capture ports %v, want %v", got.PortNames(), wantPorts)
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
	config       gosnappi.Config
	controlState gosnappi.ControlState
	captureReq   gosnappi.CaptureRequest
}

func (fg *fakeGosnappi) NewConfig() gosnappi.Config {
	return fg.config
}

func (fg *fakeGosnappi) GetConfig() (gosnappi.Config, error) {
	return fg.config, nil
}

func (fg *fakeGosnappi) SetConfig(cfg gosnappi.Config) (gosnappi.Warning, error) {
	fg.config = cfg
	return gosnappi.NewWarning(), nil
}

func (fg *fakeGosnappi) NewControlState() gosnappi.ControlState {
	return new(fakeControlState)
}

func (fg *fakeGosnappi) SetControlState(state gosnappi.ControlState) (gosnappi.Warning, error) {
	fg.controlState = state
	return gosnappi.NewWarning(), nil
}

func (fg *fakeGosnappi) NewCaptureRequest() gosnappi.CaptureRequest {
	return new(fakeCaptureRequest)
}

func (fg *fakeGosnappi) GetCapture(request gosnappi.CaptureRequest) ([]byte, error) {
	fg.captureReq = request
	return nil, nil
}

type fakeControlState struct {
	gosnappi.ControlState
	port     gosnappi.StatePort
	protocol gosnappi.StateProtocol
	traffic  gosnappi.StateTraffic
}

func (fc *fakeControlState) Port() gosnappi.StatePort {
	if fc.port == nil {
		fc.port = new(fakePort)
	}
	return fc.port
}

func (fc *fakeControlState) Protocol() gosnappi.StateProtocol {
	if fc.protocol == nil {
		fc.protocol = new(fakeProtocol)
	}
	return fc.protocol
}

func (fc *fakeControlState) Traffic() gosnappi.StateTraffic {
	if fc.traffic == nil {
		fc.traffic = new(fakeTraffic)
	}
	return fc.traffic
}

type fakePort struct {
	gosnappi.StatePort
	capture gosnappi.StatePortCapture
}

func (fp *fakePort) Capture() gosnappi.StatePortCapture {
	if fp.capture == nil {
		fp.capture = new(fakeCapture)
	}
	return fp.capture
}

type fakeCapture struct {
	gosnappi.StatePortCapture
	state gosnappi.StatePortCaptureStateEnum
	ports []string
}

func (fc *fakeCapture) PortNames() []string {
	return fc.ports
}

func (fc *fakeCapture) SetPortNames(ports []string) gosnappi.StatePortCapture {
	fc.ports = ports
	return fc
}

func (fc *fakeCapture) SetState(state gosnappi.StatePortCaptureStateEnum) gosnappi.StatePortCapture {
	fc.state = state
	return fc
}

func (fc *fakeCapture) State() gosnappi.StatePortCaptureStateEnum {
	return fc.state
}

type fakeProtocol struct {
	gosnappi.StateProtocol
	all   gosnappi.StateProtocolAll
	lacp  gosnappi.StateProtocolLacp
	route gosnappi.StateProtocolRoute
}

func (fp *fakeProtocol) All() gosnappi.StateProtocolAll {
	if fp.all == nil {
		fp.all = new(fakeProtocolAll)
	}
	return fp.all
}

func (fp *fakeProtocol) Lacp() gosnappi.StateProtocolLacp {
	if fp.lacp == nil {
		fp.lacp = new(fakeProtocolLacp)
	}
	return fp.lacp
}

func (fp *fakeProtocol) Route() gosnappi.StateProtocolRoute {
	if fp.route == nil {
		fp.route = new(fakeProtocolRoute)
	}
	return fp.route
}

type fakeProtocolAll struct {
	gosnappi.StateProtocolAll
	state gosnappi.StateProtocolAllStateEnum
}

func (fa *fakeProtocolAll) SetState(state gosnappi.StateProtocolAllStateEnum) gosnappi.StateProtocolAll {
	fa.state = state
	return fa
}

func (fa *fakeProtocolAll) State() gosnappi.StateProtocolAllStateEnum {
	return fa.state
}

type fakeProtocolLacp struct {
	gosnappi.StateProtocolLacp
	admin gosnappi.StateProtocolLacpAdmin
}

func (fl *fakeProtocolLacp) Admin() gosnappi.StateProtocolLacpAdmin {
	if fl.admin == nil {
		fl.admin = new(fakeProtocolLacpAdmin)
	}
	return fl.admin
}

type fakeProtocolLacpAdmin struct {
	gosnappi.StateProtocolLacpAdmin
	state gosnappi.StateProtocolLacpAdminStateEnum
	names []string
}

func (fp *fakeProtocolLacpAdmin) LagMemberNames() []string {
	return fp.names
}

func (fp *fakeProtocolLacpAdmin) SetLagMemberNames(names []string) gosnappi.StateProtocolLacpAdmin {
	fp.names = names
	return fp
}

func (fp *fakeProtocolLacpAdmin) SetState(state gosnappi.StateProtocolLacpAdminStateEnum) gosnappi.StateProtocolLacpAdmin {
	fp.state = state
	return fp
}

func (fp *fakeProtocolLacpAdmin) State() gosnappi.StateProtocolLacpAdminStateEnum {
	return fp.state
}

type fakeProtocolRoute struct {
	gosnappi.StateProtocolRoute
	state gosnappi.StateProtocolRouteStateEnum
	names []string
}

func (fr *fakeProtocolRoute) Names() []string {
	return fr.names
}

func (fr *fakeProtocolRoute) SetNames(names []string) gosnappi.StateProtocolRoute {
	fr.names = names
	return fr
}

func (fr *fakeProtocolRoute) SetState(state gosnappi.StateProtocolRouteStateEnum) gosnappi.StateProtocolRoute {
	fr.state = state
	return fr
}

func (fr *fakeProtocolRoute) State() gosnappi.StateProtocolRouteStateEnum {
	return fr.state
}

type fakeTraffic struct {
	gosnappi.StateTraffic
	flowTransmit gosnappi.StateTrafficFlowTransmit
}

func (ft *fakeTraffic) FlowTransmit() gosnappi.StateTrafficFlowTransmit {
	if ft.flowTransmit == nil {
		ft.flowTransmit = new(fakeFlowTransmit)
	}
	return ft.flowTransmit
}

type fakeFlowTransmit struct {
	gosnappi.StateTrafficFlowTransmit
	state gosnappi.StateTrafficFlowTransmitStateEnum
}

func (ft *fakeFlowTransmit) SetState(state gosnappi.StateTrafficFlowTransmitStateEnum) gosnappi.StateTrafficFlowTransmit {
	ft.state = state
	return ft
}

func (ft *fakeFlowTransmit) State() gosnappi.StateTrafficFlowTransmitStateEnum {
	return ft.state
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
