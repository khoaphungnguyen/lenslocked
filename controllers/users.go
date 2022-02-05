package controllers

import (
	"net/http"

	"github.com/khoaphungnguyen/lenslocked/views"
)

type Users struct {
	Template struct {
		New views.Template
	}
}

func (user Users) New(w http.ResponseWriter, r *http.Request) {
	// We need a view to render the users
	user.Template.New.Execute(w, nil)
}
