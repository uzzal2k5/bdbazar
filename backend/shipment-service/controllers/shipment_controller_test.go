package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"shipment-service/models"
	"shipment-service/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockShipmentService struct {
	mock.Mock
}

func (m *MockShipmentService) Create(s *models.Shipment) error {
	args := m.Called(s)
	return args.Error(0)
}
func (m *MockShipmentService) GetByOrderID(orderID uint) (*models.Shipment, error) {
	args := m.Called(orderID)
	return args.Get(0).(*models.Shipment), args.Error(1)
}
func (m *MockShipmentService) GetBySeller(sellerID uint) ([]models.Shipment, error) {
	args := m.Called(sellerID)
	return args.Get(0).([]models.Shipment), args.Error(1)
}
func (m *MockShipmentService) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}
func (m *MockShipmentService) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupRouter(ctrl *ShipmentController, role string, userID uint) *gin.Engine {
	r := gin.Default()
	// Middleware to set user_id and role
	r.Use(func(c *gin.Context) {
		c.Set("role", role)
		c.Set("user_id", userID)
		c.Next()
	})
	r.POST("/api/shipments", ctrl.CreateShipment)
	r.GET("/api/shipments/seller", ctrl.GetShipmentsBySeller)
	r.GET("/api/shipments/order/:order_id", ctrl.GetShipmentByOrder)
	r.PUT("/api/shipments/:id/status", ctrl.UpdateStatus)
	r.DELETE("/api/shipments/:id", ctrl.DeleteShipment)
	return r
}

func TestCreateShipment_Seller(t *testing.T) {
	mockSvc := new(MockShipmentService)
	ctrl := NewShipmentController(mockSvc)
	router := setupRouter(ctrl, "seller", 10)

	shipment := models.Shipment{OrderID: 1, BuyerID: 20, Address: "Addr"}
	mockSvc.On("Create", mock.AnythingOfType("*models.Shipment")).Return(nil).Run(func(args mock.Arguments) {
		s := args.Get(0).(*models.Shipment)
		s.Status = "pending"
	})

	jsonValue, _ := json.Marshal(shipment)
	req, _ := http.NewRequest("POST", "/api/shipments", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestCreateShipment_NotSeller(t *testing.T) {
	mockSvc := new(MockShipmentService)
	ctrl := NewShipmentController(mockSvc)
	router := setupRouter(ctrl, "buyer", 10)

	shipment := models.Shipment{OrderID: 1, BuyerID: 20, Address: "Addr"}
	jsonValue, _ := json.Marshal(shipment)
	req, _ := http.NewRequest("POST", "/api/shipments", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestGetShipmentsBySeller(t *testing.T) {
	mockSvc := new(MockShipmentService)
	ctrl := NewShipmentController(mockSvc)
	router := setupRouter(ctrl, "seller", 10)

	expected := []models.Shipment{{ID: 1, SellerID: 10}}
	mockSvc.On("GetBySeller", uint(10)).Return(expected, nil)

	req, _ := http.NewRequest("GET", "/api/shipments/seller", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetShipmentByOrder(t *testing.T) {
	mockSvc := new(MockShipmentService)
	ctrl := NewShipmentController(mockSvc)
	router := setupRouter(ctrl, "seller", 10)

	shipment := &models.Shipment{ID: 1, OrderID: 1, SellerID: 10, BuyerID: 20}
	mockSvc.On("GetByOrderID", uint(1)).Return(shipment, nil)

	req, _ := http.NewRequest("GET", "/api/shipments/order/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestGetShipmentByOrder_Forbidden(t *testing.T) {
	mockSvc := new(MockShipmentService)
	ctrl := NewShipmentController(mockSvc)
	router := setupRouter(ctrl, "seller", 999) // Different sellerID

	shipment := &models.Shipment{ID: 1, OrderID: 1, SellerID: 10, BuyerID: 20}
	mockSvc.On("GetByOrderID", uint(1)).Return(shipment, nil)

	req, _ := http.NewRequest("GET", "/api/shipments/order/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestUpdateStatus_Seller(t *testing.T) {
	mockSvc := new(MockShipmentService)
	ctrl := NewShipmentController(mockSvc)
	router := setupRouter(ctrl, "seller", 10)

	mockSvc.On("UpdateStatus", uint(1), "shipped").Return(nil)

	body := `{"status":"shipped"}`
	req, _ := http.NewRequest("PUT", "/api/shipments/1/status", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestUpdateStatus_NotSeller(t *testing.T) {
	mockSvc := new(MockShipmentService)
	ctrl := NewShipmentController(mockSvc)
	router := setupRouter(ctrl, "buyer", 10)

	body := `{"status":"shipped"}`
	req, _ := http.NewRequest("PUT", "/api/shipments/1/status", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestDeleteShipment_Seller(t *testing.T) {
	mockSvc := new(MockShipmentService)
	ctrl := NewShipmentController(mockSvc)
	router := setupRouter(ctrl, "seller", 10)

	mockSvc.On("Delete", uint(1)).Return(nil)

	req, _ := http.NewRequest("DELETE", "/api/shipments/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestDeleteShipment_NotSeller(t *testing.T) {
	mockSvc := new(MockShipmentService)
	ctrl := NewShipmentController(mockSvc)
	router := setupRouter(ctrl, "buyer", 10)

	req, _ := http.NewRequest("DELETE", "/api/shipments/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}
