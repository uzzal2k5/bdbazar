package routes

import (
    "admin-service/controllers"
    "admin-service/middleware"
    "github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(router *gin.Engine, adminController *controllers.AdminController) {
    admins := router.Group("/api/admins")
    admins.POST("/spadm/login", adminController.LoginSuperAdmin)
    admins.Use(middleware.AdminOnlyAuth())
    {
        admins.GET("", adminController.ListAdmins)
    	admins.GET("/:id", adminController.ListAdminByID)
    	admins.POST("/register", adminController.CreateAdmin)
    	admins.PUT("/:id", adminController.UpdateAdmin)
    	admins.DELETE("/:id", adminController.DeleteAdmin)

        admins.GET("/metrics", adminController.GetMetrics)
        admins.GET("/dashboard", adminController.Dashboard)
        admins.PATCH("/user/:id/block", adminController.BlockUser)
        admins.PATCH("/user/:id/approve", adminController.ApproveUser)
        admins.POST("/user/:id/reset-password", adminController.ResetAdminPassword)
        admins.DELETE("/user/:id", adminController.DeleteUser)

        admins.PATCH("/shop/:id/approve", adminController.ApproveShop)
        admins.PATCH("/shop/:id/block", adminController.BlockShop)

    }
    router.GET("/health", func(c *gin.Context) {
		c.String(200, "admin-service is healthy")
	})

}