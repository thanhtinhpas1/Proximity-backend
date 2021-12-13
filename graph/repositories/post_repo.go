package repositories

import (
	"context"
	"proximity/graph/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostRepo struct {
	BaseRepo
}

func (repo *PostRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "posts")
}

func (repo *PostRepo) Insert(ctx context.Context, entity *entity.Post) (*entity.Post, error) {
	repo.DB = repo.DB.WithContext(ctx)
	rs := repo.BaseRepo.Insert(entity)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return entity, nil
}

func (repo *PostRepo) FindById(ctx context.Context, id uuid.UUID) (*entity.Post, error) {
	post := &entity.Post{}
	repo.DB = repo.DB.WithContext(ctx)
	err := repo.DB.Preload("Comments.Author").Preload(clause.Associations).Find(post, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *PostRepo) FindAllByUserId(ctx context.Context, userId uuid.UUID, offset, limit int) ([]*entity.Post, error) {
	posts := []*entity.Post{}
	repo.DB = repo.DB.WithContext(ctx)
	rs := repo.DB.Offset(offset).Limit(limit).Preload(clause.Associations).Find(&posts, "user_id = ?", userId)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return posts, nil
}

func (repo *PostRepo) FindUserFeeds(ctx context.Context, userId uuid.UUID, offset, limit int) ([]*entity.Post, error) {
	posts := []*entity.Post{}
	repo.DB = repo.DB.WithContext(ctx)
	rs := repo.DB.Offset(offset).Limit(limit).Preload(clause.Associations).Joins("LEFT JOIN user_followings ON posts.user_id = user_followings.following_id AND (user_followings.following_id = ? OR posts.user_id = ?)", userId, userId).Find(&posts)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return posts, nil
}

func (repo *PostRepo) ReportPost(ctx context.Context, postID uuid.UUID) (*entity.Post, error) {
	repo.DB = repo.DB.WithContext(ctx)
	post, err := repo.FindById(ctx, postID)
	if err != nil {
		return nil, err
	}

	post.Reports++
	err = repo.Update(ctx, postID, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *PostRepo) EditPost(ctx context.Context, postID uuid.UUID, caption string) (*entity.Post, error) {
	repo.DB = repo.DB.WithContext(ctx)
	post := &entity.Post{}
	err := repo.DB.Model(post).Where("id = ?", postID).Update("caption", caption).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (repo *PostRepo) DeletePost(ctx context.Context, postID uuid.UUID) (*entity.Post, error) {
	repo.DB = repo.DB.WithContext(ctx)
	post, err := repo.FindById(ctx, postID)
	if err != nil {
		return nil, err
	}

	err = repo.DB.Delete(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}
