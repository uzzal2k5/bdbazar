package routes

import (
    "shop-service/controllers"
    "shop-service/middleware"

    "github.com/gin-gonic/gin"
)

func RegisterShopRoutes(r *gin.Engine, shopController *controllers.ShopController) {
    shop := r.Group("/api/shops")
    {
        shop.GET("/", shopController.ListShops)        // List all approved & unblocked shops
        shop.GET("/search", shopController.SearchShops) // Search shops by name
        shop.GET("/:id", shopController.GetShop)       // Get shop by ID


        // Authenticated endpoints
        protected := shop.Group("/")
        protected.Use(middleware.RequireAuth()) // Example middleware to protect routes
        {
            protected.POST("/", shopController.CreateShop)    // Create new shop (seller)
            protected.PUT("/:id", shopController.UpdateShop)  // Update shop
            protected.DELETE("/:id", shopController.DeleteShop) // Delete shop
            protected.GET("/dashboard", shopController.GetDashboard)
        }
    }
}
