package main

import (
	"os"

	"github.com/danielmiessler/fabric/cli"
)

// main is the entry point for the Fabric CLI application.
// Fabric is an open-source AI augmentation framework designed to help
// humans apply AI to everyday tasks using a crowdsourced set of patterns.
func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
