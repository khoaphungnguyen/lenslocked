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
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	user.Template.New.Execute(w, data)
}

func (user Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Email:", r.FormValue("email"))
	fmt.Fprint(w, "Password:", r.FormValue("password"))
}
