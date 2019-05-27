package repositories

import (
	"user-service/models"
)

type BaseRepo interface {
	Searchable
	Updatable
	Creatable
	CanFindById
	CanCreateOrUpdate
}

type Searchable interface {
	Search(val interface{}) error
}

type Updatable interface {
	Update(m models.Model, attrs ...interface{}) error
}

type Creatable interface {
	Create(m models.Model) error
}

type CanFindById interface {
	FindById(m models.Model, id string) error
}

type CanCreateOrUpdate interface {
	CreateOrUpdate(m models.Model, query interface{}, attrs ...interface{}) error
}