package db

import (
	"context"
	"fmt"

	"go.mod/src/connect"
	"go.mod/src/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser cria um usuario.
func CreateUser(ctx context.Context, person entity.Usuario) (entity.Usuario, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return entity.Usuario{}, err
	}
	collection := client.Database("mydb").Collection("people")
	_, err = collection.InsertOne(ctx, person)
	return person, nil
}

// FindUsers retorna todos os usuarios.
func FindUsers(ctx context.Context) ([]entity.Usuario, error) {
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

// FindById retorna usuarios a partir do id.
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

// RemoveById apaga o usuario a partir do id.
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

// UpdateById atualiza os dados do user a partir do id.
func UpdateById(ctx context.Context, id string, person entity.Usuario) error {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return err
	}
	resultId, _ := primitive.ObjectIDFromHex(id)
	// filter objeto de consulta (filtro)
	filter := bson.D{{Key: "_id", Value: resultId}}
	update := bson.D{{Key: "$set", Value: person.Name}}
	collection := client.Database("mydb").Collection("people")
	result, err := collection.UpdateOne(ctx, filter, update)
	fmt.Println(result)
	return err
}

func CheckIfUserExists(ctx context.Context, name string) (bool, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return false, err
	}

	collection := client.Database("mydb").Collection("people")

	filter := bson.M{"name": name}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func CheckIfUserEmailExists(ctx context.Context, email string) (bool, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return false, err
	}
	collection := client.Database("mydb").Collection("people")
	filter := bson.M{"email": email}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
