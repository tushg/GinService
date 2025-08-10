package logger

import (
	"os"
	"path/filepath"
)

// Config holds logging configuration
type Config struct {
	Level      Level  `mapstructure:"level" yaml:"level"`
	Format     string `mapstructure:"format" yaml:"format"`
	Output     string `mapstructure:"output" yaml:"output"`
	FilePath   string `mapstructure:"file_path" yaml:"file_path"`
	MaxSize    int    `mapstructure:"max_size" yaml:"max_size"`       // MB
	MaxBackups int    `mapstructure:"max_backups" yaml:"max_backups"`
	MaxAge     int    `mapstructure:"max_age" yaml:"max_age"`         // days
	Compress   bool   `mapstructure:"compress" yaml:"compress"`
	AddCaller  bool   `mapstructure:"add_caller" yaml:"add_caller"`
	AddStack   bool   `mapstructure:"add_stack" yaml:"add_stack"`
}

// DefaultConfig returns default logging configuration
func DefaultConfig() *Config {
	return &Config{
		Level:      InfoLevel,
		Format:     "json",
		Output:     "stdout",
		FilePath:   "logs/app.log",
		MaxSize:    100,    // 100MB
		MaxBackups: 3,
		MaxAge:     28,     // 28 days
		Compress:   true,
		AddCaller:  true,
		AddStack:   false,
	}
}

// Validate validates the logging configuration
func (c *Config) Validate() error {
	// Validate level
	if c.Level < DebugLevel || c.Level > FatalLevel {
		c.Level = InfoLevel
	}

	// Validate format
	if c.Format != "json" && c.Format != "text" {
		c.Format = "json"
	}

	// Validate output
	if c.Output != "stdout" && c.Output != "stderr" && c.Output != "file" {
		c.Output = "stdout"
	}

	// If file output is selected, ensure directory exists
	if c.Output == "file" {
		dir := filepath.Dir(c.FilePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	return nil
}
