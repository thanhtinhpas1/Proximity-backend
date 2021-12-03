package repositories

import (
	"context"
	"proximity/graph/model"

	"gorm.io/gorm"
)

type ChatRepo struct {
	BaseRepo
}

func (repo *ChatRepo) Init(db *gorm.DB) {
	repo = &ChatRepo{}
	repo.BaseRepo.Init(db, "chats")
}

func (repo *ChatRepo) FindById(ctx context.Context, id uint) (*model.Chat, error) {
	chat := &model.Chat{ID: id}
	err := repo.BaseRepo.FindById(chat)
	if err != nil {
		return nil, err
	}
	return chat, nil
}
