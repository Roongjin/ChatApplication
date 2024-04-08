package usecase

import (
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository"
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository/postgres"
	"github.com/uptrace/bun"
)

type ConversationUseCase struct {
	ConversationRepo repository.Conversation
}

func NewConversationUseCase(db *bun.DB) *ConversationUseCase {
	return &ConversationUseCase{
		ConversationRepo: postgres.NewConversationDB(db),
	}
}
