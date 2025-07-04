// ------------------------------
// controllers/user_controller.go
// ------------------------------

package controllers

import (
	"bdbazar/database"
	"bdbazar/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	email := c.MustGet("email").(string)
	var user models.User
	database.DB.Where("email = ?", email).First(&user)
	c.JSON(http.StatusOK, user)
}