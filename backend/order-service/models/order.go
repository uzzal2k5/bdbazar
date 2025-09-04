package models

import (
	"time"
)

type Order struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	BuyerID     uint           `json:"buyer_id"`
	ShopID      uint           `json:"shop_id"` // if ordering from specific vendor/shop
	Status      string         `json:"status"` // e.g., "pending", "paid", "shipped", "cancelled"
	TotalAmount float64        `json:"total_amount"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	OrderItems  []OrderItem    `json:"order_items" gorm:"foreignKey:OrderID"`
}
