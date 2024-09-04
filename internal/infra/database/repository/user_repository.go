package repository

import (
	"first-project/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
