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
	"fmt"
)

// updateAllXPathsAndRefs is a convenience function for updating XPaths and
// refs from the root of the config.
func (cfg *Ixnetwork) updateAllXPathsAndRefs() {
	cfg.updateXPaths(&XPath{parentXPath: "/"})
	cfg.updateRefs()
}

// Copy returns a new deep copy of the IxNetwork config.
func (cfg *Ixnetwork) Copy() *Ixnetwork {
	return (cfg.copyCfg(map[any]any{})).(*Ixnetwork)
}

func removeNilsMap(m map[string]any) {
	var removeNils func(any)
	removeNils = func(v any) {
		if vl, ok := v.([]any); ok {
			for _, v := range vl {
				removeNils(v)
			}
		} else if vm, ok := v.(map[string]any); ok {
			removeNilsMap(vm)
		}
	}
	for k, v := range m {
		if v == nil {
			delete(m, k)
			continue
		}
		removeNils(v)
	}
}

// MarshalJSON implements the encoding/json.Marshaler interface.
// We use this custom implementation rather than default handling for Ixnetwork
// objects because we need to remove keys with 'null' values from the final JSON
// output.
func (cfg *Ixnetwork) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(*cfg)
	if err != nil {
		return nil, err
	}

	var jsonMap map[string]any
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return nil, fmt.Errorf("could not unmarshal the already marshaled Ixnetwork config: %w", err)
	}

	// Remove any explicit 'null' values because the IxNetwork config import API will not accept configs containing them.
	removeNilsMap(jsonMap)
	b, err = json.Marshal(jsonMap)
	if err != nil {
		return nil, fmt.Errorf("could not marshal map with nulls removed back into JSON: %w", err)
	}
	return b, nil
}

// ResolvedConfig returns a copy of the config node with XPaths updated and
// object references resolved. It does not modify the original config.
func (cfg *Ixnetwork) ResolvedConfig(node IxiaCfgNode) (IxiaCfgNode, error) {
	origToCopy := map[any]any{}
	cpy := (cfg.copyCfg(origToCopy)).(*Ixnetwork)
	cpy.updateAllXPathsAndRefs()
	nodeCpy, ok := (origToCopy[node]).(IxiaCfgNode)
	if !ok {
		return nil, fmt.Errorf("could not find provided config at %q node from root config", node.XPath().String())
	}
	return nodeCpy, nil
}

// MarshalJSON implements the encoding/json.Marshaler interface.
// We use this custom implementation rather than default handling for Ixnetwork
// objects because we need to remove unconfigured traffic fields from the final
// JSON output.
func (stack *TrafficTrafficItemConfigElementStack) MarshalJSON() ([]byte, error) {
	var cfgFields []*TrafficTrafficItemConfigElementStackField
	for _, f := range stack.Field {
		if f.Auto != nil && *(f.Auto) == false {
			cfgFields = append(cfgFields, f)
		}
	}
	return json.Marshal(map[string]any{
		"xpath": stack.Xpath,
		"field": cfgFields,
	})
}
