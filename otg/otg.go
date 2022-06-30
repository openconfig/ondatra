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
	"golang.org/x/net/context"
	"fmt"
	"sync"
	"testing"

	"google.golang.org/grpc"
	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/debugger"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ondatra/telemetry/otg/device"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

var (
	mu    sync.Mutex
	apis  = make(map[binding.ATE]gosnappi.GosnappiApi)
	gnmis = make(map[binding.ATE]gpb.GNMIClient)
)

// New constructs a new OTG instance.
// Tests should not call this directly; call ate.OTG() instead.
func New(ate binding.ATE) *OTG {
	return &OTG{ate: ate}
}

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
	debugger.ActionStarted(t, "Creating new config for %s", o.ate)
	cfg, err := newConfig(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("NewConfig(t) on %s: %v", o.ate, err)
	}
	return cfg
}

func newConfig(ctx context.Context, ate binding.ATE) (gosnappi.Config, error) {
	api, err := fetchAPI(ctx, ate)
	if err != nil {
		return nil, err
	}
	return api.NewConfig(), nil
}

// PushConfig pushes config to the ATE.
func (o *OTG) PushConfig(t testing.TB, cfg gosnappi.Config) {
	t.Helper()
	debugger.ActionStarted(t, "Pushing config to %s", o.ate)
	warns, err := pushConfig(context.Background(), o.ate, cfg)
	if err != nil {
		t.Fatalf("PushConfig(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("PushConfig(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func pushConfig(ctx context.Context, ate binding.ATE, cfg gosnappi.Config) ([]string, error) {
	api, err := fetchAPI(ctx, ate)
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
	debugger.ActionStarted(t, "Fetching config from %s", o.ate)
	cfg, err := fetchConfig(context.Background(), o.ate)
	if err != nil {
		t.Fatalf("FetchConfig(t) on %s: %v", o.ate, err)
	}
	return cfg
}

func fetchConfig(ctx context.Context, ate binding.ATE) (gosnappi.Config, error) {
	api, err := fetchAPI(ctx, ate)
	if err != nil {
		return nil, err
	}
	return api.GetConfig()
}

// StartProtocols starts protocols on the ATE.
func (o *OTG) StartProtocols(t testing.TB) {
	t.Helper()
	debugger.ActionStarted(t, "Starting protocols on %s", o.ate)
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
	debugger.ActionStarted(t, "Stopping protocols on %s", o.ate)
	warns, err := setProtocolState(context.Background(), o.ate, gosnappi.ProtocolStateState.STOP)
	if err != nil {
		t.Fatalf("StopProtocols(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopProtocols(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func setProtocolState(ctx context.Context, ate binding.ATE, state gosnappi.ProtocolStateStateEnum) ([]string, error) {
	api, err := fetchAPI(ctx, ate)
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
	debugger.ActionStarted(t, "Starting traffic on %s", o.ate)
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
	debugger.ActionStarted(t, "Stopping traffic on %s", o.ate)
	warns, err := setTransmitState(context.Background(), o.ate, gosnappi.TransmitStateState.STOP)
	if err != nil {
		t.Fatalf("StopTraffic(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopTraffic(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

func setTransmitState(ctx context.Context, ate binding.ATE, state gosnappi.TransmitStateStateEnum) ([]string, error) {
	api, err := fetchAPI(ctx, ate)
	if err != nil {
		return nil, err
	}
	resp, err := api.SetTransmitState(api.NewTransmitState().SetState(state))
	if err != nil {
		return nil, err
	}
	return resp.Warnings(), nil
}

func fetchAPI(ctx context.Context, ate binding.ATE) (gosnappi.GosnappiApi, error) {
	mu.Lock()
	defer mu.Unlock()
	api, ok := apis[ate]
	if !ok {
		var err error
		api, err = ate.DialOTG(ctx)
		if err != nil {
			return nil, err
		}
		apis[ate] = api
	}
	return api, nil
}

// Telemetry returns a telemetry path root for the device.
func (o *OTG) Telemetry() *device.DevicePath {
	root := device.DeviceRoot(o.ate.Name())
	root.PutCustomData(genutil.DefaultClientKey, func(ctx context.Context) (gpb.GNMIClient, error) {
		return fetchGNMI(ctx, o.ate)
	})
	return root
}

func fetchGNMI(ctx context.Context, ate binding.ATE) (gpb.GNMIClient, error) {
	mu.Lock()
	defer mu.Unlock()
	gnmi, ok := gnmis[ate]
	if !ok {
		var err error
		gnmi, err = ate.DialGNMI(ctx, grpc.WithBlock())
		if err != nil {
			return nil, err
		}
		gnmis[ate] = gnmi
	}
	return gnmi, nil
}
