# OpenCode.md

## Build, Lint, and Test Commands
- **Build**: `go build` to build the project.
- **Lint**: `golint ./...` for linting the code.
- **Test All**: `go test ./...` to run all tests.
- **Run a Single Test**: `go test -run TestFunctionName` to execute a specific test. Replace `TestFunctionName` with the name of the test to be run.

## Code Style Guidelines
- **Imports**: Use double quotes for package imports. Group standard library and third-party imports separately.
- **Formatting**: Follow `go fmt` for consistent formatting across the codebase.
- **Types**: Prefer using the built-in types; use structs for complex data models. Structure codes for clarity and maintainability.
- **Naming Conventions**: Use CamelCase for exported identifiers and lowercase for unexported identifiers. Function names should be descriptive.
- **Error Handling**: Always check for errors immediately after function calls that return an error. Handle errors gracefully and provide useful feedback to the user.