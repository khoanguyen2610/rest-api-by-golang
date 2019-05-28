package mysql

import (
	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	BaseRepo
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		BaseRepo{
			db: db,
		},
	}
}