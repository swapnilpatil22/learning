package model

import (
	"time"

	"gorm.io/gorm"
)

// Order represents an order entity in the database
type Order struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Description string         `json:"description" gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName specifies the table name for Order model
func (Order) TableName() string {
	return "orders"
}
