package controllers

import (
	"errors"
	"net/http"
	"payment-service/models"
	"payment-service/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentController struct {
	Service   services.PaymentService
	Validator *validator.Validate
}

func NewPaymentController(service services.PaymentService) *PaymentController {
	return &PaymentController{
		Service:   service,
		Validator: validator.New(),
	}
}

// =======================
// 🔐 Helper: Get Auth Info
// =======================
func getAuthUser(c *gin.Context) (uint, string, error) {
	idVal, ok := c.Get("user_id")
	if !ok {
		return 0, "", errors.New("user ID not found in context")
	}
	id, ok := idVal.(uint)
	if !ok {
		return 0, "", errors.New("user ID is not a valid number")
	}

	roleVal, ok := c.Get("role")
	if !ok {
		return 0, "", errors.New("role not found in context")
	}
	role, ok := roleVal.(string)
	if !ok {
		return 0, "", errors.New("role is not a valid string")
	}

	return id, role, nil
}

// ===========================
// 💳 POST /api/payments
// ===========================
func (pc *PaymentController) CreatePayment(c *gin.Context) {
	var payment models.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate struct
	if err := pc.Validator.Struct(payment); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
		return
	}

	// Auth check
	userID, role, err := getAuthUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if role != "buyer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only buyers can initiate payments"})
		return
	}

	payment.BuyerID = userID
	payment.Status = models.StatusPending // Set default status

	if err := pc.Service.Create(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Payment created", "payment": payment})
}

// ===========================
// 📥 GET /api/payments/buyer
// ===========================
func (pc *PaymentController) GetPaymentsByBuyer(c *gin.Context) {
	userID, role, err := getAuthUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if role != "buyer" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	payments, err := pc.Service.GetByBuyer(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

// ===========================
// 📤 GET /api/payments/seller
// ===========================
func (pc *PaymentController) GetPaymentsBySeller(c *gin.Context) {
	userID, role, err := getAuthUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if role != "seller" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	payments, err := pc.Service.GetBySeller(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

// ===========================
// ✅ PUT /api/payments/:id/complete
// ===========================
func (pc *PaymentController) CompletePayment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}

	_, role, err := getAuthUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if role != "seller" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can complete payments"})
		return
	}

	// You can also check here if the seller actually owns this payment (optional)
	if err := pc.Service.CompletePayment(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment marked as completed"})
}
