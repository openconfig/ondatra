// Copyright 2021 Google LLC
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

package ixweb

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/context"
)

const resourceManagerPath = "resourceManager"

// xpathFieldsRE matches words immediately after a slash (including the slash).
var xpathFieldsRE = regexp.MustCompile(`/(\w+)`)

// Config represents the IxNetwork config API.
type Config struct {
	sess *Session
}

// Export exports the current full configuration of the IxNetwork session.
func (c *Config) Export(ctx context.Context) (string, error) {
	const exportConfigPath = resourceManagerPath + "/operations/exportconfig"
	exportReqArgs := OpArgs{
		c.sess.AbsPath(resourceManagerPath),
		[]string{"/descendant-or-self::*"},
		false,
		"json",
	}
	var cfg string
	if err := c.sess.Post(ctx, exportConfigPath, exportReqArgs, &cfg); err != nil {
		return "", fmt.Errorf("failed to export IxNetwork config from session: %w", err)
	}
	return cfg, nil
}

// Import imports the specified config into the IxNetwork session.
// If overwrite is 'true', the existing config is completely replaced with the
// specified config; otherwise the config is updated at or below the node
// represented by the specified config. For values that are a list of nodes,
// only the nodes that are specified are updated. (E.g. you cannot remove a
// config node from a list with overwrite set to 'false'.)
func (c *Config) Import(ctx context.Context, cfg string, overwrite bool) error {
	const importConfigPath = resourceManagerPath + "/operations/importconfig"
	importReqData := OpArgs{
		c.sess.AbsPath(resourceManagerPath),
		cfg,
		overwrite,
		"suppressNothing",
		false,
	}
	if err := c.sess.Post(ctx, importConfigPath, importReqData, nil); err != nil {
		return fmt.Errorf("failed import IxNetwork config to session: %w", err)
	}
	return nil
}

// QueryIDs resolves the specified XPaths to config node IDs, returning a map
// from xpath to ID. For better performance, callers prefer to group multiple
// xpaths into one call rather than multiple calls with fewer XPaths per call.
func (c *Config) QueryIDs(ctx context.Context, xpaths ...string) (map[string]string, error) {
	if len(xpaths) == 0 {
		return nil, nil
	}

	// Fields names to query for in the select request
	fieldSet := make(map[string]bool)
	for _, xp := range xpaths {
		for _, m := range xpathFieldsRE.FindAllStringSubmatch(xp, -1) {
			fieldSet[m[1]] = true
		}
	}
	fieldList := make([]string, 0, len(fieldSet))
	for fn := range fieldSet {
		fieldList = append(fieldList, fn)
	}

	const selectOpPath = "operations/select?xpath=true"
	selectReq := map[string][]any{
		"selects": []any{
			map[string]any{
				// Base path in IxNetwork REST hierarchy to query from.
				// Since we can't assume we know any IDs we can't make this narrower.
				"from":       "/",
				"properties": []string{""},
				"children": []any{
					map[string]any{
						// Regex of field names of objects to return IDs for. This operation
						// will not recurse on a field name in the hierarchy that is not
						// included here, so the field names of all parent objects of a
						// target object must be included.
						"child":      fmt.Sprintf("^(%s)$", strings.Join(fieldList, "|")),
						"properties": []string{"id"},
						"filters": []map[string]string{
							map[string]string{
								"property": "",
								"regex":    "",
							},
						},
					},
				},
				"inlines": []map[string]any{
					map[string]any{
						"child":      "",
						"properties": []string{""},
					},
				},
			},
		},
	}
	var resp []map[string]any
	if err := c.sess.Post(ctx, selectOpPath, selectReq, &resp); err != nil {
		return nil, fmt.Errorf("error querying for node IDs: %w", err)
	}
	if len(resp) != 1 {
		return nil, fmt.Errorf("expected one element in select response, got %d (%v)", len(resp), resp)
	}

	xpathToID := make(map[string]string)
	for _, xp := range xpaths {
		xpathToID[xp] = ""
	}
	updateIDsFromMap(resp[0], xpathToID)
	var notFound []string
	for xp, id := range xpathToID {
		if id == "" {
			notFound = append(notFound, xp)
		}
	}
	if len(notFound) > 0 {
		return nil, fmt.Errorf("did not find node IDs for the following XPaths: %v", notFound)
	}
	return xpathToID, nil
}

// updateIDsFromMap recurses through a nested map response from the IxNetwork
// 'select' operation, updating the IDs for xpaths in the given xpathToID map.
func updateIDsFromMap(m map[string]any, xpathToID map[string]string) {
	xp, hasXPath := m["xpath"].(string)
	id, hasID := m["href"].(string)
	if hasXPath && hasID {
		if _, wantXPath := xpathToID[xp]; wantXPath {
			xpathToID[xp] = id
		}
	}

	// Update on an sublists/maps.
	var updateIDs func(any)
	updateIDs = func(mapOrList any) {
		switch val := mapOrList.(type) {
		case []any:
			for _, v := range val {
				updateIDs(v)
			}
		case map[string]any:
			updateIDsFromMap(val, xpathToID)
		}
	}
	for _, v := range m {
		updateIDs(v)
	}
}
