package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Category represents a product category stored in DB
type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"uniqueIndex;not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Product represents the product with foreign key to Category
type Product struct {
	gorm.Model
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Price       float64        `gorm:"not null" json:"price"`
	Quantity    int            `gorm:"not null" json:"quantity"`
	ImageURL    string         `json:"image_url"`
	Stock       int            `json:"stock"`
	CategoryID  uint           `gorm:"not null;index" json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID" json:"category"`
	SellerID    uint           `gorm:"index;not null" json:"seller_id"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate validates that the CategoryID exists before inserting product
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	var cat Category
	if err := tx.First(&cat, p.CategoryID).Error; err != nil {
		return errors.New("invalid category_id: category does not exist")
	}
	return nil
}

// BeforeUpdate validates CategoryID before updating product
func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	var cat Category
	if err := tx.First(&cat, p.CategoryID).Error; err != nil {
		return errors.New("invalid category_id: category does not exist")
	}
	return nil
}
