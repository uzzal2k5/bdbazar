package routes

import (
    "shipment-service/controllers"
    "shipment-service/middleware"

    "github.com/gin-gonic/gin"
)

func RegisterShipmentRoutes(router *gin.Engine, shipmentController *controllers.ShipmentController) {
    shipment := router.Group("/api/shipments")
    shipment.Use(middleware.RequireAuth())
    {
        // Buyer endpoints
        shipment.POST("/",shipmentController.CreateShipment)                   
        shipment.GET("/seller",shipmentController.GetShipmentsBySeller)          
        shipment.GET("/order/:order_id",shipmentController.GetShipmentByOrder)        
        shipment.PUT("/:id/status",shipmentController.UpdateStatus)     
        shipment.DELETE("/:id",shipmentController.DeleteShipment)     


    }

}
