package resolvers

import (
	"context"
	"proximity/graph/entity"
	"proximity/graph/model"
	"proximity/graph/repositories"
	"proximity/graph/utils"

	"github.com/google/uuid"
)

type ChatResolver struct {
	ChatRepo        *repositories.ChatRepo
	ParticipantRepo *repositories.ParticipantRepo
}

func (rsv *ChatResolver) CreateTemporaryChat(ctx context.Context, userID string) (*model.Chat, error) {
	userId, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	chat := entity.Chat{UserID: userId}
	rs, err := rsv.ChatRepo.Insert(ctx, &chat)
	if err != nil {
		return nil, err
	}
	chatRs, err := rsv.ChatRepo.FindById(ctx, rs.ID)
	if err != nil {
		return nil, err
	}
	return utils.ToChatModel(chatRs)
}

func (rsv *ChatResolver) ConnectChatToUsers(ctx context.Context, chatID string, userID string, targetID string) (*model.Chat, error) {
	chatId, _ := uuid.Parse(chatID)
	userId, _ := uuid.Parse(userID)
	targetId, _ := uuid.Parse(targetID)

	err := rsv.ParticipantRepo.ConnectChatToUsers(ctx, chatId, userId, targetId)
	if err != nil {
		return nil, err
	}

	return rsv.Chat(ctx, chatID)
}

func (rsv *ChatResolver) FindAllByUserId(ctx context.Context, userId string) ([]*model.Chat, error) {
	parsedUserId, _ := uuid.Parse(userId)
	rs, err := rsv.ChatRepo.FindAllByUserId(ctx, parsedUserId)
	if err != nil {
		return nil, err
	}
	return utils.ToChatModels(rs)
}

func (rsv *ChatResolver) DeleteChat(ctx context.Context, chatID string) (*model.Chat, error) {
	chatId, err := uuid.Parse(chatID)
	if err != nil {
		return nil, err
	}

	rs, err := rsv.ChatRepo.DeleteChat(ctx, chatId)
	if err != nil {
		return nil, err
	}
	return utils.ToChatModel(rs)
}

func (rsv *ChatResolver) Chat(ctx context.Context, chatID string) (*model.Chat, error) {
	chatId, err := uuid.Parse(chatID)
	if err != nil {
		return nil, err
	}

	rs, err := rsv.ChatRepo.FindById(ctx, chatId)
	if err != nil {
		return nil, err
	}

	return utils.ToChatModel(rs)
}

func (rsv *ChatResolver) ChatExists(ctx context.Context, userID string, targetID string) (*model.Chat, error) {
	userId, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	targetId, err := uuid.Parse(targetID)
	if err != nil {
		return nil, err
	}

	rs, err := rsv.ChatRepo.ChatExists(ctx, userId, targetId)
	if err != nil {
		return nil, nil
	}
	return utils.ToChatModel(rs)
}
