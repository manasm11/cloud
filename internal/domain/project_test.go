package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProject_Validate(t *testing.T) {
	tests := []struct {
		name    string
		project Project
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid project",
			project: Project{
				Name:        "my-app",
				Description: "A test application",
				Problem:     "Solves testing problems",
			},
			wantErr: false,
		},
		{
			name: "valid with numbers",
			project: Project{
				Name:        "app123",
				Description: "Test app",
				Problem:     "Test problem",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			project: Project{
				Name:        "",
				Description: "Test app",
			},
			wantErr: true,
			errMsg:  "project name is required",
		},
		{
			name: "invalid name with spaces",
			project: Project{
				Name:        "my app",
				Description: "Test app",
			},
			wantErr: true,
			errMsg:  "project name can only contain lowercase letters, numbers, and hyphens",
		},
		{
			name: "invalid name with uppercase",
			project: Project{
				Name:        "MyApp",
				Description: "Test app",
			},
			wantErr: true,
			errMsg:  "project name can only contain lowercase letters, numbers, and hyphens",
		},
		{
			name: "invalid name with special chars",
			project: Project{
				Name:        "my@app",
				Description: "Test app",
			},
			wantErr: true,
			errMsg:  "project name can only contain lowercase letters, numbers, and hyphens",
		},
		{
			name: "name too short",
			project: Project{
				Name:        "ab",
				Description: "Test app",
			},
			wantErr: true,
			errMsg:  "project name must be at least 3 characters",
		},
		{
			name: "name too long",
			project: Project{
				Name:        strings.Repeat("a", 65),
				Description: "Test app",
			},
			wantErr: true,
			errMsg:  "project name must be at most 64 characters",
		},
		{
			name: "empty description",
			project: Project{
				Name:        "my-app",
				Description: "",
			},
			wantErr: true,
			errMsg:  "project description is required",
		},
		{
			name: "name starting with hyphen",
			project: Project{
				Name:        "-myapp",
				Description: "Test app",
			},
			wantErr: true,
			errMsg:  "project name cannot start or end with a hyphen",
		},
		{
			name: "name ending with hyphen",
			project: Project{
				Name:        "myapp-",
				Description: "Test app",
			},
			wantErr: true,
			errMsg:  "project name cannot start or end with a hyphen",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.project.Validate()
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAppType_String(t *testing.T) {
	tests := []struct {
		appType AppType
		want    string
	}{
		{AppTypeWebApp, "Web Application"},
		{AppTypeAPI, "REST API / Backend Service"},
		{AppTypeCLI, "CLI Tool"},
		{AppTypeDesktop, "Desktop Application"},
		{AppTypeMobile, "Mobile Application"},
		{AppTypeLibrary, "Library / Package"},
		{AppTypeFullStack, "Full-Stack Application"},
		{AppTypeMicroservices, "Microservices System"},
		{AppTypeDataPipeline, "Data Pipeline / ETL"},
		{AppTypeOther, "Other"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.appType.String())
		})
	}
}

func TestAppType_Description(t *testing.T) {
	tests := []struct {
		appType AppType
		contain string
	}{
		{AppTypeWebApp, "Browser-based"},
		{AppTypeAPI, "HTTP API"},
		{AppTypeCLI, "Command-line"},
		{AppTypeDesktop, "Native application"},
		{AppTypeMobile, "iOS and/or Android"},
		{AppTypeLibrary, "Reusable code"},
		{AppTypeFullStack, "Combined frontend and backend"},
		{AppTypeMicroservices, "Multiple coordinated services"},
		{AppTypeDataPipeline, "Data processing"},
	}

	for _, tt := range tests {
		t.Run(tt.appType.String(), func(t *testing.T) {
			assert.Contains(t, tt.appType.Description(), tt.contain)
		})
	}
}

func TestNewProject(t *testing.T) {
	p := NewProject("test-app", "A test application")

	assert.Equal(t, "test-app", p.Name)
	assert.Equal(t, "A test application", p.Description)
	assert.NotNil(t, p.Features)
	assert.NotNil(t, p.UserRoles)
	assert.NotNil(t, p.Entities)
	assert.NotNil(t, p.Integrations)
}
