package command

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVersionCommand(t *testing.T) {
	cmd := NewVersionCommand()

	t.Run("has correct use", func(t *testing.T) {
		assert.Equal(t, "version", cmd.Use)
	})

	t.Run("has short description", func(t *testing.T) {
		assert.NotEmpty(t, cmd.Short)
		assert.Contains(t, cmd.Short, "version")
	})
}

func TestVersionCommand_Execute(t *testing.T) {
	t.Run("shows version information", func(t *testing.T) {
		cmd := NewVersionCommand()
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)

		err := cmd.Execute()

		require.NoError(t, err)
		output := buf.String()
		assert.Contains(t, output, "cloud version")
	})

	t.Run("shows build time when available", func(t *testing.T) {
		// Set version info
		origVersion := Version
		origBuildTime := BuildTime
		Version = "1.0.0"
		BuildTime = "2024-01-01_12:00:00"
		defer func() {
			Version = origVersion
			BuildTime = origBuildTime
		}()

		cmd := NewVersionCommand()
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)

		err := cmd.Execute()

		require.NoError(t, err)
		output := buf.String()
		assert.Contains(t, output, "1.0.0")
		assert.Contains(t, output, "2024-01-01")
	})
}

func TestVersionCommand_SetVersion(t *testing.T) {
	t.Run("can set version at compile time", func(t *testing.T) {
		origVersion := Version
		defer func() { Version = origVersion }()

		Version = "2.0.0-beta"

		cmd := NewVersionCommand()
		buf := new(bytes.Buffer)
		cmd.SetOut(buf)

		err := cmd.Execute()

		require.NoError(t, err)
		assert.Contains(t, buf.String(), "2.0.0-beta")
	})
}
