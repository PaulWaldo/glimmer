# System Patterns

## System Architecture
The application will follow a layered architecture:

1. **UI Layer:**  Handles user interaction and display. Built using Go's `net/http` package for simplicity and cross-platform compatibility.  Will use HTML templates for rendering views.
2. **Application Logic Layer:** Contains the core application logic, including:
    - Flickr API interaction (authentication, data fetching).
    - Data processing and caching.
    - Management of photo views (group photos, contact photos).
3. **Data Access Layer:**  Abstracts Flickr API interactions. Provides functions to fetch photos, groups, contacts, etc.  This layer will handle API request construction, response parsing, and error handling.

## Key Technical Decisions
- **Go for UI:** Using Go's `net/http` package for the UI to minimize dependencies and ensure cross-platform compatibility.
- **Flickr API Library:**  Choose a suitable Go library for interacting with the Flickr API or implement custom API client functions.
- **Caching Strategy:** Implement a caching mechanism (in-memory or persistent) to reduce API calls and improve performance.
- **Concurrency:** Utilize goroutines and channels for concurrent API requests and data processing to enhance responsiveness.

## Design Patterns
- **MVC (Model-View-Controller) or similar:**  Organize UI, application logic, and data handling into distinct components.  The UI Layer will act as the View and Controller, while the Application Logic Layer will manage the Model and business logic.
- **Repository Pattern (Data Access Layer):**  Isolate data access logic in a dedicated layer to improve maintainability and testability.
- **Strategy Pattern (API Authentication):** Potentially use the Strategy pattern to handle different Flickr API authentication methods if needed in the future.

## Component Relationships
- UI Layer interacts with the Application Logic Layer to request data and trigger actions.
- Application Logic Layer uses the Data Access Layer to fetch data from the Flickr API.
- Data Access Layer directly interacts with the Flickr API.

## Critical Implementation Paths
1. **Flickr Authentication Flow:** Implement the OAuth 1.0a authentication flow to obtain user authorization.
2. **Photo Fetching:** Implement functions to fetch group photos and contact photos from the Flickr API, handling pagination and rate limits.
3. **UI Rendering:**  Set up HTML templates to display photos and application UI elements.
4. **Data Caching:** Implement a basic caching mechanism to store fetched photo data.
