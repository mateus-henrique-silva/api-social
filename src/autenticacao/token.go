package autenticacao

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"go.mod/src/connect"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Criar Token cria o token de autenticacao.
func CreateToken(ctx context.Context, usuarioID primitive.ObjectID) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	id, err := ExtractNumberFromObjectID(usuarioID)
	if err != nil {
		return "", err
	}
	fmt.Println(id)
	permission["usuarioId"] = id
	// Secret Key
	// fmt.Println(connect.SecretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)
	str, err := token.SignedString([]byte(connect.SecretKey))
	if err != nil {
		return "", err
	}
	return str, nil
}

func ValidateToken(r *http.Request) (bool, error) {
	tokenString := extraction(r)
	token, err := jwt.Parse(tokenString, returnKeyVerification)
	if err != nil {
		return false, err
	}
	fmt.Println(token)
	return true, nil
}

func extraction(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnKeyVerification(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	}
	return []byte("Secret"), nil
}

func ExtractNumberFromObjectID(objectID primitive.ObjectID) (int64, error) {
	objectIDString := objectID.Hex()
	bytes, err := hex.DecodeString(objectIDString)
	if err != nil {
		return 0, err
	}

	var number big.Int
	number.SetBytes(bytes)
	return number.Int64(), nil
}
