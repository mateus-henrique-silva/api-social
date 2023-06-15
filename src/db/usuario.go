package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mod/src/connect"
	"go.mod/src/entity"
)

func GetName(name string) (entity.Usuario, error) {
	ctx := context.TODO()
	client, err := connect.ConfigDataBase()
	if err != nil {
		return entity.Usuario{}, err
	}
	collection := client.Database("mydb").Collection("people")
	person := entity.Usuario{
		Id:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	_, err = collection.InsertOne(ctx, person)
	// defer connect.CloseDB()
	return person, err
}
