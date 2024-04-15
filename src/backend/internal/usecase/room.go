package usecase

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository"
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository/postgres"
	"github.com/uptrace/bun"
)

type RoomUseCase struct {
	RoomRepo repository.Room
}

func NewRoomUseCase(db *bun.DB) *RoomUseCase {
	return &RoomUseCase{
		RoomRepo: postgres.NewRoomDB(db),
	}
}

func (r *RoomUseCase) InitBroadcastRoom(ctx context.Context) error {
	exist, err := r.RoomRepo.CheckBroadcastRoomExistence(ctx)
	if err != nil {
		return err
	}

	if !exist {
		if err := r.RoomRepo.CreateBroadcastRoom(ctx); err != nil {
			return err
		}
	}

	bcstRoom, err := r.RoomRepo.GetBroadcastRoom(ctx)
	if err != nil {
		return err
	}

	model.BroadcastRoomId = bcstRoom.Id
	return nil
}
