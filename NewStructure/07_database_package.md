# Step 7: Database Package Setup

## Create Database Manager
Create `pkg/database/manager.go`:

```go
package database

import (
	"context"
	"fmt"
	"time"

	"gin-service/pkg/config"
	"gin-service/pkg/database/postgresql"
	"gin-service/pkg/logger"
)

// Manager manages database connections
type Manager struct {
	config  *config.Config
	logger  logger.Logger
	postgres *postgresql.Connection
}

// NewManager creates a new database manager
func NewManager(cfg *config.Config, log logger.Logger) (*Manager, error) {
	manager := &Manager{
		config: cfg,
		logger: log,
	}

	// Initialize PostgreSQL connection
	postgres, err := postgresql.NewConnection(cfg.Database, log)
	if err != nil {
		return nil, fmt.Errorf("failed to create PostgreSQL connection: %w", err)
	}
	manager.postgres = postgres

	return manager, nil
}

// GetPostgres returns PostgreSQL connection
func (m *Manager) GetPostgres() *postgresql.Connection {
	return m.postgres
}

// Close closes all database connections
func (m *Manager) Close() error {
	var errors []error

	if m.postgres != nil {
		if err := m.postgres.Close(); err != nil {
			errors = append(errors, fmt.Errorf("failed to close PostgreSQL: %w", err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("database manager close errors: %v", errors)
	}

	return nil
}

// HealthCheck performs health check on all databases
func (m *Manager) HealthCheck(ctx context.Context) error {
	var errors []error

	// Check PostgreSQL
	if m.postgres != nil {
		if err := m.postgres.HealthCheck(ctx); err != nil {
			errors = append(errors, fmt.Errorf("PostgreSQL health check failed: %w", err))
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("database health check errors: %v", errors)
	}

	return nil
}
```

## Create PostgreSQL Connection
Create `pkg/database/postgresql/connection.go`:

```go
package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gin-service/pkg/config"
	"gin-service/pkg/logger"
	_ "github.com/lib/pq"
)

// Connection represents a PostgreSQL connection
type Connection struct {
	db     *sql.DB
	config config.DatabaseConfig
	logger logger.Logger
}

// NewConnection creates a new PostgreSQL connection
func NewConnection(cfg config.DatabaseConfig, log logger.Logger) (*Connection, error) {
	dsn := cfg.GetDSN()
	
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	conn := &Connection{
		db:     db,
		config: cfg,
		logger: log,
	}

	log.Info(ctx, "PostgreSQL connection established", logger.Fields{
		"host": cfg.Host,
		"port": cfg.Port,
		"name": cfg.Name,
	})

	return conn, nil
}

// GetDB returns the underlying sql.DB
func (c *Connection) GetDB() *sql.DB {
	return c.db
}

// Close closes the database connection
func (c *Connection) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

// HealthCheck performs a health check on the database
func (c *Connection) HealthCheck(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

// Begin starts a new transaction
func (c *Connection) Begin() (*sql.Tx, error) {
	return c.db.Begin()
}

// BeginTx starts a new transaction with context
func (c *Connection) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return c.db.BeginTx(ctx, opts)
}

// Exec executes a query without returning rows
func (c *Connection) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.db.Exec(query, args...)
}

// ExecContext executes a query without returning rows with context
func (c *Connection) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.db.ExecContext(ctx, query, args...)
}

// Query executes a query that returns rows
func (c *Connection) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.Query(query, args...)
}

// QueryContext executes a query that returns rows with context
func (c *Connection) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.QueryContext(ctx, query, args...)
}

// QueryRow executes a query that returns a single row
func (c *Connection) QueryRow(query string, args ...interface{}) *sql.Row {
	return c.db.QueryRow(query, args...)
}

// QueryRowContext executes a query that returns a single row with context
func (c *Connection) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return c.db.QueryRowContext(ctx, query, args...)
}
```

## Create PostgreSQL Repository Base
Create `pkg/database/postgresql/repository.go`:

```go
package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"gin-service/pkg/logger"
)

// Repository provides common database operations
type Repository struct {
	db     *Connection
	logger logger.Logger
}

// NewRepository creates a new repository
func NewRepository(db *Connection, log logger.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: log,
	}
}

// GetDB returns the database connection
func (r *Repository) GetDB() *Connection {
	return r.db
}

// WithTransaction executes a function within a transaction
func (r *Repository) WithTransaction(ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction failed: %v, rollback failed: %w", err, rbErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// Execute executes a query and returns the result
func (r *Repository) Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Error(ctx, "Failed to execute query", err, logger.Fields{
			"query": query,
			"args":  args,
		})
		return nil, err
	}

	return result, nil
}

// Query executes a query and returns rows
func (r *Repository) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		r.logger.Error(ctx, "Failed to execute query", err, logger.Fields{
			"query": query,
			"args":  args,
		})
		return nil, err
	}

	return rows, nil
}

// QueryRow executes a query and returns a single row
func (r *Repository) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return r.db.QueryRowContext(ctx, query, args...)
}

// Close closes the repository
func (r *Repository) Close() error {
	return r.db.Close()
}
```

## Verify Database Package
```bash
# Check if all database files are created
ls -la pkg/database/
ls -la pkg/database/postgresql/

# Expected output should show:
# pkg/database/
#   manager.go
# pkg/database/postgresql/
#   connection.go
#   repository.go
```

## Next Steps
After creating the database package, proceed to the next file to create the middleware package.
