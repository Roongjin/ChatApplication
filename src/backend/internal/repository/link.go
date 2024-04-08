package repository

import "github.com/Roongjin/ChatApplication/src/backend/internal/model"

type Link interface {
	BaseRepo[model.UserRoomLink]
}
