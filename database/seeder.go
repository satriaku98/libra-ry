package database

import (
	"encoding/json"
	"libra-ry/internal/domain"
	"libra-ry/pkg"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SeedAdmin memastikan admin pertama kali ada di database
func SeedAdmin(db *gorm.DB, log *zap.Logger, defaultAdminUser, defaultAdminPassword string) {
	var count int64
	db.Model(&domain.User{}).Where("role = ?", "admin").Count(&count)

	if count == 0 { // Jika belum ada admin, buat admin baru
		hashedPassword, err := pkg.HashPassword(defaultAdminPassword)
		if err != nil {
			log.Fatal("Failed to hash password: " + err.Error())
		}

		permissions, _ := json.Marshal([]string{"buku_read", "buku_write", "user_read", "user_write"})
		admin := domain.User{
			Username:    defaultAdminUser,
			Password:    string(hashedPassword),
			Role:        "admin",
			Permissions: permissions,
		}
		if err := db.Create(&admin).Error; err != nil {
			log.Fatal("Failed to create admin: " + err.Error())
		}

		log.Info("Admin user created successfully!")
	} else {
		log.Info("Admin user already exists, skipping seeding...")
	}
}
