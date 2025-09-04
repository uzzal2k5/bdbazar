package services

import (
	"order-service/models"

	"github.com/stretchr/testify/mock"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) CreateOrder(order *models.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderService) GetOrdersByBuyer(buyerID uint) ([]models.Order, error) {
	args := m.Called(buyerID)
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *MockOrderService) GetOrdersBySeller(sellerID uint) ([]models.Order, error) {
	args := m.Called(sellerID)
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *MockOrderService) MarkAsShipped(orderID uint) error {
	args := m.Called(orderID)
	return args.Error(0)
}
