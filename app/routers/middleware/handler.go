package middleware

import (
	"net/http"

	"user-service/handlers"
	"user-service/handlers/response"
)

// MakeHandler Create handler
func MakeHandler(handlerFunc handlers.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := getEnvFromCtx(r.Context())
		ctx := handlers.NewContext(e, w, r)
		res := handlerFunc(ctx)
		response.RenderJson(w, res)
	}
}
