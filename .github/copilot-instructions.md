# GitHub Copilot Instructions

## Project Overview
This is a command line utility for Linux written in Go. The project follows Go best practices and emphasizes test-driven development (TDD), clean code, and maintainability.

## Code Style and Standards

### Go Best Practices
- Follow the [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use `gofmt` and `goimports` for consistent formatting
- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Keep functions small and focused (single responsibility)
- Use meaningful variable and function names
- Avoid naked returns in long functions

### Error Handling
- Always handle errors explicitly
- Use `fmt.Errorf` with `%w` for error wrapping
- Provide context in error messages
- Return errors rather than panicking in library code
- Only panic for truly unrecoverable errors in main package

### Testing (TDD Approach)
- Write tests BEFORE implementing functionality
- Use table-driven tests for multiple test cases
- Test file naming: `*_test.go`
- Aim for high test coverage (>80%)
- Test both success and failure paths
- Use subtests with `t.Run()` for related test cases
- Mock external dependencies
- Use testify/assert for better assertions when appropriate

### Testing Structure
```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    InputType
        expected ExpectedType
        wantErr  bool
    }{
        // test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Code Organization
- Use the `cmd/` directory for command-line applications
- Use the `internal/` directory for private application code
- Use the `pkg/` directory for public library code (if applicable)
- Keep packages focused and cohesive
- Avoid circular dependencies

### Documentation
- Write clear godoc comments for all exported functions, types, and packages
- Package comments should describe the purpose of the package
- Use complete sentences starting with the name of the element
- Example:
  ```go
  // ProcessData processes the input data and returns the result.
  // It returns an error if the data is invalid.
  func ProcessData(data string) (string, error) {
  ```

### Concurrency
- Use goroutines and channels appropriately
- Always consider race conditions
- Use `sync` package primitives when needed
- Close channels when done sending
- Use context for cancellation and timeouts

### Dependencies
- Minimize external dependencies
- Use go modules for dependency management
- Pin dependencies to specific versions
- Regularly update dependencies for security

### Security
- Validate all inputs
- Use prepared statements for database queries
- Avoid command injection vulnerabilities
- Handle sensitive data appropriately
- Use crypto/rand for random number generation

### Performance
- Profile before optimizing
- Use benchmarks to measure performance
- Consider memory allocations
- Use appropriate data structures

## Project-Specific Guidelines

### CLI Commands
- Use cobra for CLI framework
- Keep command logic in `cmd/cliutil/` directory
- Business logic should be in `internal/` packages
- Commands should be composable and follow Unix philosophy

### Configuration
- Store config in `internal/config/`
- Support environment variables
- Provide sensible defaults
- Use `~/.config/cliutil/` for user configuration

### Versioning
- Version information in `internal/version/`
- Use semantic versioning
- Set version via ldflags during build

## Common Patterns

### Constructor Pattern
```go
// New creates a new Instance with default values
func New() *Instance {
    return &Instance{
        // defaults
    }
}
```

### Options Pattern
```go
type Option func(*Config)

func WithVerbose(v bool) Option {
    return func(c *Config) {
        c.Verbose = v
    }
}
```

## Running Tests

The project uses `task test` to run all tests. This is also included in the GitHub Actions workflow.

### Test Commands
- `task test` - Run all tests with race detection and coverage
- `task test-coverage` - Run tests and generate HTML coverage report
- `go test -v ./internal/config` - Run tests for a specific package
- `go test -v -race -coverprofile=coverage.out ./...` - Run tests manually with all options

### Test Requirements
- Tests must pass before committing
- Tests are run automatically in GitHub Actions CI
- The project requires Go 1.22 or higher
- GitHub Actions tests on Go 1.22 and 1.23

### Writing Tests
- Follow TDD approach: write tests before implementing functionality
- Use table-driven tests with `t.Run()` for subtests
- Test both success and failure paths
- Keep tests in `*_test.go` files next to the code they test

## Before Committing
- Run `task fmt` to format code
- Run `task lint` to check for issues
- Run `task test` to ensure all tests pass
- Run `task check` to run all checks
- Ensure no security vulnerabilities

## When Generating Code
1. Write the test first (TDD)
2. Implement the minimal code to pass the test
3. Refactor for clarity and maintainability
4. Add documentation
5. Run all checks
