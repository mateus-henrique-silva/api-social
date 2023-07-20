package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
	StatusA Status = "aproved"
	StatusB Status = "rejected"
)

type Category struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Name         string             `bson:"name" json:"name"`
	SlugCategory string             `bson:"slugCategory" json:"slugCategory"`
	IsAtHome     bool               `bson:"isAtHome" json:"isAtHome"`
	Status       string             `bson:"status" json:"status"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
}
