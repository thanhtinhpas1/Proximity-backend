package resolvers

import (
	"context"
	"proximity/graph/model"
	"proximity/graph/repositories"
	"time"
)

type ChatResolver struct {
	ChatRepo *repositories.ChatRepo
}

func (rsv *ChatResolver) CreateTemporaryChat(ctx context.Context) (*model.Chat, error) {
	chat := model.Chat{}
	chat.CreatedAt = time.Now()
	chat.UpdatedAt = time.Now()

	err := rsv.ChatRepo.BaseRepo.Insert(chat)
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (rs *ChatResolver) ConnectChatToUsers(ctx context.Context, chatID uint, userID uint, targetID string) (*model.Chat, error) {
	chat, err := rs.ChatRepo.FindById(ctx, chatID)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (rsv *ChatResolver) AddChatMessage(ctx context.Context, chatID string, authorID string, body string) (*model.Chat, error) {
	panic("")
}
