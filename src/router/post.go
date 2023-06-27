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
	r.Get("/index", getIndexHandler)

	return r
}

func postAddHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type personJson struct {
		Title            string                    `json:"title"`
		TitleSlug        string                    `json:"titleSlug"`
		Text             string                    `json:"text"`
		NameUser         string                    `json:"name"`
		Category         string                    `json:"category"`
		Image            entity.PostImageMultiples `json:"image"`
		BannerAltText    string                    `json:"bannerAltText"`
		CommentsQuantity uint64                    `json:"commentsQuantity"`
	}
	var body personJson
	if err := rest.ParseBody(w, r, &body); err != nil {
		rest.SendError(w, err)
		return
	}
	post := entity.Post{
		ID:            primitive.NewObjectID(),
		Title:         body.Title,
		TitleSlug:     body.TitleSlug,
		Text:          body.Text,
		NameUser:      body.NameUser,
		Category:      body.Category,
		Image:         body.Image,
		BannerAltText: body.BannerAltText,
		CreatedAt:     time.Now(),
	}
	manager := core.NewPostManager()
	err := manager.InsertPost(ctx, post)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, nil)
}

func getIndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	manager := core.NewPostManager()
	send, err := manager.GetIndexHandler(ctx)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, send)

}