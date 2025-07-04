// ------------------------------
// models/product.go
// ------------------------------

package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	VendorID    uint    `json:"vendor_id"`
}