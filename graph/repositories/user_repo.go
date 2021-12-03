package repositories

import (
	"context"
	"proximity/graph/constants"
	"proximity/graph/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	BaseRepo
}

func (repo *UserRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "users")
}

func (repo *UserRepo) Insert(ctx context.Context, data *entity.User) (*entity.User, error) {
	repo.DB = repo.DB.WithContext(ctx)

	rs := repo.BaseRepo.Insert(data)
	if rs.Error != nil {
		return nil, rs.Error
	}
	return data, nil
}

func (repo *UserRepo) FindByToken(ctx context.Context, token string) (*entity.User, error) {
	repo.DB = repo.DB.WithContext(ctx)

	user := &entity.User{}
	rs := repo.DB.First(user, "token = ?", token)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return user, nil
}

func (repo *UserRepo) FindById(ctx context.Context, userId uuid.UUID) (*entity.User, error) {
	repo.DB = repo.DB.WithContext(ctx)
	user := &entity.User{}

	err := repo.DB.Preload(clause.Associations).First(user, "id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) FindAllFollowings(ctx context.Context, userId uuid.UUID) ([]*entity.User, error) {
	repo.DB = repo.DB.WithContext(ctx)

	users := []*entity.User{}
	rs := repo.DB.Joins("JOIN user_followings ON user_followings.user_id = users.id AND users.id = ?", userId).Model(&users)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return users, nil
}

func (repo *UserRepo) FindAllFollowers(ctx context.Context, userId uuid.UUID) ([]*entity.User, error) {
	repo.DB = repo.DB.WithContext(ctx)

	users := []*entity.User{}
	rs := repo.DB.Joins("JOIN user_followers ON user_followers.user_id = users.id AND users.id = ?", userId).Model(&users)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return users, nil
}

func (repo *UserRepo) SearchUsers(ctx context.Context, userId uuid.UUID, name string) ([]*entity.User, error) {
	repo.DB = repo.DB.WithContext(ctx)

	users := make([]*entity.User, constants.DEFAULT_LIMIT)
	// TODO: Optimize query
	rs := repo.DB.Limit(constants.DEFAULT_LIMIT).Find(&users, "name LIKE ?", name+"%")
	if rs.Error != nil {
		return nil, rs.Error
	}
	return users, nil
}

func (repo *UserRepo) Update(ctx context.Context, userId uuid.UUID, user *entity.User) (*entity.User, error) {
	err := repo.BaseRepo.Update(ctx, userId, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) UpdateLastSeen(ctx context.Context, userId uuid.UUID) (*entity.User, error) {
	repo.DB = repo.DB.WithContext(ctx)
	user := &entity.User{}
	err := repo.BaseRepo.DB.Model(user).Where("id = ?", userId).Update("last_seen", time.Now().UTC().UnixMilli()).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) LikeUsers(ctx context.Context, likes []uuid.UUID) ([]*entity.User, error) {
	repo.DB = repo.DB.WithContext(ctx)
	users := []*entity.User{}
	err := repo.DB.Find(&users, likes).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepo) FindAllBlockedUsers(ctx context.Context, userID uuid.UUID) ([]*entity.User, error) {
	repo.DB = repo.DB.WithContext(ctx)
	users := []*entity.User{}
	err := repo.DB.Joins("JOIN block_lists ON users.id = block_lists.to AND block_lists.from = ?", userID).Preload(clause.Associations).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepo) IsHandleAvailable(ctx context.Context, userID uuid.UUID, handle string) (bool, error) {
	repo.DB = repo.DB.WithContext(ctx)
	user := &entity.User{}
	err := repo.DB.Find(&user, "id = ? AND handle = ?", userID, handle).Error
	if err != nil {
		return false, err
	}

	return true, err
}
