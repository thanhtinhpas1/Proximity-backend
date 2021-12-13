package repositories

import (
	"context"
	"proximity/graph/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlockListRepo struct {
	BaseRepo
}

func (repo *BlockListRepo) Init(db *gorm.DB) {
	repo.BaseRepo.Init(db, "block_lists")
}

func (repo *BlockListRepo) BlockUser(ctx context.Context, from, to uuid.UUID) (bool, error) {
	blockUser := entity.BlockList{From: from, To: to}
	repo.DB = repo.DB.WithContext(ctx)
	err := repo.DB.Create(&blockUser).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *BlockListRepo) UnblockUser(ctx context.Context, from, to uuid.UUID) (bool, error) {
	blockUser := &entity.BlockList{}
	repo.DB = repo.DB.WithContext(ctx)
	err := repo.DB.Find(blockUser, "from = ? AND to = ?", from, to).Error
	if err != nil {
		return false, err
	}

	err = repo.DB.Delete(blockUser).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
