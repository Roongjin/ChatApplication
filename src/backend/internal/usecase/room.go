package usecase

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository"
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository/postgres"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RoomUseCase struct {
	RoomRepo repository.Room
	LinkRepo repository.Link
}

func NewRoomUseCase(db *bun.DB) *RoomUseCase {
	return &RoomUseCase{
		RoomRepo: postgres.NewRoomDB(db),
		LinkRepo: postgres.NewLinkDB(db),
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

func (r *RoomUseCase) GetRoomsByUserId(ctx context.Context, userId uuid.UUID) ([]*model.Room, error) {
	links, err := r.LinkRepo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	roomIds := []uuid.UUID{}
	for _, link := range links {
		if link.RoomId != model.BroadcastRoomId {
			roomIds = append(roomIds, link.RoomId)
		}
	}

	rooms, err := r.RoomRepo.FindByIds(ctx, roomIds...)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
