package main

import (
	"log"
	"net/http"

	"example.com/hello/Desktop/go-workspace/pkg/mod/github.com/!akhil!sharma90/!golang-!my!s!q!l-!c!r!u!d-!bookstore-!management-!a!p!i@v0.0.0-20211006083836-71e755118f4f/pkg/routes"
	"example.com/hello/Desktop/go-workspace/pkg/mod/github.com/gorilla/mux@v1.8.0"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
