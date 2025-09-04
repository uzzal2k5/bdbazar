package seed

import (
	"log"
	"math/rand"
	"payment-service/models"
	"time"
	"strconv"

	"gorm.io/gorm"
)

func RandomSeedPayments(db *gorm.DB) error {
	now := time.Now()

	methods := []string{
		models.MethodBkash,
		models.MethodNagad,
		models.MethodCOD,
		models.MethodCard,
		models.MethodStripe,
	}

	statuses := []string{
		models.StatusPending,
		models.StatusCompleted,
		models.StatusFailed,
		models.StatusRefunded,
	}

	for i := 1; i <= 100; i++ {
		payment := models.Payment{
			Amount:        float64(rand.Intn(1000) + 100), // 100-1099
			Status:        statuses[rand.Intn(len(statuses))],
			OrderID:       uint(1000 + i),
			BuyerID:       uint(rand.Intn(10) + 1), // 1-10
			SellerID:      uint(rand.Intn(5) + 1),  // 1-5
			Method:        methods[rand.Intn(len(methods))],
			TransactionID: generateTransactionID(i),
			Description:   "Seeded payment for testing",
		}

		if payment.Status == models.StatusCompleted {
			payment.PaymentTime = &now
		}

		if err := db.Create(&payment).Error; err != nil {
			log.Printf("Failed to seed payment %d: %v", i, err)
		}
	}

	log.Println("✅ Seeded 100 payments successfully.")
	return nil
}

func generateTransactionID(index int) string {
	prefixes := []string{"TXN-BK", "TXN-NG", "TXN-COD", "TXN-CD", "TXN-ST"}
	return prefixes[rand.Intn(len(prefixes))] + "-" + time.Now().Format("20060102150405") + "-" + strconv.Itoa(index)
}
