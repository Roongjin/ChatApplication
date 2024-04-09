package usecase

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository"
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository/postgres"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserUseCase struct {
	UserRepo repository.User
}

func NewUserUseCase(db *bun.DB) *UserUseCase {
	return &UserUseCase{
		UserRepo: postgres.NewUserDB(db),
	}
}

func (u *UserUseCase) MustGetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	existed, err := u.UserRepo.CheckUserExistenceByName(ctx, username)
	if err != nil {
		return nil, err
	}

	if existed {
		user, err := u.UserRepo.FindOneByName(ctx, username)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

	newUser := &model.User{
		Id:       uuid.New(),
		Username: username,
		IsOnline: false,
	}

	if err := u.UserRepo.AddOne(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}
