// Package generator provides file generation from project specifications.
package generator

import (
	"os"
	"path/filepath"

	"github.com/manas/cloud/internal/domain"
	"github.com/manas/cloud/internal/suggester"
)

// Generator generates project files from a project specification.
type Generator interface {
	Generate(project *domain.Project, outputDir string) error
}

// FileGenerator generates all project files.
type FileGenerator struct {
	suggester *suggester.DefaultSuggester
}

// NewFileGenerator creates a new FileGenerator.
func NewFileGenerator() *FileGenerator {
	return &FileGenerator{
		suggester: suggester.NewDefaultSuggester(),
	}
}

// GenerateAll generates all project files.
func (g *FileGenerator) GenerateAll(project *domain.Project, outputDir string) error {
	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	// Generate tech stack suggestions if not already set
	if project.TechStack.Language == "" {
		project.TechStack = *g.suggester.GenerateTechStack(project)
	}

	// Generate CLAUDE.md
	claudeGen := NewClaudeMDGenerator()
	if err := claudeGen.Generate(project, outputDir); err != nil {
		return err
	}

	// Generate Makefile
	makefileGen := NewMakefileGenerator()
	if err := makefileGen.Generate(project, outputDir); err != nil {
		return err
	}

	// Generate Dockerfile if Docker is enabled
	if project.Deployment.UseDocker {
		dockerGen := NewDockerfileGenerator()
		if err := dockerGen.Generate(project, outputDir); err != nil {
			return err
		}

		composeGen := NewComposeGenerator()
		if err := composeGen.Generate(project, outputDir); err != nil {
			return err
		}
	}

	return nil
}

// GeneratedFile represents a generated file.
type GeneratedFile struct {
	Path    string
	Content string
}

// WriteFile writes a generated file to disk.
func WriteFile(outputDir, filename, content string) error {
	path := filepath.Join(outputDir, filename)

	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(content), 0644)
}
