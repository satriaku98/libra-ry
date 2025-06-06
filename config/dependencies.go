package config

import (
	"libra-ry/internal/handler"
	"libra-ry/internal/repository"
	"libra-ry/internal/usecase"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Dependencies struct {
	BukuHandler    *handler.BukuHandler
	AuthHandler    *handler.AuthHandler
	UserHandler    *handler.UserHandler
	AnggotaHandler *handler.AnggotaHandler
}

func InitDependencies(db *gorm.DB, logger *zap.Logger) *Dependencies {
	// Initialize Buku
	bukuRepo := repository.NewBukuRepository(db)
	bukuUseCase := usecase.NewBukuUseCase(bukuRepo)
	bukuHandler := handler.NewBukuHandler(bukuUseCase)

	// Initialize Auth
	expiry, _ := strconv.Atoi(GetEnv("JWT_EXPIRY", "3"))
	secret := GetEnv("JWT_SECRET", "default_secret")
	authRepo := repository.NewUserRepository(db)
	authUseCase := usecase.NewAuthUseCase(authRepo)
	authHandler := handler.NewAuthHandler(authUseCase, expiry, secret)

	// Initialize User
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	// Initialize Anggota
	anggotaRepo := repository.NewAnggotaRepository(db)
	anggotaUseCase := usecase.NewAnggotaUseCase(anggotaRepo)
	anggotaHandler := handler.NewAnggotaHandler(anggotaUseCase)

	return &Dependencies{
		BukuHandler:    bukuHandler,
		AuthHandler:    authHandler,
		UserHandler:    userHandler,
		AnggotaHandler: anggotaHandler,
	}
}
