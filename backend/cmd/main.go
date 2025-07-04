// ------------------------------
// 11. cmd/main.go
// ------------------------------
// ------------------------------
// Swagger Setup (optional - basic)
// ------------------------------

// Add swaggo packages:
// go get -u github.com/swaggo/gin-swagger
// go get -u github.com/swaggo/files

// In main.go, add:
// import _ "bdbazar/docs"
// import ginSwagger "github.com/swaggo/gin-swagger"
// import swaggerFiles "github.com/swaggo/files"
// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// Then run:
// swag init  (you’ll need to document your handlers with comments)

// Example controller annotation:
// @Summary Login
// @Description Logs in a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body models.User true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]

// ------------------------------
// Done ✅
// ----------------------
package main

import (
	"bdbazar/config"
	"bdbazar/database"
	"bdbazar/routes"
	"os"
)

func main() {
	config.LoadEnv()
	database.Connect()
	database.Migrate()
	r := routes.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
