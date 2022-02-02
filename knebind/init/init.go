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

// Package init installs the Ondatra binding for testing with kne clusters.
// It also installs a --config flag, which must specify a kne config file. See
// config.go for syntax details.
package init

import (
	"flag"

	"github.com/pkg/errors"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/knebind"
)

var (
	configFile = flag.String("config", "", "YAML configuration file")
)

// Init provides a generator for a KNE bind instance which uses the
// configuration set by the flag --config. To be used with ondatra.RunTests.
func Init() (binding.Binding, error) {
	if *configFile == "" {
		return nil, errors.New("No --config flag specified")
	}
	cfg, err := knebind.ParseConfigFile(*configFile)
	if err != nil {
		return nil, err
	}
	return knebind.New(cfg)
}
