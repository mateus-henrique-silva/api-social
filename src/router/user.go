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
	r.Post("/", postUserHandler)
	r.Get("/", getUserHandler)
	return r
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
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

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	manager := core.NewUser()
	searchId := r.URL.Query().Get("id")
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
	// jsonWrite(w, http.StatusOK, result)
}
