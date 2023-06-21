package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Number    string             `bson:"number"`
	Password  string             `bson:"password"`
	City      string             `bson:"city"`
	State     string             `bson:"state"`
	Coutry    string             `bson:"coutry"`
	Cep       string             `bson:"cep"`
	CreatedAt time.Time          `bson:"created"`
}
