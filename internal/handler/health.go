package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check requests
type HealthHandler struct{}

// NewHealthHandler creates a new health handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
}

// GetHealth handles general health check requests
func (h *HealthHandler) GetHealth(c *gin.Context) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "gin-service",
		Version:   "1.0.0",
	}

	c.JSON(http.StatusOK, response)
}

// GetReadiness handles readiness probe requests
func (h *HealthHandler) GetReadiness(c *gin.Context) {
	response := HealthResponse{
		Status:    "ready",
		Timestamp: time.Now(),
		Service:   "gin-service",
		Version:   "1.0.0",
	}

	c.JSON(http.StatusOK, response)
}

// GetLiveness handles liveness probe requests
func (h *HealthHandler) GetLiveness(c *gin.Context) {
	response := HealthResponse{
		Status:    "alive",
		Timestamp: time.Now(),
		Service:   "gin-service",
		Version:   "1.0.0",
	}

	c.JSON(http.StatusOK, response)
}
