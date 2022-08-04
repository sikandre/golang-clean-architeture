package memdb

import (
	"cleanArch/internal/domain/model"
	"github.com/hashicorp/go-memdb"
)

type memDb struct {
	db *memdb.MemDB
}

func (m *memDb) FindAll() ([]*model.User, error) {
	txn := m.db.Txn(false)
	defer txn.Abort()

	var res []*model.User

	result, err := txn.Get("user", "id")

	for obj := result.Next(); obj != nil; obj = result.Next() {
		res = append(res, obj.(*model.User))
	}

	return res, err
}

func (m *memDb) Close() error {
	return nil
}
