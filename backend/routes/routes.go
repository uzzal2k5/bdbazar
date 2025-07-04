// ------------------------------
// 10. routes/routes.go
// ------------------------------

package routes

import (
	"bdbazar/controllers"
	"bdbazar/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
    // Vendor-only route example
	vendorGroup := r.Group("/vendor")
	vendorGroup.Use(middleware.AuthorizeRole("vendor"))
	{
		vendorGroup.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome, Vendor!"})
		})
        vendorGroup.POST("/products", controllers.AddProduct)
	}

    // Admin-only route example
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.AuthorizeRole("admin"))
	{
		adminGroup.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome, Admin!"})
		})
	}

	// Customer-only route
	customerGroup := r.Group("/customer")
	customerGroup.Use(middleware.AuthorizeRole("customer"))
	{
	    customerGroup.GET("/profile", controllers.GetUserProfile)
// 		customerGroup.GET("/profile", func(c *gin.Context) {
// 			c.JSON(200, gin.H{"message": "Welcome, Customer!"})
// 		})
	}
    // Public product display route
	r.GET("/products", controllers.GetAllProducts)
	return r
}