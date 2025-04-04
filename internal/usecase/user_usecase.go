package usecase

import (
	"encoding/json"
	"errors"
	"libra-ry/internal/domain"
	"libra-ry/internal/repository"
	"libra-ry/pkg"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	GetAll(username string, role string, permissions []string, limit int, offset int) ([]domain.User, int, error)
	GetByID(id uint) (*domain.User, error)
	Create(user *domain.User) error
	ChangePassword(userID uint, oldPassword, newPassword string) error
	Update(user *domain.User) error
	Delete(id uint) error
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) GetAll(username, role string, permissions []string, limit int, offset int) ([]domain.User, int, error) {
	// Filter permissions kosong yang tidak valid
	cleaned := []string{}
	for _, p := range permissions {
		if p != "" {
			cleaned = append(cleaned, p)
		}
	}
	return u.repo.GetAll(username, role, cleaned, limit, offset)
}

func (u *userUseCase) GetByID(id uint) (*domain.User, error) {
	return u.repo.GetByID(id)
}

func (u *userUseCase) Create(user *domain.User) error {
	permissionsJSON, err := json.Marshal(user.Permissions)
	if err != nil {
		return err
	}
	user.Permissions = permissionsJSON

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return u.repo.Create(user)
}

func (u *userUseCase) ChangePassword(userID uint, oldPassword, newPassword string) error {
	user, err := u.repo.GetByID(userID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("old password is incorrect")
	}

	hashed, err := pkg.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user.Password = hashed
	return u.repo.Update(user)
}

func (u *userUseCase) Update(user *domain.User) error {
	permissionsJSON, err := json.Marshal(user.Permissions)
	if err != nil {
		return err
	}
	user.Permissions = permissionsJSON

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		return err
	}

	return u.repo.Update(user)
}

func (u *userUseCase) Delete(id uint) error {
	return u.repo.Delete(id)
}
