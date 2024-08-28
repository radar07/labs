package main

import (
	"log"
	"net/http"
)

func main() {
	// If `NewServeMux` is specified it will use `DefaultServeMux`
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Serve Static Files
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	// http.HandleFunc("/posts", posts)
	// http.HandleFunc("/posts?id=3", getPost)
	log.Fatal(http.ListenAndServe(":8090", mux))
}
