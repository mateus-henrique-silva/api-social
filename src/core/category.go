package core

import (
	"context"

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
