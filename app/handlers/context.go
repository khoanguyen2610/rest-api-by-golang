package handlers

import (
	"net/http"

	"user-service/handlers/response"

	"github.com/gorilla/mux"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) Context {
	return Context{w, r}
}

// URLParam Get Url param
func (c *Context) URLParam(key string) string {
	vars := mux.Vars(c.r)
	return vars[key]
}

// Func Handler function
type HandlerFunc func(ctx Context) response.ApiResponse
