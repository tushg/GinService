# Gin Service

A clean, well-structured Go service using the Gin web framework with health check endpoints.

## Project Structure

```
gin-service/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go        # Configuration management
│   ├── handler/
│   │   └── health.go        # Health check handlers
│   ├── middleware/
│   │   └── middleware.go    # HTTP middleware
│   └── server/
│       └── server.go        # HTTP server configuration
├── configs/
│   └── config.yaml          # Configuration file
├── go.mod                   # Go module file
└── README.md               # This file
```

## Features

- **Health Check Endpoints**: `/api/v1/health`, `/api/v1/health/ready`, `/api/v1/health/live`
- **Configuration Management**: Using Viper for flexible configuration
- **Middleware**: Logging, Recovery, and CORS middleware
- **Graceful Shutdown**: Proper signal handling and graceful server shutdown
- **Clean Architecture**: Well-organized folder structure following Go best practices

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd gin-service
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run cmd/server/main.go
```

The server will start on port 8080 by default.

### Configuration

The application can be configured using:

1. **Configuration file**: `configs/config.yaml`
2. **Environment variables**: 
   - `SERVER_PORT` (default: 8080)
   - `SERVER_MODE` (default: debug)

### API Endpoints

#### Health Check
- `GET /api/v1/health` - General health check
- `GET /api/v1/health/ready` - Readiness probe
- `GET /api/v1/health/live` - Liveness probe

Example response:
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00Z",
  "service": "gin-service",
  "version": "1.0.0"
}
```

### Building

Build the application:
```bash
go build -o bin/server cmd/server/main.go
```

### Running in Production

Set the mode to release:
```bash
export SERVER_MODE=release
go run cmd/server/main.go
```

## Development

### Project Structure Explanation

- **`cmd/`**: Contains the main applications of the project
- **`internal/`**: Private application and library code
- **`configs/`**: Configuration files
- **`pkg/`**: Library code that's ok to use by external applications (not used in this simple example)

### Adding New Endpoints

1. Create a new handler in `internal/handler/`
2. Add the route in `cmd/server/main.go`
3. Follow the existing patterns for consistency

## License

This project is licensed under the MIT License.
