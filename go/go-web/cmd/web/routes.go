package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// If `NewServeMux` is specified it will use `DefaultServeMux`
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	// Serve Static Files
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	// mux.HandleFunc("/posts", posts)
	// mux.HandleFunc("/posts?id=3", getPost)

	return mux
}
