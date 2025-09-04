package config
// var ExternalServiceMap = map[string]string{
// 	"AuthRegister": "${AUTH_SERVICE_URL}/api/auth/v2/register",
// 	"AuthLogin":    "${AUTH_SERVICE_URL}/api/auth/v2/login",
// 	// Add more APIs from other services here
// }
//
//
//
// package controllers
//
// import (
// 	"admin-service/config"
// 	"admin-service/utils"
// 	"net/http"
//
// 	"github.com/gin-gonic/gin"
// )
//
// type AdminController struct{}
//
// func NewAdminController() *AdminController {
// 	return &AdminController{}
// }
//
// type RegisterPayload struct {
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }
//
// func (ctrl *AdminController) RegisterUserThroughAuthService(c *gin.Context) {
// 	var payload RegisterPayload
// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}
//
// 	url := config.ExternalServiceMap["AuthRegister"]
// 	resp, err := utils.PostJSON(url, payload)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to contact auth-service"})
// 		return
// 	}
//
// 	body, _ := utils.ReadBody(resp)
// 	c.Data(resp.StatusCode, "application/json", body)
// }
