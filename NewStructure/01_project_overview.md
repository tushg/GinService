# Gin Service Project - New Structure Overview

## Project Description
A RESTful API service built with Go and Gin framework, featuring clean architecture with PostgreSQL database, structured logging, and containerization support.

## Technologies Used
- **Go 1.21+** - Programming language
- **Gin** - HTTP web framework
- **PostgreSQL** - Database
- **Viper** - Configuration management
- **Docker** - Containerization
- **Make** - Build automation

## New Project Structure
```
GinService/
├── cmd/
│   └── server/
│       └── main.go
├── configs/
│   └── config.yaml
├── internal/
│   ├── health/
│   │   ├── handler.go
│   │   ├── interfaces.go
│   │   ├── models.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   └── service_test.go
│   └── product/
│       ├── handler.go
│       ├── interfaces.go
│       ├── models.go
│       ├── repository.go
│       └── service.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   ├── manager.go
│   │   └── postgresql/
│   │       ├── connection.go
│   │       └── repository.go
│   ├── logger/
│   │   ├── config.go
│   │   ├── formatters.go
│   │   ├── handlers.go
│   │   ├── interfaces.go
│   │   ├── logger.go
│   │   ├── logger_test.go
│   │   └── middleware.go
│   ├── middleware/
│   │   └── middleware.go
│   ├── server/
│   │   └── server.go
│   └── utils/
│       ├── string.go
│       └── time.go
├── scripts/
│   ├── build.sh
│   ├── run.sh
│   └── test.sh
├── docker-compose.yml
├── docker-compose.dev.yml
├── Dockerfile
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## Architecture Principles
- **Clean Architecture**: Separation of concerns with clear layers
- **Dependency Injection**: Interfaces for loose coupling
- **Configuration Management**: Environment-based configuration
- **Structured Logging**: Centralized logging with different levels
- **Database Abstraction**: Repository pattern for data access
- **Middleware Support**: CORS, logging, and custom middleware
- **Containerization**: Docker support for development and production

## Prerequisites
- Go 1.21 or higher
- PostgreSQL 12 or higher
- Docker and Docker Compose
- Make (optional, for build automation)
- Git for version control

## Next Steps
Proceed to the next file to start creating the project structure step by step.
