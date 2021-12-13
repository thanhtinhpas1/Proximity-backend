package utils

import (
	"encoding/json"
	"proximity/graph/entity"
	"proximity/graph/model"
)

func ToUserModel(user *entity.User) (*model.User, error) {
	bytes, err := json.Marshal(&user)
	if err != nil {
		return nil, err
	}

	userModel := model.User{}
	err = json.Unmarshal(bytes, &userModel)
	if err != nil {
		return nil, err
	}
	return &userModel, nil
}

func ToUsersModel(users []*entity.User) ([]*model.User, error) {
	userModels := make([]*model.User, len(users))

	for i, v := range users {
		userModels[i], _ = ToUserModel(v)
	}
	return userModels, nil
}

func ToChatModel(chat *entity.Chat) (*model.Chat, error) {
	bytes, err := json.Marshal(&chat)
	if err != nil {
		return nil, err
	}

	chatModel := model.Chat{}
	err = json.Unmarshal(bytes, &chatModel)
	if err != nil {
		return nil, err
	}
	return &chatModel, nil
}

func ToChatModels(chats []*entity.Chat) ([]*model.Chat, error) {
	chatModels := make([]*model.Chat, len(chats))
	for i, v := range chats {
		chatModels[i], _ = ToChatModel(v)
	}

	return chatModels, nil
}

func ToMessageModel(message *entity.Message) (*model.Message, error) {
	bytes, err := json.Marshal(&message)
	if err != nil {
		return nil, err
	}

	model := model.Message{}
	err = json.Unmarshal(bytes, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func ToPostModel(post *entity.Post) (*model.Post, error) {
	bytes, err := json.Marshal(&post)
	if err != nil {
		return nil, err
	}

	model := model.Post{}
	err = json.Unmarshal(bytes, &model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func ToPostModels(posts []*entity.Post) ([]*model.Post, error) {
	postModels := make([]*model.Post, len(posts))
	for i, v := range posts {
		model, err := ToPostModel(v)
		if err != nil {
			return nil, err
		}
		postModels[i] = model
	}

	return postModels, nil
}

func ToNotificationModel(notification *entity.Notification) (*model.Notification, error) {
	bytes, err := json.Marshal(&notification)
	if err != nil {
		return nil, err
	}

	model := model.Notification{}
	err = json.Unmarshal(bytes, &model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func ToNotificationModels(notifications []*entity.Notification) ([]*model.Notification, error) {
	notificationModels := make([]*model.Notification, len(notifications))
	for i, v := range notifications {
		notificationModels[i], _ = ToNotificationModel(v)
	}

	return notificationModels, nil
}
