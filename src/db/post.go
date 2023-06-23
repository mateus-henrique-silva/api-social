package db

import (
	"context"

	"go.mod/src/connect"
	"go.mod/src/entity"
)

func CreatePost(ctx context.Context, body entity.Post) error {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return err
	}
	collection := client.Database("mydb").Collection("post")
	_, err = collection.InsertOne(ctx, body)
	if err != nil {
		return err
	}
	return nil
}
