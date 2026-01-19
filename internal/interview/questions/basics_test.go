package questions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasicsQuestions_Ask(t *testing.T) {
	t.Run("collects all basic information", func(t *testing.T) {
		mock := NewMockPrompter()
		mock.TextResponses["Project Name"] = "my-app"
		mock.TextResponses["Project Description"] = "A test application"
		mock.TextResponses["Problem Statement"] = "It solves testing problems"
		mock.TextResponses["Target Users"] = "Developers and testers"
		mock.ConfirmResponses["Is this a new project?"] = true

		questions := NewBasicsQuestions(mock)
		result, err := questions.Ask()

		require.NoError(t, err)
		assert.Equal(t, "my-app", result.Name)
		assert.Equal(t, "A test application", result.Description)
		assert.Equal(t, "It solves testing problems", result.Problem)
		assert.Equal(t, "Developers and testers", result.TargetUsers)
		assert.True(t, result.IsNewProject)
	})

	t.Run("validates project name", func(t *testing.T) {
		mock := NewMockPrompter()
		mock.TextResponses["Project Name"] = "Invalid Name" // Has uppercase and space

		questions := NewBasicsQuestions(mock)
		_, err := questions.Ask()

		// Should fail validation
		assert.Error(t, err)
	})

	t.Run("accepts existing project", func(t *testing.T) {
		mock := NewMockPrompter()
		mock.TextResponses["Project Name"] = "legacy-app"
		mock.TextResponses["Project Description"] = "An existing application"
		mock.TextResponses["Problem Statement"] = "Needs modernization"
		mock.TextResponses["Target Users"] = "Internal teams"
		mock.ConfirmResponses["Is this a new project?"] = false

		questions := NewBasicsQuestions(mock)
		result, err := questions.Ask()

		require.NoError(t, err)
		assert.Equal(t, "legacy-app", result.Name)
		assert.False(t, result.IsNewProject)
	})
}

func TestBasicsResult_Validate(t *testing.T) {
	tests := []struct {
		name    string
		result  BasicsResult
		wantErr bool
	}{
		{
			name: "valid result",
			result: BasicsResult{
				Name:        "my-app",
				Description: "A test app",
				Problem:     "Testing",
				TargetUsers: "Developers",
			},
			wantErr: false,
		},
		{
			name: "missing name",
			result: BasicsResult{
				Name:        "",
				Description: "A test app",
			},
			wantErr: true,
		},
		{
			name: "missing description",
			result: BasicsResult{
				Name:        "my-app",
				Description: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.result.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
