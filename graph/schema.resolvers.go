package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"proximity/graph/constants"
	"proximity/graph/generated"
	"proximity/graph/model"
	"proximity/graph/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, token string, avatar *string, name string, email string) (*model.User, error) {
	return r.UserResolver.CreateUser(ctx, token, avatar, name, email)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID string, avatar string, name string, handle string, about string) (*model.User, error) {
	return r.UserResolver.UpdateUser(ctx, userID, avatar, name, handle, about)
}

func (r *mutationResolver) CreateTemporaryChat(ctx context.Context, userID string) (*model.Chat, error) {
	return r.ChatResolver.CreateTemporaryChat(ctx, userID)
}

func (r *mutationResolver) ConnectChatToUsers(ctx context.Context, chatID string, userID string, targetID string) (*model.Chat, error) {
	return r.ChatResolver.ConnectChatToUsers(ctx, chatID, userID, targetID)
}

func (r *mutationResolver) AddChatMessage(ctx context.Context, chatID string, authorID string, body string) (*model.Chat, error) {
	msg, err := r.MessageResolver.AddChatMessage(ctx, chatID, authorID, body)
	if err != nil {
		return nil, err
	}
	chat, err := r.ChatResolver.ChatRepo.FindById(ctx, msg.ChatId)
	if err != nil {
		return nil, err
	}
	chatModel, err := utils.ToChatModel(chat)
	if err != nil {
		return nil, err
	}
	chann, ok := r.MessageChannels[chatID]
	if ok {
		chann <- chatModel
	}
	return chatModel, nil
}

func (r *mutationResolver) UpdateFollowing(ctx context.Context, userID string, targetID string, action model.UpdateFollowingAction) (bool, error) {
	return r.UserResolver.UpdateFollowing(ctx, userID, targetID, action)
}

func (r *mutationResolver) CreatePost(ctx context.Context, userID string, uri string, caption *string) (*model.Post, error) {
	return r.PostResolver.CreatePost(ctx, userID, uri, caption)
}

func (r *mutationResolver) UpdateFcmToken(ctx context.Context, userID string, fcmToken string) (*model.User, error) {
	return r.UserResolver.UpdateFcmToken(ctx, userID, fcmToken)
}

func (r *mutationResolver) AddComment(ctx context.Context, userID string, postID string, body string) (*model.Post, error) {
	post, err := r.CommentResolver.AddComment(ctx, userID, postID, body)
	if err != nil {
		return nil, err
	}

	postChan, ok := r.PostChannels[postID]
	if ok {
		postChan <- post
	}

	return post, nil
}

func (r *mutationResolver) LikeInteraction(ctx context.Context, userID string, postID string, action model.LikeAction) (*model.Post, error) {
	post, err := r.PostResolver.LikeInteraction(ctx, userID, postID, action)
	if err != nil {
		return nil, err
	}

	postChan, ok := r.PostChannels[postID]
	if ok {
		postChan <- post
	}

	return post, nil
}

func (r *mutationResolver) MessageSeen(ctx context.Context, messageID string) (*model.Message, error) {
	return r.MessageResolver.MessageSeen(ctx, messageID)
}

func (r *mutationResolver) UpdateLastSeen(ctx context.Context, userID string) (*model.User, error) {
	return r.UserResolver.UpdateLastSeen(ctx, userID)
}

func (r *mutationResolver) DeleteChat(ctx context.Context, chatID string) (*model.Chat, error) {
	return r.ChatResolver.DeleteChat(ctx, chatID)
}

func (r *mutationResolver) ReportPost(ctx context.Context, postID string) (*model.Post, error) {
	return r.PostResolver.ReportPost(ctx, postID)
}

func (r *mutationResolver) EditPost(ctx context.Context, postID string, caption string) (*model.Post, error) {
	return r.PostResolver.EditPost(ctx, postID, caption)
}

func (r *mutationResolver) DeletePost(ctx context.Context, postID string) (*model.Post, error) {
	return r.PostResolver.DeletePost(ctx, postID)
}

func (r *mutationResolver) DeleteComment(ctx context.Context, postID string, commentID string) (*model.Post, error) {
	return r.PostResolver.DeleteComment(ctx, postID, commentID)
}

func (r *mutationResolver) DeleteNotification(ctx context.Context, notificationID string) (*model.Notification, error) {
	return r.NotificationResolver.DeleteNotification(ctx, notificationID)
}

