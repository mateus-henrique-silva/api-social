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

func CommentsRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", postCommentsHandler)
	return r
}

func postCommentsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type body struct {
		IdPost       string `json:"idPost"`
		UserComments string `json:"userComments"`
		CommentsText string `json:"commentsText"`
	}

	var bodyJson body
	if err := rest.ParseBody(w, r, &bodyJson); err != nil {
		return
	}
	send := entity.Comments{
		Id:           primitive.NewObjectID(),
		IdPost:       bodyJson.IdPost,
		UserComments: bodyJson.UserComments,
		CommentsText: bodyJson.CommentsText,
		CreatedAt:    time.Now(),
	}
	manager := core.NewCommentsManager()
	err := manager.CreateCommentManager(ctx, send)
	if err != nil {
		rest.SendError(w, err)
		return
	}
}
