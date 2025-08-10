# Step 3: Update Main.go and Test the Restructured Project

## Overview
Now that we've moved all the files to their new locations, let's update the main.go file with the new import paths and test that everything works.

## Step 3.1: Update Main.go Import Paths

### File: `cmd/server/main.go`
Update the import statements to use the new pkg paths:

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-service/pkg/config"
	"gin-service/pkg/logger"
	"gin-service/pkg/middleware"
	"gin-service/internal/resources/health"
	"gin-service/internal/resources/product"
	"gin-service/pkg/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Convert string level to logger.Level
	var logLevel logger.Level
	switch cfg.Log.Level {
	case "debug":
		logLevel = logger.DebugLevel
	case "info":
		logLevel = logger.InfoLevel
	case "warn":
		logLevel = logger.WarnLevel
	case "error":
		logLevel = logger.ErrorLevel
	case "fatal":
		logLevel = logger.FatalLevel
	default:
		logLevel = logger.InfoLevel
	}

	// Initialize logger
	logConfig := &logger.Config{
		Level:      logLevel,
		Format:     cfg.Log.Format,
		Output:     cfg.Log.Output,
		FilePath:   cfg.Log.FilePath,
		MaxSize:    cfg.Log.MaxSize,
		MaxBackups: cfg.Log.MaxBackups,
		MaxAge:     cfg.Log.MaxAge,
		Compress:   cfg.Log.Compress,
		AddCaller:  cfg.Log.AddCaller,
		AddStack:   cfg.Log.AddStack,
	}

	appLogger, err := logger.NewLogger(logConfig)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Set Gin mode
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	router := gin.New()

	// Add middleware
	router.Use(logger.RequestLogger(appLogger))
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	// Initialize repositories
	healthRepo := health.NewHealthRepository()
	productRepo := product.NewProductRepository()

	// Initialize services
	healthService := health.NewHealthService(healthRepo, appLogger)
	productService := product.NewProductService(productRepo)

	// Initialize handlers
	healthHandler := health.NewHealthHandler(healthService)
	productHandler := product.NewProductHandler(productService)

	// Setup routes
	api := router.Group("/api/v1")
	{
		// Health endpoints
		healthGroup := api.Group("/health")
		{
			healthGroup.GET("", healthHandler.GetHealth)
			healthGroup.GET("/ready", healthHandler.GetReadiness)
			healthGroup.GET("/live", healthHandler.GetLiveness)
		}

		// Product endpoints
		productGroup := api.Group("/products")
		{
			productGroup.POST("", productHandler.CreateProduct)
			productGroup.GET("", productHandler.GetAllProducts)
			productGroup.GET("/:id", productHandler.GetProduct)
			productGroup.PUT("/:id", productHandler.UpdateProduct)
			productGroup.DELETE("/:id", productHandler.DeleteProduct)
		}
	}

	// Create server
	srv := server.New(cfg.Server.Port, router)

	// Start server in a goroutine
	go func() {
		appLogger.Info(context.Background(), "Starting server", logger.Fields{
			"port": cfg.Server.Port,
			"mode": cfg.Server.Mode,
		})
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.Fatal(context.Background(), "Failed to start server", err, logger.Fields{
				"port": cfg.Server.Port,
			})
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	appLogger.Info(context.Background(), "Shutting down server", logger.Fields{})

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		appLogger.Fatal(context.Background(), "Server forced to shutdown", err, logger.Fields{})
	}

	appLogger.Info(context.Background(), "Server exited", logger.Fields{})
}
```

## Step 3.2: Update Other Import Paths

### Update `internal/health/*.go` files
Change any imports from `gin-service/internal/...` to `gin-service/pkg/...`

### Update `internal/product/*.go` files
Change any imports from `gin-service/internal/...` to `gin-service/pkg/...`

## Step 3.3: Test the Restructured Project

### Test 1: Check Import Paths
```bash
# This should show any import errors
go mod tidy
```

### Test 2: Compile the Project
```bash
# Try to build the project
go build ./...
```

### Test 3: Run Tests
```bash
# Run all tests
go test ./...
```

### Test 4: Build the Server
```bash
# Build the main server
go build -o bin/server cmd/server/main.go
```

## Step 3.4: Common Import Path Issues

If you encounter import errors, here are the typical fixes:

### Logger Import Issues
```go
// OLD
"gin-service/internal/logger"

// NEW
"gin-service/pkg/logger"
```

### Config Import Issues
```go
// OLD
"gin-service/internal/config"

// NEW
"gin-service/pkg/config"
```

### Middleware Import Issues
```go
// OLD
"gin-service/internal/middleware"

// NEW
"gin-service/pkg/middleware"
```

### Server Import Issues
```go
// OLD
"gin-service/internal/server"

// NEW
"gin-service/pkg/server"
```

## Step 3.5: Verify Project Structure

After all updates, your project should look like this:
```
gin-service/
├── cmd/server/main.go          ✓ (updated imports)
├── configs/config.yaml         ✓
├── internal/
│   ├── health/                 ✓ (updated imports)
│   └── product/                ✓ (updated imports)
├── pkg/
│   ├── config/config.go        ✓
│   ├── database/               ✓
│   ├── logger/                 ✓
│   ├── middleware/             ✓
│   ├── server/                 ✓
│   └── utils/                  ✓
└── scripts/                    ✓
```

## Action Items
- [ ] Update `cmd/server/main.go` with new import paths
- [ ] Update all other Go files with new import paths
- [ ] Run `go mod tidy` to clean up
- [ ] Test compilation with `go build ./...`
- [ ] Run tests with `go test ./...`
- [ ] Build the server with `go build -o bin/server cmd/server/main.go`

## Troubleshooting
- **Import errors**: Double-check all import paths are updated
- **Build errors**: Make sure all files were moved correctly
- **Test failures**: Check that all dependencies are properly imported

## Next Steps
Once the project compiles and tests pass, we'll:
1. Test the configuration loading
2. Test the logger initialization
3. Test the server startup
4. Create additional utility functions

**Complete this step and let me know when you're ready to continue!**
