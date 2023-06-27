package autenticacao

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Criar Token cria o token de autenticacao.
func CreateToken(ctx context.Context, usuarioID uint64) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permission["usuarioId"] = usuarioID
	//Secret
	token := jwt.NewWithClaims(jwt.SigningMethodES256, permission)
	return token.SignedString([]byte("Secret"))

}
