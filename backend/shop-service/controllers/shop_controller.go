package controllers

import (
    "fmt"
    "net/http"
    "shop-service/models"
    "shop-service/services"
    "strconv"

    "github.com/gin-gonic/gin"
)

// ShopController handles all shop-related endpoints
type ShopController struct {
    Service services.ShopService
}

// NewShopController initializes the controller
func NewShopController(service services.ShopService) *ShopController {
    return &ShopController{Service: service}
}

// ======================
// 🔐 Helper: Get User ID & Role from context safely
// ======================
func getAuthUser(c *gin.Context) (uint, string, error) {
    rawID, exists := c.Get("id")
    if !exists || rawID == nil {
        return 0, "", fmt.Errorf("user ID not found in context")
    }

    // JWT typically parses numeric values as float64
    uidFloat, ok := rawID.(float64)
    if !ok {
        return 0, "", fmt.Errorf("user ID is not a valid number")
    }

    rawRole, exists := c.Get("role")
    if !exists || rawRole == nil {
        return 0, "", fmt.Errorf("role not found in context")
    }

    roleStr, ok := rawRole.(string)
    if !ok {
        return 0, "", fmt.Errorf("role is not a valid string")
    }

    return uint(uidFloat), roleStr, nil
}

// ======================
// 🧑‍💼 Create Shop (Seller Only)
// ======================
func (ctrl *ShopController) CreateShop(c *gin.Context) {
    var shop models.Shop
    if err := c.ShouldBindJSON(&shop); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, role, err := getAuthUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    if role != "seller" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can create shops"})
        return
    }

    shop.OwnerID = userID

    if err := ctrl.Service.CreateShop(&shop); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, shop)
}

// ======================
// 📦 Get Shop By ID
// ======================
func (ctrl *ShopController) GetShop(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shop ID"})
        return
    }

    shop, err := ctrl.Service.GetShopByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Shop not found"})
        return
    }

    c.JSON(http.StatusOK, shop)
}

// ======================
// ✏️ Update Shop (Seller Only, Owner Only)
// ======================
func (ctrl *ShopController) UpdateShop(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shop ID"})
        return
    }

    var shop models.Shop
    if err := c.ShouldBindJSON(&shop); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, role, err := getAuthUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    if role != "seller" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can update shops"})
        return
    }

    existing, err := ctrl.Service.GetShopByID(uint(id))
    if err != nil || existing.OwnerID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own shop"})
        return
    }

    shop.ID = uint(id)
    shop.OwnerID = userID

    if err := ctrl.Service.UpdateShop(&shop); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, shop)
}

// ======================
// ❌ Delete Shop (Seller Only, Owner Only)
// ======================
func (ctrl *ShopController) DeleteShop(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shop ID"})
        return
    }

    userID, role, err := getAuthUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    if role != "seller" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can delete shops"})
        return
    }

    existing, err := ctrl.Service.GetShopByID(uint(id))
    if err != nil || existing.OwnerID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own shop"})
        return
    }

    if err := ctrl.Service.DeleteShop(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Shop deleted successfully"})
}

// ======================
// 📃 List All Approved Shops
// ======================
func (ctrl *ShopController) ListShops(c *gin.Context) {
    shops, err := ctrl.Service.ListShops()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, shops)
}

// ======================
// 🔍 Search Shop by Name
// ======================
func (ctrl *ShopController) SearchShops(c *gin.Context) {
    name := c.Query("name")
    if name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'name' is required"})
        return
    }

    shops, err := ctrl.Service.SearchShops(name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, shops)
}

// ======================
// 📊 Shop Dashboard (Seller Only)
// ======================
func (ctrl *ShopController) GetDashboard(c *gin.Context) {
    userID, role, err := getAuthUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    if role != "seller" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can access dashboard"})
        return
    }

    dashboard, err := ctrl.Service.GetShopDashboard(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, dashboard)
}
