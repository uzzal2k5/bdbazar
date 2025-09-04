package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"shop-service/config"
	"shop-service/controllers"
	"shop-service/models"
	"shop-service/repository"
	"shop-service/routes"
	"shop-service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var sellerToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMiIsInJvbGUiOiJzZWxsZXIifQ.6yxDxI7r6lFTyx4b0StMWCJvP_POKiEerAUZ4-Fjcsw"

func setupRouter() *gin.Engine {
	_ = godotenv.Load("../.env")

	os.Setenv("DB_SOURCE", "host=localhost user=postgres password=postgres123 dbname=bdbazar_shop_test port=55436 sslmode=disable")
	db := config.ConnectDB()
	db.AutoMigrate(&models.Shop{})

	repo := repository.NewShopRepository(db)
	svc := services.NewShopService(repo)
	ctrl := controllers.NewShopController(svc)

	r := gin.Default()
	routes.SetupShopRoutes(r, ctrl)
	return r
}

func TestCreateAndFetchShops(t *testing.T) {
	r := setupRouter()

	// ✅ Create a shop
	shop := models.Shop{
		Name:        "Test Shop",
		Description: "Best shop for integration test",
	}

	payload, _ := json.Marshal(shop)

	req, _ := http.NewRequest("POST", "/api/shops", bytes.NewBuffer(payload))
	req.Header.Set("Authorization", sellerToken)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("Expected 201 Created, got %d", w.Code)
	}

	// ✅ Public: GET /api/shops
	publicReq, _ := http.NewRequest("GET", "/api/shops", nil)
	publicResp := httptest.NewRecorder()
	r.ServeHTTP(publicResp, publicReq)

	if publicResp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK from public GET /api/shops, got %d", publicResp.Code)
	}

	// ✅ Seller: GET /api/shops/me
	sellerReq, _ := http.NewRequest("GET", "/api/shops/me", nil)
	sellerReq.Header.Set("Authorization", sellerToken)
	sellerResp := httptest.NewRecorder()
	r.ServeHTTP(sellerResp, sellerReq)

	if sellerResp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK from seller GET /api/shops/me, got %d", sellerResp.Code)
	}
}
