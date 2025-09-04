package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"admin-service/config"
	"admin-service/controllers"
	"admin-service/models"
	"admin-service/repository"
	"admin-service/routes"
	"admin-service/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	_ = godotenv.Load("../.env")

	os.Setenv("DB_SOURCE", "host=localhost user=postgres password=postgres123 dbname=bdbazar_admin_test port=55434 sslmode=disable")
	db := config.ConnectDB()
	db.AutoMigrate(&models.User{}, &models.Shop{}, &models.Order{})

	repo := repository.NewAdminRepository(db)
	svc := services.NewAdminService(repo)
	ctrl := controllers.NewAdminController(svc)

	r := gin.Default()
	routes.SetupAdminRoutes(r, ctrl)
	return r
}

var adminToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMSIsInJvbGUiOiJhZG1pbiJ9.2TP-KLrTnNk94JQ3SijvMGc6EReWyw8Lu3OyPGXYcVI"

func TestGetPlatformMetrics(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/api/admin/metrics", nil)
	req.Header.Set("Authorization", adminToken)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp.Code)
	}
}

func TestBlockUser(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("PUT", "/api/admin/users/1/block", nil)
	req.Header.Set("Authorization", adminToken)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp.Code)
	}
}

func TestApproveShop(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("PUT", "/api/admin/shops/1/approve", nil)
	req.Header.Set("Authorization", adminToken)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp.Code)
	}
}

func TestResetPassword(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("PUT", "/api/admin/users/1/reset-password", nil)
	req.Header.Set("Authorization", adminToken)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("DELETE", "/api/admin/users/1", nil)
	req.Header.Set("Authorization", adminToken)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK, got %d", resp.Code)
	}
}
