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
//   - exit code 1 on error; stderr used for error output (not stdout)
//   - TODO: explore adding a --dry-run flag to preview pattern output
//   - TODO: look into piping output directly to pbcopy on macOS for clipboard support
//   - TODO: investigate --watch mode to re-run pattern on file change (inotify/fsevents)
func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
