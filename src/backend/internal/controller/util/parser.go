package util

import (
	"context"
	"time"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/Roongjin/ChatApplication/src/backend/internal/usecase"
	"github.com/google/uuid"
)

func ParseConversationToMessage(ctx context.Context, userUsecase usecase.UserUseCase, conversations []*model.Conversation) ([]*model.Message, error) {
	seenUserId := make(map[uuid.UUID]bool)
	allUserIds := []uuid.UUID{}
	for _, conv := range conversations {
		exist := seenUserId[conv.UserId]
		if !exist {
			allUserIds = append(allUserIds, conv.UserId)
			seenUserId[conv.UserId] = true
		}
	}

	allUsers, err := userUsecase.UserRepo.FindByIds(ctx, allUserIds...)
	if err != nil {
		return nil, err
	}

	userIdMapping := make(map[uuid.UUID]*model.User)
	for _, user := range allUsers {
		userIdMapping[user.Id] = user
	}

	messages := []*model.Message{}
	for _, conv := range conversations {
		messages = append(messages, &model.Message{
			Id:         conv.Id,
			Text:       conv.Text,
			SenderName: userIdMapping[conv.UserId].Username,
			Timestamp:  conv.CreatedAt.Add(time.Hour * 7),
		})
	}

	return messages, nil
}
