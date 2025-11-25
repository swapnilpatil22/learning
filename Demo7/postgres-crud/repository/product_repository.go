package repository

import (
	"postgres-crud/database"
	"postgres-crud/model"

	"gorm.io/gorm"
)

// ProductRepository defines the interface for product data operations
type ProductRepository interface {
	Create(product *model.Product) error
	GetByID(id uint) (*model.Product, error)
	GetAll() ([]model.Product, error)
	GetByCondition(condition string, args ...interface{}) ([]model.Product, error)
	Update(product *model.Product) error
	UpdateField(id uint, field string, value interface{}) error
	Delete(id uint) error
	DeleteByModel(product *model.Product) error
	GetProductsByOrderID(orderID uint) ([]model.Product, error)
	AddProductToOrder(orderID uint, productID uint, quantity int, price float64) error
	RemoveProductFromOrder(orderID uint, productID uint) error
	FilterProducts(name, description string, minPrice, maxPrice *float64, minStock, maxStock *int) ([]model.Product, error)
	GetProductsWithOrders() ([]model.Product, error)
}

// productRepository implements ProductRepository interface
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository() ProductRepository {
	return &productRepository{
		db: database.DB,
	}
}

// Create inserts a new product into the database
func (r *productRepository) Create(product *model.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a product by its ID
func (r *productRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAll retrieves all products from the database
func (r *productRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// GetByCondition retrieves products matching a condition
func (r *productRepository) GetByCondition(condition string, args ...interface{}) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Where(condition, args...).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// Update updates an existing product
func (r *productRepository) Update(product *model.Product) error {
	if err := r.db.Save(product).Error; err != nil {
		return err
	}
	return nil
}

// UpdateField updates a specific field of a product
func (r *productRepository) UpdateField(id uint, field string, value interface{}) error {
	if err := r.db.Model(&model.Product{}).Where("id = ?", id).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

// Delete removes a product by ID
func (r *productRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}

// DeleteByModel removes a product using the model instance
func (r *productRepository) DeleteByModel(product *model.Product) error {
	if err := r.db.Delete(product).Error; err != nil {
		return err
	}
	return nil
}

// GetProductsByOrderID retrieves all products associated with an order
func (r *productRepository) GetProductsByOrderID(orderID uint) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Joins("JOIN order_products ON order_products.product_id = products.id").
		Where("order_products.order_id = ?", orderID).
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// AddProductToOrder adds a product to an order
func (r *productRepository) AddProductToOrder(orderID uint, productID uint, quantity int, price float64) error {
	orderProduct := model.OrderProduct{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
	}
	if err := r.db.Create(&orderProduct).Error; err != nil {
		return err
	}
	return nil
}

// RemoveProductFromOrder removes a product from an order
func (r *productRepository) RemoveProductFromOrder(orderID uint, productID uint) error {
	if err := r.db.Where("order_id = ? AND product_id = ?", orderID, productID).
		Delete(&model.OrderProduct{}).Error; err != nil {
		return err
	}
	return nil
}

// FilterProducts retrieves products based on multiple filter criteria
func (r *productRepository) FilterProducts(name, description string, minPrice, maxPrice *float64, minStock, maxStock *int) ([]model.Product, error) {
	var products []model.Product
	query := r.db.Model(&model.Product{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if description != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}
	if minPrice != nil {
		query = query.Where("price >= ?", *minPrice)
	}
	if maxPrice != nil {
		query = query.Where("price <= ?", *maxPrice)
	}
	if minStock != nil {
		query = query.Where("stock >= ?", *minStock)
	}
	if maxStock != nil {
		query = query.Where("stock <= ?", *maxStock)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductsWithOrders retrieves all products with their associated orders
func (r *productRepository) GetProductsWithOrders() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Preload("Orders").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}


