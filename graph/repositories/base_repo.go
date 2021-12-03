package repositories

import (
	"gorm.io/gorm"
)

type BaseRepo struct {
	DB *gorm.DB
}

func (repo *BaseRepo) Init(db *gorm.DB, tableName string) {
	repo.DB = db.Table(tableName)
}

func (repo *BaseRepo) Insert(data interface{}) error {
	return repo.DB.Create(data).Error
}

func (repo *BaseRepo) Update(id uint, data interface{}) error {
	return repo.DB.Where("id = ?", id).Updates(data).Error
}

func (repo *BaseRepo) FindById(dest interface{}) error {
	return repo.DB.First(dest).Error
}
