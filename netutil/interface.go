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

package netutil

import (
	"fmt"
	"testing"

	"github.com/openconfig/entity-naming/entname"
	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/gnmi"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
)

// LoopbackInterface returns the vendor-specific name of the loopback interface with
// the specified zero-based index.
func LoopbackInterface(t *testing.T, dut *ondatra.DUTDevice, index int) string {
	t.Helper()
	name, err := loopbackInterface(dut, index)
	if err != nil {
		t.Fatalf("LoopbackInterface(t, %s, %d): %v", dut.Name(), index, err)
	}
	return name
}

func loopbackInterface(dut dutIntf, id int) (string, error) {
	dev, err := deviceParams(dut)
	if err != nil {
		return "", err
	}
	return entname.LoopbackInterface(dev, id)
}

// NextBundleInterface returns the aggregate interface with the lowest non-zero
// index that, according to the device's telemetry, is not yet configured.
// Deprecated: Use NextAggregateInterface instead.
func NextBundleInterface(t *testing.T, dut *ondatra.DUTDevice) string {
	t.Helper()
	batch := gnmi.OCBatch()
	batch.AddPaths(gnmi.OC().InterfaceAny().Aggregation())
	name, err := nextAggregateInterface(dut, gnmi.Lookup(t, dut, batch.State()))
	if err != nil {
		t.Fatalf("NextBundleInterface(t, %s): %v", dut.Name(), err)
	}
	return name
}

// NextAggregateInterface returns the aggregate interface with the lowest
// non-zero index that according to the dut's telemetry is not yet configured.
func NextAggregateInterface(t *testing.T, dut *ondatra.DUTDevice) string {
	t.Helper()
	batch := gnmi.OCBatch()
	batch.AddPaths(gnmi.OC().InterfaceAny().Aggregation())
	name, err := nextAggregateInterface(dut, gnmi.Lookup(t, dut, batch.State()))
	if err != nil {
		t.Fatalf("NextAggregateInterface(t, %s): %v", dut.Name(), err)
	}
	return name
}

func nextAggregateInterface(dut dutIntf, currAggsRoot *ygnmi.Value[*oc.Root]) (string, error) {
	dev, err := deviceParams(dut)
	if err != nil {
		return "", err
	}
	var aggIntfs map[string]*oc.Interface
	if root, present := currAggsRoot.Val(); present {
		aggIntfs = root.Interface
	}
	for index := 0; ; index++ {
		name, err := entname.AggregateInterface(dev, index)
		if err != nil {
			return "", err
		}
		if _, ok := aggIntfs[name]; !ok {
			return name, nil
		}
	}
}

var namingVendors map[ondatra.Vendor]entname.Vendor = map[ondatra.Vendor]entname.Vendor{
	ondatra.ARISTA:  entname.VendorArista,
	ondatra.CISCO:   entname.VendorCisco,
	ondatra.JUNIPER: entname.VendorJuniper,
	ondatra.NOKIA:   entname.VendorNokia,
}

func deviceParams(dut dutIntf) (*entname.DeviceParams, error) {
	ov, ok := namingVendors[dut.Vendor()]
	if !ok {
		return nil, fmt.Errorf("DUT vendor %v not supported by OC entity naming library", dut.Vendor())
	}
	return &entname.DeviceParams{
		Vendor:        ov,
		HardwareModel: dut.Model(),
	}, nil
}

// dutIntf is an interface to a DUT used to enable dependency injection.
type dutIntf interface {
	Vendor() ondatra.Vendor
	Model() string
}
