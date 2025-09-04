package services

import (
	"errors"
	"shipment-service/models"
	"shipment-service/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockShipmentRepo struct {
	mock.Mock
}

func (m *MockShipmentRepo) Create(s *models.Shipment) error {
	args := m.Called(s)
	return args.Error(0)
}

func (m *MockShipmentRepo) GetByOrderID(orderID uint) (*models.Shipment, error) {
	args := m.Called(orderID)
	return args.Get(0).(*models.Shipment), args.Error(1)
}

func (m *MockShipmentRepo) GetBySeller(sellerID uint) ([]models.Shipment, error) {
	args := m.Called(sellerID)
	return args.Get(0).([]models.Shipment), args.Error(1)
}

func (m *MockShipmentRepo) UpdateStatus(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockShipmentRepo) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateShipment(t *testing.T) {
	mockRepo := new(MockShipmentRepo)
	service := NewShipmentService(mockRepo)

	shipment := &models.Shipment{OrderID: 1, SellerID: 2, BuyerID: 3, Address: "Addr"}

	mockRepo.On("Create", shipment).Return(nil)

	err := service.Create(shipment)
	assert.NoError(t, err)
	assert.Equal(t, "pending", shipment.Status)

	mockRepo.AssertExpectations(t)
}

func TestGetByOrderID(t *testing.T) {
	mockRepo := new(MockShipmentRepo)
	service := NewShipmentService(mockRepo)

	expected := &models.Shipment{ID: 1, OrderID: 1}
	mockRepo.On("GetByOrderID", uint(1)).Return(expected, nil)

	result, err := service.GetByOrderID(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockRepo.AssertExpectations(t)
}

func TestGetBySeller(t *testing.T) {
	mockRepo := new(MockShipmentRepo)
	service := NewShipmentService(mockRepo)

	expected := []models.Shipment{{ID: 1, SellerID: 2}, {ID: 2, SellerID: 2}}
	mockRepo.On("GetBySeller", uint(2)).Return(expected, nil)

	result, err := service.GetBySeller(2)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockRepo.AssertExpectations(t)
}

func TestUpdateStatus(t *testing.T) {
	mockRepo := new(MockShipmentRepo)
	service := NewShipmentService(mockRepo)

	mockRepo.On("UpdateStatus", uint(1), "shipped").Return(nil)

	err := service.UpdateStatus(1, "shipped")
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteShipment(t *testing.T) {
	mockRepo := new(MockShipmentRepo)
	service := NewShipmentService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := service.Delete(1)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
