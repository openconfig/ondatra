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

// Package init installs the Ondatra binding for testing with Drivenets devices.
package init

import (
	"flag"

	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/dnbind"
	"github.com/openconfig/ondatra/dnbind/creds"
)

var (
	configFile = flag.String("config", "", "YAML config file")
	credFlags  = creds.DefineFlags()
)

// Init provides a generator for a Drivenets bind instance.
// To be used with ondatra.RunTests.
func Init() (binding.Binding, error) {
	var cfg *dnbind.Config
	var err error

	if *configFile != "" {
		cfg, err = dnbind.ParseConfigFile(*configFile)
	} else {
		cfg, err = parseFlags()
	}
	if err != nil {
		return nil, err
	}
	return dnbind.New(cfg)
}

func parseFlags() (*dnbind.Config, error) {
	cred, err := credFlags.Parse()
	if err != nil {
		return nil, err
	}
	return &dnbind.Config{
		Credentials: cred,
	}, nil
}
