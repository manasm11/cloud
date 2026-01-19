package suggester

import (
	"github.com/manas/cloud/internal/domain"
)

// InfrastructureSuggester suggests infrastructure components.
type InfrastructureSuggester struct{}

// NewInfrastructureSuggester creates a new InfrastructureSuggester.
func NewInfrastructureSuggester() *InfrastructureSuggester {
	return &InfrastructureSuggester{}
}

// Suggest returns infrastructure suggestions for the project.
func (s *InfrastructureSuggester) Suggest(project *domain.Project) InfrastructureSuggestions {
	return InfrastructureSuggestions{
		Cache:        s.suggestCache(project),
		MessageQueue: s.suggestMessageQueue(project),
		Search:       s.suggestSearch(project),
		Monitoring:   s.suggestMonitoring(project),
		Logging:      s.suggestLogging(project),
	}
}

func (s *InfrastructureSuggester) suggestCache(project *domain.Project) []domain.Suggestion {
	suggestions := []domain.Suggestion{
		{
			Name:        "Redis",
			Category:    "cache",
			Description: "In-memory data store, excellent for caching and session storage.",
			Pros:        []string{"Very fast", "Versatile", "Widely supported", "Data structures"},
			Cons:        []string{"Memory bound", "Requires management"},
			WhenToUse:   "Most caching needs, session storage, rate limiting.",
			Confidence:  0.9,
			IsRecommended: true,
		},
	}

	// For simpler apps, suggest in-memory
	if project.Performance.ConcurrentUsers == "small" {
		suggestions = append([]domain.Suggestion{
			{
				Name:        "In-Memory Cache",
				Category:    "cache",
				Description: "Simple in-process caching, no external dependencies.",
				Pros:        []string{"No dependencies", "Simple", "Fast"},
				Cons:        []string{"Not shared between instances", "Lost on restart"},
				WhenToUse:   "Small applications with single instance deployment.",
				Confidence:  0.7,
				IsRecommended: true,
			},
		}, suggestions...)
	}

	return suggestions
}

func (s *InfrastructureSuggester) suggestMessageQueue(project *domain.Project) []domain.Suggestion {
	var suggestions []domain.Suggestion

	// Only suggest message queues for certain app types
	if project.AppType != domain.AppTypeAPI &&
		project.AppType != domain.AppTypeMicroservices &&
		project.AppType != domain.AppTypeFullStack {
		return suggestions
	}

	suggestions = []domain.Suggestion{
		{
			Name:        "Redis Streams",
			Category:    "queue",
			Description: "Redis-based message streaming, good for simpler use cases.",
			Pros:        []string{"If already using Redis", "Simple", "Fast"},
			Cons:        []string{"Less features than dedicated MQ"},
			WhenToUse:   "Simple message passing when already using Redis.",
			Confidence:  0.7,
			IsRecommended: false,
		},
		{
			Name:        "RabbitMQ",
			Category:    "queue",
			Description: "Robust message broker with flexible routing.",
			Pros:        []string{"Mature", "Flexible routing", "Many patterns", "Good UI"},
			Cons:        []string{"Erlang dependency", "Resource usage"},
			WhenToUse:   "Complex routing needs, enterprise patterns.",
			Confidence:  0.8,
			IsRecommended: true,
		},
	}

	// For large scale, suggest Kafka
	if project.Performance.ConcurrentUsers == "xlarge" {
		suggestions = append(suggestions, domain.Suggestion{
			Name:        "Apache Kafka",
			Category:    "queue",
			Description: "Distributed event streaming platform for high-throughput.",
			Pros:        []string{"High throughput", "Durability", "Replay capability"},
			Cons:        []string{"Complex", "Resource intensive", "Operational overhead"},
			WhenToUse:   "High-volume event streaming, event sourcing.",
			Confidence:  0.8,
			IsRecommended: true,
		})
	}

	return suggestions
}

func (s *InfrastructureSuggester) suggestSearch(project *domain.Project) []domain.Suggestion {
	if !project.DataCharacteristics.FullTextSearch {
		return nil
	}

	return []domain.Suggestion{
		{
			Name:        "Meilisearch",
			Category:    "search",
			Description: "Fast, typo-tolerant search engine. Easy to deploy and use.",
			Pros:        []string{"Very fast", "Easy setup", "Typo tolerance", "Good DX"},
			Cons:        []string{"Smaller dataset limits", "Fewer features"},
			WhenToUse:   "Product search, site search, simpler use cases.",
			Confidence:  0.8,
			IsRecommended: true,
		},
		{
			Name:        "Elasticsearch",
			Category:    "search",
			Description: "Powerful search and analytics engine.",
			Pros:        []string{"Feature rich", "Scalable", "Analytics", "Ecosystem"},
			Cons:        []string{"Complex", "Resource heavy", "JVM"},
			WhenToUse:   "Complex search needs, log analytics, large datasets.",
			Confidence:  0.7,
			IsRecommended: false,
		},
	}
}

func (s *InfrastructureSuggester) suggestMonitoring(project *domain.Project) []domain.Suggestion {
	suggestions := []domain.Suggestion{
		{
			Name:        "Prometheus + Grafana",
			Category:    "monitoring",
			Description: "Open-source monitoring stack. Industry standard for metrics.",
			Pros:        []string{"Open source", "Flexible", "Large community", "Good dashboards"},
			Cons:        []string{"Self-hosted complexity", "Storage management"},
			WhenToUse:   "Self-hosted monitoring with full control.",
			Confidence:  0.8,
			IsRecommended: true,
		},
	}

	// Cloud-based options
	if project.Deployment.CloudProvider != "selfhost" {
		suggestions = append(suggestions, domain.Suggestion{
			Name:        "DataDog",
			Category:    "monitoring",
			Description: "Cloud-based monitoring and analytics platform.",
			Pros:        []string{"Easy setup", "Great UI", "APM included", "Many integrations"},
			Cons:        []string{"Expensive at scale", "Vendor lock-in"},
			WhenToUse:   "When you want managed monitoring without operational overhead.",
			Confidence:  0.7,
			IsRecommended: false,
		})
	}

	return suggestions
}

func (s *InfrastructureSuggester) suggestLogging(project *domain.Project) []domain.Suggestion {
	return []domain.Suggestion{
		{
			Name:        "Structured Logging (JSON)",
			Category:    "logging",
			Description: "JSON-formatted logs for easy parsing and analysis.",
			Pros:        []string{"Machine readable", "Easy to search", "Context rich"},
			Cons:        []string{"Less human readable", "Slightly larger"},
			WhenToUse:   "Production applications with log aggregation.",
			Confidence:  0.9,
			IsRecommended: true,
		},
		{
			Name:        "Loki",
			Category:    "logging",
			Description: "Log aggregation system designed to work with Grafana.",
			Pros:        []string{"Works with Grafana", "Label-based", "Cost effective"},
			Cons:        []string{"Less powerful than ELK", "Newer"},
			WhenToUse:   "When using Grafana for monitoring.",
			Confidence:  0.7,
			IsRecommended: false,
		},
	}
}
