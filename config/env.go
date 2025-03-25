package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv membaca file .env dan menyimpan ke variabel environment
func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}
}

// GetEnv mengambil nilai dari environment variable dengan default value jika kosong
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
