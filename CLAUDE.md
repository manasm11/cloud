# CLAUDE.md - Cloud CLI Project

## Project Identity

**Name:** cloud  
**Type:** Command Line Interface (CLI) Tool  
**Language:** Go 1.22+  
**License:** MIT  
**Version:** 0.1.0

## Executive Summary

`cloud` is an interactive CLI tool that helps developers create complex, production-ready applications by systematically gathering requirements through guided questioning, suggesting appropriate technology stacks with educational explanations, and generating comprehensive project specifications that Claude Code can use to build the application.

The tool embodies the philosophy: **"Ask everything, assume nothing, build to last."**

---

## Problem Statement

### The Challenge

Developers face several challenges when starting new projects:

1. **Incomplete Requirements**: Projects often start with vague ideas, leading to scope creep and architectural rewrites
2. **Technology Overwhelm**: The abundance of frameworks, databases, and tools makes informed decisions difficult
3. **Knowledge Gaps**: Developers may not know about technologies that would be perfect for their use case
4. **AI Context Poverty**: AI coding assistants produce better results with comprehensive context, but providing that context is tedious
5. **Inconsistent Practices**: Without structure, projects lack consistent testing, documentation, and deployment strategies
6. **Reinventing the Wheel**: Common project scaffolding is recreated for each new project

### The Solution

`cloud` addresses these challenges by:

- **Systematic Questioning**: Guided interviews extract complete requirements before a single line of code is written
- **Educational Suggestions**: Explains technology options, their trade-offs, and when to use them
- **Zero Assumptions**: Every decision is explicitly confirmed with the user
- **Comprehensive Output**: Generates CLAUDE.md files that give AI assistants complete project context
- **Step-by-Step Breakdown**: Decomposes complex requirements into atomic, testable implementation steps
- **Production-Ready Scaffolding**: Includes Makefile, Docker configurations, and CI/CD templates
- **Specialized Agents**: Identifies areas requiring domain expertise and defines specialized AI agents

---

## Core Principles

### 1. Never Assume
Every piece of information must come from explicit user input. When in doubt, ask. Provide sensible defaults but always confirm them.

### 2. Educate First
Before asking users to choose, explain what each option means, when it's appropriate, and what trade-offs it involves. Users should learn while using the tool.

### 3. Build to Last
Recommend architectures and practices that scale. Prefer boring, proven technologies over trendy ones. Emphasize maintainability over cleverness.

### 4. TDD Always
Every feature in this CLI is implemented using Test-Driven Development. Every generated project includes comprehensive testing strategies.

### 5. Incremental Clarity
Keep questioning until the user explicitly confirms requirements are complete. It's better to ask one more question than to build the wrong thing.

### 6. Complete Context
The generated CLAUDE.md should contain everything needed to build the project. Another developer (or AI) should be able to understand and implement the entire system from this document alone.

---

## Detailed Requirements

### Functional Requirements

#### FR-1: Command Structure

##### FR-1.1: Root Command (`cloud`)
- Display help information when run without subcommands
- Show available commands and global flags
- Display version with `--version` flag

##### FR-1.2: New Project Command (`cloud new`)
- Start interactive interview session
- Support `--output` flag to specify output directory (default: current directory)
- Support `--name` flag to pre-set project name (still confirm with user)
- Support `--minimal` flag for quick setup with fewer questions
- Generate CLAUDE.md and Makefile upon completion
- Generate Docker files if cloud hosting is selected

##### FR-1.3: Edit Command (`cloud edit <path>`)
- Accept path to existing CLAUDE.md file
- Parse and validate the file structure
- Present current configuration organized by sections
- Allow interactive modification of any section
- Support `--section` flag to jump directly to a section
- Create backup before saving changes
- Regenerate dependent files (Makefile, Docker) if needed

##### FR-1.4: Validate Command (`cloud validate <path>`)
- Validate CLAUDE.md file structure and completeness
- Report missing required sections
- Warn about potential issues
- Suggest improvements

##### FR-1.5: Version Command (`cloud version`)
- Display version number
- Display build date
- Display Go version used
- Display git commit hash if available

---

#### FR-2: Interactive Interview System

##### FR-2.1: Project Basics Section
Questions to ask:
1. "What is the name of your project?" (validate: lowercase, alphanumeric, hyphens allowed)
2. "Describe your project in one sentence:"
3. "What problem does this project solve?"
4. "Who are the target users of this application?"
5. "Is this a new project or are you modernizing an existing system?"

##### FR-2.2: Application Type Section
Present options with explanations:

```
What type of application are you building?

1. Web Application
   → Browser-based application with UI
   → Examples: Dashboards, SaaS products, portals
   
2. REST API / Backend Service  
   → HTTP API serving data to clients
   → Examples: Microservices, mobile backends, integrations
   
3. CLI Tool
   → Command-line application
   → Examples: Developer tools, automation scripts, utilities
   
4. Desktop Application
   → Native application for Windows/macOS/Linux
   → Examples: Editors, productivity tools, system utilities
   
5. Mobile Application
   → iOS and/or Android application
   → Examples: Consumer apps, enterprise mobile tools
   
6. Library / Package
   → Reusable code for other developers
   → Examples: SDKs, utility libraries, frameworks
   
7. Full-Stack Application
   → Combined frontend and backend
   → Examples: Complete web applications with API
   
8. Microservices System
   → Multiple coordinated services
   → Examples: Distributed systems, scalable architectures
   
9. Data Pipeline / ETL
   → Data processing and transformation
   → Examples: Analytics pipelines, data integration
   
10. Other (please describe)
```

Follow-up questions based on selection:
- Web App: "Single Page Application (SPA) or Server-Side Rendered (SSR)?"
- API: "Public API, internal service, or both?"
- CLI: "Interactive tool or automation script?"
- Desktop: "Which platforms must be supported?"
- Mobile: "iOS only, Android only, or both? Native or cross-platform?"

