package logger

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// HTTPMiddleware creates a logging middleware for HTTP requests
func HTTPMiddleware(log Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get client IP
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		bodySize := c.Writer.Size()

		// Create log fields
		fields := Fields{
			"method":     method,
			"path":       path,
			"raw_query":  raw,
			"client_ip":  clientIP,
			"status":     statusCode,
			"latency":    latency.String(),
			"body_size":  bodySize,
			"user_agent": c.Request.UserAgent(),
		}

		// Add error information if any
		if len(c.Errors) > 0 {
			fields["errors"] = c.Errors.String()
		}

		// Log based on status code
		ctx := c.Request.Context()
		switch {
		case statusCode >= 500:
			log.Error(ctx, "HTTP Server Error", nil, fields)
		case statusCode >= 400:
			log.Warn(ctx, "HTTP Client Error", fields)
		case statusCode >= 300:
			log.Info(ctx, "HTTP Redirect", fields)
		default:
			log.Info(ctx, "HTTP Request", fields)
		}
	}
}

// RequestLogger creates a detailed request logger middleware
func RequestLogger(log Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		// Add request ID to context
		ctx := context.WithValue(c.Request.Context(), "request_id", requestID)
		c.Request = c.Request.WithContext(ctx)

		// Log request start
		log.Info(ctx, "Request started", Fields{
			"request_id": requestID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"client_ip":  c.ClientIP(),
		})

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log request completion
		fields := Fields{
			"request_id": requestID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     c.Writer.Status(),
			"latency":    latency.String(),
			"body_size":  c.Writer.Size(),
		}

		// Add error information if any
		if len(c.Errors) > 0 {
			fields["errors"] = c.Errors.String()
		}

		ctx = c.Request.Context()
		switch {
		case c.Writer.Status() >= 500:
			log.Error(ctx, "Request failed with server error", nil, fields)
		case c.Writer.Status() >= 400:
			log.Warn(ctx, "Request failed with client error", fields)
		default:
			log.Info(ctx, "Request completed", fields)
		}
	}
}

// generateRequestID generates a simple request ID
func generateRequestID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// randomString generates a random string of given length
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
