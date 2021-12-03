package repositories

import (
	"context"
	"proximity/graph/entity"
	"proximity/graph/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LikeRepo struct {
	BaseRepo
}

func (repo *LikeRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "likes")
}

func (repo *LikeRepo) LikeInteraction(ctx context.Context, userID, postID uuid.UUID, action model.LikeAction) (*entity.Like, error) {
	like := &entity.Like{UserID: userID, PostID: postID, Action: action.String()}
	repo.DB = repo.DB.WithContext(ctx)
	rs := repo.BaseRepo.Insert(like)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return like, nil
}
