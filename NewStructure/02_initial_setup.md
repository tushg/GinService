# Step 2: Initial Project Setup

## Create Root Directory
```bash
# Create the main project directory
mkdir GinService
cd GinService
```

## Initialize Go Module
```bash
# Initialize the Go module
go mod init gin-service

# Expected output:
# go: creating new go.mod for module gin-service
# go: to add module, create a .gitignore first, run: go mod tidy after adding code
```

## Create .gitignore File
```bash
# Create .gitignore file
echo "# Binaries for programs and plugins
bin/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work

# Environment variables
.env
.env.local

# IDE files
.vscode/
.idea/
*.swp
*.swo

# OS generated files
.DS_Store
.DS_Store?
._*
.Spotlight-V100
.Trashes
ehthumbs.db
Thumbs.db

# Logs
*.log

# Docker
.dockerignore" > .gitignore
```

## Initialize Git Repository
```bash
# Initialize git repository
git init

# Add all files
git add .

# Initial commit
git commit -m "Initial project setup"
```

## Create Basic Directory Structure
```bash
# Create main directories
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

## Verify Directory Structure
```bash
# Check the created structure (Windows)
tree /f

# Expected output should show all directories created
```

## Next Steps
After completing this step, proceed to the next file to add Go dependencies and create the configuration files.
