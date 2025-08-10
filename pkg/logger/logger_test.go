package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger_Levels(t *testing.T) {
	config := &Config{
		Level:     InfoLevel,
		Format:    "json",
		Output:    "stdout",
		AddCaller: false,
		AddStack:  false,
	}

	log, err := NewLogger(config)
	assert.NoError(t, err)
	assert.NotNil(t, log)

	ctx := context.Background()

	// Test different log levels
	log.Debug(ctx, "Debug message", Fields{"key": "value"})
	log.Info(ctx, "Info message", Fields{"key": "value"})
	log.Warn(ctx, "Warning message", Fields{"key": "value"})
	log.Error(ctx, "Error message", nil, Fields{"key": "value"})
}

func TestLogger_WithFields(t *testing.T) {
	config := &Config{
		Level:     DebugLevel,
		Format:    "text",
		Output:    "stdout",
		AddCaller: false,
		AddStack:  false,
	}

	log, err := NewLogger(config)
	assert.NoError(t, err)

	ctx := context.Background()

	// Test with fields
	logWithFields := log.WithFields(Fields{
		"service": "test",
		"version": "1.0.0",
	})

	logWithFields.Info(ctx, "Message with fields", Fields{"additional": "data"})
}

func TestLogger_WithContext(t *testing.T) {
	config := &Config{
		Level:     InfoLevel,
		Format:    "json",
		Output:    "stdout",
		AddCaller: false,
		AddStack:  false,
	}

	log, err := NewLogger(config)
	assert.NoError(t, err)

	ctx := context.WithValue(context.Background(), "request_id", "test-123")
	logWithContext := log.WithContext(ctx)

	logWithContext.Info(ctx, "Message with context", Fields{})
}

func TestLevel_String(t *testing.T) {
	assert.Equal(t, "DEBUG", DebugLevel.String())
	assert.Equal(t, "INFO", InfoLevel.String())
	assert.Equal(t, "WARN", WarnLevel.String())
	assert.Equal(t, "ERROR", ErrorLevel.String())
	assert.Equal(t, "FATAL", FatalLevel.String())
}
