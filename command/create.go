package command

import "fmt"

// CreateCommand is a Command that creates a new stack.
type CreateCommand struct {
	Meta
}

// Help gives the help of the create command.
func (c *CreateCommand) Help() string {
	return "Explain how to create a stack"
}

// Run executes the create command.
func (c *CreateCommand) Run(args []string) int {
	c.Ui.Output(fmt.Sprintf("Creating stack"))
	return 0
}

// Synopsis gives a short description of the create command.
func (c *CreateCommand) Synopsis() string {
	return "Creates a new stack"
}
