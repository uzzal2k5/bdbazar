package repository

import (
    "gorm.io/gorm"
    "auth-service/models"
    "errors"
    "time"
)

type UserRepository interface {
    CreateUser(user *models.User) error
    FindByEmail(email string) (*models.User, error)
    FindByID(userID uint) (*models.User, error)
    FindByEmailOrMobile(email, mobile string) (models.User, error)
    StoreRefreshToken(userID uint, token string, expiresAt time.Time) error
    FindRefreshToken(token string) (models.RefreshToken, error)
    DeleteRefreshToken(token string) error
}

type userRepo struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil // ✅ return a pointer
}

// FindByEmailOrMobile checks for existing user by email or mobile
func (r *userRepo) FindByEmailOrMobile(email, mobile string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ? OR mobile = ?", email, mobile).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
       return models.User{}, gorm.ErrRecordNotFound
    }
	return user, err
}

// FindByID returns user by ID
func (r *userRepo) FindByID(userID uint) (*models.User, error) {
    var user models.User
    if err := r.db.First(&user, userID).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

// StoreRefreshToken saves a refresh token for a user
func (r *userRepo) StoreRefreshToken(userID uint, token string, expiresAt time.Time) error {
    rt := models.RefreshToken{
        Token:     token,
        UserID:    userID,
        ExpiresAt: expiresAt,
    }
    return r.db.Create(&rt).Error
}

// FindRefreshToken fetches refresh token record
func (r *userRepo) FindRefreshToken(token string) (models.RefreshToken, error) {
	var rt models.RefreshToken
	err := r.db.Where("token = ?", token).First(&rt).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.RefreshToken{}, errors.New("refresh token not found")
	}
	return rt, err
}

// DeleteRefreshToken deletes a token
func (r *userRepo) DeleteRefreshToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&models.RefreshToken{}).Error
}