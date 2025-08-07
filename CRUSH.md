# healthchecks-go Development Guide

## Commands
```bash
# Run all tests
 go test ./...

# Run specific test
 go test -run TestHealthCheckFunc_Run

# Run tests with coverage
 go test -cover ./...

# Format code
 go fmt ./...

# Run go vet for linting
 go vet ./...

# Install dependencies
 go mod tidy
```

## Code Style Guidelines
- **Package**: Use `package healthchecks` for all files
- **Imports**: Group stdlib imports first, then external deps (none currently)
- **Interfaces**: Define minimal interfaces (e.g., HealthCheck with single Run method)
- **Naming**: Use descriptive names (HealthCheckFunc, HealthCheckNamespace)
- **Error Handling**: Return multiple values (live, ready, status) instead of errors
- **Testing**: Use table-driven tests where appropriate, test files end with `_test.go`
- **Functions**: Exported functions start with capital letter, unexported with lowercase
- **Comments**: Minimal comments, code should be self-documenting
- **Type Assertions**: Use comma-ok idiom for safe type assertions
- **Maps**: Initialize maps before use, check for nil/empty maps
- **HTTP Handlers**: Return appropriate status codes (200 OK, 503 Service Unavailable)

## Reminders
Always check the go.mod file at the start of a session to avoid referring to the module by its package name.
