package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mod/src/core"
	"go.mod/src/entity"
	"go.mod/src/middleware"
	"go.mod/src/rest"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func postRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", postAddHandler)
	r.Get("/index", middleware.AuthMiddleware(getIndexHandlerFunc(getIndexHandler))) // Aplicando o middleware apenas nesta rota
	r.Get("/index/post-card", getPostsCards)

	return r
}

func postAddHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type personJson struct {
		Title            string                    `json:"title"`
		TitleSlug        string                    `json:"titleSlug"`
		Text             string                    `json:"text"`
		Article          string                    `json:"article"`
		NameUser         string                    `json:"name"`
		AuthorImage      string                    `json:"authorImage"`
		LinkYoutube      string                    `json:"linkYoutube"`
		Category         string                    `json:"category"`
		CategorySlug     string                    `json:"categorySlug"`
		Image            entity.PostImageMultiples `json:"image"`
		BannerAltText    string                    `json:"bannerAltText"`
		CommentsQuantity uint64                    `json:"commentsQuantity"`
		Summary          string                    `json:"summary"`
		Approved         bool                      `json:"approved"`
	}
	var body personJson
	if err := rest.ParseBody(w, r, &body); err != nil {
		rest.SendError(w, err)
		return
	}
	post := entity.Post{
		ID:               primitive.NewObjectID(),
		Title:            body.Title,
		TitleSlug:        body.TitleSlug,
		Text:             body.Text,
		Article:          body.Article,
		NameUser:         body.NameUser,
		Category:         body.Category,
		CategorySlug:     body.CategorySlug,
		Image:            body.Image,
		BannerAltText:    body.BannerAltText,
		CommentsQuantity: body.CommentsQuantity,
		Approved:         body.Approved,
		Summary:          body.Summary,
		CreatedAt:        time.Now(),
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
	idParams := r.URL.Query().Get("slug")
	if idParams == "" {
		send, err := manager.GetIndexHandler(ctx)
		if err != nil {
			rest.SendError(w, err)
			return
		}
		rest.Send(w, send)
	} else {
		send, err := manager.GetIndexPostById(ctx, idParams)
		if err != nil {
			rest.SendError(w, err)
			return
		}
		rest.Send(w, send)
	}

}

type getIndexHandlerFunc func(w http.ResponseWriter, r *http.Request)

func (f getIndexHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func getPostsCards(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	manager := core.NewPostManager()
	sendManager, err := manager.GetIndexPostCards(ctx)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.Send(w, sendManager)

}
