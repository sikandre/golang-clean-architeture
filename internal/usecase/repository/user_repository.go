package repository

import (
	"cleanArch/internal/domain/model"
	"cleanArch/pkg/datastore"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
}

type userRepository struct {
	db datastore.Database
}

func NewUserRepository(db datastore.Database) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll() ([]*model.User, error) {
	users, err := ur.db.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}
