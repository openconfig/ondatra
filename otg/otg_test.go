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
	"golang.org/x/net/context"
	"strings"
	"testing"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/testt"
)

var (
	fakeSnappi = new(fakeGosnappi)
	fakeATE    = &fakebind.ATE{
		AbstractATE: &binding.AbstractATE{&binding.Dims{
			Ports: map[string]*binding.Port{
				"port1": &binding.Port{},
			},
		}},
		DialOTGFn: func(context.Context) (gosnappi.GosnappiApi, error) {
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
	fakeSnappi.protocol = gosnappi.ProtocolStateState.STOP
	otgAPI.StartProtocols(t)
	if got, want := fakeSnappi.protocol, gosnappi.ProtocolStateState.START; got != want {
		t.Errorf("StartProtocols got unexpected protocol state %v, want %v", got, want)
	}
}

func TestStopProtocols(t *testing.T) {
	fakeSnappi.protocol = gosnappi.ProtocolStateState.START
	otgAPI.StopProtocols(t)
	if got, want := fakeSnappi.protocol, gosnappi.ProtocolStateState.STOP; got != want {
		t.Errorf("StopProtocols got unexpected protocol state %v, want %v", got, want)
	}
}

func TestStartTraffic(t *testing.T) {
	fakeSnappi.transmit = gosnappi.TransmitStateState.STOP
	otgAPI.StartTraffic(t)
	if got, want := fakeSnappi.transmit, gosnappi.TransmitStateState.START; got != want {
		t.Errorf("StartTraffic got unexpected transmit state %v, want %v", got, want)
	}
}

func TestStopTraffic(t *testing.T) {
	fakeSnappi.transmit = gosnappi.TransmitStateState.START
	otgAPI.StopTraffic(t)
	if got, want := fakeSnappi.transmit, gosnappi.TransmitStateState.STOP; got != want {
		t.Errorf("StopTraffic got unexpected transmit state %v, want %v", got, want)
	}
}

type fakeGosnappi struct {
	gosnappi.GosnappiApi
	cfg      gosnappi.Config
	protocol gosnappi.ProtocolStateStateEnum
	transmit gosnappi.TransmitStateStateEnum
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
	fg.protocol = state.(*fakeProtocolState).state
	return gosnappi.NewResponseWarning(), nil
}

func (fg *fakeGosnappi) NewTransmitState() gosnappi.TransmitState {
	return new(fakeTransmitState)
}

func (fg *fakeGosnappi) SetTransmitState(state gosnappi.TransmitState) (gosnappi.ResponseWarning, error) {
	fg.transmit = state.(*fakeTransmitState).state
	return gosnappi.NewResponseWarning(), nil
}

type fakeProtocolState struct {
	gosnappi.ProtocolState
	state gosnappi.ProtocolStateStateEnum
}

func (fp *fakeProtocolState) SetState(state gosnappi.ProtocolStateStateEnum) gosnappi.ProtocolState {
	fp.state = state
	return fp
}

type fakeTransmitState struct {
	gosnappi.TransmitState
	state gosnappi.TransmitStateStateEnum
}

func (ft *fakeTransmitState) SetState(state gosnappi.TransmitStateStateEnum) gosnappi.TransmitState {
	ft.state = state
	return ft
}
