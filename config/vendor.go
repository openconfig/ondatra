// Copyright 2022 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package config provides an API for setting native config on DUTs via
// vendor-specific (non-gNMI) protocols.
//
// We use the following example snippet to understand various aspects of the
// Config API.
//
//	dut.Config().New().
//		WithAristaText(`
//	interface {{ port "port1" }}
//	 no switchport
//	 ip address 192.0.2.1/30
//	 ip address {{ var "ipv4addr" }}
//	 ipv6 address 2001:DB8::1/126
//	 load-interval 30`).
//		WithCiscoText(`
//	interface {{ port "port1" }}
//	 ipv4 address {{ var "ipv4addr" }}
//	 ipv6 address 2001:DB8::1/126
//	 load-interval 30`).
//		WithJuniperFile(`path/to/juniper.config`).
//		WithVarValue("ipv4addr", computeV4Address()).
//		Push(t)
//
// # Per-Vendor Configs
//
// In the example above, one of three possible native configs are pushed to the
// device depending upon the vendor of the device at runtime. If the vendor is
// not one of those three, the [VendorConfig.Push] function fails fatally. The
// `With<Vendor>Text` methods will push a vendor-specific text string and the
// `With<Vendor>File` methods will push the text in a vendor-specific file.
//
// If the test has already pre-computed the native config in a way that has
// taken the vendor into account, perhaps by calling an external
// config-generation service, the test can push that config without regard to
// the vendor using the functions [VendorConfig.WithText] or
// [VendorConfig.WithFile].
//
// # Templated Variables
//
// As shown, the syntax for the config allows templated port names like `{{ port
// "port1" }}`. That placeholder will be replaced at runtime with the name given
// to `port1`, where `port1` is the ID of a port in the testbed file. The syntax
// may also contain instances of `{{ var "<key>" }}`, which will be replaced at
// runtime by calling `WithVarValue` or `WithVarMap` on the config object.
//
// # Push vs Append
//
// The effect of Push is to reset the device back to its original config (the
// config the device had when it was reserved) and then to append the specified
// config. As a result, test cases that use native config don't usually need to
// reset the device config to be hermetic; it suffices to just start the test
// case with a call to Push. If a test wants only to append config and skip
// the reset behavior, it should call [VendorConfig.Append] instead.
package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"

	"golang.org/x/net/context"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/events"
	"github.com/openconfig/ondatra/internal/testbed"

	opb "github.com/openconfig/ondatra/proto"
)

// Truncate text at this length when String()-ing.
const configTrunc = 10

// NewVendorConfig returns a new vendor configuration for a DUT.
// Tests must not call this method directly.
func NewVendorConfig(dut binding.DUT) *VendorConfig {
	return &VendorConfig{
		dut:       dut,
		perVendor: make(map[opb.Device_Vendor]configProvider),
		vars:      make(map[string]string),
	}
}

// VendorConfig is a vendor configuration for a DUT.
type VendorConfig struct {
	dut       binding.DUT
	allVendor configProvider
	perVendor map[opb.Device_Vendor]configProvider
	vars      map[string]string
}

func (c *VendorConfig) String() string {
	return fmt.Sprintf("VendorConfig%+v", *c)
}

// WithText sets the config to be pushed regardless of the DUT vendor.
// This should only be used when the DUT vendor was already taken into account
// in the generation of the config and only when no per-vendor configs are set.
func (c *VendorConfig) WithText(text string) *VendorConfig {
	c.allVendor = configText(text)
	return c
}

// WithAristaText sets the config to be pushed if the DUT vendor is Arista.
func (c *VendorConfig) WithAristaText(text string) *VendorConfig {
	c.perVendor[opb.Device_ARISTA] = configText(text)
	return c
}

// WithCiscoText sets the config to be pushed if the DUT vendor is Cisco.
func (c *VendorConfig) WithCiscoText(text string) *VendorConfig {
	c.perVendor[opb.Device_CISCO] = configText(text)
	return c
}

// WithJuniperText sets the config to be pushed if the DUT vendor is Juniper.
func (c *VendorConfig) WithJuniperText(text string) *VendorConfig {
	c.perVendor[opb.Device_JUNIPER] = configText(text)
	return c
}

// WithFile sets the config to be pushed regardless of the DUT vendor.
// This should only be used when the DUT vendor was already taken into account
// in the generation of the config and only when no per-vendor configs are set.
func (c *VendorConfig) WithFile(path string) *VendorConfig {
	c.allVendor = configFile(path)
	return c
}

