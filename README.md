# Cloud CLI

An interactive CLI tool that helps developers create comprehensive project specifications through guided questioning. It suggests appropriate technology stacks with educational explanations and generates `CLAUDE.md` files that Claude Code can use to build applications.

**Philosophy**: *"Ask everything, assume nothing, build to last."*

## Features

- **Interactive Interview System**: Guided questions to gather complete project requirements
- **Technology Suggestions**: Recommends languages, frameworks, databases with explanations
- **Cross-Platform Support**: Build for desktop, mobile, and web from a single codebase
- **File Generation**: Automatically generates CLAUDE.md, Makefile, Dockerfile, and docker-compose.yml
- **Session Persistence**: Resume interrupted interviews where you left off

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/manasm11/cloud.git
cd cloud

# Build the binary
make build

# Or install directly
go install ./cmd/cloud
```

### Prerequisites

- Go 1.22 or later

## Usage

### Create a New Project Specification

```bash
# Start interactive interview
cloud new

# Specify output directory
cloud new --output ./my-project

# Pre-set project name
cloud new --name my-awesome-app

# Resume a previous session
cloud new --resume
```

### Commands

| Command | Description |
|---------|-------------|
| `cloud new` | Start a new project specification interview |
| `cloud version` | Display version information |
| `cloud --help` | Show help information |

### Interview Sections

The `cloud new` command guides you through:

1. **Project Basics** - Name, description, problem statement, target users
2. **Application Type** - Web, API, CLI, Desktop, Mobile, Cross-Platform, etc.
3. **Features** - Detailed feature requirements with dependencies
4. **User Roles** - User types and permissions
5. **Data Requirements** - Entities, relationships, data characteristics
6. **Integrations** - External services and APIs
7. **Performance** - Scale and response time requirements
8. **Security** - Authentication, data sensitivity, compliance
9. **Deployment** - Hosting, Docker, environments
10. **Development Practices** - Testing, CI/CD, code quality

### Supported Application Types

| Type | Description |
|------|-------------|
| Web Application | Browser-based apps with UI |
| REST API | HTTP APIs serving data to clients |
| CLI Tool | Command-line applications |
| Desktop Application | Native apps for Windows/macOS/Linux |
| Mobile Application | iOS and/or Android apps |
| Cross-Platform | Single codebase for multiple platforms |
| Library/Package | Reusable code for other developers |
| Full-Stack | Combined frontend and backend |
| Microservices | Multiple coordinated services |
| Data Pipeline | Data processing and transformation |

### Cross-Platform Frameworks

When building cross-platform applications, `cloud` supports:

| Framework | Language | Platforms |
|-----------|----------|-----------|
| Flutter | Dart | iOS, Android, Web, Windows, macOS, Linux |
| React Native | TypeScript | iOS, Android |
| Electron | TypeScript | Windows, macOS, Linux |
| Tauri | Rust + TypeScript | Windows, macOS, Linux |
| Wails | Go + TypeScript | Windows, macOS, Linux |
| .NET MAUI | C# | iOS, Android, Windows, macOS |
| Avalonia | C# | Windows, macOS, Linux |
| Compose Multiplatform | Kotlin | Android, iOS, Desktop, Web |

## Generated Files

After completing the interview, `cloud` generates:

### CLAUDE.md

A comprehensive project specification containing:
- Project overview and requirements
- Technology stack with rationale
- Architecture diagrams (Mermaid)
- Data models and API design
- Implementation steps
- Testing strategy
- Security considerations

### Makefile

Build and development commands tailored to your tech stack:
- `make build` - Build the application
- `make test` - Run tests
- `make lint` - Run linter
- `make run` - Run the application

### Dockerfile (optional)

Optimized multi-stage Dockerfile for containerized deployment.

### docker-compose.yml (optional)

Service orchestration including databases, caching, and your application.

## Development

### Building

```bash
# Build for current platform
make build

# Run tests
make test

# Run tests with coverage
make test-coverage

# Run linter
make lint

# Format code
make fmt
```

### Project Structure

```
cloud/
├── cmd/cloud/          # Entry point
├── internal/
│   ├── command/        # CLI commands (root, new, version)
│   ├── domain/         # Core domain models
│   ├── generator/      # File generators
│   ├── interview/      # Interview system
│   │   ├── questions/  # Question modules
│   │   └── validation/ # Input validation
│   └── suggester/      # Tech stack suggestions
├── Makefile
└── go.mod
```

### Running Tests

```bash
# Run all tests
make test

# Run with verbose output
go test -v ./...

# Run specific package tests
go test -v ./internal/generator/...
```

## Example Workflow

1. **Start a new project**:
   ```bash
   cloud new --output ./my-saas-app
   ```

2. **Answer interview questions** about your project requirements

3. **Review generated files**:
   ```
   my-saas-app/
   ├── CLAUDE.md           # Project specification
   ├── Makefile            # Build commands
   ├── Dockerfile          # Container config
   └── docker-compose.yml  # Service orchestration
   ```

4. **Use with Claude Code**:
   ```bash
   cd my-saas-app
   claude   # Claude Code will read CLAUDE.md for context
   ```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes with tests
4. Run `make ci` to ensure all checks pass
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

## Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Huh](https://github.com/charmbracelet/huh) - Interactive forms
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
