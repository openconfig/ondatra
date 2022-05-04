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
	"github.com/openconfig/ondatra/telemetry"
)

var (
	bundlePrefixes = map[ondatra.Vendor]string{
		ondatra.ARISTA:  "Port-Channel",
		ondatra.CISCO:   "Bundle-Ether",
		ondatra.JUNIPER: "ae",
	}
	vlanPrefixes = map[ondatra.Vendor]string{
		ondatra.ARISTA:  "Vlan",
		ondatra.CISCO:   "BVI",
		ondatra.JUNIPER: "irb.",
	}
)

// BundleInterface returns the vendor-specific name of the bundle interface with
// the specified integer id.
func BundleInterface(t *testing.T, dut *ondatra.DUTDevice, id int) string {
	t.Helper()
	name, err := bundleInterface(dut.Vendor(), id)
	if err != nil {
		t.Fatalf("BundleInterface(t, %s, %d): %v", dut.Name(), id, err)
	}
	return name
}

// VLANInterface returns the vendor-specific name of the VLAN interface with
// the specified integer id.
func VLANInterface(t *testing.T, dut *ondatra.DUTDevice, id int) string {
	t.Helper()
	name, err := vlanInterface(dut.Vendor(), id)
	if err != nil {
		t.Fatalf("VLANInterface(t, %s, %d): %v", dut.Name(), id, err)
	}
	return name
}

// NextBundleInterface returns the bundled interface with the lowest positive ID
// that, according to the device's telemetry, is not yet configured.
func NextBundleInterface(t *testing.T, dut *ondatra.DUTDevice) string {
	t.Helper()
	batch := dut.Telemetry().NewBatch()
	dut.Telemetry().InterfaceAny().Aggregation().Batch(t, batch)
	bundleIntfs := batch.Lookup(t).Val(t).Interface
	name, err := nextBundleInterface(dut.Vendor(), bundleIntfs)
	if err != nil {
		t.Fatalf("NextBundleInterface(t, %s): %v", dut.Name(), err)
	}
	return name
}

// NextVLANInterface returns the VLAN interface with the lowest positive ID
// that, according to the device's telemetry, is not yet configured.
func NextVLANInterface(t *testing.T, dut *ondatra.DUTDevice) string {
	t.Helper()
	batch := dut.Telemetry().NewBatch()
	dut.Telemetry().InterfaceAny().RoutedVlan().Batch(t, batch)
	vlanIntfs := batch.Lookup(t).Val(t).Interface
	name, err := nextVLANInterface(dut.Vendor(), vlanIntfs)
	if err != nil {
		t.Fatalf("NextVLANInterface(t, %s): %v", dut.Name(), err)
	}
	return name
}

func bundleInterface(vendor ondatra.Vendor, id int) (string, error) {
	if id < 0 {
		return "", fmt.Errorf("bundle interface cannot have negative number: %d", id)
	}
	prefix, err := bundlePrefix(vendor)
	if err != nil {
		return "", err
	}
	return intfName(prefix, id), nil
}

func vlanInterface(vendor ondatra.Vendor, id int) (string, error) {
	if id < 0 {
		return "", fmt.Errorf("VLAN interface cannot have negative number: %d", id)
	}
	prefix, err := vlanPrefix(vendor)
	if err != nil {
		return "", err
	}
	return intfName(prefix, id), nil
}

func nextBundleInterface(vendor ondatra.Vendor, bundleIntfs map[string]*telemetry.Interface) (string, error) {
	prefix, err := bundlePrefix(vendor)
	if err != nil {
		return "", err
	}
	for id := 1; ; id++ {
		bundleName := intfName(prefix, id)
		if _, ok := bundleIntfs[bundleName]; !ok {
			return bundleName, nil
		}
	}
}

func nextVLANInterface(vendor ondatra.Vendor, vlanIntfs map[string]*telemetry.Interface) (string, error) {
	prefix, err := vlanPrefix(vendor)
	if err != nil {
		return "", err
	}
	for id := 1; ; id++ {
		vlanName := intfName(prefix, id)
		if _, ok := vlanIntfs[vlanName]; !ok {
			return vlanName, nil
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

func vlanPrefix(vendor ondatra.Vendor) (string, error) {
	prefix, ok := vlanPrefixes[vendor]
	if !ok {
		return "", fmt.Errorf("no VLAN interface prefix for DUT vendor %v", vendor)
	}
	return prefix, nil
}

func intfName(prefix string, id int) string {
	return fmt.Sprintf("%s%d", prefix, id)
}
