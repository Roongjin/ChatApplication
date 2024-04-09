package repository

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
)

type User interface {
	BaseRepo[model.User]
	CheckUserExistenceByName(ctx context.Context, username string) (bool, error)
	FindOneByName(ctx context.Context, username string) (*model.User, error)
	FindOnlineUsers(ctx context.Context) ([]*model.User, error)
}
