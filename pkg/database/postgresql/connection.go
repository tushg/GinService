package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

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

// Connection represents a database connection
type Connection interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	Ping(ctx context.Context) error
	BeginTx(ctx context.Context) (Transaction, error)
	IsHealthy(ctx context.Context) (bool, error)
	GetDriverInfo() string
	GetDB() *sql.DB
}

// Transaction represents a database transaction
type Transaction interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	GetUnderlyingTx() interface{}
}

// PostgreSQLConnection implements Connection for PostgreSQL
type PostgreSQLConnection struct {
	db     *sql.DB
	config *DatabaseConfig
}

// NewPostgreSQLConnection creates a new PostgreSQL connection
func NewPostgreSQLConnection(config *DatabaseConfig) *PostgreSQLConnection {
	return &PostgreSQLConnection{
		config: config,
	}
}

// Connect establishes a connection to PostgreSQL
func (p *PostgreSQLConnection) Connect(ctx context.Context) error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.config.Host,
		p.config.Port,
		p.config.Username,
		p.config.Password,
		p.config.Database,
		p.config.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(p.config.MaxConnections)
	db.SetMaxIdleConns(p.config.MaxIdleConnections)
	db.SetConnMaxLifetime(p.config.ConnectionTimeout)

	// Test the connection
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	p.db = db
	return nil
}

// Close closes the database connection
func (p *PostgreSQLConnection) Close(ctx context.Context) error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

// Ping checks if the database is reachable
func (p *PostgreSQLConnection) Ping(ctx context.Context) error {
	if p.db == nil {
		return fmt.Errorf("database connection not established")
	}
	return p.db.PingContext(ctx)
}

// BeginTx starts a new transaction
func (p *PostgreSQLConnection) BeginTx(ctx context.Context) (Transaction, error) {
	if p.db == nil {
		return nil, fmt.Errorf("database connection not established")
	}

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	return &PostgreSQLTransaction{tx: tx}, nil
}

// IsHealthy checks if the database is healthy
func (p *PostgreSQLConnection) IsHealthy(ctx context.Context) (bool, error) {
	if err := p.Ping(ctx); err != nil {
		return false, err
	}

	// Check if we can execute a simple query
	var result int
	err := p.db.QueryRowContext(ctx, "SELECT 1").Scan(&result)
	if err != nil {
		return false, fmt.Errorf("health check query failed: %w", err)
	}

	return result == 1, nil
}

// GetDriverInfo returns information about the database driver
func (p *PostgreSQLConnection) GetDriverInfo() string {
	return "PostgreSQL"
}

// GetDB returns the underlying sql.DB for direct access if needed
func (p *PostgreSQLConnection) GetDB() *sql.DB {
	return p.db
}

// PostgreSQLTransaction implements Transaction for PostgreSQL
type PostgreSQLTransaction struct {
	tx *sql.Tx
}

// Commit commits the transaction
func (pt *PostgreSQLTransaction) Commit(ctx context.Context) error {
	return pt.tx.Commit()
}

// Rollback rolls back the transaction
func (pt *PostgreSQLTransaction) Rollback(ctx context.Context) error {
	return pt.tx.Rollback()
}

// GetUnderlyingTx returns the underlying transaction object
func (pt *PostgreSQLTransaction) GetUnderlyingTx() interface{} {
	return pt.tx
}
