package command

import (
	"github.com/eveld/stackctl/version"
	"github.com/mitchellh/cli"
)

// VersionCommand is a Command implementation prints the version.
type VersionCommand struct {
	Info *version.Info
	Ui   cli.Ui
}

// Help gives the help of the version command.
func (c *VersionCommand) Help() string {
	return ""
}

// Run executes the version command.
func (c *VersionCommand) Run(_ []string) int {
	c.Ui.Output(c.Info.String())
	return 0
}

// Synopsis gives a short description of the version command.
func (c *VersionCommand) Synopsis() string {
	return "Prints the Stackctl version"
}
