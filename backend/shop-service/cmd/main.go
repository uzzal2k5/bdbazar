package main

import (
    "log"
    "shop-service/config"
    "shop-service/controllers"
    "shop-service/models"
    "shop-service/repository"
    "shop-service/routes"
    "shop-service/services"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    cfg := config.LoadConfig()

    db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto migrate
    db.AutoMigrate(&models.Shop{})

    shopRepo := repository.NewShopRepository(db)
    shopService := services.NewShopService(shopRepo)
    shopController := controllers.NewShopController(shopService)

    route := gin.Default()

    routes.RegisterShopRoutes(route, shopController)

    log.Printf("Starting Shop Service on port %s", cfg.Port)
    route.Run(":" + cfg.Port)
}
