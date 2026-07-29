# Contributing

Thank you for your interest in improving Wind Flight Predictor.

## Development

Clone the repository:

```bash
git clone https://github.com/irenmun/wind-flight-predictor.git
```

Install Go:

```bash
go version
```

Build the project:

```bash
go build ./cmd
```

Run tests:

```bash
go test ./...
```

## Code Style

- Follow standard Go formatting (`gofmt`)
- Keep functions focused on a single responsibility
- Prefer descriptive variable names
- Write comments for exported functions

## Pull Requests

Before submitting a pull request:

- Ensure all tests pass
- Run `go vet ./...`
- Run `gofmt`
- Update documentation if necessary

## Commit Message Example

```
feat(predictor): improve wind compensation algorithm
```

## Reporting Issues

Please include:

- Go version
- Operating system
- Steps to reproduce
- Expected behavior
- Actual behavior
