package repository

import "libra-ry/internal/domain"

type BukuRepository interface {
	GetAll(title string, author string, publisher string, year int, limit int, offset int) ([]domain.Buku, int, error)
	GetByID(id uint) (*domain.Buku, error)
	Create(buku *domain.Buku) error
	Update(buku *domain.Buku) error
	Delete(id uint) error
}
