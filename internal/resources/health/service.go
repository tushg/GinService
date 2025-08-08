package health

import (
	"context"
	"time"

	"gin-service/internal/logger"
)

// healthService implements HealthService interface
type healthService struct {
	repository HealthRepository
	logger     logger.Logger
}

// NewHealthService creates a new health service instance
func NewHealthService(repository HealthRepository, log logger.Logger) HealthService {
	return &healthService{
		repository: repository,
		logger:     log,
	}
}

// GetHealth handles general health check business logic
func (s *healthService) GetHealth(ctx context.Context) (*HealthResponse, error) {
	s.logger.Debug(ctx, "Health check requested", logger.Fields{})
	
	// Get system status from repository
	systemStatus, err := s.repository.GetSystemStatus(ctx)
	if err != nil {
		s.logger.Error(ctx, "Failed to get system status", err, logger.Fields{})
		return nil, err
	}

	// Check database connection
	if err := s.repository.CheckDatabaseConnection(ctx); err != nil {
		s.logger.Warn(ctx, "Database connection failed", logger.Fields{
			"error": err.Error(),
		})
		return &HealthResponse{
			Status:    "unhealthy",
			Timestamp: time.Now(),
			Service:   "gin-service",
			Version:   systemStatus.Version,
			Details: &HealthDetails{
				Database: "unavailable",
			},
		}, nil
	}

	// Check external services
	if err := s.repository.CheckExternalServices(ctx); err != nil {
		s.logger.Warn(ctx, "External services check failed", logger.Fields{
			"error": err.Error(),
		})
		return &HealthResponse{
			Status:    "degraded",
			Timestamp: time.Now(),
			Service:   "gin-service",
			Version:   systemStatus.Version,
			Details: &HealthDetails{
				Database:         "healthy",
				ExternalServices: []string{"some services unavailable"},
			},
		}, nil
	}

	return &HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "gin-service",
		Version:   systemStatus.Version,
		Details: &HealthDetails{
			Database:         "healthy",
			ExternalServices: []string{"all services healthy"},
			Uptime:           time.Since(systemStatus.Uptime).String(),
		},
	}, nil
}

// GetReadiness handles readiness probe business logic
func (s *healthService) GetReadiness(ctx context.Context) (*HealthResponse, error) {
	// For readiness, we check if the service is ready to accept traffic
	systemStatus, err := s.repository.GetSystemStatus(ctx)
	if err != nil {
		return nil, err
	}

	// Check database connection (critical for readiness)
	if err := s.repository.CheckDatabaseConnection(ctx); err != nil {
		return &HealthResponse{
			Status:    "not ready",
			Timestamp: time.Now(),
			Service:   "gin-service",
			Version:   systemStatus.Version,
		}, nil
	}

	return &HealthResponse{
		Status:    "ready",
		Timestamp: time.Now(),
		Service:   "gin-service",
		Version:   systemStatus.Version,
	}, nil
}

// GetLiveness handles liveness probe business logic
func (s *healthService) GetLiveness(ctx context.Context) (*HealthResponse, error) {
	// For liveness, we just check if the service is running
	systemStatus, err := s.repository.GetSystemStatus(ctx)
	if err != nil {
		return nil, err
	}

	return &HealthResponse{
		Status:    "alive",
		Timestamp: time.Now(),
		Service:   "gin-service",
		Version:   systemStatus.Version,
	}, nil
}
