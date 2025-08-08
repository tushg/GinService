package database

import (
	"context"
	"fmt"

	"gin-service/internal/config"
	"gin-service/internal/database/postgresql"
)

// Manager manages database connections and repositories
type Manager struct {
	config *config.DatabaseConfig
	conn   postgresql.Connection
}

// NewManager creates a new database manager
func NewManager(config *config.DatabaseConfig) (*Manager, error) {
	var conn postgresql.Connection

	switch config.Type {
	case "postgresql":
		conn = postgresql.NewPostgreSQLConnection(&postgresql.DatabaseConfig{
			Type:               config.Type,
			Host:               config.Host,
			Port:               config.Port,
			Username:           config.Username,
			Password:           config.Password,
			Database:           config.Database,
			SSLMode:            config.SSLMode,
			MaxConnections:     config.MaxConnections,
			MaxIdleConnections: config.MaxIdleConnections,
			ConnectionTimeout:  config.ConnectionTimeout,
		})
	default:
		return nil, fmt.Errorf("unsupported database type: %s", config.Type)
	}

	return &Manager{
		config: config,
		conn:   conn,
	}, nil
}

// Connect establishes a connection to the database
func (m *Manager) Connect(ctx context.Context) error {
	return m.conn.Connect(ctx)
}

// Close closes the database connection
func (m *Manager) Close(ctx context.Context) error {
	return m.conn.Close(ctx)
}

// GetConnection returns the database connection
func (m *Manager) GetConnection() postgresql.Connection {
	return m.conn
}

// GetRepository returns a repository for the specified table
func (m *Manager) GetRepository(tableName string) postgresql.Repository {
	pgConn := m.conn.(*postgresql.PostgreSQLConnection)
	return postgresql.NewPostgreSQLRepository(pgConn, tableName)
}

// IsHealthy checks if the database is healthy
func (m *Manager) IsHealthy(ctx context.Context) (bool, error) {
	return m.conn.IsHealthy(ctx)
}

// Migrate runs database migrations
func (m *Manager) Migrate(ctx context.Context) error {
	// This would contain migration logic
	// For now, just return nil
	return nil
}

// GetDatabaseType returns the database type
func (m *Manager) GetDatabaseType() string {
	return m.config.Type
}
