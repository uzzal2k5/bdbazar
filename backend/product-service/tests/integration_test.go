package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"product-service/config"
	"product-service/controllers"
	"product-service/models"
	"product-service/repository"
	"product-service/routes"
	"product-service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var sellerToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsInJvbGUiOiJzZWxsZXIifQ.6yxDxI7r6lFTyx4b0StMWCJvP_POKiEerAUZ4-Fjcsw"

func setupRouter() *gin.Engine {
	_ = godotenv.Load("../.env")

	os.Setenv("DB_SOURCE", "host=localhost user=postgres password=postgres123 dbname=bdbazar_product_test port=55435 sslmode=disable")
	db := config.ConnectDB()
	db.AutoMigrate(&models.Product{})

	repo := repository.NewProductRepository(db)
	svc := services.NewProductService(repo)
	ctrl := controllers.NewProductController(svc)

	r := gin.Default()
	routes.SetupProductRoutes(r, ctrl)
	return r
}

func TestCreateAndFetchProducts(t *testing.T) {
	r := setupRouter()

	product := models.Product{
		Name:        "Integration Test Product",
		Description: "Test desc",
		Price:       49.99,
		Stock:       5,
	}

	payload, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", sellerToken)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d", w.Code)
	}

	// GET /api/products
	getReq, _ := http.NewRequest("GET", "/api/products", nil)
	getResp := httptest.NewRecorder()
	r.ServeHTTP(getResp, getReq)

	if getResp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", getResp.Code)
	}

	// GET /api/products/seller (auth required)
	sellerReq, _ := http.NewRequest("GET", "/api/products/seller", nil)
	sellerReq.Header.Set("Authorization", sellerToken)
	sellerResp := httptest.NewRecorder()
	r.ServeHTTP(sellerResp, sellerReq)

	if sellerResp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", sellerResp.Code)
	}
}
