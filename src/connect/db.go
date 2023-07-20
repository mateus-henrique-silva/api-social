package connect

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
)

// tfxY81frdh82wnia
func ConfigDataBase() (*mongo.Client, error) {
	ctx := context.TODO()
	Load()
	mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func CloseDB() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conexão com o MongoDB fechada")
}
