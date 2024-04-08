package repository

import "github.com/Roongjin/ChatApplication/src/backend/internal/model"

type User interface {
	BaseRepo[model.User]
}
