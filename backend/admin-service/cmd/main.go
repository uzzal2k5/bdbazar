package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"admin-service/config"
	"admin-service/controllers"
	"admin-service/models"
	"admin-service/repository"
	"admin-service/routes"
	"admin-service/seed"
	"admin-service/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Auto-migrate models
	if err := db.AutoMigrate(&models.SuperAdmin{}, &models.Admin{}, &models.AdminActivityLog{}); err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}

	// Run seeders
	seed.SeedAdmin(db)
	seed.SeedAdminActivityLogs(db)

	// Initialize repo, service, and controller
	adminRepo := repository.NewAdminRepository(db)
	adminService := services.NewAdminService(adminRepo)
	adminController := controllers.NewAdminController(adminService)

    // 🚨 Create super admin on startup
	if err := adminService.SetupSuperAdmin(cfg.SuperAdmin); err != nil {
		log.Fatalf("❌ Super admin setup failed: %v", err)
	}

	// Start Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterAdminRoutes(router, adminController)

	log.Printf("🚀 Starting Admin Service on port %s", cfg.Port)

	// Run server
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
