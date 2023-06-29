package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/uticket/rest"
	"go.mod/src/core"
	"go.mod/src/entity"
)

func LoginRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", loginHandler)

	return r
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type LoginJson struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var body LoginJson
	rest.ParseBody(w, r, &body)
	person := entity.Login{
		Email:    body.Email,
		Password: body.Password,
	}
	manager := core.NewManagerLogin()
	boolean, err := manager.ManagerLoginVerify(ctx, person)
	if err != nil {
		rest.SendError(w, err)
		return
	}

	w.Write([]byte(boolean))

}
