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

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/knebind/creds"
	"gopkg.in/yaml.v2"
)

// Config contains parameters to configure the KNE binding.
type Config struct {
	// TODO(team): Deprecate username and password fields.
	Username, Password string
	Credentials        *creds.Credentials `yaml:"credentials"`
	Topology           string             `yaml:"topology"`
	Kubeconfig         string             `yaml:"kubecfg"`
	SkipReset          bool               `yaml:"skip_reset"`
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
	if c.Topology == "" {
		return nil, fmt.Errorf("no topology path specified in config: %v", c)
	}
	if c.Username != "" || c.Password != "" {
		log.Error("The top-level 'username' and 'password' keys are deprecated; " +
			"Use the 'credentials' key instead.\n" +
			"See https://github.com/openconfig/ondatra/tree/main/knebind#device-credentials")
	}
	return c, nil
}
