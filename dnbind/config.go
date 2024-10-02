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

package dnbind

import (
	"fmt"
	"os"

	"github.com/openconfig/ondatra/dnbind/creds"
	"gopkg.in/yaml.v2"
)

type Node struct {
	Id              string         `yaml:"id"`
	Hostname        string         `yaml:"hostname"`
	Vendor          string         `yaml:"vendor,omitempty"`
	HardwareModel   string         `yaml:"model,omitempty"`
	SoftwareVersion string         `yaml:"version,omitempty"`
	Credentials     creds.UserPass `yaml:"credentials,omitempty"`
}

// Config contains parameters to configure the Drivenets binding.
type Config struct {
	Credentials *creds.Credentials
	Nodes       []Node `yaml:"nodes"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%+v", *c)
}

func (c *Config) Lookup(id string) *Node {
	for _, node := range c.Nodes {
		if node.Id == id {
			return &node
		}
	}
	return nil
}

// ParseCredFile parses a yaml file containing a serialized Config.
func ParseConfigFile(credFile string) (*Config, error) {
	data, err := os.ReadFile(credFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	c := &Config{
		Credentials: &creds.Credentials{
			Node: make(map[string]*creds.UserPass),
		},
	}
	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, fmt.Errorf("error unmarshalling config YAML: %w", err)
	}
	for _, node := range c.Nodes {
		c.Credentials.Node[node.Hostname] = &node.Credentials
	}

	return c, nil
}
