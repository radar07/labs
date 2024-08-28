package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	file, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	infoLog := log.New(file, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// If `NewServeMux` is specified it will use `DefaultServeMux`
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	// Serve Static Files
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static"))))

	// mux.HandleFunc("/posts", posts)
	// mux.HandleFunc("/posts?id=3", getPost)
	infoLog.Printf("Listening on port", *addr)
	errorLog.Fatal(http.ListenAndServe(*addr, mux))
}
