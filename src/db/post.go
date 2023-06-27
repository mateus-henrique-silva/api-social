package db

import (
	"context"

	"go.mod/src/connect"
	"go.mod/src/entity"
	"go.mongodb.org/mongo-driver/bson"
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

func CheckIfPostNameExists(ctx context.Context, title string) (bool, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return false, err
	}

	collection := client.Database("mydb").Collection("post")

	filter := bson.M{"title": title}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
func GetIndexPost(ctx context.Context) ([]entity.PostReturnResponse, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return nil, err
	}

	collection := client.Database("mydb").Collection("post")
	sendCollection, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer sendCollection.Close(ctx)

	var data []entity.PostReturnResponse
	for sendCollection.Next(ctx) {
		var post entity.Post
		var postReturnResponse entity.PostReturnResponse

		err := sendCollection.Decode(&post)
		if err != nil {
			return nil, err
		}

		postReturnResponse.Category = post.Category
		postReturnResponse.Post = post

		data = append(data, postReturnResponse)
	}

	if err := sendCollection.Err(); err != nil {
		return nil, err
	}

	return data, nil
}