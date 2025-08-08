package config

import (
	"fmt"
	"strings"

	"gin-service/internal/logger"
	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Log    LogConfig    `mapstructure:"log"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// LogConfig holds logging configuration
type LogConfig struct {
	Level      logger.Level `mapstructure:"level"`
	Format     string       `mapstructure:"format"`
	Output     string       `mapstructure:"output"`
	FilePath   string       `mapstructure:"file_path"`
	MaxSize    int          `mapstructure:"max_size"`
	MaxBackups int          `mapstructure:"max_backups"`
	MaxAge     int          `mapstructure:"max_age"`
	Compress   bool         `mapstructure:"compress"`
	AddCaller  bool         `mapstructure:"add_caller"`
	AddStack   bool         `mapstructure:"add_stack"`
}

// Load reads configuration from file or environment variables
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// Set default values
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")
	
	// Set default logging values
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "json")
	viper.SetDefault("log.output", "stdout")
	viper.SetDefault("log.file_path", "logs/app.log")
	viper.SetDefault("log.max_size", 100)
	viper.SetDefault("log.max_backups", 3)
	viper.SetDefault("log.max_age", 28)
	viper.SetDefault("log.compress", true)
	viper.SetDefault("log.add_caller", true)
	viper.SetDefault("log.add_stack", false)

	// Read environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
