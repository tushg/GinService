# Step 2: Setting Up Dependencies

## Overview
Now that you have the folder structure, let's add the necessary Go dependencies to your project.

## Dependencies We Need
- **Gin** - HTTP web framework
- **Viper** - Configuration management
- **PostgreSQL driver** - Database connectivity
- **UUID generator** - For unique identifiers
- **Lumberjack** - Log rotation
- **CORS middleware** - Cross-origin resource sharing
- **Testing framework** - For unit tests

## Action Items
- [ ] Add all required dependencies
- [ ] Run `go mod tidy` to clean up
- [ ] Verify dependencies are properly installed

## Commands to Run
```bash
# Add main dependencies
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/spf13/viper
go get github.com/lib/pq
go get github.com/google/uuid
go get gopkg.in/natefinch/lumberjack.v2

# Add development dependencies
go get github.com/stretchr/testify

# Clean up and verify
go mod tidy
go mod download
```

## Expected go.mod Content
Your `go.mod` file should look similar to this:
```go
module gin-service

go 1.21

require (
    github.com/gin-contrib/cors v1.4.0
    github.com/gin-gonic/gin v1.9.1
    github.com/google/uuid v1.4.0
    github.com/lib/pq v1.10.9
    github.com/spf13/viper v1.17.0
    github.com/stretchr/testify v1.8.4
    gopkg.in/natefinch/lumberjack.v2 v2.2.1
)
```

## Verification
After running the commands:
1. Check that `go.mod` file was created
2. Verify `go.sum` file exists
3. Run `go mod verify` to ensure dependencies are valid

## Troubleshooting
- If you get version conflicts, try using `go get -u` to get the latest versions
- If you're behind a corporate firewall, you might need to configure Go proxy settings
- Make sure your Go version is 1.21 or higher

## Next Steps
Once dependencies are set up, we'll move to:
1. Creating the configuration structure
2. Setting up the logger
3. Creating basic models

**Complete this step and let me know when you're ready to continue!**
