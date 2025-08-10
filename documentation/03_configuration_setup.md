# Step 3: Configuration Setup

## Overview
Now let's create the configuration management system using Viper. This will handle loading configuration from YAML files and environment variables.

## What We're Creating
- Configuration structures for server, logging, and database
- YAML configuration file
- Configuration loading functions

## Files to Create
1. `internal/config/config.go` - Configuration structures and loading logic
2. `configs/config.yaml` - Configuration values

## Step 3.1: Create Configuration Structures

### File: `internal/config/config.go`
Create this file with the following content:

```go
package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Log      LogConfig      `mapstructure:"log"`
	Database DatabaseConfig `mapstructure:"database"`
}

// ServerConfig represents server configuration
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// LogConfig represents logging configuration
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

// DatabaseConfig represents database configuration
type DatabaseConfig struct {
	Type                string        `mapstructure:"type"`
	Host                string        `mapstructure:"host"`
	Port                int           `mapstructure:"port"`
	Username            string        `mapstructure:"username"`
	Password            string        `mapstructure:"password"`
	Database            string        `mapstructure:"database"`
	SSLMode             string        `mapstructure:"ssl_mode"`
	MaxConnections      int           `mapstructure:"max_connections"`
	MaxIdleConnections  int           `mapstructure:"max_idle_connections"`
	ConnectionTimeout   time.Duration `mapstructure:"connection_timeout"`
}

// Load loads configuration from file and environment variables
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")

	// Set default values
	setDefaults()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Bind environment variables
	bindEnvVars()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// setDefaults sets default configuration values
func setDefaults() {
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "json")
	viper.SetDefault("log.output", "stdout")
	viper.SetDefault("log.max_size", 100)
	viper.SetDefault("log.max_backups", 3)
	viper.SetDefault("log.max_age", 28)
	viper.SetDefault("log.compress", true)
	viper.SetDefault("log.add_caller", true)
	viper.SetDefault("log.add_stack", false)
	viper.SetDefault("database.type", "postgresql")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.ssl_mode", "disable")
	viper.SetDefault("database.max_connections", 10)
	viper.SetDefault("database.max_idle_connections", 5)
	viper.SetDefault("database.connection_timeout", "30s")
}

// bindEnvVars binds environment variables to configuration keys
func bindEnvVars() {
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("server.mode", "SERVER_MODE")
	viper.BindEnv("log.level", "LOG_LEVEL")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.username", "DB_USERNAME")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.database", "DB_NAME")
}
```

## Step 3.2: Create Configuration File

### File: `configs/config.yaml`
Create this file with the following content:

```yaml
server:
  port: "8080"
  mode: "debug"

log:
  level: "info"
  format: "json"
  output: "stdout"
  file_path: "logs/app.log"
  max_size: 100
  max_backups: 3
  max_age: 28
  compress: true
  add_caller: true
  add_stack: false

database:
  type: "postgresql"
  host: "localhost"
  port: 5432
  username: "postgres"
  password: "password"
  database: "gin_service"
  ssl_mode: "disable"
  max_connections: 10
  max_idle_connections: 5
  connection_timeout: "30s"
```

## Action Items
- [ ] Create `internal/config/config.go` file
- [ ] Create `configs/config.yaml` file
- [ ] Verify the files are created correctly
- [ ] Test that the configuration loads without errors

## Testing the Configuration
Create a simple test file to verify configuration loading:

```go
// internal/config/config_test.go
package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	if cfg.Server.Port != "8080" {
		t.Errorf("Expected server port 8080, got %s", cfg.Server.Port)
	}

	if cfg.Log.Level != "info" {
		t.Errorf("Expected log level info, got %s", cfg.Log.Level)
	}
}
```

## Next Steps
Once configuration is set up, we'll move to:
1. Creating the logger system
2. Setting up database connection
3. Creating middleware

**Complete this step and let me know when you're ready to continue!**
