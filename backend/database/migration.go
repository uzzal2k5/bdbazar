// ------------------------------
// 5. database/migration.go
// ------------------------------

package database

import "bdbazar/models"

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Product{})
}
