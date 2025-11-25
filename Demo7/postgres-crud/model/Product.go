package model

import (
	"time"

	"gorm.io/gorm"
)

// Product represents a product entity in the database
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string         `json:"name" gorm:"type:varchar(255);not null"`
	Description string         `json:"description" gorm:"type:text"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock       int            `json:"stock" gorm:"type:int;default:0"`
	Orders      []Order        `json:"orders,omitempty" gorm:"many2many:order_products;"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName specifies the table name for Product model
func (Product) TableName() string {
	return "products"
}

// OrderProduct represents the join table for Order-Product many-to-many relationship
type OrderProduct struct {
	OrderID   uint      `gorm:"primaryKey"`
	ProductID uint      `gorm:"primaryKey"`
	Quantity  int       `gorm:"type:int;not null;default:1"`
	Price     float64   `gorm:"type:decimal(10,2);not null"` // Price at time of order
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// TableName specifies the table name for OrderProduct model
func (OrderProduct) TableName() string {
	return "order_products"
}


