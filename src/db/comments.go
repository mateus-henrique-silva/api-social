package db

import (
	"context"

	"go.mod/src/connect"
	"go.mod/src/entity"
)

func CrateComment(ctx context.Context, body entity.Comments) error {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return err
	}
	collection := client.Database("mydb").Collection("comments")
	_, err = collection.InsertOne(ctx, body)
	if err != nil {
		return err
	}
	return nil
}
