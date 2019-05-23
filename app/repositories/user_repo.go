package repositories

import (
	"user-service/app/models"
)

type UserRepo interface {
	FindById(id int) (models.User, error)
}
