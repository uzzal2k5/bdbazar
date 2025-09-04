package main

import (
	"fmt"
	"log"
	"os"
// 	"time"

	"auth-service/config"
	"auth-service/routes"
	"auth-service/controllers"
	"auth-service/services"
// 	"auth-service/middleware"
// 	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
    hash, _ := bcrypt.GenerateFromPassword([]byte("strongpassword123"), bcrypt.DefaultCost)
	fmt.Println("Hash:", string(hash))
    // Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
    address := os.Getenv("ADDRESS")
    if address == "" {
		address = ":8080" // Default address if not specified
	}
//  // Initialize Redis client
//     redisClient := redis.NewClient(&redis.Options{
//         Addr:     "localhost:6379", // your Redis address
//         Password: "",               // set password if any
//         DB:       0,                // default DB
//     })
//
//     // Test Redis connection
//     if err := redisClient.Ping(context.Background()).Err(); err != nil {
//         log.Fatalf("Failed to connect to Redis: %v", err)
//     }
    // Load config & DB
	cfg := config.LoadConfig()

    // Initialize the AuthService with config
	authService := services.NewAuthService(cfg)

    // Controller
    authController := controllers.NewAuthController(authService)

    // Setup Gin router
	router := gin.Default()
	routes.AuthRoutes(router, authController, cfg.JWTSecret)

    // Start the server
	fmt.Printf("Server is running at %s\n", cfg.Address)
	if err := router.Run(cfg.Address); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
