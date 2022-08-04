package memdb

import (
	"cleanArch/internal/domain/model"
	"cleanArch/pkg/datastore"
	"github.com/hashicorp/go-memdb"
	"log"
	"time"
)

func NewDb() datastore.Database {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"user": &memdb.TableSchema{
				Name: "user",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"age": &memdb.IndexSchema{
						Name:    "age",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Age"},
					},
					"createdAt": &memdb.IndexSchema{
						Name:    "createdAt",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "CreatedAt"},
					},
					"updatedAt": &memdb.IndexSchema{
						Name:    "updatedAt",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "UpdatedAt"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Fatalln(err)
	}

	user := &model.User{
		Id:        1,
		Name:      "Bob",
		Age:       "30",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	txn := db.Txn(true)

	err = txn.Insert("user", user)
	if err != nil {
		log.Fatalln(err)
	}

	txn.Commit()

	return &memDb{db: db}

}