##### FR-2.3: Features Deep Dive Section
Implement iterative feature gathering:

```
Let's list all the features your application needs.

I'll keep asking until you say "done". Be as specific as possible.
For each feature, I'll ask clarifying questions.

Feature 1: [user input]
  → Can you describe what this feature does in detail?
  → Who uses this feature? (which user type)
  → Is this feature critical for launch (MVP) or can it come later?
  → Does this feature depend on any other features?
  → What happens when this feature fails? (error handling needs)

Feature 2: [user input]
...

Type "done" when you've listed all features.
Type "review" to see features listed so far.
Type "remove N" to remove feature N.
Type "edit N" to modify feature N.
```

##### FR-2.4: User Roles Section
If applicable (skip for libraries/CLI tools):

```
Let's define the user types/roles in your application.

Common patterns (select all that apply or add custom):
[ ] Anonymous users (not logged in)
[ ] Registered users (standard accounts)
[ ] Premium/Paid users
[ ] Administrators
[ ] Super administrators
[ ] API consumers (other systems)
[ ] Support staff
[ ] Content moderators
[ ] Custom: ___________

For each role, I'll ask:
→ What can this role do that others cannot?
→ What are they restricted from doing?
→ How do users get this role?
```

##### FR-2.5: Data Requirements Section

```
Let's understand your data needs.

1. What are the main entities/objects in your system?
   (Example: Users, Products, Orders, Comments)
   
   For each entity:
   → What are the key attributes?
   → What relationships exist with other entities?
   → How many records do you expect? (ballpark)
   → How often is this data read vs written?
   → Does this data need to be searchable?
   → Is there sensitive/PII data?

2. Data characteristics:
   → Is your data mostly structured, semi-structured, or unstructured?
   → Do you need full-text search capabilities?
   → Do you need geospatial queries?
   → Do you need time-series data handling?
   → Do you need graph relationships?

3. Data lifecycle:
   → How long should data be retained?
   → Do you need audit trails?
   → Do you need soft deletes?
   → Do you need data versioning?
```

##### FR-2.6: Integration Requirements Section

```
What external systems or services will your application integrate with?

Categories to consider:
[ ] Authentication providers (Google, GitHub, SAML, LDAP)
[ ] Payment processors (Stripe, PayPal, etc.)
[ ] Email services (SendGrid, SES, etc.)
[ ] SMS/Notifications (Twilio, Firebase, etc.)
[ ] Cloud storage (S3, GCS, etc.)
[ ] Third-party APIs (specify which)
[ ] Existing internal systems (specify which)
[ ] Analytics services
[ ] Monitoring services
[ ] Other: ___________

For each integration:
→ Is this required for MVP or later phase?
→ Do you have existing accounts/credentials?
→ What's the expected volume of integration calls?
```

##### FR-2.7: Performance Requirements Section

```
Let's understand your performance needs.

These questions help choose the right architecture:

1. Expected concurrent users:
   ( ) < 100 (small internal tool)
   ( ) 100-1,000 (small business)
   ( ) 1,000-10,000 (medium scale)
   ( ) 10,000-100,000 (large scale)
   ( ) > 100,000 (massive scale)
   ( ) Unknown/Unsure - help me decide

2. Acceptable response time for typical operations:
   ( ) < 100ms (real-time critical)
   ( ) < 500ms (interactive)
   ( ) < 2 seconds (standard web)
   ( ) < 10 seconds (background acceptable)

3. Do you have any operations that take a long time (> 30 seconds)?
   → If yes, these need background processing

4. Expected data volume:
   ( ) < 1 GB
   ( ) 1-10 GB
   ( ) 10-100 GB
   ( ) 100 GB - 1 TB
   ( ) > 1 TB

5. Traffic patterns:
   ( ) Steady throughout day
   ( ) Business hours only
   ( ) Spiky/unpredictable
   ( ) Seasonal peaks
```

##### FR-2.8: Security Requirements Section

```
Security is crucial. Let's understand your needs:

1. Authentication requirements:
   [ ] Username/password
   [ ] Email/password
   [ ] Social login (Google, GitHub, etc.)
   [ ] SSO/SAML/OIDC
   [ ] API keys
   [ ] JWT tokens
   [ ] Multi-factor authentication (MFA)
   [ ] Passwordless (magic links)

2. Data sensitivity:
   ( ) Public data only
   ( ) Internal/private but not regulated
   ( ) Contains PII (Personally Identifiable Information)
   ( ) Contains financial data
   ( ) Contains health data (HIPAA)
   ( ) Contains payment card data (PCI-DSS)

3. Security features needed:
   [ ] Rate limiting
   [ ] Brute force protection
   [ ] CORS configuration
   [ ] Input validation/sanitization
   [ ] SQL injection prevention
   [ ] XSS prevention
   [ ] CSRF protection
   [ ] Encryption at rest
   [ ] Encryption in transit
   [ ] Audit logging
   [ ] IP allowlisting

4. Compliance requirements:
   [ ] GDPR
   [ ] CCPA
   [ ] HIPAA
   [ ] SOC 2
   [ ] PCI-DSS
   [ ] Other: ___________
   [ ] None/Not sure
```

##### FR-2.9: Deployment & Hosting Section

```
Where will your application be hosted?

1. Hosting preference:
   ( ) Self-hosted (on-premise servers)
   ( ) Cloud provider - AWS
   ( ) Cloud provider - Google Cloud
   ( ) Cloud provider - Azure
   ( ) Cloud provider - DigitalOcean
   ( ) Platform as a Service (Heroku, Railway, Render, Fly.io)
   ( ) Serverless (Lambda, Cloud Functions)
   ( ) Edge (Cloudflare Workers, Vercel)
   ( ) Not sure - help me decide
   
2. Containerization:
   ( ) Yes, I want Docker
   ( ) No containers
   ( ) Not sure - explain the options
   
   [If unsure, explain]:
   Docker packages your application with its dependencies, making deployment
   consistent across environments. Recommended for:
   - Team projects (works the same on everyone's machine)
   - Complex dependencies
   - Microservices
   - Cloud deployments
   
   Not necessary for:
   - Simple scripts/tools
   - Libraries
   - Serverless functions

3. Orchestration needs (if Docker):
   ( ) Single container (Docker only)
   ( ) Multiple containers, simple setup (docker-compose)
   ( ) Kubernetes
   ( ) Not sure - help me decide

4. Environment stages needed:
   [ ] Local development
   [ ] Development/Integration
   [ ] Staging/QA
   [ ] Production
   [ ] Demo environment
```

