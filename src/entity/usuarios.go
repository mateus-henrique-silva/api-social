package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Optional *string

type Usuario struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	UserImage   string             `bson:"userImage" json:"userImage"`
	Email       string             `bson:"email" json:"email"`
	Number      string             `bson:"number" json:"number"`
	Password    string             `bson:"password" json:"password"`
	City        Optional           `bson:"city" json:"city"`
	State       Optional           `bson:"state" json:"state"`
	Coutry      Optional           `bson:"coutry" json:"coutry"`
	Cep         Optional           `bson:"cep" json:"cep"`
	Role        string             `bson:"role" json:"role"`
	IsSubscribe bool               `bson:"isSubscribe" json:"isSubscribe"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}
