package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"shipment-service/config"
	"shipment-service/controllers"
	"shipment-service/models"
	"shipment-service/repository"
	"shipment-service/routes"
	"shipment-service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Example tokens - replace with valid tokens from your auth-service JWT secret
var sellerToken = "Bearer <SELLER_JWT_TOKEN>"
var buyerToken = "Bearer <BUYER_JWT_TOKEN>"

func setupRouter() *gin.Engine {
	_ = godotenv.Load("../.env") // load .env or set env vars externally

	os.Setenv("DB_SOURCE", "host=localhost user=postgres password=postgres123 dbname=bdbazar_shipping_test port=5432 sslmode=disable")
	db := config.ConnectDB()
	db.AutoMigrate(&models.Shipment{})

	repo := repository.NewShipmentRepo(db)
	svc := services.NewShipmentService(repo)
	ctrl := controllers.NewShipmentController(svc)

	r := gin.Default()
	routes.SetupRoutes(r, ctrl)

	return r
}

func TestShippingServiceIntegration(t *testing.T) {
	router := setupRouter()

	// 1) Seller creates a shipment
	shipment := models.Shipment{
		OrderID:  1001,
		BuyerID:  2001,
		Address:  "123 BDBazar St, Dhaka",
		Status:   "pending",
	}
	payload, _ := json.Marshal(shipment)
	req1, _ := http.NewRequest("POST", "/api/shipments", bytes.NewBuffer(payload))
	req1.Header.Set("Authorization", sellerToken)
	req1.Header.Set("Content-Type", "application/json")
	resp1 := httptest.NewRecorder()
	router.ServeHTTP(resp1, req1)

	if resp1.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d: %s", resp1.Code, resp1.Body.String())
	}

	var createdShipment models.Shipment
	json.Unmarshal(resp1.Body.Bytes(), &createdShipment)

	// 2) Seller fetches all shipments
	req2, _ := http.NewRequest("GET", "/api/shipments/seller", nil)
	req2.Header.Set("Authorization", sellerToken)
	resp2 := httptest.NewRecorder()
	router.ServeHTTP(resp2, req2)

	if resp2.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp2.Code)
	}

	// 3) Buyer tries to get shipment by order ID
	url := "/api/shipments/order/" + string(rune(createdShipment.OrderID))
	req3, _ := http.NewRequest("GET", url, nil)
	req3.Header.Set("Authorization", buyerToken)
	resp3 := httptest.NewRecorder()
	router.ServeHTTP(resp3, req3)

	if resp3.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK for buyer fetching shipment, got %d", resp3.Code)
	}

	// 4) Seller updates shipment status
	updatePayload := []byte(`{"status":"shipped"}`)
	urlUpdate := "/api/shipments/" + string(rune(createdShipment.ID)) + "/status"
	req4, _ := http.NewRequest("PUT", urlUpdate, bytes.NewBuffer(updatePayload))
	req4.Header.Set("Authorization", sellerToken)
	req4.Header.Set("Content-Type", "application/json")
	resp4 := httptest.NewRecorder()
	router.ServeHTTP(resp4, req4)

	if resp4.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK for status update, got %d", resp4.Code)
	}

	// 5) Seller deletes shipment
	urlDelete := "/api/shipments/" + string(rune(createdShipment.ID))
	req5, _ := http.NewRequest("DELETE", urlDelete, nil)
	req5.Header.Set("Authorization", sellerToken)
	resp5 := httptest.NewRecorder()
	router.ServeHTTP(resp5, req5)

	if resp5.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK for delete, got %d", resp5.Code)
	}
}