##### FR-2.10: Development Practices Section

```
Let's establish development practices:

1. Version control:
   → Repository host: (GitHub / GitLab / Bitbucket / Other)
   → Branching strategy: (Git Flow / GitHub Flow / Trunk-based / Not sure)

2. Testing requirements:
   [ ] Unit tests (required - TDD approach)
   [ ] Integration tests
   [ ] End-to-end tests
   [ ] Performance tests
   [ ] Security tests
   [ ] Contract tests (for APIs)
   
   Minimum test coverage target: ___%

3. Code quality:
   [ ] Linting
   [ ] Formatting
   [ ] Static analysis
   [ ] Pre-commit hooks
   [ ] Code review required

4. CI/CD requirements:
   [ ] Automated testing on PR
   [ ] Automated deployment to staging
   [ ] Automated deployment to production
   [ ] Manual approval for production
   [ ] Rollback capability

5. Documentation needs:
   [ ] API documentation (OpenAPI/Swagger)
   [ ] Code documentation
   [ ] Architecture documentation
   [ ] User documentation
   [ ] Runbooks/Operations documentation
```

---

#### FR-3: Tech Stack Suggestion Engine

##### FR-3.1: Language Suggestions
Based on application type and requirements, suggest appropriate languages:

| Application Type | Primary Suggestions | Alternatives |
|------------------|---------------------|--------------|
| Web Backend | Go, Node.js, Python | Rust, Java, C# |
| Web Frontend | TypeScript/React, Vue | Svelte, Angular |
| CLI Tool | Go, Rust | Python, Node.js |
| Desktop App | Go+Wails, Electron | Tauri, Flutter |
| Mobile App | Flutter, React Native | Swift/Kotlin native |
| Library | Depends on ecosystem | - |
| Data Pipeline | Python, Go | Scala, Rust |

For each suggestion, explain:
- Why it fits the requirements
- Pros and cons
- Community and ecosystem
- Learning curve
- Performance characteristics

##### FR-3.2: Framework Suggestions
Based on language and application type:

**Go Web Frameworks:**
- **Standard library (net/http)**: "Best for learning, simple APIs, maximum control"
- **Chi**: "Lightweight, idiomatic, good middleware ecosystem"
- **Gin**: "Fast, popular, lots of examples available"
- **Echo**: "Feature-rich, good documentation"
- **Fiber**: "Express-like, very fast, good for Node.js developers"

**Go CLI Frameworks:**
- **Cobra**: "Industry standard, excellent subcommand support"
- **urfave/cli**: "Simple, less boilerplate"
- **Bubble Tea**: "For interactive terminal UIs"

##### FR-3.3: Database Suggestions
Based on data requirements:

| Requirement | Suggestion | Explanation |
|-------------|------------|-------------|
| Relational data, ACID needed | PostgreSQL | "Most capable open-source RDBMS" |
| Simple relational, embedded | SQLite | "Zero configuration, great for CLI/desktop" |
| Document storage, flexible schema | MongoDB | "Good for rapid development, flexible" |
| Key-value, caching | Redis | "In-memory, extremely fast" |
| Time-series data | TimescaleDB, InfluxDB | "Optimized for time-based queries" |
| Full-text search | Elasticsearch, Meilisearch | "Purpose-built for search" |
| Graph relationships | Neo4j, DGraph | "When relationships are primary concern" |

##### FR-3.4: Additional Infrastructure
Based on requirements, suggest:

- **Message Queues**: RabbitMQ, Redis Streams, NATS, Kafka
- **Caching**: Redis, Memcached
- **File Storage**: S3, MinIO, local filesystem
- **CDN**: CloudFlare, AWS CloudFront
- **Monitoring**: Prometheus + Grafana, DataDog
- **Logging**: ELK Stack, Loki
- **Tracing**: Jaeger, Zipkin

---

#### FR-4: CLAUDE.md Generation

##### FR-4.1: Document Structure
Generate a comprehensive markdown document with these sections:

```markdown
# CLAUDE.md - [Project Name]

## Project Overview
## Problem Statement
## Target Users
## Core Principles for This Project

## Detailed Requirements
### Functional Requirements
### Non-Functional Requirements
### Constraints and Assumptions

## Tech Stack
### Rationale for Each Choice

## Architecture
### System Architecture Diagram (Mermaid)
### Component Descriptions
### Data Flow

## Data Model
### Entity Descriptions
### Relationships
### Database Schema

## API Design (if applicable)
### Endpoints
### Authentication
### Error Handling

## Implementation Steps
### Phase 1: Foundation
### Phase 2: Core Features
### Phase 3: Additional Features
### Phase 4: Polish and Launch

## Testing Strategy
### Unit Testing Approach
### Integration Testing Approach
### E2E Testing Approach

## Security Considerations
## Deployment Strategy

## Specialized Agents
### [Agent Name]
#### Purpose
#### Context
#### Tasks

## Development Guidelines
### Code Style
### Git Workflow
### Review Process

## Makefile Targets
## Docker Configuration (if applicable)

## Open Questions and Decisions
## Glossary
```

##### FR-4.2: Implementation Steps Generation
Break down requirements into atomic, testable steps:

