package postgres

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ConversationDB struct {
	*BaseDB[model.Conversation]
}

func NewConversationDB(db *bun.DB) *ConversationDB {
	type T = model.Conversation

	return &ConversationDB{
		BaseDB: NewBaseDB[T](db),
	}
}

func (c *ConversationDB) GetConversationsByRoomId(ctx context.Context, roomId uuid.UUID) ([]*model.Conversation, error) {
	var conversations []*model.Conversation
	if err := c.db.NewSelect().Model(&conversations).Where("room_id = ?", roomId).Scan(ctx, &conversations); err != nil {
		return nil, err
	}

	return conversations, nil
}
