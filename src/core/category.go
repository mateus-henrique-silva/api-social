package core

import (
	"context"

	"github.com/uticket/rest"
	"go.mod/src/db"
	"go.mod/src/entity"
)

type CategoryManager struct {
}

func NewCategoryManager() *CategoryManager {
	return &CategoryManager{}
}

func (m *CategoryManager) CreateCategoryManager(ctx context.Context, body entity.Category) (entity.Category, error) {
	result, err := db.CreateCategory(ctx, body)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *CategoryManager) FindCategoryManager(ctx context.Context) ([]entity.Category, error) {
	result, err := db.FindCategoryAll(ctx)
	if err != nil {
		return nil, &rest.Error{Status: 400, Code: "invalid_consult", Message: "error ao realizar a consulta de retorno de dados pela api"}
	}
	return result, nil
}

func (m *CategoryManager) UpdateCategoryManager(ctx context.Context, body entity.Category, id string) error {
	err := db.UpdateCategory(ctx, id, body)
	if err != nil {
		return err
	}
	return nil
}
