// Package domain contains the core domain models for the cloud CLI.
package domain

import (
	"errors"
	"regexp"
	"strings"
)

// AppType represents the type of application being built.
type AppType int

const (
	AppTypeUnknown AppType = iota
	AppTypeWebApp
	AppTypeAPI
	AppTypeCLI
	AppTypeDesktop
	AppTypeMobile
	AppTypeCrossPlatform // Desktop + Mobile from single codebase
	AppTypeLibrary
	AppTypeFullStack
	AppTypeMicroservices
	AppTypeDataPipeline
	AppTypeOther
)

// String returns the human-readable name of the app type.
func (a AppType) String() string {
	names := map[AppType]string{
		AppTypeWebApp:        "Web Application",
		AppTypeAPI:           "REST API / Backend Service",
		AppTypeCLI:           "CLI Tool",
		AppTypeDesktop:       "Desktop Application",
		AppTypeMobile:        "Mobile Application",
		AppTypeCrossPlatform: "Cross-Platform Application",
		AppTypeLibrary:       "Library / Package",
		AppTypeFullStack:     "Full-Stack Application",
		AppTypeMicroservices: "Microservices System",
		AppTypeDataPipeline:  "Data Pipeline / ETL",
		AppTypeOther:         "Other",
	}
	if name, ok := names[a]; ok {
		return name
	}
	return "Unknown"
}

// Description returns a detailed description of the app type.
func (a AppType) Description() string {
	descriptions := map[AppType]string{
		AppTypeWebApp:        "Browser-based application with UI. Examples: Dashboards, SaaS products, portals.",
		AppTypeAPI:           "HTTP API serving data to clients. Examples: Microservices, mobile backends, integrations.",
		AppTypeCLI:           "Command-line application. Examples: Developer tools, automation scripts, utilities.",
		AppTypeDesktop:       "Native application for Windows/macOS/Linux. Examples: Editors, productivity tools.",
		AppTypeMobile:        "iOS and/or Android application. Examples: Consumer apps, enterprise mobile tools.",
		AppTypeCrossPlatform: "Single codebase for multiple platforms (desktop, mobile, web). Examples: Flutter apps, Electron apps, Tauri apps.",
		AppTypeLibrary:       "Reusable code for other developers. Examples: SDKs, utility libraries, frameworks.",
		AppTypeFullStack:     "Combined frontend and backend. Examples: Complete web applications with API.",
		AppTypeMicroservices: "Multiple coordinated services. Examples: Distributed systems, scalable architectures.",
		AppTypeDataPipeline:  "Data processing and transformation. Examples: Analytics pipelines, data integration.",
		AppTypeOther:         "Custom application type not listed above.",
	}
	if desc, ok := descriptions[a]; ok {
		return desc
	}
	return "Unknown application type."
}

// Project represents a software project being defined.
type Project struct {
	// Basic information
	Name         string `json:"name" yaml:"name"`
	Description  string `json:"description" yaml:"description"`
	Problem      string `json:"problem" yaml:"problem"`
	TargetUsers  string `json:"target_users" yaml:"target_users"`
	IsNewProject bool   `json:"is_new_project" yaml:"is_new_project"`

	// Application type
	AppType        AppType           `json:"app_type" yaml:"app_type"`
	AppTypeDetails map[string]string `json:"app_type_details" yaml:"app_type_details"`

	// Features
	Features []Feature `json:"features" yaml:"features"`

	// User roles
	UserRoles []UserRole `json:"user_roles" yaml:"user_roles"`

	// Data requirements
	Entities            []Entity            `json:"entities" yaml:"entities"`
	DataCharacteristics DataCharacteristics `json:"data_characteristics" yaml:"data_characteristics"`

	// Integrations
	Integrations []Integration `json:"integrations" yaml:"integrations"`

	// Performance requirements
	Performance PerformanceRequirements `json:"performance" yaml:"performance"`

	// Security requirements
	Security SecurityRequirements `json:"security" yaml:"security"`

	// Deployment
	Deployment DeploymentConfig `json:"deployment" yaml:"deployment"`

	// Development practices
	DevPractices DevPractices `json:"dev_practices" yaml:"dev_practices"`

	// Tech stack (after suggestions are confirmed)
	TechStack TechStack `json:"tech_stack" yaml:"tech_stack"`

	// Generated specialized agents
	Agents []Agent `json:"agents" yaml:"agents"`

	// Implementation steps
	Steps []Step `json:"steps" yaml:"steps"`
}

// NewProject creates a new Project with the given name and description.
func NewProject(name, description string) *Project {
	return &Project{
		Name:           name,
		Description:    description,
		IsNewProject:   true,
		Features:       make([]Feature, 0),
		UserRoles:      make([]UserRole, 0),
		Entities:       make([]Entity, 0),
		Integrations:   make([]Integration, 0),
		Agents:         make([]Agent, 0),
		Steps:          make([]Step, 0),
		AppTypeDetails: make(map[string]string),
	}
}

// projectNameRegex validates project names (lowercase letters, numbers, hyphens).
var projectNameRegex = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// Validate validates the project configuration.
func (p *Project) Validate() error {
	// Validate name
	if p.Name == "" {
		return errors.New("project name is required")
	}

	if len(p.Name) < 3 {
		return errors.New("project name must be at least 3 characters")
	}

	if len(p.Name) > 64 {
		return errors.New("project name must be at most 64 characters")
	}

	if strings.HasPrefix(p.Name, "-") || strings.HasSuffix(p.Name, "-") {
		return errors.New("project name cannot start or end with a hyphen")
	}

	if !projectNameRegex.MatchString(p.Name) {
		return errors.New("project name can only contain lowercase letters, numbers, and hyphens")
	}

	// Validate description
	if p.Description == "" {
		return errors.New("project description is required")
	}

	return nil
}

