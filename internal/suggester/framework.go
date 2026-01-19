package suggester

import (
	"github.com/manas/cloud/internal/domain"
)

// FrameworkSuggester suggests frameworks based on language and project requirements.
type FrameworkSuggester struct {
	frameworks map[string][]frameworkInfo
}

type frameworkInfo struct {
	name        string
	purpose     string
	description string
	pros        []string
	cons        []string
	whenToUse   string
}

// NewFrameworkSuggester creates a new FrameworkSuggester.
func NewFrameworkSuggester() *FrameworkSuggester {
	return &FrameworkSuggester{
		frameworks: initFrameworks(),
	}
}

func initFrameworks() map[string][]frameworkInfo {
	return map[string][]frameworkInfo{
		"Go": {
			{
				name:        "Standard Library (net/http)",
				purpose:     "web",
				description: "Go's built-in HTTP package. Maximum control, no dependencies.",
				pros:        []string{"No dependencies", "Full control", "Best for learning", "Stable API"},
				cons:        []string{"More boilerplate", "No routing helpers", "Manual middleware"},
				whenToUse:   "Simple APIs, learning Go, or when you want minimal dependencies.",
			},
			{
				name:        "Chi",
				purpose:     "web",
				description: "Lightweight, idiomatic HTTP router with middleware support.",
				pros:        []string{"Lightweight", "Idiomatic Go", "Good middleware", "Compatible with net/http"},
				cons:        []string{"Less features than Gin", "Smaller community"},
				whenToUse:   "When you want a minimal router that feels like Go.",
			},
			{
				name:        "Gin",
				purpose:     "web",
				description: "Fast HTTP web framework with a martini-like API.",
				pros:        []string{"Very fast", "Popular", "Good documentation", "Many examples"},
				cons:        []string{"Not fully idiomatic", "Larger dependency"},
				whenToUse:   "High-performance APIs with good community support.",
			},
			{
				name:        "Echo",
				purpose:     "web",
				description: "High performance, minimalist web framework.",
				pros:        []string{"Fast", "Good documentation", "Built-in middleware", "Easy to use"},
				cons:        []string{"Similar to Gin", "Medium community"},
				whenToUse:   "When you need a balance of features and performance.",
			},
			{
				name:        "Fiber",
				purpose:     "web",
				description: "Express-inspired web framework built on fasthttp.",
				pros:        []string{"Very fast", "Express-like API", "Good for Node.js developers"},
				cons:        []string{"Not net/http compatible", "Different from idiomatic Go"},
				whenToUse:   "Maximum performance or when coming from Express.js.",
			},
			{
				name:        "Cobra",
				purpose:     "cli",
				description: "A library for creating powerful modern CLI applications.",
				pros:        []string{"Industry standard", "Great subcommand support", "Auto-generated help"},
				cons:        []string{"Can be verbose", "Learning curve"},
				whenToUse:   "Any CLI application with subcommands.",
			},
			{
				name:        "GORM",
				purpose:     "orm",
				description: "The fantastic ORM library for Go.",
				pros:        []string{"Feature-rich", "Popular", "Good documentation", "Auto migrations"},
				cons:        []string{"Magic behavior", "Performance overhead", "Complex queries difficult"},
				whenToUse:   "When you want rapid development with an ORM.",
			},
			{
				name:        "sqlx",
				purpose:     "database",
				description: "Extensions to database/sql for easier querying.",
				pros:        []string{"Lightweight", "No magic", "Good performance", "Struct scanning"},
				cons:        []string{"Manual SQL", "No migrations"},
				whenToUse:   "When you prefer writing SQL but want convenience helpers.",
			},
			{
				name:        "Wails",
				purpose:     "crossplatform",
				description: "Build desktop apps using Go and web technologies.",
				pros:        []string{"Go backend power", "Small bundle size", "System webview", "Fast development", "Native look"},
				cons:        []string{"Desktop only", "Smaller community", "Webview differences"},
				whenToUse:   "Desktop apps when you want Go backend with web frontend.",
			},
		},
		"TypeScript": {
			{
				name:        "Express",
				purpose:     "web",
				description: "Fast, unopinionated, minimalist web framework for Node.js.",
				pros:        []string{"Very popular", "Huge ecosystem", "Simple", "Flexible"},
				cons:        []string{"Minimal by default", "Callback-based", "No structure"},
				whenToUse:   "Simple APIs or when you want maximum flexibility.",
			},
			{
				name:        "NestJS",
				purpose:     "web",
				description: "A progressive Node.js framework for building efficient, scalable server-side applications.",
				pros:        []string{"Angular-like structure", "TypeScript-first", "Modular", "Good for large apps"},
				cons:        []string{"Learning curve", "Opinionated", "Heavier"},
				whenToUse:   "Enterprise applications or when you want strong architecture.",
			},
			{
				name:        "Fastify",
				purpose:     "web",
				description: "Fast and low overhead web framework for Node.js.",
				pros:        []string{"Very fast", "Schema-based validation", "Plugin system", "TypeScript support"},
				cons:        []string{"Smaller ecosystem than Express", "Different patterns"},
				whenToUse:   "High-performance APIs with good TypeScript support.",
			},
			{
				name:        "Next.js",
				purpose:     "fullstack",
				description: "The React framework for production.",
				pros:        []string{"SSR/SSG", "Great DX", "Vercel integration", "API routes"},
				cons:        []string{"React lock-in", "Complex caching", "Vercel-optimized"},
				whenToUse:   "Full-stack React applications with SSR needs.",
			},
			{
				name:        "React",
				purpose:     "frontend",
				description: "A JavaScript library for building user interfaces.",
				pros:        []string{"Huge ecosystem", "Component-based", "Large community", "Many jobs"},
				cons:        []string{"Just a library", "Many choices to make", "Frequent changes"},
				whenToUse:   "Most web applications, especially SPAs.",
			},
			{
				name:        "Vue",
				purpose:     "frontend",
				description: "The progressive JavaScript framework.",
				pros:        []string{"Easy to learn", "Good documentation", "Flexible", "Single-file components"},
				cons:        []string{"Smaller ecosystem", "Fewer jobs than React"},
				whenToUse:   "When you want an approachable, well-documented framework.",
			},
			{
				name:        "Prisma",
				purpose:     "orm",
				description: "Next-generation ORM for Node.js and TypeScript.",
				pros:        []string{"Great TypeScript support", "Schema-first", "Migrations", "Studio UI"},
				cons:        []string{"Performance concerns at scale", "Lock-in", "Query limitations"},
				whenToUse:   "TypeScript projects that want type-safe database access.",
			},
			{
				name:        "Electron",
				purpose:     "crossplatform",
				description: "Build cross-platform desktop apps with JavaScript, HTML, and CSS.",
				pros:        []string{"Huge ecosystem", "Web tech familiarity", "Chromium power", "Mature and stable", "Many apps built with it"},
				cons:        []string{"Large bundle size", "High memory usage", "Not native feel", "Security concerns"},
				whenToUse:   "Desktop apps where bundle size isn't critical and you want web tech.",
			},
			{
				name:        "React Native",
				purpose:     "crossplatform",
				description: "Build mobile apps using React and JavaScript that render to native components.",
				pros:        []string{"React ecosystem", "Native components", "Hot reload", "Large community", "Code sharing with web"},
				cons:        []string{"Native bridges can be slow", "Platform-specific code sometimes needed", "Breaking changes"},
				whenToUse:   "Mobile apps when your team knows React and wants native feel.",
			},
		},
		"Python": {
			{
				name:        "FastAPI",
				purpose:     "web",
				description: "Modern, fast web framework for building APIs with Python.",
				pros:        []string{"Very fast", "Auto documentation", "Type hints", "Async support"},
				cons:        []string{"Newer", "Async can be complex"},
				whenToUse:   "Modern Python APIs with automatic OpenAPI docs.",
			},
			{
				name:        "Django",
				purpose:     "web",
				description: "The web framework for perfectionists with deadlines.",
				pros:        []string{"Batteries included", "Admin panel", "ORM", "Mature"},
				cons:        []string{"Monolithic", "Can be slow", "Django way"},
				whenToUse:   "Full-featured web applications with admin needs.",
			},
			{
				name:        "Flask",
				purpose:     "web",
				description: "A lightweight WSGI web application framework.",
				pros:        []string{"Simple", "Flexible", "Good for learning", "Many extensions"},
				cons:        []string{"Manual setup", "No async", "Fewer features"},
				whenToUse:   "Simple APIs or when you want to choose your own components.",
			},
		},
		"Rust": {
			{
				name:        "Actix-web",
				purpose:     "web",
				description: "Powerful, pragmatic, and extremely fast web framework for Rust.",
				pros:        []string{"Extremely fast", "Actor model", "Mature", "Good ecosystem"},
				cons:        []string{"Complex", "Steep learning curve"},
				whenToUse:   "Maximum performance web services in Rust.",
			},
			{
				name:        "Axum",
				purpose:     "web",
				description: "Ergonomic and modular web framework built with Tokio.",
				pros:        []string{"Tokio ecosystem", "Type-safe extractors", "Tower middleware"},
				cons:        []string{"Newer", "Fewer examples"},
				whenToUse:   "When using Tokio ecosystem or want ergonomic API.",
			},
			{
				name:        "Clap",
				purpose:     "cli",
				description: "A full featured, fast Command Line Argument Parser for Rust.",
				pros:        []string{"Derive macros", "Full featured", "Good documentation"},
				cons:        []string{"Compile time cost", "Complex for simple CLIs"},
				whenToUse:   "Any Rust CLI application.",
			},
			{
				name:        "Tauri",
				purpose:     "crossplatform",
				description: "Build smaller, faster, and more secure desktop apps with a web frontend and Rust backend.",
				pros:        []string{"Tiny bundle size", "Memory efficient", "Secure by default", "System webview", "Rust backend power"},
				cons:        []string{"Rust learning curve", "Webview differences across platforms", "Younger ecosystem"},
				whenToUse:   "Desktop apps prioritizing small size and security with web frontend.",
			},
		},
		"Dart": {
			{
				name:        "Flutter",
				purpose:     "crossplatform",
				description: "Google's UI toolkit for building natively compiled applications for mobile, web, and desktop.",
				pros:        []string{"Single codebase for all platforms", "Hot reload", "Beautiful custom UIs", "Growing community", "Great performance"},
				cons:        []string{"Large app size", "Platform-specific still needed sometimes", "Dart learning curve"},
				whenToUse:   "Cross-platform apps targeting mobile, desktop, and web with custom UIs.",
			},
		},
		"Kotlin": {
			{
				name:        "Compose Multiplatform",
				purpose:     "crossplatform",
				description: "JetBrains' declarative UI framework for building native UIs across Android, iOS, desktop, and web.",
				pros:        []string{"Kotlin-first", "Shared UI and logic", "Native performance", "JetBrains support", "Android Compose familiar"},
				cons:        []string{"iOS support still maturing", "Smaller community than Flutter", "JVM overhead on desktop"},
				whenToUse:   "When your team knows Kotlin or comes from Android Compose development.",
			},
		},
		"C#": {
			{
				name:        ".NET MAUI",
				purpose:     "crossplatform",
				description: "Microsoft's cross-platform framework for building native mobile and desktop apps with C# and XAML.",
				pros:        []string{"Official Microsoft support", "Single codebase", "Native UI controls", ".NET ecosystem", "Good Windows integration"},
				cons:        []string{"macOS/Linux support weaker", "Larger app size", "XAML learning curve", "Slower iOS development"},
				whenToUse:   "When targeting Windows primarily with cross-platform needs, or for enterprise .NET shops.",
			},
			{
				name:        "Avalonia",
				purpose:     "crossplatform",
				description: "Cross-platform .NET UI framework inspired by WPF, with consistent UI across all platforms.",
				pros:        []string{"Consistent UI everywhere", "WPF-like XAML", "True cross-platform", "Active community", "Good Linux support"},
				cons:        []string{"Smaller ecosystem than MAUI", "Custom controls needed", "Not Microsoft-official"},
				whenToUse:   "When you want consistent UI across Windows, macOS, and Linux with .NET.",
			},
		},
	}
}

