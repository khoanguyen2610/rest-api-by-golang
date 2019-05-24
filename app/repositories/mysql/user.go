package mysql

import (
	"github.com/jinzhu/gorm"
)

type ProductRepo struct {
	BaseRepo
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{
		BaseRepo{
			db:  db,
		},
	}
}
