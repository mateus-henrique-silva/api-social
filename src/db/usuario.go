package db

import (
	"context"

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
	// defer connect.CloseDB()
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
