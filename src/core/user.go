package core

import (
	"context"

	"github.com/uticket/rest"
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

func (m *UserManager) UserGetManager(ctx context.Context) ([]entity.Usuario, error) {
	request, err := db.FindUser(ctx)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (m *UserManager) FindUserByIdManager(ctx context.Context, id string) (entity.Usuario, error) {
	request, err := db.FindById(ctx, id)
	if err != nil {
		return request, &rest.Error{Status: 400, Code: "id_consult_invalid", Message: "INVALID ID"}
	}
	return request, nil
}