// Suggest returns framework suggestions for the project and language.
func (s *FrameworkSuggester) Suggest(project *domain.Project, language string) []domain.Suggestion {
	var suggestions []domain.Suggestion

	frameworks, ok := s.frameworks[language]
	if !ok {
		return suggestions
	}

	// Filter and score based on project requirements
	for _, fw := range frameworks {
		// Check if this framework purpose matches what we need
		if !s.matchesPurpose(project, fw.purpose) {
			continue
		}

		suggestion := domain.Suggestion{
			Name:          fw.name,
			Category:      fw.purpose,
			Description:   fw.description,
			Pros:          fw.pros,
			Cons:          fw.cons,
			WhenToUse:     fw.whenToUse,
			Confidence:    0.8,
			IsRecommended: true,
		}
		suggestions = append(suggestions, suggestion)
	}

	return suggestions
}

func (s *FrameworkSuggester) matchesPurpose(project *domain.Project, purpose string) bool {
	switch project.AppType {
	case domain.AppTypeWebApp:
		return purpose == "web" || purpose == "frontend"
	case domain.AppTypeAPI:
		return purpose == "web" || purpose == "orm" || purpose == "database"
	case domain.AppTypeCLI:
		return purpose == "cli"
	case domain.AppTypeFullStack:
		return purpose == "fullstack" || purpose == "web" || purpose == "frontend" || purpose == "orm"
	case domain.AppTypeMobile:
		return purpose == "mobile" || purpose == "crossplatform"
	case domain.AppTypeDesktop:
		return purpose == "desktop" || purpose == "crossplatform"
	case domain.AppTypeCrossPlatform:
		return purpose == "crossplatform"
	default:
		return true
	}
}
