package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/khoaphungnguyen/lenslocked/views"
)

func executeTemplate(w http.ResponseWriter, filename string) {
	t, err := views.Parse(filename)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "template/home.gohtml")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "template/contact.gohtml")

}

func fagHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "template/fag.gohtml")
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
