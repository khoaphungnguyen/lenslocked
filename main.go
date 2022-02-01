package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/khoaphungnguyen/lenslocked/controllers"
	"github.com/khoaphungnguyen/lenslocked/views"
)

type User struct {
	Name string
	Bio  string
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	home, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(home))

	contact, err := views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(contact))

	tpl, err := views.Parse(filepath.Join("templates", "fag.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/fag", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on: 3000....")
	http.ListenAndServe(":3000", r)
}
