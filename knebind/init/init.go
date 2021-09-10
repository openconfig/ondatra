// Package init installs the Ondatra binding for testing with kne clusters.
// It also installs a --config flag, which must specify a kne config file. See
// config.go for syntax details.
package init

import (
	"flag"

	"github.com/pkg/errors"
	"github.com/openconfig/ondatra/internal/binding"
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