// WithAristaFile sets the config to be pushed if the DUT vendor is Arista.
func (c *VendorConfig) WithAristaFile(path string) *VendorConfig {
	c.perVendor[opb.Device_ARISTA] = configFile(path)
	return c
}

// WithCiscoFile sets the config to be pushed if the DUT vendor is Cisco.
func (c *VendorConfig) WithCiscoFile(path string) *VendorConfig {
	c.perVendor[opb.Device_CISCO] = configFile(path)
	return c
}

// WithJuniperFile sets the config to be pushed if the DUT vendor is Juniper.
func (c *VendorConfig) WithJuniperFile(path string) *VendorConfig {
	c.perVendor[opb.Device_JUNIPER] = configFile(path)
	return c
}

// WithVarValue replaces each occurrence of {{ var "key" }} in the pushed config
// with the specified value.
func (c *VendorConfig) WithVarValue(key, value string) *VendorConfig {
	c.vars[key] = value
	return c
}

// WithVarMap sets the map used to replace each occurrence of {{ var "key" }} in the pushed config.
func (c *VendorConfig) WithVarMap(m map[string]string) *VendorConfig {
	c.vars = m
	return c
}

// Push resets the device to its base config and appends the specified config.
func (c *VendorConfig) Push(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Pushing config to %s", c.dut)
	if err := c.pushConfig(context.Background(), true); err != nil {
		t.Fatalf("Push(t) on %s: %v", c.dut, err)
	}
}

// Append appends the specified config to the current config.
func (c *VendorConfig) Append(t testing.TB) {
	t.Helper()
	t = events.ActionStarted(t, "Appending config to %s", c.dut)
	if err := c.pushConfig(context.Background(), false); err != nil {
		t.Fatalf("Append(t) on %s: %v", c.dut, err)
	}
}

func (c *VendorConfig) pushConfig(ctx context.Context, reset bool) error {
	if c.allVendor != nil && len(c.perVendor) > 0 {
		return errors.New("cannot specify both all-vendor and per-vendor config")
	}
	var prov configProvider
	if c.allVendor != nil {
		prov = c.allVendor
	} else if pv, ok := c.perVendor[c.dut.Vendor()]; ok {
		prov = pv
	} else {
		return fmt.Errorf("no config specified for device %v", c.dut)
	}
	text, err := prov.get()
	if err != nil {
		return fmt.Errorf("error getting config from provider %v: %w", prov, err)
	}
	config, err := c.interpolate(text)
	if err != nil {
		return err
	}
	return c.dut.PushConfig(ctx, config, reset)
}

// interpolateConfig substitutes templated variables in device config text.
// The following Go template functions are allowed in config:
// - {{ port "<portID>" }}: replaced with the physical port name
// - {{ secrets "<arg1>" "<arg2>" }}: left untouched, returned as-is
// - {{ var "<key>" }}: returns the value for the key in the vars map
func (c *VendorConfig) interpolate(config string) (string, error) {
	funcMap := map[string]any{
		"port": func(portID string) (string, error) {
			port, err := testbed.Port(c.dut, portID)
			if err != nil {
				return "", err
			}
			return port.Name, nil
		},
		// "secrets" function should be a noop
		"secrets": func(secrets ...string) string {
			var args []string
			for _, s := range secrets {
				args = append(args, fmt.Sprintf("%q", s))
			}
			return fmt.Sprintf("{{ secrets %s }}", strings.Join(args, " "))
		},
		"var": func(key string) (string, error) {
			v, ok := c.vars[key]
			if !ok {
				return "", fmt.Errorf("no value for key %q in vars map", key)
			}
			return v, nil
		},
	}
	template, err := template.New(c.dut.Name()).Funcs(funcMap).Parse(config)
	if err != nil {
		return "", fmt.Errorf("invalid template in config %q: %w", config, err)
	}
	var b strings.Builder
	if err = template.Execute(&b, nil); err != nil {
		return "", fmt.Errorf("invalid template in config %q: %w", config, err)
	}
	return b.String(), nil
}

type configProvider interface {
	get() (string, error)
}

type configText string

func (t configText) String() string {
	s := string(t)
	if len(s) < configTrunc {
		return s
	}
	return s[:configTrunc] + "..."
}

func (t configText) get() (string, error) {
	return string(t), nil
}

type configFile string

func (f configFile) get() (string, error) {
	c, err := os.ReadFile(string(f))
	if err != nil {
		return "", err
	}
	return string(c), nil
}
