package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"auth-service/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config holds application configuration and DB instance
type Config struct {
	DBUrl     string
	JWTSecret string
	Port      string
	APIKey    string
	Address   string
	DB        *gorm.DB
}

// LoadConfig loads environment variables, connects to DB, and returns config
func LoadConfig() Config {
	// Allow specifying a custom .env file
	envFile := flag.String("envFile", ".env", "Path to the .env file")
	flag.Parse()

	log.Printf("Loading env file from: %s", *envFile)
	if err := godotenv.Load(*envFile); err != nil {
		log.Printf("⚠️ Warning: Error loading .env file: %v", err)
	}

	// Required variables
	host := mustGetEnv("DB_HOST")
	user := mustGetEnv("DB_USER")
	pass := mustGetEnv("DB_PASS")
	name := mustGetEnv("DB_NAME")
	dbPort := mustGetEnv("DB_PORT")
	jwtSecret := mustGetEnv("JWT_SECRET")
	apiKey := mustGetEnv("API_KEY")

	// Optional with default
	port := getEnv("PORT", "8080")
	address := getEnv("ADDRESS", ":"+port)

	// Construct DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, pass, name, dbPort,
	)

	// Connect to DB
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to PostgreSQL: %v", err)
	}
	log.Println("✅ Connected to PostgreSQL successfully")

	// Run auto-migration
	migrateDB(db)

	return Config{
		DBUrl:     dsn,
		JWTSecret: jwtSecret,
		Port:      port,
		APIKey:    apiKey,
		Address:   address,
		DB:        db,
	}
}

// migrateDB auto-migrates DB tables
func migrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{},&models.RefreshToken{})
	if err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}
	log.Println("✅ Database migration completed")
}

// mustGetEnv fetches required environment variable or logs fatal error
func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("❌ Required environment variable %s not set", key)
	}
	return value
}

// getEnv fetches env variable with fallback
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
