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
│   ├── middleware/
│   │   └── middleware.go    # HTTP middleware
│   ├── server/
│   │   └── server.go        # HTTP server configuration
│   └── resources/           # Resource-based modules
│       ├── health/          # Health check resource
│       │   ├── interfaces.go # Service & Repository interfaces
│       │   ├── models.go     # Request/Response models
│       │   ├── repository.go # Data access layer
│       │   ├── service.go    # Business logic layer
│       │   ├── handler.go    # HTTP request handling
│       │   └── service_test.go # Unit tests
│       └── product/         # Product resource
│           ├── interfaces.go # Service & Repository interfaces
│           ├── models.go     # Request/Response models
│           ├── repository.go # Data access layer
│           ├── service.go    # Business logic layer
│           └── handler.go    # HTTP request handling
├── configs/
│   └── config.yaml          # Configuration file
├── go.mod                   # Go module file
└── README.md               # This file
```

## Features

- **Modular Architecture**: Resource-based organization with clear separation of concerns
- **Interface-Based Design**: Easy dependency injection and unit testing
- **Health Check Endpoints**: `/api/v1/health`, `/api/v1/health/ready`, `/api/v1/health/live`
- **Product Management**: Full CRUD operations for products with validation
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
  "version": "1.0.0",
  "details": {
    "database": "healthy",
    "external_services": ["all services healthy"],
    "uptime": "1h30m45s"
  }
}
```

#### Product Management
- `POST /api/v1/products` - Create a new product
- `GET /api/v1/products` - Get all products (with pagination)
- `GET /api/v1/products/:id` - Get a specific product
- `PUT /api/v1/products/:id` - Update a product
- `DELETE /api/v1/products/:id` - Delete a product

Example product creation:
```json
POST /api/v1/products
{
  "name": "Test Product",
  "description": "A test product",
  "price": 29.99,
  "category": "Electronics",
  "stock": 10
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

### Architecture Explanation

#### **Layered Architecture**
- **Handler Layer**: HTTP request/response handling, validation, and routing
- **Service Layer**: Business logic, orchestration, and domain rules
- **Repository Layer**: Data access and persistence logic
- **Interface Layer**: Contracts for dependency injection and testing

#### **Resource-Based Organization**
Each resource (health, product, etc.) has its own folder with:
- `interfaces.go`: Service and Repository interfaces
- `models.go`: Request/Response data structures
- `repository.go`: Data access implementation
- `service.go`: Business logic implementation
- `handler.go`: HTTP handling implementation
- `*_test.go`: Unit tests

#### **Benefits**
- **Separation of Concerns**: Each layer has a single responsibility
- **Testability**: Interfaces enable easy mocking and unit testing
- **Maintainability**: Clear structure makes code easy to understand and modify
- **Scalability**: Easy to add new resources following the same pattern
- **Dependency Injection**: Services depend on interfaces, not concrete implementations

### Adding New Resources

1. Create a new folder in `internal/resources/` (e.g., `user/`)
2. Create the required files: `interfaces.go`, `models.go`, `repository.go`, `service.go`, `handler.go`
3. Implement the interfaces following the existing patterns
4. Add routes in `cmd/server/main.go`
5. Add unit tests for each layer

## License

This project is licensed under the MIT License.
