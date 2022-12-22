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

// FetchConfig fetches config from the ATE.
func (o *OTG) FetchConfig(t testing.TB) gosnappi.Config {
	t.Helper()
	t = events.ActionStarted(t, "Fetching config from %s", o.ate)
	cfg, err := fetchConfig(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("FetchConfig(t) on %s: %v", o.ate, err)
	}
	return cfg
}

func fetchConfig(ctx context.Context, ate binding.ATE) (gosnappi.Config, error) {
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
	warns, err := setProtocolState(context.Background(), o.ate, gosnappi.ProtocolStateState.START)
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
	warns, err := setProtocolState(context.Background(), o.ate, gosnappi.ProtocolStateState.STOP)
	if err != nil {
		t.Fatalf("StopProtocols(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopProtocols(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func setProtocolState(ctx context.Context, ate binding.ATE, state gosnappi.ProtocolStateStateEnum) ([]string, error) {
	api, err := rawapis.FetchOTG(ctx, ate)
	if err != nil {
		return nil, err
	}
	resp, err := api.SetProtocolState(api.NewProtocolState().SetState(state))
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
	warns, err := setTransmitState(context.Background(), o.ate, gosnappi.TransmitStateState.START)
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
	warns, err := setTransmitState(context.Background(), o.ate, gosnappi.TransmitStateState.STOP)
	if err != nil {
		t.Fatalf("StopTraffic(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopTraffic(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func setTransmitState(ctx context.Context, ate binding.ATE, state gosnappi.TransmitStateStateEnum) ([]string, error) {
	api, err := rawapis.FetchOTG(ctx, ate)
	if err != nil {
		return nil, err
	}
	resp, err := api.SetTransmitState(api.NewTransmitState().SetState(state))
	if err != nil {
		return nil, err
	}
	return resp.Warnings(), nil
}

// AdvertiseRoutes advertises routes on the ATE.
func (o *OTG) AdvertiseRoutes(t testing.TB, routes []string) {
	t.Helper()
	t = events.ActionStarted(t, "Advertising routes on %v", o.ate)
	warns, err := setRouteState(context.Background(), o.ate, routes, gosnappi.RouteStateState.ADVERTISE)
	if err != nil {
		t.Fatalf("AdvertiseRoutes(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("AdvertiseRoutes(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// WithdrawRoutes withdraws routes on the ATE.
func (o *OTG) WithdrawRoutes(t testing.TB, routes []string) {
	t.Helper()
	t = events.ActionStarted(t, "Withdrawing routes for %v", o.ate)
	warns, err := setRouteState(context.Background(), o.ate, routes, gosnappi.RouteStateState.WITHDRAW)
	if err != nil {
		t.Fatalf("WithdrawRoutes(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("WithdrawRoutes(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func setRouteState(ctx context.Context, ate binding.ATE, routes []string, state gosnappi.RouteStateStateEnum) ([]string, error) {
	api, err := rawapis.FetchOTG(ctx, ate)
	if err != nil {
		return nil, err
	}
	resp, err := api.SetRouteState(api.NewRouteState().SetState(state).SetNames(routes))
	if err != nil {
		return nil, err
	}
	return resp.Warnings(), nil
}

// EnableLACPMembers enables lacp member ports on the ATE.
func (o *OTG) EnableLACPMembers(t testing.TB, ports ...string) {
	t.Helper()
	t = events.ActionStarted(t, "EnableLACPMembers on %v", o.ate)
	warns, err := setLACPMemberState(context.Background(), o.ate, gosnappi.LacpMemberStateState.UP, ports)
	if err != nil {
		t.Fatalf("EnableLACPMembers(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("EnableLACPMembers(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// DisableLACPMembers disables lacp member ports on the ATE.
func (o *OTG) DisableLACPMembers(t testing.TB, ports ...string) {
	t.Helper()
	t = events.ActionStarted(t, "DisableLACPMembers on %v", o.ate)
	warns, err := setLACPMemberState(context.Background(), o.ate, gosnappi.LacpMemberStateState.DOWN, ports)
	if err != nil {
		t.Fatalf("DisableLACPMembers(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("DisableLACPMembers(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func setLACPMemberState(ctx context.Context, ate binding.ATE, state gosnappi.LacpMemberStateStateEnum, lagMemberPorts []string) ([]string, error) {
	api, err := rawapis.FetchOTG(ctx, ate)
	if err != nil {
		return nil, err
	}
	resp, err := api.SetDeviceState(api.NewDeviceState().SetLacpMemberState(gosnappi.NewLacpMemberState().SetLagMemberPortNames(lagMemberPorts).SetState(state)))
	if err != nil {
		return nil, err
	}
	return resp.Warnings(), nil
}

// StartCapture starts capturing bytes on the specified ports.
func (o *OTG) StartCapture(t testing.TB, portNames ...string) {
	t.Helper()
	t = events.ActionStarted(t, "StartCapture on %v", o.ate)
	warns, err := setCaptureState(context.Background(), o.ate, gosnappi.CaptureStateState.START, portNames)
	if err != nil {
		t.Fatalf("StartCapture(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StartCapture(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// StopCapture starts capturing bytes on the specified ports.
func (o *OTG) StopCapture(t testing.TB, portNames ...string) {
	t.Helper()
	t = events.ActionStarted(t, "StopCapture on %v", o.ate)
	warns, err := setCaptureState(context.Background(), o.ate, gosnappi.CaptureStateState.STOP, portNames)
	if err != nil {
		t.Fatalf("StopCapture(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopCapture(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func setCaptureState(ctx context.Context, ate binding.ATE, state gosnappi.CaptureStateStateEnum, portNames []string) ([]string, error) {
	api, err := rawapis.FetchOTG(ctx, ate)
	if err != nil {
		return nil, err
	}
	resp, err := api.SetCaptureState(api.NewCaptureState().SetState(state).SetPortNames(portNames))
	if err != nil {
		return nil, err
	}
	return resp.Warnings(), nil
}

// FetchCapture fetches the captured bytes on the specified port.
func (o *OTG) FetchCapture(t testing.TB, portName string) []byte {
	t.Helper()
	t = events.ActionStarted(t, "FetchCapture on %v", o.ate)
	bytes, err := fetchCapture(context.Background(), o.ate, portName)
	if err != nil {
		t.Fatalf("FetchCapture(t) on %s: %v", o.ate, err)
	}
	return bytes
}

func fetchCapture(ctx context.Context, ate binding.ATE, postName string) ([]byte, error) {
	api, err := rawapis.FetchOTG(ctx, ate)
	if err != nil {
		return nil, err
	}
	return api.GetCapture(api.NewCaptureRequest().SetPortName(postName))
}

// GNMIOpts returns a new set of options to customize gNMI queries.
func (o *OTG) GNMIOpts() *gnmi.Opts {
	return gnmi.NewOpts(o.ate.Name(), false, func(ctx context.Context) (gpb.GNMIClient, error) {
		return rawapis.FetchOTGGNMI(ctx, o.ate)
	})
}
