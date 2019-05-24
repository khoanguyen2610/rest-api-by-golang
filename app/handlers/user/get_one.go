package user

import (
	"net/http"

	"user-service/handlers"
	"user-service/handlers/response"
	"user-service/models"
	"user-service/connections"
	"user-service/repositories/mysql"
)

func GetOne(ctx handlers.Context) response.ApiResponse {
	id := ctx.URLParam("id")
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	mysqlConfig := connections.MysqlConfig{
		Host:   "localhost",
		Port:   3306,
		DbName: "sample_go",
		User:   "root",
		Pass:   "123456",
	}
	db := connections.ConnectDB(mysqlConfig)
	defer db.Close()

	repo := mysql.NewBaseRepo(db)
	var u models.User
	err := repo.FindById(&u, id)
	if err != nil {
		return response.RecordNotFoundError(err, "User not found", http.StatusNotFound)
	}
	return response.Ok(u)
}
