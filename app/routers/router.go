package routers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"user-service/handlers/user"
	"user-service/routers/middleware"

	"github.com/gorilla/mux"
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

func allArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	fmt.Println("Endpoint Hit: All Articles ")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err == nil {
		log.Print(err)
	}

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func InitRouter() http.Handler {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", middleware.MakeHandler(user.GetOne)).Methods("GET")
	myRouter.HandleFunc("/articles", allArticles).Methods("POST")
	return myRouter
}
