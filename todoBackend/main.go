package main

import (
	"fmt"
	"log"
	"net/http"
	handlers "todolist/handler"
	"todolist/mongodb"
)

func main() {
	connectionURI := "mongodb://localhost:27017/"
	err := mongodb.GetMongodbEngine(connectionURI)
	if err != nil {
		panic(err)
	}

	handler :=handlers.SetupHandlers()
	fmt.Println("serving at port 8081")
	log.Fatal(http.ListenAndServe(":8081", handler))
}
