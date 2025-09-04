package routes

import (
    "payment-service/controllers"
    "payment-service/middleware"

    "github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(router *gin.Engine, paymentController *controllers.PaymentController) {
    payment := router.Group("/api/payments")
    payment.Use(middleware.RequireAuth())
    {
        // Buyer endpoints
        payment.POST("/",paymentController.CreatePayment)                   // POST: Create a new payment (buyer)
        payment.GET("/buyer",paymentController.GetPaymentsByBuyer)          // GET: List payments made by buyer
        // Seller endpoints
        payment.GET("/seller",paymentController.GetPaymentsBySeller)        // GET: List payments received by seller
        payment.POST("/:id/complete",paymentController.CompletePayment)     // PUT: Mark payment as completed (seller)


    }

}
