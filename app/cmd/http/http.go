package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"user-service/configs"
	"user-service/connections"
	"user-service/routers"
	"user-service/routers/middleware"
	"user-service/utils/env"
)

func main() {
	appConf := configs.GetAppConfigFromEnv()

	// init DB Connection
	mysqlConfig := connections.MysqlConfig{
		Host:   appConf.DB.Host,
		Port:   appConf.DB.Port,
		DbName: appConf.DB.DbName,
		User:   appConf.DB.User,
		Pass:   appConf.DB.Pass,
	}
	db := connections.ConnectDB(mysqlConfig)

	envFactory := func(ctx context.Context) *env.Env {
		return env.NewEnv(&appConf, db, ctx)
	}

	// init Routes
	r := mux.NewRouter()
	r.Use(middleware.Env(envFactory))

	// init Sub-Routes with prefix "/api"
	r.PathPrefix("/api").Handler(
		http.StripPrefix(
			strings.TrimSuffix("/api", "/"),
			routers.InitRouter(),
		),
	)

	fmt.Println("Start listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}