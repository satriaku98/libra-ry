package repository

import (
	"libra-ry/internal/domain"

	"gorm.io/gorm"
)

type anggotaRepository struct {
	db *gorm.DB
}

func NewAnggotaRepository(db *gorm.DB) AnggotaRepository {
	return &anggotaRepository{db: db}
}

func (r *anggotaRepository) GetByUserID(userID uint) (*domain.Anggota, error) {
	var anggota domain.Anggota
	err := r.db.Where("user_id = ?", userID).First(&anggota).Error
	if err != nil {
		return nil, err
	}
	return &anggota, nil
}

func (r *anggotaRepository) Create(anggota *domain.Anggota) error {
	return r.db.Create(anggota).Error
}

func (r *anggotaRepository) Update(anggota *domain.Anggota) error {
	return r.db.Save(anggota).Error
}

func (r *anggotaRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&domain.Anggota{}).Error
}
