package postgres

import (
	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/uptrace/bun"
)

type LinkDB struct {
	*BaseDB[model.UserRoomLink]
}

func NewLinkDB(db *bun.DB) *LinkDB {
	type T = model.UserRoomLink

	return &LinkDB{
		BaseDB: NewBaseDB[T](db),
	}
}
