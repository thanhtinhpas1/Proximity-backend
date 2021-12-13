package repositories

import (
	"context"
	"proximity/graph/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ParticipantRepo struct {
	BaseRepo
}

func (repo *ParticipantRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "chat_participants")
}

func (repo *ParticipantRepo) JoinChat(ctx context.Context, chatId, userId uuid.UUID) error {
	repo.DB = repo.DB.WithContext(ctx)
	participant := entity.Participant{ChatID: chatId, UserID: userId}
	rs := repo.DB.Create(participant)
	if rs.Error != nil {
		return rs.Error
	}
	return nil
}

func (repo *ParticipantRepo) ConnectChatToUsers(ctx context.Context, chatID, userID, targetID uuid.UUID) error {
	repo.DB = repo.DB.WithContext(ctx)
	participant1 := entity.Participant{ChatID: chatID, UserID: userID}
	participant2 := entity.Participant{ChatID: chatID, UserID: targetID}

	rs := repo.DB.CreateInBatches([]entity.Participant{participant1, participant2}, 2)
	if rs.Error != nil {
		return rs.Error
	}

	return nil
}
