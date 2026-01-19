package interview

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/manas/cloud/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSession(t *testing.T) {
	session := NewSession()

	assert.NotNil(t, session)
	assert.NotNil(t, session.Project)
	assert.Equal(t, StepBasics, session.CurrentStep)
	assert.False(t, session.IsComplete)
}

func TestSession_SetProjectBasics(t *testing.T) {
	session := NewSession()

	session.SetProjectName("my-app")
	session.SetDescription("A test application")
	session.SetProblem("Solves testing problems")
	session.SetTargetUsers("Developers")

	assert.Equal(t, "my-app", session.Project.Name)
	assert.Equal(t, "A test application", session.Project.Description)
	assert.Equal(t, "Solves testing problems", session.Project.Problem)
	assert.Equal(t, "Developers", session.Project.TargetUsers)
}

func TestSession_SetAppType(t *testing.T) {
	session := NewSession()

	session.SetAppType(domain.AppTypeWebApp)

	assert.Equal(t, domain.AppTypeWebApp, session.Project.AppType)
}

func TestSession_AddFeature(t *testing.T) {
	session := NewSession()

	feature := domain.Feature{
		Name:        "User Authentication",
		Description: "Users can log in and out",
		IsMVP:       true,
	}
	session.AddFeature(feature)

	assert.Len(t, session.Project.Features, 1)
	assert.Equal(t, "User Authentication", session.Project.Features[0].Name)
}

func TestSession_SaveAndLoad(t *testing.T) {
	// Create a session with data
	session := NewSession()
	session.SetProjectName("test-project")
	session.SetDescription("A test project")
	session.SetAppType(domain.AppTypeCLI)
	session.AddFeature(domain.Feature{
		Name:        "Feature 1",
		Description: "First feature",
		IsMVP:       true,
	})
	session.CurrentStep = StepFeatures

	// Save to temp file
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "session.json")

	err := session.Save(path)
	require.NoError(t, err)

	// Verify file exists
	_, err = os.Stat(path)
	require.NoError(t, err)

	// Load the session
	loaded, err := LoadSession(path)
	require.NoError(t, err)

	// Verify data
	assert.Equal(t, "test-project", loaded.Project.Name)
	assert.Equal(t, "A test project", loaded.Project.Description)
	assert.Equal(t, domain.AppTypeCLI, loaded.Project.AppType)
	assert.Len(t, loaded.Project.Features, 1)
	assert.Equal(t, StepFeatures, loaded.CurrentStep)
}

func TestSession_LoadNonExistent(t *testing.T) {
	_, err := LoadSession("/nonexistent/path/session.json")
	assert.Error(t, err)
}

func TestSession_AdvanceStep(t *testing.T) {
	session := NewSession()

	assert.Equal(t, StepBasics, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepAppType, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepFeatures, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepUserRoles, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepData, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepIntegrations, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepPerformance, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepSecurity, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepDeployment, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepDevPractices, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepTechStack, session.CurrentStep)

	session.AdvanceStep()
	assert.Equal(t, StepComplete, session.CurrentStep)
	assert.True(t, session.IsComplete)
}

func TestSession_StepName(t *testing.T) {
	tests := []struct {
		step Step
		name string
	}{
		{StepBasics, "Project Basics"},
		{StepAppType, "Application Type"},
		{StepFeatures, "Features"},
		{StepUserRoles, "User Roles"},
		{StepData, "Data Requirements"},
		{StepIntegrations, "Integrations"},
		{StepPerformance, "Performance"},
		{StepSecurity, "Security"},
		{StepDeployment, "Deployment"},
		{StepDevPractices, "Development Practices"},
		{StepTechStack, "Tech Stack"},
		{StepComplete, "Complete"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.name, tt.step.String())
		})
	}
}

func TestSession_Progress(t *testing.T) {
	session := NewSession()

	// At start
	progress := session.Progress()
	assert.Equal(t, 0, progress.CurrentStep)
	assert.Equal(t, 11, progress.TotalSteps)
	assert.InDelta(t, 0.0, progress.Percentage, 0.01)

	// After a few steps
	session.AdvanceStep() // AppType
	session.AdvanceStep() // Features
	session.AdvanceStep() // UserRoles

	progress = session.Progress()
	assert.Equal(t, 3, progress.CurrentStep)
	assert.InDelta(t, 27.27, progress.Percentage, 0.1)
}

func TestSession_CanGoBack(t *testing.T) {
	session := NewSession()

	// Can't go back from first step
	assert.False(t, session.CanGoBack())

	session.AdvanceStep()
	assert.True(t, session.CanGoBack())

	session.GoBack()
	assert.Equal(t, StepBasics, session.CurrentStep)
	assert.False(t, session.CanGoBack())
}
