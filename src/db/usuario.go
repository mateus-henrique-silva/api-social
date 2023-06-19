package db

import (
	"context"
	"fmt"

	"go.mod/src/connect"
	"go.mod/src/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(ctx context.Context, person entity.Usuario) (entity.Usuario, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return entity.Usuario{}, err
	}
	collection := client.Database("mydb").Collection("people")
	_, err = collection.InsertOne(ctx, person)
	return person, err
}

func FindUser(ctx context.Context) ([]entity.Usuario, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return nil, err
	}
	collection := client.Database("mydb").Collection("people")
	filter := bson.D{{}}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var request []entity.Usuario
	for cur.Next(context.TODO()) {
		var body entity.Usuario
		err := cur.Decode(&body)
		if err != nil {
			return nil, err
		}
		request = append(request, body)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return request, nil
}

func FindById(ctx context.Context, id string) (entity.Usuario, error) {
	var body entity.Usuario
	client, err := connect.ConfigDataBase()
	if err != nil {
		return body, err
	}
	collection := client.Database("mydb").Collection("people")
	resultId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": resultId}
	single := collection.FindOne(ctx, filter)
	err = single.Decode(&body)
	if err != nil {
		return body, err
	}
	return body, nil

}

func RemoveById(ctx context.Context, id string) error {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return err
	}
	collection := client.Database("mydb").Collection("people")
	resultId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": resultId}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func UpdateById(ctx context.Context, id string, person entity.Usuario) error {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return err
	}
	resultId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: resultId}} // Objeto de consulta (filtro)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: person.Name}}}}
	collection := client.Database("mydb").Collection("people")

	result, err := collection.UpdateOne(ctx, filter, update)
	fmt.Println(result)
	return err
}
