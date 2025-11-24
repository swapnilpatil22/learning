package repository

import (
	"postgres-crud/database"
	"postgres-crud/model"

	"gorm.io/gorm"
)

// OrderRepository defines the interface for order data operations
type OrderRepository interface {
	Create(order *model.Order) error
	GetByID(id int) (*model.Order, error)
	GetAll() ([]model.Order, error)
	GetByCondition(condition string, args ...interface{}) ([]model.Order, error)
	Update(order *model.Order) error
	UpdateField(id int, field string, value interface{}) error
	Delete(id int) error
	DeleteByModel(order *model.Order) error
}

// orderRepository implements OrderRepository interface
type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository() OrderRepository {
	return &orderRepository{
		db: database.DB,
	}
}

// Create inserts a new order into the database
func (r *orderRepository) Create(order *model.Order) error {
	if err := r.db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves an order by its ID
func (r *orderRepository) GetByID(id int) (*model.Order, error) {
	var order model.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// GetAll retrieves all orders from the database
func (r *orderRepository) GetAll() ([]model.Order, error) {
	var orders []model.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// GetByCondition retrieves orders matching a condition
func (r *orderRepository) GetByCondition(condition string, args ...interface{}) ([]model.Order, error) {
	var orders []model.Order
	if err := r.db.Where(condition, args...).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// Update updates an existing order
func (r *orderRepository) Update(order *model.Order) error {
	if err := r.db.Save(order).Error; err != nil {
		return err
	}
	return nil
}

// UpdateField updates a specific field of an order
func (r *orderRepository) UpdateField(id int, field string, value interface{}) error {
	if err := r.db.Model(&model.Order{}).Where("id = ?", id).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

// Delete removes an order by ID
func (r *orderRepository) Delete(id int) error {
	if err := r.db.Delete(&model.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}

// DeleteByModel removes an order using the model instance
func (r *orderRepository) DeleteByModel(order *model.Order) error {
	if err := r.db.Delete(order).Error; err != nil {
		return err
	}
	return nil
}

