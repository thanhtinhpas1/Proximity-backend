package repositories

import "gorm.io/gorm"

type UserRepo struct {
	BaseRepo
}

func (repo *UserRepo) Init(db *gorm.DB) {
	repo = &UserRepo{}
	repo.BaseRepo.Init(db, "users")
}
