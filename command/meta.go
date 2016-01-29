package command

import (
	"bufio"
	"flag"
	"io"
	"strings"

	"github.com/mitchellh/cli"
)

// FlagSetFlags is an enum to define what flags are present in the
// default FlagSet returned by Meta.FlagSet.
type FlagSetFlags uint

// Meta contains the meta-options and functionality that nearly every
// Vault command inherits.
type Meta struct {
	Ui cli.Ui

	// These are internal and shouldn't be modified or access by anyone
	// except Meta.
	config *Config
}

// Config loads the configuration and returns it. If the configuration
// is already loaded, it is returned.
func (m *Meta) Config() (*Config, error) {
	if m.config != nil {
		return m.config, nil
	}

	var err error
	m.config, err = LoadConfig("")
	if err != nil {
		return nil, err
	}

	return m.config, nil
}

// FlagSet returns a FlagSet with the common flags that every
// command implements. The exact behavior of FlagSet can be configured
// using the flags as the second parameter, for example to disable
// server settings on the commands that don't talk to a server.
func (m *Meta) FlagSet(n string, fs FlagSetFlags) *flag.FlagSet {
	f := flag.NewFlagSet(n, flag.ContinueOnError)

	// Create an io.Writer that writes to our Ui properly for errors.
	// This is kind of a hack, but it does the job. Basically: create
	// a pipe, use a scanner to break it into lines, and output each line
	// to the UI. Do this forever.
	errR, errW := io.Pipe()
	errScanner := bufio.NewScanner(errR)
	go func() {
		for errScanner.Scan() {
			m.Ui.Error(errScanner.Text())
		}
	}()
	f.SetOutput(errW)

	return f
}

// generalOptionsUsage returns the usage documenation for commonly
// available options
func generalOptionsUsage() string {
	general := ``
	return strings.TrimSpace(general)
}
