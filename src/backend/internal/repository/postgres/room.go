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

func (r *RoomDB) CheckBroadcastRoomExistence(ctx context.Context) (bool, error) {
	var bcstRoom model.BroadcastRoom
	count, err := r.db.NewSelect().Model(&bcstRoom).Count(ctx)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}

func (r *RoomDB) CreateBroadcastRoom(ctx context.Context) error {
	bcstRoom := model.BroadcastRoom{
		Id: uuid.New(),
	}
	if _, err := r.db.NewInsert().Model(&bcstRoom).Exec(ctx); err != nil {
		return err
	}

	pttRoom := model.Room{
		Id: bcstRoom.Id,
	}
	if _, err := r.db.NewInsert().Model(&pttRoom).Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *RoomDB) GetBroadcastRoom(ctx context.Context) (*model.BroadcastRoom, error) {
	var bcstRoom model.BroadcastRoom
	if err := r.db.NewSelect().Model(&bcstRoom).Scan(ctx, &bcstRoom); err != nil {
		return nil, err
	}

	return &bcstRoom, nil
}
