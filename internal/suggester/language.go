package suggester

import (
	"github.com/manas/cloud/internal/domain"
)

// LanguageSuggester suggests programming languages based on project requirements.
type LanguageSuggester struct {
	languages map[string]languageInfo
}

type languageInfo struct {
	name        string
	description string
	pros        []string
	cons        []string
	bestFor     []domain.AppType
	whenToUse   string
}

// NewLanguageSuggester creates a new LanguageSuggester.
func NewLanguageSuggester() *LanguageSuggester {
	return &LanguageSuggester{
		languages: initLanguages(),
	}
}

func initLanguages() map[string]languageInfo {
	return map[string]languageInfo{
		"go": {
			name:        "Go",
			description: "Fast, simple, and reliable. Excellent for backend services and CLI tools.",
			pros: []string{
				"Fast compilation and execution",
				"Simple, readable syntax",
				"Excellent concurrency support",
				"Single binary deployment",
				"Strong standard library",
				"Great for microservices",
			},
			cons: []string{
				"Verbose error handling",
				"No generics until recently",
				"Smaller ecosystem than Node/Python",
				"Learning curve for OOP developers",
			},
			bestFor:   []domain.AppType{domain.AppTypeCLI, domain.AppTypeAPI, domain.AppTypeMicroservices, domain.AppTypeCrossPlatform},
			whenToUse: "When you need performance, simplicity, and easy deployment. Great for DevOps tools, APIs, and cross-platform desktop apps with Wails.",
		},
		"typescript": {
			name:        "TypeScript",
			description: "JavaScript with types. Excellent for web applications and full-stack development.",
			pros: []string{
				"Type safety catches errors early",
				"Excellent IDE support",
				"Huge ecosystem (npm)",
				"Works in browser and Node.js",
				"Easy to learn for JS developers",
				"Great for full-stack development",
			},
			cons: []string{
				"Build step required",
				"Type definitions can be complex",
				"Runtime is still JavaScript",
				"Configuration overhead",
			},
			bestFor:   []domain.AppType{domain.AppTypeWebApp, domain.AppTypeFullStack, domain.AppTypeAPI, domain.AppTypeCrossPlatform},
			whenToUse: "When building web applications, sharing code between frontend and backend, or cross-platform apps with Electron/React Native.",
		},
		"python": {
			name:        "Python",
			description: "Versatile and readable. Excellent for data processing, ML, and rapid prototyping.",
			pros: []string{
				"Very readable syntax",
				"Huge ecosystem",
				"Excellent for data science/ML",
				"Rapid development",
				"Great documentation",
				"Easy to learn",
			},
			cons: []string{
				"Slower than compiled languages",
				"GIL limits concurrency",
				"Dynamic typing can hide bugs",
				"Deployment can be complex",
			},
			bestFor:   []domain.AppType{domain.AppTypeDataPipeline, domain.AppTypeAPI, domain.AppTypeLibrary},
			whenToUse: "When you need data processing, machine learning, or rapid prototyping.",
		},
		"rust": {
			name:        "Rust",
			description: "Safe and fast. Excellent for performance-critical and systems programming.",
			pros: []string{
				"Memory safety without garbage collection",
				"Excellent performance",
				"Strong type system",
				"Great tooling (cargo)",
				"Prevents data races",
				"Growing ecosystem",
			},
			cons: []string{
				"Steep learning curve",
				"Slower compilation",
				"Complex ownership model",
				"Smaller ecosystem",
			},
			bestFor:   []domain.AppType{domain.AppTypeCLI, domain.AppTypeAPI, domain.AppTypeLibrary, domain.AppTypeCrossPlatform},
			whenToUse: "When you need maximum performance and memory safety. Good for CLI tools, systems programming, and lightweight cross-platform apps with Tauri.",
		},
		"java": {
			name:        "Java",
			description: "Mature and enterprise-ready. Excellent for large-scale applications.",
			pros: []string{
				"Mature ecosystem",
				"Excellent tooling",
				"Strong enterprise support",
				"Great for large teams",
				"Cross-platform (JVM)",
				"Many libraries available",
			},
			cons: []string{
				"Verbose syntax",
				"Slow startup (improving with GraalVM)",
				"Memory overhead",
				"Configuration complexity",
			},
			bestFor:   []domain.AppType{domain.AppTypeAPI, domain.AppTypeMicroservices, domain.AppTypeDesktop},
			whenToUse: "When building enterprise applications or need mature, stable technology.",
		},
		"dart": {
			name:        "Dart",
			description: "Optimized for UI. The language behind Flutter for cross-platform mobile development.",
			pros: []string{
				"Flutter's native language",
				"Hot reload for fast development",
				"Single codebase for iOS/Android",
				"Good performance",
				"Strong typing",
				"Growing community",
			},
			cons: []string{
				"Smaller ecosystem than JS",
				"Mainly used with Flutter",
				"Limited backend options",
				"Newer language",
			},
			bestFor:   []domain.AppType{domain.AppTypeMobile, domain.AppTypeDesktop, domain.AppTypeCrossPlatform},
			whenToUse: "When building cross-platform mobile apps with Flutter.",
		},
		"kotlin": {
			name:        "Kotlin",
			description: "Modern JVM language. Excellent for Android and backend development.",
			pros: []string{
				"Concise syntax",
				"Null safety built-in",
				"Full Java interoperability",
				"Official Android language",
				"Coroutines for async",
				"Growing ecosystem",
			},
			cons: []string{
				"Compilation slower than Java",
				"Smaller community than Java",
				"Learning curve from Java",
				"Some tooling gaps",
			},
			bestFor:   []domain.AppType{domain.AppTypeMobile, domain.AppTypeAPI, domain.AppTypeDesktop, domain.AppTypeCrossPlatform},
			whenToUse: "When building Android apps, modern JVM-based backends, or Compose Multiplatform apps.",
		},
		"csharp": {
			name:        "C#",
			description: "Powerful and versatile. Excellent for Windows development and games.",
			pros: []string{
				"Strong typing",
				"Excellent tooling (Visual Studio)",
				"Cross-platform with .NET Core",
				"Great for games (Unity)",
				"LINQ for data manipulation",
				"Async/await support",
			},
			cons: []string{
				"Historically Windows-focused",
				"Large runtime",
				"Complex ecosystem",
				"Microsoft dependency",
			},
			bestFor:   []domain.AppType{domain.AppTypeDesktop, domain.AppTypeAPI, domain.AppTypeWebApp, domain.AppTypeCrossPlatform},
			whenToUse: "When building Windows applications, games, enterprise .NET backends, or cross-platform apps with MAUI/Avalonia.",
		},
	}
}

