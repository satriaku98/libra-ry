package usecase

import (
	"errors"
	"libra-ry/internal/domain"
	"libra-ry/internal/repository"

	"gorm.io/gorm"
)

type AnggotaUseCase interface {
	GetByUserID(userID uint) (*domain.Anggota, error)
	Create(anggota *domain.Anggota) error
	Update(anggota *domain.Anggota) error
	DeleteByUserID(userID uint) error
}

type anggotaUseCase struct {
	repo repository.AnggotaRepository
}

func NewAnggotaUseCase(repo repository.AnggotaRepository) AnggotaUseCase {
	return &anggotaUseCase{repo: repo}
}

func (u *anggotaUseCase) GetByUserID(userID uint) (*domain.Anggota, error) {
	return u.repo.GetByUserID(userID)
}

func (u *anggotaUseCase) Create(anggota *domain.Anggota) error {
	return u.repo.Create(anggota)
}

func (u *anggotaUseCase) Update(anggota *domain.Anggota) error {
	// Cek apakah data anggota sudah ada
	existing, err := u.repo.GetByUserID(anggota.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existing != nil {
		// Jika sudah ada, gunakan ID dari data lama untuk update
		anggota.ID = existing.ID
		return u.repo.Update(anggota)
	}

	// Jika belum ada, buat data baru
	return u.repo.Create(anggota)
}

func (u *anggotaUseCase) DeleteByUserID(userID uint) error {
	return u.repo.DeleteByUserID(userID)
}
