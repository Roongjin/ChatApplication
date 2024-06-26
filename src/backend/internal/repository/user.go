package repository

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/google/uuid"
)

type User interface {
	BaseRepo[model.User]
	CheckUserExistenceByName(ctx context.Context, username string) (bool, error)
	FindOneByName(ctx context.Context, username string) (*model.User, error)
	FindByNames(ctx context.Context, usernames ...string) ([]*model.User, error)
	FindOnlineUsers(ctx context.Context) ([]*model.User, error)
	FindAllUsers(ctx context.Context) ([]*model.User, error)
	ChangeOnlineStatusById(ctx context.Context, userId uuid.UUID, isOnline bool) error
	InitNewUser(ctx context.Context, user *model.User) error
}
