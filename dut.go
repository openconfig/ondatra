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
	"fmt"
	"testing"

	"github.com/openconfig/ondatra/config/device"
	"github.com/openconfig/ondatra/internal/binding"
	"github.com/openconfig/ondatra/internal/cli"
	"github.com/openconfig/ondatra/internal/console"
	"github.com/openconfig/ondatra/internal/dut"
	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ondatra/internal/operations"
	"github.com/openconfig/ondatra/internal/p4rt"
	"github.com/openconfig/ondatra/internal/reservation"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	opb "github.com/openconfig/ondatra/proto"
	p4pb "github.com/p4lang/p4runtime/go/p4/v1"

)

// DUTDevice is a device under test.
type DUTDevice struct {
	*Device
}

// Config returns a handle to the DUT configuration API.
func (d *DUTDevice) Config() *Config {
	dev := device.DeviceRoot(d.ID())
	// TODO: Add field to root node in ygot instead of using custom data.
	dev.PutCustomData(genutil.DefaultClientKey, d.Device.clientFn)
	return &Config{
		dut:        d.res.(*reservation.DUT),
		DevicePath: dev,
	}
}

// Config is the DUT configuration API.
type Config struct {
	dut *reservation.DUT
	*device.DevicePath
}

// New returns an empty DUT configuration.
func (a *Config) New() *DUTConfig {
	return &DUTConfig{
		dut: a.dut,
		cfg: &dut.Config{
			VC:   make(map[opb.Device_Vendor]dut.ConfigProvider),
			Vars: make(map[string]string),
		},
	}
}

// DUTConfig is a configuration of a device under test.
type DUTConfig struct {
	dut *reservation.DUT
	cfg *dut.Config
}

func (c *DUTConfig) String() string {
	return fmt.Sprintf("DUTConfig%+v", *c)
}

// WithAristaText sets the config to be pushed to an Arista device.
func (c *DUTConfig) WithAristaText(text string) *DUTConfig {
	c.cfg.VC[opb.Device_ARISTA] = dut.ConfigText(text)
	return c
}

// WithCiscoText sets the config to be pushed to a Cisco device.
func (c *DUTConfig) WithCiscoText(text string) *DUTConfig {
	c.cfg.VC[opb.Device_CISCO] = dut.ConfigText(text)
	return c
}

// WithJuniperText sets the config to be pushed to a Juniper device.
func (c *DUTConfig) WithJuniperText(text string) *DUTConfig {
	c.cfg.VC[opb.Device_JUNIPER] = dut.ConfigText(text)
	return c
}

// WithOpenConfigText sets the openconfig to be pushed to a device
// if no vendor-specific configuration has been set.
func (c *DUTConfig) WithOpenConfigText(text string) *DUTConfig {
	c.cfg.Open = dut.ConfigText(text)
	return c
}

// WithAristaFile sets the config to be pushed to an Arista device.
func (c *DUTConfig) WithAristaFile(path string) *DUTConfig {
	c.cfg.VC[opb.Device_ARISTA] = dut.ConfigFile(path)
	return c
}

// WithCiscoFile sets the config to be pushed to a Cisco device.
func (c *DUTConfig) WithCiscoFile(path string) *DUTConfig {
	c.cfg.VC[opb.Device_CISCO] = dut.ConfigFile(path)
	return c
}

// WithJuniperFile sets the config to be pushed to a Juniper device.
func (c *DUTConfig) WithJuniperFile(path string) *DUTConfig {
	c.cfg.VC[opb.Device_JUNIPER] = dut.ConfigFile(path)
	return c
}

// WithOpenConfigFile sets the openconfig to be pushed to a device
// if no vendor-specific configuration has been set.
func (c *DUTConfig) WithOpenConfigFile(path string) *DUTConfig {
	c.cfg.Open = dut.ConfigFile(path)
	return c
}

// WithVarValue replaces each occurrence of {{ var "key" }} in the pushed config
// with the specified value.
func (c *DUTConfig) WithVarValue(key, value string) *DUTConfig {
	c.cfg.Vars[key] = value
	return c
}

// WithVarMap sets the map used to replace each occurrence of {{ var "key" }} in the pushed config.
func (c *DUTConfig) WithVarMap(m map[string]string) *DUTConfig {
	c.cfg.Vars = m
	return c
}

// Push replaces the config on the device with the specified config, prepended with the device's base config.
// It pushes vendor config to the device if it has been set; otherwise, it attempts to push openconfig.
// If neither vendor config nor openconfig has been specified, it fails the test.
func (c *DUTConfig) Push(t testing.TB) {
	t.Helper()
	logAction(t, "Pushing config to %s", c.dut)
	if err := dut.PushConfig(context.Background(), c.dut, c.cfg, false); err != nil {
		t.Fatalf("Push(t) on %s: %v", c.dut, err)
	}
}

