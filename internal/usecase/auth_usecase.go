package usecase

import (
	"errors"
	"libra-ry/internal/domain"
	"libra-ry/internal/repository"
	"libra-ry/pkg"
)

type AuthUseCase struct {
	userRepo repository.UserRepository
}

// NewAuthUseCase membuat instance baru dari AuthUseCase
func NewAuthUseCase(userRepo repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

// Login melakukan proses autentikasi
func (u *AuthUseCase) Login(username, password string) (*domain.User, error) {
	user, err := u.userRepo.GetByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !pkg.CheckPassword(user.Password, password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
