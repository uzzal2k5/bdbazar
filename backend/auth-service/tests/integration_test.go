package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"auth-service/config"
	"auth-service/controllers"
	"auth-service/models"
	"auth-service/repository"
	"auth-service/routes"
	"auth-service/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	_ = godotenv.Load("../.env")

	os.Setenv("DB_SOURCE", "host=localhost user=postgres password=postgres123 dbname=bdbazar_auth_test port=55433 sslmode=disable")

	db := config.ConnectDB()
	db.AutoMigrate(&models.User{})

	repo := repository.NewAuthRepository(db)
	svc := services.NewAuthService(repo)
	ctrl := controllers.NewAuthController(svc)

	r := gin.Default()
	routes.SetupAuthRoutes(r, ctrl)
	return r
}

func TestAuth_RegisterAndLogin(t *testing.T) {
	r := setupRouter()

	user := models.RegisterInput{
		Name:     "Test User",
		Email:    "testuser@example.com",
		Password: "testpass123",
		Role:     "buyer",
	}
	jsonBody, _ := json.Marshal(user)

	// Register user
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated && w.Code != http.StatusConflict {
		t.Fatalf("Expected 201 Created or 409 Conflict, got %d", w.Code)
	}

	// Login
	loginBody := map[string]string{
		"email":    "testuser@example.com",
		"password": "testpass123",
	}
	jsonLogin, _ := json.Marshal(loginBody)

	loginReq, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonLogin))
	loginReq.Header.Set("Content-Type", "application/json")
	loginResp := httptest.NewRecorder()
	r.ServeHTTP(loginResp, loginReq)

	if loginResp.Code != http.StatusOK {
		t.Fatalf("Expected 200 OK from login, got %d", loginResp.Code)
	}

	var resBody map[string]interface{}
	json.NewDecoder(loginResp.Body).Decode(&resBody)

	if _, exists := resBody["token"]; !exists {
		t.Fatalf("Expected token in login response")
	}
}
