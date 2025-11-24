package errors

import (
	"errors"
	"net/http"
)

// APIError represents an API error with HTTP status code
type APIError struct {
	Code    int
	Message string
	Details string
}

func (e *APIError) Error() string {
	return e.Message
}

// Predefined API errors
var (
	ErrNotFound     = &APIError{Code: http.StatusNotFound, Message: "Resource not found"}
	ErrBadRequest   = &APIError{Code: http.StatusBadRequest, Message: "Invalid request"}
	ErrUnauthorized = &APIError{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrInternal     = &APIError{Code: http.StatusInternalServerError, Message: "Internal server error"}
)

// NewAPIError creates a new API error
func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

// IsNotFound checks if error is a not found error
func IsNotFound(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code == http.StatusNotFound
	}
	return false
}

// IsBadRequest checks if error is a bad request error
func IsBadRequest(err error) bool {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.Code == http.StatusBadRequest
	}
	return false
}

