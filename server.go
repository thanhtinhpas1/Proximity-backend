package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"proximity/graph"
	"proximity/graph/entity"
	"proximity/graph/generated"
	"proximity/graph/model"
	"proximity/graph/repositories"
	"proximity/graph/resolvers"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultPort = "8080"

func connectDatabase(ctx context.Context) *gorm.DB {
	dsn := "user=postgres password=admin dbname=postgres host=localhost port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                   newLogger,
		DisableNestedTransaction: true,
		PrepareStmt:              true,
	})
	db = db.Session(&gorm.Session{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	return db
}

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Warn, // Log level
		IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
		Colorful:                  true,        // Disable color
	},
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := connectDatabase(context.Background())

	err := db.Statement.AutoMigrate(&entity.BlockList{}, &entity.User{}, &entity.Chat{}, &entity.Comment{}, &entity.Message{}, &entity.Notification{}, &entity.Post{}, &entity.Story{}, &entity.Like{}, &entity.UserFollowing{}, &entity.UserFollower{})
	if err != nil {
		log.Fatal(err)
	}

	userRepo := &repositories.UserRepo{}
	userRepo.Init(db)
	followingRepo := &repositories.FollowingRepo{}
	followingRepo.Init(db)
	followerRepo := &repositories.FollowerRepo{}
	followerRepo.Init(db)
	userResolver := &resolvers.UserResolver{UserRepo: userRepo, FollowingRepo: followingRepo, FollowerRepo: followerRepo}

	chatRepo := &repositories.ChatRepo{}
	chatRepo.Init(db)
	participantRepo := &repositories.ParticipantRepo{}
	participantRepo.Init(db)
	chatResolver := &resolvers.ChatResolver{ChatRepo: chatRepo, ParticipantRepo: participantRepo}

	messageRepo := &repositories.MessageRepo{}
	messageRepo.Init(db)
	messageResolver := &resolvers.MessageResolver{MessageRepo: messageRepo}

	postRepo := &repositories.PostRepo{}
	postRepo.Init(db)
	commentRepo := &repositories.CommentRepo{}
	commentRepo.Init(db)
	commentResolver := &resolvers.CommentResolver{CommentRepo: commentRepo, PostRepo: postRepo}

	likeRepo := &repositories.LikeRepo{}
	likeRepo.Init(db)
	postResolver := &resolvers.PostResolver{PostRepo: postRepo, LikeRepo: likeRepo, CommentRepo: commentRepo}

	blockListRepo := &repositories.BlockListRepo{}
	blockListRepo.Init(db)
	blockListResolver := &resolvers.BlockListResolver{BlockListRepo: blockListRepo}

	notificationRepo := &repositories.NotificationRepo{}
	notificationRepo.Init(db)
	notificationResolver := &resolvers.NotificationResolver{NotificationRepo: notificationRepo}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserResolver:         userResolver,
		ChatResolver:         chatResolver,
		CommentResolver:      commentResolver,
		MessageResolver:      messageResolver,
		PostResolver:         postResolver,
		NotificationResolver: notificationResolver,
		BlockListResolver:    blockListResolver,

		MessageChannels: map[string]chan *model.Chat{},
		PostChannels:    map[string]chan *model.Post{},
		UserChannels:    map[string]chan *model.User{},
	}}))

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	http.Handle("/query", playground.Handler("GraphQL playground", "/"))
	http.Handle("/", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
