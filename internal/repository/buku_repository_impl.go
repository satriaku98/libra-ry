package repository

import (
	"encoding/json"
	"libra-ry/internal/domain"

	"gorm.io/gorm"
)

type bukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) BukuRepository {
	return &bukuRepository{db: db}
}

func (r *bukuRepository) GetAll(judul string, penulis string, penerbit string, tahunTerbit int, limit int, offset int, tags []string, sortBy []string) ([]domain.Buku, int, error) {
	var books []domain.Buku
	var totalBooks int64

	validSortFields := map[string]bool{
		"judul":        true,
		"penulis":      true,
		"penerbit":     true,
		"tahun_terbit": true,
		"stok":         true,
	}

	query := r.db.Model(&domain.Buku{}).Where("deleted_at IS NULL")

	// Filter
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
	if len(tags) > 0 {
		permJSON, err := json.Marshal(tags)
		if err != nil {
			return nil, 0, err
		}
		query = query.Where("tags @> ?", permJSON)
	}

	// Sorting
	for _, field := range sortBy {
		fieldName := field
		direction := "ASC"
		if len(field) > 0 && field[0] == '-' {
			fieldName = field[1:]
			direction = "DESC"
		}

		if validSortFields[fieldName] {
			query = query.Order(fieldName + " " + direction)
		}
	}

	// Count & Find
	err := query.Count(&totalBooks).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Limit(limit).Offset(offset).Find(&books).Error
	return books, int(totalBooks), err
}

func (r *bukuRepository) GetByID(id uint) (*domain.Buku, error) {
	var book domain.Buku
	err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&book).Error
	return &book, err
}

func (r *bukuRepository) Create(buku *domain.Buku) error {
	return r.db.Create(buku).Error
}

func (r *bukuRepository) Update(buku *domain.Buku) error {
	return r.db.Model(&domain.Buku{}).Where("id = ? AND deleted_at IS NULL", buku.ID).Updates(buku).Error
}

func (r *bukuRepository) Delete(id uint) error {
	// soft delete
	return r.db.Where("id = ?", id).Delete(&domain.Buku{}).Error
}
