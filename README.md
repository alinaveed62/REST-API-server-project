# REST API Server in Go

This project is a simple REST API server written in Go. It exposes endpoints to create, list, retrieve, and delete items. Data is stored in memory (no persistent storage).

rest-api-server/
├── .github/
│   └── workflows/
│       └── ci.yml         # GitHub Actions workflow for CI/CD pipelining
├── go.mod                 # Go module file
├── main.go                # Entry point: starts the HTTP server
├── handlers.go            # HTTP handlers for REST endpoints
├── models.go              # Data model definitions
├── handlers_test.go       # Test file for our REST endpoints
└── README.md              # Project documentation


## Prerequisites

- Go 1.18 or later
- VS Code (optional, for development)

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/rest-api-server.git
   cd rest-api-server
