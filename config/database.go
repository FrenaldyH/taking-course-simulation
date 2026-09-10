package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the shared database connection used by the repository layer.
var DB *gorm.DB

// LoadEnv loads .env into the process environment.
// A missing file is not fatal: the variables may already be set by the system.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}
}

// getEnv returns an environment variable, or fallback when it is unset.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// ConnectDatabase opens the PostgreSQL connection and stores it in DB.
func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", ""),
		getEnv("DB_NAME", "alpro_db"),
		getEnv("DB_PORT", "5432"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n"+
			"Check that PostgreSQL is running and the credentials in .env are correct.", err)
	}

	DB = database
	log.Println("Database connected successfully")
}
