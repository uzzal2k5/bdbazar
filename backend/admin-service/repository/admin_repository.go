package repository

import (
	"admin-service/models"
	"admin-service/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAllAdmins() ([]models.Admin, error)
	GetAdminByID(id uint) (*models.Admin, error)
	RegisterAdmin(admin *models.Admin) error
	UpdateAdmin(admin *models.Admin) error
	DeleteAdmin(id uint) error
	ResetAdminPassword(id uint, newPassword string) error
	FindByEmail(email string, admin *models.Admin) error
	FindByMobile(mobile string, admin *models.Admin) error

	CreateSuperAdmin(superAdmin *models.SuperAdmin) error
	FindSuperAdminByEmailOrMobile(email, mobile string) (*models.SuperAdmin, error)
//     FindSuperAdminByEmailOrMobile(email, mobile string) (*models.SuperAdmin, error)
}

type adminRepo struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepo{db}
}

// CreateSuperAdmin inserts a super admin if not exists
func (r *adminRepo) CreateSuperAdmin(superAdmin *models.SuperAdmin) error {
	var existing models.SuperAdmin
	// Assuming ID is a string "1" or similar unique static ID
	if err := r.db.First(&existing, "id = ?", superAdmin.ID).Error; err == nil {
		// Already exists
		return nil
	}
	return r.db.Create(superAdmin).Error
}

// GetSuperAdmin fetches super admin by static ID
func (r *adminRepo) GetSuperAdmin() (*models.SuperAdmin, error) {
	var superAdmin models.SuperAdmin
	if err := r.db.First(&superAdmin, "id = ?", "1").Error; err != nil {
		return nil, err
	}
	return &superAdmin, nil
}

// FindSuperAdminByEmailOrMobile fetches super admin by email or mobile
func (r *adminRepo) FindSuperAdminByEmailOrMobile(email, mobile string) (*models.SuperAdmin, error) {
    var superAdmin models.SuperAdmin
    err := r.db.Where("email = ? OR mobile = ?", email, mobile).First(&superAdmin).Error
    if err != nil {
        if err == gorm.ErrRecordNotFound {
    			return nil, nil // no record found
    	}
        return nil, err
    }
    return &superAdmin, nil
}


// RegisterAdmin sends new admin registration to auth-service
func (r *adminRepo) RegisterAdmin(admin *models.Admin) error {
	reqBody, err := json.Marshal(admin)
	if err != nil {
		return err
	}
	url := config.GetAuthServiceURL() + "/register"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("auth-service returned status %d", resp.StatusCode)
	}
	return nil
}

// UpdateAdmin updates admin record
func (r *adminRepo) UpdateAdmin(updatedAdmin *models.Admin) error {
	var admin models.Admin
	if err := r.db.First(&admin, updatedAdmin.ID).Error; err != nil {
		return err
	}

	admin.Name = updatedAdmin.Name
	admin.Username = updatedAdmin.Username
	admin.Email = updatedAdmin.Email
	admin.Mobile = updatedAdmin.Mobile
	admin.Role = updatedAdmin.Role

	return r.db.Save(&admin).Error
}

// DeleteAdmin deletes admin by ID
func (r *adminRepo) DeleteAdmin(id uint) error {
	return r.db.Delete(&models.Admin{}, id).Error
}

// ResetAdminPassword resets password in DB (or via auth-service)
func (r *adminRepo) ResetAdminPassword(id uint, newPassword string) error {
	var admin models.Admin
	if err := r.db.First(&admin, id).Error; err != nil {
		return err
	}
	admin.Password = newPassword
	return r.db.Save(&admin).Error
}

func (r *adminRepo) FindByEmail(email string, admin *models.Admin) error {
	return r.db.Where("email = ?", email).First(admin).Error
}

func (r *adminRepo) FindByMobile(mobile string, admin *models.Admin) error {
	return r.db.Where("mobile = ?", mobile).First(admin).Error
}

// GetAllAdmins returns all admins
func (r *adminRepo) GetAllAdmins() ([]models.Admin, error) {
	var admins []models.Admin
	result := r.db.Find(&admins)
	return admins, result.Error
}

// GetAdminByID fetches admin by ID
func (r *adminRepo) GetAdminByID(id uint) (*models.Admin, error) {
	var admin models.Admin
	result := r.db.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}
