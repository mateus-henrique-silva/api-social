package core

import (
	"context"

	"go.mod/src/db"
	"go.mod/src/entity"
	"go.mod/src/rest"
)

type CommentsManager struct {
}

func NewCommentsManager() *CommentsManager {
	return &CommentsManager{}
}

func (c *CommentsManager) CreateCommentManager(ctx context.Context, comments entity.Comments) error {
	err := db.CrateComment(ctx, comments)
	if err != nil {
		return &rest.Error{Status: 400, Code: "erro_consult", Message: "Erro ao cadastrar"}
	}
	return nil
}
