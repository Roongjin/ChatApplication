package postgres

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/uptrace/bun"
)

type UserDB struct {
	*BaseDB[model.User]
}

func NewUserDB(db *bun.DB) *UserDB {
	type T = model.User

	return &UserDB{
		BaseDB: NewBaseDB[T](db),
	}
}

func (u *UserDB) CheckUserExistenceByName(ctx context.Context, username string) (bool, error) {
	var user model.User
	exist, err := u.db.NewSelect().Model(&user).Where("username = ?", username).Exists(ctx)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (u *UserDB) FindOneByName(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := u.db.NewSelect().Model(&user).Where("username = ?", username).Scan(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
