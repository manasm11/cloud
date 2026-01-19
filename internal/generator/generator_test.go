package generator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/manas/cloud/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClaudeMDGenerator_Render(t *testing.T) {
	gen := NewClaudeMDGenerator()

	project := &domain.Project{
		Name:        "test-app",
		Description: "A test application",
		Problem:     "Testing the generator",
		TargetUsers: "Developers",
		AppType:     domain.AppTypeCLI,
		Features: []domain.Feature{
			{Name: "Feature 1", Description: "First feature", IsMVP: true},
			{Name: "Feature 2", Description: "Second feature", IsMVP: false},
		},
		TechStack: domain.TechStack{
			Language: "Go",
			Database: "PostgreSQL",
		},
	}

	content, err := gen.Render(project)

	require.NoError(t, err)
	assert.Contains(t, content, "# CLAUDE.md - test-app")
	assert.Contains(t, content, "A test application")
	assert.Contains(t, content, "Feature 1")
	assert.Contains(t, content, "Feature 2")
	assert.Contains(t, content, "Go")
	assert.Contains(t, content, "PostgreSQL")
}

func TestMakefileGenerator_Generate(t *testing.T) {
	gen := NewMakefileGenerator()

	t.Run("generates Go makefile", func(t *testing.T) {
		project := &domain.Project{
			Name: "my-go-app",
			TechStack: domain.TechStack{
				Language: "Go",
			},
		}

		tempDir := t.TempDir()
		err := gen.Generate(project, tempDir)

		require.NoError(t, err)

		content, err := os.ReadFile(filepath.Join(tempDir, "Makefile"))
		require.NoError(t, err)

		assert.Contains(t, string(content), "my-go-app")
		assert.Contains(t, string(content), "go build")
		assert.Contains(t, string(content), "go test")
	})

	t.Run("generates TypeScript makefile", func(t *testing.T) {
		project := &domain.Project{
			Name: "my-ts-app",
			TechStack: domain.TechStack{
				Language: "TypeScript",
			},
		}

		tempDir := t.TempDir()
		err := gen.Generate(project, tempDir)

		require.NoError(t, err)

		content, err := os.ReadFile(filepath.Join(tempDir, "Makefile"))
		require.NoError(t, err)

		assert.Contains(t, string(content), "my-ts-app")
		assert.Contains(t, string(content), "npm")
	})

	t.Run("generates Flutter makefile for cross-platform", func(t *testing.T) {
		project := &domain.Project{
			Name:    "my-flutter-app",
			AppType: domain.AppTypeCrossPlatform,
			AppTypeDetails: map[string]string{
				"framework": "flutter",
			},
		}

		tempDir := t.TempDir()
		err := gen.Generate(project, tempDir)

		require.NoError(t, err)

		content, err := os.ReadFile(filepath.Join(tempDir, "Makefile"))
		require.NoError(t, err)

		assert.Contains(t, string(content), "my-flutter-app")
		assert.Contains(t, string(content), "flutter run")
		assert.Contains(t, string(content), "flutter build")
	})

	t.Run("generates Tauri makefile for cross-platform", func(t *testing.T) {
		project := &domain.Project{
			Name:    "my-tauri-app",
			AppType: domain.AppTypeCrossPlatform,
			AppTypeDetails: map[string]string{
				"framework": "tauri",
			},
		}

		tempDir := t.TempDir()
		err := gen.Generate(project, tempDir)

		require.NoError(t, err)

		content, err := os.ReadFile(filepath.Join(tempDir, "Makefile"))
		require.NoError(t, err)

		assert.Contains(t, string(content), "my-tauri-app")
		assert.Contains(t, string(content), "tauri dev")
		assert.Contains(t, string(content), "tauri build")
	})

	t.Run("generates Wails makefile for cross-platform", func(t *testing.T) {
		project := &domain.Project{
			Name:    "my-wails-app",
			AppType: domain.AppTypeCrossPlatform,
			AppTypeDetails: map[string]string{
				"framework": "wails",
			},
		}

		tempDir := t.TempDir()
		err := gen.Generate(project, tempDir)

		require.NoError(t, err)

		content, err := os.ReadFile(filepath.Join(tempDir, "Makefile"))
		require.NoError(t, err)

		assert.Contains(t, string(content), "my-wails-app")
		assert.Contains(t, string(content), "wails dev")
		assert.Contains(t, string(content), "wails build")
	})
}

func TestDockerfileGenerator_Generate(t *testing.T) {
	gen := NewDockerfileGenerator()

	t.Run("generates Go dockerfile", func(t *testing.T) {
		project := &domain.Project{
			Name: "my-go-app",
			TechStack: domain.TechStack{
				Language: "Go",
			},
		}

		tempDir := t.TempDir()
		err := gen.Generate(project, tempDir)

		require.NoError(t, err)

		content, err := os.ReadFile(filepath.Join(tempDir, "Dockerfile"))
		require.NoError(t, err)

		assert.Contains(t, string(content), "golang:1.22-alpine")
		assert.Contains(t, string(content), "my-go-app")
	})
}

func TestComposeGenerator_Generate(t *testing.T) {
	gen := NewComposeGenerator()

	t.Run("generates compose with postgres", func(t *testing.T) {
		project := &domain.Project{
			Name: "my-app",
			TechStack: domain.TechStack{
				Language: "Go",
				Database: "PostgreSQL",
				Cache:    "Redis",
			},
		}

		tempDir := t.TempDir()
		err := gen.Generate(project, tempDir)

		require.NoError(t, err)

		content, err := os.ReadFile(filepath.Join(tempDir, "docker-compose.yml"))
		require.NoError(t, err)

		assert.Contains(t, string(content), "my-app")
		assert.Contains(t, string(content), "postgres:16-alpine")
		assert.Contains(t, string(content), "redis:7-alpine")
	})
}

func TestFileGenerator_GenerateAll(t *testing.T) {
	gen := NewFileGenerator()

	project := &domain.Project{
		Name:        "full-app",
		Description: "A complete application",
		AppType:     domain.AppTypeAPI,
		Deployment: domain.DeploymentConfig{
			UseDocker: true,
		},
		TechStack: domain.TechStack{
			Language: "Go",
			Database: "PostgreSQL",
		},
	}

	tempDir := t.TempDir()
	err := gen.GenerateAll(project, tempDir)

	require.NoError(t, err)

	// Check all files were created
	files := []string{"CLAUDE.md", "Makefile", "Dockerfile", "docker-compose.yml"}
	for _, f := range files {
		_, err := os.Stat(filepath.Join(tempDir, f))
		assert.NoError(t, err, "file %s should exist", f)
	}
}
