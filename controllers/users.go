package controllers

import (
	"net/http"
)

type Users struct {
	Template struct {
		New Template
	}
}

func (user Users) New(w http.ResponseWriter, r *http.Request) {
	// We need a view to render the users
	user.Template.New.Execute(w, nil)
}
