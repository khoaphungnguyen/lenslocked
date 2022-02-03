package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/khoaphungnguyen/lenslocked/controllers"
	"github.com/khoaphungnguyen/lenslocked/templates"
	"github.com/khoaphungnguyen/lenslocked/views"
)

type User struct {
	Name string
	Bio  string
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	home := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(home))

	contact := views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(contact))

	fag := views.Must(views.ParseFS(templates.FS, "fag.gohtml", "tailwind.gohtml"))
	r.Get("/fag", controllers.FAQ(fag))

	signup := views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", controllers.FAQ(signup))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on: 3000....")
	http.ListenAndServe(":3000", r)
}
