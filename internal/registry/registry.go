package registry

import (
	"cleanArch/pkg/datastore"
	"cleanArch/pkg/server"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

func NewRegistry() {
	db := datastore.NewDB()

	// Migrate the schema
	// Create
	//db.Create(&model.User{
	//	Name:      "Bob",
	//	Age:       "10",
	//	CreatedAt: time.Time{},
	//	UpdatedAt: time.Time{},
	//})
	defer db.Close()

	r := registry{db: db}

	server.NewHttpServer(r.NewAppController())
}
