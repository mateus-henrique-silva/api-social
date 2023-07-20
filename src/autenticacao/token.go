package autenticacao

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateToken cria o token de autenticação.
func CreateToken(ctx context.Context, usuarioID primitive.ObjectID) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Hour * 6).Unix()
	id, err := ExtractNumberFromObjectID(usuarioID)
	if err != nil {
		return "", err
	}
	permission["usuarioId"] = id
	SecretKey := []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)
	str, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return str, nil
}

func ValidateToken(r *http.Request) (bool, error) {
	tokenString := extraction(r)
	if tokenString == "" {
		return false, errors.New("Token de autenticação não fornecido")
	}

	token, err := jwt.Parse(tokenString, returnKeyVerification)
	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, errors.New("Token de autenticação inválido")
	}

	return true, nil
}

func extraction(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if token != "" {
		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) == 2 {
			return splitToken[1]
		}
	}
	return ""
}

func returnKeyVerification(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado")
	}
	SecretKey := []byte(os.Getenv("SECRET_KEY"))
	return []byte(SecretKey), nil
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
