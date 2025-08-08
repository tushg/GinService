package logger

import (
	"context"
	"os"
	"sync"
	"time"
)

// logger implements the Logger interface
type logger struct {
	config  *Config
	handler Handler
	fields  Fields
	mu      sync.RWMutex
}

// NewLogger creates a new logger instance
func NewLogger(config *Config) (Logger, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Create formatter
	var formatter Formatter
	switch config.Format {
	case "json":
		formatter = &JSONFormatter{
			AddCaller: config.AddCaller,
			AddStack:  config.AddStack,
		}
	case "text":
		formatter = &TextFormatter{
			AddCaller: config.AddCaller,
			AddStack:  config.AddStack,
		}
	default:
		formatter = &JSONFormatter{
			AddCaller: config.AddCaller,
			AddStack:  config.AddStack,
		}
	}

	// Create handler based on output configuration
	var handler Handler
	switch config.Output {
	case "stdout", "stderr":
		consoleHandler, err := NewConsoleHandler(config.Output, formatter)
		if err != nil {
			return nil, err
		}
		handler = consoleHandler
	case "file":
		fileHandler, err := NewFileHandler(config, formatter)
		if err != nil {
			return nil, err
		}
		handler = fileHandler
	default:
		// Default to stdout
		consoleHandler, err := NewConsoleHandler("stdout", formatter)
		if err != nil {
			return nil, err
		}
		handler = consoleHandler
	}

	return &logger{
		config:  config,
		handler: handler,
		fields:  make(Fields),
	}, nil
}

// log logs a message at the specified level
func (l *logger) log(ctx context.Context, level Level, msg string, err error, fields Fields) {
	if level < l.config.Level {
		return
	}

	l.mu.RLock()
	baseFields := make(Fields, len(l.fields)+len(fields))
	for k, v := range l.fields {
		baseFields[k] = v
	}
	for k, v := range fields {
		baseFields[k] = v
	}
	l.mu.RUnlock()

	entry := Entry{
		Level:     level,
		Timestamp: time.Now(),
		Message:   msg,
		Fields:    baseFields,
		Error:     err,
		Context:   ctx,
	}

	// Handle the log entry
	if err := l.handler.Handle(entry); err != nil {
		// Fallback to stderr if logging fails
		fallbackHandler, _ := NewConsoleHandler("stderr", &TextFormatter{})
		fallbackHandler.Handle(Entry{
			Level:     ErrorLevel,
			Timestamp: time.Now(),
			Message:   "Failed to write log entry",
			Error:     err,
		})
	}
}

// Debug logs a debug message
func (l *logger) Debug(ctx context.Context, msg string, fields Fields) {
	l.log(ctx, DebugLevel, msg, nil, fields)
}

// Info logs an info message
func (l *logger) Info(ctx context.Context, msg string, fields Fields) {
	l.log(ctx, InfoLevel, msg, nil, fields)
}

// Warn logs a warning message
func (l *logger) Warn(ctx context.Context, msg string, fields Fields) {
	l.log(ctx, WarnLevel, msg, nil, fields)
}

// Error logs an error message
func (l *logger) Error(ctx context.Context, msg string, err error, fields Fields) {
	l.log(ctx, ErrorLevel, msg, err, fields)
}

// Fatal logs a fatal message and exits
func (l *logger) Fatal(ctx context.Context, msg string, err error, fields Fields) {
	l.log(ctx, FatalLevel, msg, err, fields)
	os.Exit(1)
}

// WithContext creates a new logger with the given context
func (l *logger) WithContext(ctx context.Context) Logger {
	return &logger{
		config:  l.config,
		handler: l.handler,
		fields:  l.fields,
	}
}

// WithFields creates a new logger with additional fields
func (l *logger) WithFields(fields Fields) Logger {
	l.mu.Lock()
	defer l.mu.Unlock()

	newFields := make(Fields, len(l.fields)+len(fields))
	for k, v := range l.fields {
		newFields[k] = v
	}
	for k, v := range fields {
		newFields[k] = v
	}

	return &logger{
		config:  l.config,
		handler: l.handler,
		fields:  newFields,
	}
}
