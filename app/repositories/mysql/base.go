package mysql

import (
	"github.com/jinzhu/gorm"

	"user-service/models"
)

type BaseRepo struct {
	db *gorm.DB
}

func NewBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{
		db: db,
	}
}

func (r *BaseRepo) FindById(m models.Model, id string) error {
	return r.db.Where("id = ?", id).First(m).Error
}

func (r *BaseRepo) CreateOrUpdate(m models.Model, query interface{}, attrs ...interface{}) error {
	return r.db.Where(query).Assign(attrs...).FirstOrCreate(m).Error
}

func (r *BaseRepo) Update(m models.Model, attrs ...interface{}) error {
	return r.db.Model(m).Update(attrs...).Error
}

func (r *BaseRepo) Create(m models.Model) error {
	return r.db.Create(m).Error
}

func (r *BaseRepo) Search(val interface{}) error {
	q := r.db.Model(val)
	//for query, val := range f.GetWhere() {
	//	q = q.Where(query, val...)
	//}
	//
	//if f.GetLimit() > 0 {
	//	q = q.Limit(f.GetLimit())
	//}
	//
	//q = q.Offset(f.GetOffset())

	return q.Find(val).Error
}
