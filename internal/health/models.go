package health

import "time"

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	Details   *HealthDetails `json:"details,omitempty"`
}

// HealthDetails contains detailed health information
type HealthDetails struct {
	Database   string `json:"database,omitempty"`
	ExternalServices []string `json:"external_services,omitempty"`
	Uptime     string `json:"uptime,omitempty"`
}

// SystemStatus represents the overall system status
type SystemStatus struct {
	IsHealthy bool      `json:"is_healthy"`
	Uptime    time.Time `json:"uptime"`
	Version   string    `json:"version"`
}

// HealthRequest represents health check request (for future use)
type HealthRequest struct {
	IncludeDetails bool `json:"include_details"`
}
