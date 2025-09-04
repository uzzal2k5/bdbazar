package services

import (
	"errors"
	"payment-service/models"
	"payment-service/repository"
)

type PaymentService interface {
	Create(payment *models.Payment) error
	GetByBuyer(buyerID uint) ([]models.Payment, error)
	GetBySeller(sellerID uint) ([]models.Payment, error)
	CompletePayment(paymentID uint) error
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(r repository.PaymentRepository) PaymentService {
	return &paymentService{repo: r}
}

// Create initializes a new payment with status "pending"
func (s *paymentService) Create(p *models.Payment) error {
	if p == nil {
		return errors.New("payment object cannot be nil")
	}
	p.Status = models.StatusPending
	return s.repo.Create(p)
}

// GetByBuyer returns payments made by a specific buyer
func (s *paymentService) GetByBuyer(buyerID uint) ([]models.Payment, error) {
	return s.repo.GetByBuyer(buyerID)
}

// GetBySeller returns payments received by a specific seller
func (s *paymentService) GetBySeller(sellerID uint) ([]models.Payment, error) {
	return s.repo.GetBySeller(sellerID)
}

// CompletePayment marks a payment as completed and sets payment_time
func (s *paymentService) CompletePayment(paymentID uint) error {
	return s.repo.UpdateStatus(paymentID, models.StatusCompleted)
}
