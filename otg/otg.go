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

// Package otg provides an API to Open Traffic Generator.
package otg

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/gnmi"
	"github.com/openconfig/ondatra/internal/events"
	"github.com/openconfig/ondatra/internal/rawapis"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// New constructs a new OTG instance.
// Tests should not call this directly; call ate.OTG() instead.
func New(ate binding.ATE) *OTG {
	return &OTG{ate: ate}
}

var _ gnmi.DeviceOrOpts = (*OTG)(nil)

// OTG provides the Open Traffic Generator API to an ATE.
// Tests should not construct this directly; call ate.OTG() instead.
type OTG struct {
	ate binding.ATE
}

func (o *OTG) String() string {
	return fmt.Sprintf("{ate: %s}", o.ate)
}

// NewConfig creates a new OTG config.
// Deprecated: Use gosnappi.NewConfig directly.
func (o *OTG) NewConfig(t testing.TB) gosnappi.Config {
	t.Helper()
	t = events.ActionStarted(t, "Creating new config for %s", o.ate)
	cfg, err := newConfig(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("NewConfig(t) on %s: %v", o.ate, err)
	}
	return cfg
}

func newConfig(ctx context.Context, ate binding.ATE) (gosnappi.Config, error) {
	api, err := rawapis.FetchOTG(ctx, ate)
	if err != nil {
		return nil, err
	}
	return api.NewConfig(), nil
}

