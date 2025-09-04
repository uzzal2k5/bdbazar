package repository

import (
	"payment-service/models"
	"time"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(*models.Payment) error
	GetByBuyer(uint) ([]models.Payment, error)
	GetBySeller(uint) ([]models.Payment, error)
	UpdateStatus(uint, string) error
}

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepo{db}
}

// Create inserts a new payment record
func (r *paymentRepo) Create(p *models.Payment) error {
	return r.db.Create(p).Error
}

// GetByBuyer returns payments made by a buyer
func (r *paymentRepo) GetByBuyer(buyerID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Where("buyer_id = ?", buyerID).Find(&payments).Error
	return payments, err
}

// GetBySeller returns payments received by a seller
func (r *paymentRepo) GetBySeller(sellerID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Where("seller_id = ?", sellerID).Find(&payments).Error
	return payments, err
}

// UpdateStatus changes the status of a payment
func (r *paymentRepo) UpdateStatus(id uint, status string) error {
	updateFields := map[string]interface{}{
		"status": status,
	}
	if status == models.StatusCompleted {
		updateFields["payment_time"] = time.Now()
	}
	return r.db.Model(&models.Payment{}).
		Where("id = ?", id).
		Updates(updateFields).Error
}
