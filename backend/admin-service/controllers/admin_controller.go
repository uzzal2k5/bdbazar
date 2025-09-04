package controllers

import (
	"admin-service/models"
	"admin-service/services"
	"admin-service/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	Service *services.AdminService
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"` // Email or mobile
	Password string `json:"password" binding:"required"`
}

func NewAdminController(service *services.AdminService) *AdminController {
	return &AdminController{Service: service}
}

// Login for Super Admin
func (ctrl *AdminController) LoginSuperAdmin(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	superAdmin, err := ctrl.Service.ExistsSuperAdminByEmailOrMobile(req.Username, req.Username)
	if err != nil || superAdmin == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(req.Password, superAdmin.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	idUint64, err := strconv.ParseUint(superAdmin.ID, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid super admin ID"})
		return
	}

	token, err := utils.GenerateJWT(uint(idUint64), "superadmin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"superadmin": gin.H{
			"id":       superAdmin.ID,
			"username": superAdmin.Username,
			"email":    superAdmin.Email,
			"role":     "superadmin",
		},
	})
}

// Helper method to extract admin from context
func (ctrl *AdminController) getCurrentAdmin(c *gin.Context) (*models.Admin, error) {
	if adminRaw, exists := c.Get("admin"); exists {
		if admin, ok := adminRaw.(*models.Admin); ok {
			return admin, nil
		}
		return nil, fmt.Errorf("invalid admin type in context")
	}

	userIDRaw, exists := c.Get("userID")
	if !exists {
		return nil, fmt.Errorf("no user info found in context")
	}
	userIDStr, ok := userIDRaw.(string)
	if !ok {
		return nil, fmt.Errorf("invalid userID type in context")
	}
	idUint64, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format")
	}
	return ctrl.Service.GetAdminByID(uint(idUint64))
}

// Admin Listing
func (ctrl *AdminController) ListAdmins(c *gin.Context) {
	admins, err := ctrl.Service.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": admins})
}

func (ctrl *AdminController) ListAdminByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	admin, err := ctrl.Service.GetAdminByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": admin})
}

// Admin User Management
func (ctrl *AdminController) CreateAdmin(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requesterAdmin, err := ctrl.getCurrentAdmin(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
		return
	}

	if err := ctrl.Service.RegisterAdminUser(requesterAdmin, &admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, admin)
}

func (ctrl *AdminController) UpdateAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Service.UpdateAdmin(uint(id), &admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (ctrl *AdminController) DeleteAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	if err := ctrl.Service.DeleteAdmin(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (ctrl *AdminController) ResetAdminPassword(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	if err := ctrl.Service.ResetAdminPassword(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// Admin Dashboard
func (ctrl *AdminController) Dashboard(c *gin.Context) {
	data := ctrl.Service.Dashboard()
	c.JSON(http.StatusOK, data)
}

func (ctrl *AdminController) GetMetrics(c *gin.Context) {
	data := ctrl.Service.GetMetrics()
	c.JSON(http.StatusOK, data)
}

// User Management
func (ctrl *AdminController) ApproveUser(c *gin.Context) {
	userID := c.Param("id")
	if err := ctrl.Service.ApproveUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User approved"})
}

func (ctrl *AdminController) BlockUser(c *gin.Context) {
	userID := c.Param("id")
	if err := ctrl.Service.BlockUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User blocked"})
}

func (ctrl *AdminController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := ctrl.Service.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// Shop Management
func (ctrl *AdminController) ApproveShop(c *gin.Context) {
	shopID := c.Param("id")
	if err := ctrl.Service.ApproveShop(shopID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shop approved"})
}

func (ctrl *AdminController) BlockShop(c *gin.Context) {
	shopID := c.Param("id")
	if err := ctrl.Service.BlockShop(shopID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shop blocked"})
}
