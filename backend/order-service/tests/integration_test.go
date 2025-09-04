package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"order-service/config"
	"order-service/controllers"
	"order-service/models"
	"order-service/repository"
	"order-service/routes"
	"order-service/services"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupIntegrationRouter() *gin.Engine {
	_ = godotenv.Load("../.env")

	os.Setenv("DB_SOURCE", "host=localhost user=postgres password=postgres123 dbname=bdbazar_test port=55432 sslmode=disable")
	db := config.ConnectDB()
	db.AutoMigrate(&models.Order{})

	repo := repository.NewOrderRepository(db)
	svc := services.NewOrderService(repo)
	ctrl := controllers.NewOrderController(svc)

	r := gin.Default()
	routes.SetupRoutes(r, ctrl)
	return r
}

func TestIntegration_CreateOrderAndGetBuyerOrders(t *testing.T) {
	r := setupIntegrationRouter()

	token := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSIsInJvbGUiOiJidXllciJ9.x030QCtUNW3kUXLg4m2fctf-6P9FnFqYl3-_6KCAM2Y"

	orderPayload := models.Order{
		ProductID: 1001,
		SellerID:  2,
		Quantity:  3,
	}
	jsonPayload, _ := json.Marshal(orderPayload)

	// POST /api/orders
	req, _ := http.NewRequest("POST", "/api/orders", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d", w.Code)
	}

	// GET /api/orders/buyer
	getReq, _ := http.NewRequest("GET", "/api/orders/buyer", nil)
	getReq.Header.Set("Authorization", token)
	getResp := httptest.NewRecorder()
	r.ServeHTTP(getResp, getReq)

	if getResp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", getResp.Code)
	}
}
