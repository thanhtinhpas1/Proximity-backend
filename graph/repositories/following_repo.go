package repositories

import (
	"context"
	"proximity/graph/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FollowingRepo struct {
	BaseRepo
}

func (repo *FollowingRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "user_followings")
}

func (repo *FollowingRepo) Insert(ctx context.Context, entity *entity.UserFollowing) (*entity.UserFollowing, error) {
	repo.DB = repo.DB.WithContext(ctx)
	rs := repo.BaseRepo.Insert(entity)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return entity, nil
}

func (repo *FollowingRepo) DoesFollow(ctx context.Context, userID, targetID uuid.UUID) (bool, error) {
	repo.DB = repo.DB.WithContext(ctx)
	err := repo.DB.First(&entity.UserFollowing{}, "user_id = ? AND following_id = ?", userID, targetID).Error
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (repo *FollowingRepo) Remove(ctx context.Context, userId, targetId uuid.UUID) (bool, error) {
	repo.DB = repo.DB.WithContext(ctx)
	following := &entity.UserFollowing{UserID: userId, FollowingID: targetId}
	err := repo.DB.Delete(following, "user_id = ? AND following_id = ?", following.UserID, following.FollowingID).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
