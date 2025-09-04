package seed

import (
	"log"
	"payment-service/models"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedPayments(db *gorm.DB) error {
	now := time.Now()

	payments := []models.Payment{
		{
			Amount:        500.00,
			Status:        models.StatusCompleted,
			OrderID:       101,
			BuyerID:       1,
			SellerID:      2,
			Method:        models.MethodBkash,
			TransactionID: "TXN-BK-001",
			PaymentTime:   &now,
			Description:   "Bkash payment for order 101",
		},
		{
			Amount:        1200.50,
			Status:        models.StatusPending,
			OrderID:       102,
			BuyerID:       2,
			SellerID:      3,
			Method:        models.MethodCOD,
			TransactionID: "TXN-COD-002",
			Description:   "Cash on delivery for order 102",
		},
		{
			Amount:        299.99,
			Status:        models.StatusRefunded,
			OrderID:       103,
			BuyerID:       3,
			SellerID:      1,
			Method:        models.MethodStripe,
			TransactionID: "TXN-STRIPE-003",
			PaymentTime:   &now,
			Description:   "Stripe payment refunded for order 103",
		},
	}

	for i, payment := range payments {
		err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&payment).Error
		if err != nil {
			log.Printf("❌ Failed to seed payment #%d (TXN: %s): %v", i+1, payment.TransactionID, err)
		}
	}

	return nil
}
