package models

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`   // foreign key from product-service
	ProductName string `json:"product_name"` // snapshot (optional)
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"` // price at order time
	Subtotal  float64 `json:"subtotal"`   // calculated = quantity * unit_price
}
