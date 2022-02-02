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

package ondatra

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
)

// OTG provides the Open Traffic Generator API to an ATE.
type OTG struct {
	ate *binding.ATE
	api gosnappi.GosnappiApi
}

func (o *OTG) String() string {
	return fmt.Sprintf("{ate: %s}", o.ate)
}

// NewConfig creates a new OTG config.
func (o *OTG) NewConfig() gosnappi.Config {
	return o.api.NewConfig()
}

// PushConfig pushes config to the ATE.
func (o *OTG) PushConfig(t testing.TB, cfg gosnappi.Config) {
	t.Helper()
	const locFormat = "service-ate%d.ceos-simple.svc.cluster.local:5555+service-ate%d.ceos-simple.svc.cluster.local:50071"
	for _, p := range cfg.Ports().Items() {
		portNum, err := strconv.Atoi(p.Name()[1:])
		if err != nil {
			t.Fatal(err)
		}
		p.SetLocation(fmt.Sprintf(locFormat, portNum, portNum))
	}
	logAction(t, "Pushing config to %s", o.ate)
	resp, err := o.api.SetConfig(cfg)
	if err != nil {
		t.Fatalf("PushConfig(t) on %s: %v", o, err)
	}
	if warns := resp.Warnings(); len(warns) > 0 {
		t.Logf("PushConfig(t) on %s non-fatal warnings: %v", o, warns)
	}
}

// FetchConfig fetches config from the ATE.
func (o *OTG) FetchConfig(t testing.TB) gosnappi.Config {
	t.Helper()
	logAction(t, "Fetching config from %s", o.ate)
	cfg, err := o.api.GetConfig()
	if err != nil {
		t.Fatalf("FetchConfig(t) from %s: %v", o, err)
	}
	return cfg
}

// StartProtocols starts protocols on the ATE.
func (o *OTG) StartProtocols(t testing.TB) {
	logAction(t, "Starting protocols on %s", o.ate)
	started := o.api.NewProtocolState().SetState(gosnappi.ProtocolStateState.START)
	resp, err := o.api.SetProtocolState(started)
	if err != nil {
		t.Fatalf("StartProtocols(t) from %s: %v", o, err)
	}
	if warns := resp.Warnings(); len(warns) > 0 {
		t.Logf("StartProtocols(t) on %s non-fatal warnings: %v", o, warns)
	}
}

// StopProtocols stops protocols on the ATE.
func (o *OTG) StopProtocols(t testing.TB) {
	logAction(t, "Stopping rotocols on %s", o.ate)
	stopped := o.api.NewProtocolState().SetState(gosnappi.ProtocolStateState.STOP)
	resp, err := o.api.SetProtocolState(stopped)
	if err != nil {
		t.Fatalf("StopProtocols(t) from %s: %v", o, err)
	}
	if warns := resp.Warnings(); len(warns) > 0 {
		t.Logf("StopProtocols(t) on %s non-fatal warnings: %v", o, warns)
	}
}

// StartTraffic starts traffic on the ATE.
func (o *OTG) StartTraffic(t testing.TB) {
	logAction(t, "Starting traffic on %s", o.ate)
	started := o.api.NewTransmitState().SetState(gosnappi.TransmitStateState.START)
	resp, err := o.api.SetTransmitState(started)
	if err != nil {
		t.Fatalf("StartTraffic(t) from %s: %v", o, err)
	}
	if warns := resp.Warnings(); len(warns) > 0 {
		t.Logf("StartTraffic(t) on %s non-fatal warnings: %v", o, warns)
	}
}

// StopTraffic stops traffic on the ATE.
func (o *OTG) StopTraffic(t testing.TB) {
	logAction(t, "Stopping traffic on %s", o.ate)
	stopped := o.api.NewTransmitState().SetState(gosnappi.TransmitStateState.STOP)
	resp, err := o.api.SetTransmitState(stopped)
	if err != nil {
		t.Fatalf("StopTraffic(t) from %s: %v", o, err)
	}
	if warns := resp.Warnings(); len(warns) > 0 {
		t.Logf("StopTraffic(t) on %s non-fatal warnings: %v", o, warns)
	}
}
