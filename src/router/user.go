package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/uticket/rest"
	"go.mod/src/core"
	"go.mod/src/entity"
)

func userRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", postUser)
	return r
}

func postUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type personJson struct {
		Name string `json:"name"`
	}
	var body personJson

	if err := rest.ParseBody(w, r, &body); err != nil {
		return
	}
	person := entity.Usuario{
		Id:        uuid.New(),
		Name:      body.Name,
		CreatedAt: time.Now(),
	}
	manager := core.NewUser()
	sendManager := manager.UserPostManager(ctx, person)
	rest.Send(w, sendManager)
}
