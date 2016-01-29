package command

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
	"github.com/mitchellh/go-homedir"
)

const (
	// DefaultConfigPath is the default path to the configuration file
	DefaultConfigPath = "~/.stackctl"

	// ConfigPathEnv is the environment variable that can be used to
	// override where the Stackctl configuration is.
	ConfigPathEnv = "STACKCTL_CONFIG_PATH"
)

// Config is the CLI configuration for Vault that can be specified via
// a `$HOME/.stackctl` file which is HCL-formatted (therefore HCL or JSON).
type Config struct {
}

// LoadConfig reads the configuration from the given path. If path is
// empty, then the default path will be used, or the environment variable
// if set.
func LoadConfig(path string) (*Config, error) {
	if path == "" {
		path = DefaultConfigPath
	}
	if v := os.Getenv(ConfigPathEnv); v != "" {
		path = v
	}

	path, err := homedir.Expand(path)
	if err != nil {
		return nil, fmt.Errorf("Error expanding config path: %s", err)
	}

	var config Config
	contents, err := ioutil.ReadFile(path)
	if !os.IsNotExist(err) {
		if err != nil {
			return nil, err
		}

		obj, err := hcl.Parse(string(contents))
		if err != nil {
			return nil, err
		}

		if err := hcl.DecodeObject(&config, obj); err != nil {
			return nil, err
		}
	}

	return &config, nil
}
