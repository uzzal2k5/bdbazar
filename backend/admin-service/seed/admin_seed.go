package seed

import (
	"log"
	"time"

	"admin-service/models"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) {
	var count int64
	db.Model(&models.Admin{}).Count(&count)
	if count == 0 {
		admins := []models.Admin{
        	{
        		Name:     "Muskan Admin",
        		Username: "muskan",
        		Password: "securepassword",
        		ID:   1,
        		Email:    "muskan@example.com",
        		Mobile:   "01710000001",
        		Status:   "",
        		Role:     "admin",
        		CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
        	},
        	{
        		Name:     "Test Admin",
        		Username: "test",
        		Password: "test123",
        		ID:   2,
        		Email:    "test@example.com",
        		Mobile:   "01710000002",
        		Status:   "",
        		Role:     "admin",
        		CreatedAt: time.Now(),
                UpdatedAt: time.Now(),
        	},
        }

		if err := db.Create(&admins).Error; err != nil {
			log.Printf("❌ Failed to seed admins: %v", err)
		} else {
			log.Println("✅ Admins seeded successfully.")
		}
	} else {
		log.Println("ℹ️ Admins already seeded. Skipping...")
	}
}

func SeedAdminActivityLogs(db *gorm.DB) {
	var count int64
	db.Model(&models.AdminActivityLog{}).Count(&count)
	if count == 0 {
		logs := []models.AdminActivityLog{
			{
				AdminID:  1,
				Action:   "Initial login",
				LoggedAt: time.Now(),
			},
			{
				AdminID:  1,
				Action:   "Created a new shop",
				LoggedAt: time.Now(),
			},
			{
				AdminID:  2,
				Action:   "Approved a user request",
				LoggedAt: time.Now(),
			},
		}
		if err := db.Create(&logs).Error; err != nil {
			log.Printf("❌ Failed to seed admin activity logs: %v", err)
		} else {
			log.Println("✅ Admin activity logs seeded successfully.")
		}
	} else {
		log.Println("ℹ️ Admin activity logs already seeded. Skipping...")
	}
}
