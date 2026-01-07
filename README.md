# bug-free-octo-succotash
An experimental command line utility for Linux in Go driven by CoPilot

## Installation

### Prerequisites
- Go 1.24 or later
- [Task](https://taskfile.dev/) (optional, for using the Taskfile)

### Building from Source

Using Go:
```bash
go build -o bug-free-octo-succotash .
```

Using Task (recommended):
```bash
task build
```

## Usage

Run the application:
```bash
./bug-free-octo-succotash
```

Display version information:
```bash
./bug-free-octo-succotash --version
# or
./bug-free-octo-succotash -v
```

## Development

This project uses [Task](https://taskfile.dev/) for build automation.

### Available Tasks

- `task build` - Build the application with version information
- `task run` - Build and run the application
- `task version` - Display version information
- `task clean` - Remove build artifacts
- `task test` - Run tests
- `task fmt` - Format Go code
- `task vet` - Run go vet
- `task lint` - Run all linters (fmt and vet)
- `task install` - Install to $GOPATH/bin
- `task tag TAG_VERSION=v0.2.0` - Create a new git tag
- `task release` - Build a release version (runs lint, test, and build)

### Version Management

The application version is automatically determined from git tags and commit information:
- Version from `git describe --tags` or defaults to `0.1.0`
- Git commit SHA from `git rev-parse --short HEAD`
- Build timestamp in ISO 8601 format

These values are injected at build time using Go's `-ldflags`.
