package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", HelloHhandler)
	mux.HandleFunc("GET /users", GetUsers)
	mux.HandleFunc("GET /", HomeRoot)

	log.Println("Server listening at http://localhost:8080")

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}