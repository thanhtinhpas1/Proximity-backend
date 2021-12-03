package graph

import (
	"proximity/graph/resolvers"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserResolver *resolvers.UserResolver
	ChatResolver *resolvers.ChatResolver
}
