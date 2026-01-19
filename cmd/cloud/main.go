// Package main is the entry point for the cloud CLI application.
package main

import (
	"os"

	"github.com/manas/cloud/internal/command"
)

// Build-time variables set via ldflags
var (
	version   = "dev"
	buildTime = "unknown"
	gitCommit = "unknown"
)

func main() {
	// Set version info in command package
	command.Version = version
	command.BuildTime = buildTime
	command.GitCommit = gitCommit

	// Create and execute root command
	rootCmd := command.NewRootCommand()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
