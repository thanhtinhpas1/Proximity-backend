package resolvers

import (
	"context"
	"proximity/graph/entity"
	"proximity/graph/model"
	"proximity/graph/repositories"
	"proximity/graph/utils"

	"github.com/google/uuid"
)

type PostResolver struct {
	PostRepo    *repositories.PostRepo
	LikeRepo    *repositories.LikeRepo
	CommentRepo *repositories.CommentRepo
}

func (rsv *PostResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	parsedId, _ := uuid.Parse(id)
	post, err := rsv.PostRepo.FindById(ctx, parsedId)
	if err != nil {
		return nil, err
	}
	return utils.ToPostModel(post)
}

func (rsv *PostResolver) Posts(ctx context.Context, userID string, offset, limit int) ([]*model.Post, error) {
	parsedUserId, _ := uuid.Parse(userID)
	posts, err := rsv.PostRepo.FindUserFeeds(ctx, parsedUserId, offset, limit)
	if err != nil {
		return nil, err
	}
	return utils.ToPostModels(posts)
}

func (rsv *PostResolver) UserFeed(ctx context.Context, userID string, offset, limit int) ([]*model.Post, error) {
	parsedUserId, _ := uuid.Parse(userID)
	posts, err := rsv.PostRepo.FindUserFeeds(ctx, parsedUserId, offset, limit)
	if err != nil {
		return nil, err
	}
	return utils.ToPostModels(posts)
}

func (rsv *PostResolver) CreatePost(ctx context.Context, userID string, uri string, caption *string) (*model.Post, error) {
	parsedUserId, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	post := entity.Post{UserID: parsedUserId, URI: uri, Caption: caption}
	rs, err := rsv.PostRepo.Insert(ctx, &post)
	if err != nil {
		return nil, err
	}

	return utils.ToPostModel(rs)
}

func (rsv *PostResolver) LikeInteraction(ctx context.Context, userID string, postID string, action model.LikeAction) (*model.Post, error) {
	userId, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	postId, err := uuid.Parse(postID)
	if err != nil {
		return nil, err
	}
	_, err = rsv.LikeRepo.LikeInteraction(ctx, userId, postId, action)
	if err != nil {
		return nil, err
	}
	post, err := rsv.PostRepo.FindById(ctx, postId)
	if err != nil {
		return nil, err
	}

	return utils.ToPostModel(post)
}

func (rsv *PostResolver) ReportPost(ctx context.Context, postID string) (*model.Post, error) {
	postId, err := uuid.Parse(postID)
	if err != nil {
		return nil, err
	}

	rs, err := rsv.PostRepo.ReportPost(ctx, postId)
	if err != nil {
		return nil, err
	}
	return utils.ToPostModel(rs)
}

func (rsv *PostResolver) EditPost(ctx context.Context, postID string, caption string) (*model.Post, error) {
	postId, err := uuid.Parse(postID)
	if err != nil {
		return nil, err
	}

	rs, err := rsv.PostRepo.EditPost(ctx, postId, caption)
	if err != nil {
		return nil, err
	}
	return utils.ToPostModel(rs)
}

func (rsv *PostResolver) DeletePost(ctx context.Context, postID string) (*model.Post, error) {
	postId, err := uuid.Parse(postID)
	if err != nil {
		return nil, err
	}

	rs, err := rsv.PostRepo.DeletePost(ctx, postId)
	if err != nil {
		return nil, err
	}
	return utils.ToPostModel(rs)
}

func (rsv *PostResolver) DeleteComment(ctx context.Context, postID string, commentID string) (*model.Post, error) {
	postId, err := uuid.Parse(postID)
	if err != nil {
		return nil, err
	}

	commentId, err := uuid.Parse(commentID)
	if err != nil {
		return nil, err
	}
	_, err = rsv.CommentRepo.DeleteComment(ctx, postId, commentId)
	if err != nil {
		return nil, err
	}

	post, err := rsv.PostRepo.FindById(ctx, postId)
	if err != nil {
		return nil, err
	}

	return utils.ToPostModel(post)
}
