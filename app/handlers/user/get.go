package user

import (
	"net/http"

	"user-service/handlers"
	"user-service/handlers/response"
	"user-service/models"
)

func Get(ctx handlers.Context) response.ApiResponse {
	id := ctx.URLParam("id")
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON

	repo := ctx.BaseRepo()
	var u models.User
	err := repo.FindById(&u, id)
	if err != nil {
		return response.RecordNotFoundError(err, "User not found", http.StatusNotFound)
	}
	return response.Ok(u)
}
