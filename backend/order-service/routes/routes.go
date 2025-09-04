package routes

import (
	"order-service/controllers"
	"order-service/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(router *gin.Engine, orderController *controllers.OrderController) {
    protected := router.Group("/api/orders")
    protected.Use(middleware.RequireAuth())
	{
		protected.POST("/", orderController.CreateOrder)         // ✅ Create new order
		protected.GET("/buyer", orderController.GetBuyerOrders)
		protected.GET("/seller",orderController.GetSellerOrders)
		protected.PUT("/:id/ship", orderController.MarkOrderShipped)       // ✏️ Update existing order
		protected.DELETE("/:id", orderController.DeleteOrder)    // ❌ Delete order

	}

}
