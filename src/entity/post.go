package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID               primitive.ObjectID `bson:"_id"`
	Title            string             `bson:"title"`
	TitleSlug        string             `bson:"titleSlug"`
	Text             string             `bson:"text"`
	NameUser         string             `bson:"name"`
	Image            PostImageMultiples `bson:"image"`
	BannerAltText    string             `bson:"bannerAltText"`
	CommentsQuantity uint64             `bson:"commentsQuantity"`
	CreatedAt        time.Time          `bson:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt"`
}

type PostImageMultiples struct {
	ImageOne   string `bson:"imageOne"`
	ImageTwo   string `bson:"imageTwo"`
	ImageThree string `bson:"imageThree"`
	ImageFour  string `bson:"imageFour"`
	ImageFive  string `bson:"imageFive"`
}
