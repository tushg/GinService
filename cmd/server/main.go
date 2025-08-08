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

	// Set Gin mode
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	router := gin.New()

	// Add middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	// Initialize repositories
	healthRepo := health.NewHealthRepository()
	productRepo := product.NewProductRepository()

	// Initialize services
	healthService := health.NewHealthService(healthRepo)
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
		log.Printf("Starting server on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
