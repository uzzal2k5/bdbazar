package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"payment-service/config"
	"payment-service/models"
	"payment-service/repository"
	"payment-service/services"
	"payment-service/controllers"
	"payment-service/routes"
	"payment-service/seed"
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
	if err := db.AutoMigrate(&models.Payment{}); err != nil {
        log.Fatalf("❌ Auto migration failed: %v", err)
    }
    seed.SeedPayments(db)
    seed.RandomSeedPayments(db)

    // Initialize repository, service, and controller
    paymentRepo := repository.NewPaymentRepository(db)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentController := controllers.NewPaymentController(paymentService)

    // Initialize Gin router
    router := gin.Default()

    // Register routes
    routes.RegisterPaymentRoutes(router, paymentController)

    log.Printf("Starting Order Service on port %s", cfg.Port)

    // Start server
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
