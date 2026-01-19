package questions

import (
	"github.com/manas/cloud/internal/domain"
)

// AppTypeResult contains the results from the app type questions.
type AppTypeResult struct {
	AppType domain.AppType
	Details map[string]string
}

// AppTypeQuestions handles the application type interview questions.
type AppTypeQuestions struct {
	prompter Prompter
}

// NewAppTypeQuestions creates a new AppTypeQuestions instance.
func NewAppTypeQuestions(prompter Prompter) *AppTypeQuestions {
	return &AppTypeQuestions{
		prompter: prompter,
	}
}

// GetAppTypeOptions returns all application type options with descriptions.
func GetAppTypeOptions() []Option {
	return []Option{
		{
			Key:         "web",
			Title:       "Web Application",
			Description: "Browser-based application with UI (dashboards, SaaS, portals)",
		},
		{
			Key:         "api",
			Title:       "REST API / Backend Service",
			Description: "HTTP API serving data to clients (microservices, mobile backends)",
		},
		{
			Key:         "cli",
			Title:       "CLI Tool",
			Description: "Command-line application (developer tools, automation scripts)",
		},
		{
			Key:         "desktop",
			Title:       "Desktop Application",
			Description: "Native application for Windows/macOS/Linux",
		},
		{
			Key:         "mobile",
			Title:       "Mobile Application",
			Description: "iOS and/or Android application",
		},
		{
			Key:         "crossplatform",
			Title:       "Cross-Platform Application",
			Description: "Single codebase for desktop, mobile, and/or web (Flutter, Electron, Tauri)",
		},
		{
			Key:         "library",
			Title:       "Library / Package",
			Description: "Reusable code for other developers (SDKs, utilities)",
		},
		{
			Key:         "fullstack",
			Title:       "Full-Stack Application",
			Description: "Combined frontend and backend",
		},
		{
			Key:         "microservices",
			Title:       "Microservices System",
			Description: "Multiple coordinated services (distributed systems)",
		},
		{
			Key:         "pipeline",
			Title:       "Data Pipeline / ETL",
			Description: "Data processing and transformation",
		},
		{
			Key:         "other",
			Title:       "Other",
			Description: "Custom application type",
		},
	}
}

// Ask runs the app type questions interview.
func (q *AppTypeQuestions) Ask() (*AppTypeResult, error) {
	result := &AppTypeResult{
		Details: make(map[string]string),
	}

	// Main app type selection
	appTypeKey, err := q.prompter.Select(
		"What type of application are you building?",
		"Select the option that best describes your project",
		GetAppTypeOptions(),
	)
	if err != nil {
		return nil, err
	}

	// Map key to domain type
	result.AppType = keyToAppType(appTypeKey)

	// Ask follow-up questions based on type
	if err := q.askFollowUp(appTypeKey, result); err != nil {
		return nil, err
	}

	return result, nil
}

func keyToAppType(key string) domain.AppType {
	mapping := map[string]domain.AppType{
		"web":           domain.AppTypeWebApp,
		"api":           domain.AppTypeAPI,
		"cli":           domain.AppTypeCLI,
		"desktop":       domain.AppTypeDesktop,
		"mobile":        domain.AppTypeMobile,
		"crossplatform": domain.AppTypeCrossPlatform,
		"library":       domain.AppTypeLibrary,
		"fullstack":     domain.AppTypeFullStack,
		"microservices": domain.AppTypeMicroservices,
		"pipeline":      domain.AppTypeDataPipeline,
		"other":         domain.AppTypeOther,
	}
	if t, ok := mapping[key]; ok {
		return t
	}
	return domain.AppTypeOther
}

