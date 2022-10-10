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
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config contains parameters to configure the KNE binding.
// They are all exported so they can be unmarhalled from YAML.
type Config struct {
	Username, Password string
	TopoPath           string `yaml:"topology"`
	CLIPath            string `yaml:"cli"`
	KubecfgPath        string `yaml:"kubecfg"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%+v", *c)
}

// ParseConfigFile parses a yaml file containing a serialized Config.
func ParseConfigFile(configFile string) (*Config, error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	c := &Config{}
	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, fmt.Errorf("error unmarshalling config YAML: %w", err)
	}
	if c.Username == "" {
		return nil, fmt.Errorf("no username specified in config: %v", c)
	}
	if c.Password == "" {
		return nil, fmt.Errorf("No password specified in config: %v", c)
	}
	if c.TopoPath == "" {
		return nil, fmt.Errorf("no topology path specified in config: %v", c)
	}
	if c.CLIPath == "" {
		// If no CLI path specified, use kne available in PATH.
		c.CLIPath = "kne"
	}
	return c, nil
}
