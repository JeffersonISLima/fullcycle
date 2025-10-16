# Go File and Directory Naming Conventions

This guide outlines the best practices for naming files and directories in Go projects, following Go's official conventions and community standards.

## Directory Naming

### General Rules

- Use **lowercase** letters only
- Use **hyphens** (`-`) to separate words
- Avoid underscores (`_`) and camelCase
- Keep names short and descriptive

### Examples

```
✅ Good:
- user-service
- data-models
- http-handlers
- config-loader

❌ Bad:
- UserService
- data_models
- HttpHandlers
- configLoader
```

### Special Directory Names

- `cmd/` - Main applications
- `pkg/` - Library code intended for external use
- `internal/` - Private application and library code
- `api/` - OpenAPI/Swagger specs, protocol buffer files
- `web/` - Web application specific components
- `configs/` - Configuration file templates
- `init/` - System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs
- `scripts/` - Scripts to perform various build, install, analysis, etc operations
- `build/` - Packaging and Continuous Integration
- `deployments/` - IaaS, PaaS, system and container orchestration deployment configurations and templates
- `test/` - Additional external test apps and test data

## File Naming

### Go Source Files

- Use **snake_case** for file names
- File name should describe the primary functionality
- Use descriptive names that indicate purpose

### Examples

```
✅ Good:
- user_service.go
- http_handler.go
- config_loader.go
- database_connection.go
- user_test.go

❌ Bad:
- userService.go
- httphandler.go
- configloader.go
- db.go
- usertest.go
```

### Test Files

- Add `_test.go` suffix
- Follow the same naming convention as the source file

### Examples

```
Source file: user_service.go
Test file:   user_service_test.go

Source file: http_handler.go
Test file:   http_handler_test.go
```

### Build Files

- `go.mod` - Module definition
- `go.sum` - Module checksums
- `Makefile` - Build automation
- `Dockerfile` - Container configuration
- `.gitignore` - Git ignore rules

## Package Naming

### Package Directory Name

- Should match the package name
- Use lowercase, single word
- Avoid underscores, hyphens, or mixedCaps

### Examples

```
✅ Good:
Directory: user
Package:   package user

Directory: httpclient
Package:   package httpclient

❌ Bad:
Directory: user-service
Package:   package user (should be userservice)

Directory: user_service
Package:   package user_service
```

## Project Structure Example

```
my-project/
├── cmd/
│   ├── server/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── internal/
│   ├── user/
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── user_test.go
│   └── config/
│       ├── config.go
│       └── config_test.go
├── pkg/
│   ├── logger/
│   │   ├── logger.go
│   │   └── logger_test.go
│   └── validator/
│       └── validator.go
├── api/
│   └── openapi.yaml
├── scripts/
│   ├── build.sh
│   └── deploy.sh
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
└── README.md
```

## Key Principles

1. **Consistency** - Follow the same naming pattern throughout the project
2. **Clarity** - Names should clearly indicate purpose and content
3. **Simplicity** - Keep names as simple as possible while maintaining clarity
4. **Go Conventions** - Follow Go's official style guide and community standards
5. **Avoid Abbreviations** - Use full words instead of abbreviations when possible

## References

- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://golang.org/doc/effective_go.html)
