package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/uticket/rest"
	"go.mod/src/core"
	"go.mod/src/entity"
)

func userRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", postUserHandler)
	r.Get("/", getUserHandler)
	r.Delete("/{id}", deleteUserHandler)
	r.Put("/{id}", putUserHandler)
	return r
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type personJson struct {
		Name   string `json:"name"`
		Email  string `json:"email"`
		Number string `json:"number"`
	}
	var body personJson

	if err := rest.ParseBody(w, r, &body); err != nil {
		return
	}
	person := entity.Usuario{
		Name:      body.Name,
		Email:     body.Email,
		Number:    body.Number,
		CreatedAt: time.Now(),
	}
	manager := core.NewUser()
	sendManager := manager.UserPostManager(ctx, person)
	rest.Send(w, sendManager)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	manager := core.NewUser()
	searchId := r.URL.Query().Get("_id")
	fmt.Println(searchId)
	if searchId == "" {
		sendManger, err := manager.UserGetManager(ctx)
		if err != nil {
			rest.SendError(w, err)
			return
		}
		rest.Send(w, sendManger)

	} else {
		manager := core.NewUser()
		sendById, err := manager.FindUserByIdManager(ctx, searchId)
		if err != nil {
			rest.SendError(w, err)
			return
		}
		rest.Send(w, sendById)

	}

}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	manager := core.NewUser()
	err := manager.RemoveByIdManager(ctx, id)
	if err != nil {
		rest.SendError(w, err)
		return
	}
}

func putUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := chi.URLParam(r, "id")
	type bodyJson struct {
		Name string `json:"name"`
	}
	var bodyJsonBody entity.Usuario
	if err := rest.ParseBody(w, r, &bodyJsonBody); err != nil {
		return
	}
	person := entity.Usuario{
		Name: bodyJsonBody.Name,
	}
	manager := core.NewUser()
	err := manager.UpdateByIdManager(ctx, id, person)
	if err != nil {
		rest.SendError(w, err)
		return
	}
}
