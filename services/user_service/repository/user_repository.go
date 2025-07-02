package repository

import (
	"user_service/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	if err := r.db.Debug().Create(user).Error; err != nil {
		return err
	}
	return nil
}
