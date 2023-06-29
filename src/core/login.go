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

func (m *ManagerLogin) ManagerLoginVerify(ctx context.Context, body entity.Login) (string, error) {
	fmt.Println(body.Email)
	fmt.Println(body.Password)
	validEmail, err := db.CheckIfUserEmailExists(ctx, body.Email)
	if err != nil {
		return "", &rest.Error{Status: 400, Code: "Email_check_in_email", Message: "Erro checking in email"}
	}
	valid, err := db.CheckInPasswordValidLogin(ctx, body.Email, body.Password)
	if err != nil {
		return "", &rest.Error{Status: 400, Code: "Invalid_check_in_password", Message: "Error checking in password"}
	}
	if !validEmail {
		return "", &rest.Error{Status: 400, Code: "erro_consult", Message: "email_invalid"}
	}
	if !valid {
		return "", &rest.Error{Status: 400, Code: "erro_consult", Message: "password_invalid"}
	}
	var validation bool
	fmt.Println(validEmail)
	fmt.Println(valid)
	fmt.Println(validation)
	var token string
	if valid == true && validEmail == true {
		validation = true
		userId, err := db.ReturnByIdLogin(ctx, body.Email)
		fmt.Println(userId)
		token, err := autenticacao.CreateToken(ctx, userId)
		fmt.Println(token)
		if err != nil {
			return "", err
		}
		fmt.Println(token)
		return "", nil
	}
	return token, nil
}
