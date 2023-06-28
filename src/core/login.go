package core

import (
	"context"
	"fmt"

	"github.com/uticket/rest"
	"go.mod/src/autenticacao"
	"go.mod/src/db"
	"go.mod/src/entity"
)

type ManagerLogin struct {
}

func NewManagerLogin() *ManagerLogin {
	return &ManagerLogin{}
}

func (m *ManagerLogin) ManagerLoginVerify(ctx context.Context, body entity.Login) (bool, error) {
	fmt.Println(body.Email)
	fmt.Println(body.Password)
	validEmail, err := db.CheckIfUserEmailExists(ctx, body.Email)
	if err != nil {
		return false, &rest.Error{Status: 400, Code: "Email_check_in_email", Message: "Erro checking in email"}
	}
	valid, err := db.CheckInPasswordValidLogin(ctx, body.Email, body.Password)
	if err != nil {
		return false, &rest.Error{Status: 400, Code: "Invalid_check_in_password", Message: "Error checking in password"}
	}

	var validation bool
	fmt.Println(validEmail)
	fmt.Println(valid)
	fmt.Println(validation)
	if valid == true && validEmail == true {
		validation = true
		userId, err := db.ReturnByIdLogin(ctx, body.Email)
		fmt.Println(userId)
		token, err := autenticacao.CreateToken(ctx, userId)
		fmt.Println(token)
		if err != nil {
			return false, err
		}
		fmt.Println(token)
		return true, nil
	}
	return validation, nil
}
