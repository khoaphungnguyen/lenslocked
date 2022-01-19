package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:rex@test.com\">rex@test.com</a>.</p>")
}

func fagHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>FAG page</h1>
	<p>
	<strong>Q: Is there a free version?</strong><br>
	Yes: We offer a free trial for 30 days on any paid plans. <br>
	<br>
	<strong>Q: What are your support hours?</strong><br>
	We have support staff answering email 24/7, though response time may be a bit slower on weekends. <br>
	<br>
	<strong>Q: How do I contact support?</strong><br>
	Email us -  <a href=\"mailto:rex@test.com\">rex@test.com</a>.</p> <br>
	</p>`)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.Path)

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
