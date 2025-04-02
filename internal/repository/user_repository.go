package repository

import "libra-ry/internal/domain"

type UserRepository interface {
	GetAll(username string, role string, permissions []string, limit int, offset int) ([]domain.User, int, error)
	GetByUsername(username string) (*domain.User, error)
	GetByID(id uint) (*domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id uint) error
}
