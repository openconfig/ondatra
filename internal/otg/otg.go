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

// Package otg provides Open Traffic Generator API business logic.
package otg

import (
	"golang.org/x/net/context"
	"fmt"
	"sync"

	"google.golang.org/grpc"
	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

var (
	mu    sync.Mutex
	apis  = make(map[binding.ATE]gosnappi.GosnappiApi)
	gnmis = make(map[binding.ATE]gpb.GNMIClient)
)

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

// NewConfig constructs a new config for the specified ATE.
func NewConfig(ctx context.Context, ate binding.ATE) (gosnappi.Config, error) {
	api, err := fetchAPI(ctx, ate)
	if err != nil {
		return nil, err
	}
	return api.NewConfig(), nil
}

// FetchConfig fetches and returns the config on the specified ATE.
func FetchConfig(ctx context.Context, ate binding.ATE) (gosnappi.Config, error) {
	api, err := fetchAPI(ctx, ate)
	if err != nil {
		return nil, err
	}
	return api.GetConfig()
}

// PushConfig pushes a config to the specified ATE.
// The first return value are any non-fatal warnings encountered.
func PushConfig(ctx context.Context, ate binding.ATE, cfg gosnappi.Config) ([]string, error) {
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

// StartProtocols starts protocol on the specified ATE.
// The first return value are any non-fatal warnings encountered.
func StartProtocols(ctx context.Context, ate binding.ATE) ([]string, error) {
	return setProtocolState(ctx, ate, gosnappi.ProtocolStateState.START)
}

// StopProtocols stops protocol on the specified ATE.
// The first return value are any non-fatal warnings encountered.
func StopProtocols(ctx context.Context, ate binding.ATE) ([]string, error) {
	return setProtocolState(ctx, ate, gosnappi.ProtocolStateState.STOP)
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

// StartTraffic starts traffic on the specified ATE.
// The first return value are any non-fatal warnings encountered.
func StartTraffic(ctx context.Context, ate binding.ATE) ([]string, error) {
	return setTransmitState(ctx, ate, gosnappi.TransmitStateState.START)
}

// StopTraffic stops traffic on the specified ATE.
// The first return value are any non-fatal warnings encountered.
func StopTraffic(ctx context.Context, ate binding.ATE) ([]string, error) {
	return setTransmitState(ctx, ate, gosnappi.TransmitStateState.STOP)
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

// FetchGNMI returns a gNMI client for the specified ATE.
func FetchGNMI(ctx context.Context, ate binding.ATE) (gpb.GNMIClient, error) {
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
