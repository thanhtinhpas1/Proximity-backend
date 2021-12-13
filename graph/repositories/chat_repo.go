package repositories

import (
	"context"
	"proximity/graph/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ChatRepo struct {
	BaseRepo
}

func (repo *ChatRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "chats")
}

func (repo *ChatRepo) Insert(ctx context.Context, chat *entity.Chat) (*entity.Chat, error) {
	repo.DB = repo.DB.WithContext(ctx)
	rs := repo.BaseRepo.Insert(chat)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return chat, nil
}

func (repo *ChatRepo) FindById(ctx context.Context, id uuid.UUID) (*entity.Chat, error) {
	chat := &entity.Chat{}
	repo.DB = repo.DB.WithContext(ctx)
	err := repo.DB.Preload("Messages.Author").Preload(clause.Associations).Find(chat, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (repo *ChatRepo) FindAllByUserId(ctx context.Context, userId uuid.UUID) ([]*entity.Chat, error) {
	chats := []*entity.Chat{}
	repo.DB = repo.DB.WithContext(ctx)
	err := repo.BaseRepo.DB.Preload("Messages.Author").Preload(clause.Associations).Joins("JOIN chat_participants ON chats.id = chat_participants.chat_id AND chat_participants.user_id = ?", userId).Find(&chats)
	if err.Error != nil {
		return nil, err.Error
	}
	return chats, nil
}

func (repo *ChatRepo) DeleteChat(ctx context.Context, chatID uuid.UUID) (*entity.Chat, error) {
	repo.DB = repo.DB.WithContext(ctx)
	chat, err := repo.FindById(ctx, chatID)
	if err != nil {
		return nil, err
	}

	err = repo.DB.Delete(chat).Error
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (repo *ChatRepo) ChatExists(ctx context.Context, userID, targetID uuid.UUID) (*entity.Chat, error) {
	repo.DB = repo.DB.WithContext(ctx)
	chat := &entity.Chat{}
	err := repo.DB.Joins("JOIN chat_participants ON chats.id = chat_participants.chat_id AND (chats.user_id = ? AND chat_participants.user_id = ? OR chats.user_id = ? AND chat_participants.user_id = ?)", userID, targetID, targetID, userID).First(chat).Error
	if err != nil {
		return nil, err
	}

	return chat, nil
}
