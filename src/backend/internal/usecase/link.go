package usecase

import (
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository"
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository/postgres"
	"github.com/uptrace/bun"
)

type LinkUseCase struct {
	LinkRepo repository.Link
}

func NewLinkUseCase(db *bun.DB) *LinkUseCase {
	return &LinkUseCase{
		LinkRepo: postgres.NewLinkDB(db),
	}
}
