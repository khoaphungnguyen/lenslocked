package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func templateHandler(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("error parsing: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing error: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler(w, "template/home.gohtml")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler(w, "template/contact.gohtml")

}

func fagHandler(w http.ResponseWriter, r *http.Request) {
	templateHandler(w, "template/fag.gohtml")
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.Path)

}

type User struct {
	Name string
	Bio  string
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/articles/{date}", getArticle)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/fag", fagHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on: 3000....")
	http.ListenAndServe(":3000", r)
}
