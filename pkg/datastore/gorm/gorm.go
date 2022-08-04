package gorm

import (
	"github.com/jinzhu/gorm"
)

type gormDB struct {
	*gorm.DB
}

func (g *gormDB) FindAll(out interface{}) error {
	g.DB.Find(out)

	return g.Error
}

func (g *gormDB) Close() error {
	return g.DB.Close()
}
