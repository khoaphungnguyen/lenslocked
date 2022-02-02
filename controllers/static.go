package controllers

import (
	"html/template"
	"net/http"

	"github.com/khoaphungnguyen/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes: We offer a free trial for 30 days on any paid plans.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have support staff answering email 24/7, though response time may be a bit slower on weekends.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us -  <a href="mailto:rex@test.com">rex@test.com</a>.</p>`,
		},
		{
			Question: "Where is your office?",
			Answer:   "Our team is entire remote",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