func (r *mutationResolver) BlockUser(ctx context.Context, from string, to string) (bool, error) {
	return r.BlockListResolver.BlockUser(ctx, from, to)
}

func (r *mutationResolver) UnblockUser(ctx context.Context, from string, to string) (bool, error) {
	return r.BlockListResolver.UnblockUser(ctx, from, to)
}

func (r *queryResolver) SignIn(ctx context.Context, token string) (*model.User, error) {
	return r.UserResolver.SignIn(ctx, token)
}

func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	return r.UserResolver.User(ctx, userID)
}

func (r *queryResolver) UserConnections(ctx context.Context, userID string, typeArg model.ConnectionsType) ([]*model.User, error) {
	if typeArg == model.ConnectionsTypeFollowing {
		users, err := r.UserResolver.FindAllFollowings(ctx, userID)
		if err != nil {
			return nil, err
		}
		return utils.ToUsersModel(users)
	}

	users, err := r.UserResolver.FindAllFollowers(ctx, userID)
	if err != nil {
		return nil, err
	}
	return utils.ToUsersModel(users)
}

func (r *queryResolver) UserExists(ctx context.Context, token *string) (bool, error) {
	user, _ := r.UserResolver.FindUserByToken(ctx, *token)
	if user == nil {
		return false, nil
	}
	return true, nil
}

func (r *queryResolver) Notifications(ctx context.Context, userID string) ([]*model.Notification, error) {
	return r.NotificationResolver.FindAllByUserId(ctx, userID)
}

func (r *queryResolver) Chat(ctx context.Context, chatID string) (*model.Chat, error) {
	return r.ChatResolver.Chat(ctx, chatID)
}

func (r *queryResolver) Chats(ctx context.Context, userID string) ([]*model.Chat, error) {
	return r.ChatResolver.FindAllByUserId(ctx, userID)
}

func (r *queryResolver) DoesFollow(ctx context.Context, userID string, targetID string) (bool, error) {
	return r.UserResolver.DoesFollow(ctx, userID, targetID)
}

func (r *queryResolver) ChatExists(ctx context.Context, userID string, targetID string) (*model.Chat, error) {
	return r.ChatResolver.ChatExists(ctx, userID, targetID)
}

func (r *queryResolver) SearchUsers(ctx context.Context, userID string, name string) ([]*model.User, error) {
	return r.UserResolver.SearchUsers(ctx, userID, name)
}

func (r *queryResolver) IsHandleAvailable(ctx context.Context, userID string, handle string) (bool, error) {
	return r.UserResolver.IsHandleAvailable(ctx, userID, handle)
}

func (r *queryResolver) Post(ctx context.Context, postID string) (*model.Post, error) {
	return r.PostResolver.Post(ctx, postID)
}

func (r *queryResolver) Posts(ctx context.Context, userID string, last int) ([]*model.Post, error) {
	return r.PostResolver.Posts(ctx, userID, constants.DEFAULT_OFFSET, constants.DEFAULT_LIMIT)
}

func (r *queryResolver) UserFeed(ctx context.Context, userID string) ([]*model.Post, error) {
	return r.PostResolver.Posts(ctx, userID, constants.DEFAULT_OFFSET, constants.DEFAULT_LIMIT)
}

func (r *queryResolver) LikeUsers(ctx context.Context, likes []string) ([]*model.User, error) {
	return r.UserResolver.LikeUsers(ctx, likes)
}

func (r *queryResolver) BlockedUsers(ctx context.Context, userID string) ([]*model.User, error) {
	return r.UserResolver.BlockedUsers(ctx, userID)
}

func (r *subscriptionResolver) Chat(ctx context.Context, chatID string) (<-chan *model.Chat, error) {
	chatChan := make(chan *model.Chat, 1)
	r.MessageChannels[chatID] = chatChan
	return chatChan, nil
}

func (r *subscriptionResolver) Post(ctx context.Context, postID string) (<-chan *model.Post, error) {
	postChan := make(chan *model.Post)
	r.PostChannels[postID] = postChan
	return postChan, nil
}

func (r *subscriptionResolver) User(ctx context.Context, userID string) (<-chan *model.User, error) {
	userChan := make(chan *model.User)
	r.UserChannels[userID] = userChan
	return userChan, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
