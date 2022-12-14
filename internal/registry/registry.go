package registry

import (
	"cleanArch/pkg/datastore"
	"cleanArch/pkg/datastore/memdb"
	"cleanArch/pkg/server"
)

type registry struct {
	db datastore.Database
}

func NewRegistry() {
	db := memdb.NewDb()
	//db := gorm.NewDB()

	defer db.Close()

	r := registry{db: db}

	server.NewHttpServer(r.NewAppController())
}
