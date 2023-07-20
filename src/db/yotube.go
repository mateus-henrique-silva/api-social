package db

import (
	"context"
	"fmt"

	"go.mod/src/connect"
	"go.mod/src/entity"
	"go.mongodb.org/mongo-driver/bson"
)

func ReturnAllVideos(ctx context.Context) ([]entity.Youtube, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return nil, err
	}
	collection := client.Database("mydb").Collection("youtube")
	sendCollection, err := collection.Find(ctx, bson.M{})
	var bodyArray []entity.Youtube
	for sendCollection.Next(ctx) {
		var body entity.Youtube
		if err := sendCollection.Decode(&body); err != nil {
			return nil, err
		}
		bodyArray = append(bodyArray, body)

	}

	return bodyArray, nil

}

func CreateYoutube(ctx context.Context, youtube entity.Youtube) error {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return err
	}
	collection := client.Database("mydb").Collection("youtube")
	sendCollection, err := collection.InsertOne(ctx, youtube)
	fmt.Println(sendCollection)
	return nil

}

func PutYotube(ctx context.Context, youtube entity.Youtube, Id string) error {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return nil
	}
	collection := client.Database("mydb").Collection("youtube")
	sendCollection, err := collection.UpdateOne(ctx, bson.M{"videID": Id}, youtube)
	fmt.Println(sendCollection)
	return nil

}
