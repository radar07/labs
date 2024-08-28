package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	// If `NewServeMux` is specified it will use `DefaultServeMux`
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Serve Static Files
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	// http.HandleFunc("/posts", posts)
	// http.HandleFunc("/posts?id=3", getPost)
	log.Fatal(http.ListenAndServe(*addr, mux))
}
