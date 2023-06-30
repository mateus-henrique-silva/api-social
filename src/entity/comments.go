package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comments struct {
	Id           primitive.ObjectID `bson:"_id"`
	IdPost       string             `bson:"idPost"`
	UserComments string             `bson:"userComments" json:"userComments"`
	CommentsText string             `bson:"commentsText"`
	CreatedAt    time.Time          `bson:"createdAt"`
}
