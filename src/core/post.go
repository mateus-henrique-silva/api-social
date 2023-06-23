package core

import (
	"context"

	"github.com/uticket/rest"
	"go.mod/src/db"
	"go.mod/src/entity"
)

type PostManager struct {
}

func NewPostManager() *PostManager {
	return &PostManager{}
}

func (m *PostManager) InsertPost(ctx context.Context, body entity.Post) error {
	err := db.CreatePost(ctx, body)
	if err != nil {
		return &rest.Error{Status: 400, Code: "error creating post", Message: "Error ao criar o post no blog"}
	}
	return nil
}
