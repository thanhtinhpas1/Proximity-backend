package graph

import (
	"proximity/graph/model"
	"proximity/graph/resolvers"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserResolver         *resolvers.UserResolver
	BlockListResolver    *resolvers.BlockListResolver
	ChatResolver         *resolvers.ChatResolver
	CommentResolver      *resolvers.CommentResolver
	MessageResolver      *resolvers.MessageResolver
	PostResolver         *resolvers.PostResolver
	NotificationResolver *resolvers.NotificationResolver

	// TODO: using other solution for messages, posts, users
	MessageChannels map[string]chan *model.Chat
	PostChannels    map[string]chan *model.Post
	UserChannels    map[string]chan *model.User
}
