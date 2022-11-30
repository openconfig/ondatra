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
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/net/context"

	"github.com/google/go-cmp/cmp"
)

type fakeSession struct {
	ixSession
	config config
}

func (s *fakeSession) Config() config {
	return s.config
}

type fakeConfig struct {
	config
	exportRes any
	importErr error
	queryRes  any
}

func (c *fakeConfig) Export(context.Context) (string, error) {
	switch v := c.exportRes.(type) {
	case string:
		return v, nil
	case error:
		return "", v
	default:
		return "", fmt.Errorf("unexpected result type: %v (%T)", v, v)
	}
}

func (c *fakeConfig) Import(context.Context, string, bool) error {
	return c.importErr
}

func (c *fakeConfig) QueryIDs(ctx context.Context, xpaths ...string) (map[string]string, error) {
	switch v := c.queryRes.(type) {
	case map[string]string:
		return v, nil
	case nil:
		return nil, nil
	case error:
		return nil, v
	default:
		return nil, fmt.Errorf("unexpected result type: %v (%T)", v, v)
	}
}

func TestExportConfig(t *testing.T) {
	toXPath := func(p string) *XPath {
		var xp XPath
		json.Unmarshal([]byte(strconv.Quote(p)), &xp)
		return &xp
	}
	cmpXPath := func(xp1, xp2 *XPath) bool {
		// Handle nils by using a format string.
		return fmt.Sprintf("%s", xp1) == fmt.Sprintf("%s", xp2)
	}

	tests := []struct {
		desc    string
		res     any
		wantErr string
		wantCfg *Ixnetwork
	}{{
		desc:    "error",
		res:     errors.New("POST error"),
		wantErr: "POST error",
	}, {
		desc: "success",
		res:  `{"xpath": "/", "availableHardware": {"xpath": "/availableHardware"}}`,
		wantCfg: &Ixnetwork{
			Xpath: toXPath("/"),
			AvailableHardware: &AvailableHardware{
				Xpath: toXPath("/availableHardware"),
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &Client{sess: &fakeSession{
				config: &fakeConfig{exportRes: test.res},
			}}
			gotCfg, gotErr := c.ExportConfig(context.Background())
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("ExportConfig: unexpected error, got err %v, want err %q", gotErr, test.wantErr)
			}
			if diff := cmp.Diff(test.wantCfg, gotCfg, cmp.Comparer(cmpXPath)); diff != "" {
				t.Errorf("ExportConfig: unexpected returned config, diff: (-want +got)\n%s.", diff)
			}
		})
	}
}

func TestImportConfig(t *testing.T) {
	cfg := &Ixnetwork{AvailableHardware: &AvailableHardware{}}
	tests := []struct {
		desc      string
		node      IxiaCfgNode
		overwrite bool
		err       error
		wantErr   string
	}{{
		desc:    "error",
		node:    cfg,
		err:     errors.New("POST error"),
		wantErr: "POST error",
	}, {
		desc:      "success full push.",
		node:      cfg,
		overwrite: true,
	}, {
		desc: "success partial push",
		node: cfg.AvailableHardware,
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &Client{sess: &fakeSession{
				config: &fakeConfig{importErr: test.err},
			}}
			gotErr := c.ImportConfig(context.Background(), cfg, test.node, test.overwrite)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("ImportConfig: unexpected error, got err %v, want err %q", gotErr, test.wantErr)
			}
		})
	}
}

