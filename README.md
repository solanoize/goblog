# GoBlog

[![Go Build Status](https://github.com/solanoize/goblog/actions/workflows/go.yml/badge.svg)](https://github.com/solanoize/goblog/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/solanoize/goblog)](https://goreportcard.com/report/github.com/solanoize/goblog)

A simple blogging application API built with Go, focusing on a clear, modular, and resource-oriented architecture.

## Features

- User management (registration, authentication)
- JWT-based authentication middleware
- Modular design with clear separation of concerns (auth, users, config)
- Centralized configuration for logging, database, and server
- Pagination and validation utilities
- PostgreSQL integration

## Important Aspects & Key Concepts

This project is structured with an emphasis on **modularity** and **resource-oriented design**, primarily within the `internal` package. Key principles include:

*   **Domain-Driven Structure:** Code is organized by domain or resource (e.g., `users`, `auth`), encapsulating all related logic and components within its dedicated package.
*   **Layered Architecture:** Within each domain, there's a clear separation of concerns, typically involving:
    *   **Contracts:** Interfaces defining operations.
    *   **Models:** Data structures.
    *   **Repositories:** Data access logic.
    *   **Services:** Business logic.
    *   **Controllers/Resources:** Handles API requests and responses.
    *   **Middleware:** Request pre-processing and post-processing.
*   **Centralized Configuration:** The `config` package consolidates all application settings, making it easy to manage and scale.
*   **Reusable Utilities:** Common functionalities like pagination and validation are abstracted into the `utils` package.

## Folder Structure (Design by Resources)

The project adheres to a "Design by Resources" approach, particularly within the `internal` directory, where each sub-directory represents a distinct domain or functional module.

```
goblog/
├── .github/                       # GitHub Actions workflows
│   └── workflows/
│       └── go.yml                 # CI/CD pipeline for Go builds
├── internal/                      # Internal application logic, not exposed externally
│   ├── apps/                      # Application bootstrapping and core initialization
│   │   └── bootstrap.go           # Entry point for app setup
│   ├── auth/                      # Authentication and authorization domain
│   │   ├── contract.go            # Interfaces for authentication operations
│   │   ├── middleware.go          # JWT authentication middleware
│   │   ├── resource.go            # Handlers for authentication-related API endpoints (e.g., login, register)
│   │   └── service.go             # Business logic for authentication
│   ├── config/                    # Application-wide configuration settings
│   │   ├── logging.go             # Logging configuration
│   │   ├── postgre.go             # PostgreSQL database connection and setup
│   │   ├── router.go              # HTTP router configuration (e.g., Chi, Gin)
│   │   └── server.go              # HTTP server configuration
│   ├── users/                     # User management domain (resource: users)
│   │   ├── contract.go            # Interfaces for user operations
│   │   ├── controller.go          # Handles HTTP requests for user-related actions
│   │   ├── middleware.go          # User-specific middleware (e.g., authorization)
│   │   ├── model.go               # User data model/entity
│   │   ├── repository.go          # Data access layer for user persistence
│   │   ├── resource.go            # Defines user-related API resources
│   │   └── service.go             # Business logic for user management
│   └── utils/                     # General utility functions and helpers
│       ├── middleware.go          # Common middleware functions
│       ├── paginate_response.go   # Helpers for structuring paginated API responses
│       ├── pagination.go          # Logic for handling pagination parameters
│       ├── render.go              # Utilities for rendering API responses
│       └── validation.go          # Request data validation utilities
├── docker-compose.yaml            # Docker Compose configuration for local development (e.g., database)
├── go.mod                         # Go module definition
├── go.sum                         # Go module checksums
├── main.go                        # Main application entry point
└── README.md                      # Project documentation
```

## Quick Setup

### Prerequisites

- Go 1.20 or higher (recommended)
- Git
- Docker (optional, for database setup)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/goblog.git
cd goblog
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o goblog
```

4. Run the application:
```bash
./goblog
```

The application will start on `http://localhost:5000`

### Docker Setup (Optional)

If you prefer to run the database using Docker:

```bash
docker-compose up -d
```

To stop the containers:
```bash
docker-compose down
```

To reset volumes:
```bash
docker-compose down -v
```

The application will be accessible at `http://localhost:5000`

## Configuration

Set environment variables as needed:

```bash
JWT_SECRET=supersecretkey

POSTGRES_HOST=localhost
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=goblog_db
POSTGRES_PORT=5432

SERVER_PORT=:5000
```


## Usage

Visit `http://localhost:5000` in your browser or use an API client to interact with the endpoints.
Detailed API documentation (e.g., Swagger) might be available or can be generated based on the resource definitions.

## Contributing

Pull requests are welcome. For major changes, please open an issue first.

## License

MIT License

Copyright (c) 2024 Solanoize

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.