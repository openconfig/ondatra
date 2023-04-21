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

func TestSetControlState(t *testing.T) {
	fakeSnappi.controlState = nil
	want := gosnappi.NewControlState()
	otgAPI.SetControlState(t, want)
	if got := fakeSnappi.controlState; got != want {
		t.Errorf("TestSetControlState got unexpected control state %v, want %v", got, want)
	}
}

func TestSetControlAction(t *testing.T) {
	fakeSnappi.controlAction = nil
	want := gosnappi.NewControlAction()
	otgAPI.SetControlAction(t, want)
	if got := fakeSnappi.controlAction; got != want {
		t.Errorf("SetControlAction got unexpected control action %v, want %v", got, want)
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

func TestGetCapture(t *testing.T) {
	want := gosnappi.NewCaptureRequest().SetPortName("port1")
	otgAPI.GetCapture(t, want)
	if got := fakeSnappi.captureReq; got != want {
		t.Errorf("GetCapture got unexpected request %v, want %v", got, want)
	}
}

type fakeGosnappi struct {
	gosnappi.GosnappiApi
	config        gosnappi.Config
	controlState  gosnappi.ControlState
	controlAction gosnappi.ControlAction
	captureReq    gosnappi.CaptureRequest
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

func (fg *fakeGosnappi) SetControlState(state gosnappi.ControlState) (gosnappi.Warning, error) {
	fg.controlState = state
	return gosnappi.NewWarning(), nil
}

func (fg *fakeGosnappi) SetControlAction(action gosnappi.ControlAction) (gosnappi.ControlActionResponse, error) {
	fg.controlAction = action
	return gosnappi.NewControlActionResponse(), nil
}

func (fg *fakeGosnappi) GetCapture(request gosnappi.CaptureRequest) ([]byte, error) {
	fg.captureReq = request
	return nil, nil
}