// Feature represents a feature of the application.
type Feature struct {
	Name          string   `json:"name" yaml:"name"`
	Description   string   `json:"description" yaml:"description"`
	UserRoles     []string `json:"user_roles" yaml:"user_roles"`
	IsMVP         bool     `json:"is_mvp" yaml:"is_mvp"`
	DependsOn     []string `json:"depends_on" yaml:"depends_on"`
	ErrorHandling string   `json:"error_handling" yaml:"error_handling"`
}

// UserRole represents a user role in the application.
type UserRole struct {
	Name         string   `json:"name" yaml:"name"`
	Description  string   `json:"description" yaml:"description"`
	Permissions  []string `json:"permissions" yaml:"permissions"`
	Restrictions []string `json:"restrictions" yaml:"restrictions"`
	HowAssigned  string   `json:"how_assigned" yaml:"how_assigned"`
}

// Entity represents a data entity in the application.
type Entity struct {
	Name           string         `json:"name" yaml:"name"`
	Attributes     []Attribute    `json:"attributes" yaml:"attributes"`
	Relationships  []Relationship `json:"relationships" yaml:"relationships"`
	ExpectedCount  string         `json:"expected_count" yaml:"expected_count"`
	ReadWriteRatio string         `json:"read_write_ratio" yaml:"read_write_ratio"`
	Searchable     bool           `json:"searchable" yaml:"searchable"`
	HasPII         bool           `json:"has_pii" yaml:"has_pii"`
}

// Attribute represents an attribute of an entity.
type Attribute struct {
	Name     string `json:"name" yaml:"name"`
	Type     string `json:"type" yaml:"type"`
	Required bool   `json:"required" yaml:"required"`
	Unique   bool   `json:"unique" yaml:"unique"`
}

// Relationship represents a relationship between entities.
type Relationship struct {
	TargetEntity string `json:"target_entity" yaml:"target_entity"`
	Type         string `json:"type" yaml:"type"` // one-to-one, one-to-many, many-to-many
	Required     bool   `json:"required" yaml:"required"`
}

// DataCharacteristics describes the characteristics of the data.
type DataCharacteristics struct {
	Structure       string `json:"structure" yaml:"structure"` // structured, semi-structured, unstructured
	FullTextSearch  bool   `json:"full_text_search" yaml:"full_text_search"`
	Geospatial      bool   `json:"geospatial" yaml:"geospatial"`
	TimeSeries      bool   `json:"time_series" yaml:"time_series"`
	GraphRelations  bool   `json:"graph_relations" yaml:"graph_relations"`
	RetentionPeriod string `json:"retention_period" yaml:"retention_period"`
	AuditTrail      bool   `json:"audit_trail" yaml:"audit_trail"`
	SoftDeletes     bool   `json:"soft_deletes" yaml:"soft_deletes"`
	Versioning      bool   `json:"versioning" yaml:"versioning"`
}

// Integration represents an external system integration.
type Integration struct {
	Name           string `json:"name" yaml:"name"`
	Type           string `json:"type" yaml:"type"` // auth, payment, email, sms, storage, api, etc.
	Provider       string `json:"provider" yaml:"provider"`
	IsMVP          bool   `json:"is_mvp" yaml:"is_mvp"`
	HasCredentials bool   `json:"has_credentials" yaml:"has_credentials"`
	ExpectedVolume string `json:"expected_volume" yaml:"expected_volume"`
}

// PerformanceRequirements describes performance needs.
type PerformanceRequirements struct {
	ConcurrentUsers string `json:"concurrent_users" yaml:"concurrent_users"`
	ResponseTime    string `json:"response_time" yaml:"response_time"`
	LongOperations  bool   `json:"long_operations" yaml:"long_operations"`
	DataVolume      string `json:"data_volume" yaml:"data_volume"`
	TrafficPattern  string `json:"traffic_pattern" yaml:"traffic_pattern"`
}

// SecurityRequirements describes security needs.
type SecurityRequirements struct {
	AuthMethods      []string `json:"auth_methods" yaml:"auth_methods"`
	DataSensitivity  string   `json:"data_sensitivity" yaml:"data_sensitivity"`
	SecurityFeatures []string `json:"security_features" yaml:"security_features"`
	Compliance       []string `json:"compliance" yaml:"compliance"`
}

// DeploymentConfig describes deployment requirements.
type DeploymentConfig struct {
	HostingType   string   `json:"hosting_type" yaml:"hosting_type"`
	CloudProvider string   `json:"cloud_provider" yaml:"cloud_provider"`
	UseDocker     bool     `json:"use_docker" yaml:"use_docker"`
	Orchestration string   `json:"orchestration" yaml:"orchestration"`
	Environments  []string `json:"environments" yaml:"environments"`
}

// DevPractices describes development practices.
type DevPractices struct {
	RepoHost       string   `json:"repo_host" yaml:"repo_host"`
	BranchStrategy string   `json:"branch_strategy" yaml:"branch_strategy"`
	TestingTypes   []string `json:"testing_types" yaml:"testing_types"`
	MinCoverage    int      `json:"min_coverage" yaml:"min_coverage"`
	CodeQuality    []string `json:"code_quality" yaml:"code_quality"`
	CICDFeatures   []string `json:"cicd_features" yaml:"cicd_features"`
	Documentation  []string `json:"documentation" yaml:"documentation"`
}
