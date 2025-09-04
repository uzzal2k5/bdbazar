package services

import (
	"errors"
	"shop-service/models"
	"shop-service/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateShop_Success(t *testing.T) {
	mockRepo := new(repository.MockShopRepository)
	service := NewShopService(mockRepo)

	shop := &models.Shop{
		Name:     "MyShop",
		SellerID: 101,
	}

	mockRepo.On("Create", shop).Return(nil)

	err := service.Create(shop)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateShop_Error(t *testing.T) {
	mockRepo := new(repository.MockShopRepository)
	service := NewShopService(mockRepo)

	shop := &models.Shop{Name: "FailShop"}

	mockRepo.On("Create", shop).Return(errors.New("DB error"))

	err := service.Create(shop)
	assert.Error(t, err)
}

func TestGetByID_Success(t *testing.T) {
	mockRepo := new(repository.MockShopRepository)
	service := NewShopService(mockRepo)

	expected := &models.Shop{ID: 1, Name: "TestShop"}

	mockRepo.On("GetByID", uint(1)).Return(expected, nil)

	result, err := service.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetByID_NotFound(t *testing.T) {
	mockRepo := new(repository.MockShopRepository)
	service := NewShopService(mockRepo)

	mockRepo.On("GetByID", uint(2)).Return(&models.Shop{}, errors.New("not found"))

	_, err := service.GetByID(2)

	assert.Error(t, err)
	assert.Equal(t, "not found", err.Error())
}

func TestGetBySellerID(t *testing.T) {
	mockRepo := new(repository.MockShopRepository)
	service := NewShopService(mockRepo)

	mockShops := []models.Shop{
		{ID: 1, Name: "Shop1", SellerID: 200},
		{ID: 2, Name: "Shop2", SellerID: 200},
	}

	mockRepo.On("GetBySellerID", uint(200)).Return(mockShops, nil)

	result, err := service.GetBySellerID(200)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
}

func TestUpdateShop_Success(t *testing.T) {
	mockRepo := new(repository.MockShopRepository)
	service := NewShopService(mockRepo)

	shop := &models.Shop{ID: 1, Name: "UpdatedShop"}

	mockRepo.On("Update", shop).Return(nil)

	err := service.Update(shop)
	assert.NoError(t, err)
}

func TestDeleteShop_Success(t *testing.T) {
	mockRepo := new(repository.MockShopRepository)
	service := NewShopService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := service.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