func (q *AppTypeQuestions) askFollowUp(appType string, result *AppTypeResult) error {
	switch appType {
	case "web":
		return q.askWebFollowUp(result)
	case "api":
		return q.askAPIFollowUp(result)
	case "cli":
		return q.askCLIFollowUp(result)
	case "mobile":
		return q.askMobileFollowUp(result)
	case "desktop":
		return q.askDesktopFollowUp(result)
	case "crossplatform":
		return q.askCrossPlatformFollowUp(result)
	case "fullstack":
		return q.askFullStackFollowUp(result)
	}
	return nil
}

func (q *AppTypeQuestions) askWebFollowUp(result *AppTypeResult) error {
	webType, err := q.prompter.Select(
		"Web Application Type",
		"How will the web application be rendered?",
		[]Option{
			{Key: "spa", Title: "Single Page Application (SPA)", Description: "Client-side rendering with JavaScript framework"},
			{Key: "ssr", Title: "Server-Side Rendered (SSR)", Description: "HTML generated on the server"},
			{Key: "hybrid", Title: "Hybrid (SSR + SPA)", Description: "Initial SSR with client-side hydration"},
			{Key: "static", Title: "Static Site", Description: "Pre-built HTML pages"},
		},
	)
	if err != nil {
		return err
	}
	result.Details["web_type"] = webType
	return nil
}

func (q *AppTypeQuestions) askAPIFollowUp(result *AppTypeResult) error {
	apiType, err := q.prompter.Select(
		"API Type",
		"Who will consume this API?",
		[]Option{
			{Key: "public", Title: "Public API", Description: "External developers and third parties"},
			{Key: "internal", Title: "Internal Service", Description: "Only internal systems and services"},
			{Key: "both", Title: "Both Public and Internal", Description: "Serves both internal and external clients"},
		},
	)
	if err != nil {
		return err
	}
	result.Details["api_type"] = apiType
	return nil
}

func (q *AppTypeQuestions) askCLIFollowUp(result *AppTypeResult) error {
	cliType, err := q.prompter.Select(
		"CLI Type",
		"What kind of CLI tool is this?",
		[]Option{
			{Key: "interactive", Title: "Interactive Tool", Description: "User interacts with prompts and menus"},
			{Key: "automation", Title: "Automation Script", Description: "Runs with flags/args, minimal interaction"},
			{Key: "both", Title: "Both", Description: "Supports both interactive and non-interactive modes"},
		},
	)
	if err != nil {
		return err
	}
	result.Details["cli_type"] = cliType
	return nil
}

func (q *AppTypeQuestions) askMobileFollowUp(result *AppTypeResult) error {
	platform, err := q.prompter.Select(
		"Mobile Platform",
		"Which platforms will you support?",
		[]Option{
			{Key: "ios", Title: "iOS Only", Description: "Apple devices only"},
			{Key: "android", Title: "Android Only", Description: "Android devices only"},
			{Key: "both", Title: "Both iOS and Android", Description: "Cross-platform support"},
		},
	)
	if err != nil {
		return err
	}
	result.Details["platform"] = platform

	approach, err := q.prompter.Select(
		"Mobile Development Approach",
		"How will you build the mobile app?",
		[]Option{
			{Key: "native", Title: "Native", Description: "Separate native codebases (Swift/Kotlin)"},
			{Key: "cross-platform", Title: "Cross-Platform", Description: "Single codebase (Flutter, React Native)"},
			{Key: "hybrid", Title: "Hybrid/WebView", Description: "Web technologies in a native wrapper"},
		},
	)
	if err != nil {
		return err
	}
	result.Details["approach"] = approach

	return nil
}

func (q *AppTypeQuestions) askDesktopFollowUp(result *AppTypeResult) error {
	platforms, err := q.prompter.MultiSelect(
		"Desktop Platforms",
		"Which desktop platforms must be supported?",
		[]Option{
			{Key: "windows", Title: "Windows", Description: "Microsoft Windows"},
			{Key: "macos", Title: "macOS", Description: "Apple macOS"},
			{Key: "linux", Title: "Linux", Description: "Linux distributions"},
		},
	)
	if err != nil {
		return err
	}
	for i, p := range platforms {
		result.Details["platform_"+string(rune('0'+i))] = p
	}
	return nil
}

