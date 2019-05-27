package validation

import (
	"gopkg.in/go-playground/validator.v9"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	return validate
}