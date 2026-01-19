package suggester

import (
	"testing"

	"github.com/manas/cloud/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestLanguageSuggester_Suggest(t *testing.T) {
	suggester := NewLanguageSuggester()

	t.Run("suggests Go for CLI tools", func(t *testing.T) {
		project := &domain.Project{
			AppType: domain.AppTypeCLI,
		}

		suggestions := suggester.Suggest(project)

		assert.NotEmpty(t, suggestions)
		// Go should be the top recommendation for CLI
		assert.Equal(t, "Go", suggestions[0].Name)
		assert.True(t, suggestions[0].IsRecommended)
	})

	t.Run("suggests Go for API backends", func(t *testing.T) {
		project := &domain.Project{
			AppType: domain.AppTypeAPI,
		}

		suggestions := suggester.Suggest(project)

		assert.NotEmpty(t, suggestions)
		// Go should be recommended for APIs
		foundGo := false
		for _, s := range suggestions {
			if s.Name == "Go" {
				foundGo = true
				break
			}
		}
		assert.True(t, foundGo)
	})

	t.Run("suggests TypeScript for web apps", func(t *testing.T) {
		project := &domain.Project{
			AppType: domain.AppTypeWebApp,
		}

		suggestions := suggester.Suggest(project)

		assert.NotEmpty(t, suggestions)
		// TypeScript should be among suggestions
		foundTS := false
		for _, s := range suggestions {
			if s.Name == "TypeScript" {
				foundTS = true
				break
			}
		}
		assert.True(t, foundTS)
	})

	t.Run("suggests Flutter/Dart for mobile", func(t *testing.T) {
		project := &domain.Project{
			AppType: domain.AppTypeMobile,
			AppTypeDetails: map[string]string{
				"approach": "cross-platform",
			},
		}

		suggestions := suggester.Suggest(project)

		assert.NotEmpty(t, suggestions)
		foundDart := false
		for _, s := range suggestions {
			if s.Name == "Dart" || s.Name == "TypeScript" {
				foundDart = true
				break
			}
		}
		assert.True(t, foundDart)
	})

	t.Run("suggestions have descriptions", func(t *testing.T) {
		project := &domain.Project{
			AppType: domain.AppTypeCLI,
		}

		suggestions := suggester.Suggest(project)

		for _, s := range suggestions {
			assert.NotEmpty(t, s.Name, "suggestion should have a name")
			assert.NotEmpty(t, s.Description, "suggestion should have a description")
			assert.NotEmpty(t, s.Pros, "suggestion should have pros")
			assert.NotEmpty(t, s.Cons, "suggestion should have cons")
		}
	})
}
