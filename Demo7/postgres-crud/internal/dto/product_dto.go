package dto

import "time"

// CreateProductRequest represents the request body for creating a product
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,min=3,max=255"`
	Description string  `json:"description" binding:"max=1000"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Stock       int     `json:"stock" binding:"min=0"`
}

// UpdateProductRequest represents the request body for updating a product
type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required,min=3,max=255"`
	Description string  `json:"description" binding:"max=1000"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Stock       int     `json:"stock" binding:"min=0"`
}

// ProductResponse represents the product data in API responses
type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ListProductsResponse represents the response for listing products
type ListProductsResponse struct {
	Products []ProductResponse `json:"products"`
	Count    int              `json:"count"`
}

// AddProductToOrderRequest represents the request to add a product to an order
type AddProductToOrderRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

// FilterProductsRequest represents the request for filtering products
type FilterProductsRequest struct {
	Name        string   `form:"name"`
	Description string   `form:"description"`
	MinPrice    *float64 `form:"min_price"`
	MaxPrice    *float64 `form:"max_price"`
	MinStock    *int     `form:"min_stock"`
	MaxStock    *int     `form:"max_stock"`
}

// FilterOrdersRequest represents the request for filtering orders
type FilterOrdersRequest struct {
	Description string `form:"description"`
	ProductID   *uint  `form:"product_id"`
	WithProducts bool  `form:"with_products"`
}


