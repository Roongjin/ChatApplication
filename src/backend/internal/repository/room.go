package repository

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
)

type Room interface {
	BaseRepo[model.Room]
	CheckBroadcastRoomExistence(ctx context.Context) (bool, error)
	CreateBroadcastRoom(ctx context.Context) error
	GetBroadcastRoom(ctx context.Context) (*model.BroadcastRoom, error)
}
