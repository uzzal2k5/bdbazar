package services

import (
	"admin-service/models"
	"admin-service/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockUser_Success(t *testing.T) {
	mockRepo := new(repository.MockAdminRepository)
	service := NewAdminService(mockRepo)

	mockRepo.On("BlockUser", uint(1)).Return(nil)

	err := service.BlockUser(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestBlockUser_Error(t *testing.T) {
	mockRepo := new(repository.MockAdminRepository)
	service := NewAdminService(mockRepo)

	mockRepo.On("BlockUser", uint(999)).Return(errors.New("user not found"))

	err := service.BlockUser(999)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}

func TestApproveShop_Success(t *testing.T) {
	mockRepo := new(repository.MockAdminRepository)
	service := NewAdminService(mockRepo)

	mockRepo.On("ApproveShop", uint(1)).Return(nil)

	err := service.ApproveShop(1)

	assert.NoError(t, err)
}

func TestApproveShop_Error(t *testing.T) {
	mockRepo := new(repository.MockAdminRepository)
	service := NewAdminService(mockRepo)

	mockRepo.On("ApproveShop", uint(404)).Return(errors.New("shop not found"))

	err := service.ApproveShop(404)

	assert.Error(t, err)
	assert.Equal(t, "shop not found", err.Error())
}

func TestForceResetPassword(t *testing.T) {
	mockRepo := new(repository.MockAdminRepository)
	service := NewAdminService(mockRepo)

	mockRepo.On("ForceResetPassword", uint(1), "newpass123").Return(nil)

	err := service.ForceResetPassword(1, "newpass123")

	assert.NoError(t, err)
}

func TestDeleteUser_Success(t *testing.T) {
	mockRepo := new(repository.MockAdminRepository)
	service := NewAdminService(mockRepo)

	mockRepo.On("DeleteUser", uint(2)).Return(nil)

	err := service.DeleteUser(2)

	assert.NoError(t, err)
}

func TestGetMetrics(t *testing.T) {
	mockRepo := new(repository.MockAdminRepository)
	service := NewAdminService(mockRepo)

	mockResult := models.AdminMetrics{
		TotalUsers:  120,
		TotalOrders: 300,
		TotalRevenue: 15000.00,
	}

	mockRepo.On("GetMetrics").Return(mockResult, nil)

	metrics, err := service.GetMetrics()

	assert.NoError(t, err)
	assert.Equal(t, 120, metrics.TotalUsers)
	assert.Equal(t, 15000.00, metrics.TotalRevenue)
}
