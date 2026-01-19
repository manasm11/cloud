package questions

import (
	"testing"

	"github.com/manas/cloud/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppTypeQuestions_Ask(t *testing.T) {
	t.Run("selects web application", func(t *testing.T) {
		mock := NewMockPrompter()
		mock.SelectResponses["What type of application are you building?"] = "web"
		mock.SelectResponses["Web Application Type"] = "spa"

		questions := NewAppTypeQuestions(mock)
		result, err := questions.Ask()

		require.NoError(t, err)
		assert.Equal(t, domain.AppTypeWebApp, result.AppType)
		assert.Equal(t, "spa", result.Details["web_type"])
	})

	t.Run("selects CLI tool", func(t *testing.T) {
		mock := NewMockPrompter()
		mock.SelectResponses["What type of application are you building?"] = "cli"
		mock.SelectResponses["CLI Type"] = "interactive"

		questions := NewAppTypeQuestions(mock)
		result, err := questions.Ask()

		require.NoError(t, err)
		assert.Equal(t, domain.AppTypeCLI, result.AppType)
		assert.Equal(t, "interactive", result.Details["cli_type"])
	})

	t.Run("selects API backend", func(t *testing.T) {
		mock := NewMockPrompter()
		mock.SelectResponses["What type of application are you building?"] = "api"
		mock.SelectResponses["API Type"] = "public"

		questions := NewAppTypeQuestions(mock)
		result, err := questions.Ask()

		require.NoError(t, err)
		assert.Equal(t, domain.AppTypeAPI, result.AppType)
		assert.Equal(t, "public", result.Details["api_type"])
	})

	t.Run("selects mobile app", func(t *testing.T) {
		mock := NewMockPrompter()
		mock.SelectResponses["What type of application are you building?"] = "mobile"
		mock.SelectResponses["Mobile Platform"] = "both"
		mock.SelectResponses["Mobile Development Approach"] = "cross-platform"

		questions := NewAppTypeQuestions(mock)
		result, err := questions.Ask()

		require.NoError(t, err)
		assert.Equal(t, domain.AppTypeMobile, result.AppType)
		assert.Equal(t, "both", result.Details["platform"])
		assert.Equal(t, "cross-platform", result.Details["approach"])
	})
}

func TestGetAppTypeOptions(t *testing.T) {
	options := GetAppTypeOptions()

	assert.NotEmpty(t, options)

	// Verify some expected options exist
	keys := make([]string, len(options))
	for i, opt := range options {
		keys[i] = opt.Key
	}

	assert.Contains(t, keys, "web")
	assert.Contains(t, keys, "api")
	assert.Contains(t, keys, "cli")
	assert.Contains(t, keys, "mobile")
	assert.Contains(t, keys, "desktop")
	assert.Contains(t, keys, "library")
	assert.Contains(t, keys, "fullstack")
	assert.Contains(t, keys, "microservices")
}
