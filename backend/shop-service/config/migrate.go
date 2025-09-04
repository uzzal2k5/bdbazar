package config

import (
	"fmt"
	"shop-service/models"
	"log"

	"gorm.io/gorm"
)

// MigrateDB performs auto migration for all models
func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Shop{},

	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Database migration completed successfully!")
}
