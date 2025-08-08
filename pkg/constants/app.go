package constants

// Application constants
const (
	// App information
	AppName    = "gin-service"
	AppVersion = "1.0.0"
	AppAuthor  = "Your Name"
	
	// API constants
	APIVersion = "v1"
	APIPrefix  = "/api/" + APIVersion
	
	// HTTP constants
	DefaultPort = "8080"
	DefaultHost = "localhost"
	
	// Database constants
	DefaultDBMaxConnections     = 10
	DefaultDBMaxIdleConnections = 5
	DefaultDBConnectionTimeout  = "30s"
	
	// Pagination constants
	DefaultPageSize  = 10
	MaxPageSize      = 100
	DefaultPage      = 1
	
	// Validation constants
	MaxStringLength = 255
	MinStringLength = 1
	
	// Time constants
	DefaultTimeout = "30s"
	
	// File constants
	MaxFileSize = 10 * 1024 * 1024 // 10MB
	AllowedFileTypes = "jpg,jpeg,png,gif,pdf,doc,docx"
	
	// Cache constants
	DefaultCacheTTL = 300 // 5 minutes in seconds
	
	// Rate limiting constants
	DefaultRateLimit = 100 // requests per minute
	BurstLimit       = 20  // burst requests
)

// Environment constants
const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
	EnvTesting     = "testing"
)

// Database types
const (
	DBTypePostgreSQL = "postgresql"
	DBTypeMySQL      = "mysql"
	DBTypeSQLite     = "sqlite"
)

// Log levels
const (
	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
	LogLevelFatal = "fatal"
)

// Log formats
const (
	LogFormatJSON = "json"
	LogFormatText = "text"
)

// Log outputs
const (
	LogOutputStdout = "stdout"
	LogOutputStderr = "stderr"
	LogOutputFile   = "file"
)

// HTTP methods
const (
	MethodGET    = "GET"
	MethodPOST   = "POST"
	MethodPUT    = "PUT"
	MethodDELETE = "DELETE"
	MethodPATCH  = "PATCH"
)

// Content types
const (
	ContentTypeJSON        = "application/json"
	ContentTypeXML         = "application/xml"
	ContentTypeFormData    = "multipart/form-data"
	ContentTypeURLEncoded  = "application/x-www-form-urlencoded"
	ContentTypeTextPlain   = "text/plain"
	ContentTypeTextHTML    = "text/html"
	ContentTypeOctetStream = "application/octet-stream"
)

// Status messages
const (
	StatusSuccess = "success"
	StatusError   = "error"
	StatusPending = "pending"
	StatusFailed  = "failed"
)

// Health check constants
const (
	HealthStatusHealthy   = "healthy"
	HealthStatusUnhealthy = "unhealthy"
	HealthStatusDegraded  = "degraded"
)

// Error messages
const (
	ErrMsgInternalServer     = "Internal server error"
	ErrMsgBadRequest         = "Bad request"
	ErrMsgNotFound           = "Resource not found"
	ErrMsgUnauthorized       = "Unauthorized"
	ErrMsgForbidden          = "Forbidden"
	ErrMsgValidationFailed   = "Validation failed"
	ErrMsgDatabaseError      = "Database error"
	ErrMsgExternalAPIError   = "External API error"
	ErrMsgRateLimitExceeded  = "Rate limit exceeded"
	ErrMsgTimeout            = "Request timeout"
	ErrMsgConflict           = "Resource conflict"
)

// Success messages
const (
	MsgCreated     = "Resource created successfully"
	MsgUpdated     = "Resource updated successfully"
	MsgDeleted     = "Resource deleted successfully"
	MsgRetrieved   = "Resource retrieved successfully"
	MsgListed      = "Resources listed successfully"
	MsgValidated   = "Validation successful"
	MsgAuthenticated = "Authentication successful"
)

// File paths
const (
	ConfigPath     = "./configs"
	LogsPath       = "./logs"
	UploadsPath    = "./uploads"
	TempPath       = "./temp"
	MigrationsPath = "./migrations"
)

// Headers
const (
	HeaderAuthorization = "Authorization"
	HeaderContentType   = "Content-Type"
	HeaderAccept        = "Accept"
	HeaderUserAgent     = "User-Agent"
	HeaderXRequestID    = "X-Request-ID"
	HeaderXAPIKey       = "X-API-Key"
	HeaderXCorrelationID = "X-Correlation-ID"
)

// Query parameters
const (
	QueryParamPage     = "page"
	QueryParamLimit    = "limit"
	QueryParamSort     = "sort"
	QueryParamOrder    = "order"
	QueryParamSearch   = "search"
	QueryParamFilter   = "filter"
	QueryParamInclude  = "include"
	QueryParamExclude  = "exclude"
	QueryParamFields   = "fields"
	QueryParamExpand   = "expand"
)

// Sort orders
const (
	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"
)

// Time formats
const (
	TimeFormatRFC3339     = "2006-01-02T15:04:05Z07:00"
	TimeFormatISO8601     = "2006-01-02T15:04:05.000Z"
	TimeFormatDate        = "2006-01-02"
	TimeFormatTime        = "15:04:05"
	TimeFormatDateTime    = "2006-01-02 15:04:05"
	TimeFormatHuman       = "January 2, 2006 at 3:04 PM"
	TimeFormatShort       = "Jan 2, 2006"
	TimeFormatTimestamp   = "20060102150405"
)
