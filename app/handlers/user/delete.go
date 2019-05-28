package user

import (
	"net/http"
	"user-service/handlers"
	"user-service/handlers/response"
	"user-service/models"
	//"user-service/repositories/mysql"
)

func Delete(ctx handlers.Context) response.ApiResponse {
	id := ctx.URLParam("id")
	repo := ctx.BaseRepo()
	var user models.User
	err := repo.FindById(&user, id)
	if err != nil {
		return response.RecordNotFoundError(err, "User not found", http.StatusNotFound)
	}

	if err := repo.Delete(&user); err != nil {
		return response.BadRequest(err)
	} else {
		return response.Ok(user)
	}



	//userRepo := mysql.NewUserRepo(ctx.DB())
	//
	//err := userRepo.FindById(&user, id)
	//if err != nil {
	//	return response.RecordNotFoundError(err, "User not found", http.StatusNotFound)
	//}
	//
	//if err := userRepo.Delete(&user); err != nil {
	//	return response.BadRequest(err)
	//} else {
	//	return response.Ok(user)
	//}
}
