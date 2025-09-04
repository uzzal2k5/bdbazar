package services

import (
	"errors"
	"payment-service/models"
	"payment-service/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockPaymentRepo struct {
	mock.Mock
}

func (m *MockPaymentRepo) Create(p *models.Payment) error {
	args := m.Called(p)
	return args.Error(0)
}
func (m *MockPaymentRepo) GetByBuyer(b uint) ([]models.Payment, error) {
	args := m.Called(b)
	return args.Get(0).([]models.Payment), args.Error(1)
}
func (m *MockPaymentRepo) GetBySeller(s uint) ([]models.Payment, error) {
	args := m.Called(s)
	return args.Get(0).([]models.Payment), args.Error(1)
}
func (m *MockPaymentRepo) UpdateStatus(id uint, s string) error {
	args := m.Called(id, s)
	return args.Error(0)
}

func TestCreatePayment(t *testing.T) {
	mockRepo := new(MockPaymentRepo)
	svc := NewPaymentService(mockRepo)

	p := &models.Payment{Amount: 100.0}
	mockRepo.On("Create", p).Return(nil)

	err := svc.Create(p)
	assert.NoError(t, err)
	assert.Equal(t, "pending", p.Status)
	mockRepo.AssertExpectations(t)
}

func TestGetByBuyer(t *testing.T) {
	mockRepo := new(MockPaymentRepo)
	svc := NewPaymentService(mockRepo)

	expected := []models.Payment{{ID: 1, BuyerID: 2, Amount: 10.0}}
	mockRepo.On("GetByBuyer", uint(2)).Return(expected, nil)

	result, err := svc.GetByBuyer(2)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetBySeller(t *testing.T) {
	mockRepo := new(MockPaymentRepo)
	svc := NewPaymentService(mockRepo)

	expected := []models.Payment{{ID: 1, SellerID: 3, Amount: 15.0}}
	mockRepo.On("GetBySeller", uint(3)).Return(expected, nil)

	result, err := svc.GetBySeller(3)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestCompletePayment(t *testing.T) {
	mockRepo := new(MockPaymentRepo)
	svc := NewPaymentService(mockRepo)

	mockRepo.On("UpdateStatus", uint(1), "completed").Return(nil)

	err := svc.CompletePayment(1)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCompletePayment_Error(t *testing.T) {
	mockRepo := new(MockPaymentRepo)
	svc := NewPaymentService(mockRepo)

	mockRepo.On("UpdateStatus", uint(1), "completed").Return(errors.New("update error"))

	err := svc.CompletePayment(1)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
