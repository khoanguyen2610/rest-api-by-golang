package user

import (
	"net/http"

	"user-service/handlers"
	"user-service/handlers/response"
	"user-service/models"
)

func List(ctx handlers.Context) response.ApiResponse {
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON

	repo := ctx.BaseRepo()
	var users []models.User
	err := repo.Search(&users)
	if err != nil {
		return response.RecordNotFoundError(err, "User not found", http.StatusNotFound)
	}
	return response.Ok(users)
}
