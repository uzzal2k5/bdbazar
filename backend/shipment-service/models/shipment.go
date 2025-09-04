package models

import "time"

type Shipment struct {
    ID           uint      `json:"id" gorm:"primaryKey"`
    OrderID      uint      `json:"order_id" gorm:"index"`
    SellerID     uint      `json:"seller_id" gorm:"index"`
    BuyerID      uint      `json:"buyer_id" gorm:"index"`
    Address      string    `json:"address"`
    Status       string    `json:"status"` // e.g. pending, shipped, delivered, cancelled
    TrackingCode string    `json:"tracking_code,omitempty"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
