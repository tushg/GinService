# Step 9: Server Package Setup

## Create Server Package
Create `pkg/server/server.go`:

```go
package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gin-service/pkg/config"
	"gin-service/pkg/logger"
	"gin-service/pkg/middleware"
)

// Server represents the HTTP server
type Server struct {
	config *config.Config
	logger logger.Logger
	router *gin.Engine
	server *http.Server
}

// NewServer creates a new server instance
func NewServer(cfg *config.Config, log logger.Logger) *Server {
	// Set Gin mode based on environment
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Add middleware
	router.Use(middleware.RecoveryMiddleware(log))
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.SecurityMiddleware())
	router.Use(middleware.LoggingMiddleware(log))
	router.Use(middleware.TimeoutMiddleware(30 * time.Second))

	server := &Server{
		config: cfg,
		logger: log,
		router: router,
	}

	return server
}

// SetupRoutes configures all the routes
func (s *Server) SetupRoutes(handlers ...interface{}) {
	// Health check endpoint
	s.router.GET("/health", s.healthCheck)

	// API v1 routes
	v1 := s.router.Group("/api/v1")
	{
		// Health routes
		if healthHandler, ok := handlers[0].(interface{ SetupRoutes(*gin.RouterGroup) }); ok {
			healthHandler.SetupRoutes(v1.Group("/health"))
		}

		// Product routes
		if productHandler, ok := handlers[1].(interface{ SetupRoutes(*gin.RouterGroup) }); ok {
			productHandler.SetupRoutes(v1.Group("/products"))
		}
	}

	// 404 handler
	s.router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "The requested resource was not found",
		})
	})
}

// healthCheck handles health check requests
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"service":   "gin-service",
		"version":   "1.0.0",
	})
}

// Start starts the server
func (s *Server) Start() error {
	addr := fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port)
	
	s.server = &http.Server{
		Addr:           addr,
		Handler:        s.router,
		ReadTimeout:    s.config.Server.ReadTimeout,
		WriteTimeout:   s.config.Server.WriteTimeout,
		MaxHeaderBytes: s.config.Server.MaxHeaderBytes,
	}

	s.logger.Info(context.Background(), "Starting server", logger.Fields{
		"address": addr,
		"env":     s.config.Env,
	})

	return s.server.ListenAndServe()
}

// Stop gracefully stops the server
func (s *Server) Stop(ctx context.Context) error {
	if s.server != nil {
		s.logger.Info(ctx, "Stopping server", nil)
		return s.server.Shutdown(ctx)
	}
	return nil
}

// GetRouter returns the Gin router for testing
func (s *Server) GetRouter() *gin.Engine {
	return s.router
}

// GetServer returns the HTTP server for testing
func (s *Server) GetServer() *http.Server {
	return s.server
}
```

## Create Server Test File
Create `pkg/server/server_test.go`:

```go
package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gin-service/pkg/config"
	"gin-service/pkg/logger"
)

// MockLogger is a mock implementation of logger.Logger
type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Debug(ctx context.Context, message string, fields logger.Fields) {}
func (m *MockLogger) Info(ctx context.Context, message string, fields logger.Fields)  {}
func (m *MockLogger) Warn(ctx context.Context, message string, fields logger.Fields)  {}
func (m *MockLogger) Error(ctx context.Context, message string, err error, fields logger.Fields) {}
func (m *MockLogger) Fatal(ctx context.Context, message string, err error, fields logger.Fields) {}

func TestNewServer(t *testing.T) {
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port:           8080,
			Host:           "localhost",
			ReadTimeout:    30 * time.Second,
			WriteTimeout:   30 * time.Second,
			MaxHeaderBytes: 1048576,
		},
		Env: "test",
	}

	logger := new(MockLogger)
	server := NewServer(cfg, logger)

	assert.NotNil(t, server)
	assert.NotNil(t, server.router)
	assert.Equal(t, cfg, server.config)
	assert.Equal(t, logger, server.logger)
}

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port: 8080,
			Host: "localhost",
		},
		Env: "test",
	}

	logger := new(MockLogger)
	server := NewServer(cfg, logger)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	
	server.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	
	// Parse response body to check content
	var response map[string]interface{}
	// You can add JSON parsing here if needed
}

func TestNoRouteHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port: 8080,
			Host: "localhost",
		},
		Env: "test",
	}

	logger := new(MockLogger)
	server := NewServer(cfg, logger)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/nonexistent", nil)
	
	server.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestServerStop(t *testing.T) {
	cfg := &config.Config{
		Server: config.ServerConfig{
			Port: 8080,
			Host: "localhost",
		},
		Env: "test",
	}

	logger := new(MockLogger)
	server := NewServer(cfg, logger)

	// Test stop without starting (should not error)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	err := server.Stop(ctx)
	assert.NoError(t, err)
}
```

## Verify Server Package
```bash
# Check if server files are created
ls -la pkg/server/

# Expected output should show:
# server.go
# server_test.go
```

## Next Steps
After creating the server package, proceed to the next file to create the utility packages.
