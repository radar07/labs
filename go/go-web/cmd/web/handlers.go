package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Query())
	if r.URL.Path != "/" {
		http.Error(w, "Page not found", 404)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusContinue)
	// w.Write([]byte("Hello, web"))
}

func posts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Page not found!"))
		return
	}
}

func getPost(w http.ResponseWriter, r *http.Request) {
	// Use `URL.Query` to get the the URL paramaters
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Page not found!"))
		return
	}

}
