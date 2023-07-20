package core

import (
	"context"

	"go.mod/src/db"
	"go.mod/src/entity"
	"go.mod/src/rest"
)

type PostManager struct {
}

func NewPostManager() *PostManager {
	return &PostManager{}
}

func (m *PostManager) InsertPost(ctx context.Context, body entity.Post) error {
	exist, err := db.CheckIfPostNameExists(ctx, body.Title)
	if err != nil {
		return &rest.Error{Status: 400, Code: "erro_consult", Message: "erro ao consultar"}
	}
	if exist {
		return &rest.Error{Status: 400, Code: "erro", Message: "nome do post já existe"}
	}
	err = db.CreatePost(ctx, body)
	if err != nil {
		return &rest.Error{Status: 400, Code: "error creating post", Message: "Error ao criar o post no blog"}
	}

	return nil
}

func (m *PostManager) GetIndexHandler(ctx context.Context) ([]entity.PostReturnResponse, error) {
	send, err := db.GetIndexPost(ctx)
	if err != nil {
		return nil, &rest.Error{Status: 400, Code: "erro_consult", Message: "erro ao realizar consulta"}
	}
	return send, nil
}

func (m *PostManager) GetIndexPostById(ctx context.Context, id string) (entity.Post, error) {
	server, err := db.GetPost(ctx, id)
	if err != nil {
		return entity.Post{}, &rest.Error{Status: 400, Code: "erro_consult", Message: "erro ao realizar consulta de retorno de dados da api"}
	}

	return server, nil
}

func (m *PostManager) GetIndexPostCards(ctx context.Context) ([]entity.PostReturnResponse, error) {
	server, err := db.GetPostCard(ctx)
	if err != nil {
		return nil, &rest.Error{Status: 400, Code: "erro_consult", Message: "erro ao realizar consulta de retorno de dados da api"}
	}
	return server, nil
}
