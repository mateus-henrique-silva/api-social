package db

import (
	"context"

	"go.mod/src/connect"
	"go.mod/src/entity"
)

func CreateCategory(ctx context.Context, body entity.Category) (entity.Category, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return entity.Category{}, err
	}
	collection := client.Database("mydb").Collection("category")
	_, err = collection.InsertOne(ctx, body)
	return body, nil
}
