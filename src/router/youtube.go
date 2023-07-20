package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.mod/src/core"
	"go.mod/src/entity"
	"go.mod/src/rest"
)

func YoutbeRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", youtubeHandler)
	r.Post("/", youtubePostHandler)
	r.Put("/{id}", youtubePutHandler)
	return r
}

func youtubeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	manager := core.NewYoutubeManager()
	sendManager, err := manager.GetYoutube(ctx)
	if err != nil {
		rest.SendError(w, err)
	}
	rest.Send(w, sendManager)
}

func youtubePostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type body struct {
		VideoID    string `json:"video"`
		Status     string `json:"status"`
		AuthorName string `json:"authorName"`
		Title      string `json:"title"`
		Summary    string `json:"summary"`
	}
	var bodyJson body
	result := entity.Youtube{
		VideoID:    bodyJson.VideoID,
		Status:     bodyJson.Status,
		AuthorName: bodyJson.AuthorName,
		Title:      bodyJson.Title,
		CreatedAt:  time.Now(),
		EditedAt:   time.Now(),
		Summary:    bodyJson.Summary,
	}
	manager := core.NewYoutubeManager()
	err := manager.PostYoutube(ctx, result)
	if err != nil {
		rest.SendError(w, err)
		return
	}
}

func youtubePutHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	type body struct {
		VideoID    string `json:"video"`
		Status     string `json:"status"`
		AuthorName string `json:"authorName"`
		Title      string `json:"title"`
		Summary    string `json:"summary"`
	}
	var bodyJson body
	result := entity.Youtube{
		VideoID:    bodyJson.VideoID,
		Status:     bodyJson.Status,
		AuthorName: bodyJson.AuthorName,
		Title:      bodyJson.Title,
		CreatedAt:  time.Now(),
		EditedAt:   time.Now(),
		Summary:    bodyJson.Summary,
	}
	manager := core.NewYoutubeManager()
	err := manager.PutYoutube(ctx, result, id)
	if err != nil {
		rest.SendError(w, err)
		return
	}
}
