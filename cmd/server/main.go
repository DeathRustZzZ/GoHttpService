package main

import (
	"fmt"
	"net/http"
)

func mainPage(res http.ResponseWriter, req *http.Request) {
	fmt.Println("mainPage:", req.URL.Path)
	res.Write([]byte("Hello !"))
}

func apiPage(res http.ResponseWriter, req *http.Request) {
	fmt.Println("apiPage:", req.URL.Path)
	res.Write([]byte("This page for /api"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(`GET /api`, apiPage)
	mux.HandleFunc(`GET /`, mainPage)

	fmt.Println("Server started: http://localhost:8080")
	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
