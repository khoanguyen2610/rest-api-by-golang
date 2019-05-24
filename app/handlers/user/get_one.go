package user

import (
	"log"
	"net/http"
	"strconv"

	"user-service/handlers"
	"user-service/handlers/response"
	"user-service/models"
	"user-service/connections"
	"user-service/repositories/mysql"
)

type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles = []Article{
	Article{Id: 1, Title: "Title 001", Desc: "Description 001", Content: "Content 001"},
	Article{Id: 2, Title: "Title 002", Desc: "Description 002", Content: "Content 002"},
	Article{Id: 3, Title: "Title 003", Desc: "Description 003", Content: "Content 003"},
	Article{Id: 4, Title: "Title 004", Desc: "Description 004", Content: "Content 004"},
	Article{Id: 5, Title: "Title 005", Desc: "Description 005", Content: "Content 005"},
}

func GetOne(ctx handlers.Context) response.ApiResponse {
	id := ctx.URLParam("id")
	key, err := strconv.Atoi(id)
	if err != nil {
		log.Print(err)
	}
	// Loop over all of our Articles
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

	repo := mysql.NewBaseRepo(db)
	var u models.User
	repo.FindById(&u, id)
	defer db.Close()

	return response.Ok(u)


	for _, article := range Articles {
		if article.Id == key {
			return response.Ok(article)
		}
	}
	return response.ErrorResponse(err, http.StatusNotFound)
}
