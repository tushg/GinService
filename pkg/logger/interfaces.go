package logger

import (
	"context"
	"time"
)

// Level represents the logging level
type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// String returns the string representation of the log level
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Fields represents key-value pairs for structured logging
type Fields map[string]interface{}

// Entry represents a log entry
type Entry struct {
	Level     Level                 `json:"level"`
	Timestamp time.Time             `json:"timestamp"`
	Message   string                `json:"message"`
	Fields    Fields                `json:"fields,omitempty"`
	Error     error                 `json:"error,omitempty"`
	Context   context.Context       `json:"-"`
}

// Logger defines the interface for logging operations
type Logger interface {
	Debug(ctx context.Context, msg string, fields Fields)
	Info(ctx context.Context, msg string, fields Fields)
	Warn(ctx context.Context, msg string, fields Fields)
	Error(ctx context.Context, msg string, err error, fields Fields)
	Fatal(ctx context.Context, msg string, err error, fields Fields)
	
	WithContext(ctx context.Context) Logger
	WithFields(fields Fields) Logger
}

// Handler defines the interface for log handlers (output destinations)
type Handler interface {
	Handle(entry Entry) error
	Close() error
}

// Formatter defines the interface for log formatters
type Formatter interface {
	Format(entry Entry) ([]byte, error)
}
