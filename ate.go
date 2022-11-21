// Copyright 2019 Google LLC
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
	"golang.org/x/net/context"
	"testing"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/binding/ixweb"
	"github.com/openconfig/ondatra/internal/events"
	"github.com/openconfig/ondatra/internal/rawapis"
	"github.com/openconfig/ondatra/otg"
)

// ATEDevice is an automated test equipment.
type ATEDevice struct {
	*Device
	otg *otg.OTG
}

// OTG returns a handle to the OTG API.
func (a *ATEDevice) OTG() *otg.OTG {
	return otg.New(a.res.(binding.ATE))
}

// Topology returns a handle to the IxNetwork Topology API.
func (a *ATEDevice) Topology() *Topology {
	return &Topology{a.res.(binding.ATE)}
}

// Traffic returns a handle to the IxNetwork Traffic API.
func (a *ATEDevice) Traffic() *Traffic {
	return &Traffic{a.res.(binding.ATE)}
}

// Actions returns a handle to the IxNetwork Actions API.
func (a *ATEDevice) Actions() *Actions {
	return &Actions{a.res.(binding.ATE)}
}

// RawAPIs returns a handle to raw protocol APIs on the ATE.
func (a *ATEDevice) RawAPIs() *RawATEAPIs {
	return &RawATEAPIs{ate: a.res.(binding.ATE)}
}

// RawATEAPIs provides access to raw ATE protocol APIs.
type RawATEAPIs struct {
	ate binding.ATE
}

// IxNetwork returns the raw IxNetwork session for the ATE.
// TODO(team): Add unit tests once raw APIs is factored out into its own package.
func (r *RawATEAPIs) IxNetwork(t testing.TB) *ixweb.Session {
	t.Helper()
	t = events.ActionStarted(t, "Fetching IxNetwork session for %s", r.ate)
	ixnet, err := rawapis.FetchIxNetwork(context.Background(), r.ate)
	if err != nil {
		t.Fatalf("IxNetwork(t) on %v: %v", r.ate, err)
	}
	return ixnet.Session
}
