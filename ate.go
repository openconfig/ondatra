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
	"testing"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/debugger"
	"github.com/openconfig/ondatra/internal/testbed"
	"golang.org/x/net/context"
)

// ATEDevice is an automated test equipment.
type ATEDevice struct {
	*Device
	otg *OTG
}

// OTG returns a handle to the OTG API.
func (a *ATEDevice) OTG(t testing.TB) *OTG {
	if a.otg == nil {
		api, err := testbed.Bind().DialOTG(context.Background(), a.res.(*binding.ATE))
		if err != nil {
			t.Error(err)
		}
		a.otg = &OTG{ate: a.res.(*binding.ATE), api: api}
	}
	return a.otg
}

// Topology returns a handle to the topology API.
func (a *ATEDevice) Topology() *Topology {
	return &Topology{a.res.(*binding.ATE)}
}

// Traffic returns a handle to the traffic API.
func (a *ATEDevice) Traffic() *Traffic {
	return &Traffic{a.res.(*binding.ATE)}
}

// RawATEAPIs returns a handle to raw protocol APIs on the ATE.
func (a *ATEDevice) RawATEAPIs() *RawATEAPIs {
	return &RawATEAPIs{ate: a.res.(*binding.ATE)}
}

// RawAPIs provides access to raw DUT protocols APIs.
type RawATEAPIs struct {
	ate *binding.ATE
}

// ATEGNMI provides access to either a new or default gNMI client.
func (r *RawATEAPIs) ATEGNMI() *GNMIATEAPI {
	return &GNMIATEAPI{ate: r.ate}
}

// GNMIATEAPI provides access for creating a default or new gNMI client on the ATE.
type GNMIATEAPI struct {
	ate *binding.ATE
}

// New returns a new gNMI client on the ATE. This client will not be cached.
func (g *GNMIATEAPI) New(t testing.TB) gpb.GNMIClient {
	t.Helper()
	debugger.ActionStarted(t, "Creating gNMI client for %s", g.ate)
	gnmi, err := newGNMI(context.Background(), g.ate)
	if err != nil {
		t.Fatalf("GNMI(t) on %v: %v", g.ate, err)
	}
	return gnmi
}

// Default returns the default gNMI client for the ate.
func (g *GNMIATEAPI) Default(t testing.TB) gpb.GNMIClient {
	t.Helper()
	debugger.ActionStarted(t, "Fetching gNMI client for %s", g.ate)
	gnmi, err := fetchGNMI(context.Background(), g.ate, nil)
	if err != nil {
		t.Fatalf("GNMI(t) on %v: %v", g.ate, err)
	}
	return gnmi
}
