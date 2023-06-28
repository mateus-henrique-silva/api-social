package db

import (
	"context"

	"go.mod/src/connect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CheckInPasswordValidLogin(ctx context.Context, email string, password string) (bool, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return false, err
	}

	collection := client.Database("mydb").Collection("people")

	filter := bson.M{"email": email}
	var person struct {
		PasswordHash string `bson:"password"`
	}
	err = collection.FindOne(ctx, filter).Decode(&person)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil // Usuário não encontrado
		}
		return false, err
	}

	return comparePasswordHash(person.PasswordHash, password), nil
}

func decodePasswordHash(passwordHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}

func ReturnByIdLogin(ctx context.Context, email string) (ID primitive.ObjectID, err error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return ID, err
	}
	collection := client.Database("mydb").Collection("people")
	filter := bson.M{"email": email}
	type personId struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	var personIdJson personId
	err = collection.FindOne(ctx, filter).Decode(&personIdJson)
	if err != nil {
		return
	}
	ID = personIdJson.ID
	return ID, nil
}

// Função para comparar a senha não criptografada com o hash da senha armazenada
func comparePasswordHash(passwordHash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
