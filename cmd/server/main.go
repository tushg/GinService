package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-service/internal/config"
	"gin-service/internal/logger"
	"gin-service/internal/middleware"
	"gin-service/internal/resources/health"
	"gin-service/internal/resources/product"
	"gin-service/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logConfig := &logger.Config{
		Level:      cfg.Log.Level,
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
