package service

import (
	"fmt"
	"postgres-crud/model"
	"postgres-crud/repository"
)

// OrderService defines the interface for order business logic
type OrderService interface {
	CreateOrder(description string) (*model.Order, error)
	GetOrderByID(id int) (*model.Order, error)
	GetAllOrders() ([]model.Order, error)
	GetOrdersByDescription(pattern string) ([]model.Order, error)
	UpdateOrder(id int, description string) (*model.Order, error)
	UpdateOrderDescription(id int, description string) error
	DeleteOrder(id int) error
}

// orderService implements OrderService interface
type orderService struct {
	repo repository.OrderRepository
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

// CreateOrder creates a new order with the given description
func (s *orderService) CreateOrder(description string) (*model.Order, error) {
	if description == "" {
		return nil, fmt.Errorf("description cannot be empty")
	}

	order := &model.Order{
		Description: description,
	}

	if err := s.repo.Create(order); err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	return order, nil
}

// GetOrderByID retrieves an order by its ID
func (s *orderService) GetOrderByID(id int) (*model.Order, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid order ID")
	}

	order, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}

	return order, nil
}

// GetAllOrders retrieves all orders
func (s *orderService) GetAllOrders() ([]model.Order, error) {
	orders, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all orders: %w", err)
	}

	return orders, nil
}

// GetOrdersByDescription retrieves orders matching a description pattern
func (s *orderService) GetOrdersByDescription(pattern string) ([]model.Order, error) {
	orders, err := s.repo.GetByCondition("description LIKE ?", "%"+pattern+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to get orders by description: %w", err)
	}

	return orders, nil
}

// UpdateOrder updates an order's description
func (s *orderService) UpdateOrder(id int, description string) (*model.Order, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid order ID")
	}

	if description == "" {
		return nil, fmt.Errorf("description cannot be empty")
	}

	order, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("order not found: %w", err)
	}

	order.Description = description
	if err := s.repo.Update(order); err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	return order, nil
}

// UpdateOrderDescription updates only the description field of an order
func (s *orderService) UpdateOrderDescription(id int, description string) error {
	if id <= 0 {
		return fmt.Errorf("invalid order ID")
	}

	if description == "" {
		return fmt.Errorf("description cannot be empty")
	}

	if err := s.repo.UpdateField(id, "description", description); err != nil {
		return fmt.Errorf("failed to update order description: %w", err)
	}

	return nil
}

// DeleteOrder deletes an order by ID
func (s *orderService) DeleteOrder(id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid order ID")
	}

	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}

	return nil
}

