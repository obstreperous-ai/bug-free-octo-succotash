# bug-free-octo-succotash

An experimental command line utility for Linux written in Go and driven by GitHub Copilot.

[![Tests](https://github.com/obstreperous-ai/bug-free-octo-succotash/workflows/Tests/badge.svg)](https://github.com/obstreperous-ai/bug-free-octo-succotash/actions)
[![Code Quality](https://github.com/obstreperous-ai/bug-free-octo-succotash/workflows/Code%20Quality/badge.svg)](https://github.com/obstreperous-ai/bug-free-octo-succotash/actions)
[![CodeQL](https://github.com/obstreperous-ai/bug-free-octo-succotash/workflows/CodeQL%20Security%20Scan/badge.svg)](https://github.com/obstreperous-ai/bug-free-octo-succotash/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/obstreperous-ai/bug-free-octo-succotash)](https://goreportcard.com/report/github.com/obstreperous-ai/bug-free-octo-succotash)

## Features

- 🚀 Fast and efficient command line interface
- 🧪 Test-driven development with high code coverage
- 🔒 Security-focused with CodeQL analysis
- 📦 Easy to build and install
- 🛠️ Extensible architecture

## Prerequisites

- Go 1.22 or higher
- [Task](https://taskfile.dev/) (optional, for using the Taskfile)

## Installation

### From Source

```bash
git clone https://github.com/obstreperous-ai/bug-free-octo-succotash.git
cd bug-free-octo-succotash
task build
# Or without Task:
go build -o cliutil .
```

### Using Go Install

```bash
go install github.com/obstreperous-ai/bug-free-octo-succotash@latest
```

## Usage

```bash
# Run the CLI utility
./cliutil

# Show help
./cliutil --help

# Show version information
./cliutil version

# Enable verbose output
./cliutil --verbose
```

## Development

### Quick Start

This project is optimized for development in GitHub Codespaces. Simply open the repository in Codespaces, and all development tools will be automatically configured.

### Development Environment

#### Local Development

1. Install Go 1.22 or higher
2. Install Task: `go install github.com/go-task/task/v3/cmd/task@latest`
3. Install golangci-lint: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
4. Clone the repository and run `go mod download`

#### Using Devcontainer (Recommended)

The project includes a devcontainer configuration for VS Code and GitHub Codespaces:

```bash
# Open in VS Code with Dev Containers extension
# Or create a GitHub Codespace from the repository
```

### Building

```bash
# Using Task (recommended)
task build

# Using Go directly
go build -o cliutil .
```

### Testing

```bash
# Run all tests
task test

# Run tests with coverage
task test-coverage

# Run tests for a specific package
go test -v ./internal/config
```

### Code Quality

```bash
# Run all checks (format, vet, lint, test)
task check

# Format code
task fmt

# Run linter
task lint

# Run go vet
task vet
```

### Available Tasks

Run `task` or `task --list` to see all available tasks:

- `build` - Build the application
- `test` - Run tests
- `test-coverage` - Run tests with coverage report
- `lint` - Run linters
- `fmt` - Format code
- `vet` - Run go vet
- `tidy` - Tidy dependencies
- `clean` - Clean build artifacts
- `run` - Run the application
- `check` - Run all checks
- `ci` - Run CI checks
- `install` - Install to $GOPATH/bin

## Project Structure

```
.
├── .devcontainer/          # Development container configuration
│   └── devcontainer.json   # VS Code devcontainer settings
├── .github/                # GitHub configuration
│   ├── copilot-instructions.md  # GitHub Copilot guidelines
│   ├── dependabot.yml      # Dependabot configuration
│   └── workflows/          # GitHub Actions workflows
│       ├── test.yml        # Test workflow
│       ├── lint.yml        # Linting workflow
│       └── codeql.yml      # Security scanning workflow
├── cmd/                    # Command-line applications
│   └── cliutil/           # CLI utility commands
│       ├── root.go        # Root command
│       └── version.go     # Version command
├── internal/              # Private application code
│   ├── config/           # Configuration handling
│   └── version/          # Version information
├── main.go               # Application entry point
├── go.mod                # Go module definition
├── go.sum                # Go module checksums
├── Taskfile.yml          # Task definitions
├── .golangci.yml         # Linter configuration
└── README.md            # This file
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Write tests for your changes (TDD approach)
4. Implement your changes
5. Run `task check` to ensure code quality
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Development Guidelines

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Write tests before implementing features (TDD)
- Maintain high test coverage (>80%)
- Use meaningful commit messages
- Keep PRs focused and small
- Update documentation as needed

See [.github/copilot-instructions.md](.github/copilot-instructions.md) for detailed coding guidelines.

## CI/CD

The project uses GitHub Actions for continuous integration:

- **Tests**: Automated testing on multiple Go versions
- **Code Quality**: Linting and formatting checks
- **Security**: CodeQL security scanning
- **Dependencies**: Automated dependency updates via Dependabot

## Security

Security issues should be reported privately. Please see our security policy for details.

The project includes:
- CodeQL security scanning
- Dependabot for dependency updates
- Regular security audits via GitHub Actions

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI framework
- Uses [Task](https://taskfile.dev/) for build automation
- Developed with [GitHub Copilot](https://github.com/features/copilot)

## Contact

Project Link: [https://github.com/obstreperous-ai/bug-free-octo-succotash](https://github.com/obstreperous-ai/bug-free-octo-succotash)
