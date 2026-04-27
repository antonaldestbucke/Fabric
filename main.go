package main

import (
	"fmt"
	"os"

	"github.com/danielmiessler/fabric/cli"
)

// main is the entry point for the Fabric CLI application.
// Fabric is an open-source AI augmentation framework designed to help
// humans apply AI to everyday tasks using a crowdsourced set of patterns.
//
// Personal fork: using this for local AI workflow automation and learning
// how Go CLI apps are structured. -- @myfork
//
// Notes to self:
//   - patterns live in ~/.config/fabric/patterns/
//   - run `fabric --list` to see available patterns
//   - `fabric --update` pulls latest patterns from upstream
func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
