package repositories

import (
	"context"
	"proximity/graph/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepo struct {
	BaseRepo
}

func (repo *CommentRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "comments")
}

func (repo *CommentRepo) AddComment(ctx context.Context, userID uuid.UUID, postID uuid.UUID, body string) (*entity.Comment, error) {
	comment := &entity.Comment{UserID: userID, PostID: postID, Body: body}
	repo.DB = repo.DB.WithContext(ctx)
	rs := repo.BaseRepo.Insert(comment)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return comment, nil
}

func (repo *CommentRepo) DeleteComment(ctx context.Context, postID, commentID uuid.UUID) (*entity.Comment, error) {
	repo.DB = repo.DB.WithContext(ctx)
	comment := &entity.Comment{}
	err := repo.FindById(comment, commentID)
	if err != nil {
		return nil, err
	}

	err = repo.DB.Delete(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}
