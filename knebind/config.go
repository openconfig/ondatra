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

package knebind

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	tpb "github.com/openconfig/kne/proto/topo"
)

// Config contains parameters to configure the KNE binding.
// They are all exported so they can be unmarhalled from YAML.
type Config struct {
	// TODO(team): Deprecate username and password fields. Add option inside credentials field.
	Username, Password string
	Credentials        *Credentials `yaml:"credentials"`
	TopoPath           string       `yaml:"topology"`
	KubecfgPath        string       `yaml:"kubecfg"`
	SkipReset          bool         `yaml:"skip_reset"`
}

// Credentials contains credential maps for nodes in the KNE topology.
type Credentials struct {
	Node   map[string]*UserPass     `yaml:"node"`
	Vendor map[tpb.Vendor]*UserPass `yaml:"vendor"`
}

// UnmarshalYAML allows the Credentials type to be correctly unmarshaled from yaml.
func (c *Credentials) UnmarshalYAML(unmarshal func(any) error) error {
	var u map[string]map[any]*UserPass
	if err := unmarshal(&u); err != nil {
		return err
	}
	for k, v := range u {
		switch key := strings.ToLower(k); key {
		case "node":
			c.Node = make(map[string]*UserPass)
			for k, v := range v {
				c.Node[k.(string)] = v
			}
		case "vendor":
			c.Vendor = make(map[tpb.Vendor]*UserPass)
			for k, v := range v {
				n, ok := tpb.Vendor_value[k.(string)]
				if !ok {
					return fmt.Errorf("kne vendor %v not recognized", k)
				}
				c.Vendor[tpb.Vendor(n)] = v
			}
		}
	}
	return nil
}

// UserPass contains a username and password combination.
type UserPass struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%+v", *c)
}

// ParseConfigFile parses a yaml file containing a serialized Config.
func ParseConfigFile(configFile string) (*Config, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	c := &Config{}
	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, fmt.Errorf("error unmarshalling config YAML: %w", err)
	}
	if c.TopoPath == "" {
		return nil, fmt.Errorf("no topology path specified in config: %v", c)
	}
	return c, nil
}
