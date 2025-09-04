package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"

	"auth-service/models"
	"auth-service/services"
)

type AuthController struct {
	authService services.AuthService
}

// NewAuthController initializes AuthController with AuthService
func NewAuthController(authService services.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

// Register handles POST /api/auth/register
func (c *AuthController) Register(ctx *gin.Context) {
	var input struct {
		Name     string   `json:"name" binding:"required"`
		Email    string   `json:"email" binding:"required,email"`
		Mobile   string   `json:"mobile" binding:"required"`
		Password string   `json:"password" binding:"required,min=6"`
		Roles    []string `json:"roles" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists by email or mobile via service (implement this in your service)
	existingUser, err := c.authService.FindByEmailOrMobile(input.Email, input.Mobile)
	if err == nil && existingUser.ID != 0 {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User already registered"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	rolesJSON, err := json.Marshal(input.Roles)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid role format"})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Mobile:   input.Mobile,
		Password: string(hashedPassword),
		Roles:    datatypes.JSON(rolesJSON),
	}

	if err := c.authService.Register(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handles POST /api/auth/login
func (c *AuthController) Login(ctx *gin.Context) {
	var input struct {
		Identifier    string `json:"identifier" binding:"required"`
		Password string   `json:"password" binding:"required,min=6"`
	}

    // Validate input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    // Delegate authentication to service
	accessToken, refreshToken, err := c.authService.Login(input.Identifier, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

    // Respond with tokens
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Refresh handles POST /api/auth/refresh
func (c *AuthController) Refresh(ctx *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := c.authService.Refresh(input.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Logout handles POST /api/auth/logout
func (c *AuthController) Logout(ctx *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.authService.Logout(input.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
