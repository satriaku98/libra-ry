package repository

import (
	"encoding/json"
	"libra-ry/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll(username string, role string, permissions []string, limit int, offset int) ([]domain.User, int, error) {
	var users []domain.User
	var totalUsers int64
	query := r.db

	// Filter jika parameter tidak kosong
	if username != "" {
		query = query.Where("username ILIKE ?", "%"+username+"%")
	}
	if role != "" {
		query = query.Where("role ILIKE ?", "%"+role+"%")
	}
	if len(permissions) > 0 {
		permJSON, err := json.Marshal(permissions)
		if err != nil {
			return nil, 0, err
		}
		query = query.Where("permissions @> ?", permJSON)
	}

	err := query.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Model(&domain.User{}).Count(&totalUsers).Error
	return users, int(totalUsers), err
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
