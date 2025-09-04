package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"payment-service/config"
	"payment-service/controllers"
	"payment-service/models"
	"payment-service/repository"
	"payment-service/routes"
	"payment-service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var sellerToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMyIsInJvbGUiOiJzZWxsZXIifQ.-LN7jo4Hog7cqC4cyXGVUTc8Xx3vDFsvfHHTPj0JUoQ"
var buyerToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsInJvbGUiOiJidXllciJ9.fuXb1iGh2zTsCqE_VKTSmV3KLF_7dgl-j9ay7HK9sYw"

func setupRouter() *gin.Engine {
	_ = godotenv.Load("../.env")

	os.Setenv("DB_SOURCE", "host=localhost user=postgres password=postgres123 dbname=bdbazar_payment_test port=5432 sslmode=disable")
	db := config.ConnectDB()
	db.AutoMigrate(&models.Payment{})

	repo := repository.NewPaymentRepo(db)
	svc := services.NewPaymentService(repo)
	ctrl := controllers.NewPaymentController(svc)

	r := gin.Default()
	routes.Setup(r, ctrl)
	return r
}

func TestPaymentServiceIntegration(t *testing.T) {
	r := setupRouter()

	// 1) Buyer creates a payment
	payment := models.Payment{
		OrderID:  101,
		SellerID: 3,
		Amount:   99.99,
	}

	payload, _ := json.Marshal(payment)
	req, _ := http.NewRequest("POST", "/api/payments", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", buyerToken)
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d, body: %s", resp.Code, resp.Body.String())
	}

	// 2) Buyer fetches payments
	req2, _ := http.NewRequest("GET", "/api/payments/buyer", nil)
	req2.Header.Set("Authorization", buyerToken)
	resp2 := httptest.NewRecorder()
	r.ServeHTTP(resp2, req2)

	if resp2.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp2.Code)
	}

	// 3) Seller fetches received payments
	req3, _ := http.NewRequest("GET", "/api/payments/seller", nil)
	req3.Header.Set("Authorization", sellerToken)
	resp3 := httptest.NewRecorder()
	r.ServeHTTP(resp3, req3)

	if resp3.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp3.Code)
	}

	// 4) Buyer marks payment complete (use ID from created payment)
	var createdPayment models.Payment
	err := json.Unmarshal(resp.Body.Bytes(), &createdPayment)
	if err != nil {
		t.Fatalf("Error unmarshaling payment create response: %v", err)
	}

	req4, _ := http.NewRequest("PUT", "/api/payments/"+string(rune(createdPayment.ID))+"/complete", nil)
	req4.Header.Set("Authorization", buyerToken)
	resp4 := httptest.NewRecorder()
	r.ServeHTTP(resp4, req4)

	if resp4.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK on complete, got %d", resp4.Code)
	}
}