```markdown
## Implementation Steps

### Phase 1: Foundation (Steps 1-15)

#### Step 1: Project Initialization
**Objective**: Set up project structure and tooling
**Tasks**:
1. Initialize Go module
2. Create directory structure
3. Set up Makefile
4. Configure linting and formatting
5. Set up pre-commit hooks
**Tests**:
- `make lint` passes
- `make build` succeeds
- Directory structure matches specification
**Acceptance Criteria**:
- Clean project compiles
- All tools configured
**Estimated Effort**: 30 minutes

#### Step 2: Domain Models
**Objective**: Define core domain types
**Tasks**:
1. Create [Entity1] struct with fields
2. Create [Entity2] struct with fields
3. Add validation methods
4. Add JSON/DB tags
**Tests**:
- Unit tests for validation logic
- Test struct serialization
**TDD Approach**:
1. Write test for valid [Entity1] creation
2. Implement [Entity1] struct to pass test
3. Write test for invalid cases
4. Implement validation to pass tests
**Acceptance Criteria**:
- All domain models defined
- 100% test coverage on models
**Estimated Effort**: 1 hour

[Continue for all steps...]
```

##### FR-4.3: Specialized Agents Section
When specialized agents are needed, generate:

```markdown
## Specialized Agents

### Database Design Agent

**Purpose**: Design optimal database schema for the application's data requirements.

**Context**:
You are a database design expert. The application requires [describe data needs].
The chosen database is [database]. Consider performance, scalability, and maintainability.

**Your Responsibilities**:
1. Analyze the entity requirements
2. Design normalized schema (3NF minimum)
3. Identify indexes needed for query patterns
4. Design for the expected data volume
5. Consider future migration paths
6. Document all design decisions

**Constraints**:
- Must support [specific requirements]
- Must handle [volume] records
- Must support [query patterns]

**Deliverables**:
1. Entity-Relationship Diagram
2. SQL schema with comments
3. Index strategy document
4. Migration files
5. Seed data for development

---

### API Design Agent

**Purpose**: Design RESTful API following best practices.

**Context**:
You are an API design expert. Design APIs for [application description].
Follow REST best practices, ensure consistency, and optimize for developer experience.

**Your Responsibilities**:
1. Define resource naming conventions
2. Design endpoint structure
3. Define request/response schemas
4. Design error response format
5. Plan API versioning strategy
6. Document authentication flow

**Constraints**:
- RESTful design principles
- JSON:API or similar specification
- Must support [specific requirements]

**Deliverables**:
1. OpenAPI 3.0 specification
2. Example requests/responses
3. Authentication documentation
4. Rate limiting documentation

---

### Security Review Agent

**Purpose**: Review and enhance application security.

**Context**:
You are a security expert. Review [application] for security vulnerabilities
and compliance with [requirements].

**Your Responsibilities**:
1. Review authentication implementation
2. Check authorization logic
3. Identify injection vulnerabilities
4. Review data validation
5. Check encryption usage
6. Verify secure configurations
7. Review dependency security

**Deliverables**:
1. Security assessment report
2. Remediation recommendations
3. Security test cases
4. Secure configuration guide

---

### Testing Agent

**Purpose**: Design comprehensive testing strategy.

**Context**:
You are a testing expert. Design tests for [application] ensuring
high quality and reliability.

**Your Responsibilities**:
1. Design unit test structure
2. Identify integration test scenarios
3. Design E2E test flows
4. Create test data strategy
5. Define coverage requirements
6. Set up test automation

**Deliverables**:
1. Test plan document
2. Test case specifications
3. Mock/fixture designs
4. CI test configuration
```

---

#### FR-5: Makefile Generation

##### FR-5.1: Standard Targets
Generate Makefile with targets based on tech stack:

```makefile
# Makefile for [Project Name]
# Generated by cloud CLI

.PHONY: help build run test lint fmt clean docker-build docker-run

# Default target
.DEFAULT_GOAL := help

# Variables
APP_NAME := [project-name]
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

## build: Build the application
build:
	go build $(LDFLAGS) -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

## run: Run the application
run:
	go run ./cmd/$(APP_NAME)

## test: Run all tests
test:
	go test -v -race -cover ./...

## test-coverage: Run tests with coverage report
test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

## lint: Run linter
lint:
	golangci-lint run ./...

## fmt: Format code
fmt:
	go fmt ./...
	goimports -w .

## vet: Run go vet
vet:
	go vet ./...

## clean: Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

## deps: Download dependencies
deps:
	go mod download
	go mod tidy

## install-tools: Install development tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest

# Docker targets (if applicable)
## docker-build: Build Docker image
docker-build:
	docker build -t $(APP_NAME):$(VERSION) .

## docker-run: Run Docker container
docker-run:
	docker run --rm -it $(APP_NAME):$(VERSION)

## docker-compose-up: Start all services
docker-compose-up:
	docker-compose up -d

## docker-compose-down: Stop all services
docker-compose-down:
	docker-compose down

## docker-compose-logs: View logs
docker-compose-logs:
	docker-compose logs -f

# Database targets (if applicable)
## db-migrate: Run database migrations
db-migrate:
	go run ./cmd/migrate up

## db-rollback: Rollback last migration
db-rollback:
	go run ./cmd/migrate down

## db-seed: Seed database with test data
db-seed:
	go run ./cmd/seed

# CI targets
## ci: Run all CI checks
ci: lint vet test

## pre-commit: Run pre-commit checks
pre-commit: fmt lint vet test
```

---

#### FR-6: Docker File Generation

##### FR-6.1: Dockerfile Generation
When Docker is selected, generate optimized Dockerfile:

```dockerfile
# Dockerfile for [Project Name]
# Generated by cloud CLI

# Build stage
FROM golang:1.22-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build arguments
ARG VERSION=dev
ARG BUILD_TIME=unknown

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-s -w -X main.version=${VERSION} -X main.buildTime=${BUILD_TIME}" \
    -o app ./cmd/[project-name]

# Final stage
FROM alpine:3.19

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/app .

# Copy any required files (configs, migrations, etc.)
# COPY --from=builder /build/migrations ./migrations

# Change ownership
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port (adjust as needed)
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
ENTRYPOINT ["./app"]
```

