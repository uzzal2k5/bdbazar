package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"shipment-service/config"
	"shipment-service/models"
	"shipment-service/repository"
	"shipment-service/services"
	"shipment-service/controllers"
	"shipment-service/routes"
	"shipment-service/seed"
)

func main() {
    // Load configuration and environment variables
    cfg := config.LoadConfig()

    // Connect to the database
    db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto migrate Order model
	if err := db.AutoMigrate(&models.Shipment{}); err != nil {
        log.Fatalf("❌ Auto migration failed: %v", err)
    }

    err = seed.SeedShipments(db)
    if err != nil {
        log.Fatalf("Failed to seed shipments: %v", err)
    }


    // Initialize repository, service, and controller
    shipmentRepo := repository.NewShipmentRepository(db)
	shipmentService := services.NewShipmentService(shipmentRepo)
	shipmentController := controllers.NewShipmentController(shipmentService)
    // Initialize Gin router
    router := gin.Default()

    // Register routes
    routes.RegisterShipmentRoutes(router, shipmentController)

    log.Printf("Starting Order Service on port %s", cfg.Port)

    // Start server
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
