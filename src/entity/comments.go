package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comments struct {
	Id           primitive.ObjectID `bson:"_id" json:"id"`
	IdPost       string             `bson:"idPost" json:"idPost"`
	UserComments string             `bson:"userComments" json:"userComments"`
	CommentsText string             `bson:"commentsText" json:"commentsText"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
}
