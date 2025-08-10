# Step 6: Logger Package Setup

## Create Logger Interfaces
Create `pkg/logger/interfaces.go`:

```go
package logger

import (
	"context"
)

// Fields represents key-value pairs for structured logging
type Fields map[string]interface{}

// Logger interface defines the logging methods
type Logger interface {
	Debug(ctx context.Context, message string, fields Fields)
	Info(ctx context.Context, message string, fields Fields)
	Warn(ctx context.Context, message string, fields Fields)
	Error(ctx context.Context, message string, err error, fields Fields)
	Fatal(ctx context.Context, message string, err error, fields Fields)
}
```

## Create Logger Configuration
Create `pkg/logger/config.go`:

```go
package logger

import (
	"time"

	"gin-service/pkg/config"
)

// LoggerConfig holds logger configuration
type LoggerConfig struct {
	Level  string
	Format string
	Output string
	File   FileConfig
}

// FileConfig holds file logging configuration
type FileConfig struct {
	Enabled   bool
	Path      string
	MaxSize   int
	MaxAge    int
	MaxBackups int
	Compress  bool
}

// NewLoggerConfig creates a new logger configuration from app config
func NewLoggerConfig(cfg *config.Config) *LoggerConfig {
	return &LoggerConfig{
		Level:  cfg.Logging.Level,
		Format: cfg.Logging.Format,
		Output: cfg.Logging.Output,
		File: FileConfig{
			Enabled:    cfg.Logging.File.Enabled,
			Path:       cfg.Logging.File.Path,
			MaxSize:    cfg.Logging.File.MaxSize,
			MaxAge:     cfg.Logging.File.MaxAge,
			MaxBackups: cfg.Logging.File.MaxBackups,
			Compress:   cfg.Logging.File.Compress,
		},
	}
}
```

## Create Logger Implementation
Create `pkg/logger/logger.go`:

```go
package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// ZapLogger implements Logger interface using zap
type ZapLogger struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
}

// NewLogger creates a new logger instance
func NewLogger(cfg *LoggerConfig) (*ZapLogger, error) {
	var core zapcore.Core

	// Configure encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	var encoder zapcore.Encoder
	if cfg.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Configure output
	if cfg.Output == "stdout" {
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLogLevel(cfg.Level))
	} else if cfg.File.Enabled {
		writer := &lumberjack.Logger{
			Filename:   cfg.File.Path,
			MaxSize:    cfg.File.MaxSize,
			MaxAge:     cfg.File.MaxAge,
			MaxBackups: cfg.File.MaxBackups,
			Compress:   cfg.File.Compress,
		}
		core = zapcore.NewCore(encoder, zapcore.AddSync(writer), getLogLevel(cfg.Level))
	} else {
		core = zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLogLevel(cfg.Level))
	}

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	sugar := logger.Sugar()

	return &ZapLogger{
		logger: logger,
		sugar:  sugar,
	}, nil
}

// Debug logs debug message
func (l *ZapLogger) Debug(ctx context.Context, message string, fields Fields) {
	l.logWithContext(ctx, l.sugar.Debugw, message, fields)
}

// Info logs info message
func (l *ZapLogger) Info(ctx context.Context, message string, fields Fields) {
	l.logWithContext(ctx, l.sugar.Infow, message, fields)
}

// Warn logs warning message
func (l *ZapLogger) Warn(ctx context.Context, message string, fields Fields) {
	l.logWithContext(ctx, l.sugar.Warnw, message, fields)
}

// Error logs error message
func (l *ZapLogger) Error(ctx context.Context, message string, err error, fields Fields) {
	if fields == nil {
		fields = Fields{}
	}
	fields["error"] = err.Error()
	l.logWithContext(ctx, l.sugar.Errorw, message, fields)
}

// Fatal logs fatal message and exits
func (l *ZapLogger) Fatal(ctx context.Context, message string, err error, fields Fields) {
	if fields == nil {
		fields = Fields{}
	}
	fields["error"] = err.Error()
	l.logWithContext(ctx, l.sugar.Fatalw, message, fields)
}

// logWithContext logs message with context and fields
func (l *ZapLogger) logWithContext(ctx context.Context, logFunc func(msg string, keysAndValues ...interface{}), message string, fields Fields) {
	if fields == nil {
		fields = Fields{}
	}

	// Add timestamp
	fields["timestamp"] = time.Now().UTC().Format(time.RFC3339)

	// Add request ID if available
	if requestID := getRequestID(ctx); requestID != "" {
		fields["request_id"] = requestID
	}

	// Convert fields to key-value pairs
	args := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		args = append(args, key, value)
	}

	logFunc(message, args...)
}

// getRequestID extracts request ID from context
func getRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	// You can implement request ID extraction logic here
	return ""
}

// getLogLevel converts string level to zapcore.Level
func getLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// Sync flushes any buffered log entries
func (l *ZapLogger) Sync() error {
	return l.logger.Sync()
}
```

## Create Logger Formatters
Create `pkg/logger/formatters.go`:

