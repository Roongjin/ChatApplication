package postgres

import (
	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/uptrace/bun"
)

type RoomDB struct {
	*BaseDB[model.Room]
}

func NewRoomDB(db *bun.DB) *RoomDB {
	type T = model.Room

	return &RoomDB{
		BaseDB: NewBaseDB[T](db),
	}
}
