package usecase

import (
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
