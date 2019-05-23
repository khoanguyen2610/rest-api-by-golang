package main

import (
	// "fmt"
	"log"
	"net/http"
	"user-service/routers"
	
)
func handleRequest() {
	
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
	// argument
	// fmt.Printf("%T", routers)
    log.Fatal(http.ListenAndServe(":8080", routers.InitRouter()))
}

func main() {
	handleRequest()
}
