package main

import (
	"log"
	"net/http"
)

func main() {
	// If `NewServeMux` is specified it will use `DefaultServeMux`
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// http.HandleFunc("/posts", posts)
	// http.HandleFunc("/posts?id=3", getPost)
	log.Fatal(http.ListenAndServe(":8090", mux))
}
