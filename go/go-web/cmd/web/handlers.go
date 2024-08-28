package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.clientError(w, http.StatusInternalServerError)
		return
	}

	// w.WriteHeader(http.StatusContinue)
	// w.Write([]byte("Hello, web"))
}

func (app *application) posts(w http.ResponseWriter, r *http.Request) {
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
