package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var templates = template.Must(template.ParseGlob("templates/*.tmpl"))

type PageData struct {
	Title   string
	Message string
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		log.Println("template execute error:", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := PageData{Title: "Go HTMX-like App", Message: "Welcome to the Go standard-library web app."}
	renderTemplate(w, "index.tmpl", data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "About", Message: "This app uses only the Go standard library: net/http and html/template."}
	renderTemplate(w, "about.tmpl", data)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Time", Message: time.Now().Format(time.RFC1123)}
	renderTemplate(w, "time.tmpl", data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/time", timeHandler)

	// Static files under /static/
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	addr := ":8080"
	log.Printf("Starting server on %s — visit http://localhost:8080/", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
