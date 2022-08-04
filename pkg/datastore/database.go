package datastore

import "cleanArch/internal/domain/model"

type Database interface {
	FindAll() ([]*model.User, error)
	Close() error
}
