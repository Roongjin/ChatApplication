package postgres

import (
	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/uptrace/bun"
)

type UserDB struct {
	*BaseDB[model.User]
}

func NewUserDB(db *bun.DB) *UserDB {
	type T = model.User

	return &UserDB{
		BaseDB: NewBaseDB[T](db),
	}
}