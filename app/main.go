package main

import (
	"fmt"
	"log"
	"net/http"
	"user-service/routers"
	
)
func handleRequest() {
	fmt.Println("Start listening...")
    log.Fatal(http.ListenAndServe(":8080", routers.InitRouter()))
}

func main() {
	handleRequest()
}
