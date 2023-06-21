package db

import (
	"context"

	"go.mod/src/connect"
	"go.mod/src/entity"
	"go.mongodb.org/mongo-driver/bson"
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

func FindCategoryAll(ctx context.Context) ([]entity.Category, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return nil, err
	}
	collection := client.Database("mydb").Collection("category")
	filter := bson.D{{}}
	findCollection, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer findCollection.Close(ctx)
	var data []entity.Category

	for findCollection.Next(ctx) {
		var body entity.Category
		err := findCollection.Decode(&body)
		if err != nil {
			return nil, err
		}
		data = append(data, body)
	}

	return data, err
}
