package handlers

import (
	"encoding/json"
	"net/http"
	"user-service/repositories"
	"user-service/repositories/mysql"
	"user-service/utils/env"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"user-service/handlers/response"
)

type Context struct {
	e *env.Env
	w http.ResponseWriter
	r *http.Request
}

func NewContext(e *env.Env, w http.ResponseWriter, r *http.Request) Context {
	return Context{e, w, r}
}

// DB Get db from ENV
func (c *Context) DB() *gorm.DB {
	return c.e.GetDB()
}

// Get ENV
func (c *Context) Env() *env.Env {
	return c.e
}

// BaseRepo Returns base repository
func (c *Context) BaseRepo() repositories.BaseRepo {
	return mysql.NewBaseRepo(c.DB())
}

// URLParam Get Url param
func (c *Context) URLParam(key string) string {
	vars := mux.Vars(c.r)
	return vars[key]
}

// DecodePayload Decode payload to target
func (c *Context) DecodePayload(target interface{}) error {
	decoder := json.NewDecoder(c.r.Body)
	return decoder.Decode(target)
}

// Func Handler function
type HandlerFunc func(ctx Context) response.ApiResponse
