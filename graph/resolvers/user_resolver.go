package resolvers

import (
	"context"
	"proximity/graph/model"
	"proximity/graph/repositories"
)

type UserResolver struct {
	UserRepo *repositories.UserRepo
}

func (rsv *UserResolver) CreateUser(ctx context.Context, token string, avatar *string, name string, email string) (*model.User, error) {
	user := model.User{Token: token, Avatar: avatar, Name: name, Email: email}
	err := rsv.UserRepo.BaseRepo.Insert(user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (rsv *UserResolver) UpdateUser(ctx context.Context, userID uint, avatar string, name string, handle string, about string) (*model.User, error) {
	user := model.User{ID: userID, Avatar: &avatar, Name: name, Handle: handle, About: &about}
	err := rsv.UserRepo.BaseRepo.Update(userID, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
