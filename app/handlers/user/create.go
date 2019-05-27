package user

import (
	"fmt"
	"user-service/handlers"
	"user-service/handlers/response"
	"user-service/models"
	"user-service/handlers/user/form"
)


func Create(ctx handlers.Context) response.ApiResponse {
	payload := form.UserPostForm{}
	// Parse input
	if err := ctx.DecodePayload(&payload); err != nil {
		return response.BadRequest(
			fmt.Errorf("cannot parse JSON to object: %s", err.Error()),
		)
	}

	// Validate input
	if err := ctx.Validate(payload); err != nil {
		return response.ValidationError(err)
	}

	repo := ctx.BaseRepo()
	var user models.User
	if err := ctx.ConvertForm2Model(payload, &user); err != nil {
		return response.BadRequest(
			fmt.Errorf("cannot parse input to object: %s", err.Error()),
		)
	}

	if err := repo.Create(&user); err != nil {
		return response.BadRequest(err)
	} else {
		return response.Ok(user)
	}

}
