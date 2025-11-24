package handler

import (
	"net/http"
	"strconv"
	"postgres-crud/internal/dto"
	"postgres-crud/internal/errors"
	"postgres-crud/service"
	"github.com/gin-gonic/gin"
)

// OrderHandler handles HTTP requests for orders
type OrderHandler struct {
	orderService service.OrderService
}

// NewOrderHandler creates a new instance of OrderHandler
func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

// CreateOrder handles POST /api/v1/orders
// @Summary Create a new order
// @Description Create a new order with description
// @Tags orders
// @Accept json
// @Produce json
// @Param order body dto.CreateOrderRequest true "Order data"
// @Success 201 {object} dto.OrderResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid request body",
			Details: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	order, err := h.orderService.CreateOrder(req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to create order",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, dto.OrderResponse{
		ID:          order.ID,
		Description: order.Description,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	})
}

// GetOrder handles GET /api/v1/orders/:id
// @Summary Get an order by ID
// @Description Get order details by ID
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} dto.OrderResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/orders/{id} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Invalid order ID",
			Code:  http.StatusBadRequest,
		})
		return
	}

	order, err := h.orderService.GetOrderByID(uint(id))
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error: "Order not found",
				Code:  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to fetch order",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.OrderResponse{
		ID:          order.ID,
		Description: order.Description,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	})
}

// ListOrders handles GET /api/v1/orders
// @Summary List all orders
// @Description Get a list of all orders
// @Tags orders
// @Produce json
// @Success 200 {object} dto.ListOrdersResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/v1/orders [get]
func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to fetch orders",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	response := make([]dto.OrderResponse, len(orders))
	for i, order := range orders {
		response[i] = dto.OrderResponse{
			ID:          order.ID,
			Description: order.Description,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, dto.ListOrdersResponse{
		Orders: response,
		Count:  len(response),
	})
}

// UpdateOrder handles PUT /api/v1/orders/:id
// @Summary Update an order
// @Description Update order details by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body dto.UpdateOrderRequest true "Order data"
// @Success 200 {object} dto.OrderResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Invalid order ID",
			Code:  http.StatusBadRequest,
		})
		return
	}

	var req dto.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "Invalid request body",
			Details: err.Error(),
			Code:    http.StatusBadRequest,
		})
		return
	}

	order, err := h.orderService.UpdateOrder(uint(id), req.Description)
	if err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error: "Order not found",
				Code:  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to update order",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.OrderResponse{
		ID:          order.ID,
		Description: order.Description,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	})
}

// DeleteOrder handles DELETE /api/v1/orders/:id
// @Summary Delete an order
// @Description Delete an order by ID
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} dto.SuccessResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "Invalid order ID",
			Code:  http.StatusBadRequest,
		})
		return
	}

	if err := h.orderService.DeleteOrder(uint(id)); err != nil {
		if errors.IsNotFound(err) {
			c.JSON(http.StatusNotFound, dto.ErrorResponse{
				Error: "Order not found",
				Code:  http.StatusNotFound,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "Failed to delete order",
			Details: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse{
		Message: "Order deleted successfully",
	})
}

