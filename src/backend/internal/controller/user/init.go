package user

import (
	"github.com/Roongjin/ChatApplication/src/backend/internal/usecase"
	"github.com/uptrace/bun"
)

type Resolver struct {
	UserUsecase usecase.UserUseCase
}

func NewResolver(db *bun.DB) *Resolver {
	return &Resolver{
		UserUsecase: *usecase.NewUserUseCase(db),
	}
}
