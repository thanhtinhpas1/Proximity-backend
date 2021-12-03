package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"proximity/graph/generated"
	"proximity/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, token string, avatar *string, name string, email string) (*model.User, error) {
	return r.UserResolver.CreateUser(ctx, token, avatar, name, email)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID uint, avatar string, name string, handle string, about string) (*model.User, error) {
	return r.UserResolver.UpdateUser(ctx, userID, avatar, name, handle, about)
}

func (r *mutationResolver) CreateTemporaryChat(ctx context.Context) (*model.Chat, error) {
	return r.ChatResolver.CreateTemporaryChat(ctx)
}

func (r *mutationResolver) ConnectChatToUsers(ctx context.Context, chatID uint, userID uint, targetID string) (*model.Chat, error) {
	return r.ChatResolver.ConnectChatToUsers(ctx, chatID, userID, targetID)
}

func (r *mutationResolver) AddChatMessage(ctx context.Context, chatID string, authorID string, body string) (*model.Chat, error) {
	return r.ChatResolver.AddChatMessage(ctx, chatID, authorID, body)
}

func (r *mutationResolver) UpdateFollowing(ctx context.Context, userID string, targetID string, action model.UpdateFollowingAction) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, userID string, uri string, caption *string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFcmToken(ctx context.Context, userID string, fcmToken string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddComment(ctx context.Context, userID string, postID string, body string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LikeInteraction(ctx context.Context, userID string, postID string, action model.LikeAction) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MessageSeen(ctx context.Context, messageID string) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateLastSeen(ctx context.Context, userID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteChat(ctx context.Context, chatID string) (*model.Chat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReportPost(ctx context.Context, postID string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditPost(ctx context.Context, postID string, caption string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePost(ctx context.Context, postID string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteComment(ctx context.Context, postID string, commentID string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteNotification(ctx context.Context, notificationID string) (*model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) BlockUser(ctx context.Context, from string, to string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UnblockUser(ctx context.Context, from string, to string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SignIn(ctx context.Context, token string) (*model.User, error) {
	return &model.User{}, nil
}

func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserConnections(ctx context.Context, userID string, typeArg model.ConnectionsType) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserExists(ctx context.Context, token *string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Notifications(ctx context.Context, userID string) ([]*model.Notification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Chat(ctx context.Context, chatID string) (*model.Chat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Chats(ctx context.Context, userID string) ([]*model.Chat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) DoesFollow(ctx context.Context, userID string, targetID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ChatExists(ctx context.Context, userID string, targetID string) (*model.Chat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SearchUsers(ctx context.Context, userID string, name string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IsHandleAvailable(ctx context.Context, userID string, handle string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Post(ctx context.Context, postID string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context, userID string, last int) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserFeed(ctx context.Context, userID string) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) LikeUsers(ctx context.Context, likes []string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BlockedUsers(ctx context.Context, userID string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Chat(ctx context.Context, chatID string) (<-chan *model.Chat, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Post(ctx context.Context, postID string) (<-chan *model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) User(ctx context.Context, userID string) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented"))
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