// PushConfig pushes config to the ATE.
func (o *OTG) PushConfig(t testing.TB, cfg gosnappi.Config) {
	t.Helper()
	t = events.ActionStarted(t, "Pushing config to %s", o.ate)
	warns, err := pushConfig(context.Background(), o.ate, cfg)
	if err != nil {
		t.Fatalf("PushConfig(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("PushConfig(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func pushConfig(ctx context.Context, ate binding.ATE, cfg gosnappi.Config) ([]string, error) {
	api, err := rawapis.FetchOTG(ctx, ate)
	if err != nil {
		return nil, err
	}
	for _, cfgPort := range cfg.Ports().Items() {
		bindPort, ok := ate.Ports()[cfgPort.Name()]
		if !ok {
			var tbPorts []string
			for p := range ate.Ports() {
				tbPorts = append(tbPorts, p)
			}
			return nil, fmt.Errorf("port %s in config is not one of the ATE ports %v", cfgPort.Name(), tbPorts)
		}
		cfgPort.SetLocation(bindPort.Name)
	}
	resp, err := api.SetConfig(cfg)
	if err != nil {
		return nil, err
	}
	return resp.Warnings(), nil
}

// GetConfig gets the current config.
func (o *OTG) GetConfig(t testing.TB) gosnappi.Config {
	t.Helper()
	t = events.ActionStarted(t, "Getting config from %s", o.ate)
	cfg, err := getConfig(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("GetConfig(t) on %s: %v", o.ate, err)
	}
	return cfg
}

// FetchConfig gets the current config.
// Deprecated: Use GetConfig instead.
func (o *OTG) FetchConfig(t testing.TB) gosnappi.Config {
	t.Helper()
	t = events.ActionStarted(t, "Fetching config from %s", o.ate)
	cfg, err := getConfig(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("FetchConfig(t) on %s: %v", o.ate, err)
	}
	return cfg
}

func getConfig(ctx context.Context, ate binding.ATE) (gosnappi.Config, error) {
	api, err := rawapis.FetchOTG(ctx, ate)
	if err != nil {
		return nil, err
	}
	return api.GetConfig()
}

// StartProtocols starts protocols on the ATE.
func (o *OTG) StartProtocols(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Starting protocols on %s", o.ate)
	warns, err := o.setProtocolState(context.Background(), gosnappi.StateProtocolAllState.START)
	if err != nil {
		t.Fatalf("StartProtocols(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StartProtocols(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// StopProtocols stops protocols on the ATE.
func (o *OTG) StopProtocols(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Stopping protocols on %s", o.ate)
	warns, err := o.setProtocolState(context.Background(), gosnappi.StateProtocolAllState.STOP)
	if err != nil {
		t.Fatalf("StopProtocols(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopProtocols(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func (o *OTG) setProtocolState(ctx context.Context, state gosnappi.StateProtocolAllStateEnum) ([]string, error) {
	controlState := gosnappi.NewControlState()
	controlState.Protocol().All().SetState(state)
	resp, err := o.setControlState(ctx, controlState)
	if err != nil {
		return nil, err
	}
	return resp.Warnings(), nil
}

// StartTraffic starts traffic on the ATE.
func (o *OTG) StartTraffic(t testing.TB) {
	t.Helper()
	// TODO(greg-dennis): Remove sleep when Keysight fixes a MAC resolution bug.
	time.Sleep(2 * time.Second)
	t = events.ActionStarted(t, "Starting traffic on %s", o.ate)
	warns, err := o.setTransmitState(context.Background(), gosnappi.StateTrafficFlowTransmitState.START)
	if err != nil {
		t.Fatalf("StartTraffic(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StartTraffic(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// StopTraffic stops traffic on the ATE.
func (o *OTG) StopTraffic(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Stopping traffic on %s", o.ate)
	warns, err := o.setTransmitState(context.Background(), gosnappi.StateTrafficFlowTransmitState.STOP)
	if err != nil {
		t.Fatalf("StopTraffic(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopTraffic(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func (o *OTG) setTransmitState(ctx context.Context, state gosnappi.StateTrafficFlowTransmitStateEnum) ([]string, error) {
	controlState := gosnappi.NewControlState()
	controlState.Traffic().FlowTransmit().SetState(state)
	resp, err := o.setControlState(ctx, controlState)
	if err != nil {
		return nil, err
	}
	return resp.Warnings(), nil
}

// SetControlState sets the operational state of configured resources.
func (o *OTG) SetControlState(t testing.TB, state gosnappi.ControlState) {
	t.Helper()
	t = events.ActionStarted(t, "SetControlState on %v", o.ate)
	resp, err := o.setControlState(context.Background(), state)
	if err != nil {
		t.Fatalf("SetControlState(t) on %s: %v", o.ate, err)
	}
	if len(resp.Warnings()) > 0 {
		t.Logf("SetControlState(t) on %s non-fatal warnings: %v", o.ate, resp.Warnings())
	}
}

func (o *OTG) setControlState(ctx context.Context, state gosnappi.ControlState) (gosnappi.Warning, error) {
	api, err := rawapis.FetchOTG(ctx, o.ate)
	if err != nil {
		return nil, err
	}
	return api.SetControlState(state)
}

// SetControlAction triggers actions against configured resources.
func (o *OTG) SetControlAction(t testing.TB, action gosnappi.ControlAction) {
	t.Helper()
	t = events.ActionStarted(t, "SetControlAction on %v", o.ate)
	resp, err := o.setControlAction(action)
	if err != nil {
		t.Fatalf("SetControlAction(t) on %s: %v", o.ate, err)
	}
	if len(resp.Warnings()) > 0 {
		t.Logf("SetControlAction(t) on %s non-fatal warnings: %v", o.ate, resp.Warnings())
	}
}

func (o *OTG) setControlAction(action gosnappi.ControlAction) (gosnappi.ControlActionResponse, error) {
	api, err := rawapis.FetchOTG(context.Background(), o.ate)
	if err != nil {
		return nil, err
	}
	return api.SetControlAction(action)
}

// DisableLACPMembers disables lacp member ports on the ATE.
// Deprecated: Use SetControlState instead.
func (o *OTG) DisableLACPMembers(t testing.TB, ports ...string) {
	t.Helper()
	t = events.ActionStarted(t, "DisableLACPMembers on %v", o.ate)
	controlState := gosnappi.NewControlState()
	controlState.Protocol().Lacp().Admin().SetState(gosnappi.StateProtocolLacpAdminState.DOWN).SetLagMemberNames(ports)
	resp, err := o.setControlState(context.Background(), controlState)
	if err != nil {
		t.Fatalf("DisableLACPMembers(t) on %s: %v", o.ate, err)
	}
	if len(resp.Warnings()) > 0 {
		t.Logf("DisableLACPMembers(t) on %s non-fatal warnings: %v", o.ate, resp.Warnings())
	}
}

// GetCapture gets the results of a port capture.
func (o *OTG) GetCapture(t testing.TB, req gosnappi.CaptureRequest) []byte {
	t.Helper()
	t = events.ActionStarted(t, "GetCapture on %v", o.ate)
	bytes, err := getCapture(context.Background(), o.ate, req)
	if err != nil {
		t.Fatalf("GetCapture(t) on %s: %v", o.ate, err)
	}
	return bytes
}

func getCapture(ctx context.Context, ate binding.ATE, req gosnappi.CaptureRequest) ([]byte, error) {
	api, err := rawapis.FetchOTG(ctx, ate)
	if err != nil {
		return nil, err
	}
	return api.GetCapture(req)
}

// GNMIOpts returns a new set of options to customize gNMI queries.
func (o *OTG) GNMIOpts() *gnmi.Opts {
	return gnmi.NewOpts(o.ate.Name(), false, func(ctx context.Context) (gpb.GNMIClient, error) {
		return rawapis.FetchOTGGNMI(ctx, o.ate)
	})
}