// Suggest returns language suggestions for the project.
func (s *LanguageSuggester) Suggest(project *domain.Project) []domain.Suggestion {
	var suggestions []domain.Suggestion

	// Score each language based on project requirements
	scored := s.scoreLanguages(project)

	// Sort by score and convert to suggestions
	for _, item := range scored {
		lang := s.languages[item.key]
		suggestions = append(suggestions, domain.Suggestion{
			Name:          lang.name,
			Category:      "language",
			Description:   lang.description,
			Pros:          lang.pros,
			Cons:          lang.cons,
			WhenToUse:     lang.whenToUse,
			Confidence:    item.score,
			IsRecommended: item.score >= 0.7,
		})
	}

	return suggestions
}

type scoredItem struct {
	key   string
	score float64
}

func (s *LanguageSuggester) scoreLanguages(project *domain.Project) []scoredItem {
	scores := make(map[string]float64)

	// Base scoring by app type
	for key, lang := range s.languages {
		for _, appType := range lang.bestFor {
			if appType == project.AppType {
				scores[key] += 0.5
				break
			}
		}
	}

	// Adjust for specific requirements
	switch project.AppType {
	case domain.AppTypeCLI:
		scores["go"] += 0.4
		scores["rust"] += 0.3
		scores["python"] += 0.1

	case domain.AppTypeAPI:
		scores["go"] += 0.3
		scores["typescript"] += 0.2
		scores["python"] += 0.2
		scores["java"] += 0.1
		scores["rust"] += 0.1

	case domain.AppTypeWebApp:
		scores["typescript"] += 0.4
		scores["python"] += 0.1

	case domain.AppTypeFullStack:
		scores["typescript"] += 0.5

	case domain.AppTypeMobile:
		if project.AppTypeDetails != nil {
			if project.AppTypeDetails["approach"] == "cross-platform" {
				scores["dart"] += 0.4
				scores["typescript"] += 0.3 // React Native
			} else if project.AppTypeDetails["platform"] == "android" {
				scores["kotlin"] += 0.5
			}
		} else {
			scores["dart"] += 0.3
			scores["kotlin"] += 0.2
			scores["typescript"] += 0.2
		}

	case domain.AppTypeDesktop:
		scores["go"] += 0.2
		scores["csharp"] += 0.3
		scores["typescript"] += 0.2 // Electron
		scores["rust"] += 0.2       // Tauri

	case domain.AppTypeMicroservices:
		scores["go"] += 0.4
		scores["java"] += 0.2
		scores["typescript"] += 0.1

	case domain.AppTypeDataPipeline:
		scores["python"] += 0.5
		scores["go"] += 0.2

	case domain.AppTypeLibrary:
		scores["go"] += 0.2
		scores["rust"] += 0.3
		scores["typescript"] += 0.2
		scores["python"] += 0.2

	case domain.AppTypeCrossPlatform:
		// Score based on selected framework if available
		if project.AppTypeDetails != nil {
			framework := project.AppTypeDetails["framework"]
			switch framework {
			case "flutter":
				scores["dart"] += 0.6
			case "react-native":
				scores["typescript"] += 0.6
			case "electron":
				scores["typescript"] += 0.6
			case "tauri":
				scores["rust"] += 0.5
				scores["typescript"] += 0.3 // frontend
			case "maui", "avalonia":
				scores["csharp"] += 0.6
			case "wails":
				scores["go"] += 0.5
				scores["typescript"] += 0.3 // frontend
			case "compose":
				scores["kotlin"] += 0.6
			default:
				// No framework chosen yet, suggest all options
				scores["dart"] += 0.3       // Flutter
				scores["typescript"] += 0.3 // Electron, React Native
				scores["rust"] += 0.2       // Tauri
				scores["csharp"] += 0.2     // MAUI, Avalonia
				scores["go"] += 0.2         // Wails
				scores["kotlin"] += 0.2     // Compose Multiplatform
			}
		} else {
			// No details, suggest all options
			scores["dart"] += 0.3
			scores["typescript"] += 0.3
			scores["rust"] += 0.2
			scores["csharp"] += 0.2
			scores["go"] += 0.2
			scores["kotlin"] += 0.2
		}
	}

	// Performance requirements boost
	if project.Performance.ResponseTime == "realtime" {
		scores["go"] += 0.2
		scores["rust"] += 0.3
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
