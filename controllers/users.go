package controllers

import (
	"fmt"
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

func (user Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email:", r.FormValue("email"))
	fmt.Fprint(w, "Password:", r.FormValue("password"))
}
