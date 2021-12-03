package resolvers

import (
	"context"
	"proximity/graph/entity"
	"proximity/graph/model"
	"proximity/graph/repositories"
	"proximity/graph/utils"

	"github.com/google/uuid"
)

type MessageResolver struct {
	MessageRepo *repositories.MessageRepo
}

func (rsv *MessageResolver) AddChatMessage(ctx context.Context, chatID string, authorID string, body string) (*entity.Message, error) {
	message := entity.Message{}
	message.ChatId, _ = uuid.Parse(chatID)
	message.UserID, _ = uuid.Parse(authorID)
	message.Body = &body

	return rsv.MessageRepo.Insert(ctx, &message)
}

func (rsv *MessageResolver) MessageSeen(ctx context.Context, messageID string) (*model.Message, error) {
	messageId, err := uuid.Parse(messageID)
	if err != nil {
		return nil, err
	}
	rs, err := rsv.MessageRepo.MessageSeen(ctx, messageId)
	if err != nil {
		return nil, err
	}

	return utils.ToMessageModel(rs)
}
