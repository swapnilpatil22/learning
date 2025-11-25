package service

import (
	"fmt"
	"postgres-crud/model"
	"postgres-crud/repository"
)

// ProductService defines the interface for product business logic
type ProductService interface {
	CreateProduct(name, description string, price float64, stock int) (*model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	GetAllProducts() ([]model.Product, error)
	GetProductsByName(pattern string) ([]model.Product, error)
	UpdateProduct(id uint, name, description string, price float64, stock int) (*model.Product, error)
	DeleteProduct(id uint) error
	AddProductToOrder(orderID uint, productID uint, quantity int) error
	RemoveProductFromOrder(orderID uint, productID uint) error
	GetOrderProducts(orderID uint) ([]model.Product, error)
	FilterProducts(name, description string, minPrice, maxPrice *float64, minStock, maxStock *int) ([]model.Product, error)
	GetProductsWithOrders() ([]model.Product, error)
}

// productService implements ProductService interface
type productService struct {
	productRepo repository.ProductRepository
	orderRepo   repository.OrderRepository
}

// NewProductService creates a new instance of ProductService
func NewProductService(productRepo repository.ProductRepository, orderRepo repository.OrderRepository) ProductService {
	return &productService{
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}

// CreateProduct creates a new product
func (s *productService) CreateProduct(name, description string, price float64, stock int) (*model.Product, error) {
	if name == "" {
		return nil, fmt.Errorf("product name cannot be empty")
	}
	if price < 0 {
		return nil, fmt.Errorf("product price cannot be negative")
	}
	if stock < 0 {
		return nil, fmt.Errorf("product stock cannot be negative")
	}

	product := &model.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}

// GetProductByID retrieves a product by its ID
func (s *productService) GetProductByID(id uint) (*model.Product, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid product ID")
	}

	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return product, nil
}

// GetAllProducts retrieves all products
func (s *productService) GetAllProducts() ([]model.Product, error) {
	products, err := s.productRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all products: %w", err)
	}

	return products, nil
}

// GetProductsByName retrieves products matching a name pattern
func (s *productService) GetProductsByName(pattern string) ([]model.Product, error) {
	products, err := s.productRepo.GetByCondition("name LIKE ?", "%"+pattern+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to get products by name: %w", err)
	}

	return products, nil
}

// UpdateProduct updates a product
func (s *productService) UpdateProduct(id uint, name, description string, price float64, stock int) (*model.Product, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid product ID")
	}

	if name == "" {
		return nil, fmt.Errorf("product name cannot be empty")
	}
	if price < 0 {
		return nil, fmt.Errorf("product price cannot be negative")
	}
	if stock < 0 {
		return nil, fmt.Errorf("product stock cannot be negative")
	}

	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("product not found: %w", err)
	}

	product.Name = name
	product.Description = description
	product.Price = price
	product.Stock = stock

	if err := s.productRepo.Update(product); err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return product, nil
}

// DeleteProduct deletes a product by ID
func (s *productService) DeleteProduct(id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid product ID")
	}

	if err := s.productRepo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

// AddProductToOrder adds a product to an order
func (s *productService) AddProductToOrder(orderID uint, productID uint, quantity int) error {
	if orderID == 0 {
		return fmt.Errorf("invalid order ID")
	}
	if productID == 0 {
		return fmt.Errorf("invalid product ID")
	}
	if quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}

	// Verify order exists
	_, err := s.orderRepo.GetByID(orderID)
	if err != nil {
		return fmt.Errorf("order not found: %w", err)
	}

	// Verify product exists and check stock
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		return fmt.Errorf("product not found: %w", err)
	}

	if product.Stock < quantity {
		return fmt.Errorf("insufficient stock: available %d, requested %d", product.Stock, quantity)
	}

	// Add product to order with current price
	if err := s.productRepo.AddProductToOrder(orderID, productID, quantity, product.Price); err != nil {
		return fmt.Errorf("failed to add product to order: %w", err)
	}

	// Update product stock
	product.Stock -= quantity
	if err := s.productRepo.Update(product); err != nil {
		return fmt.Errorf("failed to update product stock: %w", err)
	}

	return nil
}

// RemoveProductFromOrder removes a product from an order
func (s *productService) RemoveProductFromOrder(orderID uint, productID uint) error {
	if orderID == 0 {
		return fmt.Errorf("invalid order ID")
	}
	if productID == 0 {
		return fmt.Errorf("invalid product ID")
	}

	if err := s.productRepo.RemoveProductFromOrder(orderID, productID); err != nil {
		return fmt.Errorf("failed to remove product from order: %w", err)
	}

	return nil
}

// GetOrderProducts retrieves all products for an order
func (s *productService) GetOrderProducts(orderID uint) ([]model.Product, error) {
	if orderID == 0 {
		return nil, fmt.Errorf("invalid order ID")
	}

	products, err := s.productRepo.GetProductsByOrderID(orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order products: %w", err)
	}

	return products, nil
}

// FilterProducts filters products based on multiple criteria
func (s *productService) FilterProducts(name, description string, minPrice, maxPrice *float64, minStock, maxStock *int) ([]model.Product, error) {
	products, err := s.productRepo.FilterProducts(name, description, minPrice, maxPrice, minStock, maxStock)
	if err != nil {
		return nil, fmt.Errorf("failed to filter products: %w", err)
	}
	return products, nil
}

// GetProductsWithOrders retrieves all products with their associated orders
func (s *productService) GetProductsWithOrders() ([]model.Product, error) {
	products, err := s.productRepo.GetProductsWithOrders()
	if err != nil {
		return nil, fmt.Errorf("failed to get products with orders: %w", err)
	}
	return products, nil
}


