package repository

import "github.com/Roongjin/ChatApplication/src/backend/internal/model"

type Room interface {
	BaseRepo[model.Room]
}
