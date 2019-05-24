package routers

import (
	"net/http"
	"user-service/handlers/user"
	"user-service/routers/middleware"

	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user", middleware.MakeHandler(user.List)).Methods("GET")
	myRouter.HandleFunc("/user/{id}", middleware.MakeHandler(user.Get)).Methods("GET")
	myRouter.HandleFunc("/user", middleware.MakeHandler(user.Create)).Methods("POST")
	myRouter.HandleFunc("/user/{id}", middleware.MakeHandler(user.Update)).Methods("PUT")
	return myRouter
}
