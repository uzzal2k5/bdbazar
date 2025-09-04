package config

import (
	"fmt"
	"shipment-service/models"
	"log"

	"gorm.io/gorm"
)

// migrateDB performs auto migration for all models
func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.Shipment{},
	)

	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	fmt.Println("✅ Database migration completed successfully!")
}
