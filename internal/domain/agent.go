package domain

// Agent represents a specialized AI agent for a specific domain task.
type Agent struct {
	Name             string   `json:"name" yaml:"name"`
	Purpose          string   `json:"purpose" yaml:"purpose"`
	Context          string   `json:"context" yaml:"context"`
	Responsibilities []string `json:"responsibilities" yaml:"responsibilities"`
	Constraints      []string `json:"constraints" yaml:"constraints"`
	Deliverables     []string `json:"deliverables" yaml:"deliverables"`
}

// AgentType represents the type of specialized agent.
type AgentType string

const (
	AgentTypeDatabaseDesign AgentType = "database-design"
	AgentTypeAPIDesign      AgentType = "api-design"
	AgentTypeSecurityReview AgentType = "security-review"
	AgentTypeTesting        AgentType = "testing"
	AgentTypeDevOps         AgentType = "devops"
	AgentTypeFrontend       AgentType = "frontend"
	AgentTypePerformance    AgentType = "performance"
)

// NewDatabaseDesignAgent creates a database design agent based on project requirements.
func NewDatabaseDesignAgent(p *Project) Agent {
	return Agent{
		Name:    "Database Design Agent",
		Purpose: "Design optimal database schema for the application's data requirements.",
		Context: buildDatabaseAgentContext(p),
		Responsibilities: []string{
			"Analyze the entity requirements",
			"Design normalized schema (3NF minimum)",
			"Identify indexes needed for query patterns",
			"Design for the expected data volume",
			"Consider future migration paths",
			"Document all design decisions",
		},
		Constraints: buildDatabaseAgentConstraints(p),
		Deliverables: []string{
			"Entity-Relationship Diagram",
			"SQL schema with comments",
			"Index strategy document",
			"Migration files",
			"Seed data for development",
		},
	}
}

// NewAPIDesignAgent creates an API design agent based on project requirements.
func NewAPIDesignAgent(p *Project) Agent {
	return Agent{
		Name:    "API Design Agent",
		Purpose: "Design RESTful API following best practices.",
		Context: buildAPIAgentContext(p),
		Responsibilities: []string{
			"Define resource naming conventions",
			"Design endpoint structure",
			"Define request/response schemas",
			"Design error response format",
			"Plan API versioning strategy",
			"Document authentication flow",
		},
		Constraints: []string{
			"RESTful design principles",
			"JSON:API or similar specification",
		},
		Deliverables: []string{
			"OpenAPI 3.0 specification",
			"Example requests/responses",
			"Authentication documentation",
			"Rate limiting documentation",
		},
	}
}

// NewSecurityReviewAgent creates a security review agent based on project requirements.
func NewSecurityReviewAgent(p *Project) Agent {
	return Agent{
		Name:    "Security Review Agent",
		Purpose: "Review and enhance application security.",
		Context: buildSecurityAgentContext(p),
		Responsibilities: []string{
			"Review authentication implementation",
			"Check authorization logic",
			"Identify injection vulnerabilities",
			"Review data validation",
			"Check encryption usage",
			"Verify secure configurations",
			"Review dependency security",
		},
		Constraints: []string{},
		Deliverables: []string{
			"Security assessment report",
			"Remediation recommendations",
			"Security test cases",
			"Secure configuration guide",
		},
	}
}

// NewTestingAgent creates a testing agent based on project requirements.
func NewTestingAgent(p *Project) Agent {
	return Agent{
		Name:    "Testing Agent",
		Purpose: "Design comprehensive testing strategy.",
		Context: buildTestingAgentContext(p),
		Responsibilities: []string{
			"Design unit test structure",
			"Identify integration test scenarios",
			"Design E2E test flows",
			"Create test data strategy",
			"Define coverage requirements",
			"Set up test automation",
		},
		Constraints: []string{},
		Deliverables: []string{
			"Test plan document",
			"Test case specifications",
			"Mock/fixture designs",
			"CI test configuration",
		},
	}
}

// NewDevOpsAgent creates a DevOps agent based on project requirements.
func NewDevOpsAgent(p *Project) Agent {
	return Agent{
		Name:    "DevOps Agent",
		Purpose: "Design deployment and operations infrastructure.",
		Context: buildDevOpsAgentContext(p),
		Responsibilities: []string{
			"Design deployment pipeline",
			"Configure container orchestration",
			"Set up monitoring and alerting",
			"Design backup and recovery",
			"Create runbooks",
			"Implement infrastructure as code",
		},
		Constraints: []string{},
		Deliverables: []string{
			"CI/CD pipeline configuration",
			"Kubernetes manifests or Docker Compose",
			"Monitoring dashboards",
			"Runbook documentation",
			"Infrastructure as code",
		},
	}
}

// Helper functions to build agent contexts

func buildDatabaseAgentContext(p *Project) string {
	return "You are a database design expert. The application requires handling of " +
		"multiple entities with various relationships. Consider performance, scalability, and maintainability."
}

func buildDatabaseAgentConstraints(p *Project) []string {
	constraints := []string{}
	if p.Performance.DataVolume != "" {
		constraints = append(constraints, "Must handle "+p.Performance.DataVolume+" of data")
	}
	if p.Performance.ConcurrentUsers != "" {
		constraints = append(constraints, "Must support "+p.Performance.ConcurrentUsers+" concurrent users")
	}
	return constraints
}

func buildAPIAgentContext(p *Project) string {
	return "You are an API design expert. Design APIs for the application following REST best practices. " +
		"Ensure consistency and optimize for developer experience."
}

func buildSecurityAgentContext(p *Project) string {
	context := "You are a security expert. Review the application for security vulnerabilities"
	if len(p.Security.Compliance) > 0 {
		context += " and compliance requirements."
	} else {
		context += "."
	}
	return context
}

func buildTestingAgentContext(p *Project) string {
	return "You are a testing expert. Design tests for the application ensuring high quality and reliability."
}

func buildDevOpsAgentContext(p *Project) string {
	context := "You are a DevOps expert. Design deployment infrastructure for "
	if p.Deployment.HostingType != "" {
		context += p.Deployment.HostingType + " hosting."
	} else {
		context += "the application."
	}
	return context
}
