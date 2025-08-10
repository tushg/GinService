package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles HTTP requests for health endpoints
type HealthHandler struct {
	service HealthService
}

// NewHealthHandler creates a new health handler instance
func NewHealthHandler(service HealthService) *HealthHandler {
	return &HealthHandler{
		service: service,
	}
}

// GetHealth handles GET /api/v1/health requests
func (h *HealthHandler) GetHealth(c *gin.Context) {
	ctx := c.Request.Context()
	
	response, err := h.service.GetHealth(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get health status",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetReadiness handles GET /api/v1/health/ready requests
func (h *HealthHandler) GetReadiness(c *gin.Context) {
	ctx := c.Request.Context()
	
	response, err := h.service.GetReadiness(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get readiness status",
		})
		return
	}

	// Return appropriate status code based on readiness
	if response.Status == "ready" {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusServiceUnavailable, response)
	}
}

// GetLiveness handles GET /api/v1/health/live requests
func (h *HealthHandler) GetLiveness(c *gin.Context) {
	ctx := c.Request.Context()
	
	response, err := h.service.GetLiveness(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get liveness status",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
