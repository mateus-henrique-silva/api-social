package db

import (
	"context"

	"go.mod/src/connect"
	"go.mod/src/entity"
)

func CreateUser(ctx context.Context, person entity.Usuario) (entity.Usuario, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return entity.Usuario{}, err
	}
	collection := client.Database("mydb").Collection("people")
	_, err = collection.InsertOne(ctx, person)
	// defer connect.CloseDB()
	return person, err
}
