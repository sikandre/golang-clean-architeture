package gorm

import (
	"cleanArch/internal/domain/model"
	"github.com/jinzhu/gorm"
)

type gormDB struct {
	*gorm.DB
}

func (g *gormDB) FindAll() ([]*model.User, error) {
	var out []*model.User

	g.DB.Find(&out)

	return out, g.Error
}

func (g *gormDB) Close() error {
	return g.DB.Close()
}
