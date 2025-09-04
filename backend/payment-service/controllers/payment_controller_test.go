package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"payment-service/models"
	"payment-service/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPaymentService struct {
	mock.Mock
}

func (m *MockPaymentService) Create(p *models.Payment) error {
	args := m.Called(p)
	return args.Error(0)
}
func (m *MockPaymentService) GetByBuyer(b uint) ([]models.Payment, error) {
	args := m.Called(b)
	return args.Get(0).([]models.Payment), args.Error(1)
}
func (m *MockPaymentService) GetBySeller(s uint) ([]models.Payment, error) {
	args := m.Called(s)
	return args.Get(0).([]models.Payment), args.Error(1)
}
func (m *MockPaymentService) CompletePayment(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupTestRouter(ctrl *PaymentController) *gin.Engine {
	r := gin.Default()
	r.POST("/api/payments", func(c *gin.Context) {
		// fake setting user_id & role for testing
		c.Set("user_id", uint(1))
		c.Set("role", "buyer")
		ctrl.Create(c)
	})
	r.GET("/api/payments/buyer", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		ctrl.GetByBuyer(c)
	})
	r.GET("/api/payments/seller", func(c *gin.Context) {
		c.Set("user_id", uint(2))
		ctrl.GetBySeller(c)
	})
	r.PUT("/api/payments/:id/complete", func(c *gin.Context) {
		c.Set("role", "buyer")
		ctrl.Complete(c)
	})
	return r
}

func TestCreatePaymentController(t *testing.T) {
	mockSvc := new(MockPaymentService)
	ctrl := NewPaymentController(mockSvc)
	r := setupTestRouter(ctrl)

	p := models.Payment{OrderID: 10, Amount: 100.0}
	mockSvc.On("Create", mock.AnythingOfType("*models.Payment")).Return(nil).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*models.Payment)
		arg.Status = "pending"
	})

	body, _ := json.Marshal(p)
	req, _ := http.NewRequest("POST", "/api/payments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestCreatePaymentController_BadRequest(t *testing.T) {
	mockSvc := new(MockPaymentService)
	ctrl := NewPaymentController(mockSvc)
	r := setupTestRouter(ctrl)

	req, _ := http.NewRequest("POST", "/api/payments", bytes.NewBuffer([]byte("bad json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreatePaymentController_Forbidden(t *testing.T) {
	mockSvc := new(MockPaymentService)
	ctrl := NewPaymentController(mockSvc)
	r := setupTestRouter(ctrl)

	r.POST("/api/payments", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		c.Set("role", "seller") // Not buyer
		ctrl.Create(c)
	})

	p := models.Payment{OrderID: 10, Amount: 100.0}
	body, _ := json.Marshal(p)
	req, _ := http.NewRequest("POST", "/api/payments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestGetByBuyerController(t *testing.T) {
	mockSvc := new(MockPaymentService)
	ctrl := NewPaymentController(mockSvc)
	r := setupTestRouter(ctrl)

	expected := []models.Payment{{ID: 1, Amount: 10}}
	mockSvc.On("GetByBuyer", uint(1)).Return(expected, nil)

	req, _ := http.NewRequest("GET", "/api/payments/buyer", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetByBuyerController_Error(t *testing.T) {
	mockSvc := new(MockPaymentService)
	ctrl := NewPaymentController(mockSvc)
	r := setupTestRouter(ctrl)

	mockSvc.On("GetByBuyer", uint(1)).Return(nil, errors.New("db error"))

	req, _ := http.NewRequest("GET", "/api/payments/buyer", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetBySellerController(t *testing.T) {
	mockSvc := new(MockPaymentService)
	ctrl := NewPaymentController(mockSvc)
	r := setupTestRouter(ctrl)

	expected := []models.Payment{{ID: 2, Amount: 20}}
	mockSvc.On("GetBySeller", uint(2)).Return(expected, nil)

	req, _ := http.NewRequest("GET", "/api/payments/seller", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestCompletePaymentController(t *testing.T) {
	mockSvc := new(MockPaymentService)
	ctrl := NewPaymentController(mockSvc)
	r := setupTestRouter(ctrl)

	mockSvc.On("CompletePayment", uint(1)).Return(nil)

	req, _ := http.NewRequest("PUT", "/api/payments/1/complete", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestCompletePaymentController_Forbidden(t *testing.T) {
	mockSvc := new(MockPaymentService)
	ctrl := NewPaymentController(mockSvc)
	r := setupTestRouter(ctrl)

	r.PUT("/api/payments/:id/complete", func(c *gin.Context) {
		c.Set("role", "seller") // Not buyer
		ctrl.Complete(c)
	})

	req, _ := http.NewRequest("PUT", "/api/payments/1/complete", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}
