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

func categoryRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", postCategoryHandler)
	r.Get("/", getCategoryHandler)
	r.Put("/", putCategoryHandler)
	return r
}

func postCategoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type personJson struct {
		Name string `json:"name"`
	}
	var json personJson
	if err := rest.ParseBody(w, r, &json); err != nil {
		return
	}
	body := entity.Category{
		ID:        primitive.NewObjectID(),
		Name:      json.Name,
		CreatedAt: time.Now(),
	}
	manager := core.NewCategoryManager()
	sendManager, err := manager.CreateCategoryManager(ctx, body)
	if err != nil {
		return
	}
	rest.Send(w, sendManager)

}

func getCategoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	manager := core.NewCategoryManager()
	sendManager, err := manager.FindCategoryManager(ctx)
	if err != nil {
		return
	}
	rest.Send(w, sendManager)
}

func putCategoryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type personJson struct {
		Name string `json:"name"`
	}
	var body entity.Category
	if err := rest.ParseBody(w, r, &body); err != nil {
		return
	}
	send := entity.Category{
		Name: body.Name,
	}
	manager := core.NewCategoryManager()
	sendManger, err := manager.UpdateCategoryManager(ctx, send)
	if err != nil {
		return
	}
	rest.Send(w, sendManger)

}
