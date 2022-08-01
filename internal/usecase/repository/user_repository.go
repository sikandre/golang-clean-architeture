package repository

import (
	"cleanArch/internal/domain/model"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	FindAll(users []*model.User) ([]*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}
