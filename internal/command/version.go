package command

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Version information, set at build time via ldflags.
var (
	Version   = "dev"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

// NewVersionCommand creates and returns the version command.
func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version information",
		Long:  `Print the version number, build time, and other build information for cloud.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "cloud version %s\n", Version)
			fmt.Fprintf(cmd.OutOrStdout(), "  Build time: %s\n", BuildTime)
			fmt.Fprintf(cmd.OutOrStdout(), "  Git commit: %s\n", GitCommit)
			fmt.Fprintf(cmd.OutOrStdout(), "  Go version: %s\n", runtime.Version())
			fmt.Fprintf(cmd.OutOrStdout(), "  OS/Arch:    %s/%s\n", runtime.GOOS, runtime.GOARCH)
		},
	}
}
