package handler

import (
	"net/http"
	"strconv"
	"postgres-crud/internal/dto"
	"postgres-crud/internal/errors"
	"postgres-crud/service"
	"github.com/gin-gonic/gin"
)

// ProductHandler handles HTTP requests for products
type ProductHandler struct {
	productService service.ProductService
}

// NewProductHandler creates a new instance of ProductHandler
func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateProduct handles POST /api/v1/products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid request body",
			Details: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	product, err := h.productService.CreateProduct(req.Name, req.Description, req.Price, req.Stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to create product",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	})
}

// GetProduct handles GET /api/v1/products/:id
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Invalid product ID",
			Code:  http.StatusBadRequest,
		})
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error: "Product not found",
				Code:  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to fetch product",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	})
}

// ListProducts handles GET /api/v1/products
func (h *ProductHandler) ListProducts(c *gin.Context) {
	// Check if filtering is requested
	var filterReq dto.FilterProductsRequest
	if err := c.ShouldBindQuery(&filterReq); err == nil {
		// If any filter parameter is provided, use filtering
		if filterReq.Name != "" || filterReq.Description != "" || 
		   filterReq.MinPrice != nil || filterReq.MaxPrice != nil ||
		   filterReq.MinStock != nil || filterReq.MaxStock != nil {
			products, err := h.productService.FilterProducts(
				filterReq.Name,
				filterReq.Description,
				filterReq.MinPrice,
				filterReq.MaxPrice,
				filterReq.MinStock,
				filterReq.MaxStock,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
					Error:   "Failed to filter products",
					Details: err.Error(),
					Code:    http.StatusInternalServerError,
				})
				return
			}

			response := make([]dto.ProductResponse, len(products))
			for i, product := range products {
				response[i] = dto.ProductResponse{
					ID:          product.ID,
					Name:        product.Name,
					Description: product.Description,
					Price:       product.Price,
					Stock:       product.Stock,
					CreatedAt:   product.CreatedAt,
					UpdatedAt:   product.UpdatedAt,
				}
			}

			c.JSON(http.StatusOK, dto.ListProductsResponse{
				Products: response,
				Count:    len(response),
			})
			return
		}
	}

	// Default: get all products
	products, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to fetch products",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	response := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		response[i] = dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.ListProductsResponse{
		Products: response,
		Count:    len(response),
	})
}

// UpdateProduct handles PUT /api/v1/products/:id
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Invalid product ID",
			Code:  http.StatusBadRequest,
		})
		return
	}

	var req dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid request body",
			Details: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	product, err := h.productService.UpdateProduct(uint(id), req.Name, req.Description, req.Price, req.Stock)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error: "Product not found",
				Code:  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to update product",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	})
}

// DeleteProduct handles DELETE /api/v1/products/:id
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Invalid product ID",
			Code:  http.StatusBadRequest,
		})
		return
	}

	if err := h.productService.DeleteProduct(uint(id)); err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error: "Product not found",
				Code:  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to delete product",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Product deleted successfully",
	})
}

// AddProductToOrder handles POST /api/v1/orders/:id/products
func (h *ProductHandler) AddProductToOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Invalid order ID",
			Code:  http.StatusBadRequest,
		})
		return
	}

	var req dto.AddProductToOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid request body",
			Details: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	if err := h.productService.AddProductToOrder(uint(orderID), req.ProductID, req.Quantity); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Failed to add product to order",
			Details: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Product added to order successfully",
	})
}

// RemoveProductFromOrder handles DELETE /api/v1/orders/:id/products/:productId
func (h *ProductHandler) RemoveProductFromOrder(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Invalid order ID",
			Code:  http.StatusBadRequest,
		})
		return
	}

	productID, err := strconv.ParseUint(c.Param("productId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Invalid product ID",
			Code:  http.StatusBadRequest,
		})
		return
	}

	if err := h.productService.RemoveProductFromOrder(uint(orderID), uint(productID)); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Failed to remove product from order",
			Details: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Product removed from order successfully",
	})
}

// GetOrderProducts handles GET /api/v1/orders/:id/products
func (h *ProductHandler) GetOrderProducts(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Invalid order ID",
			Code:  http.StatusBadRequest,
		})
		return
	}

	products, err := h.productService.GetOrderProducts(uint(orderID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to fetch order products",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	response := make([]dto.ProductResponse, len(products))
	for i, product := range products {
		response[i] = dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.ListProductsResponse{
		Products: response,
		Count:    len(response),
	})
}


