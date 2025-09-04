package services

import (
	"fmt"
	"log"

	"admin-service/config"
	"admin-service/models"
	"admin-service/repository"
	"admin-service/utils"
)

type AdminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) *AdminService {
	return &AdminService{adminRepo: repo}
}

// GetAllAdmins returns all admins
func (s *AdminService) GetAllAdmins() ([]models.Admin, error) {
	return s.adminRepo.GetAllAdmins()
}

// GetAdminByID fetches admin by ID
func (s *AdminService) GetAdminByID(id uint) (*models.Admin, error) {
	return s.adminRepo.GetAdminByID(id)
}

// SetupSuperAdmin creates a Super Admin user with a static ID (only once)
func (s *AdminService) SetupSuperAdmin(cfg config.SuperAdminConfig) error {
    // Call ExistsSuperAdminByEmailOrMobile on the service receiver 's' directly, not s.AdminService
    spUser, err := s.ExistsSuperAdminByEmailOrMobile(cfg.Email, cfg.Mobile)
    if err != nil {
        return err
    }
   if spUser != nil {
       log.Println("Super admin already exists with given email or mobile")
       return nil
   }


    // Hash password (ensure utils.HashPassword returns (string, error))
    hashedPassword, err := utils.HashPassword(cfg.Password)
    if err != nil {
        return fmt.Errorf("failed to hash password: %w", err)
    }

    // Create SuperAdmin instance with string ID "1"
    superAdmin := &models.SuperAdmin{
        ID:       "1",  // string ID
        Name:     cfg.Name,
        Username: cfg.Username,
        Password: hashedPassword,
        Email:    cfg.Email,
        Mobile:   cfg.Mobile,
        Status:   "active",
    }

    // Save to DB via repo
    if err := s.adminRepo.CreateSuperAdmin(superAdmin); err != nil {
        return fmt.Errorf("❌ failed to create super admin: %w", err)
    }

    log.Println("✅ Super admin created on first launch")
    return nil
}


func (s *AdminService) ExistsSuperAdminByEmailOrMobile(email, mobile string) (*models.SuperAdmin, error) {
    superAdmin, err := s.adminRepo.FindSuperAdminByEmailOrMobile(email, mobile)
    if err != nil {
        return nil, err
    }
    return superAdmin, nil
}

// RegisterAdminUser checks if the admin already exists and then registers via auth-service
// RegisterAdminUser registers a new admin, only if the requester is a super admin
func (s *AdminService) RegisterAdminUser(requester *models.Admin, admin *models.Admin) error {
    if requester.Role != "superadmin" {
		return fmt.Errorf("❌ only a superadmin can create new admins")
	}
	admin.Role = "admin" // enforce admin role

	var existing models.Admin
	if err := s.adminRepo.FindByEmail(admin.Email, &existing); err == nil {
		return fmt.Errorf("email '%s' is already in use", admin.Email)
	}
	if err := s.adminRepo.FindByMobile(admin.Mobile, &existing); err == nil {
		return fmt.Errorf("mobile '%s' is already in use", admin.Mobile)
	}

	// Register to auth-service
	if err := s.adminRepo.RegisterAdmin(admin); err != nil {
		return fmt.Errorf("failed to register admin via auth-service: %w", err)
	}

	return nil
}


// UpdateAdmin updates an existing admin
func (s *AdminService) UpdateAdmin(id uint, updatedAdmin *models.Admin) error {
	updatedAdmin.ID = id
	return s.adminRepo.UpdateAdmin(updatedAdmin)
}

// DeleteAdmin removes an admin from the system
func (s *AdminService) DeleteAdmin(id uint) error {
	return s.adminRepo.DeleteAdmin(id)
}

// BlockUser calls auth-service to block a user
func (s *AdminService) BlockUser(userID string) error {
	url := fmt.Sprintf("%s/api/users/%s/block", config.GetAuthServiceURL(), userID)
	headers := map[string]string{"Content-Type": "application/json"}
	_, err := utils.HttpRequest("PATCH", url, nil, headers)
	return err
}

// ApproveUser calls auth-service to approve a user
func (s *AdminService) ApproveUser(userID string) error {
	url := fmt.Sprintf("%s/api/users/%s/approve", config.GetAuthServiceURL(), userID)
	headers := map[string]string{"Content-Type": "application/json"}
	_, err := utils.HttpRequest("PATCH", url, nil, headers)
	return err
}

// ResetAdminPassword resets the password via auth-service
func (s *AdminService) ResetAdminPassword(userID uint) error {
	url := fmt.Sprintf("%s/api/users/%d/reset-password", config.GetAuthServiceURL(), userID)
	headers := map[string]string{"Content-Type": "application/json"}
	_, err := utils.HttpRequest("POST", url, nil, headers)
	return err
}

// DeleteUser deletes a user via auth-service
func (s *AdminService) DeleteUser(userID uint) error {
	url := fmt.Sprintf("%s/api/users/%d", config.GetAuthServiceURL(), userID)
	headers := map[string]string{"Content-Type": "application/json"}
	_, err := utils.HttpRequest("DELETE", url, nil, headers)
	return err
}

// ApproveShop approves a shop via shop-service
func (s *AdminService) ApproveShop(shopID string) error {
	url := fmt.Sprintf("%s/api/shops/%s/approve", config.GetShopServiceURL(), shopID)
	headers := map[string]string{"Content-Type": "application/json"}
	_, err := utils.HttpRequest("PATCH", url, nil, headers)
	return err
}

// BlockShop blocks a shop via shop-service
func (s *AdminService) BlockShop(shopID string) error {
	url := fmt.Sprintf("%s/api/shops/%s/block", config.GetShopServiceURL(), shopID)
	headers := map[string]string{"Content-Type": "application/json"}
	_, err := utils.HttpRequest("PATCH", url, nil, headers)
	return err
}

// Dashboard returns mock dashboard data
func (s *AdminService) Dashboard() map[string]interface{} {
	return map[string]interface{}{
		"total_users": 100,
		"total_shops": 20,
	}
}

// GetMetrics returns mock system metrics
func (s *AdminService) GetMetrics() map[string]interface{} {
	return map[string]interface{}{
		"uptime":   "99.99%",
		"requests": 10324,
		"latency":  "120ms",
	}
}
