# HealthChecks-go

[![Go Reference](https://pkg.go.dev/badge/github.com/Fryuni/healthchecks-go.svg)](https://pkg.go.dev/github.com/Fryuni/healthchecks-go) [![Last Release](https://img.shields.io/github/v/release/Fryuni/healthchecks-go)](https://github.com/Fryuni/healthchecks-go/releases)

A simple Go library for running health checks across components. The package provides an interface to define health checks, use functions as checks, group checks into namespaces, and aggregate their statuses.

## Features

- Define health checks using the `HealthCheck` interface
- Use functions as health checks via `HealthCheckFunc`
- Group and manage multiple health checks with `HealthCheckNamespace`
- Aggregate individual check results into overall statuses

## Installation

```bash
go get github.com/Fryuni/healthchecks-go
```

## Usage

Register health checks and run them:

```go
package main

import (
	"fmt"
	"github.com/Fryuni/healthchecks-go"
)

func main() {
	// Register a health check function
	healthchecks.RegisterFunc("database", func() (bool, bool, interface{}) {
		// perform health check, e.g. connect to a database
		return true, true, "ok"
	})

	// Run all health checks in the root namespace
	live, ready, status := healthchecks.Run()
	fmt.Printf("Live: %v, Ready: %v, Status: %v\n", live, ready, status)
}
```

## Testing

Run tests with:

```bash
go test ./...
```

And for coverage:

```bash
go test -cover ./...
```

## Code Style

Follow the guidelines in [CRUSH.md](CRUSH.md) for development commands and code style.

## License

This project is licensed under the terms of the [LICENSE](LICENSE) file.
