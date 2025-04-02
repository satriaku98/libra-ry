package config

import (
	"fmt"
	"libra-ry/database"
	"libra-ry/internal/domain"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB menginisialisasi database
func NewDB(log *zap.Logger) *gorm.DB {
	// Ambil konfigurasi dari env
	dbUser := GetEnv("DB_USER", "default_user")
	dbPassword := GetEnv("DB_PASSWORD", "default_pass")
	dbName := GetEnv("DB_NAME", "default_db")
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "5432")
	dbSSLMode := GetEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: NewDatabaseLogger(log),
	})
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}

	// AutoMigrate
	err = db.AutoMigrate(&domain.Buku{}, &domain.User{})
	if err != nil {
		log.Fatal("Failed to migrate database", zap.Error(err))
	}

	// Seeder
	database.SeedAdmin(db, log, GetEnv("DEFAULT_ADMIN_USER", "admin"), GetEnv("DEFAULT_ADMIN_PASSWORD", "admin"))

	log.Info("Database connected and migrated")
	return db
}