##### FR-6.2: docker-compose.yml Generation
When multiple services or databases are needed:

```yaml
# docker-compose.yml for [Project Name]
# Generated by cloud CLI

version: '3.8'

services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
      args:
        VERSION: ${VERSION:-dev}
        BUILD_TIME: ${BUILD_TIME:-unknown}
    container_name: [project-name]-app
    ports:
      - "${APP_PORT:-8080}:8080"
    environment:
      - DATABASE_URL=postgres://postgres:postgres@db:5432/[project-name]?sslmode=disable
      - REDIS_URL=redis://redis:6379
      - ENV=development
    depends_on:
      db:
        condition: service_healthy
      redis:
        condition: service_healthy
    networks:
      - [project-name]-network
    restart: unless-stopped

  db:
    image: postgres:16-alpine
    container_name: [project-name]-db
    ports:
      - "${DB_PORT:-5432}:5432"
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: [project-name]
    volumes:
      - postgres-data:/var/lib/postgresql/data
      - ./scripts/init-db.sql:/docker-entrypoint-initdb.d/init.sql:ro
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 5s
      timeout: 5s
      retries: 5
    networks:
      - [project-name]-network
    restart: unless-stopped

  redis:
    image: redis:7-alpine
    container_name: [project-name]-redis
    ports:
      - "${REDIS_PORT:-6379}:6379"
    volumes:
      - redis-data:/data
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 5s
      timeout: 5s
      retries: 5
    networks:
      - [project-name]-network
    restart: unless-stopped

  # Optional: Database admin UI
  adminer:
    image: adminer
    container_name: [project-name]-adminer
    ports:
      - "${ADMINER_PORT:-8081}:8080"
    depends_on:
      - db
    networks:
      - [project-name]-network
    profiles:
      - dev-tools

volumes:
  postgres-data:
  redis-data:

networks:
  [project-name]-network:
    driver: bridge
```

##### FR-6.3: .dockerignore Generation

```
# .dockerignore for [Project Name]
# Generated by cloud CLI

# Git
.git
.gitignore

# IDE
.idea/
.vscode/
*.swp
*.swo

# Build artifacts
bin/
dist/
*.exe

# Test artifacts
coverage.out
coverage.html
*.test

# Documentation
*.md
!README.md
docs/

# Docker
Dockerfile*
docker-compose*
.docker/

# CI/CD
.github/
.gitlab-ci.yml
.travis.yml

# Misc
.env.local
.env.*.local
*.log
tmp/
```

---

### Non-Functional Requirements

#### NFR-1: Usability
- NFR-1.1: Prompts must be clear and unambiguous
- NFR-1.2: Errors must include actionable guidance
- NFR-1.3: Support terminal colors (with graceful fallback)
- NFR-1.4: Progress indication for operations > 1 second
- NFR-1.5: Allow interruption (Ctrl+C) with graceful handling
- NFR-1.6: Remember user's place if interrupted

#### NFR-2: Performance
- NFR-2.1: Cold start under 100ms
- NFR-2.2: No perceptible lag between prompts
- NFR-2.3: File generation under 2 seconds

#### NFR-3: Reliability
- NFR-3.1: Never lose user input
- NFR-3.2: Validate all input before proceeding
- NFR-3.3: Create backups before modifying files
- NFR-3.4: Generate valid markdown always

#### NFR-4: Portability
- NFR-4.1: Support Linux, macOS, Windows
- NFR-4.2: Single binary distribution
- NFR-4.3: No external runtime dependencies

#### NFR-5: Extensibility
- NFR-5.1: Plugin architecture for custom questions
- NFR-5.2: Template customization support
- NFR-5.3: Easy addition of new tech stack options

---

## Tech Stack

### Language: Go 1.22+

**Rationale:**
- Fast compilation and execution
- Single binary output simplifies distribution
- Excellent CLI tooling ecosystem
- Strong type system catches errors at compile time
- Cross-compilation built-in
- Aligns with your Go expertise and tutorial work

### Dependencies

| Dependency | Purpose | Rationale |
|------------|---------|-----------|
| `github.com/spf13/cobra` | CLI framework | Industry standard, excellent docs |
| `github.com/charmbracelet/huh` | Interactive forms | Modern, accessible, beautiful |
| `github.com/charmbracelet/lipgloss` | Terminal styling | Consistent theming |
| `github.com/charmbracelet/log` | Logging | Matches UI style |
| `gopkg.in/yaml.v3` | YAML parsing | For config files |
| `github.com/stretchr/testify` | Testing | Better assertions |

### Project Structure

