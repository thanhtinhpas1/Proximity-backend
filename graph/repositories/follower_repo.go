package repositories

import (
	"context"
	"proximity/graph/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FollowerRepo struct {
	BaseRepo
}

func (repo *FollowerRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "user_followers")
}

func (repo *FollowerRepo) Insert(ctx context.Context, entity *entity.UserFollower) (*entity.UserFollower, error) {
	repo.DB = repo.DB.WithContext(ctx)
	rs := repo.DB.Create(entity)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return entity, nil
}

func (repo *FollowerRepo) Remove(ctx context.Context, userId, targetId uuid.UUID) (bool, error) {
	repo.DB = repo.DB.WithContext(ctx)
	follower := &entity.UserFollower{UserID: userId, FollowerID: targetId}
	err := repo.DB.Delete(follower, "user_id = ? AND follower_id = ?", follower.UserID, follower.FollowerID).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