```go
package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

// Formatter interface for different log formats
type Formatter interface {
	Format(level, message string, fields Fields, timestamp time.Time) string
}

// JSONFormatter formats logs as JSON
type JSONFormatter struct{}

// Format formats log entry as JSON
func (f *JSONFormatter) Format(level, message string, fields Fields, timestamp time.Time) string {
	logEntry := map[string]interface{}{
		"timestamp": timestamp.UTC().Format(time.RFC3339),
		"level":     level,
		"message":   message,
	}

	// Add fields
	for key, value := range fields {
		logEntry[key] = value
	}

	jsonBytes, err := json.Marshal(logEntry)
	if err != nil {
		return fmt.Sprintf(`{"timestamp":"%s","level":"error","message":"failed to format log: %s"}`, 
			timestamp.UTC().Format(time.RFC3339), err.Error())
	}

	return string(jsonBytes)
}

// ConsoleFormatter formats logs for console output
type ConsoleFormatter struct{}

// Format formats log entry for console
func (f *ConsoleFormatter) Format(level, message string, fields Fields, timestamp time.Time) string {
	output := fmt.Sprintf("[%s] %s: %s", 
		timestamp.Format("2006-01-02 15:04:05"), 
		level, 
		message)

	if len(fields) > 0 {
		output += " | "
		for key, value := range fields {
			output += fmt.Sprintf("%s=%v ", key, value)
		}
	}

	return output
}
```

## Create Logger Handlers
Create `pkg/logger/handlers.go`:

```go
package logger

import (
	"os"
	"path/filepath"
)

// Handler interface for different log outputs
type Handler interface {
	Write(data []byte) error
	Close() error
}

// FileHandler handles file-based logging
type FileHandler struct {
	file *os.File
}

// NewFileHandler creates a new file handler
func NewFileHandler(filepath string) (*FileHandler, error) {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filepath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &FileHandler{file: file}, nil
}

// Write writes data to file
func (h *FileHandler) Write(data []byte) error {
	_, err := h.file.Write(data)
	return err
}

// Close closes the file
func (h *FileHandler) Close() error {
	return h.file.Close()
}

// ConsoleHandler handles console output
type ConsoleHandler struct{}

// Write writes data to console
func (h *ConsoleHandler) Write(data []byte) error {
	_, err := os.Stdout.Write(data)
	return err
}

// Close does nothing for console handler
func (h *ConsoleHandler) Close() error {
	return nil
}
```

## Create Logger Middleware
Create `pkg/logger/middleware.go`:

```go
package logger

import (
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware creates a logging middleware for Gin
func LoggingMiddleware(logger Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log request details
		fields := Fields{
			"method":     c.Request.Method,
			"path":       path,
			"raw_query":  raw,
			"status":     c.Writer.Status(),
			"latency":    latency.String(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}

		if len(c.Errors) > 0 {
			fields["errors"] = c.Errors.String()
		}

		// Log based on status code
		switch {
		case c.Writer.Status() >= 500:
			logger.Error(c.Request.Context(), "Server error", nil, fields)
		case c.Writer.Status() >= 400:
			logger.Warn(c.Request.Context(), "Client error", fields)
		default:
			logger.Info(c.Request.Context(), "Request completed", fields)
		}
	}
}
```

## Create Logger Test File
Create `pkg/logger/logger_test.go`:

```go
package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	cfg := &LoggerConfig{
		Level:  "info",
		Format: "json",
		Output: "stdout",
	}

	logger, err := NewLogger(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, logger)
}

func TestLoggerMethods(t *testing.T) {
	cfg := &LoggerConfig{
		Level:  "debug",
		Format: "console",
		Output: "stdout",
	}

	logger, err := NewLogger(cfg)
	assert.NoError(t, err)

	ctx := context.Background()
	fields := Fields{"key": "value"}

	// Test all log methods
	logger.Debug(ctx, "debug message", fields)
	logger.Info(ctx, "info message", fields)
	logger.Warn(ctx, "warn message", fields)
	logger.Error(ctx, "error message", assert.AnError, fields)

	// Test without fields
	logger.Info(ctx, "message without fields", nil)
}

func TestGetLogLevel(t *testing.T) {
	assert.Equal(t, zapcore.DebugLevel, getLogLevel("debug"))
	assert.Equal(t, zapcore.InfoLevel, getLogLevel("info"))
	assert.Equal(t, zapcore.WarnLevel, getLogLevel("warn"))
	assert.Equal(t, zapcore.ErrorLevel, getLogLevel("error"))
	assert.Equal(t, zapcore.FatalLevel, getLogLevel("fatal"))
	assert.Equal(t, zapcore.InfoLevel, getLogLevel("unknown"))
}
```

## Verify Logger Package
```bash
# Check if all logger files are created
ls -la pkg/logger/

# Expected output should show:
# config.go
# formatters.go
# handler.go
# interfaces.go
# logger.go
# logger_test.go
# middleware.go
```

## Next Steps
After creating the logger package, proceed to the next file to create the database package.
