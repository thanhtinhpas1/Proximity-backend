package resolvers

import (
	"context"
	"proximity/graph/model"
	"proximity/graph/repositories"
	"proximity/graph/utils"

	"github.com/google/uuid"
)

type CommentResolver struct {
	CommentRepo *repositories.CommentRepo
	PostRepo    *repositories.PostRepo
}

func (rsv *CommentResolver) AddComment(ctx context.Context, userID string, postID string, body string) (*model.Post, error) {
	userId, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	postId, err := uuid.Parse(postID)
	if err != nil {
		return nil, err
	}

	_, err = rsv.CommentRepo.AddComment(ctx, userId, postId, body)
	if err != nil {
		return nil, err
	}

	chat, err := rsv.PostRepo.FindById(ctx, postId)

	if err != nil {
		return nil, err
	}
	return utils.ToPostModel(chat)
}
