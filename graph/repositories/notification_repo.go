package repositories

import (
	"context"
	"proximity/graph/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepo struct {
	BaseRepo
}

func (repo *NotificationRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "notifications")
}

func (repo *NotificationRepo) FindByUserId(ctx context.Context, userId uuid.UUID) ([]*entity.Notification, error) {
	notifications := []*entity.Notification{}
	repo.DB = repo.DB.WithContext(ctx)
	rs := repo.BaseRepo.DB.Find(&notifications, "user_id = ?", userId)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return notifications, nil
}

func (repo *NotificationRepo) DeleteNotification(ctx context.Context, notificationID uuid.UUID) (*entity.Notification, error) {
	repo.DB = repo.DB.WithContext(ctx)
	notification := &entity.Notification{}
	err := repo.FindById(notification, notificationID)
	if err != nil {
		return nil, err
	}

	err = repo.DB.Delete(notification).Error
	if err != nil {
		return nil, err
	}
	return notification, nil
}
