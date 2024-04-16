package repository

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/google/uuid"
)

type Link interface {
	BaseRepo[model.UserRoomLink]
	FindByUserId(ctx context.Context, userId uuid.UUID) ([]*model.UserRoomLink, error)
	FindByRoomId(ctx context.Context, roomId uuid.UUID) ([]*model.UserRoomLink, error)
	CheckExistenceByRoomMemberIds(ctx context.Context, memberIds []uuid.UUID) (bool, error)
	FindRoomByRoomMemberIds(ctx context.Context, memberIds []uuid.UUID) (*model.Room, error)
}
