// ------------------------------
// controllers/product_controller.go
// ------------------------------

package controllers

import (
	"bdbazar/database"
	"bdbazar/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
    role := c.MustGet("role").(string)
	if role != "vendor" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only vendors can add products"})
		return
	}
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    email := c.MustGet("email").(string)
	var user models.User
	database.DB.Where("email = ?", email).First(&user)
	product.VendorID = user.ID

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}
