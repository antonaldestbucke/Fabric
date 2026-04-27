package main

import (
	"fmt"
	"os"

	"github.com/danielmiessler/fabric/cli"
)

// main is the entry point for the Fabric CLI application.
// Fabric is an open-source AI augmentation framework designed to help
// humans apply AI to everyday tasks using a crowdsourced set of patterns.
func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
