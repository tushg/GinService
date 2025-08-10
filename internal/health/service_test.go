package health

import (
	"context"
	"testing"
	"time"

	"gin-service/pkg/logger"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHealthRepository is a mock implementation of HealthRepository
type MockHealthRepository struct {
	mock.Mock
}

// MockLogger is a mock implementation of logger.Logger
type MockLogger struct{}

func (m *MockLogger) Debug(ctx context.Context, message string, fields logger.Fields)            {}
func (m *MockLogger) Info(ctx context.Context, message string, fields logger.Fields)             {}
func (m *MockLogger) Warn(ctx context.Context, message string, fields logger.Fields)             {}
func (m *MockLogger) Error(ctx context.Context, message string, err error, fields logger.Fields) {}
func (m *MockLogger) Fatal(ctx context.Context, message string, err error, fields logger.Fields) {}

func (m *MockHealthRepository) GetSystemStatus(ctx context.Context) (*SystemStatus, error) {
	args := m.Called(ctx)
	return args.Get(0).(*SystemStatus), args.Error(1)
}

func (m *MockHealthRepository) CheckDatabaseConnection(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockHealthRepository) CheckExternalServices(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestHealthService_GetHealth(t *testing.T) {
	// Arrange
	mockRepo := new(MockHealthRepository)
	mockLogger := new(MockLogger)
	service := NewHealthService(mockRepo, mockLogger)
	ctx := context.Background()

	expectedStatus := &SystemStatus{
		IsHealthy: true,
		Uptime:    time.Now(),
		Version:   "1.0.0",
	}

	mockRepo.On("GetSystemStatus", ctx).Return(expectedStatus, nil)
	mockRepo.On("CheckDatabaseConnection", ctx).Return(nil)
	mockRepo.On("CheckExternalServices", ctx).Return(nil)

	// Act
	response, err := service.GetHealth(ctx)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "healthy", response.Status)
	assert.Equal(t, "gin-service", response.Service)
	assert.Equal(t, "1.0.0", response.Version)
	assert.NotNil(t, response.Details)
	assert.Equal(t, "healthy", response.Details.Database)
	assert.Equal(t, []string{"all services healthy"}, response.Details.ExternalServices)

	mockRepo.AssertExpectations(t)
}

func TestHealthService_GetReadiness(t *testing.T) {
	// Arrange
	mockRepo := new(MockHealthRepository)
	mockLogger := new(MockLogger)
	service := NewHealthService(mockRepo, mockLogger)
	ctx := context.Background()

	expectedStatus := &SystemStatus{
		IsHealthy: true,
		Uptime:    time.Now(),
		Version:   "1.0.0",
	}

	mockRepo.On("GetSystemStatus", ctx).Return(expectedStatus, nil)
	mockRepo.On("CheckDatabaseConnection", ctx).Return(nil)

	// Act
	response, err := service.GetReadiness(ctx)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "ready", response.Status)
	assert.Equal(t, "gin-service", response.Service)
	assert.Equal(t, "1.0.0", response.Version)

	mockRepo.AssertExpectations(t)
}

func TestHealthService_GetLiveness(t *testing.T) {
	// Arrange
	mockRepo := new(MockHealthRepository)
	mockLogger := new(MockLogger)
	service := NewHealthService(mockRepo, mockLogger)
	ctx := context.Background()

	expectedStatus := &SystemStatus{
		IsHealthy: true,
		Uptime:    time.Now(),
		Version:   "1.0.0",
	}

	mockRepo.On("GetSystemStatus", ctx).Return(expectedStatus, nil)

	// Act
	response, err := service.GetLiveness(ctx)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "alive", response.Status)
	assert.Equal(t, "gin-service", response.Service)
	assert.Equal(t, "1.0.0", response.Version)

	mockRepo.AssertExpectations(t)
}
