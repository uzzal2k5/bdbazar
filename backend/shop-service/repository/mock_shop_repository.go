package repository

import (
	"shop-service/models"

	"github.com/stretchr/testify/mock"
)

type MockShopRepository struct {
	mock.Mock
}

func (m *MockShopRepository) Create(shop *models.Shop) error {
	args := m.Called(shop)
	return args.Error(0)
}

func (m *MockShopRepository) GetByID(id uint) (*models.Shop, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Shop), args.Error(1)
}

func (m *MockShopRepository) GetBySellerID(sellerID uint) ([]models.Shop, error) {
	args := m.Called(sellerID)
	return args.Get(0).([]models.Shop), args.Error(1)
}

func (m *MockShopRepository) Update(shop *models.Shop) error {
	args := m.Called(shop)
	return args.Error(0)
}

func (m *MockShopRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
