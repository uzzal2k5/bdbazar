package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"order-service/models"
	"order-service/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouterWithController(controller *OrderController) *gin.Engine {
	r := gin.Default()

	r.POST("/api/orders", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		controller.CreateOrder(c)
	})

	r.GET("/api/orders/buyer", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		controller.GetBuyerOrders(c)
	})

	r.GET("/api/orders/seller", func(c *gin.Context) {
		c.Set("user_id", uint(2))
		controller.GetSellerOrders(c)
	})

	r.PUT("/api/orders/:id/ship", func(c *gin.Context) {
		c.Set("user_id", uint(2))
		c.Set("role", "seller")
		controller.MarkOrderShipped(c)
	})

	return r
}

func TestCreateOrder_Success(t *testing.T) {
	mockService := new(services.MockOrderService)
	controller := NewOrderController(mockService)
	router := setupRouterWithController(controller)

	order := models.Order{
		ProductID: 123,
		SellerID:  2,
		Quantity:  1,
	}

	mockService.On("CreateOrder", mock.AnythingOfType("*models.Order")).Return(nil)

	jsonData, _ := json.Marshal(order)
	req, _ := http.NewRequest("POST", "/api/orders", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockService.AssertExpectations(t)
}

func TestGetBuyerOrders_Success(t *testing.T) {
	mockService := new(services.MockOrderService)
	controller := NewOrderController(mockService)
	router := setupRouterWithController(controller)

	mockOrders := []models.Order{{ID: 1, BuyerID: 1}}
	mockService.On("GetOrdersByBuyer", uint(1)).Return(mockOrders, nil)

	req, _ := http.NewRequest("GET", "/api/orders/buyer", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestGetSellerOrders_Success(t *testing.T) {
	mockService := new(services.MockOrderService)
	controller := NewOrderController(mockService)
	router := setupRouterWithController(controller)

	mockOrders := []models.Order{{ID: 1, SellerID: 2}}
	mockService.On("GetOrdersBySeller", uint(2)).Return(mockOrders, nil)

	req, _ := http.NewRequest("GET", "/api/orders/seller", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}

func TestMarkOrderShipped_Success(t *testing.T) {
	mockService := new(services.MockOrderService)
	controller := NewOrderController(mockService)
	router := setupRouterWithController(controller)

	mockService.On("MarkAsShipped", uint(42)).Return(nil)

	req, _ := http.NewRequest("PUT", "/api/orders/42/ship", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockService.AssertExpectations(t)
}
