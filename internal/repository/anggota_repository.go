package repository

import "libra-ry/internal/domain"

type AnggotaRepository interface {
	GetByUserID(userID uint) (*domain.Anggota, error)
	Create(anggota *domain.Anggota) error
	Update(anggota *domain.Anggota) error
	DeleteByUserID(userID uint) error
}
