package domain

// TechStack represents the selected technology stack for a project.
type TechStack struct {
	// Primary language
	Language        string `json:"language" yaml:"language"`
	LanguageVersion string `json:"language_version" yaml:"language_version"`

	// Frameworks
	Frameworks []Framework `json:"frameworks" yaml:"frameworks"`

	// Database
	Database        string `json:"database" yaml:"database"`
	DatabaseVersion string `json:"database_version" yaml:"database_version"`

	// Additional infrastructure
	Cache        string `json:"cache" yaml:"cache"`
	MessageQueue string `json:"message_queue" yaml:"message_queue"`
	FileStorage  string `json:"file_storage" yaml:"file_storage"`
	SearchEngine string `json:"search_engine" yaml:"search_engine"`

	// Observability
	Monitoring string `json:"monitoring" yaml:"monitoring"`
	Logging    string `json:"logging" yaml:"logging"`
	Tracing    string `json:"tracing" yaml:"tracing"`

	// Development tools
	Linter        string `json:"linter" yaml:"linter"`
	Formatter     string `json:"formatter" yaml:"formatter"`
	TestFramework string `json:"test_framework" yaml:"test_framework"`
}

// Framework represents a framework selection with rationale.
type Framework struct {
	Name      string `json:"name" yaml:"name"`
	Purpose   string `json:"purpose" yaml:"purpose"` // web, cli, orm, etc.
	Version   string `json:"version" yaml:"version"`
	Rationale string `json:"rationale" yaml:"rationale"`
}

// Suggestion represents a technology suggestion with explanation.
type Suggestion struct {
	Name          string   `json:"name" yaml:"name"`
	Category      string   `json:"category" yaml:"category"` // language, framework, database, etc.
	Description   string   `json:"description" yaml:"description"`
	Pros          []string `json:"pros" yaml:"pros"`
	Cons          []string `json:"cons" yaml:"cons"`
	WhenToUse     string   `json:"when_to_use" yaml:"when_to_use"`
	Confidence    float64  `json:"confidence" yaml:"confidence"` // 0-1, how confident we are this is a good fit
	IsRecommended bool     `json:"is_recommended" yaml:"is_recommended"`
}

// LanguageSuggestions contains all language options for a project type.
type LanguageSuggestions struct {
	Primary      []Suggestion `json:"primary" yaml:"primary"`
	Alternatives []Suggestion `json:"alternatives" yaml:"alternatives"`
}

// FrameworkSuggestions contains framework options for a language.
type FrameworkSuggestions struct {
	Language string       `json:"language" yaml:"language"`
	Options  []Suggestion `json:"options" yaml:"options"`
}

// DatabaseSuggestions contains database options based on requirements.
type DatabaseSuggestions struct {
	Options []Suggestion `json:"options" yaml:"options"`
}
