package resolvers

import (
	"context"
	"proximity/graph/model"
	"proximity/graph/repositories"
	"proximity/graph/utils"

	"github.com/google/uuid"
)

type NotificationResolver struct {
	NotificationRepo *repositories.NotificationRepo
}

func (rsv *NotificationResolver) FindAllByUserId(ctx context.Context, userID string) ([]*model.Notification, error) {
	parsedUserId, _ := uuid.Parse(userID)
	notifications, err := rsv.NotificationRepo.FindByUserId(ctx, parsedUserId)
	if err != nil {
		return nil, err
	}

	return utils.ToNotificationModels(notifications)
}

func (rsv *NotificationResolver) DeleteNotification(ctx context.Context, notificationID string) (*model.Notification, error) {
	notificationId, err := uuid.Parse(notificationID)
	if err != nil {
		return nil, err
	}

	rs, err := rsv.NotificationRepo.DeleteNotification(ctx, notificationId)
	if err != nil {
		return nil, err
	}
	return utils.ToNotificationModel(rs)
}
