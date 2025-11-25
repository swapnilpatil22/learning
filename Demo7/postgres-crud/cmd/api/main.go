package main

import (
	"fmt"
	"log"
	"postgres-crud/config"
	"postgres-crud/database"
	"postgres-crud/internal/router"
	"postgres-crud/model"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ensure database connection is closed when program exits
	defer func() {
		if err := database.Close(); err != nil {
			log.Println("Error closing database:", err)
		}
	}()

	// Run database migrations
	// Migrate OrderProduct first (join table), then Order and Product
	// This ensures the join table exists before the many-to-many relationships are set up
	if err := database.Migrate(&model.OrderProduct{}, &model.Order{}, &model.Product{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Setup router
	r := router.SetupRouter()

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("🚀 Server starting on http://%s", addr)
	log.Printf("📋 Health check: http://%s/health", addr)
	log.Printf("📦 Order API endpoints: http://%s/api/v1/orders", addr)
	log.Printf("🛍️  Product API endpoints: http://%s/api/v1/products", addr)

	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
