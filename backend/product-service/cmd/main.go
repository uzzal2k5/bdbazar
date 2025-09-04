package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

	"product-service/config"
	"product-service/models"
	"product-service/repository"
	"product-service/services"
	"product-service/controllers"
	"product-service/routes"
)

func main() {
    // Load configuration and environment variables
    cfg := config.LoadConfig()

    // Connect to the database
    db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto migrate Product model
	if err := db.AutoMigrate(&models.Product{}, &models.Category{}); err != nil {
        log.Fatalf("❌ Auto migration failed: %v", err)
    }

    // Initialize repository, service, and controller
    productRepo := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

    // Initialize Gin router
    router := gin.Default()

    // Register routes
    routes.RegisterProductRoutes(router, productController)

    log.Printf("Starting Product Service on port %s", cfg.Port)

    // Start server
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
