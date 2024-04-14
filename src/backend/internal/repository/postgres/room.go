package postgres

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/google/uuid"
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

func (r *RoomDB) InitBroadcastRoom(ctx context.Context) error {
	var bcstRoom model.BroadcastRoom
	count, err := r.db.NewSelect().Model(&bcstRoom).Count(ctx)
	if err != nil {
		return err
	}

	if count == 0 {
		bcstRoom.Id = uuid.New()
		if _, err := r.db.NewInsert().Model(&bcstRoom).Exec(ctx); err != nil {
			return err
		}

		pttRoom := model.Room{
			Id: bcstRoom.Id,
		}
		if _, err := r.db.NewInsert().Model(&pttRoom).Exec(ctx); err != nil {
			return err
		}
	}

	if err := r.db.NewSelect().Model(&bcstRoom).Scan(ctx, &bcstRoom); err != nil {
		return err
	}

	model.BroadcastRoomId = bcstRoom.Id
	return nil
}
