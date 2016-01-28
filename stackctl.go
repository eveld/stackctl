package main // import "github.com/eveld/stackctl"

import (
	"os"

	"github.com/eveld/stackctl/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
