package dto

import "time"

// CreateOrderRequest represents the request body for creating an order
type CreateOrderRequest struct {
	Description string `json:"description" binding:"required,min=3,max=255"`
}

// UpdateOrderRequest represents the request body for updating an order
type UpdateOrderRequest struct {
	Description string `json:"description" binding:"required,min=3,max=255"`
}

// OrderResponse represents the order data in API responses
type OrderResponse struct {
	ID          uint                `json:"id"`
	Description string             `json:"description"`
	Products    []ProductResponse  `json:"products,omitempty"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// ListOrdersResponse represents the response for listing orders
type ListOrdersResponse struct {
	Orders []OrderResponse `json:"orders"`
	Count  int             `json:"count"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
	Code    int    `json:"code,omitempty"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Message string `json:"message"`
}

// HealthResponse represents a health check response
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

