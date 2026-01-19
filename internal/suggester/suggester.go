// Package suggester provides technology stack suggestions based on project requirements.
package suggester

import (
	"github.com/manas/cloud/internal/domain"
)

// Suggester provides technology recommendations based on project requirements.
type Suggester interface {
	// SuggestLanguages returns language suggestions for the project
	SuggestLanguages(project *domain.Project) []domain.Suggestion

	// SuggestFrameworks returns framework suggestions for the chosen language
	SuggestFrameworks(project *domain.Project, language string) []domain.Suggestion

	// SuggestDatabases returns database suggestions based on data requirements
	SuggestDatabases(project *domain.Project) []domain.Suggestion

	// SuggestInfrastructure returns infrastructure suggestions
	SuggestInfrastructure(project *domain.Project) InfrastructureSuggestions
}

// InfrastructureSuggestions contains all infrastructure recommendations.
type InfrastructureSuggestions struct {
	Cache        []domain.Suggestion
	MessageQueue []domain.Suggestion
	Search       []domain.Suggestion
	Monitoring   []domain.Suggestion
	Logging      []domain.Suggestion
}

// DefaultSuggester is the default implementation of Suggester.
type DefaultSuggester struct {
	languageSuggester       *LanguageSuggester
	frameworkSuggester      *FrameworkSuggester
	databaseSuggester       *DatabaseSuggester
	infrastructureSuggester *InfrastructureSuggester
}

// NewDefaultSuggester creates a new DefaultSuggester.
func NewDefaultSuggester() *DefaultSuggester {
	return &DefaultSuggester{
		languageSuggester:       NewLanguageSuggester(),
		frameworkSuggester:      NewFrameworkSuggester(),
		databaseSuggester:       NewDatabaseSuggester(),
		infrastructureSuggester: NewInfrastructureSuggester(),
	}
}

// SuggestLanguages returns language suggestions for the project.
func (s *DefaultSuggester) SuggestLanguages(project *domain.Project) []domain.Suggestion {
	return s.languageSuggester.Suggest(project)
}

// SuggestFrameworks returns framework suggestions for the chosen language.
func (s *DefaultSuggester) SuggestFrameworks(project *domain.Project, language string) []domain.Suggestion {
	return s.frameworkSuggester.Suggest(project, language)
}

// SuggestDatabases returns database suggestions based on data requirements.
func (s *DefaultSuggester) SuggestDatabases(project *domain.Project) []domain.Suggestion {
	return s.databaseSuggester.Suggest(project)
}

// SuggestInfrastructure returns infrastructure suggestions.
func (s *DefaultSuggester) SuggestInfrastructure(project *domain.Project) InfrastructureSuggestions {
	return s.infrastructureSuggester.Suggest(project)
}

// GenerateTechStack generates a complete tech stack based on project requirements.
func (s *DefaultSuggester) GenerateTechStack(project *domain.Project) *domain.TechStack {
	stack := &domain.TechStack{}

	// Get language suggestions and pick the top one
	languages := s.SuggestLanguages(project)
	if len(languages) > 0 {
		stack.Language = languages[0].Name
	}

	// Get framework suggestions
	frameworks := s.SuggestFrameworks(project, stack.Language)
	for _, fw := range frameworks {
		if fw.IsRecommended {
			stack.Frameworks = append(stack.Frameworks, domain.Framework{
				Name:      fw.Name,
				Purpose:   fw.Category,
				Rationale: fw.Description,
			})
		}
	}

	// Get database suggestions
	databases := s.SuggestDatabases(project)
	if len(databases) > 0 {
		stack.Database = databases[0].Name
	}

	// Get infrastructure suggestions
	infra := s.SuggestInfrastructure(project)
	if len(infra.Cache) > 0 {
		stack.Cache = infra.Cache[0].Name
	}
	if len(infra.MessageQueue) > 0 {
		stack.MessageQueue = infra.MessageQueue[0].Name
	}
	if len(infra.Monitoring) > 0 {
		stack.Monitoring = infra.Monitoring[0].Name
	}
	if len(infra.Logging) > 0 {
		stack.Logging = infra.Logging[0].Name
	}

	return stack
}
