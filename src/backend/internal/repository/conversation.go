package repository

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/google/uuid"
)

type Conversation interface {
	BaseRepo[model.Conversation]
	GetConversationsByRoomId(ctx context.Context, roomId uuid.UUID) ([]*model.Conversation, error)
}
