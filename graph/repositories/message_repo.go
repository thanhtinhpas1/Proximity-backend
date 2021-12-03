package repositories

import (
	"context"
	"proximity/graph/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageRepo struct {
	BaseRepo
}

func (repo *MessageRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "messages")
}

func (repo *MessageRepo) Insert(ctx context.Context, message *entity.Message) (*entity.Message, error) {
	repo.DB = repo.DB.WithContext(ctx)
	rs := repo.BaseRepo.Insert(message)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return message, nil
}

func (repo *MessageRepo) MessageSeen(ctx context.Context, messageID uuid.UUID) (*entity.Message, error) {
	repo.DB = repo.DB.WithContext(ctx)
	message := &entity.Message{}
	rs := repo.BaseRepo.DB.Model(message).Where("id = ?", messageID).Update("seen", true)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return message, nil
}
