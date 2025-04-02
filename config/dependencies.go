package config

import (
	"libra-ry/internal/handler"
	"libra-ry/internal/repository"
	"libra-ry/internal/usecase"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Dependencies struct {
	BukuHandler *handler.BukuHandler
	// MemberHandler *handler.MemberHandler
}

func InitDependencies(db *gorm.DB, logger *zap.Logger) *Dependencies {
	// Initialize Buku
	bukuRepo := repository.NewBukuRepository(db)
	bukuUseCase := usecase.NewBukuUseCase(bukuRepo)
	bukuHandler := handler.NewBukuHandler(bukuUseCase)

	// Initialize Member
	// memberRepo := repository.NewMemberRepository(db)
	// memberUseCase := usecase.NewMemberUseCase(memberRepo)
	// memberHandler := handler.NewMemberHandler(memberUseCase)

	return &Dependencies{
		BukuHandler: bukuHandler,
		// MemberHandler: memberHandler,
	}
}
