package repository

import "github.com/Roongjin/ChatApplication/src/backend/internal/model"

type Conversation interface {
	BaseRepo[model.Conversation]
}
