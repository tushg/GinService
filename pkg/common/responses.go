package common

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response represents a standardized API response
type Response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     *AppError   `json:"error,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
	Path      string      `json:"path,omitempty"`
	Method    string      `json:"method,omitempty"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Response
	Pagination Pagination `json:"pagination"`
}

// Pagination represents pagination metadata
type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}

// NewResponse creates a new response
func NewResponse(success bool, message string, data interface{}) *Response {
	return &Response{
		Success:   success,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// NewSuccessResponse creates a success response
func NewSuccessResponse(message string, data interface{}) *Response {
	return NewResponse(true, message, data)
}

// NewErrorResponse creates an error response
func NewErrorResponse(err *AppError) *Response {
	return &Response{
		Success:   false,
		Error:     err,
		Timestamp: time.Now(),
	}
}

// NewPaginatedResponse creates a paginated response
func NewPaginatedResponse(message string, data interface{}, pagination Pagination) *PaginatedResponse {
	return &PaginatedResponse{
		Response: Response{
			Success:   true,
			Message:   message,
			Data:      data,
			Timestamp: time.Now(),
		},
		Pagination: pagination,
	}
}

// CalculatePagination calculates pagination metadata
func CalculatePagination(page, limit int, total int64) Pagination {
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	
	return Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
}

// SendSuccess sends a success response
func SendSuccess(c *gin.Context, message string, data interface{}) {
	response := NewSuccessResponse(message, data)
	response.Path = c.Request.URL.Path
	response.Method = c.Request.Method
	
	c.JSON(http.StatusOK, response)
}

// SendCreated sends a created response
func SendCreated(c *gin.Context, message string, data interface{}) {
	response := NewSuccessResponse(message, data)
	response.Path = c.Request.URL.Path
	response.Method = c.Request.Method
	
	c.JSON(http.StatusCreated, response)
}

// SendNoContent sends a no content response
func SendNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// SendError sends an error response
func SendError(c *gin.Context, err *AppError) {
	response := NewErrorResponse(err)
	response.Path = c.Request.URL.Path
	response.Method = c.Request.Method
	
	statusCode := err.HTTPStatus
	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}
	
	c.JSON(statusCode, response)
}

// SendValidationError sends a validation error response
func SendValidationError(c *gin.Context, message string) {
	err := NewValidationError(message)
	SendError(c, err)
}

// SendValidationErrorWithDetails sends a validation error response with details
func SendValidationErrorWithDetails(c *gin.Context, message, details string) {
	err := NewValidationErrorWithDetails(message, details)
	SendError(c, err)
}

// SendNotFound sends a not found error response
func SendNotFound(c *gin.Context, message string) {
	err := NewNotFoundError(message)
	SendError(c, err)
}

// SendUnauthorized sends an unauthorized error response
func SendUnauthorized(c *gin.Context, message string) {
	err := NewUnauthorizedError(message)
	SendError(c, err)
}

// SendForbidden sends a forbidden error response
func SendForbidden(c *gin.Context, message string) {
	err := NewForbiddenError(message)
	SendError(c, err)
}

// SendInternalError sends an internal error response
func SendInternalError(c *gin.Context, message string) {
	err := NewInternalError(message)
	SendError(c, err)
}

// SendInternalErrorWithErr sends an internal error response with underlying error
func SendInternalErrorWithErr(c *gin.Context, message string, underlyingErr error) {
	err := NewInternalErrorWithErr(message, underlyingErr)
	SendError(c, err)
}

// SendBadRequest sends a bad request error response
func SendBadRequest(c *gin.Context, message string) {
	err := NewBadRequestError(message)
	SendError(c, err)
}

// SendConflict sends a conflict error response
func SendConflict(c *gin.Context, message string) {
	err := NewConflictError(message)
	SendError(c, err)
}

// SendDatabaseError sends a database error response
func SendDatabaseError(c *gin.Context, message string) {
	err := NewDatabaseError(message)
	SendError(c, err)
}

// SendDatabaseErrorWithErr sends a database error response with underlying error
func SendDatabaseErrorWithErr(c *gin.Context, message string, underlyingErr error) {
	err := NewDatabaseErrorWithErr(message, underlyingErr)
	SendError(c, err)
}

// SendPaginatedSuccess sends a paginated success response
func SendPaginatedSuccess(c *gin.Context, message string, data interface{}, pagination Pagination) {
	response := NewPaginatedResponse(message, data, pagination)
	response.Path = c.Request.URL.Path
	response.Method = c.Request.Method
	
	c.JSON(http.StatusOK, response)
}
