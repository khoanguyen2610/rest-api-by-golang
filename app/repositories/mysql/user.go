package mysql

import (
	"github.com/jinzhu/gorm"

	"user-service/models"
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

func (r *UserRepo) Delete(m models.Model) error {
	return r.db.Delete(m).Error
}