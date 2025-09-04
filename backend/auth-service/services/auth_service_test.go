package services

import (
	"auth-service/models"
	"auth-service/repository"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("JWT_SECRET", "testsecret123")
}

func TestRegisterUser_Success(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := NewAuthService(mockRepo)

	user := &models.User{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "password123",
		Role:     "seller",
	}

	// simulate that email is not already used
	mockRepo.On("GetByEmail", user.Email).Return(&models.User{}, errors.New("not found"))
	mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(nil)

	err := service.RegisterUser(user)

	assert.NoError(t, err)
	assert.NotEmpty(t, user.Password) // should be hashed
	assert.NotEqual(t, "password123", user.Password)
	mockRepo.AssertExpectations(t)
}

func TestRegisterUser_EmailExists(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := NewAuthService(mockRepo)

	existingUser := &models.User{
		Email: "alice@example.com",
	}

	mockRepo.On("GetByEmail", existingUser.Email).Return(existingUser, nil)

	user := &models.User{
		Email: "alice@example.com",
	}

	err := service.RegisterUser(user)

	assert.Error(t, err)
	assert.Equal(t, "email already in use", err.Error())
}

func TestLoginUser_Success(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := NewAuthService(mockRepo)

	hashedPwd, _ := HashPassword("password123")

	user := &models.User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: hashedPwd,
		Role:     "buyer",
	}

	mockRepo.On("GetByEmail", "alice@example.com").Return(user, nil)

	token, err := service.LoginUser("alice@example.com", "password123")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// check token format
	parts := len(token)
	assert.Greater(t, parts, 10)
}

func TestLoginUser_InvalidPassword(t *testing.T) {
	mockRepo := new(repository.MockUserRepository)
	service := NewAuthService(mockRepo)

	hashedPwd, _ := HashPassword("correct-password")

	user := &models.User{
		ID:       1,
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: hashedPwd,
		Role:     "buyer",
	}

	mockRepo.On("GetByEmail", "alice@example.com").Return(user, nil)

	token, err := service.LoginUser("alice@example.com", "wrong-password")

	assert.Error(t, err)
	assert.Empty(t, token)
}

func TestHashPasswordAndCheck(t *testing.T) {
	password := "mypassword"
	hashed, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEqual(t, password, hashed)

	valid := CheckPasswordHash(password, hashed)
	assert.True(t, valid)
}
