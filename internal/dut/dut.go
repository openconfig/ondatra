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

// Package dut controls devices under test (DUT) for ONDATRA tests.
package dut

import (
	"golang.org/x/net/context"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/testbed"
	"github.com/openconfig/ondatra/binding/usererr"

	opb "github.com/openconfig/ondatra/proto"
)

// Truncate text at this length when String()-ing.
const configTrunc = 10

// Config stores the potential config text to push to the device.
type Config struct {
	VC   map[opb.Device_Vendor]ConfigProvider
	Open ConfigProvider
	Vars map[string]string
}

// ConfigProvider provide config text to push to the device.
type ConfigProvider interface {
	Get() (string, error)
}

// ConfigText is a literal config text string.
type ConfigText string

func (t ConfigText) String() string {
	s := string(t)
	if len(s) < configTrunc {
		return s
	}
	return s[:configTrunc] + "..."
}

// Get returns the config text string.
func (t ConfigText) Get() (string, error) {
	return string(t), nil
}

// ConfigFile is a file that contains config text.
type ConfigFile string

// Get returns the config text in the file.
func (f ConfigFile) Get() (string, error) {
	c, err := ioutil.ReadFile(string(f))
	if err != nil {
		return "", err
	}
	return string(c), nil
}

// PushConfig pushes config to a DUT.
func PushConfig(ctx context.Context, dut *binding.DUT, cfg *Config, append bool) error {
	prov, ok := cfg.VC[dut.Vendor]
	if !ok {
		if cfg.Open == nil {
			return errors.Errorf("no openconfig or vendor config for device %v", dut)
		}
		prov = cfg.Open
	}
	openConfig := prov == cfg.Open
	text, err := prov.Get()
	if err != nil {
		return errors.Wrapf(err, "error getting config from provider %v", prov)
	}
	config, err := interpolateConfig(dut, text, cfg.Vars)
	if err != nil {
		return err
	}
	opts := &binding.ConfigOptions{OpenConfig: openConfig, Append: append}
	return testbed.Bind().PushConfig(ctx, dut, config, opts)
}

// interpolateConfig substitutes templated variables in device config text.
// The following Go template functions are allowed in config:
// - {{ port "<portID>" }}: replaced with the physical port name
// - {{ secrets "<arg1>" "<arg2>" }}: left untouched, returned as-is
// - {{ var "<key>" }}: returns the value for the key in the vars map
func interpolateConfig(dut *binding.DUT, config string, vars map[string]string) (string, error) {
	funcMap := map[string]interface{}{
		"port": func(portID string) (string, error) {
			port, err := testbed.Port(dut.Dims, portID)
			if err != nil {
				return "", usererr.Wrap(err)
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
			v, ok := vars[key]
			if !ok {
				return "", usererr.New("No value for key %q in vars map", key)
			}
			return v, nil
		},
	}
	template, err := template.New(dut.Name).Funcs(funcMap).Parse(config)
	if err != nil {
		return "", usererr.Wrapf(err, "Invalid template in config: %q", config)
	}
	var b strings.Builder
	if err = template.Execute(&b, nil); err != nil {
		return "", usererr.Wrapf(err, "Invalid template in config: %q", config)
	}
	return b.String(), nil
}
