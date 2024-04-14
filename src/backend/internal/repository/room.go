package repository

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
)

type Room interface {
	BaseRepo[model.Room]
	InitBroadcastRoom(ctx context.Context) error
}
