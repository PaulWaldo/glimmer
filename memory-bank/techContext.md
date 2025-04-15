# Tech Context

## Technologies Used
- **Programming Language:** Go (version 1.20 or later)
- **UI Framework:** Go's `net/http` package with HTML templates.
- **Flickr API:** REST API (OAuth 1.0a authentication)
- **Testing Framework:** `testify` for unit testing.
- **Go Modules:** For dependency management.

## Development Setup
- **Go Development Environment:**  Standard Go development environment with Go toolchain installed.
- **Flickr API Key:**  Need to obtain Flickr API key and secret from Flickr Developer portal.
- **OAuth 1.0a Libraries:**  Use Go libraries for OAuth 1.0a authentication with Flickr API.

## Technical Constraints
- **Flickr API Rate Limits:**  Need to handle Flickr API rate limits gracefully. Implement caching and optimize API requests.
- **OAuth 1.0a Complexity:** OAuth 1.0a is more complex than OAuth 2.0. Need to carefully implement the authentication flow.
- **UI Simplicity:**  `net/http` UI will be basic. Focus on functionality over advanced UI features for initial version.

## Dependencies
- **Go Standard Library:**  `net/http`, `html/template`, `encoding/json`, `fmt`, `log`, `os`.
- **External Libraries:**
    -  OAuth 1.0a library for Go (e.g., `github.com/dghubble/oauth1`)
    -  HTTP client library (standard `net/http` or consider `github.com/go-resty/resty` for enhanced features)
    -  Testing: `github.com/stretchr/testify`

## Tool Usage Patterns
- **Go Toolchain:** `go build`, `go run`, `go test`, `go fmt`, `go mod`.
- **VSCode:**  For code editing, debugging, and Git integration.
- **Command Line:** For running Go commands, Git operations.
- **Mockery:** For generating mocks for testing interfaces.
