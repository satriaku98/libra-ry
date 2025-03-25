package repository

import (
	"libra-ry/internal/domain"

	"gorm.io/gorm"
)

type bukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) BukuRepository {
	return &bukuRepository{db: db}
}

func (r *bukuRepository) GetAll(judul string, penulis string, penerbit string, tahunTerbit int, limit int, offset int) ([]domain.Buku, int, error) {
	var books []domain.Buku
	var totalBooks int64
	query := r.db

	// Filter jika parameter tidak kosong
	if judul != "" {
		query = query.Where("judul ILIKE ?", "%"+judul+"%")
	}
	if penulis != "" {
		query = query.Where("penulis ILIKE ?", "%"+penulis+"%")
	}
	if penerbit != "" {
		query = query.Where("penerbit ILIKE ?", "%"+penerbit+"%")
	}
	if tahunTerbit != 0 {
		query = query.Where("tahun_terbit = ?", tahunTerbit)
	}

	err := query.Limit(limit).Offset(offset).Find(&books).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Model(&domain.Buku{}).Count(&totalBooks).Error
	return books, int(totalBooks), err
}

func (r *bukuRepository) GetByID(id uint) (*domain.Buku, error) {
	var book domain.Buku
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *bukuRepository) Create(buku *domain.Buku) error {
	return r.db.Create(buku).Error
}

func (r *bukuRepository) Update(buku *domain.Buku) error {
	return r.db.Save(buku).Error
}

func (r *bukuRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Buku{}, id).Error
}