```
cloud/
├── cmd/
│   └── cloud/
│       └── main.go                     # Entry point
├── internal/
│   ├── app/
│   │   ├── app.go                      # Application orchestrator
│   │   └── app_test.go
│   ├── command/
│   │   ├── root.go                     # Root command
│   │   ├── root_test.go
│   │   ├── new.go                      # new command
│   │   ├── new_test.go
│   │   ├── edit.go                     # edit command
│   │   ├── edit_test.go
│   │   ├── validate.go                 # validate command
│   │   ├── validate_test.go
│   │   ├── version.go                  # version command
│   │   └── version_test.go
│   ├── interview/
│   │   ├── interviewer.go              # Interview orchestrator
│   │   ├── interviewer_test.go
│   │   ├── session.go                  # Session state management
│   │   ├── session_test.go
│   │   ├── questions/
│   │   │   ├── basics.go               # Basic project questions
│   │   │   ├── basics_test.go
│   │   │   ├── apptype.go              # App type questions
│   │   │   ├── apptype_test.go
│   │   │   ├── features.go             # Feature gathering
│   │   │   ├── features_test.go
│   │   │   ├── data.go                 # Data requirements
│   │   │   ├── data_test.go
│   │   │   ├── integration.go          # Integration requirements
│   │   │   ├── integration_test.go
│   │   │   ├── performance.go          # Performance requirements
│   │   │   ├── performance_test.go
│   │   │   ├── security.go             # Security requirements
│   │   │   ├── security_test.go
│   │   │   ├── deployment.go           # Deployment requirements
│   │   │   ├── deployment_test.go
│   │   │   └── practices.go            # Dev practices
│   │   │   └── practices_test.go
│   │   └── validation/
│   │       ├── validator.go            # Input validation
│   │       └── validator_test.go
│   ├── suggester/
│   │   ├── suggester.go                # Suggestion engine interface
│   │   ├── language.go                 # Language suggestions
│   │   ├── language_test.go
│   │   ├── framework.go                # Framework suggestions
│   │   ├── framework_test.go
│   │   ├── database.go                 # Database suggestions
│   │   ├── database_test.go
│   │   └── infrastructure.go           # Infrastructure suggestions
│   │   └── infrastructure_test.go
│   ├── generator/
│   │   ├── generator.go                # Generator interface
│   │   ├── claude_md.go                # CLAUDE.md generator
│   │   ├── claude_md_test.go
│   │   ├── makefile.go                 # Makefile generator
│   │   ├── makefile_test.go
│   │   ├── dockerfile.go               # Dockerfile generator
│   │   ├── dockerfile_test.go
│   │   ├── compose.go                  # docker-compose generator
│   │   ├── compose_test.go
│   │   └── templates/
│   │       ├── embed.go                # Embed templates
│   │       ├── claude.md.tmpl
│   │       ├── makefile.tmpl
│   │       ├── dockerfile.tmpl
│   │       └── compose.tmpl
│   ├── parser/
│   │   ├── claude_md.go                # CLAUDE.md parser
│   │   └── claude_md_test.go
│   └── domain/
│       ├── project.go                  # Project model
│       ├── project_test.go
│       ├── requirement.go              # Requirement model
│       ├── techstack.go                # Tech stack model
│       ├── agent.go                    # Specialized agent model
│       └── step.go                     # Implementation step model
├── templates/                          # User-customizable templates
├── go.mod
├── go.sum
├── Makefile
├── README.md
├── CHANGELOG.md
└── .goreleaser.yml                     # Release configuration
```

---

## Implementation Steps

### Phase 1: Foundation (Steps 1-8)

#### Step 1: Project Initialization
**Objective**: Set up Go module and basic structure
**TDD Approach**:
1. Create go.mod with `go mod init github.com/manas/cloud`
2. Create directory structure
3. Set up Makefile with basic targets
4. Configure .golangci.yml

**Files to create**:
- go.mod
- Makefile
- .golangci.yml
- .gitignore

**Acceptance Criteria**:
- `go build ./...` succeeds
- `make lint` passes
- Directory structure matches specification

---

#### Step 2: Domain Models
**Objective**: Define core domain types with validation

**TDD Approach**:
```go
// Write test first
func TestProject_Validate(t *testing.T) {
    tests := []struct {
        name    string
        project Project
        wantErr bool
    }{
        {
            name: "valid project",
            project: Project{
                Name:        "my-app",
                Description: "A test application",
            },
            wantErr: false,
        },
        {
            name: "empty name",
            project: Project{
                Name: "",
            },
            wantErr: true,
        },
        {
            name: "invalid name with spaces",
            project: Project{
                Name: "my app",
            },
            wantErr: true,
        },
    }
    // ... run tests
}
```

**Files to create**:
- internal/domain/project.go
- internal/domain/project_test.go
- internal/domain/requirement.go
- internal/domain/techstack.go
- internal/domain/agent.go
- internal/domain/step.go

**Acceptance Criteria**:
- All domain types defined
- Validation methods implemented
- 100% test coverage on domain package

---

#### Step 3: Root Command Setup
**Objective**: Implement root command with Cobra

**TDD Approach**:
```go
func TestRootCommand(t *testing.T) {
    cmd := NewRootCommand()
    
    t.Run("has correct use", func(t *testing.T) {
        assert.Equal(t, "cloud", cmd.Use)
    })
    
    t.Run("has help flag", func(t *testing.T) {
        assert.NotNil(t, cmd.Flags().Lookup("help"))
    })
    
    t.Run("has version flag", func(t *testing.T) {
        assert.NotNil(t, cmd.Flags().Lookup("version"))
    })
}
```

**Files to create**:
- cmd/cloud/main.go
- internal/command/root.go
- internal/command/root_test.go

**Acceptance Criteria**:
- `cloud --help` shows help
- `cloud --version` shows version
- Commands can be registered

---

#### Step 4: Version Command
**Objective**: Implement version command

**TDD Approach**:
```go
func TestVersionCommand(t *testing.T) {
    var buf bytes.Buffer
    cmd := NewVersionCommand()
    cmd.SetOut(&buf)
    
    err := cmd.Execute()
    
    assert.NoError(t, err)
    assert.Contains(t, buf.String(), "cloud version")
}
```

**Files to create**:
- internal/command/version.go
- internal/command/version_test.go

**Acceptance Criteria**:
- Shows version number
- Shows build info
- Clean output format

---

#### Step 5: Input Validation Module
**Objective**: Create reusable input validators

**TDD Approach**:
```go
func TestValidateProjectName(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        wantErr bool
    }{
        {"valid simple", "myapp", false},
        {"valid with hyphen", "my-app", false},
        {"valid with numbers", "app123", false},
        {"invalid spaces", "my app", true},
        {"invalid uppercase", "MyApp", true},
        {"invalid special chars", "my@app", true},
        {"too short", "ab", true},
        {"too long", strings.Repeat("a", 65), true},
    }
    // ... run tests
}
```

**Files to create**:
- internal/interview/validation/validator.go
- internal/interview/validation/validator_test.go

**Acceptance Criteria**:
- Project name validation works
- Clear error messages
- Reusable for other validations

---

#### Step 6: Session State Management
**Objective**: Manage interview state with save/restore

