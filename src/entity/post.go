package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `bson:"id"`
	Title     string             `bson:"title"`
	Text      string             `bson:"text"`
	NameUser  string             `bson:"name"`
	Image     PostImageMultiples `bson:"image"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type PostImageMultiples struct {
	ImageOne   string `bson:"image_one"`
	ImageTwo   string `bson:"image_two"`
	ImageThree string `bson:"image_three"`
	ImageFour  string `bson:"image_four"`
	ImageFive  string `bson:"image_five"`
}
