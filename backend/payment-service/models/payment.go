package models

import (
	"time"

	"gorm.io/gorm"
)

// ================================
// ENUMS: Payment Status
// ================================
const (
	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
	StatusRefunded  = "refunded"
)
// ================================
// ENUMS: Payment Method
// ================================
const (
	MethodBkash  = "bkash"
	MethodNagad  = "nagad"
	MethodCOD    = "cod"
	MethodCard   = "card"
	MethodStripe = "stripe"
)

// ================================
// Payment Model
// ================================
type Payment struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Amount      float64        `gorm:"not null" json:"amount" validate:"required"`
	Status      string         `gorm:"type:varchar(20);default:'pending'" json:"status" validate:"required,oneof=pending completed failed refunded"`
	OrderID     uint           `gorm:"not null" json:"order_id" validate:"required"`
	BuyerID     uint           `gorm:"not null" json:"buyer_id" validate:"required"`
	SellerID    uint           `gorm:"not null" json:"seller_id" validate:"required"`
	Method        string       `gorm:"type:varchar(20);not null" json:"method"` // enum: bkash, cod, stripe, etc.
    TransactionID string       `gorm:"type:varchar(100);uniqueIndex" json:"transaction_id,omitempty"`
	PaymentTime *time.Time     `json:"payment_time,omitempty"`
	Description   string       `gorm:"type:text" json:"description,omitempty"`

	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
