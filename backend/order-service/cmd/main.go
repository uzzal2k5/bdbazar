package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"order-service/config"
	"order-service/models"
	"order-service/repository"
	"order-service/services"
	"order-service/controllers"
	"order-service/routes"
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
	if err := db.AutoMigrate(&models.Order{}, &models.OrderItem{}); err != nil {
        log.Fatalf("❌ Auto migration failed: %v", err)
    }

    // Initialize repository, service, and controller
    orderRepo := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

    // Initialize Gin router
    router := gin.Default()

    // Register routes
    routes.RegisterOrderRoutes(router, orderController)

    log.Printf("Starting Order Service on port %s", cfg.Port)

    // Start server
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
