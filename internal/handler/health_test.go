package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler_GetHealth(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.New()
	handler := NewHealthHandler()
	router.GET("/health", handler.GetHealth)

	// Create a test request
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "healthy")
	assert.Contains(t, w.Body.String(), "gin-service")
}

func TestHealthHandler_GetReadiness(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewHealthHandler()
	router.GET("/ready", handler.GetReadiness)

	req, _ := http.NewRequest("GET", "/ready", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "ready")
}

func TestHealthHandler_GetLiveness(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewHealthHandler()
	router.GET("/live", handler.GetLiveness)

	req, _ := http.NewRequest("GET", "/live", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "alive")
}
