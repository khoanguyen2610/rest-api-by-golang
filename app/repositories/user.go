package repositories

import (
	"user-service/models"
)

type UserRepo interface {
	BaseRepo
	Delete(m models.Model) error
}
