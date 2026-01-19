package suggester

import (
	"github.com/manas/cloud/internal/domain"
)

// DatabaseSuggester suggests databases based on data requirements.
type DatabaseSuggester struct {
	databases map[string]databaseInfo
}

type databaseInfo struct {
	name        string
	category    string // sql, nosql, cache, search, timeseries
	description string
	pros        []string
	cons        []string
	whenToUse   string
}

// NewDatabaseSuggester creates a new DatabaseSuggester.
func NewDatabaseSuggester() *DatabaseSuggester {
	return &DatabaseSuggester{
		databases: initDatabases(),
	}
}

func initDatabases() map[string]databaseInfo {
	return map[string]databaseInfo{
		"postgresql": {
			name:        "PostgreSQL",
			category:    "sql",
			description: "The world's most advanced open source relational database.",
			pros: []string{
				"ACID compliant",
				"Advanced features (JSON, full-text search)",
				"Excellent performance",
				"Strong community",
				"Extensions ecosystem",
				"Great for complex queries",
			},
			cons: []string{
				"More complex than MySQL",
				"Can be resource-intensive",
				"Replication setup complexity",
			},
			whenToUse: "Most applications. Best default choice for relational data.",
		},
		"mysql": {
			name:        "MySQL",
			category:    "sql",
			description: "The world's most popular open source database.",
			pros: []string{
				"Very popular",
				"Easy to use",
				"Good documentation",
				"Many hosting options",
				"Large community",
			},
			cons: []string{
				"Fewer advanced features than PostgreSQL",
				"Oracle ownership concerns",
				"Storage engine complexity",
			},
			whenToUse: "When you need wide hosting support or have MySQL experience.",
		},
		"sqlite": {
			name:        "SQLite",
			category:    "sql",
			description: "Self-contained, serverless, zero-configuration SQL database.",
			pros: []string{
				"No server needed",
				"Zero configuration",
				"Single file storage",
				"Great for embedded use",
				"Perfect for development",
			},
			cons: []string{
				"No concurrent writes",
				"Limited scalability",
				"No user management",
			},
			whenToUse: "CLI tools, desktop apps, development, or single-user applications.",
		},
		"mongodb": {
			name:        "MongoDB",
			category:    "nosql",
			description: "Document database with the scalability and flexibility you need.",
			pros: []string{
				"Flexible schema",
				"Good for rapid development",
				"Horizontal scaling",
				"JSON-like documents",
				"Good for unstructured data",
			},
			cons: []string{
				"No ACID transactions (historically)",
				"Memory intensive",
				"Joins are inefficient",
				"License concerns",
			},
			whenToUse: "When you need flexible schemas or document-oriented storage.",
		},
		"redis": {
			name:        "Redis",
			category:    "cache",
			description: "In-memory data structure store, used as database, cache, and message broker.",
			pros: []string{
				"Extremely fast",
				"Versatile data structures",
				"Pub/sub support",
				"Great for caching",
				"Simple to use",
			},
			cons: []string{
				"Memory-bound",
				"Data loss risk without persistence",
				"Single-threaded",
			},
			whenToUse: "Caching, session storage, real-time features, rate limiting.",
		},
		"elasticsearch": {
			name:        "Elasticsearch",
			category:    "search",
			description: "Distributed, RESTful search and analytics engine.",
			pros: []string{
				"Powerful full-text search",
				"Real-time analytics",
				"Horizontal scaling",
				"Rich query language",
			},
			cons: []string{
				"Resource intensive",
				"Complex to operate",
				"JVM dependency",
				"License concerns",
			},
			whenToUse: "When you need advanced full-text search or log analytics.",
		},
		"meilisearch": {
			name:        "Meilisearch",
			category:    "search",
			description: "Lightning fast, ultra-relevant, and typo-tolerant search engine.",
			pros: []string{
				"Very fast",
				"Easy to use",
				"Typo tolerance",
				"Simple deployment",
				"Good documentation",
			},
			cons: []string{
				"Smaller community",
				"Fewer features than Elasticsearch",
				"Dataset size limits",
			},
			whenToUse: "When you need simple, fast search without Elasticsearch complexity.",
		},
		"timescaledb": {
			name:        "TimescaleDB",
			category:    "timeseries",
			description: "Time-series database built on PostgreSQL.",
			pros: []string{
				"PostgreSQL compatible",
				"Automatic partitioning",
				"Compression",
				"Continuous aggregates",
			},
			cons: []string{
				"PostgreSQL overhead",
				"Learning curve for time-series concepts",
			},
			whenToUse: "IoT data, metrics, analytics with time-series data.",
		},
		"dynamodb": {
			name:        "DynamoDB",
			category:    "nosql",
			description: "Fast, flexible NoSQL database service for any scale.",
			pros: []string{
				"Fully managed",
				"Automatic scaling",
				"Low latency",
				"AWS integration",
			},
			cons: []string{
				"AWS lock-in",
				"Query limitations",
				"Can be expensive",
				"Learning curve",
			},
			whenToUse: "AWS-based applications needing managed NoSQL at scale.",
		},
	}
}

// Suggest returns database suggestions for the project.
func (s *DatabaseSuggester) Suggest(project *domain.Project) []domain.Suggestion {
	var suggestions []domain.Suggestion

	// Score databases based on requirements
	scored := s.scoreDatabases(project)

	for _, item := range scored {
		db := s.databases[item.key]
		suggestions = append(suggestions, domain.Suggestion{
			Name:          db.name,
			Category:      db.category,
			Description:   db.description,
			Pros:          db.pros,
			Cons:          db.cons,
			WhenToUse:     db.whenToUse,
			Confidence:    item.score,
			IsRecommended: item.score >= 0.6,
		})
	}

	return suggestions
}

func (s *DatabaseSuggester) scoreDatabases(project *domain.Project) []scoredItem {
	scores := make(map[string]float64)

	// Base scoring by data structure
	switch project.DataCharacteristics.Structure {
	case "structured":
		scores["postgresql"] += 0.5
		scores["mysql"] += 0.3
	case "semi":
		scores["mongodb"] += 0.4
		scores["postgresql"] += 0.3 // JSON support
	case "unstructured":
		scores["mongodb"] += 0.4
	case "mixed":
		scores["postgresql"] += 0.4
	default:
		scores["postgresql"] += 0.4 // Default to PostgreSQL
	}

	// CLI or desktop apps often use SQLite
	if project.AppType == domain.AppTypeCLI || project.AppType == domain.AppTypeDesktop {
		scores["sqlite"] += 0.5
	}

	// Full-text search requirement
	if project.DataCharacteristics.FullTextSearch {
		scores["elasticsearch"] += 0.4
		scores["meilisearch"] += 0.3
		scores["postgresql"] += 0.1 // Has built-in FTS
	}

	// Time series data
	if project.DataCharacteristics.TimeSeries {
		scores["timescaledb"] += 0.5
	}

	// Performance requirements
	if project.Performance.ResponseTime == "realtime" {
		scores["redis"] += 0.3
	}

	// Cloud provider preferences
	if project.Deployment.CloudProvider == "aws" {
		scores["dynamodb"] += 0.2
	}

	// Convert to sorted slice
	var result []scoredItem
	for key, score := range scores {
		if score > 0 {
			result = append(result, scoredItem{key: key, score: score})
		}
	}

	// Sort by score descending
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[j].score > result[i].score {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	return result
}
