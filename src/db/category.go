package db

import (
	"context"
	"fmt"

	"go.mod/src/connect"
	"go.mod/src/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func UpdateCategory(ctx context.Context, id string, body entity.Category) error {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return err
	}
	resultId, _ := primitive.ObjectIDFromHex(id)
	collection := client.Database("mydb").Collection("category")
	filter := bson.M{"_id": resultId}
	update := bson.M{"$set": body}
	sendCollection, err := collection.UpdateOne(ctx, filter, update)
	fmt.Println(sendCollection)
	return err
}

func CheckIfCategoryExists(ctx context.Context, name string) (bool, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return false, err
	}

	collection := client.Database("mydb").Collection("category")

	filter := bson.M{"name": name}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
