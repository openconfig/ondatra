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

	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/gnmi"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ygnmi/ygnmi"
)

var (
	loopbackPrefixes = map[ondatra.Vendor]string{
		ondatra.ARISTA:  "Loopback",
		ondatra.CISCO:   "Loopback",
		ondatra.JUNIPER: "lo",
		ondatra.NOKIA:   "lo",
	}
	bundlePrefixes = map[ondatra.Vendor]string{
		ondatra.ARISTA:  "Port-Channel",
		ondatra.CISCO:   "Bundle-Ether",
		ondatra.JUNIPER: "ae",
		ondatra.NOKIA:   "lag",
	}
)

// LoopbackInterface returns the vendor-specific name of the loopback interface with
// the specified integer id.
func LoopbackInterface(t *testing.T, dut *ondatra.DUTDevice, id int) string {
	t.Helper()
	name, err := loopbackInterface(dut.Vendor(), id)
	if err != nil {
		t.Fatalf("LoopbackInterface(t, %s, %d): %v", dut.Name(), id, err)
	}
	return name
}

// NextBundleInterface returns the bundled interface with the lowest positive ID
// that, according to the device's telemetry, is not yet configured.
func NextBundleInterface(t *testing.T, dut *ondatra.DUTDevice) string {
	t.Helper()
	batch := gnmi.OCBatch()
	batch.AddPaths(gnmi.OC().InterfaceAny().Aggregation())
	name, err := nextBundleInterface(t, dut.Vendor(), gnmi.Lookup(t, dut, batch.State()))
	if err != nil {
		t.Fatalf("NextBundleInterface(t, %s): %v", dut.Name(), err)
	}
	return name
}

func loopbackInterface(vendor ondatra.Vendor, id int) (string, error) {
	if id < 0 {
		return "", fmt.Errorf("loopback interface cannot have negative number: %d", id)
	}
	prefix, ok := loopbackPrefixes[vendor]
	if !ok {
		return "", fmt.Errorf("no loopback interface prefix for DUT vendor %v", vendor)
	}
	return intfName(prefix, id), nil
}

func nextBundleInterface(t *testing.T, vendor ondatra.Vendor, val *ygnmi.Value[*oc.Root]) (string, error) {
	prefix, err := bundlePrefix(vendor)
	if err != nil {
		return "", err
	}
	var bundleIntfs map[string]*oc.Interface
	if root, present := val.Val(); present {
		bundleIntfs = root.Interface
	}
	for id := 1; ; id++ {
		bundleName := intfName(prefix, id)
		if _, ok := bundleIntfs[bundleName]; !ok {
			return bundleName, nil
		}
	}
}

func bundlePrefix(vendor ondatra.Vendor) (string, error) {
	prefix, ok := bundlePrefixes[vendor]
	if !ok {
		return "", fmt.Errorf("no bundle interface prefix for DUT vendor %v", vendor)
	}
	return prefix, nil
}

func intfName(prefix string, id int) string {
	return fmt.Sprintf("%s%d", prefix, id)
}
