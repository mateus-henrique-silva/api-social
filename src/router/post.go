package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/uticket/rest"
	"go.mod/src/core"
	"go.mod/src/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func postRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", postAddHandler)

	return r
}

func postAddHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type personJson struct {
		Title    string                    `json:"title"`
		Text     string                    `json:"text"`
		NameUser string                    `json:"name"`
		Image    entity.PostImageMultiples `json:"image"`
	}
	var body personJson

	post := entity.Post{
		ID:        primitive.NewObjectID(),
		Title:     body.Title,
		Text:      body.Text,
		NameUser:  body.NameUser,
		Image:     body.Image,
		CreatedAt: time.Now(),
	}
	manager := core.NewPostManager()
	err := manager.InsertPost(ctx, post)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, nil)
}