**TDD Approach**:
```go
func TestSession_SaveAndRestore(t *testing.T) {
    session := NewSession()
    session.SetProjectName("test-project")
    session.SetAppType(AppTypeWebApp)
    
    // Save to temp file
    path := filepath.Join(t.TempDir(), "session.json")
    err := session.Save(path)
    assert.NoError(t, err)
    
    // Restore
    restored, err := LoadSession(path)
    assert.NoError(t, err)
    assert.Equal(t, "test-project", restored.ProjectName())
    assert.Equal(t, AppTypeWebApp, restored.AppType())
}
```

**Files to create**:
- internal/interview/session.go
- internal/interview/session_test.go

**Acceptance Criteria**:
- State can be saved on interrupt
- State can be restored
- Partial sessions handled gracefully

---

#### Step 7: Basic Interview Flow
**Objective**: Implement basic project questions

**TDD Approach**:
```go
func TestBasicsQuestions(t *testing.T) {
    // Mock the prompt interface
    mockPrompt := &MockPrompter{
        responses: map[string]string{
            "project_name": "test-app",
            "description":  "A test application",
            "problem":      "Solves testing problems",
        },
    }
    
    questions := NewBasicsQuestions(mockPrompt)
    result, err := questions.Ask()
    
    assert.NoError(t, err)
    assert.Equal(t, "test-app", result.Name)
    assert.Equal(t, "A test application", result.Description)
}
```

**Files to create**:
- internal/interview/questions/basics.go
- internal/interview/questions/basics_test.go

**Acceptance Criteria**:
- Asks for project name
- Validates input
- Asks for description
- Asks about problem being solved

---

#### Step 8: Application Type Questions
**Objective**: Implement app type selection with explanations

**TDD Approach**:
```go
func TestAppTypeQuestions(t *testing.T) {
    mockPrompt := &MockPrompter{
        selections: map[string]int{
            "app_type": 0, // Web Application
        },
    }
    
    questions := NewAppTypeQuestions(mockPrompt)
    result, err := questions.Ask()
    
    assert.NoError(t, err)
    assert.Equal(t, AppTypeWebApp, result.Type)
}

func TestAppTypeDescriptions(t *testing.T) {
    descriptions := GetAppTypeDescriptions()
    
    assert.Contains(t, descriptions[AppTypeWebApp], "Browser-based")
    assert.Contains(t, descriptions[AppTypeCLI], "Command-line")
    // Verify each type has description
}
```

**Files to create**:
- internal/interview/questions/apptype.go
- internal/interview/questions/apptype_test.go

**Acceptance Criteria**:
- Shows all app types with descriptions
- Asks follow-up questions based on selection
- Clear explanations for each option

---

### Phase 2: Deep Dive Questions (Steps 9-16)

#### Step 9: Feature Gathering
**Objective**: Implement iterative feature collection

**Files to create**:
- internal/interview/questions/features.go
- internal/interview/questions/features_test.go

**Key behaviors**:
- Keep asking until user says "done"
- Allow review, edit, remove
- Ask clarifying questions per feature
- Track MVP vs later features

---

#### Step 10: User Roles Questions
**Objective**: Gather user types and permissions

**Files to create**:
- internal/interview/questions/users.go
- internal/interview/questions/users_test.go

---

#### Step 11: Data Requirements Questions
**Objective**: Understand data needs

**Files to create**:
- internal/interview/questions/data.go
- internal/interview/questions/data_test.go

---

#### Step 12: Integration Requirements
**Objective**: Gather external system needs

**Files to create**:
- internal/interview/questions/integration.go
- internal/interview/questions/integration_test.go

---

#### Step 13: Performance Requirements
**Objective**: Understand scale and performance needs

**Files to create**:
- internal/interview/questions/performance.go
- internal/interview/questions/performance_test.go

---

#### Step 14: Security Requirements
**Objective**: Gather security and compliance needs

**Files to create**:
- internal/interview/questions/security.go
- internal/interview/questions/security_test.go

---

#### Step 15: Deployment Requirements
**Objective**: Understand hosting and deployment needs

**Files to create**:
- internal/interview/questions/deployment.go
- internal/interview/questions/deployment_test.go

---

#### Step 16: Development Practices
**Objective**: Gather dev workflow preferences

**Files to create**:
- internal/interview/questions/practices.go
- internal/interview/questions/practices_test.go

---

### Phase 3: Suggestion Engine (Steps 17-21)

#### Step 17: Language Suggester
**Objective**: Suggest languages based on requirements

**Files to create**:
- internal/suggester/suggester.go (interface)
- internal/suggester/language.go
- internal/suggester/language_test.go

---

#### Step 18: Framework Suggester
**Objective**: Suggest frameworks based on language

**Files to create**:
- internal/suggester/framework.go
- internal/suggester/framework_test.go

---

#### Step 19: Database Suggester
**Objective**: Suggest databases based on data needs

**Files to create**:
- internal/suggester/database.go
- internal/suggester/database_test.go

---

#### Step 20: Infrastructure Suggester
**Objective**: Suggest caching, queues, etc.

**Files to create**:
- internal/suggester/infrastructure.go
- internal/suggester/infrastructure_test.go

---

#### Step 21: Tech Stack Orchestrator
**Objective**: Coordinate all suggestions

**Files to create**:
- internal/suggester/orchestrator.go
- internal/suggester/orchestrator_test.go

---

### Phase 4: Generators (Steps 22-27)

#### Step 22: Template System
**Objective**: Set up embedded templates

**Files to create**:
- internal/generator/templates/embed.go
- internal/generator/templates/claude.md.tmpl
- internal/generator/templates/makefile.tmpl
- internal/generator/templates/dockerfile.tmpl
- internal/generator/templates/compose.tmpl

---

#### Step 23: CLAUDE.md Generator
**Objective**: Generate comprehensive markdown

**Files to create**:
- internal/generator/generator.go (interface)
- internal/generator/claude_md.go
- internal/generator/claude_md_test.go

