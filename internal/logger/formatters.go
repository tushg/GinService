package logger

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// JSONFormatter formats log entries as JSON
type JSONFormatter struct {
	AddCaller bool
	AddStack  bool
}

// Format formats a log entry as JSON
func (f *JSONFormatter) Format(entry Entry) ([]byte, error) {
	data := map[string]interface{}{
		"level":     entry.Level.String(),
		"timestamp": entry.Timestamp.Format(time.RFC3339),
		"message":   entry.Message,
	}

	// Add fields
	if len(entry.Fields) > 0 {
		for k, v := range entry.Fields {
			data[k] = v
		}
	}

	// Add error if present
	if entry.Error != nil {
		data["error"] = entry.Error.Error()
	}

	// Add caller information
	if f.AddCaller {
		if pc, file, line, ok := runtime.Caller(3); ok {
			funcName := runtime.FuncForPC(pc).Name()
			data["caller"] = fmt.Sprintf("%s:%d", filepath.Base(file), line)
			data["function"] = funcName
		}
	}

	// Add stack trace for errors
	if f.AddStack && entry.Error != nil {
		var stack []string
		for i := 0; i < 10; i++ {
			if pc, file, line, ok := runtime.Caller(i + 3); ok {
				funcName := runtime.FuncForPC(pc).Name()
				stack = append(stack, fmt.Sprintf("%s:%d %s", filepath.Base(file), line, funcName))
			}
		}
		if len(stack) > 0 {
			data["stack"] = stack
		}
	}

	return json.Marshal(data)
}

// TextFormatter formats log entries as human-readable text
type TextFormatter struct {
	AddCaller bool
	AddStack  bool
}

// Format formats a log entry as text
func (f *TextFormatter) Format(entry Entry) ([]byte, error) {
	var parts []string

	// Timestamp
	parts = append(parts, entry.Timestamp.Format("2006-01-02T15:04:05.000Z07:00"))

	// Level
	parts = append(parts, fmt.Sprintf("[%s]", strings.ToUpper(entry.Level.String())))

	// Caller information
	if f.AddCaller {
		if pc, file, line, ok := runtime.Caller(3); ok {
			funcName := runtime.FuncForPC(pc).Name()
			parts = append(parts, fmt.Sprintf("(%s:%d %s)", filepath.Base(file), line, funcName))
		}
	}

	// Message
	parts = append(parts, entry.Message)

	// Fields
	if len(entry.Fields) > 0 {
		var fieldParts []string
		for k, v := range entry.Fields {
			fieldParts = append(fieldParts, fmt.Sprintf("%s=%v", k, v))
		}
		parts = append(parts, fmt.Sprintf("{%s}", strings.Join(fieldParts, " ")))
	}

	// Error
	if entry.Error != nil {
		parts = append(parts, fmt.Sprintf("error=%s", entry.Error.Error()))
	}

	// Stack trace
	if f.AddStack && entry.Error != nil {
		var stack []string
		for i := 0; i < 5; i++ {
			if pc, file, line, ok := runtime.Caller(i + 3); ok {
				funcName := runtime.FuncForPC(pc).Name()
				stack = append(stack, fmt.Sprintf("  %s:%d %s", filepath.Base(file), line, funcName))
			}
		}
		if len(stack) > 0 {
			parts = append(parts, "\nStack trace:")
			parts = append(parts, stack...)
		}
	}

	return []byte(strings.Join(parts, " ") + "\n"), nil
}
