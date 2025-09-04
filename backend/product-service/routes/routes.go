package routes

import (
	"product-service/controllers"
	"product-service/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, productController *controllers.ProductController) {
	// Public routes
	product := r.Group("/api/products")
	{
		product.GET("/", productController.GetAll)                    // 📦 List all products
		product.GET("/:id", productController.GetByID)                // 🔍 Get product by ID
		product.GET("/search", productController.SearchProduct)       // 🔍 Search product by name or ID
	}

	// Protected routes (seller only)
	protected := r.Group("/api/products")
	protected.Use(middleware.RequireAuth())
	{
		protected.POST("/", productController.CreateProduct)         // ✅ Create new product
		protected.PUT("/:id", productController.UpdateProduct)       // ✏️ Update existing product
		protected.DELETE("/:id", productController.DeleteProduct)    // ❌ Delete product
		protected.POST("/adjust-stock", productController.AdjustStock) // 🔧 Adjust stock (internal/seller)
	}
}
