package router

import (
	"postgres-crud/internal/handler"
	"postgres-crud/internal/middleware"
	"postgres-crud/repository"
	"postgres-crud/service"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures and returns the Gin router
func SetupRouter() *gin.Engine {
	// Initialize dependencies
	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(productRepo, orderRepo)
	productHandler := handler.NewProductHandler(productService)

	// Create router
	r := gin.Default()

	// Global middleware
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API routes
	api := r.Group("/api/v1")
	{
		// Order routes
		orders := api.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrder)
			orders.GET("", orderHandler.ListOrders)
			orders.GET("/:id", orderHandler.GetOrder)
			orders.PUT("/:id", orderHandler.UpdateOrder)
			orders.DELETE("/:id", orderHandler.DeleteOrder)
			
			// Order-Product relationship routes
			orders.POST("/:id/products", productHandler.AddProductToOrder)
			orders.GET("/:id/products", productHandler.GetOrderProducts)
			orders.DELETE("/:id/products/:productId", productHandler.RemoveProductFromOrder)
		}

		// Product routes
		products := api.Group("/products")
		{
			products.POST("", productHandler.CreateProduct)
			products.GET("", productHandler.ListProducts)
			products.GET("/:id", productHandler.GetProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
			
			// Get orders containing a specific product
			products.GET("/:id/orders", orderHandler.GetOrdersByProduct)
		}
	}

	return r
}

