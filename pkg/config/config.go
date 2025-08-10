package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Log      LogConfig      `mapstructure:"log"`
	Database DatabaseConfig `mapstructure:"database"`
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Type               string        `mapstructure:"type" yaml:"type"`
	Host               string        `mapstructure:"host" yaml:"host"`
	Port               int           `mapstructure:"port" yaml:"port"`
	Username           string        `mapstructure:"username" yaml:"username"`
	Password           string        `mapstructure:"password" yaml:"password"`
	Database           string        `mapstructure:"database" yaml:"database"`
	SSLMode            string        `mapstructure:"ssl_mode" yaml:"ssl_mode"`
	MaxConnections     int           `mapstructure:"max_connections" yaml:"max_connections"`
	MaxIdleConnections int           `mapstructure:"max_idle_connections" yaml:"max_idle_connections"`
	ConnectionTimeout  time.Duration `mapstructure:"connection_timeout" yaml:"connection_timeout"`
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// LogConfig holds logging configuration
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Format     string `mapstructure:"format"`
	Output     string `mapstructure:"output"`
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
	AddCaller  bool   `mapstructure:"add_caller"`
	AddStack   bool   `mapstructure:"add_stack"`
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

	// Set default database values
	viper.SetDefault("database.type", "postgresql")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.username", "postgres")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.database", "gin_service")
	viper.SetDefault("database.ssl_mode", "disable")
	viper.SetDefault("database.max_connections", 10)
	viper.SetDefault("database.max_idle_connections", 5)
	viper.SetDefault("database.connection_timeout", "30s")

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
