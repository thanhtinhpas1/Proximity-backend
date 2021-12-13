package resolvers

import (
	"context"
	"errors"
	"proximity/graph/entity"
	"proximity/graph/model"
	"proximity/graph/repositories"
	"proximity/graph/utils"

	"github.com/google/uuid"
)

type UserResolver struct {
	UserRepo      *repositories.UserRepo
	FollowingRepo *repositories.FollowingRepo
	FollowerRepo  *repositories.FollowerRepo
}

func (rsv *UserResolver) CreateUser(ctx context.Context, token string, avatar *string, name string, email string) (*model.User, error) {
	user := entity.User{Token: token, Avatar: avatar, Name: name, Email: email}
	rs, err := rsv.UserRepo.Insert(ctx, &user)
	if err != nil {
		return nil, err
	}
	return utils.ToUserModel(rs)
}

func (rsv *UserResolver) UpdateUser(ctx context.Context, userID string, avatar string, name string, handle string, about string) (*model.User, error) {
	userId, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid UserID")
	}
	user := entity.User{ID: userId, Avatar: &avatar, Name: name, Handle: handle, About: &about}
	rs, err := rsv.UserRepo.Update(ctx, userId, &user)
	if err != nil {
		return nil, err
	}

	return utils.ToUserModel(rs)
}

func (rsv *UserResolver) UpdateFcmToken(ctx context.Context, userId string, fcmToken string) (*model.User, error) {
	user := &entity.User{FcmToken: &fcmToken}
	parsedUserId, _ := uuid.Parse(userId)
	rs, err := rsv.UserRepo.Update(ctx, parsedUserId, user)
	if err != nil {
		return nil, err
	}
	return utils.ToUserModel(rs)
}

func (rsv *UserResolver) FindUserByToken(ctx context.Context, token string) (*entity.User, error) {
	user, err := rsv.UserRepo.FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (rsv *UserResolver) FindAllFollowings(ctx context.Context, userId string) ([]*entity.User, error) {
	parsedUserId, _ := uuid.Parse(userId)
	users, err := rsv.UserRepo.FindAllFollowings(ctx, parsedUserId)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (rsv *UserResolver) FindAllFollowers(ctx context.Context, userId string) ([]*entity.User, error) {
	parsedUserId, _ := uuid.Parse(userId)
	users, err := rsv.UserRepo.FindAllFollowers(ctx, parsedUserId)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (rsv *UserResolver) SearchUsers(ctx context.Context, userID string, name string) ([]*model.User, error) {
	parsedUserId, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	users, err := rsv.UserRepo.SearchUsers(ctx, parsedUserId, name)
	if err != nil {
		return nil, err
	}
	return utils.ToUsersModel(users)
}

func (rsv *UserResolver) UpdateFollowing(ctx context.Context, userID string, targetID string, action model.UpdateFollowingAction) (bool, error) {
	parseUserId, err := uuid.Parse(userID)
	if err != nil {
		return false, err
	}
	parsedTargetId, err := uuid.Parse(targetID)
	if err != nil {
		return false, err
	}

	if action == model.UpdateFollowingActionFollow {
		following := &entity.UserFollowing{UserID: parseUserId, FollowingID: parsedTargetId}
		_, err = rsv.FollowingRepo.Insert(ctx, following)
		if err != nil {
			return false, err
		}

		follower := &entity.UserFollower{UserID: parsedTargetId, FollowerID: parseUserId}
		_, err = rsv.FollowerRepo.Insert(ctx, follower)
		if err != nil {
			return false, err
		}

		return true, nil
	} else {
		rsv.FollowingRepo.Remove(ctx, parseUserId, parsedTargetId)
		rsv.FollowerRepo.Remove(ctx, parsedTargetId, parseUserId)

		return true, nil
	}
}

func (rsv *UserResolver) SignIn(ctx context.Context, token string) (*model.User, error) {
	user, err := rsv.UserRepo.FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return utils.ToUserModel(user)
}

func (rsv *UserResolver) UpdateLastSeen(ctx context.Context, userId string) (*model.User, error) {
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}
	rs, err := rsv.UserRepo.UpdateLastSeen(ctx, parsedUserId)
	if err != nil {
		return nil, err
	}
	return utils.ToUserModel(rs)
}

func (rsv *UserResolver) User(ctx context.Context, userID string) (*model.User, error) {
	userId, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	rs, err := rsv.UserRepo.FindById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return utils.ToUserModel(rs)
}

func (rsv *UserResolver) LikeUsers(ctx context.Context, likes []string) ([]*model.User, error) {
	// likeUuids, err := utils.ToUuids(likes)
	// if err != nil {
	// 	return nil, err
	// }

	users, err := rsv.UserRepo.LikeUsers(ctx, likes)
	if err != nil {
		return nil, err
	}

	return utils.ToUsersModel(users)
}

func (rsv *UserResolver) BlockedUsers(ctx context.Context, userID string) ([]*model.User, error) {
	userId, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	users, err := rsv.UserRepo.FindAllBlockedUsers(ctx, userId)
	if err != nil {
		return nil, err
	}

	return utils.ToUsersModel(users)
}

func (rsv *UserResolver) IsHandleAvailable(ctx context.Context, userID string, handle string) (bool, error) {
	userId, err := uuid.Parse(userID)
	if err != nil {
		return false, err
	}

	rs, err := rsv.UserRepo.IsHandleAvailable(ctx, userId, handle)
	if err != nil {
		return false, err
	}

	return rs, nil
}

func (rsv *UserResolver) DoesFollow(ctx context.Context, userID string, targetID string) (bool, error) {
	userId, err := uuid.Parse(userID)
	if err != nil {
		return false, err
	}

	targetId, err := uuid.Parse(targetID)
	if err != nil {
		return false, err
	}

	return rsv.FollowingRepo.DoesFollow(ctx, userId, targetId)
}
