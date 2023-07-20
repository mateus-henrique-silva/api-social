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
		postReturnResponse.Post[0] = post

		data = append(data, postReturnResponse)
	}

	if err := sendCollection.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func GetPost(ctx context.Context, name string) (entity.Post, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return entity.Post{}, err
	}

	db := client.Database("mydb")
	postCollection := db.Collection("post")
	commentCollection := db.Collection("comments")

	post := entity.Post{}
	err = postCollection.FindOne(ctx, bson.M{"titleSlug": name}).Decode(&post)
	if err != nil {
		return entity.Post{}, err
	}

	comments := []entity.Comments{}
	cursor, err := commentCollection.Find(ctx, bson.M{"userComments": name})
	if err != nil {
		return entity.Post{}, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		comment := entity.Comments{}
		err := cursor.Decode(&comment)
		if err != nil {
			return entity.Post{}, err
		}
		comments = append(comments, comment)
	}
	if err := cursor.Err(); err != nil {
		return entity.Post{}, err
	}

	post.Comments = comments
	return post, nil
}

func GetPostCard(ctx context.Context) ([]entity.PostReturnResponse, error) {
	client, err := connect.ConfigDataBase()
	if err != nil {
		return nil, err
	}

	db := client.Database("mydb")
	postCollection := db.Collection("post")

	// Encontrar todos os posts
	result, err := postCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// Mapa para armazenar os posts agrupados por categoria
	postMap := make(map[string][]entity.Post)
	for result.Next(ctx) {
		post := entity.Post{}
		err := result.Decode(&post)
		if err != nil {
			return nil, err
		}
		postMap[post.Category] = append(postMap[post.Category], post)
	}

	// Criar a resposta final com a estrutura desejada
	var response []entity.PostReturnResponse
	for category, posts := range postMap {
		response = append(response, entity.PostReturnResponse{
			Category: category,
			Post:     posts,
		})
	}

	return response, nil
}
