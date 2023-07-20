package core

import (
	"context"

	"go.mod/src/db"
	"go.mod/src/entity"
	"go.mod/src/rest"
	"go.mod/src/slug"
)

type CategoryManager struct {
}

func NewCategoryManager() *CategoryManager {
	return &CategoryManager{}
}

func (m *CategoryManager) CreateCategoryManager(ctx context.Context, body entity.Category) (entity.Category, error) {
	exist, err := db.CheckIfCategoryExists(ctx, body.Name)
	if err != nil {
		return entity.Category{}, &rest.Error{Status: 400, Code: "error_consult", Message: "erro ao realizar check"}
	}
	if exist {
		return entity.Category{}, &rest.Error{Status: 400, Code: "error_name_exists", Message: "nome ja existe"}
	}
	body.SlugCategory = slug.RemoveAcentoEspacoCaracterEspecial(body.Name)
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
