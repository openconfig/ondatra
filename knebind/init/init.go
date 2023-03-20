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
	"errors"
	"flag"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/knebind"
	"github.com/openconfig/ondatra/knebind/creds"
)

var (
	configFile = flag.String("config", "", "YAML configuration file")
	topology   = flag.String("topology", "", "Path to the topology proto file")
	kubeConfig = flag.String("kubeconfig", "",
		"Optional path to a kubeconfig file; Defaults to ~/.kube/config")
	skipReset = flag.Bool("skip_reset", false,
		"If true, skip initial device reset that happens during reservation")
	credFlags = creds.DefineFlags()
)

// Init provides a generator for a KNE bind instance which uses the
// configuration set by the flag --config. To be used with ondatra.RunTests.
func Init() (binding.Binding, error) {
	if *topology == "" && *configFile == "" {
		return nil, errors.New("either --topology or --config flag must be provided")
	}
	var cfg *knebind.Config
	var err error
	if *configFile != "" {
		log.Errorf("The --config flag is deprecated. Please use separate flags instead.")
		cfg, err = knebind.ParseConfigFile(*configFile)
	} else {
		cfg, err = parseFlags()
	}
	if err != nil {
		return nil, err
	}
	return knebind.New(cfg)
}

func parseFlags() (*knebind.Config, error) {
	cred, err := credFlags.Parse()
	if err != nil {
		return nil, err
	}
	return &knebind.Config{
		Topology:    *topology,
		Credentials: cred,
		Kubeconfig:  *kubeConfig,
		SkipReset:   *skipReset,
	}, nil
}
