package core

import (
	"context"

	"go.mod/src/db"
	"go.mod/src/entity"
	"go.mod/src/rest"
	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
}

func NewUser() *UserManager {
	return &UserManager{}
}
func (m *UserManager) UserPostManager(ctx context.Context, person entity.Usuario) error {
	password, err := HashPassword(person.Password)
	if err != nil {
		return err
	}
	person.Password = password
	consult, err := db.CheckIfUserExists(ctx, person.Name)
	if err != nil {
		return &rest.Error{Status: 400, Code: "error", Message: "consult error "}
	}
	if consult {
		return &rest.Error{Status: 400, Code: "user_exists", Message: "consulta retorna que o nome ja existe"}
	}
	existEmail, err := db.CheckIfUserEmailExists(ctx, person.Email)
	if err != nil {
		return &rest.Error{Status: 400, Code: "error", Message: "consult error "}
	}
	if existEmail {
		return &rest.Error{Status: 400, Code: "email_exists", Message: "consulta retorna que o email ja existe"}
	}
	_, err = db.CreateUser(ctx, person)
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (m *UserManager) UserGetManager(ctx context.Context) ([]entity.Usuario, error) {
	request, err := db.FindUsers(ctx)
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

func (m *UserManager) RemoveByIdManager(ctx context.Context, id string) error {
	err := db.RemoveById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserManager) UpdateByIdManager(ctx context.Context, id string, person entity.Usuario) error {
	err := db.UpdateById(ctx, id, person)
	if err != nil {
		return err
	}
	return nil
}
