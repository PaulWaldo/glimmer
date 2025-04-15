# Technical Context: Glimmer

## Technologies Used

### Programming Language
- **Go**: The entire application is written in Go, leveraging its concurrency features and strong typing.

### UI Framework
- **Fyne**: A cross-platform UI toolkit for Go that provides widgets, layouts, and event handling.
  - Used for creating the application window, tabs, cards, and other UI components.
  - Provides canvas elements for displaying images.
  - Handles user interactions like scrolling, clicking, and navigation.

### External APIs
- **Flickr API**: Used for authentication and retrieving photos, contacts, and groups.
  - OAuth-based authentication.
  - Endpoints for retrieving user's contacts, groups, and photos.
  - Image URLs for downloading and displaying photos.

### Libraries
- **gopkg.in/masci/flickr.v3**: Go client for the Flickr API.
  - Handles API requests and response parsing.
  - Provides structures for Flickr entities (photos, contacts, groups).
- **stretchr/testify**: Used for assertions in tests.

## Development Setup

### Project Structure
- **api/**: Contains code for interacting with the Flickr API.
  - `api.go`: Core API functionality and utilities.
  - `auth.go`: Authentication-related code.
  - `contactphotos.go`: Code for retrieving contact photos.
  - `groups.go`: Code for retrieving groups and group photos.
- **ui/**: Contains UI components and logic.
  - `main.go`: Main application setup and window creation.
  - `apptabs.go`: Tabbed interface implementation.
  - `authenticate.go`: UI for authentication.
  - `contactphotos.go`: UI for displaying contact photos.
  - `groupphotos.go`: UI for displaying group photos.
  - `photoview.go`: Photo viewing component.
  - `prefs.go`: Preferences management.
  - `viewstack.go`: Navigation stack for views.
- **mocks/**: Mock implementations for testing.
- **tester/**: Test applications.

### Build and Run
- Standard Go build tools are used.
- The application is started from `main.go` which calls `ui.Run()`.

## Technical Constraints

### Performance Considerations
- **Asynchronous Image Loading**: Images are loaded asynchronously to keep the UI responsive.
- **Semaphore Pattern**: Limits the number of concurrent image downloads to prevent overwhelming the system.
- **Batch Loading**: Photos are loaded in batches to minimize network usage and improve performance.

### UI Responsiveness
- **Background Operations**: API calls and image loading are performed in background goroutines.
- **Lifecycle Management**: The application waits for the UI to be fully drawn before starting operations.
- **Infinite Scrolling**: Contact photos implement infinite scrolling for efficient browsing.

### Authentication Security
- **OAuth**: Uses OAuth for secure authentication with Flickr.
- **Secure Storage**: Credentials are securely stored in application preferences.
- **Token Management**: Handles token acquisition, storage, and validation.

## Dependencies

### External Dependencies
- **gopkg.in/masci/flickr.v3**: Flickr API client.
- **fyne.io/fyne/v2**: UI toolkit.
- **stretchr/testify**: Testing library.

### Internal Dependencies
- **api**: Package for Flickr API interactions.
- **ui**: Package for user interface components.
- **mocks**: Package for mock implementations used in testing.

## Tool Usage Patterns

### Testing
- **Test-Driven Development**: Tests are written before implementation.
- **Mock HTTP Clients**: Used to simulate Flickr API responses in tests.
- **UI Component Testing**: Tests for UI components verify rendering and behavior.

### Concurrency
- **Goroutines**: Used for background operations.
- **Channels**: Used for synchronization and communication between goroutines.
- **Semaphores**: Used to limit concurrent operations.

### Error Handling
- **Error Propagation**: Errors are propagated up the call stack.
- **Logging**: Errors are logged using Fyne's logging system.
- **User Feedback**: Error states are communicated to the user through the UI.

## Configuration
- **Application Preferences**: Used to store user settings and credentials.
- **Batch Size**: Configurable parameter for the number of photos to load in a batch.
- **Concurrent Downloads**: Configurable parameter for the number of simultaneous image downloads.
