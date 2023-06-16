package core

import (
	"context"

	"go.mod/src/db"
	"go.mod/src/entity"
)

type UserManager struct {
}

func NewUser() *UserManager {
	return &UserManager{}
}
func (m *UserManager) UserPostManager(ctx context.Context, person entity.Usuario) error {
	_, err := db.CreateUser(ctx, person)
	if err != nil {
		return err
	}
	return nil
}