// Append appends the specific config to the device's current config.
// It pushes vendor config to the device if it has been set; otherwise, it attempts to push openconfig.
// If neither vendor config nor openconfig has been specified, it fails the test.
func (c *DUTConfig) Append(t testing.TB) {
	t.Helper()
	logAction(t, "Appending config to %s", c.dut)
	if err := dut.PushConfig(context.Background(), c.dut, c.cfg, true); err != nil {
		t.Fatalf("Append(t) on %s: %v", c.dut, err)
	}
}

// RawAPIs returns a handle to raw protocol APIs on the DUT.
func (d *DUTDevice) RawAPIs() *RawAPIs {
	return &RawAPIs{dut: d.res.(*reservation.DUT)}
}

// RawAPIs provides access to raw DUT protocols APIs.
type RawAPIs struct {
	dut *reservation.DUT
}

// GNMI provides access to either a new or default gNMI client.
func (r *RawAPIs) GNMI() *GNMIAPI {
	return &GNMIAPI{dut: r.dut}
}

// GNOI provides access to either a new or default gNOI client.
func (r *RawAPIs) GNOI() *GNOIAPI {
	return &GNOIAPI{dut: r.dut}
}

// GNMIAPI provides access for creating a default or new gNMI client on the DUT.
type GNMIAPI struct {
	dut *reservation.DUT
}

// GNOIAPI provides access for creating a default or new gNOI client on the DUT.
type GNOIAPI struct {
	dut *reservation.DUT
}

// New returns a new gNMI client on the DUT. This client will not be cached.
func (g *GNMIAPI) New(t testing.TB) gpb.GNMIClient {
	t.Helper()
	logAction(t, "Creating gNMI client for %s", g.dut)
	gnmi, err := newGNMI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("GNMI(t) on %v: %v", g.dut, err)
	}
	return gnmi
}

// Default returns the default gNMI client for the DUT.
func (g *GNMIAPI) Default(t testing.TB) gpb.GNMIClient {
	t.Helper()
	logAction(t, "Fetching gNMI client for %s", g.dut)
	gnmi, err := fetchGNMI(context.Background(), g.dut, nil)
	if err != nil {
		t.Fatalf("GNMI(t) on %v: %v", g.dut, err)
	}
	return gnmi
}

// New returns a new gNOI client on the DUT.
func (g *GNOIAPI) New(t testing.TB) GNOI {
	t.Helper()
	logAction(t, "Creating gNOI client for %s", g.dut)
	bgnoi, err := operations.NewGNOI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("GNOI(t) on %v: %v", g.dut, err)
	}
	return bgnoi
}

// Default returns the default gNOI client for the DUT.
func (g *GNOIAPI) Default(t testing.TB) GNOI {
	t.Helper()
	logAction(t, "Fetching gNOI client for %s", g.dut)
	bgnoi, err := operations.FetchGNOI(context.Background(), g.dut)
	if err != nil {
		t.Fatalf("GNOI(t) on %v: %v", g.dut, err)
	}
	return bgnoi
}

// GNOI stores gNOI clients to a DUT.
type GNOI interface {
	// Embed an unexported interface that wraps binding.GNOIClients,
	// so as to not expose the binding.GNOIClients instance directly.
	privateGNOI
}

type privateGNOI interface {
	binding.GNOIClients
}

// P4RT returns a P4RT client on the DUT.
func (r *RawAPIs) P4RT(t testing.TB) p4pb.P4RuntimeClient {
	t.Helper()
	logAction(t, "Creating P4RT client for %s", r.dut)
	p4rtClient, err := p4rt.NewP4RT(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to create P4RT client on %v: %v", r.dut, err)
	}
	return p4rtClient
}

// StreamClient provides the interface for streaming IO to DUT.
type StreamClient interface {
	// Embed an unexported interface that wraps binding.GNOIClients,
	// so as to not expose the binding.GNOIClients instance directly.
	privateStreamClient
}

type privateStreamClient interface {
	binding.StreamClient
}

// CLI returns a streaming CLI client on the DUT.
func (r *RawAPIs) CLI(t testing.TB) StreamClient {
	t.Helper()
	logAction(t, "Creating CLI client for %s", r.dut)
	c, err := cli.NewCLI(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to create CLI client on %v: %v", r.dut, err)
	}
	return c
}

// Console returns a transactional CLI client on the DUT.
func (r *RawAPIs) Console(t testing.TB) StreamClient {
	t.Helper()
	logAction(t, "Creating console client for %s", r.dut)
	c, err := console.NewConsole(context.Background(), r.dut)
	if err != nil {
		t.Fatalf("Failed to create console client on %v: %v", r.dut, err)
	}
	return c
}