func (q *AppTypeQuestions) askCrossPlatformFollowUp(result *AppTypeResult) error {
	// Ask which platforms to target
	platforms, err := q.prompter.MultiSelect(
		"Target Platforms",
		"Which platforms will your app run on?",
		[]Option{
			{Key: "windows", Title: "Windows", Description: "Microsoft Windows desktop"},
			{Key: "macos", Title: "macOS", Description: "Apple macOS desktop"},
			{Key: "linux", Title: "Linux", Description: "Linux desktop"},
			{Key: "ios", Title: "iOS", Description: "iPhone and iPad"},
			{Key: "android", Title: "Android", Description: "Android phones and tablets"},
			{Key: "web", Title: "Web", Description: "Browser-based access"},
		},
	)
	if err != nil {
		return err
	}
	result.Details["platforms"] = joinPlatforms(platforms)

	// Ask which framework to use
	framework, err := q.prompter.Select(
		"Cross-Platform Framework",
		"Which framework will you use? (We'll suggest based on your needs)",
		[]Option{
			{Key: "flutter", Title: "Flutter", Description: "Google's UI toolkit. Best for mobile + desktop. Dart language."},
			{Key: "react-native", Title: "React Native", Description: "Meta's framework. Best for mobile. JavaScript/TypeScript."},
			{Key: "electron", Title: "Electron", Description: "Desktop apps with web tech. JavaScript/TypeScript. Large bundles."},
			{Key: "tauri", Title: "Tauri", Description: "Lightweight desktop apps. Rust backend, web frontend. Small bundles."},
			{Key: "maui", Title: ".NET MAUI", Description: "Microsoft's cross-platform. C#. Good Windows integration."},
			{Key: "avalonia", Title: "Avalonia", Description: "Cross-platform .NET UI. C#. WPF-like."},
			{Key: "wails", Title: "Wails", Description: "Go + web frontend. Lightweight desktop apps."},
			{Key: "compose", Title: "Compose Multiplatform", Description: "JetBrains. Kotlin. Desktop + Android + iOS."},
			{Key: "undecided", Title: "Help me decide", Description: "Get recommendations based on requirements"},
		},
	)
	if err != nil {
		return err
	}
	result.Details["framework"] = framework

	// Ask about UI requirements
	uiStyle, err := q.prompter.Select(
		"UI Style",
		"What kind of UI does your app need?",
		[]Option{
			{Key: "native", Title: "Native Look & Feel", Description: "Matches each platform's design language"},
			{Key: "custom", Title: "Custom/Branded UI", Description: "Consistent look across all platforms"},
			{Key: "hybrid", Title: "Hybrid", Description: "Custom UI with native elements where needed"},
		},
	)
	if err != nil {
		return err
	}
	result.Details["ui_style"] = uiStyle

	return nil
}

func joinPlatforms(platforms []string) string {
	if len(platforms) == 0 {
		return ""
	}
	result := platforms[0]
	for i := 1; i < len(platforms); i++ {
		result += "," + platforms[i]
	}
	return result
}

func (q *AppTypeQuestions) askFullStackFollowUp(result *AppTypeResult) error {
	// Ask about frontend
	if err := q.askWebFollowUp(result); err != nil {
		return err
	}

	// Ask about API type
	apiType, err := q.prompter.Select(
		"Backend API Style",
		"How will the frontend communicate with the backend?",
		[]Option{
			{Key: "rest", Title: "REST API", Description: "Traditional RESTful endpoints"},
			{Key: "graphql", Title: "GraphQL", Description: "Single endpoint with query language"},
			{Key: "trpc", Title: "tRPC / RPC", Description: "Type-safe RPC calls"},
		},
	)
	if err != nil {
		return err
	}
	result.Details["api_style"] = apiType

	return nil
}