---

#### Step 24: Makefile Generator
**Objective**: Generate project Makefile

**Files to create**:
- internal/generator/makefile.go
- internal/generator/makefile_test.go

---

#### Step 25: Dockerfile Generator
**Objective**: Generate optimized Dockerfile

**Files to create**:
- internal/generator/dockerfile.go
- internal/generator/dockerfile_test.go

---

#### Step 26: docker-compose Generator
**Objective**: Generate compose file

**Files to create**:
- internal/generator/compose.go
- internal/generator/compose_test.go

---

#### Step 27: Agent Definition Generator
**Objective**: Generate specialized agent specs

**Files to create**:
- internal/generator/agents.go
- internal/generator/agents_test.go

---

### Phase 5: Commands (Steps 28-31)

#### Step 28: New Command
**Objective**: Implement `cloud new`

**Files to create**:
- internal/command/new.go
- internal/command/new_test.go

---

#### Step 29: CLAUDE.md Parser
**Objective**: Parse existing CLAUDE.md files

**Files to create**:
- internal/parser/claude_md.go
- internal/parser/claude_md_test.go

---

#### Step 30: Edit Command
**Objective**: Implement `cloud edit`

**Files to create**:
- internal/command/edit.go
- internal/command/edit_test.go

---

#### Step 31: Validate Command
**Objective**: Implement `cloud validate`

**Files to create**:
- internal/command/validate.go
- internal/command/validate_test.go

---

### Phase 6: Polish (Steps 32-35)

#### Step 32: Error Handling
**Objective**: Consistent error presentation

**Files to create**:
- internal/errors/errors.go
- internal/errors/errors_test.go

---

#### Step 33: Progress Indicators
**Objective**: Add spinners and progress bars

**Files to modify**:
- Add progress to generation steps

---

#### Step 34: Interrupt Handling
**Objective**: Graceful Ctrl+C handling

**Files to modify**:
- internal/app/app.go
- Signal handling throughout

---

#### Step 35: Documentation
**Objective**: Complete documentation

**Files to create**:
- README.md (comprehensive)
- CONTRIBUTING.md
- docs/ folder with guides

---

## Testing Strategy

### Unit Tests
- Every package has corresponding `_test.go` files
- Use table-driven tests
- Mock external dependencies
- Target 80%+ coverage

### Integration Tests
- Test command execution end-to-end
- Use temporary directories
- Verify file generation

### Test Helpers

```go
// testutil/mock_prompter.go
type MockPrompter struct {
    TextResponses   map[string]string
    SelectResponses map[string]int
    MultiResponses  map[string][]int
    ConfirmResponses map[string]bool
}

func (m *MockPrompter) Text(key, prompt string) (string, error) {
    if resp, ok := m.TextResponses[key]; ok {
        return resp, nil
    }
    return "", fmt.Errorf("no mock response for %s", key)
}
```

---

## Specialized Agents

### Database Design Agent

**When to invoke**: Project has complex data requirements (>5 entities, specific query patterns, or scaling needs)

**Agent Definition**:
```markdown
You are a database design expert. Design the optimal schema for this application.

Context:
- Database: [chosen database]
- Entities: [list from requirements]
- Relationships: [from data questions]
- Expected volume: [from performance questions]
- Query patterns: [from feature descriptions]

Your tasks:
1. Design normalized schema
2. Define indexes for query patterns
3. Plan for the expected scale
4. Create migration strategy
5. Design seed data structure

Output:
- Entity-relationship diagram (Mermaid)
- SQL schema with comments
- Index recommendations
- Migration files structure
```

### API Design Agent

**When to invoke**: Project is API-centric or has >10 API endpoints

**Agent Definition**:
```markdown
You are an API design expert. Design RESTful APIs for this application.

Context:
- Features: [list from requirements]
- User roles: [from user questions]
- Authentication: [from security questions]

Your tasks:
1. Define resource structure
2. Design endpoint naming
3. Specify request/response formats
4. Define error response format
5. Plan versioning strategy

Output:
- OpenAPI 3.0 specification
- Example requests for each endpoint
- Authentication flow documentation
```

### Security Review Agent

**When to invoke**: Application handles sensitive data or has compliance requirements

### Testing Agent

**When to invoke**: Complex testing requirements or >50 implementation steps

### DevOps Agent

**When to invoke**: Kubernetes deployment or multi-environment setup needed

---

## Development Guidelines

### Code Style
- Follow Go conventions
- Use `gofmt` and `goimports`
- Lint with `golangci-lint`
- Document exported functions

### Git Workflow
- Feature branches from `main`
- Conventional commits
- PR required for merge
- All tests must pass

### Testing Requirements
- Write test first (TDD)
- No PR without tests
- Coverage must not decrease

---

## Makefile Targets

| Target | Description |
|--------|-------------|
| `make build` | Build the binary |
| `make test` | Run all tests |
| `make test-coverage` | Run tests with coverage |
| `make lint` | Run linter |
| `make fmt` | Format code |
| `make install` | Install to GOPATH/bin |
| `make clean` | Remove build artifacts |
| `make ci` | Run all CI checks |

---

## Open Questions

These should be resolved during development:

1. Should we support a config file for default preferences?
2. Should templates be customizable by users?
3. Should we support plugins for custom question sets?
4. Should we integrate with Claude API directly for smart suggestions?
5. Should we generate CI/CD configs (GitHub Actions, GitLab CI)?

---

## Glossary

| Term | Definition |
|------|------------|
| CLAUDE.md | Markdown file containing complete project specification for Claude Code |
| Interview | The interactive session gathering project requirements |
| Session | The state of an interview, can be saved and restored |
| Suggester | Component that recommends technology choices |
| Generator | Component that creates output files |
| Specialized Agent | AI agent with domain-specific expertise |

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 0.1.0 | TBD | Initial release |

---

*This CLAUDE.md was generated by cloud CLI v0.1.0*