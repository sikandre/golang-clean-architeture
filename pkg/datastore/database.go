package datastore

type Database interface {
	FindAll(out interface{}) error
	Close() error
}
