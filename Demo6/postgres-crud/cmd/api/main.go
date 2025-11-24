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
	if err := database.Migrate(&model.Order{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Setup router
	r := router.SetupRouter()

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("🚀 Server starting on http://%s", addr)
	log.Printf("📋 Health check: http://%s/health", addr)
	log.Printf("📦 API endpoints: http://%s/api/v1/orders", addr)

	if err := r.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