func TestUpdateIDs(t *testing.T) {
	const (
		topologyID    = "/fake/abs/path/to/topology/2"
		deviceGroupID = "/fake/abs/path/to/topology/2/deviceGroup/3"
	)
	deviceGroup := &TopologyDeviceGroup{}
	topology := &Topology{
		DeviceGroup: []*TopologyDeviceGroup{deviceGroup},
	}
	cfg := &Ixnetwork{
		Topology: []*Topology{topology},
	}

	cfg.updateAllXPathsAndRefs() // Needed for mapping XPaths to expected IDs.
	topologyXPath := topology.XPath().String()
	deviceGroupXPath := deviceGroup.XPath().String()

	tests := []struct {
		desc    string
		nodes   []IxiaCfgNode
		res     any
		ids     map[string]string
		wantIDs map[string]string
		wantErr string
	}{{
		desc: "empty list noop",
	}, {
		desc:    "error",
		nodes:   []IxiaCfgNode{topology},
		res:     errors.New("query error"),
		wantErr: "query error",
	}, {
		desc:    "update object from cache",
		nodes:   []IxiaCfgNode{deviceGroup},
		ids:     map[string]string{deviceGroupXPath: deviceGroupID},
		wantIDs: map[string]string{deviceGroupXPath: deviceGroupID},
	}, {
		desc:  "query multiple objects and update non-empty cache",
		nodes: []IxiaCfgNode{topology, deviceGroup},
		ids:   map[string]string{"/fake/xpath": "/fakeID"},
		res: map[string]string{
			topologyXPath:    topologyID,
			deviceGroupXPath: deviceGroupID,
		},
		wantIDs: map[string]string{
			"/fake/xpath":    "/fakeID",
			topologyXPath:    topologyID,
			deviceGroupXPath: deviceGroupID,
		},
	}, {
		desc:    "query single object",
		nodes:   []IxiaCfgNode{deviceGroup},
		ids:     map[string]string{},
		res:     map[string]string{deviceGroupXPath: deviceGroupID},
		wantIDs: map[string]string{deviceGroupXPath: deviceGroupID},
	}}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := &Client{
				sess: &fakeSession{
					config: &fakeConfig{queryRes: test.res},
				},
				xPathToID: test.ids,
			}
			gotErr := c.UpdateIDs(context.Background(), cfg, test.nodes...)
			if (gotErr == nil) != (test.wantErr == "") || (gotErr != nil && !strings.Contains(gotErr.Error(), test.wantErr)) {
				t.Errorf("UpdateIDs: unexpected error, got err %v, want err %q", gotErr, test.wantErr)
			}
			if diff := cmp.Diff(test.wantIDs, c.xPathToID); diff != "" {
				t.Errorf("UpdateIDs: unexpected diff in ID cache (-want,+got): %s\n", diff)
			}
		})
	}
}

// This case tests that ID updates succeed even with no XPaths set,
// as is the case for Ixnetwork objects returned from LastImportedConfig.
func TestUpdateIDsOfLastImported(t *testing.T) {
	const (
		fakeXPath     = "/fake/xpath"
		fakeID        = "/fakeID"
		topoID        = "/fake/abs/path/to/topology/2"
		deviceGroupID = "/fake/abs/path/to/topology/2/deviceGroup/3"
	)

	cfg := &Ixnetwork{
		Topology: []*Topology{{
			DeviceGroup: []*TopologyDeviceGroup{{}},
		}},
	}
	cfg.updateAllXPathsAndRefs()
	topoXPath := cfg.Topology[0].XPath().String()
	deviceGroupXPath := cfg.Topology[0].DeviceGroup[0].XPath().String()

	c := &Client{
		sess: &fakeSession{config: &fakeConfig{
			queryRes: map[string]string{
				topoXPath:        topoID,
				deviceGroupXPath: deviceGroupID,
			},
		}},
		xPathToID:    map[string]string{fakeXPath: fakeID},
		lastImported: cfg.Copy(),
	}

	lastCfg := c.LastImportedConfig()
	lastCfgTopo := lastCfg.Topology[0]
	lastCfgDeviceGroup := lastCfgTopo.DeviceGroup[0]

	if err := c.UpdateIDs(context.Background(), lastCfg, lastCfgTopo, lastCfgDeviceGroup); err != nil {
		t.Fatalf("UpdateIDs: unexpected error updating config objects with unset xpaths: %v", err)
	}

	// Instantiate after UpdateIDs when xpaths should be set
	wantIDs := map[string]string{
		fakeXPath:                           fakeID,
		lastCfgTopo.XPath().String():        topoID,
		lastCfgDeviceGroup.XPath().String(): deviceGroupID,
	}
	if diff := cmp.Diff(wantIDs, c.xPathToID); diff != "" {
		t.Errorf("UpdateIDs: unexpected diff in ID cache (-want,+got): %s\n", diff)
	}
}
