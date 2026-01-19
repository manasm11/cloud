// Package command contains the CLI commands for the cloud application.
package command

import (
	"github.com/spf13/cobra"
)

// NewRootCommand creates and returns the root command.
func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cloud",
		Short: "A CLI tool for creating comprehensive project specifications",
		Long: `cloud is an interactive CLI tool that helps developers create complex,
production-ready applications by systematically gathering requirements through
guided questioning, suggesting appropriate technology stacks with educational
explanations, and generating comprehensive project specifications.

The tool embodies the philosophy: "Ask everything, assume nothing, build to last."

Use 'cloud new' to start creating a new project specification.`,
		// When run without arguments, show help
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// Add subcommands
	rootCmd.AddCommand(NewVersionCommand())
	rootCmd.AddCommand(NewNewCommand())

	return rootCmd
}
