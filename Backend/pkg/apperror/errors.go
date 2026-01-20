// pkg/apperror/errors.go
package apperror

import (
	"fmt"
	"net/http"
)

// AppError represents a structured application error
type AppError struct {
	Code       string                 `json:"code"`
	Message    string                 `json:"message"`
	Details    map[string]interface{} `json:"details,omitempty"`
	HTTPStatus int                    `json:"-"`
	Err        error                  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s - %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Common error codes
const (
	// Client errors (4xx)
	ErrCodeValidation     = "VALIDATION_ERROR"
	ErrCodeUnauthorized   = "UNAUTHORIZED"
	ErrCodeForbidden      = "FORBIDDEN"
	ErrCodeNotFound       = "NOT_FOUND"
	ErrCodeConflict       = "CONFLICT"
	ErrCodeRateLimited    = "RATE_LIMITED"
	ErrCodeInvalidRequest = "INVALID_REQUEST"

	// Server errors (5xx)
	ErrCodeInternal    = "INTERNAL_ERROR"
	ErrCodeDatabase    = "DATABASE_ERROR"
	ErrCodeExternal    = "EXTERNAL_SERVICE_ERROR"
	ErrCodeUnavailable = "SERVICE_UNAVAILABLE"
)

// New creates a new AppError
func New(code, message string, status int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: status,
	}
}

// Wrap wraps an existing error with an AppError
func Wrap(err error, code, message string, status int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: status,
		Err:        err,
	}
}

// WithDetails adds details to the error
func (e *AppError) WithDetails(details map[string]interface{}) *AppError {
	e.Details = details
	return e
}

// Common error constructors

// BadRequest returns a 400 error
func BadRequest(message string) *AppError {
	return New(ErrCodeInvalidRequest, message, http.StatusBadRequest)
}

// ValidationError returns a 400 validation error
func ValidationError(message string, details map[string]interface{}) *AppError {
	return New(ErrCodeValidation, message, http.StatusBadRequest).WithDetails(details)
}

// Unauthorized returns a 401 error
func Unauthorized(message string) *AppError {
	if message == "" {
		message = "Authentication required"
	}
	return New(ErrCodeUnauthorized, message, http.StatusUnauthorized)
}

// Forbidden returns a 403 error
func Forbidden(message string) *AppError {
	if message == "" {
		message = "Access denied"
	}
	return New(ErrCodeForbidden, message, http.StatusForbidden)
}

// NotFound returns a 404 error
func NotFound(resource string) *AppError {
	return New(ErrCodeNotFound, fmt.Sprintf("%s not found", resource), http.StatusNotFound)
}

// Conflict returns a 409 error
func Conflict(message string) *AppError {
	return New(ErrCodeConflict, message, http.StatusConflict)
}

// RateLimited returns a 429 error
func RateLimited(message string) *AppError {
	if message == "" {
		message = "Too many requests"
	}
	return New(ErrCodeRateLimited, message, http.StatusTooManyRequests)
}

// InternalError returns a 500 error
func InternalError(err error) *AppError {
	return Wrap(err, ErrCodeInternal, "An internal error occurred", http.StatusInternalServerError)
}

// DatabaseError returns a 500 database error
func DatabaseError(err error) *AppError {
	return Wrap(err, ErrCodeDatabase, "Database operation failed", http.StatusInternalServerError)
}

// ExternalError returns a 502 external service error
func ExternalError(service string, err error) *AppError {
	return Wrap(err, ErrCodeExternal, fmt.Sprintf("External service '%s' failed", service), http.StatusBadGateway)
}

// Unavailable returns a 503 error
func Unavailable(message string) *AppError {
	if message == "" {
		message = "Service temporarily unavailable"
	}
	return New(ErrCodeUnavailable, message, http.StatusServiceUnavailable)
}
