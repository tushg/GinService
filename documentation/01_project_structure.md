# Step 1: Project Structure and Initial Setup

## Overview
This document will guide you through creating a complete Gin service project structure from scratch.

## What We're Building
A RESTful API service using:
- **Gin** - HTTP web framework
- **PostgreSQL** - Database
- **Structured logging** - Custom logger implementation
- **Configuration management** - Using Viper
- **Docker** - Containerization
- **Clean Architecture** - Separation of concerns

## Project Structure
```
gin-service/
├── cmd/
│   └── server/
│       └── main.go
├── configs/
│   └── config.yaml
├── internal/
│   ├── config/
│   ├── database/
│   ├── logger/
│   ├── middleware/
│   ├── resources/
│   ├── server/
│   └── shared/
├── pkg/
│   ├── common/
│   ├── constants/
│   └── utils/
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
└── docker-compose.yml
```

## Prerequisites
Before starting, ensure you have:
- Go 1.21+ installed
- Git installed
- A code editor (VS Code recommended)
- Docker (optional, for later steps)

## Next Steps
1. Create the root directory
2. Initialize Go module
3. Create the folder structure
4. Set up basic dependencies

## Action Items
- [ ] Create a new directory called `gin-service`
- [ ] Navigate to the directory
- [ ] Run `go mod init gin-service`
- [ ] Create all the folders listed above

## Commands to Run
```bash
mkdir gin-service
cd gin-service
go mod init gin-service
mkdir -p cmd/server
mkdir -p configs
mkdir -p internal/config
mkdir -p internal/database
mkdir -p internal/logger
mkdir -p internal/middleware
mkdir -p internal/resources
mkdir -p internal/server
mkdir -p internal/shared
mkdir -p pkg/common
mkdir -p pkg/constants
mkdir -p pkg/utils
```

**Complete these steps and let me know when you're ready for the next part!**
