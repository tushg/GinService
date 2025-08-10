# Step 1: New Project Structure and Initial Setup

## Overview
This document will guide you through creating a complete Gin service project with the new restructured layout.

## New Project Structure
```
gin-service/
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
│   │   └── service.go
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
├── go.mod
├── go.sum
├── Makefile
├── Dockerfile
└── docker-compose.yml
```

## Key Changes from Original Structure
1. **Moved config, database, logger, middleware, server** from `internal/` to `pkg/`
2. **Simplified internal structure** to only contain business logic (health, product)
3. **Added scripts folder** for build and deployment scripts
4. **Cleaner separation** between infrastructure code (pkg) and business logic (internal)

## Prerequisites
Before starting, ensure you have:
- Go 1.21+ installed
- Git installed
- A code editor (VS Code recommended)
- Docker (optional, for later steps)

## Action Items
- [ ] Create the new directory structure
- [ ] Move existing files to new locations
- [ ] Update import paths in all Go files
- [ ] Test that the project compiles

## Commands to Run
```bash
# Create new directory structure
mkdir -p cmd/server
mkdir -p configs
mkdir -p internal/health
mkdir -p internal/product
mkdir -p pkg/config
mkdir -p pkg/database/postgresql
mkdir -p pkg/logger
mkdir -p pkg/middleware
mkdir -p pkg/server
mkdir -p pkg/utils
mkdir -p scripts
```

## File Movement Plan
1. Move `internal/config/*` → `pkg/config/*`
2. Move `internal/database/*` → `pkg/database/*`
3. Move `internal/logger/*` → `pkg/logger/*`
4. Move `internal/middleware/*` → `pkg/middleware/*`
5. Move `internal/server/*` → `pkg/server/*`
6. Move `pkg/utils/*` → `pkg/utils/*` (already in correct place)
7. Keep `internal/health/*` and `internal/product/*` in place
8. Keep `cmd/server/*` in place

## Next Steps
1. Create the new folder structure
2. Move files to new locations
3. Update import paths
4. Test compilation

**Complete these steps and let me know when you're ready for the next part!**
