package main

import (
	"fmt"
	"log"
	"postgres-crud/config"
	"postgres-crud/database"
	"postgres-crud/model"
	"postgres-crud/repository"
	"postgres-crud/service"
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

	// Initialize repository and service layers
	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepo)

	fmt.Println("=== GORM CRUD Operations for Order Table ===\n")

	// CREATE - Insert new orders
	fmt.Println("1. CREATE Operation:")
	order1, err := orderService.CreateOrder("First order - Laptop")
	if err != nil {
		log.Printf("Error creating order1: %v\n", err)
	} else {
		fmt.Printf("Created order with ID: %d, Description: %s\n", order1.ID, order1.Description)
	}

	order2, err := orderService.CreateOrder("Second order - Mouse")
	if err != nil {
		log.Printf("Error creating order2: %v\n", err)
	} else {
		fmt.Printf("Created order with ID: %d, Description: %s\n", order2.ID, order2.Description)
	}

	order3, err := orderService.CreateOrder("Third order - Keyboard")
	if err != nil {
		log.Printf("Error creating order3: %v\n", err)
	} else {
		fmt.Printf("Created order with ID: %d, Description: %s\n\n", order3.ID, order3.Description)
	}

	// READ - Get single order by ID
	fmt.Println("2. READ Operation (Single):")
	foundOrder, err := orderService.GetOrderByID(int(order1.ID))
	if err != nil {
		log.Printf("Error finding order: %v\n", err)
	} else {
		fmt.Printf("Found order - ID: %d, Description: %s\n\n", foundOrder.ID, foundOrder.Description)
	}

	// READ - Get all orders
	fmt.Println("3. READ Operation (All):")
	orders, err := orderService.GetAllOrders()
	if err != nil {
		log.Printf("Error fetching all orders: %v\n", err)
	} else {
		fmt.Printf("Total orders: %d\n", len(orders))
		for _, order := range orders {
			fmt.Printf("  - ID: %d, Description: %s\n", order.ID, order.Description)
		}
		fmt.Println()
	}

	// READ - Get orders with condition
	fmt.Println("4. READ Operation (With Condition):")
	filteredOrders, err := orderService.GetOrdersByDescription("order")
	if err != nil {
		log.Printf("Error fetching filtered orders: %v\n", err)
	} else {
		fmt.Printf("Filtered orders: %d\n", len(filteredOrders))
		for _, order := range filteredOrders {
			fmt.Printf("  - ID: %d, Description: %s\n", order.ID, order.Description)
		}
		fmt.Println()
	}

	// UPDATE - Update an order
	fmt.Println("5. UPDATE Operation:")
	updatedOrder, err := orderService.UpdateOrder(int(order1.ID), "Updated order - Gaming Laptop")
	if err != nil {
		log.Printf("Error updating order: %v\n", err)
	} else {
		fmt.Printf("Updated order - ID: %d, New Description: %s\n\n", updatedOrder.ID, updatedOrder.Description)
	}

	// UPDATE - Update specific field
	fmt.Println("6. UPDATE Operation (Specific Field):")
	if err := orderService.UpdateOrderDescription(int(order2.ID), "Updated - Wireless Mouse"); err != nil {
		log.Printf("Error updating order field: %v\n", err)
	} else {
		fmt.Printf("Updated order ID %d's description\n\n", order2.ID)
	}

	// DELETE - Delete an order
	fmt.Println("7. DELETE Operation:")
	if err := orderService.DeleteOrder(int(order3.ID)); err != nil {
		log.Printf("Error deleting order: %v\n", err)
	} else {
		fmt.Printf("Deleted order with ID: %d\n\n", order3.ID)
	}

	// Verify final state
	fmt.Println("8. Final State - All Remaining Orders:")
	finalOrders, err := orderService.GetAllOrders()
	if err != nil {
		log.Printf("Error fetching final orders: %v\n", err)
	} else {
		fmt.Printf("Remaining orders: %d\n", len(finalOrders))
		for _, order := range finalOrders {
			fmt.Printf("  - ID: %d, Description: %s\n", order.ID, order.Description)
		}
	}

	fmt.Println("\n=== CRUD Operations Completed ===")
}
