package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseRepo struct {
	DB *gorm.DB
}

func (repo *BaseRepo) Init(db *gorm.DB, tableName string) {
	repo.DB = db.Table(tableName)
}

func (repo *BaseRepo) Insert(data interface{}) *gorm.DB {
	return repo.DB.Create(data)
}

func (repo *BaseRepo) Update(ctx context.Context, id uuid.UUID, data interface{}) error {
	repo.DB = repo.DB.WithContext(ctx)
	return repo.DB.Where("id = ?", id).Updates(data).Error
}

func (repo *BaseRepo) FindById(dest interface{}, id uuid.UUID) error {
	return repo.DB.First(dest, "id = ?", id).Error
}
