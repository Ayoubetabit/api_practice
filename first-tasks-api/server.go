package main

import (
	"fmt"
	"net/http"
)

func HelloHhandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received a request :", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Hello ayoub the go api is working well!")
}

func HomeRoot(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received a request :", r.Method, r.URL.Path)
	w.Header().Set("Contetnt-type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Ayoube The home root is working")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a request :", r.Method, r.URL.Path)
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	for i := 0; i < 10; i++{
		fmt.Fprintln(w, "User", i)
	}
}
