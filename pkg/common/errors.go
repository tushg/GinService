package common

import (
	"fmt"
	"net/http"
)

// ErrorCode represents different types of errors
type ErrorCode string

const (
	// Common error codes
	ErrorCodeValidation     ErrorCode = "VALIDATION_ERROR"
	ErrorCodeNotFound       ErrorCode = "NOT_FOUND"
	ErrorCodeUnauthorized   ErrorCode = "UNAUTHORIZED"
	ErrorCodeForbidden      ErrorCode = "FORBIDDEN"
	ErrorCodeInternal       ErrorCode = "INTERNAL_ERROR"
	ErrorCodeBadRequest     ErrorCode = "BAD_REQUEST"
	ErrorCodeConflict       ErrorCode = "CONFLICT"
	ErrorCodeTimeout        ErrorCode = "TIMEOUT"
	ErrorCodeDatabase       ErrorCode = "DATABASE_ERROR"
	ErrorCodeExternalAPI    ErrorCode = "EXTERNAL_API_ERROR"
	ErrorCodeRateLimit      ErrorCode = "RATE_LIMIT_EXCEEDED"
)

// AppError represents a standardized application error
type AppError struct {
	Code       ErrorCode `json:"code"`
	Message    string    `json:"message"`
	Details    string    `json:"details,omitempty"`
	HTTPStatus int       `json:"-"`
	Err        error     `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%s)", e.Code, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new application error
func NewAppError(code ErrorCode, message string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
	}
}

// NewAppErrorWithDetails creates a new application error with details
func NewAppErrorWithDetails(code ErrorCode, message, details string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		Details:    details,
		HTTPStatus: httpStatus,
	}
}

// NewAppErrorWithErr creates a new application error with underlying error
func NewAppErrorWithErr(code ErrorCode, message string, httpStatus int, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
		Err:        err,
	}
}

// Common error constructors
func NewValidationError(message string) *AppError {
	return NewAppError(ErrorCodeValidation, message, http.StatusBadRequest)
}

func NewValidationErrorWithDetails(message, details string) *AppError {
	return NewAppErrorWithDetails(ErrorCodeValidation, message, details, http.StatusBadRequest)
}

func NewNotFoundError(message string) *AppError {
	return NewAppError(ErrorCodeNotFound, message, http.StatusNotFound)
}

func NewUnauthorizedError(message string) *AppError {
	return NewAppError(ErrorCodeUnauthorized, message, http.StatusUnauthorized)
}

func NewForbiddenError(message string) *AppError {
	return NewAppError(ErrorCodeForbidden, message, http.StatusForbidden)
}

func NewInternalError(message string) *AppError {
	return NewAppError(ErrorCodeInternal, message, http.StatusInternalServerError)
}

func NewInternalErrorWithErr(message string, err error) *AppError {
	return NewAppErrorWithErr(ErrorCodeInternal, message, http.StatusInternalServerError, err)
}

func NewBadRequestError(message string) *AppError {
	return NewAppError(ErrorCodeBadRequest, message, http.StatusBadRequest)
}

func NewConflictError(message string) *AppError {
	return NewAppError(ErrorCodeConflict, message, http.StatusConflict)
}

func NewTimeoutError(message string) *AppError {
	return NewAppError(ErrorCodeTimeout, message, http.StatusRequestTimeout)
}

func NewDatabaseError(message string) *AppError {
	return NewAppError(ErrorCodeDatabase, message, http.StatusInternalServerError)
}

func NewDatabaseErrorWithErr(message string, err error) *AppError {
	return NewAppErrorWithErr(ErrorCodeDatabase, message, http.StatusInternalServerError, err)
}

func NewExternalAPIError(message string) *AppError {
	return NewAppError(ErrorCodeExternalAPI, message, http.StatusBadGateway)
}

func NewRateLimitError(message string) *AppError {
	return NewAppError(ErrorCodeRateLimit, message, http.StatusTooManyRequests)
}

// IsAppError checks if an error is an AppError
func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

// GetAppError extracts AppError from an error
func GetAppError(err error) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	return nil
}

// GetHTTPStatus returns the HTTP status code for an error
func GetHTTPStatus(err error) int {
	if appErr := GetAppError(err); appErr != nil {
		return appErr.HTTPStatus
	}
	return http.StatusInternalServerError
}
