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
	"testing"

	"github.com/open-traffic-generator/snappi/gosnappi"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/debugger"
	"github.com/openconfig/ondatra/internal/otg"
)

// OTG provides the Open Traffic Generator API to an ATE.
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
	cfg, err := otg.NewConfig(o.ate)
	if err != nil {
		t.Fatalf("NewConfig(t) on %s: %v", o.ate, err)
	}
	return cfg
}

// PushConfig pushes config to the ATE.
func (o *OTG) PushConfig(t testing.TB, cfg gosnappi.Config) {
	t.Helper()
	debugger.ActionStarted(t, "Pushing config to %s", o.ate)
	warns, err := otg.PushConfig(o.ate, cfg)
	if err != nil {
		t.Fatalf("PushConfig(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("PushConfig(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// FetchConfig fetches config from the ATE.
func (o *OTG) FetchConfig(t testing.TB) gosnappi.Config {
	t.Helper()
	debugger.ActionStarted(t, "Fetching config from %s", o.ate)
	cfg, err := otg.FetchConfig(o.ate)
	if err != nil {
		t.Fatalf("FetchConfig(t) on %s: %v", o.ate, err)
	}
	return cfg
}

// StartProtocols starts protocols on the ATE.
func (o *OTG) StartProtocols(t testing.TB) {
	t.Helper()
	debugger.ActionStarted(t, "Starting protocols on %s", o.ate)
	warns, err := otg.StartProtocols(o.ate)
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
	warns, err := otg.StopProtocols(o.ate)
	if err != nil {
		t.Fatalf("StopProtocols(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopProtocols(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}

// StartTraffic starts traffic on the ATE.
func (o *OTG) StartTraffic(t testing.TB) {
	t.Helper()
	debugger.ActionStarted(t, "Starting traffic on %s", o.ate)
	warns, err := otg.StartTraffic(o.ate)
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
	warns, err := otg.StopTraffic(o.ate)
	if err != nil {
		t.Fatalf("StopTraffic(t) on %s: %v", o.ate, err)
	}
	if len(warns) > 0 {
		t.Logf("StopTraffic(t) on %s non-fatal warnings: %v", o.ate, warns)
	}
}
