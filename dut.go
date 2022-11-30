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
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/config"
	"github.com/openconfig/ondatra/operations"
	"github.com/openconfig/ondatra/raw"
)

// DUTDevice is a device under test.
type DUTDevice struct {
	*Device
}

// Config returns a handle to the DUT configuration API.
func (d *DUTDevice) Config() *Config {
	return &Config{
		dut: d.res.(binding.DUT),
	}
}

// Config is the DUT configuration API.
type Config struct {
	dut binding.DUT
}

// New returns an empty DUT configuration.
func (c *Config) New() *config.VendorConfig {
	return config.NewVendorConfig(c.dut)
}

// Operations returns a handle to the DUT operations API.
func (d *DUTDevice) Operations() *operations.Operations {
	return operations.New(d.res.(binding.DUT))
}

// RawAPIs returns a handle to raw protocol APIs on the DUT.
func (d *DUTDevice) RawAPIs() *raw.DUTAPIs {
	return raw.NewDUTAPIs(d.res.(binding.DUT))
}
