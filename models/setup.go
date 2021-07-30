package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB the db instance exported
var DB *gorm.DB

// getEnv helper function for fetching OS variables
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// ConnectDatabase function to connect to a pgsql database
func ConnectDatabase() {
	pgUser := getEnv("POSTGRES_USER", "fizzbuzz")
	pgPWD := getEnv("POSTGRES_PASSWORD", "fizzbuzz")
	pgDB := getEnv("POSTGRES_DB", "fizzbuzz")
	pgHost := getEnv("POSTGRES_HOST", "localhost")
	pgPort := getEnv("POSTGRES_PORT", "5432")
	dsn := "host=" + pgHost + " user=" + pgUser + " password=" + pgPWD + " dbname=" + pgDB + " port=" + pgPort
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Stats{})

	DB = database
}
