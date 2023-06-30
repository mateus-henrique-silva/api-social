package core

import (
	"context"

	"github.com/uticket/rest"
	"go.mod/src/db"
	"go.mod/src/entity"
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
