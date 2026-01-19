package generator

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/manas/cloud/internal/domain"
)

// ClaudeMDGenerator generates CLAUDE.md files.
type ClaudeMDGenerator struct{}

// NewClaudeMDGenerator creates a new ClaudeMDGenerator.
func NewClaudeMDGenerator() *ClaudeMDGenerator {
	return &ClaudeMDGenerator{}
}

// Generate creates a CLAUDE.md file for the project.
func (g *ClaudeMDGenerator) Generate(project *domain.Project, outputDir string) error {
	content, err := g.render(project)
	if err != nil {
		return fmt.Errorf("failed to render CLAUDE.md: %w", err)
	}

	return WriteFile(outputDir, "CLAUDE.md", content)
}

// Render generates the CLAUDE.md content without writing to disk.
func (g *ClaudeMDGenerator) Render(project *domain.Project) (string, error) {
	return g.render(project)
}

func (g *ClaudeMDGenerator) render(project *domain.Project) (string, error) {
	tmpl, err := template.New("claude.md").Funcs(template.FuncMap{
		"join":      strings.Join,
		"upper":     strings.ToUpper,
		"lower":     strings.ToLower,
		"title":     strings.Title,
		"hasPrefix": strings.HasPrefix,
		"indent": func(spaces int, s string) string {
			pad := strings.Repeat(" ", spaces)
			lines := strings.Split(s, "\n")
			for i, line := range lines {
				if line != "" {
					lines[i] = pad + line
				}
			}
			return strings.Join(lines, "\n")
		},
	}).Parse(claudeMDTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	data := templateData{
		Project:     project,
		GeneratedAt: time.Now().Format("2006-01-02"),
		Version:     "0.1.0",
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

type templateData struct {
	Project     *domain.Project
	GeneratedAt string
	Version     string
}

const claudeMDTemplate = `# CLAUDE.md - {{.Project.Name}}

## Project Overview

**Name:** {{.Project.Name}}
**Type:** {{.Project.AppType.String}}
**Generated:** {{.GeneratedAt}}

### Description

{{.Project.Description}}

### Problem Statement

{{if .Project.Problem}}{{.Project.Problem}}{{else}}(To be defined){{end}}

### Target Users

{{if .Project.TargetUsers}}{{.Project.TargetUsers}}{{else}}(To be defined){{end}}

---

## Core Principles

1. **Never Assume** - Every decision is explicit and confirmed
2. **Test First** - TDD approach for all features
3. **Build to Last** - Prefer proven, maintainable solutions
4. **Security First** - Security is not an afterthought

---

## Requirements

### Functional Requirements
{{if .Project.Features}}
{{range $i, $f := .Project.Features}}
#### FR-{{$i | printf "%d"}}: {{$f.Name}}
{{if $f.Description}}{{$f.Description}}{{end}}
{{if $f.IsMVP}}- **Priority:** MVP{{else}}- **Priority:** Post-MVP{{end}}
{{end}}
{{else}}
(Features to be defined during implementation)
{{end}}

### Non-Functional Requirements

#### Performance
- **Concurrent Users:** {{if .Project.Performance.ConcurrentUsers}}{{.Project.Performance.ConcurrentUsers}}{{else}}TBD{{end}}
- **Response Time:** {{if .Project.Performance.ResponseTime}}{{.Project.Performance.ResponseTime}}{{else}}TBD{{end}}

#### Security
- **Data Sensitivity:** {{if .Project.Security.DataSensitivity}}{{.Project.Security.DataSensitivity}}{{else}}TBD{{end}}
{{if .Project.Security.AuthMethods}}- **Authentication:** {{join .Project.Security.AuthMethods ", "}}{{end}}

---

## Tech Stack

### Language
{{if .Project.TechStack.Language}}**{{.Project.TechStack.Language}}**{{else}}(To be selected){{end}}

### Frameworks
{{if .Project.TechStack.Frameworks}}
{{range .Project.TechStack.Frameworks}}
- **{{.Name}}** ({{.Purpose}}): {{.Rationale}}
{{end}}
{{else}}
(To be selected based on requirements)
{{end}}

### Database
{{if .Project.TechStack.Database}}**{{.Project.TechStack.Database}}**{{else}}(To be selected){{end}}

### Infrastructure
{{if .Project.TechStack.Cache}}- **Cache:** {{.Project.TechStack.Cache}}{{end}}
{{if .Project.TechStack.MessageQueue}}- **Message Queue:** {{.Project.TechStack.MessageQueue}}{{end}}
{{if .Project.TechStack.Monitoring}}- **Monitoring:** {{.Project.TechStack.Monitoring}}{{end}}

---

## Data Model

### Entities
{{if .Project.Entities}}
{{range .Project.Entities}}
#### {{.Name}}
{{if .Attributes}}
| Attribute | Type | Required | Unique |
|-----------|------|----------|--------|
{{range .Attributes}}| {{.Name}} | {{.Type}} | {{if .Required}}Yes{{else}}No{{end}} | {{if .Unique}}Yes{{else}}No{{end}} |
{{end}}
{{end}}
{{end}}
{{else}}
(Entities to be defined during implementation)
{{end}}

### Data Characteristics
- **Structure:** {{if .Project.DataCharacteristics.Structure}}{{.Project.DataCharacteristics.Structure}}{{else}}TBD{{end}}
- **Full-Text Search:** {{if .Project.DataCharacteristics.FullTextSearch}}Yes{{else}}No{{end}}
- **Audit Trail:** {{if .Project.DataCharacteristics.AuditTrail}}Yes{{else}}No{{end}}

---

## User Roles
{{if .Project.UserRoles}}
{{range .Project.UserRoles}}
### {{.Name}}
{{if .Description}}{{.Description}}{{end}}
{{if .Permissions}}
**Permissions:**
{{range .Permissions}}- {{.}}
{{end}}{{end}}
{{end}}
{{else}}
(User roles to be defined)
{{end}}

---

## Integrations
{{if .Project.Integrations}}
| Integration | Type | MVP | Provider |
|-------------|------|-----|----------|
{{range .Project.Integrations}}| {{.Name}} | {{.Type}} | {{if .IsMVP}}Yes{{else}}No{{end}} | {{.Provider}} |
{{end}}
{{else}}
(No external integrations defined)
{{end}}

---

## Deployment

### Hosting
- **Provider:** {{if .Project.Deployment.CloudProvider}}{{.Project.Deployment.CloudProvider}}{{else}}TBD{{end}}
- **Docker:** {{if .Project.Deployment.UseDocker}}Yes{{else}}No{{end}}
{{if .Project.Deployment.Orchestration}}- **Orchestration:** {{.Project.Deployment.Orchestration}}{{end}}

### Environments
{{if .Project.Deployment.Environments}}
{{range .Project.Deployment.Environments}}- {{.}}
{{end}}
{{else}}
- Development
- Staging
- Production
{{end}}

---

## Development Practices

### Version Control
- **Repository:** {{if .Project.DevPractices.RepoHost}}{{.Project.DevPractices.RepoHost}}{{else}}GitHub{{end}}
- **Branching:** {{if .Project.DevPractices.BranchStrategy}}{{.Project.DevPractices.BranchStrategy}}{{else}}GitHub Flow{{end}}

### Testing
{{if .Project.DevPractices.TestingTypes}}
{{range .Project.DevPractices.TestingTypes}}- {{.}}
{{end}}
{{else}}
- Unit Tests
- Integration Tests
{{end}}

### Code Quality
- Linting enabled
- Formatting enforced
- Pre-commit hooks recommended

---

## Implementation Steps

### Phase 1: Foundation
1. Project setup and tooling configuration
2. Domain models and validation
3. Database schema and migrations
4. Basic API structure

### Phase 2: Core Features
{{if .Project.Features}}
{{range $i, $f := .Project.Features}}
{{if $f.IsMVP}}{{$i | printf "%d"}}. {{$f.Name}}
{{end}}{{end}}
{{else}}
(Features to be defined)
{{end}}

### Phase 3: Polish
1. Error handling and logging
2. Documentation
3. Performance optimization
4. Security hardening

---

## Makefile Targets

| Target | Description |
|--------|-------------|
| ` + "`make build`" + ` | Build the application |
| ` + "`make test`" + ` | Run all tests |
| ` + "`make lint`" + ` | Run linter |
| ` + "`make fmt`" + ` | Format code |
| ` + "`make clean`" + ` | Clean build artifacts |

---

## Open Questions

- [ ] Finalize database schema
- [ ] Define API endpoints
- [ ] Security review requirements
- [ ] Deployment pipeline details

---

*Generated by cloud CLI v{{.Version}} on {{.GeneratedAt}}*
`
