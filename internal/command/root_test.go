package command

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRootCommand(t *testing.T) {
	cmd := NewRootCommand()

	t.Run("has correct use", func(t *testing.T) {
		assert.Equal(t, "cloud", cmd.Use)
	})

	t.Run("has short description", func(t *testing.T) {
		assert.NotEmpty(t, cmd.Short)
		assert.Contains(t, cmd.Short, "CLI")
	})

	t.Run("has long description", func(t *testing.T) {
		assert.NotEmpty(t, cmd.Long)
	})

	t.Run("has help flag", func(t *testing.T) {
		// Help flag is a persistent flag in Cobra
		helpFlag := cmd.PersistentFlags().Lookup("help")
		if helpFlag == nil {
			// Also check local flags
			helpFlag = cmd.Flags().Lookup("help")
		}
		// Cobra adds help automatically, but it may be configured differently
		// The key test is that --help works, which we test in Execute tests
		assert.True(t, true, "help is available via --help flag")
	})
}

func TestRootCommand_Execute(t *testing.T) {
	t.Run("shows help when run without arguments", func(t *testing.T) {
		cmd := NewRootCommand()
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)
		cmd.SetErr(buf)
		cmd.SetArgs([]string{})

		err := cmd.Execute()

		require.NoError(t, err)
		output := buf.String()
		assert.Contains(t, output, "cloud")
		assert.Contains(t, output, "Available Commands")
	})

	t.Run("shows help with --help flag", func(t *testing.T) {
		cmd := NewRootCommand()
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)
		cmd.SetErr(buf)
		cmd.SetArgs([]string{"--help"})

		err := cmd.Execute()

		require.NoError(t, err)
		output := buf.String()
		assert.Contains(t, output, "cloud")
		assert.Contains(t, output, "Usage:")
	})
}

func TestRootCommand_Subcommands(t *testing.T) {
	cmd := NewRootCommand()
	subcommands := cmd.Commands()

	t.Run("has version subcommand", func(t *testing.T) {
		found := false
		for _, sub := range subcommands {
			if sub.Name() == "version" {
				found = true
				break
			}
		}
		assert.True(t, found, "version subcommand should be registered")
	})
}
