package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	Title            string             `bson:"title" json:"title"`
	TitleSlug        string             `bson:"titleSlug" json:"titleSlug"`
	Text             string             `bson:"text" json:"text"`
	Article          string             `bson:"article" json:"article"`
	NameUser         string             `bson:"name" json:"nameUser"`
	AuthorImage      string             `bson:"authorImage" json:"authorImage"`
	Category         string             `bson:"category" json:"category"`
	CategorySlug     string             `bson:"categorySlug" json:"categorySlug"`
	Image            PostImageMultiples `bson:"image" json:"image"`
	BannerAltText    string             `bson:"bannerAltText" json:"bannerAltText"`
	CommentsQuantity uint64             `bson:"-" json:"-"`
	LinkYoutube      string             `bson:"linkYoutube" json:"linkYoutube"`
	Approved         bool               `bson:"approved" json:"approved"`
	Comments         []Comments         `bson:"comments" json:"comments"`
	Summary          string             `bson:"summary" json:"summary"`
	CreatedAt        time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt        time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type PostImageMultiples struct {
	ImageOne   string `bson:"imageOne" json:"imageOne"`
	ImageTwo   string `bson:"imageTwo" json:"imageTwo"`
	ImageThree string `bson:"imageThree" json:"imageThree"`
	ImageFour  string `bson:"imageFour" json:"imageFour"`
	ImageFive  string `bson:"imageFive" json:"imageFive"`
}

type PostReturnResponse struct {
	Category string `bson:"category" json:"category"`
	Post     []Post `bson:"post" json:"post"`
}

type PostCard struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Title         string             `bson:"title" json:"title"`
	TitleSlug     string             `bson:"titleSlug" json:"titleSlug"`
	Text          string             `bson:"text" json:"text"`
	NameUser      string             `bson:"name" json:"nameUser"`
	Category      string             `bson:"category" json:"category"`
	Image         PostImageMultiples `bson:"image" json:"image"`
	BannerAltText string             `bson:"bannerAltText" json:"bannerAltText"`
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`
}
