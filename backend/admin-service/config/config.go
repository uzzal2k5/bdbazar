package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type SuperAdminConfig struct {
	Name     string
	Username string
	Password string
	Email    string
	Mobile   string
}


// Config holds application configuration and DB instance
type Config struct {
	DBSource     string
	JWTSecret string
	Port      string
	APIKey    string
	Address   string
	DB        *gorm.DB
	SuperAdmin  SuperAdminConfig
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

    // 	Require Variables to Create Super User
    spName := mustGetEnv("SUPERUSER_NAME")
    spUser := mustGetEnv("SUPERUSER_USERNAME")
    spPass := mustGetEnv("SUPERUSER_PASSWORD")
    spEmail := mustGetEnv("SUPERUSER_EMAIL")
    spMobile := mustGetEnv("SUPERUSER_MOBILE")

	// Optional with default
	port := getEnv("PORT", "8087")
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
	MigrateDB(db)

	return Config{
		DBSource:     dsn,
		JWTSecret: jwtSecret,
		Port:      port,
		APIKey:    apiKey,
		Address:   address,
		DB:        db,
        // Nest superadmin config inside Config
		SuperAdmin: SuperAdminConfig{
			Name:     spName,
			Username: spUser,
			Password: spPass,
			Email:    spEmail,
			Mobile:   spMobile,
		},
	}
}

// GetAuthServiceURL fetches the auth service URL from env vars
func GetAuthServiceURL() string {
    return os.Getenv("AUTH_SERVICE_URL")
}

// GetShopServiceURL fetches the shop service URL from env vars
func GetShopServiceURL() string {
    return os.Getenv("SHOP_SERVICE_URL")
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
