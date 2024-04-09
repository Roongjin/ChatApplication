package postgres

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type LinkDB struct {
	*BaseDB[model.UserRoomLink]
}

func NewLinkDB(db *bun.DB) *LinkDB {
	type T = model.UserRoomLink

	return &LinkDB{
		BaseDB: NewBaseDB[T](db),
	}
}

func (l *LinkDB) FindByUserId(ctx context.Context, userId uuid.UUID) ([]*model.UserRoomLink, error) {
	var subscriptions []*model.UserRoomLink
	if err := l.db.NewSelect().Model(&subscriptions).Where("user_id = ?", userId).Scan(ctx, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}
