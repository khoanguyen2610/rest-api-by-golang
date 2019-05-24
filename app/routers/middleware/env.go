package middleware

import (
	"context"
	"net/http"

	"user-service/utils/env"
)

type envKeyType int

var envKey envKeyType

// EnvFactory Function that creates env from a contex
type EnvFactory func(ctx context.Context) *env.Env

// Env Creates new env for request and put it in context
func Env(envFactory EnvFactory) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			r = r.WithContext(context.WithValue(ctx, envKey, envFactory(ctx)))
			h.ServeHTTP(w, r)
		})
	}
}

// keep it private to avoid use in strange places
func getEnvFromCtx(ctx context.Context) *env.Env {
	return ctx.Value(envKey).(*env.Env)
}
