// Copyright 2021 Google Inc.
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

package ixconfig

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestMarshalUnmarshalConfig tests that a sequence of (unmarshal from JSON -> transform -> marshal to JSON) produces the
// expected JSON representation.
func TestMarshalUnmarshalConfig(t *testing.T) {
	tests := []struct {
		desc         string
		fromJSONFile string
		// An in-place modification of an IxNetwork config object
		transform    func(*Ixnetwork)
		wantJSONFile string
	}{{
		desc:         "No transformations applied.",
		fromJSONFile: "ixia_full_conf_modified.json",
		transform:    func(_ *Ixnetwork) {},
		wantJSONFile: "ixia_full_conf_modified.json",
	}, {
		desc:         "No change from updating XPaths w/o any config changes.",
		fromJSONFile: "ixia_full_conf_modified.json",
		transform:    (*Ixnetwork).updateAllXPathsAndRefs,
		wantJSONFile: "ixia_full_conf_modified.json",
	}, {
		desc:         "Fix missing XPaths.",
		fromJSONFile: "missing_xpath_conf.json",
		transform:    (*Ixnetwork).updateAllXPathsAndRefs,
		wantJSONFile: "filled_xpath_conf.json",
	}}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			// Fetch JSON configs from files.
			fromJSON, err := os.ReadFile(filepath.Join("testdata", tc.fromJSONFile))
			if err != nil {
				t.Fatalf("Could not load file %s: %v", tc.fromJSONFile, err)
			}
			wantJSON, err := os.ReadFile(filepath.Join("testdata", tc.wantJSONFile))
			if err != nil {
				t.Fatalf("Could not load file %s: %v", tc.wantJSONFile, err)
			}

			// Unmarshal, apply transformations, and marshal again into JSON. The latter step is added to
			// catch issues with custom Marshal/Unmarshal implementations (eg. for XPaths).
			cfg := &Ixnetwork{}
			if err := json.Unmarshal([]byte(fromJSON), cfg); err != nil {
				t.Fatalf("Could not unmarshal JSON to an IxNetwork config: %v", err)
			}
			tc.transform(cfg)
			jsonBytes, err := json.Marshal(cfg)
			if err != nil {
				t.Fatalf("Could not marshal IxNetwork config to JSON: %v", err)
			}

			var got, want map[string]any
			if err := json.Unmarshal(jsonBytes, &got); err != nil {
				t.Fatalf("Could not unmarshal transformed JSON to a map: %v", err)
			}
			if err := json.Unmarshal([]byte(wantJSON), &want); err != nil {
				t.Fatalf("Could not unmarshal expected JSON to a map: %v", err)
			}

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("Unexpected JSON representation, diff: (-got +want)\n%s.", diff)
			}
		})
	}
}

// TestCopy tests that Copy produces a valid deep copy of IxNetwork config objects.
func TestCopy(t *testing.T) {
	orig := &Ixnetwork{
		Topology: []*Topology{{
			Name:   String("someName"),
			Vports: []string{"/vports[1]"},
			DeviceGroup: []*TopologyDeviceGroup{{
				Enabled: MultivalueTrue(),
			}},
		}},
		Vport: []*Vport{{Name: String("somePort")}},
	}
	orig.Topology[0].SetVportsRefs([]IxiaCfgNode{orig.Vport[0]})

	cp := orig.Copy()
	if diff := cmp.Diff(orig, cp, cmp.Exporter(func(reflect.Type) bool { return true })); diff != "" {
		t.Fatalf("Incorrect data in copied config, diff (-got +want):\n%s", diff)
	}
	// Check that the copy actually points to different subconfigs than the original.
	testPtrDiff := func(a, b any, itemDesc string) {
		if reflect.ValueOf(a).Pointer() == reflect.ValueOf(b).Pointer() {
			t.Errorf("Pointers for config object %q are the same", itemDesc)
		}
	}
	testPtrDiff(orig.Topology, cp.Topology, "topology list")
	origTopo := orig.Topology[0]
	cpTopo := cp.Topology[0]
	testPtrDiff(origTopo, cpTopo, "topology")
	testPtrDiff(origTopo.Name, cpTopo.Name, "topology name")
	testPtrDiff(origTopo.Vports, cpTopo.Vports, "topology port list")
	testPtrDiff(origTopo.vportsRefs[0], cpTopo.vportsRefs[0], "topology referenced port")
	origDg := origTopo.DeviceGroup[0]
	cpDg := cpTopo.DeviceGroup[0]
	testPtrDiff(origDg, cpDg, "device group")
	origMV := origDg.Enabled
	cpMV := cpDg.Enabled
	testPtrDiff(origMV, cpMV, "enabled multivalue")
	testPtrDiff(origMV.SingleValue, cpMV.SingleValue, "multivalue single value")

	// Test the Vport/reference in the topology are the same object in the copied config.
	if reflect.ValueOf(cpTopo.vportsRefs[0]).Pointer() != reflect.ValueOf(cp.Vport[0]).Pointer() {
		t.Error("Configured vport and vport referenced in topology are different objects")
	}
}

// TestRefUpdate tests that configured config references are correctly updated.
func TestRefUpdate(t *testing.T) {
	cfg := &Ixnetwork{
		Topology: []*Topology{{
			Name: String("someName"),
			DeviceGroup: []*TopologyDeviceGroup{{
				Enabled: MultivalueTrue(),
			}},
		}},
		Vport: []*Vport{{Name: String("somePort")}},
	}
	topo := cfg.Topology[0]
	vport := cfg.Vport[0]
	topo.SetVportsRefs([]IxiaCfgNode{vport})
	vport.SetConnectedToRef(topo) // This is not a valid config, just an example for testing singular refs.
	cfg.updateAllXPathsAndRefs()
	if len(topo.Vports) != 1 || topo.Vports[0] != vport.XPath().String() {
		t.Errorf("Vport reference in Topology config not correctly updated: got %v, want %v", topo.Vports, []string{vport.XPath().String()})
	}
	if *(vport.ConnectedTo) != topo.XPath().String() {
		t.Errorf("Topology reference in Vport config not correctly updated: got %q, want %q", *(vport.ConnectedTo), topo.XPath().String())
	}

	// Update refs again to ensure resolved refs are not duplicated.
	cfg.updateAllXPathsAndRefs()
	if len(topo.Vports) != 1 || topo.Vports[0] != vport.XPath().String() {
		t.Errorf("Vport reference in Topology config not correctly updated after second ref update: got %v, want %v", topo.Vports, []string{vport.XPath().String()})
	}
}
