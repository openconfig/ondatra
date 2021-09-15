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

	"github.com/pkg/errors"
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
		return nil, errors.Wrapf(err, "Error reading config file")
	}
	c := &Config{}
	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, errors.Wrapf(err, "Error unmarshalling config YAML")
	}
	if c.Username == "" {
		return nil, errors.Errorf("No username specified in config: %v", c)
	}
	if c.Password == "" {
		return nil, errors.Errorf("No password specified in config: %v", c)
	}
	if c.TopoPath == "" {
		return nil, errors.Errorf("No topology path specified in config: %v", c)
	}
	if c.CLIPath == "" {
		// If no CLI path specified, use kne_cli available in PATH.
		c.CLIPath = "kne_cli"
	}
	return c, nil
}
