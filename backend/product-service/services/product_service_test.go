package services

import (
	"errors"
	"product-service/models"
	"product-service/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductService_GetAll(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := NewProductService(mockRepo)

	expected := []models.Product{
		{Name: "Product1"}, {Name: "Product2"},
	}

	mockRepo.On("GetAll").Return(expected, nil)

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetByID(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := NewProductService(mockRepo)

	expected := &models.Product{ID: 1, Name: "Product1"}
	mockRepo.On("GetByID", uint(1)).Return(expected, nil)

	result, err := service.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestProductService_Create(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := NewProductService(mockRepo)

	product := &models.Product{Name: "New Product"}
	mockRepo.On("Create", product).Return(nil)

	err := service.Create(product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_Update(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := NewProductService(mockRepo)

	product := &models.Product{ID: 1, Name: "Updated Product"}
	mockRepo.On("Update", product).Return(nil)

	err := service.Update(product)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_Delete(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := NewProductService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := service.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetByID_Error(t *testing.T) {
	mockRepo := new(repository.MockProductRepository)
	service := NewProductService(mockRepo)

	mockRepo.On("GetByID", uint(2)).Return(&models.Product{}, errors.New("not found"))

	_, err := service.GetByID(2)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}
