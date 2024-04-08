package controller

import (
	"github.com/Roongjin/ChatApplication/src/backend/internal/controller/chat"
	"github.com/uptrace/bun"
)

type Handler struct {
	Chat chat.Resolver
}

func NewHandler(db *bun.DB) *Handler {
	return &Handler{
		Chat: *chat.NewResolver(db),
	}
}
