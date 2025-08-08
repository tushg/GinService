package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

// ConsoleHandler handles logging to console (stdout/stderr)
type ConsoleHandler struct {
	output   *os.File
	formatter Formatter
	mu       sync.Mutex
}

// NewConsoleHandler creates a new console handler
func NewConsoleHandler(output string, formatter Formatter) (*ConsoleHandler, error) {
	var file *os.File
	switch output {
	case "stdout":
		file = os.Stdout
	case "stderr":
		file = os.Stderr
	default:
		return nil, fmt.Errorf("unsupported console output: %s", output)
	}

	return &ConsoleHandler{
		output:    file,
		formatter: formatter,
	}, nil
}

// Handle writes the log entry to console
func (h *ConsoleHandler) Handle(entry Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	data, err := h.formatter.Format(entry)
	if err != nil {
		return err
	}

	_, err = h.output.Write(data)
	return err
}

// Close closes the console handler
func (h *ConsoleHandler) Close() error {
	// Don't close stdout/stderr
	return nil
}

// FileHandler handles logging to file with rotation
type FileHandler struct {
	writer    *lumberjack.Logger
	formatter Formatter
	mu        sync.Mutex
}

// NewFileHandler creates a new file handler with rotation
func NewFileHandler(config *Config, formatter Formatter) (*FileHandler, error) {
	writer := &lumberjack.Logger{
		Filename:   config.FilePath,
		MaxSize:    config.MaxSize,    // MB
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,     // days
		Compress:   config.Compress,
	}

	return &FileHandler{
		writer:    writer,
		formatter: formatter,
	}, nil
}

// Handle writes the log entry to file
func (h *FileHandler) Handle(entry Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	data, err := h.formatter.Format(entry)
	if err != nil {
		return err
	}

	_, err = h.writer.Write(data)
	return err
}

// Close closes the file handler
func (h *FileHandler) Close() error {
	return h.writer.Close()
}

// MultiHandler handles multiple output destinations
type MultiHandler struct {
	handlers []Handler
}

// NewMultiHandler creates a new multi handler
func NewMultiHandler(handlers ...Handler) *MultiHandler {
	return &MultiHandler{
		handlers: handlers,
	}
}

// Handle writes the log entry to all handlers
func (h *MultiHandler) Handle(entry Entry) error {
	var lastErr error
	for _, handler := range h.handlers {
		if err := handler.Handle(entry); err != nil {
			lastErr = err
		}
	}
	return lastErr
}

// Close closes all handlers
func (h *MultiHandler) Close() error {
	var lastErr error
	for _, handler := range h.handlers {
		if err := handler.Close(); err != nil {
			lastErr = err
		}
	}
	return lastErr
}
