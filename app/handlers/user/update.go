package user

import (
	"fmt"
	"net/http"
	"user-service/handlers"
	"user-service/handlers/response"
	"user-service/models"
	"user-service/handlers/user/form"
)

func Update(ctx handlers.Context) response.ApiResponse {
	payload := form.UserPutForm{}
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

	id := ctx.URLParam("id")
	repo := ctx.BaseRepo()
	var user models.User
	err := repo.FindById(&user, id)
	if err != nil {
		return response.RecordNotFoundError(err, "User not found", http.StatusNotFound)
	}

	if err := repo.Update(&user, payload); err != nil {
		return response.BadRequest(err)
	} else {
		return response.Ok(user)
	}
}
