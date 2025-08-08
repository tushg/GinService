package health

import "context"

// HealthService defines the interface for health business logic
type HealthService interface {
	GetHealth(ctx context.Context) (*HealthResponse, error)
	GetReadiness(ctx context.Context) (*HealthResponse, error)
	GetLiveness(ctx context.Context) (*HealthResponse, error)
}

// HealthRepository defines the interface for health data access
type HealthRepository interface {
	GetSystemStatus(ctx context.Context) (*SystemStatus, error)
	CheckDatabaseConnection(ctx context.Context) error
	CheckExternalServices(ctx context.Context) error
}
