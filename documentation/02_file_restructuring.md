# Step 2: File Restructuring and Movement

## Overview
Now let's move all the existing files to their new locations in the restructured project layout.

## File Movement Commands

### Step 2.1: Create New Directory Structure
```bash
# Create all the new directories
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

### Step 2.2: Move Files to New Locations

#### Move Configuration Files
```bash
# Move config files
mv internal/config/* pkg/config/
rmdir internal/config
```

#### Move Database Files
```bash
# Move database files
mv internal/database/* pkg/database/
rmdir internal/database
```

#### Move Logger Files
```bash
# Move logger files
mv internal/logger/* pkg/logger/
rmdir internal/logger
```

#### Move Middleware Files
```bash
# Move middleware files
mv internal/middleware/* pkg/middleware/
rmdir internal/middleware
```

#### Move Server Files
```bash
# Move server files
mv internal/server/* pkg/server/
rmdir internal/server
```

#### Move Utils Files
```bash
# Move utils files (if they exist in internal)
mv internal/utils/* pkg/utils/ 2>/dev/null || true
rmdir internal/utils 2>/dev/null || true
```

#### Keep Business Logic Files
```bash
# Keep health and product in internal (they're already there)
# No movement needed for these
```

## Step 2.3: Update Import Paths

After moving files, you'll need to update all import paths in Go files. Here are the main changes:

### Old Import Paths → New Import Paths
```go
// OLD
"gin-service/internal/config"
"gin-service/internal/database"
"gin-service/internal/logger"
"gin-service/internal/middleware"
"gin-service/internal/server"

// NEW
"gin-service/pkg/config"
"gin-service/pkg/database"
"gin-service/pkg/logger"
"gin-service/pkg/middleware"
"gin-service/pkg/server"
```

## Step 2.4: Files to Update Import Paths

1. **`cmd/server/main.go`** - Update all internal imports to pkg
2. **`internal/health/*.go`** - Update imports to use pkg paths
3. **`internal/product/*.go`** - Update imports to use pkg paths
4. **`pkg/logger/middleware.go`** - Update logger imports
5. **`pkg/database/manager.go`** - Update database imports

## Step 2.5: Create Scripts Folder

Create basic shell scripts in the new `scripts/` folder:

### `scripts/build.sh`
```bash
#!/bin/bash
echo "Building gin-service..."
go build -o bin/server cmd/server/main.go
echo "Build complete!"
```

### `scripts/run.sh`
```bash
#!/bin/bash
echo "Running gin-service..."
go run cmd/server/main.go
```

### `scripts/test.sh`
```bash
#!/bin/bash
echo "Running tests..."
go test -v ./...
```

## Action Items
- [ ] Create new directory structure
- [ ] Move all files to new locations
- [ ] Update import paths in all Go files
- [ ] Create basic scripts
- [ ] Test that the project compiles

## Verification Commands
```bash
# Check the new structure
tree -I 'bin|.git|documentation' -a

# Test compilation
go build ./...

# Run tests
go test ./...
```

## Troubleshooting
- If you get import errors, double-check all import paths
- Make sure all files were moved correctly
- Verify that the directory structure matches the expected layout

## Next Steps
Once restructuring is complete, we'll:
1. Update the main.go file with new import paths
2. Test the configuration system
3. Set up the logger
4. Create the server

**Complete this restructuring step and let me know when you're ready to continue!**
