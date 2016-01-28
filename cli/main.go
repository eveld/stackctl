package cli

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// Run calls the custom Run function.
func Run(args []string) int {
	return RunCustom(args, Commands(nil))
}

// RunCustom starts the CLI and populates the help.
func RunCustom(args []string, commands map[string]cli.CommandFactory) int {
	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	// Build the commands to include in the help now. This is pretty...
	// tedious, but we don't have a better way at the moment.
	commandsInclude := make([]string, 0, len(commands))
	for k := range commands {
		switch k {
		case "hidden":
		default:
			commandsInclude = append(commandsInclude, k)
		}
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: commands,
		HelpFunc: cli.FilteredHelpFunc(commandsInclude, HelpFunc),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
