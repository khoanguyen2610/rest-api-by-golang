package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/routers"
)

func handleRequest() {
	fmt.Println("Start listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", routers.InitRouter()))
}

func main() {
	handleRequest()
}
