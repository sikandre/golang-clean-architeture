package datastore

type Database interface {
	FindAll() []interface{}
}
