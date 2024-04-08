package chat

import (
	"github.com/uptrace/bun"

	"github.com/Roongjin/ChatApplication/src/backend/internal/usecase"
)

type Resolver struct {
	RoomUsecase         usecase.RoomUseCase
	LinkUsecase         usecase.LinkUseCase
	ConversationUsecase usecase.ConversationUseCase
	UserUsecase         usecase.UserUseCase
}

func NewResolver(db *bun.DB) *Resolver {
	return &Resolver{
		RoomUsecase:         *usecase.NewRoomUseCase(db),
		LinkUsecase:         *usecase.NewLinkUseCase(db),
		ConversationUsecase: *usecase.NewConversationUseCase(db),
		UserUsecase:         *usecase.NewUserUseCase(db),
	}
}
