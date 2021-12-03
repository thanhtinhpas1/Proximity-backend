package main

import (
	"log"
	"net/http"
	"os"
	"proximity/graph"
	"proximity/graph/generated"
	"proximity/graph/model"
	"proximity/graph/repositories"
	"proximity/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func connectDatabase() *gorm.DB {
	dsn := "user=postgres password=admin dbname=postgres host=localhost port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := connectDatabase()

	err := db.Statement.AutoMigrate(&model.BlockList{}, &model.User{}, &model.Chat{}, &model.Comment{}, &model.Message{}, &model.Notification{}, &model.Post{}, &model.Story{})
	if err != nil {
		log.Fatal(err)
	}

	userRepo := &repositories.UserRepo{}
	userRepo.Init(db)
	userResolver := &resolvers.UserResolver{UserRepo: userRepo}

	chatRepo := &repositories.ChatRepo{}
	chatRepo.Init(db)
	chatResolver := &resolvers.ChatResolver{ChatRepo: chatRepo}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{UserResolver: userResolver, ChatResolver: chatResolver}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
