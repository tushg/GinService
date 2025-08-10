# Step 3: Dependencies Setup

## Add Required Go Dependencies
```bash
# Add Gin framework
go get github.com/gin-gonic/gin

# Add CORS middleware
go get github.com/gin-contrib/cors

# Add configuration management
go get github.com/spf13/viper

# Add PostgreSQL driver
go get github.com/lib/pq

# Add UUID generation
go get github.com/google/uuid

# Add structured logging
go get go.uber.org/zap

# Add log rotation
go get gopkg.in/natefinch/lumberjack.v2

# Add testing framework
go get github.com/stretchr/testify

# Add environment variables
go get github.com/joho/godotenv

# Add graceful shutdown
go get golang.org/x/net/context
```

## Verify go.mod File
After running the above commands, your `go.mod` file should look similar to this:

```go
module gin-service

go 1.21

require (
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-gonic/gin v1.9.1
	github.com/google/uuid v1.3.1
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.7
	github.com/spf13/viper v1.16.0
	github.com/stretchr/testify v1.8.4
	go.uber.org/zap v1.24.0
	golang.org/x/net v0.10.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)
```

## Download Dependencies
```bash
# Download all dependencies
go mod download

# Tidy up the module
go mod tidy

# Verify no errors
go mod verify
```

## Expected Output
```bash
# go mod download should complete without errors
# go mod tidy should show any removed dependencies
# go mod verify should show "all modules verified"
```

## Troubleshooting
If you encounter any errors:

1. **Version conflicts**: Check Go version compatibility
2. **Network issues**: Ensure you have internet access
3. **Proxy issues**: Set GOPROXY if needed:
   ```bash
   go env -w GOPROXY=https://proxy.golang.org,direct
   ```

## Next Steps
After successfully setting up dependencies, proceed to the next file to create configuration files.
