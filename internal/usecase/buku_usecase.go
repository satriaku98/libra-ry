package usecase

import (
	"libra-ry/internal/domain"
	"libra-ry/internal/repository"
)

type BukuUseCase interface {
	GetAllBooks(page int, title string, author string, publisher string, year int) ([]domain.Buku, int, error)
	GetBookByID(id uint) (*domain.Buku, error)
	CreateBook(buku *domain.Buku) error
	UpdateBook(buku *domain.Buku) error
	DeleteBook(id uint) error
}

type bukuUseCase struct {
	repo repository.BukuRepository
}

func NewBukuUseCase(repo repository.BukuRepository) BukuUseCase {
	return &bukuUseCase{repo: repo}
}

func (uc *bukuUseCase) GetAllBooks(page int, title string, author string, publisher string, year int) ([]domain.Buku, int, error) {
	limit := 100
	offset := (page - 1) * limit
	return uc.repo.GetAll(title, author, publisher, year, limit, offset)
}

func (uc *bukuUseCase) GetBookByID(id uint) (*domain.Buku, error) {
	return uc.repo.GetByID(id)
}

func (uc *bukuUseCase) CreateBook(buku *domain.Buku) error {
	return uc.repo.Create(buku)
}

func (uc *bukuUseCase) UpdateBook(buku *domain.Buku) error {
	return uc.repo.Update(buku)
}

func (uc *bukuUseCase) DeleteBook(id uint) error {
	return uc.repo.Delete(id)
}
